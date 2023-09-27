package dsp56k

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

//go:generate env GOOS=windows GOARCH=386 go run -exec wine generate.go

const EntrySize = 40
const Entries = 0x1000000

type Entry [EntrySize]byte

func (e Entry) Len() int        { return int(e[0]) }
func (e *Entry) String() string { return string(e[e[0]:]) }

type TestWriter struct {
	Options
	ProgramCounter uint32
	bytes.Buffer
}

func (w *TestWriter) Write(args ...Token) error {
	for _, arg := range args {
		_, err := arg.WriteOperand(&w.Buffer, w.Options, w.ProgramCounter)
		if err != nil {
			return err
		}
	}

	return nil
}

func cstring(b []byte) string {
	i := bytes.IndexByte(b, 0)
	if i == -1 {
		i = len(b)
	}
	return string(b[:i])
}

func TestSorted(t *testing.T) {
	sorted := sort.SliceIsSorted(decodeTable, func(i, j int) bool {
		if decodeTable[i].Constant == decodeTable[i].Constant {
			return decodeTable[i].Mask < decodeTable[i].Mask
		}

		return decodeTable[i].Constant > decodeTable[j].Constant
	})

	for _, row := range decodeTable {
		t.Logf("%024[1]b (%06[1]x) %024[2]b (%06[2]x)", row.Mask, row.Constant)
	}

	if !sorted {
		t.Fatal("decode table must be sorted from highest constant to lowest")
	}
}

func TestSortedMove(t *testing.T) {
	sorted := sort.SliceIsSorted(decodeMoveTable, func(i, j int) bool {
		return decodeMoveTable[i].Mask > decodeMoveTable[j].Mask
	})

	for _, row := range decodeMoveTable {
		t.Logf("%016[1]b (%04[1]x) %016[2]b (%04[2]x)", row.Mask, row.Constant)
	}

	if !sorted {
		t.Fatal("decode table must be sorted from highest constant to lowest")
	}
}

func TestSortedALU(t *testing.T) {
	sorted := sort.SliceIsSorted(decodeALUTable, func(i, j int) bool {
		return decodeALUTable[i].Mask > decodeALUTable[j].Mask
	})

	for _, row := range decodeALUTable {
		t.Logf("%06x %06x", row.Mask, row.Constant)
	}

	if !sorted {
		t.Fatal("decode table must be sorted from highest constant to lowest")
	}
}

func TestMultiMask(t *testing.T) {
	mask := []int{8, 1, 6, 1} // Treat bits 6 and 8 as being adjacent

	if multimask(0b101000000, mask...) != 0b11 {
		t.Fatal()
	}

	if multimask(0b100000000, mask...) != 0b10 {
		t.Fatal()
	}

	if multimask(0b001000000, mask...) != 1 {
		t.Fatal()
	}

	if multimask(0b010111111, mask...) != 0 {
		t.Fatal()
	}
}

type TestCase struct {
	Opcode uint32
	Words  int
	Asm    string
}

func testDisassemble(t *testing.T, ins Instruction, tests []TestCase) {
	t.Helper()

	name := "unknown"
	if ins != nil {
		rt := reflect.TypeOf(ins)
		name = rt.Elem().Name()
	}

	t.Run(name, func(t *testing.T) {
		t.Helper()

		if ins == nil {
			t.SkipNow()
		}

		if len(tests) == 0 {
			t.SkipNow()
		}

		for _, test := range tests {
			// Exit the test if it was failed
			if t.Failed() {
				t.FailNow()
			}

			decodedIns := Decode(test.Opcode, 0xdeface)

			if decodedIns == nil {
				if test.Words != 0 {
					// Don't fail now, we could still print some ASM that might be helpful
					t.Errorf("failed to decode valid opcode %06x %q", test.Opcode, test.Asm)
				}

				continue
			}

			// Test that the correct opcode is decoded
			if reflect.TypeOf(decodedIns) != reflect.TypeOf(ins) {
				t.Errorf("decoded incorrect instruction from opcode %06x want=%T got=%T", test.Opcode, ins, decodedIns)
			}

			// Test word count
			wordCount := 1
			if decodedIns.UsesExtensionWord() {
				wordCount = 2
			}

			if wordCount != test.Words {
				t.Errorf("decoded wrong number of words for opcode %06x %q want=%d got=%d", test.Opcode, test.Asm, test.Words, wordCount)
			}

			// Test assembly text
			w := &TestWriter{}
			if err := decodedIns.Disassemble(w); err != nil {
				t.Errorf("error while disassembling opcode %06x %q: %s", test.Opcode, test.Asm, err)
			}

			if got := w.String(); got != test.Asm {
				// Print ASM context for the error
				t.Errorf("asm for opcode %06x mismatch want=%q got=%q", test.Opcode, test.Asm, got)
			}
		}
	})
}

