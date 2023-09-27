package dsp56k

func bit16(byte uint16, shift int) uint32 {
	return uint32(byte>>shift) & 1
}

func mask16(word uint16, shift, width int) uint32 {
	mask := uint16(1<<width - 1)
	return uint32((word >> shift) & mask)
}

func multimask16(word uint16, maskPairs ...int) uint32 {
	if l := len(maskPairs); l < 4 || l%2 > 0 {
		panic("multimask16 requires two or more pairs of values")
	}

	value := uint32(0)

	for i := 0; i < len(maskPairs)/2; i++ {
		shift, width := maskPairs[i*2], maskPairs[i*2+1]
		mask := uint32(1<<width - 1)
		field := uint32(word>>shift) & mask
		value = (value << width) | field
	}

	return value
}

func init() {
	registerParallelMove(0b1000000000000000, 0b1000000000000000, func() ParallelMove { return new(Movexy) })
	registerParallelMove(0b1100100001000000, 0b0100000000000000, func() ParallelMove { return new(Movex_aa) })
	registerParallelMove(0b1100100001000000, 0b0100000001000000, func() ParallelMove { return new(Movex_ea) })
	registerParallelMove(0b1100100001000000, 0b0100100000000000, func() ParallelMove { return new(Movey_aa) })
	registerParallelMove(0b1100100001000000, 0b0100100001000000, func() ParallelMove { return new(Movey_ea) })
	registerParallelMove(0b1100100001111111, 0b0100000001110000, func() ParallelMove { return new(Movex_ea_Abs) })
	registerParallelMove(0b1100100011111111, 0b0100000011110100, func() ParallelMove { return new(Movex_ea_Imm) })
	registerParallelMove(0b1100100001111111, 0b0100100001110000, func() ParallelMove { return new(Movey_ea_Abs) })
	registerParallelMove(0b1100100011111111, 0b0100100011110100, func() ParallelMove { return new(Movey_ea_Imm) })
	registerParallelMove(0b1110000000000000, 0b0010000000000000, func() ParallelMove { return new(Move_xx) })
	registerParallelMove(0b1111000001000000, 0b0001000000000000, func() ParallelMove { return new(Movexr_ea) })
	registerParallelMove(0b1111000001000000, 0b0001000001000000, func() ParallelMove { return new(Moveyr_ea) })
	registerParallelMove(0b1111000001111111, 0b0001000000110000, func() ParallelMove { return new(Movexr_ea_Abs) })
	registerParallelMove(0b1111000001111111, 0b0001000001110000, func() ParallelMove { return new(Moveyr_ea_Abs) })
	registerParallelMove(0b1111000011111111, 0b0001000010110100, func() ParallelMove { return new(Movexr_ea_Imm) })
	registerParallelMove(0b1111000011111111, 0b0001000011110100, func() ParallelMove { return new(Moveyr_ea_Imm) })
	registerParallelMove(0b1111010001000000, 0b0100000000000000, func() ParallelMove { return new(Movel_aa) })
	registerParallelMove(0b1111010001000000, 0b0100000001000000, func() ParallelMove { return new(Movel_ea) })
	registerParallelMove(0b1111010001111111, 0b0100000001110000, func() ParallelMove { return new(Movel_ea_Abs) })
	registerParallelMove(0b1111110000000000, 0b0010000000000000, func() ParallelMove { return new(Mover) })
	registerParallelMove(0b1111111011000000, 0b0000100000000000, func() ParallelMove { return new(Movexr_A) })
	registerParallelMove(0b1111111011000000, 0b0000100010000000, func() ParallelMove { return new(Moveyr_A) })
	registerParallelMove(0b1111111111100000, 0b0010000001000000, func() ParallelMove { return new(Move_ea) })
	registerParallelMove(0b1111111111110000, 0b0010000000100000, func() ParallelMove { return new(Ifcc) })
	registerParallelMove(0b1111111111110000, 0b0010000000110000, func() ParallelMove { return new(Ifcc_U) })
	registerParallelMove(0b1111111111111111, 0b0010000000000000, func() ParallelMove { return new(Move_Nop) })
}

// Page 109
// 001000000010CCCC
// IFcc
type Ifcc struct {
	Condition Condition
}

func (*Ifcc) UsesExtensionWord() bool {
	return false
}

func (ins *Ifcc) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Condition = condition(mask16(opcode, 0, 4))
	return true
}

func (ins *Ifcc) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		MnemonicIf,
		ins.Condition,
	)
}

// Page 110
// 001000000011CCCC
// IFcc.U
type Ifcc_U struct {
	Condition Condition
}

