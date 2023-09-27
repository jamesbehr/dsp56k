package dsp56k

// Function Idx: func170
// 0000000101iiiiii10ood000
// ADD #xx,D
// Page 27
type Add_xx struct {
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Add_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014080, func() Instruction { return new(Add_xx) })
}

func (ins *Add_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Add_xx) Disassemble(w TokenWriter) error {
	// Example: add #<$0,a
	return w.Write(
		MnemonicAdd,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood000
// ADD #xxxx,D
// Page 27
type Add_xxxx struct {
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Add_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c0, func() Instruction { return new(Add_xxxx) })
}

func (ins *Add_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *Add_xxxx) Disassemble(w TokenWriter) error {
	// Example: add #>$deface,a
	return w.Write(
		MnemonicAdd,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func170
// 0000000101iiiiii10ood110
// AND #xx,D
// Page 31
type And_xx struct {
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*And_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014086, func() Instruction { return new(And_xx) })
}

func (ins *And_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *And_xx) Disassemble(w TokenWriter) error {
	// Example: and #<$0,a
	return w.Write(
		MnemonicAnd,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood110
// AND #xxxx,D
// Page 31
type And_xxxx struct {
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*And_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c6, func() Instruction { return new(And_xxxx) })
}

func (ins *And_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *And_xxxx) Disassemble(w TokenWriter) error {
	// Example: and #>$deface,a
	return w.Write(
		MnemonicAnd,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func96
// 00000000iiiiiiii101110EE
// AND(I) #xx,D
// Page 33
type Andi struct {
	Immediate         uint32
	ProgramController Register
}

// WordCount implements Instruction
func (*Andi) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff00fc, 0x0000b8, func() Instruction { return new(Andi) })
}

func (ins *Andi) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 8))
	ins.ProgramController = programControlUnitRegister(mask(opcode, 0, 2))

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Andi) Disassemble(w TokenWriter) error {
	// Example: andi #$0,mr
	return w.Write(
		MnemonicAndi,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.ProgramController,
	)
}

// Function Idx: func158
// 0000110000011101SiiiiiiD
// ASL #ii,S2,D
// Page 35
type Asl_ii struct {
	Source      Register
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Asl_ii) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffd00, 0x0c1d00, func() Instruction { return new(Asl_ii) })
}

func (ins *Asl_ii) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 7))
	ins.Immediate = (mask(opcode, 1, 6))
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Asl_ii) Disassemble(w TokenWriter) error {
	// Example: asl #$0,a,a
	return w.Write(
		MnemonicAsl,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func159
// 0000110000011110010SsssD
// ASL S1,S2,D
// Page 35
type Asl_S1S2D struct {
	Source      Register
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Asl_S1S2D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffe0, 0x0c1e40, func() Instruction { return new(Asl_S1S2D) })
}

func (ins *Asl_S1S2D) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Asl_S1S2D) Disassemble(w TokenWriter) error {
	// Example: asl y1,b,b
	return w.Write(
		MnemonicAsl,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func158
// 0000110000011100SiiiiiiD
// ASR #ii,S2,D
// Page 38
type Asr_ii struct {
	Source      Register
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Asr_ii) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffd00, 0x0c1c00, func() Instruction { return new(Asr_ii) })
}

func (ins *Asr_ii) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 7))
	ins.Immediate = (mask(opcode, 1, 6))
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Asr_ii) Disassemble(w TokenWriter) error {
	// Example: asr #$0,a,a
	return w.Write(
		MnemonicAsr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func159
// 0000110000011110011SsssD
// ASR S1,S2,D
// Page 38
type Asr_S1S2D struct {
	Source      Register
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Asr_S1S2D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffe0, 0x0c1e60, func() Instruction { return new(Asr_S1S2D) })
}

func (ins *Asr_S1S2D) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Asr_S1S2D) Disassemble(w TokenWriter) error {
	// Example: asr y1,b,b
	return w.Write(
		MnemonicAsr,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func131
// 00001101000100000100CCCC
// Bcc xxxx
// Page 41
type Bcc_xxxx struct {
	Condition    Condition
	Displacement int32
}

// WordCount implements Instruction
func (*Bcc_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1040, func() Instruction { return new(Bcc_xxxx) })
}

func (ins *Bcc_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 0, 4))
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Bcc_xxxx) Disassemble(w TokenWriter) error {
	// Example: bcc >*-$210532
	return w.Write(
		MnemonicBcc,
		ins.Condition,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func135
// 00000101CCCC01aaaa0aaaaa
// Bcc xxx
// Page 41
type Bcc_xxx struct {
	Condition    Condition
	Displacement int32
}

// WordCount implements Instruction
func (*Bcc_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0c20, 0x050400, func() Instruction { return new(Bcc_xxx) })
}

func (ins *Bcc_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.Displacement = signExtend(9, multimask(opcode, 6, 4, 0, 5))
	return true
}

func (ins *Bcc_xxx) Disassemble(w TokenWriter) error {
	// Example: bcc <*
	return w.Write(
		MnemonicBcc,
		ins.Condition,
		ColumnSeparator,
		ForceShortAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func127
// 0000110100011RRR0100CCCC
// Bcc Rn
// Page 41
type Bcc_Rn struct {
	Address   Register
	Condition Condition
}

// WordCount implements Instruction
func (*Bcc_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1840, func() Instruction { return new(Bcc_Rn) })
}

func (ins *Bcc_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Bcc_Rn) Disassemble(w TokenWriter) error {
	// Example: bcc r0
	return w.Write(
		MnemonicBcc,
		ins.Condition,
		ColumnSeparator,
		ins.Address,
	)
}

// Function Idx: func39
// 0000101101MMMRRR0S0bbbbb
// BCHG #n,[X or Y]:ea
// Page 43
type Bchg_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
}

// WordCount implements Instruction
func (*Bchg_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b4000, func() Instruction { return new(Bchg_ea) })
}

func (ins *Bchg_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bchg_ea) Disassemble(w TokenWriter) error {
	// Example: bchg #$0,x:(r0)-n0
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
	)
}

// Function Idx: func39
// 0000101101MMMRRR0S0bbbbb
// BCHG #n,[X or Y]:ea
// Page 43
type Bchg_ea_Abs struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bchg_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0b4000|(0b110000<<8), func() Instruction { return new(Bchg_ea_Abs) })
}

func (ins *Bchg_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bchg_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceLongAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func43
// 0000101100aaaaaa0S0bbbbb
// BCHG #n,[X or Y]:aa
// Page 43
type Bchg_aa struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bchg_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b0000, func() Instruction { return new(Bchg_aa) })
}

func (ins *Bchg_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bchg_aa) Disassemble(w TokenWriter) error {
	// Example: bchg #$0,x:<$0
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func35
// 0000101110pppppp0S0bbbbb
// BCHG #n,[X or Y]:pp
// Page 43
type Bchg_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bchg_pp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b8000, func() Instruction { return new(Bchg_pp) })
}

func (ins *Bchg_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bchg_pp) Disassemble(w TokenWriter) error {
	// Example: bchg #$0,x:<<$ffffc0
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func186
// 0000000101qqqqqq0S0bbbbb
// BCHG #n,[X or Y]:qq
// Page 43
type Bchg_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bchg_qq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x014000, func() Instruction { return new(Bchg_qq) })
}

func (ins *Bchg_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bchg_qq) Disassemble(w TokenWriter) error {
	// Example: bchg #$0,x:<<$ffff80
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func29
// 0000101111DDDDDD010bbbbb
// BCHG #n,D
// Page 43
type Bchg_D struct {
	Destination Register
	BitNumber   uint32
}

// WordCount implements Instruction
func (*Bchg_D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0e0, 0x0bc040, func() Instruction { return new(Bchg_D) })
}

func (ins *Bchg_D) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	if ins.Destination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Bchg_D) Disassemble(w TokenWriter) error {
	// Example: bchg #$17,lc
	return w.Write(
		MnemonicBchg,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func57
// 0000101001MMMRRR0S0bbbbb
// BCLR #n,[X or Y]:ea
// Page 46
type Bclr_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
}

// WordCount implements Instruction
func (*Bclr_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a4000, func() Instruction { return new(Bclr_ea) })
}

func (ins *Bclr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bclr_ea) Disassemble(w TokenWriter) error {
	// Example: bclr #$0,x:(r0)-n0
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
	)
}

// Function Idx: func57
// 0000101001MMMRRR0S0bbbbb
// BCLR #n,[X or Y]:ea
// Page 46
type Bclr_ea_Abs struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bclr_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0a4000|(0b110000<<8), func() Instruction { return new(Bclr_ea_Abs) })
}

func (ins *Bclr_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bclr_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceLongAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func61
// 0000101000aaaaaa0S0bbbbb
// BCLR #n,[X or Y]:aa
// Page 46
type Bclr_aa struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bclr_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a0000, func() Instruction { return new(Bclr_aa) })
}

func (ins *Bclr_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bclr_aa) Disassemble(w TokenWriter) error {
	// Example: bclr #$0,x:<$0
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func53
// 0000101010pppppp0S0bbbbb
// BCLR #n,[X or Y]:pp
// Page 46
type Bclr_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bclr_pp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a8000, func() Instruction { return new(Bclr_pp) })
}

func (ins *Bclr_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bclr_pp) Disassemble(w TokenWriter) error {
	// Example: bclr #$0,x:<<$ffffc0
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func188
// 0000000100qqqqqq0S0bbbbb
// BCLR #n,[X or Y]:qq
// Page 46
type Bclr_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bclr_qq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x010000, func() Instruction { return new(Bclr_qq) })
}

func (ins *Bclr_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bclr_qq) Disassemble(w TokenWriter) error {
	// Example: bclr #$0,x:<<$ffff80
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func47
// 0000101011DDDDDD010bbbbb
// BCLR #n,D
// Page 46
type Bclr_D struct {
	Destination Register
	BitNumber   uint32
}

// WordCount implements Instruction
func (*Bclr_D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0e0, 0x0ac040, func() Instruction { return new(Bclr_D) })
}

func (ins *Bclr_D) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Destination == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Bclr_D) Disassemble(w TokenWriter) error {
	// Example: bclr #$17,lc
	return w.Write(
		MnemonicBclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func133
// 000011010001000011000000
// BRA xxxx
// Page 49
type Bra_xxxx struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Bra_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff8c0, 0x0d10c0, func() Instruction { return new(Bra_xxxx) })
}

func (ins *Bra_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Bra_xxxx) Disassemble(w TokenWriter) error {
	// Example: bra >*-$210532
	return w.Write(
		MnemonicBra,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func137
// 00000101000011aaaa0aaaaa
// BRA xxx
// Page 49
type Bra_xxx struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Bra_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0c20, 0x050c00, func() Instruction { return new(Bra_xxx) })
}

func (ins *Bra_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(9, multimask(opcode, 6, 4, 0, 5))
	return true
}

func (ins *Bra_xxx) Disassemble(w TokenWriter) error {
	// Example: bra <*
	return w.Write(
		MnemonicBra,
		ColumnSeparator,
		ForceShortAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func129
// 0000110100011RRR11000000
// BRA Rn
// Page 49
type Bra_Rn struct {
	Address Register
}

// WordCount implements Instruction
func (*Bra_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff8c0, 0x0d18c0, func() Instruction { return new(Bra_Rn) })
}

func (ins *Bra_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	return true
}

func (ins *Bra_Rn) Disassemble(w TokenWriter) error {
	// Example: bra r0
	return w.Write(
		MnemonicBra,
		ColumnSeparator,
		ins.Address,
	)
}

// Function Idx: func117
// 0000110010MMMRRR0S0bbbbb
// BRCLR #n,[X or Y]:ea,xxxx
// Page 51
type Brclr_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	Displacement     int32
}

// WordCount implements Instruction
func (*Brclr_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0c8000, func() Instruction { return new(Brclr_ea) })
}
func (ins *Brclr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brclr_ea) Disassemble(w TokenWriter) error {
	// Example: brclr #$0,x:(r0)-n0,>*-$210532
	return w.Write(
		MnemonicBrclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func121
// 0000110010aaaaaa1S0bbbbb
// BRCLR #n,[X or Y]:aa,xxxx
// Page 51
type Brclr_aa struct {
	Address      uint32
	Memory       Memory
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Brclr_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0c8080, func() Instruction { return new(Brclr_aa) })
}
func (ins *Brclr_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brclr_aa) Disassemble(w TokenWriter) error {
	// Example: brclr #$0,x:<$0,>*-$210532
	return w.Write(
		MnemonicBrclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func113
// 0000110011pppppp0S0bbbbb
// BRCLR #n,[X or Y]:pp,xxxx
// Page 51
type Brclr_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Brclr_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0cc000, func() Instruction { return new(Brclr_pp) })
}

func (ins *Brclr_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brclr_pp) Disassemble(w TokenWriter) error {
	// Example: brclr #$0,x:<<$ffffc0,>*-$210532
	return w.Write(
		MnemonicBrclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func192
// 0000010010qqqqqq0S0bbbbb
// BRCLR #n,[X or Y]:qq,xxxx
// Page 51
type Brclr_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Brclr_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x048000, func() Instruction { return new(Brclr_qq) })
}

func (ins *Brclr_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brclr_qq) Disassemble(w TokenWriter) error {
	// Example: brclr #$0,x:<<$ffff80,>*-$210532
	return w.Write(
		MnemonicBrclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func125
// 0000110011DDDDDD100bbbbb
// BRCLR #n,S,xxxx
// Page 51
type Brclr_S struct {
	Source       Register
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Brclr_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0cc080, func() Instruction { return new(Brclr_S) })
}

func (ins *Brclr_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Source == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Brclr_S) Disassemble(w TokenWriter) error {
	// Example: brclr #$17,lc,>*-$210532
	return w.Write(
		MnemonicBrclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func200
// 00000000000000100001CCCC
// BRKcc
// Page 54
type BRKcc struct {
	Condition Condition
}

// WordCount implements Instruction
func (*BRKcc) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x000210, func() Instruction { return new(BRKcc) })
}

func (ins *BRKcc) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *BRKcc) Disassemble(w TokenWriter) error {
	// Example: brkcc
	return w.Write(
		MnemonicBrk,
		ins.Condition,
	)
}

// Function Idx: func116
// 0000110010MMMRRR0S1bbbbb
// BRSET #n,[X or Y]:ea,xxxx
// Page 55
type Brset_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	Displacement     int32
}

// WordCount implements Instruction
func (*Brset_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0c8020, func() Instruction { return new(Brset_ea) })
}

