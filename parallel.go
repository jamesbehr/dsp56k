package dsp56k

import (
	"sort"
)

type decodeALURow struct {
	Mask, Constant uint8
	Builder        func() ALUOperation
}

var decodeALUTable []decodeALURow

func registerALUOperation(mask, constant uint8, builder func() ALUOperation) {
	if constant&mask != constant {
		panic("constant may only have bits set where bits in the mask are set")
	}

	// Insert into a sorted slice of decode rows
	i := sort.Search(len(decodeALUTable), func(i int) bool {
		return mask > decodeALUTable[i].Mask
	})

	decodeALUTable = append(decodeALUTable, decodeALURow{})
	copy(decodeALUTable[i+1:], decodeALUTable[i:])
	decodeALUTable[i] = decodeALURow{mask, constant, builder}
}

type decodeMoveRow struct {
	Mask, Constant uint16
	Builder        func() ParallelMove
}

var decodeMoveTable []decodeMoveRow

func registerParallelMove(mask, constant uint16, builder func() ParallelMove) {
	if constant&mask != constant {
		panic("constant may only have bits set where bits in the mask are set")
	}

	// Insert into a sorted slice of decode rows
	i := sort.Search(len(decodeMoveTable), func(i int) bool {
		return mask > decodeMoveTable[i].Mask
	})

	decodeMoveTable = append(decodeMoveTable, decodeMoveRow{})
	copy(decodeMoveTable[i+1:], decodeMoveTable[i:])
	decodeMoveTable[i] = decodeMoveRow{mask, constant, builder}
}

// AccumulatorWrites records what portions of the acummulator have been written
// to. Parallel moves can't write to portions of the accumulator if they have
// already been written to by the ALU operation.
//
// The accumulators are 56-bits and have multiple aliases for different
// segments.
//
// For example, the A accummulator is actually 3 smaller registers concatenated
// together: A2:A1:A0, so a write to any of these will preclude a write to A
// from happening.
//
// Tracking what parts of the acummulators are available for writing is
// critical to correctly decoding the parallel instructions.
type AccumulatorWrites uint32

const (
	writeA0  = 1 << 0
	writeA1  = 1 << 1
	writeA2  = 1 << 2
	writeB0  = 1 << 3
	writeB1  = 1 << 4
	writeB2  = 1 << 5
	writeA   = writeA2 | writeA1 | writeA0
	writeB   = writeB2 | writeB1 | writeB0
	writeA10 = writeA1 | writeA0
	writeB10 = writeB1 | writeB0
	writeAB  = writeA | writeB
)

// HasWritten returns true if any portion of register r has had a write recorded to it.
// If r is not an accumulator, then it always returns false.
func (w AccumulatorWrites) HasWritten(r Register) bool {
	switch r {
	case RegisterAB, RegisterBA:
		// Only need to check for writes to A or B, as the only difference
		// between these registers is the order the bits appear.
		return w&writeAB != 0
	case RegisterA10:
		return w&writeA10 != 0
	case RegisterA:
		return w&writeA != 0
	case RegisterA0:
		return w&writeA0 != 0
	case RegisterA1:
		return w&writeA1 != 0
	case RegisterA2:
		return w&writeA2 != 0
	case RegisterB10:
		return w&writeB10 != 0
	case RegisterB:
		return w&writeB != 0
	case RegisterB0:
		return w&writeB0 != 0
	case RegisterB1:
		return w&writeB1 != 0
	case RegisterB2:
		return w&writeB2 != 0
	}

	return false
}

// RecordWrite records a write to the specified register.
func (w *AccumulatorWrites) RecordWrite(r Register) {
	switch r {
	case RegisterAB, RegisterBA:
		*w = writeAB
	case RegisterA10:
		*w = writeA10
	case RegisterA:
		*w = writeA
	case RegisterA0:
		*w = writeA0
	case RegisterA1:
		*w = writeA1
	case RegisterA2:
		*w = writeA2
	case RegisterB10:
		*w = writeB10
	case RegisterB:
		*w = writeB
	case RegisterB0:
		*w = writeB0
	case RegisterB1:
		*w = writeB1
	case RegisterB2:
		*w = writeB2
	}
}

// Some instructions only affect bits 47-24 of the destination accumulator
// (i.e. A1 or B1), even though they name the full 56-bit accumulator. This
// means that parallel moves to A2/B2 or A0/B0 should still be possible.
// This function simply fixes which part of the accumulator is recorded as
// being written to.
func (w *AccumulatorWrites) RecordPartialWrite(accumulator Register) {
	switch accumulator {
	case RegisterA:
		w.RecordWrite(RegisterA1)
	case RegisterB:
		w.RecordWrite(RegisterB1)
	}
}

