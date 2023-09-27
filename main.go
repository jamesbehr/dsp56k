package dsp56k

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
)

type MultiplyMode int

const (
	MultiplyModeInvalid MultiplyMode = iota
	MultiplyModeSigned
	MultiplyModeUnsigned
	MultiplyModeMixed
)

func (m MultiplyMode) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	switch m {
	case MultiplyModeUnsigned:
		return fmt.Fprint(w, "uu")
	case MultiplyModeSigned:
		return fmt.Fprint(w, "ss")
	case MultiplyModeMixed:
		return fmt.Fprint(w, "su")
	}

	return 0, errors.New("invalid multiply mode")
}

func multiplyMode2Bit(value uint32) MultiplyMode {
	switch value {
	case 0b00:
		return MultiplyModeSigned
	case 0b10:
		return MultiplyModeMixed
	case 0b11:
		return MultiplyModeUnsigned
	}

	// This is undocumented, but it is how the assembler behaves
	return MultiplyModeUnsigned
}

func multiplyMode1Bit(value uint32) MultiplyMode {
	switch value {
	case 0b0:
		return MultiplyModeMixed
	case 0b1:
		return MultiplyModeUnsigned
	}

	// This is undocumented, but it is how the assembler behaves
	return MultiplyModeUnsigned
}

func boolean(value uint32) bool {
	if value == 0 {
		return false
	}

	return true
}

type RegisterPair struct {
	First, Second Register
}

// Table A-30 on Page 243
func multiplyAllPairs(value uint32) RegisterPair {
	switch value {
	case 0b0000:
		return RegisterPair{RegisterX0, RegisterX0}
	case 0b0100:
		return RegisterPair{RegisterX0, RegisterY1}
	case 0b0001:
		return RegisterPair{RegisterY0, RegisterY0}
	case 0b0101:
		return RegisterPair{RegisterY0, RegisterX0}
	case 0b0010:
		return RegisterPair{RegisterX1, RegisterX0}
	case 0b0110:
		return RegisterPair{RegisterX1, RegisterY0}
	case 0b0011:
		return RegisterPair{RegisterY1, RegisterY0}
	case 0b0111:
		return RegisterPair{RegisterY1, RegisterX1}
	case 0b1000:
		return RegisterPair{RegisterX1, RegisterX1}
	case 0b1100:
		return RegisterPair{RegisterY1, RegisterX0}
	case 0b1001:
		return RegisterPair{RegisterY1, RegisterY1}
	case 0b1101:
		return RegisterPair{RegisterX0, RegisterY0}
	case 0b1010:
		return RegisterPair{RegisterX0, RegisterX1}
	case 0b1110:
		return RegisterPair{RegisterY0, RegisterX1}
	case 0b1011:
		return RegisterPair{RegisterY0, RegisterY1}
	case 0b1111:
		return RegisterPair{RegisterX1, RegisterY1}
	}

	panic("unsupported")
}

// Table A-26 on Page 242
func multiplyPairs(value uint32) RegisterPair {
	switch value {
	case 0b000:
		return RegisterPair{RegisterX0, RegisterX0}
	case 0b001:
		return RegisterPair{RegisterY0, RegisterY0}
	case 0b010:
		return RegisterPair{RegisterX1, RegisterX0}
	case 0b011:
		return RegisterPair{RegisterY1, RegisterY0}
	case 0b100:
		return RegisterPair{RegisterX0, RegisterY1}
	case 0b101:
		return RegisterPair{RegisterY0, RegisterX0}
	case 0b110:
		return RegisterPair{RegisterX1, RegisterY0}
	case 0b111:
		return RegisterPair{RegisterY1, RegisterX1}
	}

	panic("unsupported")
}

// Table A-38 on Page 246
func addressMode2Bit(mm, rrr uint32) EffectiveAddress {
	switch mm {
	case 0b00:
		return EffectiveAddress{AddressModeNoUpdate, rrr}
	case 0b01:
		return EffectiveAddress{AddressModePostIncrementOffset, rrr}
	case 0b10:
		return EffectiveAddress{AddressModePostDecrement, rrr}
	case 0b11:
		return EffectiveAddress{AddressModePostIncrement, rrr}
	}

	panic("not supported")
}