func (ins *Brset_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brset_ea) Disassemble(w TokenWriter) error {
	// Example: brset #$0,x:(r0)-n0,>*-$210532
	return w.Write(
		MnemonicBrset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func120
// 0000110010aaaaaa1S1bbbbb
// BRSET #n,[X or Y]:aa,xxxx
// Page 55
type Brset_aa struct {
	Address      uint32
	Memory       Memory
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Brset_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0c80a0, func() Instruction { return new(Brset_aa) })
}

func (ins *Brset_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brset_aa) Disassemble(w TokenWriter) error {
	// Example: brset #$0,x:<$0,>*-$210532
	return w.Write(
		MnemonicBrset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func112
// 0000110011pppppp0S1bbbbb
// BRSET #n,[X or Y]:pp,xxxx
// Page 55
type Brset_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Brset_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0cc020, func() Instruction { return new(Brset_pp) })
}

func (ins *Brset_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brset_pp) Disassemble(w TokenWriter) error {
	// Example: brset #$0,x:<<$ffffc0,>*-$210532
	return w.Write(
		MnemonicBrset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func191
// 0000010010qqqqqq0S1bbbbb
// BRSET #n,[X or Y]:qq,xxxx
// Page 55
type Brset_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Brset_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x048020, func() Instruction { return new(Brset_qq) })
}

func (ins *Brset_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Brset_qq) Disassemble(w TokenWriter) error {
	// Example: brset #$0,x:<<$ffff80,>*-$210532
	return w.Write(
		MnemonicBrset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func124
// 0000110011DDDDDD101bbbbb
// BRSET #n,S,xxxx
// Page 55
type Brset_S struct {
	Source       Register
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Brset_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0cc0a0, func() Instruction { return new(Brset_S) })
}

func (ins *Brset_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Source == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Brset_S) Disassemble(w TokenWriter) error {
	// Example: brset #$17,lc,>*-$210532
	return w.Write(
		MnemonicBrset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func130
// 00001101000100000000CCCC
// BScc xxxx
// Page 58
type BScc_xxxx struct {
	Condition    Condition
	Displacement int32
}

// WordCount implements Instruction
func (*BScc_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1000, func() Instruction { return new(BScc_xxxx) })
}

func (ins *BScc_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 0, 4))
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *BScc_xxxx) Disassemble(w TokenWriter) error {
	// Example: bscc >*-$210532
	return w.Write(
		MnemonicBs,
		ins.Condition,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func134
// 00000101CCCC00aaaa0aaaaa
// BScc xxx
// Page 58
type BScc_xxx struct {
	Condition    Condition
	Displacement int32
}

// WordCount implements Instruction
func (*BScc_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0c20, 0x050000, func() Instruction { return new(BScc_xxx) })
}

func (ins *BScc_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.Displacement = signExtend(9, multimask(opcode, 6, 4, 0, 5))
	return true
}

func (ins *BScc_xxx) Disassemble(w TokenWriter) error {
	// Example: bscc <*
	return w.Write(
		MnemonicBs,
		ins.Condition,
		ColumnSeparator,
		ForceShortAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func126
// 0000110100011RRR0000CCCC
// BScc Rn
// Page 58
type BScc_Rn struct {
	Address   Register
	Condition Condition
}

// WordCount implements Instruction
func (*BScc_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1800, func() Instruction { return new(BScc_Rn) })
}

func (ins *BScc_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *BScc_Rn) Disassemble(w TokenWriter) error {
	// Example: bscc r0
	return w.Write(
		MnemonicBs,
		ins.Condition,
		ColumnSeparator,
		ins.Address,
	)
}

// Function Idx: func115
// 0000110110MMMRRR0S0bbbbb
// BSCLR #n,[X or Y]:ea,xxxx
// Page 60
type Bsclr_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	Displacement     int32
}

// WordCount implements Instruction
func (*Bsclr_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0d8000, func() Instruction { return new(Bsclr_ea) })
}

func (ins *Bsclr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bsclr_ea) Disassemble(w TokenWriter) error {
	// Example: bsclr #$0,x:(r0)-n0,>*-$210532
	return w.Write(
		MnemonicBsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func119
// 0000110110aaaaaa1S0bbbbb
// BSCLR #n,[X or Y]:aa,xxxx
// Page 60
type Bsclr_aa struct {
	Address      uint32
	Memory       Memory
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Bsclr_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0d8080, func() Instruction { return new(Bsclr_aa) })
}

func (ins *Bsclr_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsclr_aa) Disassemble(w TokenWriter) error {
	// Example: bsclr #$0,x:<$0,>*-$210532
	return w.Write(
		MnemonicBsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func111
// 0000110111pppppp0S0bbbbb
// BSCLR #n,[X or Y]:pp,xxxx
// Page 60
type Bsclr_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Bsclr_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0dc000, func() Instruction { return new(Bsclr_pp) })
}

func (ins *Bsclr_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsclr_pp) Disassemble(w TokenWriter) error {
	// Example: bsclr #$0,x:<<$ffffc0,>*-$210532
	return w.Write(
		MnemonicBsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func190
// 0000010010qqqqqq1S0bbbbb
// BSCLR #n,[X or Y]:qq,xxxx
// Page 60
type Bsclr_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Bsclr_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x048080, func() Instruction { return new(Bsclr_qq) })
}

func (ins *Bsclr_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsclr_qq) Disassemble(w TokenWriter) error {
	// Example: bsclr #$0,x:<<$ffff80,>*-$210532
	return w.Write(
		MnemonicBsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func123
// 0000110111DDDDDD100bbbbb
// BSCLR #n,S,xxxx
// Page 60
type Bsclr_S struct {
	Source       Register
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Bsclr_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0dc080, func() Instruction { return new(Bsclr_S) })
}

func (ins *Bsclr_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)

	if ins.BitNumber >= 24 {
		return false
	}

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Bsclr_S) Disassemble(w TokenWriter) error {
	// Example: bsclr #$17,lc,>*-$210532
	return w.Write(
		MnemonicBsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func56
// 0000101001MMMRRR0S1bbbbb
// BSET #n,[X or Y]:ea
// Page 63
type Bset_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
}

// WordCount implements Instruction
func (*Bset_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a4020, func() Instruction { return new(Bset_ea) })
}

func (ins *Bset_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))

	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bset_ea) Disassemble(w TokenWriter) error {
	// Example: bset #$0,x:(r0)-n0
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
	)
}

// Function Idx: func56
// 0000101001MMMRRR0S1bbbbb
// BSET #n,[X or Y]:ea
// Page 63
type Bset_ea_Abs struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bset_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0a4020|(0b110000<<8), func() Instruction { return new(Bset_ea_Abs) })
}

func (ins *Bset_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Bset_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceLongAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func60
// 0000101000aaaaaa0S1bbbbb
// BSET #n,[X or Y]:aa
// Page 63
type Bset_aa struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Bset_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a0020, func() Instruction { return new(Bset_aa) })
}

func (ins *Bset_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bset_aa) Disassemble(w TokenWriter) error {
	// Example: bset #$0,x:<$0
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func52
// 0000101010pppppp0S1bbbbb
// BSET #n,[X or Y]:pp
// Page 63
type Bset_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bset_pp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0a8020, func() Instruction { return new(Bset_pp) })
}

func (ins *Bset_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bset_pp) Disassemble(w TokenWriter) error {
	// Example: bset #$0,x:<<$ffffc0
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func187
// 0000000100qqqqqq0S1bbbbb
// BSET #n,[X or Y]:qq
// Page 63
type Bset_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Bset_qq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x010020, func() Instruction { return new(Bset_qq) })
}

func (ins *Bset_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bset_qq) Disassemble(w TokenWriter) error {
	// Example: bset #$0,x:<<$ffff80
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func46
// 0000101011DDDDDD011bbbbb
// BSET #n,D
// Page 63
type Bset_D struct {
	Destination Register
	BitNumber   uint32
}

// WordCount implements Instruction
func (*Bset_D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0e0, 0x0ac060, func() Instruction { return new(Bset_D) })
}

func (ins *Bset_D) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Destination == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Bset_D) Disassemble(w TokenWriter) error {
	// Example: bset #$17,lc
	return w.Write(
		MnemonicBset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func132
// 000011010001000010000000
// BSR xxxx
// Page 66
type Bsr_xxxx struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Bsr_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1080, func() Instruction { return new(Bsr_xxxx) })
}

func (ins *Bsr_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Bsr_xxxx) Disassemble(w TokenWriter) error {
	// Example: bsr >*-$210532
	return w.Write(
		MnemonicBsr,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func136
// 00000101000010aaaa0aaaaa
// BSR xxx
// Page 66
type Bsr_xxx struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Bsr_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0c20, 0x050800, func() Instruction { return new(Bsr_xxx) })
}

func (ins *Bsr_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(9, multimask(opcode, 6, 4, 0, 5))
	return true
}

func (ins *Bsr_xxx) Disassemble(w TokenWriter) error {
	// Example: bsr <*
	return w.Write(
		MnemonicBsr,
		ColumnSeparator,
		ForceShortAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func128
// 0000110100011RRR10000000
// BSR Rn
// Page 66
type Bsr_Rn struct {
	Address Register
}

// WordCount implements Instruction
func (*Bsr_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff8c0, 0x0d1880, func() Instruction { return new(Bsr_Rn) })
}

func (ins *Bsr_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	return true
}

func (ins *Bsr_Rn) Disassemble(w TokenWriter) error {
	// Example: bsr r0
	return w.Write(
		MnemonicBsr,
		ColumnSeparator,
		ins.Address,
	)
}

// Function Idx: func114
// 0000110110MMMRRR0S1bbbbb
// BSSET #n,[X or Y]:ea,xxxx
// Page 68
type Bsset_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	Displacement     int32
}

// WordCount implements Instruction
func (*Bsset_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0d8020, func() Instruction { return new(Bsset_ea) })
}

func (ins *Bsset_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsset_ea) Disassemble(w TokenWriter) error {
	// Example: bsset #$0,x:(r0)-n0,>*-$210532
	return w.Write(
		MnemonicBsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func118
// 0000110110aaaaaa1S1bbbbb
// BSSET #n,[X or Y]:aa,xxxx
// Page 68
type Bsset_aa struct {
	Address      uint32
	Memory       Memory
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Bsset_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0d80a0, func() Instruction { return new(Bsset_aa) })
}

func (ins *Bsset_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsset_aa) Disassemble(w TokenWriter) error {
	// Example: bsset #$0,x:<$0,>*-$210532
	return w.Write(
		MnemonicBsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func110
// 0000110111pppppp0S1bbbbb
// BSSET #n,[X or Y]:pp,xxxx
// Page 68
type Bsset_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Bsset_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0dc020, func() Instruction { return new(Bsset_pp) })
}

func (ins *Bsset_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsset_pp) Disassemble(w TokenWriter) error {
	// Example: bsset #$0,x:<<$ffffc0,>*-$210532
	return w.Write(
		MnemonicBsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func189
// 0000010010qqqqqq1S1bbbbb
// BSSET #n,[X or Y]:qq,xxxx
// Page 68
type Bsset_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	Displacement      int32
}

// WordCount implements Instruction
func (*Bsset_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0480a0, func() Instruction { return new(Bsset_qq) })
}

func (ins *Bsset_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Bsset_qq) Disassemble(w TokenWriter) error {
	// Example: bsset #$0,x:<<$ffff80,>*-$210532
	return w.Write(
		MnemonicBsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func122
// 0000110111DDDDDD101bbbbb
// BSSET #n,S,xxxx
// Page 68
type Bsset_S struct {
	Source       Register
	BitNumber    uint32
	Displacement int32
}

// WordCount implements Instruction
func (*Bsset_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0dc0a0, func() Instruction { return new(Bsset_S) })
}

func (ins *Bsset_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Source == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Bsset_S) Disassemble(w TokenWriter) error {
	// Example: bsset #$17,lc,>*-$210532
	return w.Write(
		MnemonicBsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func38
// 0000101101MMMRRR0S1bbbbb
// BTST #n,[X or Y]:ea
// Page 71
type Btst_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
}

// WordCount implements Instruction
func (*Btst_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b4020, func() Instruction { return new(Btst_ea) })
}

func (ins *Btst_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Btst_ea) Disassemble(w TokenWriter) error {
	// Example: btst #$0,x:(r0)-n0
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
	)
}

// Function Idx: func38
// 0000101101MMMRRR0S1bbbbb
// BTST #n,[X or Y]:ea
// Page 71
type Btst_ea_Abs struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Btst_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0b4020|(0b110000<<8), func() Instruction { return new(Btst_ea_Abs) })
}

func (ins *Btst_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Btst_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceLongAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func42
// 0000101100aaaaaa0S1bbbbb
// BTST #n,[X or Y]:aa
// Page 71
type Btst_aa struct {
	Address   uint32
	Memory    Memory
	BitNumber uint32
}

// WordCount implements Instruction
func (*Btst_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b0020, func() Instruction { return new(Btst_aa) })
}

func (ins *Btst_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Btst_aa) Disassemble(w TokenWriter) error {
	// Example: btst #$0,x:<$0
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func34
// 0000101110pppppp0S1bbbbb
// BTST #n,[X or Y]:pp
// Page 71
type Btst_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Btst_pp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0b8020, func() Instruction { return new(Btst_pp) })
}

func (ins *Btst_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Btst_pp) Disassemble(w TokenWriter) error {
	// Example: btst #$0,x:<<$ffffc0
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func185
// 0000000101qqqqqq0S1bbbbb
// BTST #n,[X or Y]:qq
// Page 71
type Btst_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
}

// WordCount implements Instruction
func (*Btst_qq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x014020, func() Instruction { return new(Btst_qq) })
}

func (ins *Btst_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Btst_qq) Disassemble(w TokenWriter) error {
	// Example: btst #$0,x:<<$ffff80
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
	)
}

// Function Idx: func28
// 0000101111DDDDDD011bbbbb
// BTST #n,D
// Page 71
type Btst_D struct {
	Destination Register
	BitNumber   uint32
}

// WordCount implements Instruction
func (*Btst_D) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0e0, 0x0bc060, func() Instruction { return new(Btst_D) })
}

func (ins *Btst_D) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Destination == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Btst_D) Disassemble(w TokenWriter) error {
	// Example: btst #$17,lc
	return w.Write(
		MnemonicBtst,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func167
// 0000110000011110000000SD
// CLB S,D
// Page 73
type Clb struct {
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Clb) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x0c1e00, func() Instruction { return new(Clb) })
}

func (ins *Clb) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 1))
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Clb) Disassemble(w TokenWriter) error {
	// Example: clb a,a
	return w.Write(
		MnemonicClb,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func170
// 0000000101iiiiii10ood101
// CMP #xx, S2
// Page 76
type Cmp_xxS2 struct {
	Immediate uint32
	Source    Register
}

// WordCount implements Instruction
func (*Cmp_xxS2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014085, func() Instruction { return new(Cmp_xxS2) })
}

func (ins *Cmp_xxS2) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Source = accumulator(bit(opcode, 3))
	return true
}

func (ins *Cmp_xxS2) Disassemble(w TokenWriter) error {
	// Example: cmp #<$0,a
	return w.Write(
		MnemonicCmp,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood101
// CMP #xxxx,S2
// Page 76
type Cmp_xxxxS2 struct {
	Source    Register
	Immediate uint32
}

// WordCount implements Instruction
func (*Cmp_xxxxS2) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c5, func() Instruction { return new(Cmp_xxxxS2) })
}