// Some instructions don't allow duplicate writes at all. This helper function
// returns true if any portion of an accumulator has been written to.
func (w AccumulatorWrites) HasPartiallyWritten(accumulator Register) bool {
	switch accumulator {
	case RegisterA, RegisterA0, RegisterA1, RegisterA2:
		return w.HasWritten(RegisterA)
	case RegisterB, RegisterB0, RegisterB1, RegisterB2:
		return w.HasWritten(RegisterB)
	}

	return false
}

type ParallelInstruction struct {
	Move ParallelMove
	ALU  ALUOperation
}

// Silences the go vet warnings about unkeyed struct literals from other packages
func Parallel(move ParallelMove, alu ALUOperation) Instruction {
	return &ParallelInstruction{move, alu}
}

func (ins *ParallelInstruction) UsesExtensionWord() bool {
	return ins.Move.UsesExtensionWord()
}

func init() {
	registerInstruction(0x800000, 0x800000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x704000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x700000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x604000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x600000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x504000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x500000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf44000, 0x444000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf44000, 0x440000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf44000, 0x404000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf44000, 0x400000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf80000, 0x380000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf80000, 0x300000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf80000, 0x280000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xfc0000, 0x240000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xfe0000, 0x220000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xfd0000, 0x210000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xfc8000, 0x208000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffe000, 0x204000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffe000, 0x202000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffc000, 0x200000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x104000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xf04000, 0x100000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffc000, 0x098000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffc000, 0x090000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffc000, 0x088000, func() Instruction { return new(ParallelInstruction) })
	registerInstruction(0xffc000, 0x080000, func() Instruction { return new(ParallelInstruction) })
}

var decodeALUCache [0x100]int

func decodeALUOperation(opcode uint8) ALUOperation {
	if c := decodeALUCache[opcode]; c < 0 {
		// Cached failure
		return nil
	} else if c > 0 {
		// Cache hit
		return decodeALUTable[c-1].Builder()
	}

	for i, candidate := range decodeALUTable {
		if opcode&candidate.Mask == candidate.Constant {
			decodeALUCache[opcode] = i + 1
			return candidate.Builder()
		}
	}

	decodeALUCache[opcode] = -1
	return nil
}

var decodeMoveCache [0x10000]int

func decodeParallelMove(opcode uint16) ParallelMove {
	if c := decodeMoveCache[opcode]; c < 0 {
		// Cached failure
		return nil
	} else if c > 0 {
		// Cache hit
		return decodeMoveTable[c-1].Builder()
	}

	// Cache miss
	for i, candidate := range decodeMoveTable {
		if opcode&candidate.Mask == candidate.Constant {
			decodeMoveCache[opcode] = i + 1
			return candidate.Builder()
		}
	}

	decodeMoveCache[opcode] = -1
	return nil
}

func (ins *ParallelInstruction) Decode(opcode uint32, extensionWord uint32) bool {
	aluOpcode := uint8(opcode)
	ins.ALU = decodeALUOperation(aluOpcode)
	if ins.ALU == nil {
		return false
	}

	writes, ok := ins.ALU.DecodeALU(aluOpcode)
	if !ok {
		return false
	}

	parallelMoveOpcode := uint16(opcode >> 8)
	ins.Move = decodeParallelMove(parallelMoveOpcode)
	if ins.Move == nil {
		return false
	}

	if !ins.Move.DecodeParallelMove(parallelMoveOpcode, extensionWord, writes) {
		return false
	}

	_, isMoveNop := ins.Move.(*Move_Nop)
	_, isAluNop := ins.ALU.(*AluNop)
	if isAluNop && isMoveNop {
		return false
	}

	return true
}

func (ins *ParallelInstruction) Disassemble(w TokenWriter) error {
	if err := ins.ALU.Disassemble(w); err != nil {
		return err
	}

	if err := ins.Move.Disassemble(w); err != nil {
		return err
	}

	return nil
}

type ParallelMove interface {
	Disassemble(w TokenWriter) error
	DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool
	UsesExtensionWord() bool
}

type ALUOperation interface {
	Disassemble(w TokenWriter) error
	DecodeALU(opcode uint8) (writes AccumulatorWrites, ok bool)
}