type SymbolTableMap map[Location]string

func (s SymbolTableMap) LookupSymbol(m Memory, addr uint32) (string, bool) {
	str, ok := s[Location{m, addr}]
	return str, ok
}

func TestSymbolsInDisassembly(t *testing.T) {
	type SymbolTestCase struct {
		Instruction Instruction
		ASM         string
		SymbolicASM string
	}

	symbolTable := SymbolTableMap{
		{MemoryX, 0xcafe}: "xmem",
		{MemoryY, 0xcafe}: "ymem",
		{MemoryY, 0xdead}: "ymem2",
		{MemoryP, 0xcafe}: "pmem",
		{MemoryP, 0xdead}: "pmem2",
		{MemoryL, 0xcafe}: "lmem",
	}

	// Pick some random value
	programCounter := 0x1234

	displacement := int32(0xcafe - programCounter)

	tests := []SymbolTestCase{
		{&Bcc_xxxx{
			Condition:    ConditionGE,
			Displacement: displacement,
		}, "bge >*+$b8ca", "bge >pmem"},
		{&Bcc_xxx{
			Condition:    ConditionGE,
			Displacement: displacement,
		}, "bge <*+$b8ca", "bge <pmem"},
		{&Bchg_ea_Abs{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bchg #$c,y:>$cafe", "bchg #$c,y:>ymem"},
		{&Bchg_aa{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bchg #$c,y:<$cafe", "bchg #$c,y:<ymem"},
		{&Bchg_pp{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bchg #$c,y:<<$cafe", "bchg #$c,y:<<ymem"},
		{&Bchg_qq{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bchg #$c,y:<<$cafe", "bchg #$c,y:<<ymem"},
		{&Bclr_ea_Abs{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bclr #$c,y:>$cafe", "bclr #$c,y:>ymem"},
		{&Bclr_aa{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bclr #$c,y:<$cafe", "bclr #$c,y:<ymem"},
		{&Bclr_pp{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bclr #$c,y:<<$cafe", "bclr #$c,y:<<ymem"},
		{&Bclr_qq{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bclr #$c,y:<<$cafe", "bclr #$c,y:<<ymem"},
		{&Bra_xxxx{
			Displacement: displacement,
		}, "bra >*+$b8ca", "bra >pmem"},
		{&Bra_xxx{
			Displacement: displacement,
		}, "bra <*+$b8ca", "bra <pmem"},
		{&Brclr_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, "brclr #$c,y:(r5)+,>*+$b8ca", "brclr #$c,y:(r5)+,>pmem"},
		{&Brclr_aa{
			Address:      0xdead,
			Memory:       MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, "brclr #$c,y:<$dead,>*+$b8ca", "brclr #$c,y:<ymem2,>pmem"},
		{&Brclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "brclr #$c,y:<<$dead,>*+$b8ca", "brclr #$c,y:<<ymem2,>pmem"},
		{&Brclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "brclr #$c,y:<<$dead,>*+$b8ca", "brclr #$c,y:<<ymem2,>pmem"},
		{&Brclr_S{
			Source:       RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, "brclr #$c,b,>*+$b8ca", "brclr #$c,b,>pmem"},
		{&Brset_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, "brset #$c,y:(r5)+,>*+$b8ca", "brset #$c,y:(r5)+,>pmem"},
		{&Brset_aa{
			Address:      0xdead,
			Memory:       MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, "brset #$c,y:<$dead,>*+$b8ca", "brset #$c,y:<ymem2,>pmem"},
		{&Brset_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "brset #$c,y:<<$dead,>*+$b8ca", "brset #$c,y:<<ymem2,>pmem"},
		{&Brset_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "brset #$c,y:<<$dead,>*+$b8ca", "brset #$c,y:<<ymem2,>pmem"},
		{&Brset_S{
			Source:       RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, "brset #$c,b,>*+$b8ca", "brset #$c,b,>pmem"},
		{&BScc_xxxx{
			Condition:    ConditionGE,
			Displacement: displacement,
		}, "bsge >*+$b8ca", "bsge >pmem"},
		{&BScc_xxx{
			Condition:    ConditionGE,
			Displacement: displacement,
		}, "bsge <*+$b8ca", "bsge <pmem"},
		{&Bsclr_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, "bsclr #$c,y:(r5)+,>*+$b8ca", "bsclr #$c,y:(r5)+,>pmem"},
		{&Bsclr_aa{
			Address:      0xdead,
			Memory:       MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, "bsclr #$c,y:<$dead,>*+$b8ca", "bsclr #$c,y:<ymem2,>pmem"},
		{&Bsclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "bsclr #$c,y:<<$dead,>*+$b8ca", "bsclr #$c,y:<<ymem2,>pmem"},
		{&Bsclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "bsclr #$c,y:<<$dead,>*+$b8ca", "bsclr #$c,y:<<ymem2,>pmem"},
		{&Bsclr_S{
			Source:       RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, "bsclr #$c,b,>*+$b8ca", "bsclr #$c,b,>pmem"},
		{&Bset_ea_Abs{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bset #$c,y:>$cafe", "bset #$c,y:>ymem"},
		{&Bset_aa{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "bset #$c,y:<$cafe", "bset #$c,y:<ymem"},
		{&Bset_pp{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bset #$c,y:<<$cafe", "bset #$c,y:<<ymem"},
		{&Bset_qq{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "bset #$c,y:<<$cafe", "bset #$c,y:<<ymem"},
		{&Bsr_xxxx{
			Displacement: displacement,
		}, "bsr >*+$b8ca", "bsr >pmem"},
		{&Bsr_xxx{
			Displacement: displacement,
		}, "bsr <*+$b8ca", "bsr <pmem"},
		{&Bsset_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, "bsset #$c,y:(r5)+,>*+$b8ca", "bsset #$c,y:(r5)+,>pmem"},
		{&Bsset_aa{
			Address:      0xdead,
			Memory:       MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, "bsset #$c,y:<$dead,>*+$b8ca", "bsset #$c,y:<ymem2,>pmem"},
		{&Bsset_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "bsset #$c,y:<<$dead,>*+$b8ca", "bsset #$c,y:<<ymem2,>pmem"},
		{&Bsset_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, "bsset #$c,y:<<$dead,>*+$b8ca", "bsset #$c,y:<<ymem2,>pmem"},
		{&Bsset_S{
			Source:       RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, "bsset #$c,b,>*+$b8ca", "bsset #$c,b,>pmem"},
		{&Btst_ea_Abs{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "btst #$c,y:>$cafe", "btst #$c,y:>ymem"},
		{&Btst_aa{
			Address:   0xcafe,
			Memory:    MemoryY,
			BitNumber: 12,
		}, "btst #$c,y:<$cafe", "btst #$c,y:<ymem"},
		{&Btst_pp{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "btst #$c,y:<<$cafe", "btst #$c,y:<<ymem"},
		{&Btst_qq{
			PeripheralAddress: 0xcafe,
			Memory:            MemoryY,
			BitNumber:         12,
		}, "btst #$c,y:<<$cafe", "btst #$c,y:<<ymem"},
		{&Do_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			LoopAddress:      0xcafd,
		}, "do y:(r5)+,$cafe", "do y:(r5)+,pmem"},
		{&Do_aa{
			Address:     0xdead,
			Memory:      MemoryY,
			LoopAddress: 0xcafd,
		}, "do y:<$dead,$cafe", "do y:<ymem2,pmem"},
		{&Do_xxx{
			Immediate:   0x1234,
			LoopAddress: 0xcafd,
		}, "do #<$1234,$cafe", "do #<$1234,pmem"},
		{&Do_S{
			Source:      RegisterB,
			LoopAddress: 0xcafd,
		}, "do b,$cafe", "do b,pmem"},
		{&DoForever{
			LoopAddress: 0xcafd,
		}, "do forever,$cafe", "do forever,pmem"},
		{&Dor_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			LoopDisplacement: displacement - 1,
		}, "dor y:(r5)+,>*+$b8ca", "dor y:(r5)+,>pmem"},
		{&Dor_aa{
			Address:          0xdead,
			Memory:           MemoryY,
			LoopDisplacement: displacement - 1,
		}, "dor y:<$dead,>*+$b8ca", "dor y:<ymem2,>pmem"},
		{&Dor_xxx{
			Immediate:        0x1234,
			LoopDisplacement: displacement - 1,
		}, "dor #<$1234,>*+$b8ca", "dor #<$1234,>pmem"},
		{&Dor_S{
			Source:           RegisterB,
			LoopDisplacement: displacement - 1,
		}, "dor b,>*+$b8ca", "dor b,>pmem"},
		{&DorForever{
			LoopDisplacement: displacement - 1,
		}, "dor forever,>*+$b8ca", "dor forever,>pmem"},
		{&Jcc_xxx{
			Condition:   ConditionGE,
			JumpAddress: 0xcafe,
		}, "jge <$cafe", "jge <pmem"},
		{&Jcc_ea_Abs{
			JumpAddress: 0xcafe,
			Condition:   ConditionGE,
		}, "jge >$cafe", "jge >pmem"},
		{&Jclr_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, "jclr #$c,y:(r5)+,$cafe", "jclr #$c,y:(r5)+,pmem"},
		{&Jclr_aa{
			Address:     0xdead,
			Memory:      MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jclr #$c,y:<$dead,$cafe", "jclr #$c,y:<ymem2,pmem"},
		{&Jclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jclr #$c,y:<<$dead,$cafe", "jclr #$c,y:<<ymem2,pmem"},
		{&Jclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jclr #$c,y:<<$dead,$cafe", "jclr #$c,y:<<ymem2,pmem"},
		{&Jclr_S{
			Source:      RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jclr #$c,b,$cafe", "jclr #$c,b,pmem"},
		{&Jmp_ea_Abs{
			JumpAddress: 0xcafe,
		}, "jmp >$cafe", "jmp >pmem"},
		{&Jmp_xxx{
			JumpAddress: 0xcafe,
		}, "jmp <$cafe", "jmp <pmem"},
		{&Jscc_xxx{
			Condition:   ConditionGE,
			JumpAddress: 0xcafe,
		}, "jsge <$cafe", "jsge <pmem"},
		{&Jscc_ea_Abs{
			JumpAddress: 0xcafe,
			Condition:   ConditionGE,
		}, "jsge >$cafe", "jsge >pmem"},
		{&Jsclr_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, "jsclr #$c,y:(r5)+,$cafe", "jsclr #$c,y:(r5)+,pmem"},
		{&Jsclr_aa{
			Address:     0xdead,
			Memory:      MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jsclr #$c,y:<$dead,$cafe", "jsclr #$c,y:<ymem2,pmem"},
		{&Jsclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jsclr #$c,y:<<$dead,$cafe", "jsclr #$c,y:<<ymem2,pmem"},
		{&Jsclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jsclr #$c,y:<<$dead,$cafe", "jsclr #$c,y:<<ymem2,pmem"},
		{&Jsclr_S{
			Source:      RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jsclr #$c,b,$cafe", "jsclr #$c,b,pmem"},
		{&Jset_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, "jset #$c,y:(r5)+,$cafe", "jset #$c,y:(r5)+,pmem"},
		{&Jset_aa{
			Address:     0xdead,
			Memory:      MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jset #$c,y:<$dead,$cafe", "jset #$c,y:<ymem2,pmem"},
		{&Jset_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jset #$c,y:<<$dead,$cafe", "jset #$c,y:<<ymem2,pmem"},
		{&Jset_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jset #$c,y:<<$dead,$cafe", "jset #$c,y:<<ymem2,pmem"},
		{&Jset_S{
			Source:      RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jset #$c,b,$cafe", "jset #$c,b,pmem"},
		{&Jsr_ea_Abs{
			JumpAddress: 0xcafe,
		}, "jsr >$cafe", "jsr >pmem"},
		{&Jsr_xxx{
			JumpAddress: 0xcafe,
		}, "jsr <$cafe", "jsr <pmem"},
		{&Jsset_ea{
			EffectiveAddress: EffectiveAddress{AddressModePostIncrement, 5},
			Memory:           MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, "jsset #$c,y:(r5)+,$cafe", "jsset #$c,y:(r5)+,pmem"},
		{&Jsset_aa{
			Address:     0xdead,
			Memory:      MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jsset #$c,y:<$dead,$cafe", "jsset #$c,y:<ymem2,pmem"},
		{&Jsset_pp{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jsset #$c,y:<<$dead,$cafe", "jsset #$c,y:<<ymem2,pmem"},
		{&Jsset_qq{
			PeripheralAddress: 0xdead,
			Memory:            MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, "jsset #$c,y:<<$dead,$cafe", "jsset #$c,y:<<ymem2,pmem"},
		{&Jsset_S{
			Source:      RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, "jsset #$c,b,$cafe", "jsset #$c,b,pmem"},
		{&Lra_xxxx{
			DestinationAddress: RegisterB,
			Displacement:       displacement,
		}, "lra >*+$b8ca,b", "lra >pmem,b"},
		{&Movec_ea_Abs{
			IsWrite:           false,
			Memory:            MemoryY,
			Address:           0xcafe,
			ProgramController: RegisterB,
		}, "move b,y:>$cafe", "move b,y:>ymem"},
		{&Movec_aa{
			IsWrite:           false,
			Address:           0xcafe,
			Memory:            MemoryY,
			ProgramController: RegisterB,
		}, "move b,y:<$cafe", "move b,y:<ymem"},
		{&Movem_ea_Abs{
			IsWrite:             false,
			Address:             0xcafe,
			SourceOrDestination: RegisterB,
		}, "move b,p:>$cafe", "move b,p:>pmem"},
		{&Movem_aa{
			IsWrite:             false,
			Address:             0xcafe,
			SourceOrDestination: RegisterB,
		}, "move b,p:<$cafe", "move b,p:<pmem"},
		{&Movep_ppea{
			PeripheralMemory:  MemoryY,
			IsWrite:           false,
			EffectiveAddress:  EffectiveAddress{AddressModePostIncrement, 5},
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,y:(r5)+", "movep y:<<ymem,y:(r5)+"},
		{&Movep_ppea_Abs{
			PeripheralMemory:  MemoryY,
			IsWrite:           false,
			Address:           0xdead,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,y:>$dead", "movep y:<<ymem,y:>ymem2"},
		{&Movep_ppea_Imm{
			PeripheralMemory:  MemoryY,
			Immediate:         0x1234,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep #>$1234,y:<<$cafe", "movep #>$1234,y:<<ymem"},
		{&Movep_Xqqea{
			IsWrite:           false,
			EffectiveAddress:  EffectiveAddress{AddressModePostIncrement, 5},
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep x:<<$cafe,y:(r5)+", "movep x:<<xmem,y:(r5)+"},
		{&Movep_Xqqea_Abs{
			IsWrite:           false,
			Address:           0xdead,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep x:<<$cafe,y:>$dead", "movep x:<<xmem,y:>ymem2"},
		{&Movep_Xqqea_Imm{
			Immediate:         0x1234,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep #>$1234,x:<<$cafe", "movep #>$1234,x:<<xmem"},
		{&Movep_Yqqea{
			IsWrite:           false,
			EffectiveAddress:  EffectiveAddress{AddressModePostIncrement, 5},
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,y:(r5)+", "movep y:<<ymem,y:(r5)+"},
		{&Movep_Yqqea_Abs{
			IsWrite:           false,
			Address:           0xdead,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,y:>$dead", "movep y:<<ymem,y:>ymem2"},
		{&Movep_Yqqea_Imm{
			Immediate:         0x1234,
			Memory:            MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep #>$1234,y:<<$cafe", "movep #>$1234,y:<<ymem"},
		{&Movep_eapp{
			PeripheralMemory:  MemoryY,
			IsWrite:           false,
			EffectiveAddress:  EffectiveAddress{AddressModePostIncrement, 5},
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,p:(r5)+", "movep y:<<ymem,p:(r5)+"},
		{&Movep_eapp_Abs{
			PeripheralMemory:  MemoryY,
			IsWrite:           false,
			Address:           0xdead,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,p:>$dead", "movep y:<<ymem,p:>pmem2"},
		{&Movep_eapp_Imm{
			PeripheralMemory:  MemoryY,
			Immediate:         0x1234,
			PeripheralAddress: 0xcafe,
		}, "movep #>$1234,y:<<$cafe", "movep #>$1234,y:<<ymem"},
		{&Movep_eaqq{
			IsWrite:           false,
			EffectiveAddress:  EffectiveAddress{AddressModePostIncrement, 5},
			PeripheralMemory:  MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,p:(r5)+", "movep y:<<ymem,p:(r5)+"},
		{&Movep_eaqq_Abs{
			IsWrite:           false,
			Address:           0xdead,
			PeripheralMemory:  MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep y:<<$cafe,p:>$dead", "movep y:<<ymem,p:>pmem2"},
		{&Movep_eaqq_Imm{
			Immediate:         0x1234,
			PeripheralMemory:  MemoryY,
			PeripheralAddress: 0xcafe,
		}, "movep #>$1234,y:<<$cafe", "movep #>$1234,y:<<ymem"},
		{&Movep_Spp{
			PeripheralMemory:    MemoryY,
			IsWrite:             false,
			SourceOrDestination: RegisterB,
			PeripheralAddress:   0xcafe,
		}, "movep y:<<$cafe,b", "movep y:<<ymem,b"},
		{&Movep_SXqq{
			IsWrite:             false,
			SourceOrDestination: RegisterB,
			PeripheralAddress:   0xcafe,
		}, "movep x:<<$cafe,b", "movep x:<<xmem,b"},
		{&Movep_SYqq{
			IsWrite:             false,
			SourceOrDestination: RegisterB,
			PeripheralAddress:   0xcafe,
		}, "movep y:<<$cafe,b", "movep y:<<ymem,b"},
		{&Plock_Abs{
			Address: 0xcafe,
		}, "plock $cafe", "plock pmem"},
		{&Plockr{
			Displacement: displacement,
		}, "plockr >*+$b8ca", "plockr >pmem"},
		{&Punlock_Abs{
			Address: 0xcafe,
		}, "punlock $cafe", "punlock pmem"},
		{&Punlockr{
			Displacement: displacement,
		}, "punlockr >*+$b8ca", "punlockr >pmem"},
		{&Rep_aa{
			Address: 0xcafe,
			Memory:  MemoryY,
		}, "rep y:<$cafe", "rep y:<ymem"},
		{&Vsl_Abs{
			Source:  RegisterB,
			Address: 0xcafe,
			Bit:     12,
		}, "vsl b,1,l:$cafe", "vsl b,1,l:lmem"},

		// Parallel Instructions
		{&ParallelInstruction{&Movex_aa{
			SourceOrDestination: RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &Max{}}, "max a,b a,x:<$cafe", "max a,b a,x:<xmem"},
		{&ParallelInstruction{&Movey_aa{
			SourceOrDestination: RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &Max{}}, "max a,b a,y:<$cafe", "max a,b a,y:<ymem"},
		{&ParallelInstruction{&Movex_ea_Abs{
			SourceOrDestination: RegisterA,
			IsWrite:             false,
			Address:             0xcafe,
		}, &Max{}}, "max a,b a,x:>$cafe", "max a,b a,x:>xmem"},
		{&ParallelInstruction{&Movey_ea_Abs{
			SourceOrDestination: RegisterA,
			IsWrite:             false,
			Address:             0xcafe,
		}, &Max{}}, "max a,b a,y:>$cafe", "max a,b a,y:>ymem"},
		{&ParallelInstruction{&Movexr_ea_Abs{
			Address:             0xcafe,
			IsWrite:             false,
			Source:              RegisterX,
			Destination:         RegisterY,
			SourceOrDestination: RegisterA,
		}, &Max{}}, "max a,b a,x:>$cafe x,y", "max a,b a,x:>xmem x,y"},
		{&ParallelInstruction{&Moveyr_ea_Abs{
			Address:             0xcafe,
			IsWrite:             false,
			Source:              RegisterX,
			Destination:         RegisterY,
			SourceOrDestination: RegisterA,
		}, &Max{}}, "max a,b x,y a,y:>$cafe", "max a,b x,y a,y:>ymem"},
		{&ParallelInstruction{&Movel_aa{
			Address:             0xcafe,
			SourceOrDestination: RegisterA,
			IsWrite:             false,
		}, &Max{}}, "max a,b a,l:<$cafe", "max a,b a,l:<lmem"},
		{&ParallelInstruction{&Movel_ea_Abs{
			SourceOrDestination: RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &Max{}}, "max a,b a,l:>$cafe", "max a,b a,l:>lmem"},
	}

	disassemble := func(ins Instruction, pc uint32, opt Options) string {
		w := &TestWriter{Options: opt, ProgramCounter: pc}
		if err := ins.Disassemble(w); err != nil {
			t.Fatalf("error while disassembling %#v: %s", ins, err)
		}

		return w.String()
	}

	options := Options{}
	optionsWithSymbols := Options{symbolTable}
	for _, test := range tests {
		if got := disassemble(test.Instruction, uint32(programCounter), options); got != test.ASM {
			t.Errorf("asm mismatch for %#v want=%q got=%q", test.Instruction, test.ASM, got)
		}

		if got := disassemble(test.Instruction, uint32(programCounter), optionsWithSymbols); got != test.SymbolicASM {
			t.Fatalf("symbolic asm mismatch for %#v want=%q got=%q", test.Instruction, test.SymbolicASM, got)
		}
	}
}

func TestDisassembleEveryInstruction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test")
	}

	f, err := os.Open("opcodes.bin")
	if err != nil {
		t.Fatal(err, "did you run go generate?")
	}

	defer f.Close()

	entries := make([]byte, EntrySize*Entries)
	if _, err := io.ReadFull(f, entries); err != nil {
		t.Fatal(err)
	}

	chunkSize := Entries / 4
	for n := 0; n < Entries; n += chunkSize {
		start := n
		end := start + chunkSize
		if end > Entries {
			end = Entries
		}

		t.Run(fmt.Sprintf("opcodes-%06x-%06x", n, end-1), func(t *testing.T) {
			t.Parallel()

			for i := start; i < end; i++ {
				entry := entries[i*EntrySize : i*EntrySize+EntrySize]
				words := int(entry[0])
				if words > 2 {
					t.Fail()
				}

				opcode := uint32(i)
				decodedIns := Decode(opcode, 0xdeface)
				asm := cstring(entry[1:])

				if decodedIns == nil {
					if words != 0 {
						t.Fatalf("failed to decode valid opcode %06x %q", opcode, asm)
					}

					continue
				}

				// Test word count
				wordCount := 1
				if decodedIns.UsesExtensionWord() {
					wordCount = 2
				}

				if wordCount != words {
					t.Errorf("decoded wrong number of words for opcode %06x %q want=%d got=%d (%T)", opcode, asm, words, wordCount, decodedIns)
				}

				// Test assembly text
				w := &TestWriter{}
				if err := decodedIns.Disassemble(w); err != nil {
					t.Fatalf("error while disassembling opcode %06x %q: %s (%T)", opcode, asm, err, decodedIns)
				}

				// Special cases where the Motorolla Disassembler outputs
				// ambiguous or incorrect disassembly
				strictMatch := true

				// Movex_Rnxxx
				if opcode&0b111111100000000010100000 == 0b000000100000000010000000 {
					strictMatch = false
				}

				// Movey_Rnxxx
				if opcode&0b111111100000000010100000 == 0b000000100000000010100000 {
					strictMatch = false
				}

				// Movex_Rnxxxx
				if opcode&0b111111111111100010000000 == 0b000010100111000010000000 {
					strictMatch = false
				}

				// Movey_Rnxxxx
				if opcode&0b111111111111100010000000 == 0b000010110111000010000000 {
					strictMatch = false
				}

				if got := w.String(); got != asm {
					if strictMatch {
						t.Fatalf("asm for opcode %06x mismatch want=%q got=%q", opcode, asm, got)
					} else {
						replacer := strings.NewReplacer(
							">", "",
							"<", "",
						)

						if got := replacer.Replace(got); got != asm {
							t.Fatalf("asm for opcode %06x mismatch want=%q got=%q", opcode, asm, got)
						}
					}
				}
			}
		})
	}
}