func (ins *Cmp_xxxxS2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *Cmp_xxxxS2) Disassemble(w TokenWriter) error {
	// Example: cmp #>$deface,a
	return w.Write(
		MnemonicCmp,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
	)
}

// Function Idx: func203
// 00001100000111111111gggd
// CMPU S1, S2
// Page 80
type Cmpu_S1S2 struct {
	Source1 Register
	Source2 Register
}

// WordCount implements Instruction
func (*Cmpu_S1S2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x0c1ff0, func() Instruction { return new(Cmpu_S1S2) })
}

func (ins *Cmpu_S1S2) Decode(opcode, extensionWord uint32) bool {
	ins.Source1 = dataALUOperands3(mask(opcode, 1, 3), bit(opcode, 0))
	ins.Source2 = accumulator(bit(opcode, 0))

	if ins.Source1 == RegisterInvalid {
		return false
	}

	if ins.Source2 == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Cmpu_S1S2) Disassemble(w TokenWriter) error {
	// Example: cmpu b,a
	return w.Write(
		MnemonicCmpu,
		ColumnSeparator,
		ins.Source1,
		OperandSeparator,
		ins.Source2,
	)
}

// Function Idx: func94
// 000000000000001000000000
// DEBUG
// Page 81
type Debug struct {
}

// WordCount implements Instruction
func (*Debug) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000200, func() Instruction { return new(Debug) })
}

func (ins *Debug) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Debug) Disassemble(w TokenWriter) error {
	// Example: debug
	return w.Write(
		MnemonicDebug,
	)
}

// Function Idx: func93
// 00000000000000110000CCCC
// DEBUGcc
// Page 82
type Debugcc struct {
	Condition Condition
}

// WordCount implements Instruction
func (*Debugcc) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x000300, func() Instruction { return new(Debugcc) })
}

func (ins *Debugcc) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Debugcc) Disassemble(w TokenWriter) error {
	// Example: debugcc
	return w.Write(
		MnemonicDebug,
		ins.Condition,
	)
}

// Function Idx: func102
// 00000000000000000000101d
// DEC D
// Page 83
type Dec struct {
	Destination Register
}

// WordCount implements Instruction
func (*Dec) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffffe, 0x00000a, func() Instruction { return new(Dec) })
}

func (ins *Dec) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Dec) Disassemble(w TokenWriter) error {
	// Example: dec a
	return w.Write(
		MnemonicDec,
		ColumnSeparator,
		ins.Destination,
	)
}

// Function Idx: func92
// 0000000110oooooo01JJdooo
// DIV S,D
// Page 84
type Div struct {
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Div) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c0, 0x018040, func() Instruction { return new(Div) })
}

func (ins *Div) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUSourceOperands(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Div) Disassemble(w TokenWriter) error {
	// Example: div x0,a
	return w.Write(
		MnemonicDiv,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func173
// 000000010010010s1SdkQQQQ
// DMAC (+/-)S1,S2,D
// Page 87
type Dmac struct {
	Mode        MultiplyMode
	Destination Register
	Sign        bool
	Sources     RegisterPair
}

// WordCount implements Instruction
func (*Dmac) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffe80, 0x012480, func() Instruction { return new(Dmac) })
}

func (ins *Dmac) Decode(opcode, extensionWord uint32) bool {
	ins.Mode = multiplyMode2Bit(multimask(opcode, 8, 1, 6, 1))
	ins.Destination = accumulator(bit(opcode, 5))
	ins.Sign = boolean(bit(opcode, 4))
	ins.Sources = multiplyAllPairs(mask(opcode, 0, 4))
	return true
}

func (ins *Dmac) Disassemble(w TokenWriter) error {
	// Example: dmacss x0,x0,a
	return w.Write(
		MnemonicDmac,
		ins.Mode,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Sources.First,
		OperandSeparator,
		ins.Sources.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func74
// 0000011001MMMRRR0S000000
// DO [X or Y]:ea, expr
// Page 89
type Do_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	LoopAddress      uint32
}

// WordCount implements Instruction
func (*Do_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x064000, func() Instruction { return new(Do_ea) })
}

func (ins *Do_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.LoopAddress = extensionWord
	return true
}

func (ins *Do_ea) Disassemble(w TokenWriter) error {
	// Example: do x:(r0)-n0,$defacf
	return w.Write(
		MnemonicDo,
		ColumnSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		absAddr(MemoryP, ins.LoopAddress+1),
	)
}

// Function Idx: func78
// 0000011000aaaaaa0S000000
// DO [X or Y]:aa, expr
// Page 89
type Do_aa struct {
	Address     uint32
	Memory      Memory
	LoopAddress uint32
}

// WordCount implements Instruction
func (*Do_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x060000, func() Instruction { return new(Do_aa) })
}

func (ins *Do_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.LoopAddress = extensionWord
	return true
}

func (ins *Do_aa) Disassemble(w TokenWriter) error {
	// Example: do x:<$0,$defacf
	return w.Write(
		MnemonicDo,
		ColumnSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		absAddr(MemoryP, ins.LoopAddress+1),
	)
}

// Function Idx: func76
// 00000110iiiiiiii1000hhhh
// DO #xxx, expr
// Page 89
type Do_xxx struct {
	Immediate   uint32
	LoopAddress uint32
}

// WordCount implements Instruction
func (*Do_xxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff00b0, 0x060080, func() Instruction { return new(Do_xxx) })
}

func (ins *Do_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = multimask(opcode, 0, 4, 8, 8)
	ins.LoopAddress = extensionWord
	return true
}

func (ins *Do_xxx) Disassemble(w TokenWriter) error {
	// Example: do #<$0,$defacf
	return w.Write(
		MnemonicDo,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		absAddr(MemoryP, ins.LoopAddress+1),
	)
}

// Function Idx: func72
// 0000011011DDDDDD00000000
// DO S, expr
// Page 89
type Do_S struct {
	Source      Register
	LoopAddress uint32
}

// WordCount implements Instruction
func (*Do_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x06c000, func() Instruction { return new(Do_S) })
}

func (ins *Do_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.LoopAddress = extensionWord

	if ins.Source == RegisterInvalid {
		return false
	}

	if ins.Source == RegisterSSH {
		return false
	}

	return true
}

func (ins *Do_S) Disassemble(w TokenWriter) error {
	// Example: do lc,$defacf
	return w.Write(
		MnemonicDo,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		absAddr(MemoryP, ins.LoopAddress+1),
	)
}

// Function Idx: func201
// 000000000000001000000011
// DO FOREVER
// Page 93
type DoForever struct {
	LoopAddress uint32
}

// WordCount implements Instruction
func (*DoForever) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffffff, 0x000203, func() Instruction { return new(DoForever) })
}

func (ins *DoForever) Decode(opcode, extensionWord uint32) bool {
	ins.LoopAddress = extensionWord
	return true
}

func (ins *DoForever) Disassemble(w TokenWriter) error {
	// Example: do forever,$defacf
	return w.Write(
		MnemonicDoForever,
		OperandSeparator,
		absAddr(MemoryP, ins.LoopAddress+1),
	)
}

// Function Idx: func138
// 0000011001MMMRRR0S010000
// DOR [X or Y]:ea,label
// Page 95
type Dor_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	LoopDisplacement int32
}

// WordCount implements Instruction
func (*Dor_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x064010, func() Instruction { return new(Dor_ea) })
}

func (ins *Dor_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.LoopDisplacement = signExtend(24, extensionWord)
	return true
}

func (ins *Dor_ea) Disassemble(w TokenWriter) error {
	// Example: dor x:(r0)-n0,>*-$210531
	return w.Write(
		MnemonicDor,
		ColumnSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.LoopDisplacement+1),
	)
}

// Function Idx: func140
// 0000011000aaaaaa0S010000
// DOR [X or Y]:aa,label
// Page 95
type Dor_aa struct {
	Address          uint32
	Memory           Memory
	LoopDisplacement int32
}

// WordCount implements Instruction
func (*Dor_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x060010, func() Instruction { return new(Dor_aa) })
}

func (ins *Dor_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.LoopDisplacement = signExtend(24, extensionWord)
	return true
}

func (ins *Dor_aa) Disassemble(w TokenWriter) error {
	// Example: dor x:<$0,>*-$210531
	return w.Write(
		MnemonicDor,
		ColumnSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.LoopDisplacement+1),
	)
}

// Function Idx: func139
// 00000110iiiiiiii1001hhhh
// DOR #xxx, label
// Page 95
type Dor_xxx struct {
	Immediate        uint32
	LoopDisplacement int32
}

// WordCount implements Instruction
func (*Dor_xxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff00b0, 0x060090, func() Instruction { return new(Dor_xxx) })
}

func (ins *Dor_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = multimask(opcode, 0, 4, 8, 8)
	ins.LoopDisplacement = signExtend(24, extensionWord)
	return true
}

func (ins *Dor_xxx) Disassemble(w TokenWriter) error {
	// Example: dor #<$0,>*-$210531
	return w.Write(
		MnemonicDor,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.LoopDisplacement+1),
	)
}

// Function Idx: func141
// 0000011011DDDDDD00010000
// DOR S, label
// Page 95
type Dor_S struct {
	Source           Register
	LoopDisplacement int32
}

// WordCount implements Instruction
func (*Dor_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0b0, 0x06c010, func() Instruction { return new(Dor_S) })
}

func (ins *Dor_S) Decode(opcode, extensionWord uint32) bool {
	ins.LoopDisplacement = signExtend(24, extensionWord)
	ins.Source = register6Bit(mask(opcode, 8, 6))
	if ins.Source == RegisterInvalid {
		return false
	}
	if ins.Source == RegisterSSH {
		return false
	}
	return true
}

func (ins *Dor_S) Disassemble(w TokenWriter) error {
	// Example: dor lc,>*-$210531
	return w.Write(
		MnemonicDor,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.LoopDisplacement+1),
	)
}

// Function Idx: func202
// 000000000000001000000010
// DOR FOREVER
// Page 98
type DorForever struct {
	LoopDisplacement int32
}

// WordCount implements Instruction
func (*DorForever) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffffff, 0x000202, func() Instruction { return new(DorForever) })
}

func (ins *DorForever) Decode(opcode, extensionWord uint32) bool {
	ins.LoopDisplacement = signExtend(24, extensionWord)
	return true
}

func (ins *DorForever) Disassemble(w TokenWriter) error {
	// Example: dor forever,>*-$210531
	return w.Write(
		MnemonicDorForever,
		OperandSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.LoopDisplacement+1),
	)
}

// Function Idx: func97
// 00000000000000001o0o1100
// ENDDO
// Page 100
type Enddo struct {
}

// WordCount implements Instruction
func (*Enddo) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffaf, 0x00008c, func() Instruction { return new(Enddo) })
}

func (ins *Enddo) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Enddo) Disassemble(w TokenWriter) error {
	// Example: enddo
	return w.Write(
		MnemonicEnddo,
	)
}

// Function Idx: func170
// 0000000101iiiiii10ood011
// EOR #xx,D
// Page 101
type Eor_xx struct {
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Eor_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014083, func() Instruction { return new(Eor_xx) })
}

func (ins *Eor_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Eor_xx) Disassemble(w TokenWriter) error {
	// Example: eor #<$0,a
	return w.Write(
		MnemonicEor,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood011
// EOR #xxxx,D
// Page 101
type Eor_xxxx struct {
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Eor_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c3, func() Instruction { return new(Eor_xxxx) })
}

func (ins *Eor_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *Eor_xxxx) Disassemble(w TokenWriter) error {
	// Example: eor #>$deface,a
	return w.Write(
		MnemonicEor,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func162
// 0000110000011010000sSSSD
// EXTRACT S1,S2,D
// Page 103
type Extract_S1S2 struct {
	Source      Register
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Extract_S1S2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x0c1a00, func() Instruction { return new(Extract_S1S2) })
}

func (ins *Extract_S1S2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Extract_S1S2) Disassemble(w TokenWriter) error {
	// Example: extract y1,b,b
	return w.Write(
		MnemonicExtract,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func160
// 0000110000011000000s000D
// EXTRACT #CO,S2,D
// Page 103
type Extract_CoS2 struct {
	Source      Register
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Extract_CoS2) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffff80, 0x0c1800, func() Instruction { return new(Extract_CoS2) })
}

func (ins *Extract_CoS2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Destination = accumulator(bit(opcode, 0))
	ins.Immediate = extensionWord
	return true
}

func (ins *Extract_CoS2) Disassemble(w TokenWriter) error {
	// Example: extract #$deface,a,a
	return w.Write(
		MnemonicExtract,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func162
// 0000110000011010100sSSSD
// EXTRACTU S1,S2,D
// Page 103
type Extractu_S1S2 struct {
	Source      Register
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Extractu_S1S2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x0c1a80, func() Instruction { return new(Extractu_S1S2) })
}

func (ins *Extractu_S1S2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Extractu_S1S2) Disassemble(w TokenWriter) error {
	// Example: extractu y1,b,b
	return w.Write(
		MnemonicExtractu,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func160
// 0000110000011000100s000D
// EXTRACTU #CO,S2,D
// Page 103
type Extractu_CoS2 struct {
	Source      Register
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Extractu_CoS2) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffff80, 0x0c1880, func() Instruction { return new(Extractu_CoS2) })
}

func (ins *Extractu_CoS2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 4))
	ins.Destination = accumulator(bit(opcode, 0))
	ins.Immediate = extensionWord
	return true
}

func (ins *Extractu_CoS2) Disassemble(w TokenWriter) error {
	// Example: extractu #$deface,a,a
	return w.Write(
		MnemonicExtractu,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func105
// 000000000000000000000101
// ILLEGAL
// Page 111
type Illegal struct {
}

// WordCount implements Instruction
func (*Illegal) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000005, func() Instruction { return new(Illegal) })
}

func (ins *Illegal) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Illegal) Disassemble(w TokenWriter) error {
	// Example: illegal
	return w.Write(
		MnemonicIllegal,
	)
}

// Function Idx: func103
// 00000000000000000000100d
// INC D
// Page 113
type Inc struct {
	Destination Register
}

// WordCount implements Instruction
func (*Inc) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffffe, 0x000008, func() Instruction { return new(Inc) })
}

func (ins *Inc) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Inc) Disassemble(w TokenWriter) error {
	// Example: inc a
	return w.Write(
		MnemonicInc,
		ColumnSeparator,
		ins.Destination,
	)
}

// Function Idx: func164
// 00001100000110110qqqSSSD
// INSERT S1,S2,D
// Page 114
type Insert_S1S2 struct {
	Source      Register
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Insert_S1S2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x0c1b00, func() Instruction { return new(Insert_S1S2) })
}