func (*Ifcc_U) UsesExtensionWord() bool {
	return false
}

func (ins *Ifcc_U) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Condition = condition(mask16(opcode, 0, 4))
	return true
}

func (ins *Ifcc_U) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		MnemonicIf,
		ins.Condition,
		UpdateCCRToken,
	)
}

// Page 157
// 0010000000000000
// MOVE S,D
type Move_Nop struct {
}

func (*Move_Nop) UsesExtensionWord() bool {
	return false
}

func (ins *Move_Nop) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	return true
}

func (ins *Move_Nop) Disassemble(w TokenWriter) error {
	return w.Write(
	// ColumnSeparator,
	)
}

// Page 158
// 001dddddiiiiiiii
// (...) #xx,D
type Move_xx struct {
	Immediate   uint32
	Destination Register
}

func (*Move_xx) UsesExtensionWord() bool {
	return false
}

func (ins *Move_xx) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Immediate = mask16(opcode, 0, 8)
	ins.Destination = register5Bit(mask16(opcode, 8, 5))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	return true
}

func (ins *Move_xx) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Page 160
// 001000eeeeeddddd
// (...) S,D
type Mover struct {
	Source      Register
	Destination Register
}

func (*Mover) UsesExtensionWord() bool {
	return false
}

func (ins *Mover) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Source = register5Bit(mask16(opcode, 5, 5))
	ins.Destination = register5Bit(mask16(opcode, 0, 5))

	if ins.Source == RegisterInvalid {
		return false
	}

	if ins.Destination == RegisterInvalid {
		return false
	}

	if writes.HasWritten(ins.Destination) {
		return false
	}

	return true
}

func (ins *Mover) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Page 162
// 00100000010MMRRR
// (...) ea
type Move_ea struct {
	EffectiveAddress EffectiveAddress
}

func (*Move_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Move_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 2), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	return true
}

func (ins *Move_ea) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Page 163
// 01dd0dddW1MMMRRR
// (...) X:ea,D / (...) S,X:ea
type Movex_ea struct {
	SourceOrDestination Register
	IsWrite             bool
	EffectiveAddress    EffectiveAddress
}

func (ins *Movex_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Movex_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movex_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryX, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, ins.EffectiveAddress,
	)
}

// Page 163
// 01dd0dddW1MMMRRR
// (...) #xxxxxx,D
// Absolute addressing mode
type Movex_ea_Abs struct {
	SourceOrDestination Register
	IsWrite             bool
	Address             uint32
}

func (ins *Movex_ea_Abs) UsesExtensionWord() bool {
	return true
}

func (ins *Movex_ea_Abs) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Address = extensionWord

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movex_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryX, ForceLongAddressMode, absAddr(MemoryX, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, ForceLongAddressMode, absAddr(MemoryX, ins.Address),
	)
}

// Page 163
// 01dd0dddW1MMMRRR
// (...) #xxxxxx,D
// Absolute addressing mode
type Movex_ea_Imm struct {
	Destination Register
	Immediate   uint32
}

func (ins *Movex_ea_Imm) UsesExtensionWord() bool {
	return true
}

func (ins *Movex_ea_Imm) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Destination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.Immediate = extensionWord

	if writes.HasWritten(ins.Destination) {
		return false
	}

	return true
}

func (ins *Movex_ea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Page 163
// 01dd0dddW0aaaaaa
// (...) X:aa,D / (...) S,X:aa
type Movex_aa struct {
	SourceOrDestination Register
	Address             uint32
	IsWrite             bool
}

func (*Movex_aa) UsesExtensionWord() bool {
	return false
}

func (ins *Movex_aa) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Address = mask16(opcode, 0, 6)

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	if ins.IsWrite && writes.HasPartiallyWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movex_aa) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryX, ForceShortAddressMode, absAddr(MemoryX, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, ForceShortAddressMode, absAddr(MemoryX, ins.Address),
	)
}

// Page 166
// 0001ffdFW0MMMRRR
// (...) X:ea,D1 S2,D2 / (...) S1,X:ea S2, D2 / (...) #xxxx,D1 S2,D2
type Movexr_ea struct {
	EffectiveAddress    EffectiveAddress
	IsWrite             bool
	Source              Register
	Destination         Register
	SourceOrDestination Register
}