// Table A-16, A-18, A-19 and A-20 on Pages 239-240
func addressMode(mmm, rrr uint32) EffectiveAddress {
	switch mmm {
	case 0b000:
		return EffectiveAddress{AddressModePostDecrementOffset, rrr}
	case 0b001:
		return EffectiveAddress{AddressModePostIncrementOffset, rrr}
	case 0b010:
		return EffectiveAddress{AddressModePostDecrement, rrr}
	case 0b011:
		return EffectiveAddress{AddressModePostIncrement, rrr}
	case 0b100:
		return EffectiveAddress{AddressModeNoUpdate, rrr}
	case 0b101:
		return EffectiveAddress{AddressModeIndexed, rrr}
	case 0b111:
		return EffectiveAddress{AddressModePreDecrement, rrr}
	}

	return EffectiveAddress{AddressModeInvalid, 0}
}

type Memory int

const (
	MemoryInvalid Memory = iota
	MemoryP
	MemoryX
	MemoryY
	MemoryL
)

func (m Memory) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	switch m {
	case MemoryP:
		return fmt.Fprint(w, "p:")
	case MemoryX:
		return fmt.Fprint(w, "x:")
	case MemoryY:
		return fmt.Fprint(w, "y:")
	case MemoryL:
		return fmt.Fprint(w, "l:")
	}

	return 0, errors.New("invalid memory")
}

type Instruction interface {
	Decode(opcode, extensionWord uint32) bool
	UsesExtensionWord() bool
	Disassemble(w TokenWriter) error
}

type Decoder interface {
	Decode(value uint32) bool
}

func highPeripheral(value uint32) uint32 {
	return 0xffffc0 + value
}

func lowPeripheral(value uint32) uint32 {
	return 0xffff80 + value
}

func xyspace(value uint32) Memory {
	switch value {
	case 0:
		return MemoryX
	case 1:
		return MemoryY
	}
	return MemoryInvalid
}

// sign-extend
func signExtend(width int, x uint32) int32 {
	shift := 32 - width
	return int32(x<<shift) >> shift
}

// shift and mask
func mask(word uint32, shift, width int) uint32 {
	mask := uint32(1<<width - 1)
	return (word >> shift) & mask
}

// extract single bit
func bit(word uint32, shift int) uint32 {
	return (word >> shift) & 1
}

func multimask(word uint32, maskPairs ...int) uint32 {
	if l := len(maskPairs); l < 4 || l%2 > 0 {
		panic("multimask requires two or more pairs of values")
	}

	value := uint32(0)

	for i := 0; i < len(maskPairs)/2; i++ {
		shift, width := maskPairs[i*2], maskPairs[i*2+1]
		mask := uint32(1<<width - 1)
		field := (word >> shift) & mask
		value = (value << width) | field
	}

	return value
}

type decodeRow struct {
	Mask, Constant uint32
	Builder        func() Instruction
}

var decodeTable []decodeRow

func registerInstruction(mask, constant uint32, builder func() Instruction) {
	// Check that the constant fits the mask
	if constant&mask != constant {
		panic("constant may only have bits set where bits in the mask are set")
	}

	// Insert into a sorted slice of decode rows
	i := sort.Search(len(decodeTable), func(i int) bool {
		// Use the mask as a tiebreaker
		if constant == decodeTable[i].Constant {
			return mask < decodeTable[i].Mask
		}

		return constant > decodeTable[i].Constant
	})

	decodeTable = append(decodeTable, decodeRow{})
	copy(decodeTable[i+1:], decodeTable[i:])
	decodeTable[i] = decodeRow{mask, constant, builder}
}