func (ins *Insert_S1S2) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUOperands2(mask(opcode, 4, 3))
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Source == RegisterInvalid {
		return false
	}

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Insert_S1S2) Disassemble(w TokenWriter) error {
	// Example: insert y1,y1,b
	return w.Write(
		MnemonicInsert,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func163
// 00001100000110010qqq000D
// INSERT #CO,S2,D
// Page 114
type Insert_CoS2 struct {
	Source      Register
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Insert_CoS2) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffff80, 0x0c1900, func() Instruction { return new(Insert_CoS2) })
}

func (ins *Insert_CoS2) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = extensionWord
	ins.Source = dataALUOperands2(mask(opcode, 4, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Insert_CoS2) Disassemble(w TokenWriter) error {
	// Example: insert #$deface,y1,b
	return w.Write(
		MnemonicInsert,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func23
// 00001110CCCCaaaaaaaaaaaa
// Jcc xxx
// Page 177
type Jcc_xxx struct {
	Condition   Condition
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jcc_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0000, 0x0e0000, func() Instruction { return new(Jcc_xxx) })
}

func (ins *Jcc_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.JumpAddress = (mask(opcode, 0, 12))
	return true
}

func (ins *Jcc_xxx) Disassemble(w TokenWriter) error {
	// Example: jcc <$0
	return w.Write(
		MnemonicJcc,
		ins.Condition,
		ColumnSeparator,
		ForceShortAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func44
// 0000101011MMMRRR1010CCCC
// Jcc ea
// Page 117
type Jcc_ea struct {
	EffectiveAddress EffectiveAddress
	Condition        Condition
}

// WordCount implements Instruction
func (*Jcc_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0ac0a0, func() Instruction { return new(Jcc_ea) })
}

func (ins *Jcc_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Jcc_ea) Disassemble(w TokenWriter) error {
	// Example: jcc (r0)-n0
	return w.Write(
		MnemonicJcc,
		ins.Condition,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func44
// 0000101011MMMRRR1010CCCC
// Jcc ea
// Page 117
type Jcc_ea_Abs struct {
	JumpAddress uint32
	Condition   Condition
}

// WordCount implements Instruction
func (*Jcc_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0ac0a0|(0b110000<<8), func() Instruction { return new(Jcc_ea_Abs) })
}

func (ins *Jcc_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = extensionWord
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Jcc_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicJcc,
		ins.Condition,
		ColumnSeparator,
		ForceLongAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func55, func152
// 0000101001MMMRRR1S0bbbbb
// JCLR #n,[X or Y]:ea,xxxx
// Page 119
type Jclr_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	JumpAddress      uint32
}

// WordCount implements Instruction
func (*Jclr_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a4080, func() Instruction { return new(Jclr_ea) })
}

func (ins *Jclr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jclr_ea) Disassemble(w TokenWriter) error {
	// Example: jclr #$0,x:(r0)-n0,$deface
	return w.Write(
		MnemonicJclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func59
// 0000101000aaaaaa1S0bbbbb
// JCLR #n,[X or Y]:aa,xxxx
// Page 119
type Jclr_aa struct {
	Address     uint32
	Memory      Memory
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jclr_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a0080, func() Instruction { return new(Jclr_aa) })
}

func (ins *Jclr_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jclr_aa) Disassemble(w TokenWriter) error {
	// Example: jclr #$0,x:<$0,$deface
	return w.Write(
		MnemonicJclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func51
// 0000101010pppppp1S0bbbbb
// JCLR #n,[X or Y]:pp,xxxx
// Page 119
type Jclr_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jclr_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a8080, func() Instruction { return new(Jclr_pp) })
}

func (ins *Jclr_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jclr_pp) Disassemble(w TokenWriter) error {
	// Example: jclr #$0,x:<<$ffffc0,$deface
	return w.Write(
		MnemonicJclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func184
// 0000000110qqqqqq1S0bbbbb
// JCLR #n,[X or Y]:qq,xxxx
// Page 119
type Jclr_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jclr_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x018080, func() Instruction { return new(Jclr_qq) })
}

func (ins *Jclr_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jclr_qq) Disassemble(w TokenWriter) error {
	// Example: jclr #$0,x:<<$ffff80,$deface
	return w.Write(
		MnemonicJclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func49
// 0000101011DDDDDD000bbbbb
// JCLR #n,S,xxxx
// Page 119
type Jclr_S struct {
	Source      Register
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jclr_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0e0, 0x0ac000, func() Instruction { return new(Jclr_S) })
}

func (ins *Jclr_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Jclr_S) Disassemble(w TokenWriter) error {
	// Example: jclr #$17,lc,$deface
	return w.Write(
		MnemonicJclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func45
// 0000101011MMMRRR10000000
// JMP ea
// Page 121
type Jmp_ea struct {
	EffectiveAddress EffectiveAddress
}

// WordCount implements Instruction
func (*Jmp_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0ac080, func() Instruction { return new(Jmp_ea) })
}

func (ins *Jmp_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	return true
}

func (ins *Jmp_ea) Disassemble(w TokenWriter) error {
	// Example: jmp (r0)-n0
	return w.Write(
		MnemonicJmp,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func45
// 0000101011MMMRRR10000000
// JMP ea
// Page 121
type Jmp_ea_Abs struct {
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jmp_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0ac080|(0b110000<<8), func() Instruction { return new(Jmp_ea_Abs) })
}

func (ins *Jmp_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = extensionWord
	return true
}

func (ins *Jmp_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicJmp,
		ColumnSeparator,
		ForceLongAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func25
// 000011000000aaaaaaaaaaaa
// JMP xxx
// Page 121
type Jmp_xxx struct {
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jmp_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff000, 0x0c0000, func() Instruction { return new(Jmp_xxx) })
}

func (ins *Jmp_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = (mask(opcode, 0, 12))
	return true
}

func (ins *Jmp_xxx) Disassemble(w TokenWriter) error {
	// Example: jmp <$0
	return w.Write(
		MnemonicJmp,
		ColumnSeparator,
		ForceShortAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func22
// 00001111CCCCaaaaaaaaaaaa
// JScc xxx
// Page 122
type Jscc_xxx struct {
	Condition   Condition
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jscc_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0000, 0x0f0000, func() Instruction { return new(Jscc_xxx) })
}

func (ins *Jscc_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.JumpAddress = (mask(opcode, 0, 12))
	return true
}

func (ins *Jscc_xxx) Disassemble(w TokenWriter) error {
	// Example: jscc <$0
	return w.Write(
		MnemonicJs,
		ins.Condition,
		ColumnSeparator,
		ForceShortAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func26
// 0000101111MMMRRR1010CCCC
// JScc ea
// Page 122
type Jscc_ea struct {
	EffectiveAddress EffectiveAddress
	Condition        Condition
}

// WordCount implements Instruction
func (*Jscc_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0bc0a0, func() Instruction { return new(Jscc_ea) })
}

func (ins *Jscc_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Jscc_ea) Disassemble(w TokenWriter) error {
	// Example: jscc (r0)-n0
	return w.Write(
		MnemonicJs,
		ins.Condition,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func26
// 0000101111MMMRRR1010CCCC
// JScc ea
// Page 122
type Jscc_ea_Abs struct {
	JumpAddress uint32
	Condition   Condition
}

// WordCount implements Instruction
func (*Jscc_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0bc0a0|(0b110000<<8), func() Instruction { return new(Jscc_ea_Abs) })
}

func (ins *Jscc_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = extensionWord
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Jscc_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicJs,
		ins.Condition,
		ColumnSeparator,
		ForceLongAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func37, func152
// 0000101101MMMRRR1S0bbbbb
// JSCLR #n,[X or Y]:ea,xxxx
// Page 124
type Jsclr_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	JumpAddress      uint32
}

// WordCount implements Instruction
func (*Jsclr_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b4080, func() Instruction { return new(Jsclr_ea) })
}

func (ins *Jsclr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jsclr_ea) Disassemble(w TokenWriter) error {
	// Example: jsclr #$0,x:(r0)-n0,$deface
	return w.Write(
		MnemonicJsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func41
// 0000101100aaaaaa1S0bbbbb
// JSCLR #n,[X or Y]:aa,xxxx
// Page 124
type Jsclr_aa struct {
	Address     uint32
	Memory      Memory
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsclr_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b0080, func() Instruction { return new(Jsclr_aa) })
}

func (ins *Jsclr_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jsclr_aa) Disassemble(w TokenWriter) error {
	// Example: jsclr #$0,x:<$0,$deface
	return w.Write(
		MnemonicJsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func33
// 0000101110pppppp1S0bbbbb
// JSCLR #n,[X or Y]:pp,xxxx
// Page 124
type Jsclr_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jsclr_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b8080, func() Instruction { return new(Jsclr_pp) })
}

func (ins *Jsclr_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	return true
}

func (ins *Jsclr_pp) Disassemble(w TokenWriter) error {
	// Example: jsclr #$0,x:<<$ffffc0,$deface
	return w.Write(
		MnemonicJsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func182
// 0000000111qqqqqq1S0bbbbb
// JSCLR #n,[X or Y]:qq,xxxx
// Page 124
type Jsclr_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jsclr_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x01c080, func() Instruction { return new(Jsclr_qq) })
}

func (ins *Jsclr_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jsclr_qq) Disassemble(w TokenWriter) error {
	// Example: jsclr #$0,x:<<$ffff80,$deface
	return w.Write(
		MnemonicJsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func31
// 0000101111DDDDDD000bbbbb
// JSCLR #n,S,xxxx
// Page 124
type Jsclr_S struct {
	Source      Register
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsclr_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0e0, 0x0bc000, func() Instruction { return new(Jsclr_S) })
}

func (ins *Jsclr_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Jsclr_S) Disassemble(w TokenWriter) error {
	// Example: jsclr #$17,lc,$deface
	return w.Write(
		MnemonicJsclr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func54, func152
// 0000101001MMMRRR1S1bbbbb
// JSET #n,[X or Y]:ea,xxxx
// Page 127
type Jset_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	JumpAddress      uint32
}

// WordCount implements Instruction
func (*Jset_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a40a0, func() Instruction { return new(Jset_ea) })
}

func (ins *Jset_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jset_ea) Disassemble(w TokenWriter) error {
	// Example: jset #$0,x:(r0)-n0,$deface
	return w.Write(
		MnemonicJset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func58
// 0000101000aaaaaa1S1bbbbb
// JSET #n,[X or Y]:aa,xxxx
// Page 127
type Jset_aa struct {
	Address     uint32
	Memory      Memory
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jset_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a00a0, func() Instruction { return new(Jset_aa) })
}

func (ins *Jset_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jset_aa) Disassemble(w TokenWriter) error {
	// Example: jset #$0,x:<$0,$deface
	return w.Write(
		MnemonicJset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func50
// 0000101010pppppp1S1bbbbb
// JSET #n,[X or Y]:pp,xxxx
// Page 127
type Jset_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jset_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0a80a0, func() Instruction { return new(Jset_pp) })
}

func (ins *Jset_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jset_pp) Disassemble(w TokenWriter) error {
	// Example: jset #$0,x:<<$ffffc0,$deface
	return w.Write(
		MnemonicJset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func183
// 0000000110qqqqqq1S1bbbbb
// JSET #n,[X or Y]:qq,xxxx
// Page 127
type Jset_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jset_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0180a0, func() Instruction { return new(Jset_qq) })
}

func (ins *Jset_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jset_qq) Disassemble(w TokenWriter) error {
	// Example: jset #$0,x:<<$ffff80,$deface
	return w.Write(
		MnemonicJset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func48
// 0000101011DDDDDD001bbbbb
// JSET #n,S,xxxx
// Page 127
type Jset_S struct {
	Source      Register
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jset_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0e0, 0x0ac020, func() Instruction { return new(Jset_S) })
}

func (ins *Jset_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Source == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Jset_S) Disassemble(w TokenWriter) error {
	// Example: jset #$17,lc,$deface
	return w.Write(
		MnemonicJset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func27
// 0000101111MMMRRR10000000
// JSR ea
// Page 129
type Jsr_ea struct {
	EffectiveAddress EffectiveAddress
}

// WordCount implements Instruction
func (*Jsr_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x0bc080, func() Instruction { return new(Jsr_ea) })
}

func (ins *Jsr_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	return true
}

func (ins *Jsr_ea) Disassemble(w TokenWriter) error {
	// Example: jsr (r0)-n0
	return w.Write(
		MnemonicJsr,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func27
// 0000101111MMMRRR10000000
// JSR ea
// Page 129
type Jsr_ea_Abs struct {
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsr_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0|(0b111111<<8), 0x0bc080|(0b110000<<8), func() Instruction { return new(Jsr_ea_Abs) })
}

func (ins *Jsr_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = extensionWord
	return true
}

func (ins *Jsr_ea_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicJsr,
		ColumnSeparator,
		ForceLongAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func24
// 000011010000aaaaaaaaaaaa
// JSR xxx
// Page 129
type Jsr_xxx struct {
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsr_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff000, 0x0d0000, func() Instruction { return new(Jsr_xxx) })
}

func (ins *Jsr_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.JumpAddress = (mask(opcode, 0, 12))
	return true
}

func (ins *Jsr_xxx) Disassemble(w TokenWriter) error {
	// Example: jsr <$0
	return w.Write(
		MnemonicJsr,
		ColumnSeparator,
		ForceShortAddressMode,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func36, func152
// 0000101101MMMRRR1S1bbbbb
// JSSET #n,[X or Y]:ea,xxxx
// Page 131
type Jsset_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
	BitNumber        uint32
	JumpAddress      uint32
}

// WordCount implements Instruction
func (*Jsset_ea) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b40a0, func() Instruction { return new(Jsset_ea) })
}

func (ins *Jsset_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jsset_ea) Disassemble(w TokenWriter) error {
	// Example: jsset #$0,x:(r0)-n0,$deface
	return w.Write(
		MnemonicJsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ins.EffectiveAddress,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func40
// 0000101100aaaaaa1S1bbbbb
// JSSET #n,[X or Y]:aa,xxxx
// Page 131
type Jsset_aa struct {
	Address     uint32
	Memory      Memory
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsset_aa) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b00a0, func() Instruction { return new(Jsset_aa) })
}

func (ins *Jsset_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jsset_aa) Disassemble(w TokenWriter) error {
	// Example: jsset #$0,x:<$0,$deface
	return w.Write(
		MnemonicJsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func32
// 0000101110pppppp1S1bbbbb
// JSSET #n,[X or Y]:pp,xxxx
// Page 131
type Jsset_pp struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jsset_pp) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x0b80a0, func() Instruction { return new(Jsset_pp) })
}

func (ins *Jsset_pp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = highPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jsset_pp) Disassemble(w TokenWriter) error {
	// Example: jsset #$0,x:<<$ffffc0,$deface
	return w.Write(
		MnemonicJsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func181
// 0000000111qqqqqq1S1bbbbb
// JSSET #n,[X or Y]:qq,xxxx
// Page 131
type Jsset_qq struct {
	PeripheralAddress uint32
	Memory            Memory
	BitNumber         uint32
	JumpAddress       uint32
}

// WordCount implements Instruction
func (*Jsset_qq) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0a0, 0x01c0a0, func() Instruction { return new(Jsset_qq) })
}

func (ins *Jsset_qq) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	return true
}

func (ins *Jsset_qq) Disassemble(w TokenWriter) error {
	// Example: jsset #$0,x:<<$ffff80,$deface
	return w.Write(
		MnemonicJsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Memory,
		ForceIOShortAddressMode,
		absAddr(ins.Memory, ins.PeripheralAddress),
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func30
// 0000101111DDDDDD001bbbbb
// JSSET #n,S,xxxx
// Page 131
type Jsset_S struct {
	Source      Register
	BitNumber   uint32
	JumpAddress uint32
}

// WordCount implements Instruction
func (*Jsset_S) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0e0, 0x0bc020, func() Instruction { return new(Jsset_S) })
}

func (ins *Jsset_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))
	ins.BitNumber = (mask(opcode, 0, 5))
	ins.JumpAddress = extensionWord

	if ins.BitNumber >= 24 {
		return false
	}
	if ins.Source == RegisterInvalid {
		return false
	}
	return true
}

func (ins *Jsset_S) Disassemble(w TokenWriter) error {
	// Example: jsset #$17,lc,$deface
	return w.Write(
		MnemonicJsset,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.BitNumber),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		absAddr(MemoryP, ins.JumpAddress),
	)
}

// Function Idx: func144
// 0000010011000RRR000ddddd
// LRA Rn,D
// Page 134
type Lra_Rn struct {
	Address            Register
	DestinationAddress Register
}

// WordCount implements Instruction
func (*Lra_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0e0, 0x04c000, func() Instruction { return new(Lra_Rn) })
}

func (ins *Lra_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.DestinationAddress = register5BitUndocumented(mask(opcode, 0, 5))
	return true
}

func (ins *Lra_Rn) Disassemble(w TokenWriter) error {
	// Example: lra r0,x0
	return w.Write(
		MnemonicLra,
		ColumnSeparator,
		ins.Address,
		OperandSeparator,
		ins.DestinationAddress,
	)
}

// Function Idx: func145
// 0000010001oooooo010ddddd
// LRA xxxx,D
// Page 134
type Lra_xxxx struct {
	DestinationAddress Register
	Displacement       int32
}

// WordCount implements Instruction
func (*Lra_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0e0, 0x044040, func() Instruction { return new(Lra_xxxx) })
}

func (ins *Lra_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.DestinationAddress = register5BitUndocumented(mask(opcode, 0, 5))
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Lra_xxxx) Disassemble(w TokenWriter) error {
	// Example: lra >*-$210532,x0
	return w.Write(
		MnemonicLra,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
		OperandSeparator,
		ins.DestinationAddress,
	)
}

// Function Idx: func161
// 000011000001111010iiiiiD
// LSL #ii,D
// Page 136
type Lsl_ii struct {
	ShiftAmount uint32
	Destination Register
}

// WordCount implements Instruction
func (*Lsl_ii) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffc0, 0x0c1e80, func() Instruction { return new(Lsl_ii) })
}

func (ins *Lsl_ii) Decode(opcode, extensionWord uint32) bool {
	ins.ShiftAmount = (mask(opcode, 1, 5))
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Lsl_ii) Disassemble(w TokenWriter) error {
	// Example: lsl #$0,a
	return w.Write(
		MnemonicLsl,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.ShiftAmount),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func169
// 00001100000111100001sssD
// LSL S,D
// Page 136
type Lsl_SD struct {
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Lsl_SD) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x0c1e10, func() Instruction { return new(Lsl_SD) })
}

func (ins *Lsl_SD) Decode(opcode, extensionWord uint32) bool {
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Lsl_SD) Disassemble(w TokenWriter) error {
	// Example: lsl y1,b
	return w.Write(
		MnemonicLsl,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func161
// 000011000001111011iiiiiD
// LSR #ii,D
// Page 139
type Lsr_ii struct {
	ShiftAmount uint32
	Destination Register
}

// WordCount implements Instruction
func (*Lsr_ii) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffc0, 0x0c1ec0, func() Instruction { return new(Lsr_ii) })
}

func (ins *Lsr_ii) Decode(opcode, extensionWord uint32) bool {
	ins.ShiftAmount = (mask(opcode, 1, 5))
	ins.Destination = accumulator(bit(opcode, 0))
	return true
}

func (ins *Lsr_ii) Disassemble(w TokenWriter) error {
	// Example: lsr #$0,a
	return w.Write(
		MnemonicLsr,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.ShiftAmount),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func169
// 00001100000111100011sssD
// LSR S,D
// Page 139
type Lsr_SD struct {
	Control     Register
	Destination Register
}

// WordCount implements Instruction
func (*Lsr_SD) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x0c1e30, func() Instruction { return new(Lsr_SD) })
}

func (ins *Lsr_SD) Decode(opcode, extensionWord uint32) bool {
	ins.Control = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Control == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Lsr_SD) Disassemble(w TokenWriter) error {
	// Example: lsr y1,b
	return w.Write(
		MnemonicLsr,
		ColumnSeparator,
		ins.Control,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func88, func87
// 00000100010MMRRR000ddddd
// LUA/LEA ea,D
// Page 142
type Lua_ea struct {
	EffectiveAddress   EffectiveAddress
	DestinationAddress Register
}

// WordCount implements Instruction
func (*Lua_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff60a8, 0x044000, func() Instruction { return new(Lua_ea) })
	registerInstruction(0xff60a8, 0x044008, func() Instruction { return new(Lua_ea) })
}

func (ins *Lua_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 2), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.DestinationAddress = register5BitUndocumented(mask(opcode, 0, 5))
	return true
}

func (ins *Lua_ea) Disassemble(w TokenWriter) error {
	// Example: lua (r0)-n0,x0
	return w.Write(
		MnemonicLua,
		ColumnSeparator,
		ins.EffectiveAddress,
		OperandSeparator,
		ins.DestinationAddress,
	)
}

// Function Idx: func143, func142
// 0000010000aaaRRRaaaadddd
// LUA/LEA (Rn + aa),D
// Page 142
type Lua_Rn struct {
	Address            Register
	Offset             int32
	DestinationAddress Register
}

// WordCount implements Instruction
func (*Lua_Rn) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc000, 0x040000, func() Instruction { return new(Lua_Rn) })
}

func (ins *Lua_Rn) Decode(opcode, extensionWord uint32) bool {
	ins.Offset = signExtend(7, multimask(opcode, 11, 3, 4, 4))
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.DestinationAddress = addressAndOffsetRegister(mask(opcode, 0, 4))
	return true
}

func (ins *Lua_Rn) Disassemble(w TokenWriter) error {
	// Example: lua (r0+$0),r0
	return w.Write(
		MnemonicLua,
		ColumnSeparator,
		OpenBrace,
		ins.Address,
		numberToken{ins.Offset, true},
		CloseBrace,
		OperandSeparator,
		ins.DestinationAddress,
	)
}

// Function Idx: func109
// 00000001000sssss11QQdk10
// MAC (+/-)S,#n,D
// Page 144
type Mac_S struct {
	Immediate   uint32
	Source      Register
	Destination Register
	Sign        bool
}

// WordCount implements Instruction
func (*Mac_S) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffe0c3, 0x0100c2, func() Instruction { return new(Mac_S) })
}