func (ins *Movexr_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Movexr_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Destination = xrDestination2(bit16(opcode, 8))
	ins.Source = accumulator(bit16(opcode, 9))
	ins.SourceOrDestination = xrSourceDestination(mask16(opcode, 10, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movexr_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryX, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
			ColumnSeparator,
			ins.Source,
			OperandSeparator,
			ins.Destination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, ins.EffectiveAddress,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Page 166
// 0001ffdFW0MMMRRR
// (...) X:ea,D1 S2,D2 / (...) S1,X:ea S2, D2 / (...) #xxxx,D1 S2,D2
type Movexr_ea_Abs struct {
	Address             uint32
	IsWrite             bool
	Source              Register
	Destination         Register
	SourceOrDestination Register
}

func (ins *Movexr_ea_Abs) UsesExtensionWord() bool {
	return true
}

func (ins *Movexr_ea_Abs) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Address = extensionWord
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Destination = xrDestination2(bit16(opcode, 8))
	ins.Source = accumulator(bit16(opcode, 9))
	ins.SourceOrDestination = xrSourceDestination(mask16(opcode, 10, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movexr_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryX, ForceLongAddressMode, absAddr(MemoryX, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
			ColumnSeparator,
			ins.Source,
			OperandSeparator,
			ins.Destination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, ForceLongAddressMode, absAddr(MemoryX, ins.Address),
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Page 166
// 0001ffdFW0MMMRRR
// (...) X:ea,D1 S2,D2 / (...) S1,X:ea S2, D2 / (...) #xxxx,D1 S2,D2
type Movexr_ea_Imm struct {
	Immediate    uint32
	Source       Register
	Destination  Register
	Destination2 Register
}

func (ins *Movexr_ea_Imm) UsesExtensionWord() bool {
	return true
}

func (ins *Movexr_ea_Imm) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Immediate = extensionWord
	ins.Destination = xrDestination2(bit16(opcode, 8))
	ins.Source = accumulator(bit16(opcode, 9))
	ins.Destination2 = xrSourceDestination(mask16(opcode, 10, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if writes.HasWritten(ins.Destination2) {
		return false
	}

	return true
}

func (ins *Movexr_ea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination2,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Page 166
// 0000100d00MMMRRR
// (...) A -> X:ea X0 -> A / (...) B -> X:ea X0 -> B
type Movexr_A struct {
	EffectiveAddress EffectiveAddress
	Accumulator      Register
	Source           Register
}

func (*Movexr_A) UsesExtensionWord() bool {
	return false
}

func (ins *Movexr_A) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Accumulator = accumulator(bit16(opcode, 8))
	ins.Source = RegisterX0

	if writes.HasWritten(ins.Accumulator) {
		return false
	}

	return true
}

func (ins *Movexr_A) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ins.Accumulator,
		OperandSeparator,
		MemoryX,
		ins.EffectiveAddress,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Accumulator,
	)
}

// Page 168
// 01dd1dddW1MMMRRR
// (...) Y:ea,D / (...) S,Y:ea
type Movey_ea struct {
	SourceOrDestination Register
	IsWrite             bool
	EffectiveAddress    EffectiveAddress
}

func (ins *Movey_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Movey_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movey_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryY, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, ins.EffectiveAddress,
	)
}

// Page 168
// 01dd1dddW1MMMRRR
// (...) #xxxx,D
type Movey_ea_Abs struct {
	SourceOrDestination Register
	IsWrite             bool
	Address             uint32
}

func (ins *Movey_ea_Abs) UsesExtensionWord() bool {
	return true
}

func (ins *Movey_ea_Abs) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Address = extensionWord

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movey_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryY, ForceLongAddressMode, absAddr(MemoryY, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, ForceLongAddressMode, absAddr(MemoryY, ins.Address),
	)
}

// Page 168
// 01dd1dddW1MMMRRR
// (...) #xxxx,D
type Movey_ea_Imm struct {
	Destination Register
	Immediate   uint32
}

func (ins *Movey_ea_Imm) UsesExtensionWord() bool {
	return true
}

func (ins *Movey_ea_Imm) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Destination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.Immediate = extensionWord

	if writes.HasWritten(ins.Destination) {
		return false
	}

	return true
}

func (ins *Movey_ea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Page 168
// 01dd1dddW0aaaaaa
// (...) Y:aa,D / (...) S,Y:aa
type Movey_aa struct {
	SourceOrDestination Register
	Address             uint32
	IsWrite             bool
}

func (*Movey_aa) UsesExtensionWord() bool {
	return false
}

func (ins *Movey_aa) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = register5Bit(multimask16(opcode, 12, 2, 8, 3))
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Address = mask16(opcode, 0, 6)

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	if ins.IsWrite && writes.HasPartiallyWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movey_aa) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryY, ForceShortAddressMode, absAddr(MemoryY, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, ForceShortAddressMode, absAddr(MemoryY, ins.Address),
	)
}

// Page 171
// 0001deffW1MMMRRR
// (...) S1,D1 Y:ea,D2 / (...) S1,D1 S2,Y:ea
type Moveyr_ea struct {
	EffectiveAddress    EffectiveAddress
	IsWrite             bool
	Source              Register
	Destination         Register
	SourceOrDestination Register
}

func (ins *Moveyr_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Moveyr_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Destination = ryDestination(bit16(opcode, 10))
	ins.Source = accumulator(bit16(opcode, 11))
	ins.SourceOrDestination = rySourceDestination(mask16(opcode, 8, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Moveyr_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			ins.Source,
			OperandSeparator,
			ins.Destination,
			ColumnSeparator,
			MemoryY, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, ins.EffectiveAddress,
	)
}

// Page 171
// 0001deffW1MMMRRR
// (...) S1,D1 #xxxx,D2
type Moveyr_ea_Abs struct {
	Address             uint32
	IsWrite             bool
	Source              Register
	Destination         Register
	SourceOrDestination Register
}

func (ins *Moveyr_ea_Abs) UsesExtensionWord() bool {
	return true
}

func (ins *Moveyr_ea_Abs) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Address = extensionWord
	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Destination = ryDestination(bit16(opcode, 10))
	ins.Source = accumulator(bit16(opcode, 11))
	ins.SourceOrDestination = rySourceDestination(mask16(opcode, 8, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Moveyr_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			ins.Source,
			OperandSeparator,
			ins.Destination,
			ColumnSeparator,
			MemoryY, ForceLongAddressMode, absAddr(MemoryY, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, ForceLongAddressMode, absAddr(MemoryY, ins.Address),
	)
}

// Page 171
// 0001deffW1MMMRRR
// (...) S1,D1 #xxxx,D2
type Moveyr_ea_Imm struct {
	Immediate    uint32
	IsWrite      bool
	Source       Register
	Destination  Register
	Destination2 Register
}

func (ins *Moveyr_ea_Imm) UsesExtensionWord() bool {
	return true
}

func (ins *Moveyr_ea_Imm) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Immediate = extensionWord
	ins.Destination = ryDestination(bit16(opcode, 10))
	ins.Source = accumulator(bit16(opcode, 11))
	ins.Destination2 = rySourceDestination(mask16(opcode, 8, 2))

	if writes.HasWritten(ins.Destination) {
		return false
	}

	if writes.HasWritten(ins.Destination2) {
		return false
	}

	return true
}

func (ins *Moveyr_ea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination2,
	)
}

// Page 171
// 0000100d10MMMRRR
// (...) Y0 -> A A -> Y:ea / (...) Y0 -> B B -> Y:ea
type Moveyr_A struct {
	EffectiveAddress EffectiveAddress
	Accumulator      Register
	Source           Register
}

func (*Moveyr_A) UsesExtensionWord() bool {
	return false
}

func (ins *Moveyr_A) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Accumulator = accumulator(bit16(opcode, 8))
	ins.Source = RegisterY0

	if writes.HasWritten(ins.Accumulator) {
		return false
	}

	return true
}