// Decode an instruction from the provided opcode word and the optional
// extension word. It will return nil for invalid instructions.
func Decode(opcode, extensionWord uint32) Instruction {
	i := sort.Search(len(decodeTable), func(i int) bool {
		return opcode >= decodeTable[i].Constant
	})

	for _, row := range decodeTable[i:] {
		if opcode < row.Constant {
			continue
		}

		if opcode&row.Mask == row.Constant {
			ins := row.Builder()
			if ins.Decode(opcode, extensionWord) {
				return ins
			}

			// Decode failed
			return nil
		}
	}

	// Matched nothing
	return nil
}

type tokenWriter struct {
	Options        Options
	Buffer         *bytes.Buffer
	ProgramCounter uint32
}

func (tw *tokenWriter) Write(tokens ...Token) error {
	for _, token := range tokens {
		if _, err := token.WriteOperand(tw.Buffer, tw.Options, tw.ProgramCounter); err != nil {
			return err
		}
	}

	return nil
}

func Disassemble(ins Instruction, symbols SymbolTable, pc uint32) (string, error) {
	tw := tokenWriter{
		Buffer:         &bytes.Buffer{},
		ProgramCounter: pc,
		Options: Options{
			SymbolTable: symbols,
		},
	}

	if err := ins.Disassemble(&tw); err != nil {
		return "", err
	}

	return tw.Buffer.String(), nil
}

// Table A-22 on Page 241
func programControlRegister8(value uint32) Register {
	switch value {
	case 0b000:
		return RegisterSZ
	case 0b001:
		return RegisterSR
	case 0b010:
		return RegisterOMR
	case 0b011:
		return RegisterSP
	case 0b100:
		return RegisterSSH
	case 0b101:
		return RegisterSSL
	case 0b110:
		return RegisterLA
	case 0b111:
		return RegisterLC
	}

	return RegisterInvalid
}

// Table A-22 on Page 241
func programControlRegister2(value uint32) Register {
	switch value {
	case 0:
		return RegisterVBA
	case 1:
		return RegisterSC
	}

	return RegisterInvalid
}

// Table A-22 on Page 241
func addressRegisterInAGU(value uint32) Register {
	switch value {
	case 0b010:
		return RegisterEP
	}

	return RegisterInvalid
}

func addressAndOffsetRegister(value uint32) Register {
	if value&0b1000 == 0 {
		return addressRegister(value & 0b111)
	}

	return addressOffsetRegister(value & 0b111)
}

func register5BitUndocumented(value uint32) Register {
	switch value {
	case 0b00000: // undocumented
		return RegisterX0
	case 0b00001: // undocumented
		return RegisterX1
	case 0b00010: // undocumented
		return RegisterY0
	case 0b00011: // undocumented
		return RegisterY1
	}

	return register5Bit(value)
}

// Table A-31 on Page 243
func register5Bit(value uint32) Register {
	switch value {
	case 0b00100:
		return RegisterX0
	case 0b00101:
		return RegisterX1
	case 0b00110:
		return RegisterY0
	case 0b00111:
		return RegisterY1
	case 0b01000:
		return RegisterA0
	case 0b01001:
		return RegisterB0
	case 0b01010:
		return RegisterA2
	case 0b01011:
		return RegisterB2
	case 0b01100:
		return RegisterA1
	case 0b01101:
		return RegisterB1
	case 0b01110:
		return RegisterA
	case 0b01111:
		return RegisterB
	case 0b10000, 0b10001, 0b10010, 0b10011, 0b10100, 0b10101, 0b10110, 0b10111:
		return addressRegister(value & 0b111)
	case 0b11000, 0b11001, 0b11010, 0b11011, 0b11100, 0b11101, 0b11110, 0b11111:
		return addressOffsetRegister(value & 0b111)
	}

	return RegisterInvalid
}