func (ins *Mac_S) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 5))
	ins.Source = dataALUMultiplyOperands1(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	return true
}

func (ins *Mac_S) Disassemble(w TokenWriter) error {
	// Example: mac y1,#$0,a
	return w.Write(
		MnemonicMac,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source,
		OperandSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func155
// 0000000101ooooo111qqdk10
// MACI (+/-)#xxxx,S,D
// Page 146
type Maci_xxxx struct {
	Source      Register
	Destination Register
	Sign        bool
	Immediate   uint32
}

// WordCount implements Instruction
func (*Maci_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c3, 0x0141c2, func() Instruction { return new(Maci_xxxx) })
}

func (ins *Maci_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUMultiplyOperands2(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	ins.Immediate = extensionWord
	return true
}

func (ins *Maci_xxxx) Disassemble(w TokenWriter) error {
	// Example: maci #>$deface,x0,a
	return w.Write(
		MnemonicMaci,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ImmediateAddressMode,
		ForceLongAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func174
// 00000001001001101sdkQQQQ
// MACsu (+/-)S1,S2,D / MACuu (+/-)S1,S2,D
// Page 147
type Macsu struct {
	Mode        MultiplyMode
	Destination Register
	Sign        bool
	Sources     RegisterPair
}

// WordCount implements Instruction
func (*Macsu) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x012680, func() Instruction { return new(Macsu) })
}

func (ins *Macsu) Decode(opcode, extensionWord uint32) bool {
	ins.Mode = multiplyMode1Bit(bit(opcode, 6))
	ins.Destination = accumulator(bit(opcode, 5))
	ins.Sign = boolean(bit(opcode, 4))
	ins.Sources = multiplyAllPairs(mask(opcode, 0, 4))
	return true
}

func (ins *Macsu) Disassemble(w TokenWriter) error {
	// Example: macsu x0,x0,a
	return w.Write(
		MnemonicMac,
		ins.Mode,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Sources.First,
		OperandSeparator,
		ins.Sources.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func109
// 00000001000sssss11QQdk11
// MACR (+/-)S,#n,D
// Page 148
type Macr_S struct {
	Immediate   uint32
	Source      Register
	Destination Register
	Sign        bool
}

// WordCount implements Instruction
func (*Macr_S) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffe0c3, 0x0100c3, func() Instruction { return new(Macr_S) })
}

func (ins *Macr_S) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 5))
	ins.Source = dataALUMultiplyOperands1(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	return true
}

func (ins *Macr_S) Disassemble(w TokenWriter) error {
	// Example: macr y1,#$0,a
	return w.Write(
		MnemonicMacr,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source,
		OperandSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func155
// 0000000101ooooo111qqdk11
// MACRI (+/-)#xxxx,S,D
// Page 150
type Macri_xxxx struct {
	Source      Register
	Destination Register
	Sign        bool
	Immediate   uint32
}

// WordCount implements Instruction
func (*Macri_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c3, 0x0141c3, func() Instruction { return new(Macri_xxxx) })
}

func (ins *Macri_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUMultiplyOperands2(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	ins.Immediate = extensionWord
	return true
}

func (ins *Macri_xxxx) Disassemble(w TokenWriter) error {
	// Example: macri #>$deface,x0,a
	return w.Write(
		MnemonicMacri,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ImmediateAddressMode,
		ForceLongAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func166
// 00001100000110111000SSSD
// MERGE S,D
// Page 154
type Merge struct {
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Merge) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x0c1b80, func() Instruction { return new(Merge) })
}

func (ins *Merge) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Merge) Disassemble(w TokenWriter) error {
	// Example: merge y1,b
	return w.Write(
		MnemonicMerge,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func152
// 0000101001110RRR1WDDDDDD
// MOVE X:(Rn + xxxx),D / MOVE S,X:(Rn + xxxx)
// Page 163
type Movex_Rnxxxx struct {
	Address             Register
	IsWrite             bool
	SourceOrDestination Register
	Offset              uint32
}

// WordCount implements Instruction
func (*Movex_Rnxxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff880, 0x0a7080, func() Instruction { return new(Movex_Rnxxxx) })
}

func (ins *Movex_Rnxxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.IsWrite = boolean(bit(opcode, 6))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 6))
	ins.Offset = extensionWord

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movex_Rnxxxx) Disassemble(w TokenWriter) error {
	// Example: move x:(r7+>$deface),lc
	// Example: move n5,x:(r5+>$deface)
	if ins.IsWrite {
		// Using ForceLongAddressMode here instead of ForceLongImmediateAddressMode
		// The immediate seems to be implied and causes an assembly error
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryX, OpenBrace, ins.Address, Plus, ForceLongAddressMode, immediate(ins.Offset), CloseBrace,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Using ForceLongAddressMode here instead of ForceLongImmediateAddressMode
	// The immediate seems to be implied and causes an assembly error
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, OpenBrace, ins.Address, Plus, ForceLongAddressMode, immediate(ins.Offset), CloseBrace,
	)
}

// Function Idx: func151
// 0000001aaaaaaRRR1a0WDDDD
// MOVE X:(Rn + xxx),D / MOVE S,X:(Rn + xxx)
// Page 163
type Movex_Rnxxx struct {
	Offset              int32
	Address             Register
	IsWrite             bool
	SourceOrDestination Register
}

// WordCount implements Instruction
func (*Movex_Rnxxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfe00a0, 0x020080, func() Instruction { return new(Movex_Rnxxx) })
}

func (ins *Movex_Rnxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Offset = signExtend(7, multimask(opcode, 11, 6, 6, 1))
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.IsWrite = boolean(bit(opcode, 4))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 4))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movex_Rnxxx) Disassemble(w TokenWriter) error {
	// Example: move x:(r7-<$1),b
	// Example: move y1,x:(r2-<$5)
	if ins.IsWrite {
		// Using ForceShortAddressMode here instead of ForceShortImmediateAddressMode
		// The immediate seems to be implied and causes an assembly error
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryX, OpenBrace, ins.Address,
			ifThen(ins.Offset < 0, Minus, Plus),
			ForceShortAddressMode,
			immediate(abs(ins.Offset)),
			CloseBrace,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Using ForceShortAddressMode here instead of ForceShortImmediateAddressMode
	// The immediate seems to be implied and causes an assembly error
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryX, OpenBrace, ins.Address,
		ifThen(ins.Offset < 0, Minus, Plus),
		ForceShortAddressMode,
		immediate(abs(ins.Offset)),
		CloseBrace,
	)
}

// Function Idx: func152
// 0000101101110RRR1WDDDDDD
// MOVE Y:(Rn + xxxx),D / MOVE D,Y:(Rn + xxxx)
// Page 168
type Movey_Rnxxxx struct {
	Address             Register
	IsWrite             bool
	SourceOrDestination Register
	Offset              uint32
}

// WordCount implements Instruction
func (*Movey_Rnxxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfff880, 0x0b7080, func() Instruction { return new(Movey_Rnxxxx) })
}

func (ins *Movey_Rnxxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.IsWrite = boolean(bit(opcode, 6))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 6))
	ins.Offset = extensionWord

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movey_Rnxxxx) Disassemble(w TokenWriter) error {
	// Example: move n5,y:(r5+$deface)
	// Example: move y:(r7+$deface),lc
	if ins.IsWrite {
		// Using ForceLongAddressMode here instead of ForceLongImmediateAddressMode
		// The immediate seems to be implied and causes an assembly error
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryY, OpenBrace, ins.Address, Plus, ForceLongAddressMode, immediate(ins.Offset), CloseBrace,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Using ForceLongAddressMode here instead of ForceLongImmediateAddressMode
	// The immediate seems to be implied and causes an assembly error
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, OpenBrace, ins.Address, Plus, ForceLongAddressMode, immediate(ins.Offset), CloseBrace,
	)
}

