//go:build exclude

package dsp56k

import (
	"bytes"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// The test data is generated using the Motorola Disassembler.
//
// The Symphony Studio provides libdsp56300.dll which exports the procedure
// dspt_unasm_563, which is able to disassemble an array of opcodes.
// The 56300 is a 24-bit word system. Each opcode is one word long, followed by
// an optional extension word, which is used for long immediate data.
//
// dspt_unasm_563 takes an array of up to 3 opcodes, returns the number of
// words used and stores the resulting assembler as up to 120 ASCII characters.
// This call signature is to support other platforms as well, so the other
// arguments and 3rd opcode are unnecessary. It will return a data constant in
// the event of an invalid opode (e.g. "dc $123456").
//
// The following Go program was used to create the test data. It iterates
// through every possible opcode. It uses the same extension word each time.
func main() {
	dllPath := "C:\\Symphony-Studio\\dsp56720-devtools\\dist\\dsp56720\\libdsp56300.dll"

	var mod = syscall.NewLazyDLL(dllPath)
	proc := mod.NewProc("dspt_unasm_563")

	replacements := map[uint32][]string{
		0x202015: {"maxm a,b ifcc", "max b,a ifcc"},
		0x202115: {"maxm a,b ifge", "max b,a ifge"},
		0x202215: {"maxm a,b ifne", "max b,a ifne"},
		0x202315: {"maxm a,b ifpl", "max b,a ifpl"},
		0x202415: {"maxm a,b ifnn", "max b,a ifnn"},
		0x202515: {"maxm a,b ifec", "max b,a ifec"},
		0x202615: {"maxm a,b iflc", "max b,a iflc"},
		0x202715: {"maxm a,b ifgt", "max b,a ifgt"},
		0x202815: {"maxm a,b ifcs", "max b,a ifcs"},
		0x202915: {"maxm a,b iflt", "max b,a iflt"},
		0x202a15: {"maxm a,b ifeq", "max b,a ifeq"},
		0x202b15: {"maxm a,b ifmi", "max b,a ifmi"},
		0x202c15: {"maxm a,b ifnr", "max b,a ifnr"},
		0x202d15: {"maxm a,b ifes", "max b,a ifes"},
		0x202e15: {"maxm a,b ifls", "max b,a ifls"},
		0x202f15: {"maxm a,b ifle", "max b,a ifle"},
		0x203015: {"maxm a,b ifcc.u", "max b,a ifcc.u"},
		0x203115: {"maxm a,b ifge.u", "max b,a ifge.u"},
		0x203215: {"maxm a,b ifne.u", "max b,a ifne.u"},
		0x203315: {"maxm a,b ifpl.u", "max b,a ifpl.u"},
		0x203415: {"maxm a,b ifnn.u", "max b,a ifnn.u"},
		0x203515: {"maxm a,b ifec.u", "max b,a ifec.u"},
		0x203615: {"maxm a,b iflc.u", "max b,a iflc.u"},
		0x203715: {"maxm a,b ifgt.u", "max b,a ifgt.u"},
		0x203815: {"maxm a,b ifcs.u", "max b,a ifcs.u"},
		0x203915: {"maxm a,b iflt.u", "max b,a iflt.u"},
		0x203a15: {"maxm a,b ifeq.u", "max b,a ifeq.u"},
		0x203b15: {"maxm a,b ifmi.u", "max b,a ifmi.u"},
		0x203c15: {"maxm a,b ifnr.u", "max b,a ifnr.u"},
		0x203d15: {"maxm a,b ifes.u", "max b,a ifes.u"},
		0x203e15: {"maxm a,b ifls.u", "max b,a ifls.u"},
		0x203f15: {"maxm a,b ifle.u", "max b,a ifle.u"},
	}

	entry := [40]byte{}

	f, err := os.Create("opcodes.bin")
	if err != nil {
		log.Fatal(err)
	}

	for opcode := uint32(0); opcode <= 0xffffff; opcode++ {
		ops := []uint32{opcode, 0xdeface, 0}
		str := make([]uint8, 120)
		ret, _, _ := proc.Call(
			uintptr(unsafe.Pointer(&ops[0])),
			uintptr(unsafe.Pointer(&str[0])),
			0, 0, 0)

		numWords := int(ret)
		if numWords > 2 {
			log.Fatal("too many words")
		}

		asm := bytes.TrimRight(str, "\x00")
		asmStr := strings.TrimSpace(string(asm))

		if len(asmStr) > len(entry) {
			log.Fatal("string too long")
		}

		if r, ok := replacements[opcode]; ok && asmStr == r[1] {
			asmStr = r[0]
		}

		entry[0] = byte(numWords)
		copy(entry[1:], append([]byte(asmStr), 0))

		if _, err := f.Write(entry[:]); err != nil {
			log.Fatal(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