// Table A-22 on Page 241
func register6Bit(value uint32) Register {
	switch (value >> 3) & 7 {
	case 0b000:
		switch value & 7 {
		case 0b100:
			return RegisterX0
		case 0b101:
			return RegisterX1
		case 0b110:
			return RegisterY0
		case 0b111:
			return RegisterY1
		}
	case 0b001:
		switch value & 7 {
		case 0b000:
			return RegisterA0
		case 0b001:
			return RegisterB0
		case 0b010:
			return RegisterA2
		case 0b011:
			return RegisterB2
		case 0b100:
			return RegisterA1
		case 0b101:
			return RegisterB1
		case 0b110:
			return RegisterA
		case 0b111:
			return RegisterB
		}
	case 0b010:
		return addressRegister(value & 7)
	case 0b011:
		return addressOffsetRegister(value & 7)
	case 0b100:
		return addressModifierRegister(value & 7)
	case 0b101:
		return addressRegisterInAGU(value & 7)
	case 0b110:
		return programControlRegister2(value & 7)
	case 0b111:
		return programControlRegister8(value & 7)
	}

	return RegisterInvalid
}

func accumulator(data uint32) Register {
	switch data {
	case 0:
		return RegisterA
	case 1:
		return RegisterB
	}

	return RegisterInvalid
}

func oppositeAccumulator(data uint32) Register {
	switch data {
	case 0:
		return RegisterB
	case 1:
		return RegisterA
	}

	return RegisterInvalid
}

func addressModifierRegister(data uint32) Register {
	switch data {
	case 0:
		return RegisterM0
	case 1:
		return RegisterM1
	case 2:
		return RegisterM2
	case 3:
		return RegisterM3
	case 4:
		return RegisterM4
	case 5:
		return RegisterM5
	case 6:
		return RegisterM6
	case 7:
		return RegisterM7
	}

	return RegisterInvalid
}

func addressOffsetRegister(data uint32) Register {
	switch data {
	case 0:
		return RegisterN0
	case 1:
		return RegisterN1
	case 2:
		return RegisterN2
	case 3:
		return RegisterN3
	case 4:
		return RegisterN4
	case 5:
		return RegisterN5
	case 6:
		return RegisterN6
	case 7:
		return RegisterN7
	}

	return RegisterInvalid
}

func addressRegister(data uint32) Register {
	switch data {
	case 0:
		return RegisterR0
	case 1:
		return RegisterR1
	case 2:
		return RegisterR2
	case 3:
		return RegisterR3
	case 4:
		return RegisterR4
	case 5:
		return RegisterR5
	case 6:
		return RegisterR6
	case 7:
		return RegisterR7
	}

	return RegisterInvalid
}

func programControlUnitRegister(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterMR
	case 0b01:
		return RegisterCCR
	case 0b10:
		return RegisterCOM
	case 0b11:
		return RegisterEOM
	}

	return RegisterInvalid
}

func condition(value uint32) Condition {
	switch value {
	case 0b0000:
		return ConditionCC
	case 0b1000:
		return ConditionCS
	case 0b0001:
		return ConditionGE
	case 0b1001:
		return ConditionLT
	case 0b0010:
		return ConditionNE
	case 0b1010:
		return ConditionEQ
	case 0b0011:
		return ConditionPL
	case 0b1011:
		return ConditionMI
	case 0b0100:
		return ConditionNN
	case 0b1100:
		return ConditionNR
	case 0b0101:
		return ConditionEC
	case 0b1101:
		return ConditionES
	case 0b0110:
		return ConditionLC
	case 0b1110:
		return ConditionLS
	case 0b0111:
		return ConditionGT
	case 0b1111:
		return ConditionLE
	}
	return ConditionInvalid
}

// Table A-15 on Page 238
func dataALUOperands1(value uint32) Register {
	switch value {
	case 0b010:
		return RegisterA1
	case 0b011:
		return RegisterB1
	case 0b100:
		return RegisterX0
	case 0b101:
		return RegisterY0
	case 0b110:
		return RegisterX1
	case 0b111:
		return RegisterY1
	}

	return RegisterInvalid
}

// Table A-15 on Page 238
func dataALUOperands2(ggg uint32) Register {
	switch ggg {
	case 0b010:
		return RegisterA0
	case 0b011:
		return RegisterB0
	case 0b100:
		return RegisterX0
	case 0b101:
		return RegisterY0
	case 0b110:
		return RegisterX1
	case 0b111:
		return RegisterY1
	}

	return RegisterInvalid
}