// Function Idx: func151
// 0000001aaaaaaRRR1a1WDDDD
// MOVE Y:(Rn + xxx),D / MOVE D,Y:(Rn + xxx)
// Page 168
type Movey_Rnxxx struct {
	Address             Register
	Offset              int32
	IsWrite             bool
	SourceOrDestination Register
}

// WordCount implements Instruction
func (*Movey_Rnxxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfe00a0, 0x0200a0, func() Instruction { return new(Movey_Rnxxx) })
}

func (ins *Movey_Rnxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Offset = signExtend(7, multimask(opcode, 11, 6, 6, 1))
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.IsWrite = boolean(bit(opcode, 4))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 4))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movey_Rnxxx) Disassemble(w TokenWriter) error {
	// Example: move y1,y:(r2-$5)
	// Example: move y:(r7-$1),b
	if ins.IsWrite {
		// Using ForceShortAddressMode here instead of ForceShortImmediateAddressMode
		// The immediate seems to be implied and causes an assembly error
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryY, OpenBrace, ins.Address,
			ifThen(ins.Offset < 0, Minus, Plus),
			ForceShortAddressMode,
			immediate(abs(ins.Offset)),
			CloseBrace,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Using ForceShortAddressMode here instead of ForceShortImmediateAddressMode
	// The immediate seems to be implied and causes an assembly error
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryY, OpenBrace, ins.Address,
		ifThen(ins.Offset < 0, Minus, Plus),
		ForceShortAddressMode,
		immediate(abs(ins.Offset)),
		CloseBrace,
	)
}

// Function Idx: func80, func79, func194, func195
// 00000101W1MMMRRR0S1DDDDD
// MOVE(C) [X or Y]:ea,D1 / MOVE(C) S1,[X or Y]:ea / MOVE(C) #xxxx,D1
// Page 178
type Movec_ea struct {
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	Memory            Memory
	ProgramController Register
}

// WordCount implements Instruction
func (*Movec_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff40a0, 0x054020, func() Instruction { return new(Movec_ea) })
}

func (ins *Movec_ea) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movec_ea) Disassemble(w TokenWriter) error {
	// Example: move m0,x:(r0)-n0
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			ins.Memory, ins.EffectiveAddress,
			OperandSeparator,
			ins.ProgramController,
		)
	}

	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.ProgramController,
		OperandSeparator,
		ins.Memory, ins.EffectiveAddress,
	)
}

// Function Idx: func80, func79, func194, func195
// 00000101W1MMMRRR0S1DDDDD
// MOVE(C) [X or Y]:ea,D1 / MOVE(C) S1,[X or Y]:ea / MOVE(C) #xxxx,D1
// Page 178
type Movec_ea_Abs struct {
	IsWrite           bool
	Memory            Memory
	Address           uint32
	ProgramController Register
}

// WordCount implements Instruction
func (*Movec_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff40a0|(0b111111<<8), 0x054020|(0b110000<<8), func() Instruction { return new(Movec_ea_Abs) })
}

func (ins *Movec_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))
	ins.Address = extensionWord

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movec_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
			OperandSeparator,
			ins.ProgramController,
		)
	}

	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.ProgramController,
		OperandSeparator,
		ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func80, func79, func194, func195
// 00000101W1MMMRRR0S1DDDDD
// MOVE(C) [X or Y]:ea,D1 / MOVE(C) S1,[X or Y]:ea / MOVE(C) #xxxx,D1
// Page 178
type Movec_ea_Imm struct {
	Memory            Memory
	Immediate         uint32
	ProgramController Register
}

// WordCount implements Instruction
func (*Movec_ea_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff40a0|(0b111111<<8)|(1<<15), 0x054020|(0b110100<<8)|(1<<15), func() Instruction { return new(Movec_ea_Imm) })
}

func (ins *Movec_ea_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.Memory = xyspace(bit(opcode, 6))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))
	ins.Immediate = extensionWord

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movec_ea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.ProgramController,
	)
}

// Function Idx: func83, func198, func199, func84
// 00000101W0aaaaaa0S1DDDDD
// MOVE(C) [X or Y]:aa,D1 / MOVE(C) S1,[X or Y]:aa
// Page 178
type Movec_aa struct {
	IsWrite           bool
	Address           uint32
	Memory            Memory
	ProgramController Register
}

// WordCount implements Instruction
func (*Movec_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff40a0, 0x050020, func() Instruction { return new(Movec_aa) })
}

func (ins *Movec_aa) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movec_aa) Disassemble(w TokenWriter) error {
	// Example: move m0,x:<$0
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			ins.Memory, ForceShortAddressMode, absAddr(ins.Memory, ins.Address),
			OperandSeparator,
			ins.ProgramController,
		)
	}

	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.ProgramController,
		OperandSeparator,
		ins.Memory, ForceShortAddressMode, absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func85, func172, func193, func86
// 00000100W1eeeeee1o1DDDDD
// MOVE(C) S1,D2 / MOVE(C) S2,D1
// Page 178
type Movec_S1D2 struct {
	IsWrite             bool
	SourceOrDestination Register
	ProgramController   Register
}

// WordCount implements Instruction
func (*Movec_S1D2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff40a0, 0x0440a0, func() Instruction { return new(Movec_S1D2) })
}

func (ins *Movec_S1D2) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.SourceOrDestination = register6Bit(mask(opcode, 8, 6))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	if ins.ProgramController == RegisterSSH && ins.SourceOrDestination == RegisterSSH {
		return false
	}

	return true
}

func (ins *Movec_S1D2) Disassemble(w TokenWriter) error {
	// Example: move lc,lc
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			ins.SourceOrDestination,
			OperandSeparator,
			ins.ProgramController,
		)
	}

	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.ProgramController,
		OperandSeparator,
		ins.SourceOrDestination,
	)
}

// Function Idx: func82, func81, func196, func197
// 00000101iiiiiiii101DDDDD
// MOVE(C) #xx,D1
// Page 178
type Movec_xx struct {
	Immediate         uint32
	ProgramController Register
}

// WordCount implements Instruction
func (*Movec_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff00b8, 0x0500a0, func() Instruction { return new(Movec_xx) })
	registerInstruction(0xff00b8, 0x0500b8, func() Instruction { return new(Movec_xx) })
	registerInstruction(0xff00b8, 0x0500b0, func() Instruction { return new(Movec_xx) })
	registerInstruction(0xff00b8, 0x0500a8, func() Instruction { return new(Movec_xx) })
}

func (ins *Movec_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 8))
	ins.ProgramController = programControllerRegister5bit(mask(opcode, 0, 5))

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movec_xx) Disassemble(w TokenWriter) error {
	// Example: move #$0,m0
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.ProgramController,
	)
}

// Function Idx: func69
// 00000111W1MMMRRR10dddddd
// MOVE(M) S,P:ea / MOVE(M) P:ea,D
// Page 181
type Movem_ea struct {
	IsWrite             bool
	EffectiveAddress    EffectiveAddress
	SourceOrDestination Register
}

// WordCount implements Instruction
func (*Movem_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff4080, 0x074080, func() Instruction { return new(Movem_ea) })
}

func (ins *Movem_ea) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 6))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movem_ea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryP, ins.EffectiveAddress,
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Example: move p:-(r7),lc
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryP, ins.EffectiveAddress,
	)
}

// Function Idx: func69
// 00000111W1MMMRRR10dddddd
// MOVE(M) S,P:ea / MOVE(M) P:ea,D
// Page 181
type Movem_ea_Abs struct {
	IsWrite             bool
	Address             uint32
	SourceOrDestination Register
}

// WordCount implements Instruction
func (*Movem_ea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff4080|(0b111111<<8), 0x074080|(0b110000<<8), func() Instruction { return new(Movem_ea_Abs) })
}

func (ins *Movem_ea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 6))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movem_ea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func70
// 00000111W0aaaaaa00dddddd
// MOVE(M) S,P:aa / MOVE(M) P:aa,D
// Page 181
type Movem_aa struct {
	IsWrite             bool
	Address             uint32
	SourceOrDestination Register
}

// WordCount implements Instruction
func (*Movem_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff4080, 0x070000, func() Instruction { return new(Movem_aa) })
}

func (ins *Movem_aa) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Address = (mask(opcode, 8, 6))
	ins.SourceOrDestination = register6Bit(mask(opcode, 0, 6))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movem_aa) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMove,
			ColumnSeparator,
			MemoryP, ForceShortAddressMode, absAddr(MemoryP, ins.Address),
			OperandSeparator,
			ins.SourceOrDestination,
		)
	}

	// Example: move p:<$3f,lc
	return w.Write(
		MnemonicMove,
		ColumnSeparator,
		ins.SourceOrDestination,
		OperandSeparator,
		MemoryP, ForceShortAddressMode, absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func65
// 0000100sW1MMMRRR1Spppppp
// MOVEP [X or Y]:pp,[X or Y]:ea / MOVEP [X or Y]:ea,[X or Y]:pp
// Page 183
type Movep_ppea struct {
	PeripheralMemory  Memory
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_ppea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfe4080, 0x084080, func() Instruction { return new(Movep_ppea) })
}

func (ins *Movep_ppea) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_ppea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ins.EffectiveAddress,
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(ins.PeripheralMemory, ins.PeripheralAddress),
		)
	}

	// Example: movep x:<<$ffffc0,x:(r0)-n0
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(ins.PeripheralMemory, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ins.EffectiveAddress,
	)
}

// Function Idx: func65
// 0000100sW1MMMRRR1Spppppp
// MOVEP [X or Y]:pp,[X or Y]:ea / MOVEP [X or Y]:ea,[X or Y]:pp
// Page 183
type Movep_ppea_Abs struct {
	PeripheralMemory  Memory
	IsWrite           bool
	Address           uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_ppea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfe4080|(0b111111<<8), 0x084080|(0b110000<<8), func() Instruction { return new(Movep_ppea_Abs) })
}

func (ins *Movep_ppea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_ppea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(ins.PeripheralMemory, ins.PeripheralAddress),
		)
	}

	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(ins.PeripheralMemory, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func65
// 0000100sW1MMMRRR1Spppppp
// MOVEP [X or Y]:pp,[X or Y]:ea / MOVEP [X or Y]:ea,[X or Y]:pp
// Page 183
type Movep_ppea_Imm struct {
	PeripheralMemory  Memory
	Immediate         uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_ppea_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfe4080|(0b111111<<8)|(1<<15), 0x084080|(0b110100<<8)|(1<<15), func() Instruction { return new(Movep_ppea_Imm) })
}

func (ins *Movep_ppea_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.Immediate = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_ppea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		absAddr(ins.Memory, ins.Immediate),
		OperandSeparator,
		ins.PeripheralMemory,
		ForceIOShortAddressMode,
		absAddr(ins.PeripheralMemory, ins.PeripheralAddress),
	)
}

// Function Idx: func176
// 00000111W1MMMRRR0Sqqqqqq
// MOVEP X:qq,[X or Y]:ea / MOVEP [X or Y]:ea,X:qq
// Page 183
type Movep_Xqqea struct {
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Xqqea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff4080, 0x074000, func() Instruction { return new(Movep_Xqqea) })
}

func (ins *Movep_Xqqea) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Xqqea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ins.EffectiveAddress,
			OperandSeparator,
			MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		)
	}

	// Example: movep x:<<$ffff80,x:(r0)-n0
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ins.EffectiveAddress,
	)
}

// Function Idx: func176
// 00000111W1MMMRRR0Sqqqqqq
// MOVEP X:qq,[X or Y]:ea / MOVEP [X or Y]:ea,X:qq
// Page 183
type Movep_Xqqea_Abs struct {
	IsWrite           bool
	Address           uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Xqqea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff4080|(0b111111<<8), 0x074000|(0b110000<<8), func() Instruction { return new(Movep_Xqqea_Abs) })
}

func (ins *Movep_Xqqea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Xqqea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
			OperandSeparator,
			MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		)
	}

	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func176
// 00000111W1MMMRRR0Sqqqqqq
// MOVEP X:qq,[X or Y]:ea / MOVEP [X or Y]:ea,X:qq
// Page 183
type Movep_Xqqea_Imm struct {
	Immediate         uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Xqqea_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff4080|(0b111111<<8)|(1<<15), 0x074000|(0b110100<<8)|(1<<15), func() Instruction { return new(Movep_Xqqea_Imm) })
}

func (ins *Movep_Xqqea_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Xqqea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		MemoryX,
		ForceIOShortAddressMode,
		absAddr(MemoryX, ins.PeripheralAddress),
	)
}

// Function Idx: func177
// 00000111W0MMMRRR1Sqqqqqq
// MOVEP Y:qq,[X or Y]:ea / MOVEP [X or Y]:ea,Y:qq
// Page 183
type Movep_Yqqea struct {
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Yqqea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff4080, 0x070080, func() Instruction { return new(Movep_Yqqea) })
}

func (ins *Movep_Yqqea) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Yqqea) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ins.EffectiveAddress,
			OperandSeparator,
			MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	// Example: movep y:<<$ffff80,x:(r0)-n0
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ins.EffectiveAddress,
	)
}

// Function Idx: func177
// 00000111W0MMMRRR1Sqqqqqq
// MOVEP Y:qq,[X or Y]:ea / MOVEP [X or Y]:ea,Y:qq
// Page 183
type Movep_Yqqea_Abs struct {
	IsWrite           bool
	Address           uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Yqqea_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff4080|(0b111111<<8), 0x070080|(0b110000<<8), func() Instruction { return new(Movep_Yqqea_Abs) })
}

func (ins *Movep_Yqqea_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.Address = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Yqqea_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
			OperandSeparator,
			MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		ins.Memory, ForceLongAddressMode, absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func177
// 00000111W0MMMRRR1Sqqqqqq
// MOVEP Y:qq,[X or Y]:ea / MOVEP [X or Y]:ea,Y:qq
// Page 183
type Movep_Yqqea_Imm struct {
	Immediate         uint32
	Memory            Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_Yqqea_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff4080|(0b111111<<8)|(1<<15), 0x070080|(0b110100<<8)|(1<<15), func() Instruction { return new(Movep_Yqqea_Imm) })
}

func (ins *Movep_Yqqea_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = extensionWord
	ins.Memory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_Yqqea_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		MemoryY,
		ForceIOShortAddressMode,
		absAddr(MemoryY, ins.PeripheralAddress),
	)
}

// Function Idx: func66
// 0000100sW1MMMRRR01pppppp
// MOVEP P:ea,[X or Y]:pp / MOVEP [X or Y]:pp,P:ea
// Page 183
type Movep_eapp struct {
	PeripheralMemory  Memory
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eapp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfe40c0, 0x084040, func() Instruction { return new(Movep_eapp) })
}