func (ins *Moveyr_A) Disassemble(w TokenWriter) error {
	return w.Write(
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Accumulator,
		ColumnSeparator,
		ins.Accumulator,
		OperandSeparator,
		MemoryY,
		ins.EffectiveAddress,
	)
}

// Page 174
// 0100L0LLW1MMMRRR
// (...) L:ea,D / (...) S,L:ea
type Movel_ea struct {
	SourceOrDestination Register
	EffectiveAddress    EffectiveAddress
	IsWrite             bool
}

func (*Movel_ea) UsesExtensionWord() bool {
	return false
}

func (ins *Movel_ea) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.EffectiveAddress = addressMode(mask16(opcode, 3, 3), mask16(opcode, 0, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.SourceOrDestination = longMoveRegister(multimask16(opcode, 11, 1, 8, 2))
	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))
	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movel_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryL, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryL, ins.EffectiveAddress,
	)
}

// Page 174
// 0100L0LLW1MMMRRR
// (...) L:ea,D / (...) S,L:ea
type Movel_ea_Abs struct {
	SourceOrDestination Register
	Address             uint32
	IsWrite             bool
}

func (*Movel_ea_Abs) UsesExtensionWord() bool {
	return true
}

func (ins *Movel_ea_Abs) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.Address = extensionWord

	ins.SourceOrDestination = longMoveRegister(multimask16(opcode, 11, 1, 8, 2))
	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))
	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movel_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryL, ForceLongAddressMode, absAddr(MemoryL, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryL, ForceLongAddressMode, absAddr(MemoryL, ins.Address),
	)
}

