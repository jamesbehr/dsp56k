# DSP56K
A library for disassembling DSP56000 instructions in Go.

The DSP56000 is a specialised architecture for digital signal processing (DSP).

It has a 24-bit word size, with each instruction being a single word long.
Some instructions use an optional extension word, which contains addresses or immediate data.
The instruction type (and the presence of an extension word) is always identifiable from looking at the first word.
The architecture is unusual that it defines multiple regions of memory.
Code is executed from the region known as P (program) memory.
Because 24-bit is an unusual word size, this library uses 32-bit words in its memory representation.
The upper 8 bits of each 32-bit word are ignored.

You can use this library to build a disassembler, for example to disassemble
the instructions starting at address `P:$0`, continuing until an instruction
that halts execution (e.g. unconditional branches and return instructions)

```go
func Disassemble(memory []uint32) {
    pc := 0
    for {
        instruction := dsp56k.Decode(memory[pc], memory[pc + 1])
        if instruction != nil {
            asm, err := dsp56k.Disassemble(instruction, nil, pc)
            if err != nil {
                log.Fatal("failed to disassemble instruction: ", err)
            }

            fmt.Println(asm)

            if instruction.UsesExtensionWord() {
                pc += 2
            } else {
                pc++
            }

            if staticanalysis.Halts(instruction) {
                break
            }
        } else {
            pc++
        }
    }
}
```