func (ins *Movep_eapp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eapp) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			MemoryP, ins.EffectiveAddress,
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	// Example: movep x:<<$ffffc0,p:(r0)-n0
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		MemoryP, ins.EffectiveAddress,
	)
}

// Function Idx: func66
// 0000100sW1MMMRRR01pppppp
// MOVEP P:ea,[X or Y]:pp / MOVEP [X or Y]:pp,P:ea
// Page 183
type Movep_eapp_Abs struct {
	PeripheralMemory  Memory
	IsWrite           bool
	Address           uint32
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eapp_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfe40c0|(0b111111<<8), 0x084040|(0b110000<<8), func() Instruction { return new(Movep_eapp_Abs) })
}

func (ins *Movep_eapp_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eapp_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func66
// 0000100sW1MMMRRR01pppppp
// MOVEP P:ea,[X or Y]:pp / MOVEP [X or Y]:pp,P:ea
// Page 183
type Movep_eapp_Imm struct {
	PeripheralMemory  Memory
	Immediate         uint32
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eapp_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfe40c0|(0b111111<<8)|(1<<15), 0x084040|(0b110100<<8)|(1<<15), func() Instruction { return new(Movep_eapp_Imm) })
}

func (ins *Movep_eapp_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = extensionWord
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eapp_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.PeripheralMemory,
		ForceIOShortAddressMode,
		absAddr(MemoryY, ins.PeripheralAddress),
	)
}

// Function Idx: func180
// 000000001WMMMRRR0Sqqqqqq
// MOVEP P:ea,[X or Y]:qq / MOVEP [X or Y]:qq,P:ea
// Page 183
type Movep_eaqq struct {
	IsWrite           bool
	EffectiveAddress  EffectiveAddress
	PeripheralMemory  Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eaqq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff8080, 0x008000, func() Instruction { return new(Movep_eaqq) })
}

func (ins *Movep_eaqq) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 14))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.PeripheralMemory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eaqq) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			MemoryP, ins.EffectiveAddress,
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	// Example: movep x:<<$ffff80,p:(r0)-n0
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		MemoryP, ins.EffectiveAddress,
	)
}

// Function Idx: func180
// 000000001WMMMRRR0Sqqqqqq
// MOVEP P:ea,[X or Y]:qq / MOVEP [X or Y]:qq,P:ea
// Page 183
type Movep_eaqq_Abs struct {
	IsWrite           bool
	Address           uint32
	PeripheralMemory  Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eaqq_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff8080|(0b111111<<8), 0x008000|(0b110000<<8), func() Instruction { return new(Movep_eaqq_Abs) })
}

func (ins *Movep_eaqq_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	ins.IsWrite = boolean(bit(opcode, 14))
	ins.PeripheralMemory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eaqq_Abs) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		MemoryP, ForceLongAddressMode, absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func180
// 000000001WMMMRRR0Sqqqqqq
// MOVEP P:ea,[X or Y]:qq / MOVEP [X or Y]:qq,P:ea
// Page 183
type Movep_eaqq_Imm struct {
	Immediate         uint32
	PeripheralMemory  Memory
	PeripheralAddress uint32
}

// WordCount implements Instruction
func (*Movep_eaqq_Imm) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xff8080|(0b111111<<8)|(1<<14), 0x008000|(0b110100<<8)|(1<<14), func() Instruction { return new(Movep_eaqq_Imm) })
}

func (ins *Movep_eaqq_Imm) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = extensionWord
	ins.PeripheralMemory = xyspace(bit(opcode, 6))
	ins.PeripheralAddress = lowPeripheral(mask(opcode, 0, 6))
	return true
}

func (ins *Movep_eaqq_Imm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.PeripheralMemory,
		ForceIOShortAddressMode,
		absAddr(MemoryY, ins.PeripheralAddress),
	)
}

// Function Idx: func67
// 0000100sW1dddddd00pppppp
// MOVEP S,[X or Y]:pp / MOVEP [X or Y]:pp,D
// Page 183
type Movep_Spp struct {
	PeripheralMemory    Memory
	IsWrite             bool
	SourceOrDestination Register
	PeripheralAddress   uint32
}

// WordCount implements Instruction
func (*Movep_Spp) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfe40c0, 0x084000, func() Instruction { return new(Movep_Spp) })
}

func (ins *Movep_Spp) Decode(opcode, extensionWord uint32) bool {
	ins.PeripheralMemory = xyspace(bit(opcode, 16))
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.SourceOrDestination = register6Bit(mask(opcode, 8, 6))
	ins.PeripheralAddress = highPeripheral(mask(opcode, 0, 6))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movep_Spp) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.SourceOrDestination,
			OperandSeparator,
			ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	// Example: movep lc,y:<<$ffffff
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		ins.PeripheralMemory, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		ins.SourceOrDestination,
	)
}

// Function Idx: func178
// 00000100W1dddddd1q0qqqqq
// MOVEP S,X:qq / MOVEP X:qq,D
// Page 183
type Movep_SXqq struct {
	IsWrite             bool
	SourceOrDestination Register
	PeripheralAddress   uint32
}

// WordCount implements Instruction
func (*Movep_SXqq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff40a0, 0x044080, func() Instruction { return new(Movep_SXqq) })
}

func (ins *Movep_SXqq) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.SourceOrDestination = register6Bit(mask(opcode, 8, 6))
	ins.PeripheralAddress = lowPeripheral(multimask(opcode, 6, 1, 0, 5))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movep_SXqq) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.SourceOrDestination,
			OperandSeparator,
			MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		)
	}

	// Example: movep lc,x:<<$ffffbf
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryX, ForceIOShortAddressMode, absAddr(MemoryX, ins.PeripheralAddress),
		OperandSeparator,
		ins.SourceOrDestination,
	)
}

// Function Idx: func179
// 00000100W1dddddd0q1qqqqq
// MOVEP S,Y:qq / MOVEP Y:qq,D
// Page 183
type Movep_SYqq struct {
	IsWrite             bool
	SourceOrDestination Register
	PeripheralAddress   uint32
}

// WordCount implements Instruction
func (*Movep_SYqq) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff40a0, 0x044020, func() Instruction { return new(Movep_SYqq) })
}

func (ins *Movep_SYqq) Decode(opcode, extensionWord uint32) bool {
	ins.IsWrite = boolean(bit(opcode, 15))
	ins.SourceOrDestination = register6Bit(mask(opcode, 8, 6))
	ins.PeripheralAddress = lowPeripheral(multimask(opcode, 6, 1, 0, 5))

	if ins.SourceOrDestination == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Movep_SYqq) Disassemble(w TokenWriter) error {
	if ins.IsWrite {
		return w.Write(
			MnemonicMovep,
			ColumnSeparator,
			ins.SourceOrDestination,
			OperandSeparator,
			MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		)
	}

	// Example: movep lc,y:<<$ffffbf
	return w.Write(
		MnemonicMovep,
		ColumnSeparator,
		MemoryY, ForceIOShortAddressMode, absAddr(MemoryY, ins.PeripheralAddress),
		OperandSeparator,
		ins.SourceOrDestination,
	)
}

// Function Idx: func109
// 00000001000sssss11QQdk00
// MPY (+/-)S,#n,D
// Page 186
type Mpy_SD struct {
	Immediate   uint32
	Source      Register
	Destination Register
	Sign        bool
}

// WordCount implements Instruction
func (*Mpy_SD) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffe0c3, 0x0100c0, func() Instruction { return new(Mpy_SD) })
}

func (ins *Mpy_SD) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 5))
	ins.Source = dataALUMultiplyOperands1(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	return true
}

func (ins *Mpy_SD) Disassemble(w TokenWriter) error {
	// Example: mpy y1,#$0,a
	return w.Write(
		MnemonicMpy,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source,
		OperandSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func175
// 00000001001001111sdkQQQQ
// MPY su (+/-)S1,S2,D / MPY uu (+/-)S1,S2,D
// Page 188
type Mpy_su struct {
	Mode        MultiplyMode
	Destination Register
	Sign        bool
	Sources     RegisterPair
}

// WordCount implements Instruction
func (*Mpy_su) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffff80, 0x012780, func() Instruction { return new(Mpy_su) })
}

func (ins *Mpy_su) Decode(opcode, extensionWord uint32) bool {
	ins.Mode = multiplyMode1Bit(bit(opcode, 6))
	ins.Destination = accumulator(bit(opcode, 5))
	ins.Sign = boolean(bit(opcode, 4))
	ins.Sources = multiplyAllPairs(mask(opcode, 0, 4))
	return true
}

func (ins *Mpy_su) Disassemble(w TokenWriter) error {
	// Example: mpysu x0,x0,a
	return w.Write(
		MnemonicMpy,
		ins.Mode,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Sources.First,
		OperandSeparator,
		ins.Sources.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func155
// 0000000101ooooo111qqdk00
// MPYI (+/-)#xxxx,S,D
// Page 189
type Mpyi struct {
	Source      Register
	Destination Register
	Sign        bool
	Immediate   uint32
}

// WordCount implements Instruction
func (*Mpyi) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c3, 0x0141c0, func() Instruction { return new(Mpyi) })
}

func (ins *Mpyi) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUMultiplyOperands2(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	ins.Immediate = extensionWord
	return true
}

func (ins *Mpyi) Disassemble(w TokenWriter) error {
	// Example: mpyi #>$deface,x0,a
	return w.Write(
		MnemonicMpyi,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ImmediateAddressMode,
		ForceLongAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func109
// 00000001000sssss11QQdk01
// MPYR (+/-)S,#n,D
// Page 190
type Mpyr_SD struct {
	Immediate   uint32
	Source      Register
	Destination Register
	Sign        bool
}

// WordCount implements Instruction
func (*Mpyr_SD) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffe0c3, 0x0100c1, func() Instruction { return new(Mpyr_SD) })
}

func (ins *Mpyr_SD) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 5))
	ins.Source = dataALUMultiplyOperands1(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	return true
}

func (ins *Mpyr_SD) Disassemble(w TokenWriter) error {
	// Example: mpyr y1,#$0,a
	return w.Write(
		MnemonicMpyr,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source,
		OperandSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func155
// 0000000101ooooo111qqdk01
// MPYRI (+/-)#xxxx,S,D
// Page 192
type Mpyri struct {
	Source      Register
	Destination Register
	Sign        bool
	Immediate   uint32
}

// WordCount implements Instruction
func (*Mpyri) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c3, 0x0141c1, func() Instruction { return new(Mpyri) })
}

func (ins *Mpyri) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUMultiplyOperands2(mask(opcode, 4, 2))
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Sign = boolean(bit(opcode, 2))
	ins.Immediate = extensionWord
	return true
}

func (ins *Mpyri) Disassemble(w TokenWriter) error {
	// Example: mpyri #>$deface,x0,a
	return w.Write(
		MnemonicMpyri,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ImmediateAddressMode,
		ForceLongAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func107
// 000000000000000000000000
// NOP
// Page 195
type Nop struct {
}

// WordCount implements Instruction
func (*Nop) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000000, func() Instruction { return new(Nop) })
}

func (ins *Nop) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Nop) Disassemble(w TokenWriter) error {
	// Example: nop
	return w.Write(
		MnemonicNop,
	)
}

// Function Idx: func91
// 0000000111011RRR0001d101
// NORM Rn,D
// Page 196
type Norm struct {
	Address     Register
	Destination Register
}

// WordCount implements Instruction
func (*Norm) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfff8f7, 0x01d815, func() Instruction { return new(Norm) })
}

func (ins *Norm) Decode(opcode, extensionWord uint32) bool {
	ins.Address = addressRegister(mask(opcode, 8, 3))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Norm) Disassemble(w TokenWriter) error {
	// Example: norm r0,a
	return w.Write(
		MnemonicNorm,
		ColumnSeparator,
		ins.Address,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func168
// 00001100000111100010sssD
// NORMF S,D
// Page 198
type Normf struct {
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Normf) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x0c1e20, func() Instruction { return new(Normf) })
}

func (ins *Normf) Decode(opcode, extensionWord uint32) bool {
	ins.Source = dataALUOperands1(mask(opcode, 1, 3))
	ins.Destination = accumulator(bit(opcode, 0))

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Normf) Disassemble(w TokenWriter) error {
	// Example: normf y1,b
	return w.Write(
		MnemonicNormf,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func170
// 0000000101iiiiii10ood010
// OR #xx,D
// Page 201
type Or_xx struct {
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Or_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014082, func() Instruction { return new(Or_xx) })
}

func (ins *Or_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Or_xx) Disassemble(w TokenWriter) error {
	// Example: or #<$0,a
	return w.Write(
		MnemonicOr,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood010
// OR #xxxx,D
// Page 201
type Or_xxxx struct {
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Or_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c2, func() Instruction { return new(Or_xxxx) })
}

func (ins *Or_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *Or_xxxx) Disassemble(w TokenWriter) error {
	// Example: or #>$deface,a
	return w.Write(
		MnemonicOr,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func95
// 00000000iiiiiiii111110EE
// OR(I) #xx,D
// Page 203
type Ori struct {
	Immediate         uint32
	ProgramController Register
}

// WordCount implements Instruction
func (*Ori) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff00fc, 0x0000f8, func() Instruction { return new(Ori) })
}

func (ins *Ori) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 8))
	ins.ProgramController = programControlUnitRegister(mask(opcode, 0, 2))

	if ins.ProgramController == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Ori) Disassemble(w TokenWriter) error {
	// Example: ori #$0,mr
	return w.Write(
		MnemonicOri,
		ColumnSeparator,
		ImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.ProgramController,
	)
}

// Function Idx: func147
// 000000000000000000000011
// PFLUSH
// Page 205
type Pflush struct {
}

// WordCount implements Instruction
func (*Pflush) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000003, func() Instruction { return new(Pflush) })
}

func (ins *Pflush) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Pflush) Disassemble(w TokenWriter) error {
	// Example: pflush
	return w.Write(
		MnemonicPflush,
	)
}

// Function Idx: func204
// 000000000000000000000001
// PFLUSHUN
// Page 206
type Pflushun struct {
}

// WordCount implements Instruction
func (*Pflushun) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000001, func() Instruction { return new(Pflushun) })
}

func (ins *Pflushun) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Pflushun) Disassemble(w TokenWriter) error {
	// Example: pflushun
	return w.Write(
		MnemonicPflushun,
	)
}

// Function Idx: func146
// 000000000000000000000010
// PFREE
// Page 207
type Pfree struct {
}

// WordCount implements Instruction
func (*Pfree) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000002, func() Instruction { return new(Pfree) })
}

func (ins *Pfree) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Pfree) Disassemble(w TokenWriter) error {
	// Example: pfree
	return w.Write(
		MnemonicPfree,
	)
}

// Function Idx: func154
// 0000101111MMMRRR10000001
// PLOCK ea
// Undocumented
type Plock struct {
	EffectiveAddress EffectiveAddress
}

// WordCount implements Instruction
func (*Plock) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0ff, 0x0bc081, func() Instruction { return new(Plock) })
}