// Page 174
// 0100L0LLW0aaaaaa
// (...) L:aa,D / (...) S,L:aa
type Movel_aa struct {
	Address             uint32
	SourceOrDestination Register
	IsWrite             bool
}

func (*Movel_aa) UsesExtensionWord() bool {
	return false
}

func (ins *Movel_aa) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	ins.SourceOrDestination = longMoveRegister(multimask16(opcode, 11, 1, 8, 2))
	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	ins.IsWrite = boolean(bit16(opcode, 7))
	ins.Address = mask16(opcode, 0, 6)

	if ins.IsWrite && writes.HasWritten(ins.SourceOrDestination) {
		return false
	}

	return true
}

func (ins *Movel_aa) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			ColumnSeparator,
			MemoryL, ForceShortAddressMode, absAddr(MemoryL, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryL, ForceShortAddressMode, absAddr(MemoryL, ins.Address),
	)
}

type MoveEffectiveAddressAndRegister struct {
	EffectiveAddress    EffectiveAddress
	SourceOrDestination Register
	IsWrite             bool
}

// Page 176
// 1wmmeeffWrrMMRRR
// (...) X:<eax>,D1 Y:<eay>,D2 / (...) X:<eax>,D1 S2,Y:<eay> / (...) S1,X:<eax> Y:<eay>,D2 / (...) S1,X:<eax> S2,Y:<eay>
type Movexy struct {
	X, Y MoveEffectiveAddressAndRegister
}

func (*Movexy) UsesExtensionWord() bool {
	return false
}

func (ins *Movexy) DecodeParallelMove(opcode uint16, extensionWord uint32, writes AccumulatorWrites) bool {
	addresssRegisterX := mask16(opcode, 0, 3)
	ins.X.EffectiveAddress = addressMode2Bit(mask16(opcode, 3, 2), addresssRegisterX)
	ins.X.SourceOrDestination = xrSourceDestination(mask16(opcode, 10, 2))
	ins.X.IsWrite = boolean(bit16(opcode, 7))

	if ins.X.IsWrite && writes.HasWritten(ins.X.SourceOrDestination) {
		return false
	}

	// This is only two bits, so it can't refer to all 8 address registers. It
	// instead splits them into two banks, r4-r7 and r0-r3. It uses the
	// opposite bank from the address register used in the X bus.
	addressRegisterY := mask16(opcode, 5, 2)
	if addresssRegisterX < 4 {
		addressRegisterY += 4
	}

	ins.Y.EffectiveAddress = addressMode2Bit(mask16(opcode, 12, 2), addressRegisterY)
	ins.Y.SourceOrDestination = rySourceDestination(mask16(opcode, 8, 2))
	ins.Y.IsWrite = boolean(bit16(opcode, 14))

	if ins.Y.IsWrite && writes.HasWritten(ins.Y.SourceOrDestination) {
		return false
	}

	if ins.X.IsWrite && ins.Y.IsWrite && ins.X.SourceOrDestination == ins.Y.SourceOrDestination {
		return false
	}

	return true
}

func (ins *Movexy) Disassemble(w TokenWriter) error {
	switch {
	case ins.X.IsWrite && ins.Y.IsWrite:
		return w.Write(
			ColumnSeparator,

			MemoryX, ins.X.EffectiveAddress,
			OperandSeparator,
			ins.X.SourceOrDestination,

			ColumnSeparator,

			MemoryY, ins.Y.EffectiveAddress,
			OperandSeparator,
			ins.Y.SourceOrDestination,
		)
	case ins.X.IsWrite:
		return w.Write(
			ColumnSeparator,

			MemoryX, ins.X.EffectiveAddress,
			OperandSeparator,
			ins.X.SourceOrDestination,

			ColumnSeparator,

			ins.Y.SourceOrDestination,
			OperandSeparator,
			MemoryY, ins.Y.EffectiveAddress,
		)
	case ins.Y.IsWrite:
		return w.Write(
			ColumnSeparator,

			ins.X.SourceOrDestination,
			OperandSeparator,
			MemoryX, ins.X.EffectiveAddress,

			ColumnSeparator,

			MemoryY, ins.Y.EffectiveAddress,
			OperandSeparator,
			ins.Y.SourceOrDestination,
		)
	default:
		return w.Write(
			ColumnSeparator,

			ins.X.SourceOrDestination,
			OperandSeparator,
			MemoryX, ins.X.EffectiveAddress,

			ColumnSeparator,

			ins.Y.SourceOrDestination,
			OperandSeparator,
			MemoryY, ins.Y.EffectiveAddress,
		)
	}
}