// Table A-14 on Page 238
func jjj(ggg uint32, d uint32) Register {
	switch ggg {
	case 0b001:
		if d == 0 {
			return RegisterB
		}

		return RegisterA
	case 0b010:
		return RegisterX
	case 0b011:
		return RegisterY
	case 0b100:
		return RegisterX0
	case 0b101:
		return RegisterY0
	case 0b110:
		return RegisterX1
	case 0b111:
		return RegisterY1
	}

	return RegisterInvalid
}

// Table A-15 on Page 238
func dataALUOperands3(ggg uint32, d uint32) Register {
	switch ggg {
	case 0b000:
		if d == 0 {
			return RegisterB
		}

		return RegisterA
	case 0b100:
		return RegisterX0
	case 0b101:
		return RegisterY0
	case 0b110:
		return RegisterX1
	case 0b111:
		return RegisterY1
	}

	return RegisterInvalid
}

// Table A-12 on Page 237
func dataALUSourceOperands(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterX0
	case 0b01:
		return RegisterY0
	case 0b10:
		return RegisterX1
	case 0b11:
		return RegisterY1
	}

	panic("unsupported")
}

// Table A-41 on Page 247
func programControllerRegister5bit(value uint32) Register {
	switch value {
	case 0b00000, 0b00001, 0b00010, 0b00011, 0b00100, 0b00101, 0b00110, 0b00111:
		return addressModifierRegister(value & 7)
	case 0b01010:
		return RegisterEP
	case 0b10000:
		return RegisterVBA
	case 0b10001:
		return RegisterSC
	case 0b11000:
		return RegisterSZ
	case 0b11001:
		return RegisterSR
	case 0b11010:
		return RegisterOMR
	case 0b11011:
		return RegisterSP
	case 0b11100:
		return RegisterSSH
	case 0b11101:
		return RegisterSSL
	case 0b11110:
		return RegisterLA
	case 0b11111:
		return RegisterLC
	}

	return RegisterInvalid
}

// Table A-28 on Page 242
func dataALUMultiplyOperands2(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterX0
	case 0b01:
		return RegisterY0
	case 0b10:
		return RegisterX1
	case 0b11:
		return RegisterY1
	}

	panic("unsupported")
}

// Table A-27 on Page 242
func dataALUMultiplyOperands1(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterY1
	case 0b01:
		return RegisterX0
	case 0b10:
		return RegisterY0
	case 0b11:
		return RegisterX1
	}

	panic("unsupported")
}

// Table A-11 on Page 237
func inputRegister(value uint32) Register {
	switch value {
	case 0:
		return RegisterX
	case 1:
		return RegisterY
	}

	panic("unsupported")
}

// Table A-35 on Page 245
func xrSourceDestination(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterX0
	case 0b01:
		return RegisterX1
	case 0b10:
		return RegisterA
	case 0b11:
		return RegisterB
	}
	panic("unsupported")
}

// Table A-35 on Page 245
func xrDestination2(value uint32) Register {
	switch value {
	case 0b0:
		return RegisterY0
	case 0b1:
		return RegisterY1
	}
	panic("unsupported")
}

// Table A-36 on Page 245
func rySourceDestination(value uint32) Register {
	switch value {
	case 0b00:
		return RegisterY0
	case 0b01:
		return RegisterY1
	case 0b10:
		return RegisterA
	case 0b11:
		return RegisterB
	}
	panic("unsupported")
}

// Table A-36 on Page 245
func ryDestination(value uint32) Register {
	switch value {
	case 0b0:
		return RegisterX0
	case 0b1:
		return RegisterX1
	}
	panic("unsupported")
}

// Table A-23 on Page 241
func longMoveRegister(value uint32) Register {
	switch value {
	case 0b000:
		return RegisterA10
	case 0b001:
		return RegisterB10
	case 0b010:
		return RegisterX
	case 0b011:
		return RegisterY
	case 0b100:
		return RegisterA
	case 0b101:
		return RegisterB
	case 0b110:
		return RegisterAB
	case 0b111:
		return RegisterBA
	}

	panic("unsupported")
}