func (ins *Plock) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	return true
}

func (ins *Plock) Disassemble(w TokenWriter) error {
	// Example: plock (r0)-n0
	return w.Write(
		MnemonicPlock,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func154
// 0000101111MMMRRR10000001
// PLOCK ea
// Undocumented
type Plock_Abs struct {
	Address uint32
}

// WordCount implements Instruction
func (*Plock_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0ff|(0b111111<<8), 0x0bc081|(0b110000<<8), func() Instruction { return new(Plock_Abs) })
}

func (ins *Plock_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	return true
}

func (ins *Plock_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicPlock,
		ColumnSeparator,
		absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func148
// 000000000000000000001111
// PLOCKR xxxx
// Page 208
type Plockr struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Plockr) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffffff, 0x00000f, func() Instruction { return new(Plockr) })
}

func (ins *Plockr) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Plockr) Disassemble(w TokenWriter) error {
	// Example: plockr >*-$210532
	return w.Write(
		MnemonicPlockr,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func153
// 0000101011MMMRRR10000001
// PUNLOCK ea
// Page 209
type Punlock struct {
	EffectiveAddress EffectiveAddress
}

// WordCount implements Instruction
func (*Punlock) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0ff, 0x0ac081, func() Instruction { return new(Punlock) })
}

func (ins *Punlock) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	return true
}

func (ins *Punlock) Disassemble(w TokenWriter) error {
	// Example: punlock (r0)-n0
	return w.Write(
		MnemonicPunlock,
		ColumnSeparator,
		ins.EffectiveAddress,
	)
}

// Function Idx: func153
// 0000101011MMMRRR10000001
// PUNLOCK ea
// Page 209
type Punlock_Abs struct {
	Address uint32
}

// WordCount implements Instruction
func (*Punlock_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc0ff|(0b111111<<8), 0x0ac081|(0b110000<<8), func() Instruction { return new(Punlock_Abs) })
}

func (ins *Punlock_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Address = extensionWord
	return true
}

func (ins *Punlock_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicPunlock,
		ColumnSeparator,
		absAddr(MemoryP, ins.Address),
	)
}

// Function Idx: func149
// 000000000000000000001110
// PUNLOCKR xxxx
// Page 210
type Punlockr struct {
	Displacement int32
}

// WordCount implements Instruction
func (*Punlockr) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffffff, 0x00000e, func() Instruction { return new(Punlockr) })
}

func (ins *Punlockr) Decode(opcode, extensionWord uint32) bool {
	ins.Displacement = signExtend(24, extensionWord)
	return true
}

func (ins *Punlockr) Disassemble(w TokenWriter) error {
	// Example: punlockr >*-$210532
	return w.Write(
		MnemonicPunlockr,
		ColumnSeparator,
		ForceLongAddressMode,
		relativeAddr(ins.Displacement),
	)
}

// Function Idx: func73
// 0000011001MMMRRR0S100000
// REP [X or Y]:ea
// Page 211
type Rep_ea struct {
	EffectiveAddress EffectiveAddress
	Memory           Memory
}

// WordCount implements Instruction
func (*Rep_ea) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x064020, func() Instruction { return new(Rep_ea) })
}

func (ins *Rep_ea) Decode(opcode, extensionWord uint32) bool {
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Memory = xyspace(bit(opcode, 6))
	return true
}

func (ins *Rep_ea) Disassemble(w TokenWriter) error {
	// Example: rep x:(r0)-n0
	return w.Write(
		MnemonicRep,
		ColumnSeparator,
		ins.Memory,
		ins.EffectiveAddress,
	)
}

// Function Idx: func77
// 0000011000aaaaaa0S100000
// REP [X or Y]:aa
// Page 211
type Rep_aa struct {
	Address uint32
	Memory  Memory
}

// WordCount implements Instruction
func (*Rep_aa) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x060020, func() Instruction { return new(Rep_aa) })
}

func (ins *Rep_aa) Decode(opcode, extensionWord uint32) bool {
	ins.Address = (mask(opcode, 8, 6))
	ins.Memory = xyspace(bit(opcode, 6))
	return true
}

func (ins *Rep_aa) Disassemble(w TokenWriter) error {
	// Example: rep x:<$0
	return w.Write(
		MnemonicRep,
		ColumnSeparator,
		ins.Memory,
		ForceShortAddressMode,
		absAddr(ins.Memory, ins.Address),
	)
}

// Function Idx: func75
// 00000110iiiiiiii1o1ohhhh
// REP #xxx
// Page 211
type Rep_xxx struct {
	Immediate uint32
}

// WordCount implements Instruction
func (*Rep_xxx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff00a0, 0x0600a0, func() Instruction { return new(Rep_xxx) })
}

func (ins *Rep_xxx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (multimask(opcode, 0, 4, 8, 8))
	return true
}

func (ins *Rep_xxx) Disassemble(w TokenWriter) error {
	// Example: rep #<$0
	return w.Write(
		MnemonicRep,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
	)
}

// Function Idx: func71
// 0000011011dddddd00100000
// REP S
// Page 211
type Rep_S struct {
	Source Register
}

// WordCount implements Instruction
func (*Rep_S) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0a0, 0x06c020, func() Instruction { return new(Rep_S) })
}

func (ins *Rep_S) Decode(opcode, extensionWord uint32) bool {
	ins.Source = register6Bit(mask(opcode, 8, 6))

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Rep_S) Disassemble(w TokenWriter) error {
	// Example: rep lc
	return w.Write(
		MnemonicRep,
		ColumnSeparator,
		ins.Source,
	)
}

// Function Idx: func100
// 00000000000000001o0o0100
// RESET
// Page 213
type Reset struct {
}

// WordCount implements Instruction
func (*Reset) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffaf, 0x000084, func() Instruction { return new(Reset) })
}

func (ins *Reset) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Reset) Disassemble(w TokenWriter) error {
	// Example: reset
	return w.Write(
		MnemonicReset,
	)
}

// Function Idx: func106
// 000000000000000000000100
// RTI
// Page 220
type Rti struct {
}

// WordCount implements Instruction
func (*Rti) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000004, func() Instruction { return new(Rti) })
}

func (ins *Rti) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Rti) Disassemble(w TokenWriter) error {
	// Example: rti
	return w.Write(
		MnemonicRti,
	)
}

// Function Idx: func101
// 000000000000000000001100
// RTS
// Page 221
type Rts struct {
}

// WordCount implements Instruction
func (*Rts) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x00000c, func() Instruction { return new(Rts) })
}

func (ins *Rts) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Rts) Disassemble(w TokenWriter) error {
	// Example: rts
	return w.Write(
		MnemonicRts,
	)
}

// Function Idx: func98
// 00000000000000001o0o0111
// STOP
// Page 223
type Stop struct {
}

// WordCount implements Instruction
func (*Stop) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffaf, 0x000087, func() Instruction { return new(Stop) })
}

func (ins *Stop) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Stop) Disassemble(w TokenWriter) error {
	// Example: stop
	return w.Write(
		MnemonicStop,
	)
}

// Function Idx: func170
// 0000000101iiiiii10ood100
// SUB #xx,D
// Page 225
type Sub_xx struct {
	Immediate   uint32
	Destination Register
}

// WordCount implements Instruction
func (*Sub_xx) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffc0c7, 0x014084, func() Instruction { return new(Sub_xx) })
}

func (ins *Sub_xx) Decode(opcode, extensionWord uint32) bool {
	ins.Immediate = (mask(opcode, 8, 6))
	ins.Destination = accumulator(bit(opcode, 3))
	return true
}

func (ins *Sub_xx) Disassemble(w TokenWriter) error {
	// Example: sub #<$0,a
	return w.Write(
		MnemonicSub,
		ColumnSeparator,
		ForceShortImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func171
// 0000000101ooooo011ood100
// SUB #xxxx,D
// Page 225
type Sub_xxxx struct {
	Destination Register
	Immediate   uint32
}

// WordCount implements Instruction
func (*Sub_xxxx) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xffc1c7, 0x0140c4, func() Instruction { return new(Sub_xxxx) })
}

func (ins *Sub_xxxx) Decode(opcode, extensionWord uint32) bool {
	ins.Destination = accumulator(bit(opcode, 3))
	ins.Immediate = extensionWord
	return true
}

func (ins *Sub_xxxx) Disassemble(w TokenWriter) error {
	// Example: sub #>$deface,a
	return w.Write(
		MnemonicSub,
		ColumnSeparator,
		ForceLongImmediateAddressMode,
		immediate(ins.Immediate),
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func90
// 00000010CCCC0ooo0JJJdooo
// Tcc S1,D1
// Page 230
type Tcc_S1D1 struct {
	Condition   Condition
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Tcc_S1D1) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0880, 0x020000, func() Instruction { return new(Tcc_S1D1) })
}

func (ins *Tcc_S1D1) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.Source = dataALUOperands3(mask(opcode, 4, 3), bit(opcode, 3))
	ins.Destination = accumulator(bit(opcode, 3))

	if ins.Source == RegisterInvalid {
		return false
	}

	return true
}

func (ins *Tcc_S1D1) Disassemble(w TokenWriter) error {
	// Example: tcc b,a
	return w.Write(
		MnemonicTcc,
		ins.Condition,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func89
// 00000011CCCCottt0JJJdTTT
// Tcc S1,D1 S2,D2
// Page 230
type Tcc_S1D1S2D2 struct {
	Condition          Condition
	SourceAddress      Register
	Source             Register
	Destination        Register
	DestinationAddress Register
}

// WordCount implements Instruction
func (*Tcc_S1D1S2D2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0080, 0x030000, func() Instruction { return new(Tcc_S1D1S2D2) })
}

func (ins *Tcc_S1D1S2D2) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.SourceAddress = addressRegister(mask(opcode, 8, 3))
	ins.Source = dataALUOperands3(mask(opcode, 4, 3), bit(opcode, 3))

	if ins.Source == RegisterInvalid {
		return false
	}

	ins.Destination = accumulator(bit(opcode, 3))
	ins.DestinationAddress = addressRegister(mask(opcode, 0, 3))
	return true
}

func (ins *Tcc_S1D1S2D2) Disassemble(w TokenWriter) error {
	// Example: tcc b,a r0,r0
	return w.Write(
		MnemonicTcc,
		ins.Condition,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
		ColumnSeparator,
		ins.SourceAddress,
		OperandSeparator,
		ins.DestinationAddress,
	)
}

// Function Idx: func156
// 00000010CCCC1ttt0ooooTTT
// Tcc S2,D2
// Page 230
type Tcc_S2D2 struct {
	Condition   Condition
	Source      Register
	Destination Register
}

// WordCount implements Instruction
func (*Tcc_S2D2) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xff0880, 0x020800, func() Instruction { return new(Tcc_S2D2) })
}

func (ins *Tcc_S2D2) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 12, 4))
	ins.Source = addressRegister(mask(opcode, 8, 3))
	ins.Destination = addressRegister(mask(opcode, 0, 3))
	return true
}

func (ins *Tcc_S2D2) Disassemble(w TokenWriter) error {
	// Example: tcc r0,r0
	return w.Write(
		MnemonicTcc,
		ins.Condition,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// Function Idx: func104
// 000000000000000000000110
// TRAP
// Page 233
type Trap struct {
}

// WordCount implements Instruction
func (*Trap) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffff, 0x000006, func() Instruction { return new(Trap) })
}

func (ins *Trap) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Trap) Disassemble(w TokenWriter) error {
	// Example: swi
	return w.Write(
		MnemonicSwi,
	)
}

// Function Idx: func150
// 00000000000000000001CCCC
// TRAPcc
// Page 234
type Trapcc struct {
	Condition Condition
}

// WordCount implements Instruction
func (*Trapcc) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfffff0, 0x000010, func() Instruction { return new(Trapcc) })
}

func (ins *Trapcc) Decode(opcode, extensionWord uint32) bool {
	ins.Condition = condition(mask(opcode, 0, 4))
	return true
}

func (ins *Trapcc) Disassemble(w TokenWriter) error {
	// Example: trapcc
	return w.Write(
		MnemonicTrap,
		ins.Condition,
	)
}

// Viterbi Shift Left
// Function Idx: func205
// 0000101S11MMMRRR110i0000
// VSL S,i,L:ea
// Undocumented
type Vsl struct {
	Source           Register
	EffectiveAddress EffectiveAddress
	Bit              uint32 // Bit to append after shifting
}

// WordCount implements Instruction
func (*Vsl) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xfec0ef, 0x0ac0c0, func() Instruction { return new(Vsl) })
	registerInstruction(0xffc0ef, 0x0bc0c0, func() Instruction { return new(Vsl) })
}

func (ins *Vsl) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 16))
	ins.EffectiveAddress = addressMode(mask(opcode, 11, 3), mask(opcode, 8, 3))
	if ins.EffectiveAddress.Mode == AddressModeInvalid {
		return false
	}

	ins.Bit = (bit(opcode, 4))
	return true
}

func (ins *Vsl) Disassemble(w TokenWriter) error {
	// Example: vsl a,0,l:(r0)-n0
	return w.Write(
		MnemonicVsl,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ifThen(ins.Bit == 0, ZeroToken, OneToken),
		OperandSeparator,
		MemoryL,
		ins.EffectiveAddress,
	)
}

// Viterbi Shift Left
// Function Idx: func205
// 0000101S11MMMRRR110i0000
// VSL S,i,L:ea
// Undocumented
type Vsl_Abs struct {
	Source  Register
	Address uint32
	Bit     uint32 // Bit to append after shifting
}

// WordCount implements Instruction
func (*Vsl_Abs) UsesExtensionWord() bool {
	return true
}

func init() {
	registerInstruction(0xfec0ef|(0b111111<<8), 0x0ac0c0|(0b110000<<8), func() Instruction { return new(Vsl_Abs) })
	registerInstruction(0xffc0ef|(0b111111<<8), 0x0bc0c0|(0b110000<<8), func() Instruction { return new(Vsl_Abs) })
}

func (ins *Vsl_Abs) Decode(opcode, extensionWord uint32) bool {
	ins.Source = accumulator(bit(opcode, 16))
	ins.Address = extensionWord
	ins.Bit = (bit(opcode, 4))
	return true
}

func (ins *Vsl_Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicVsl,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ifThen(ins.Bit == 0, ZeroToken, OneToken),
		OperandSeparator,
		MemoryL,
		absAddr(MemoryL, ins.Address),
	)
}

// Function Idx: func99
// 00000000000000001o0o0110
// WAIT
// Page 236
type Wait struct{}

// WordCount implements Instruction
func (*Wait) UsesExtensionWord() bool {
	return false
}

func init() {
	registerInstruction(0xffffaf, 0x000086, func() Instruction { return new(Wait) })
}

func (ins *Wait) Decode(opcode, extensionWord uint32) bool {
	return true
}

func (ins *Wait) Disassemble(w TokenWriter) error {
	// Example: wait
	return w.Write(
		MnemonicWait,
	)
}
