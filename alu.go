package dsp56k

func init() {
	registerALUOperation(0b10000011, 0b10000000, func() ALUOperation { return new(Mpy_S1S2D) })
	registerALUOperation(0b10000011, 0b10000001, func() ALUOperation { return new(Mpyr_S1S2D) })
	registerALUOperation(0b10000011, 0b10000010, func() ALUOperation { return new(Mac_S1S2) })
	registerALUOperation(0b10000011, 0b10000011, func() ALUOperation { return new(Macr_S1S2) })
	registerALUOperation(0b10000111, 0b00000000, func() ALUOperation { return new(Add_SD) })
	registerALUOperation(0b10000111, 0b00000001, func() ALUOperation { return new(Tfr) })
	registerALUOperation(0b10000111, 0b00000100, func() ALUOperation { return new(Sub_SD) })
	registerALUOperation(0b10000111, 0b00000101, func() ALUOperation { return new(Cmp_S1S2) })
	registerALUOperation(0b10000111, 0b00000111, func() ALUOperation { return new(Cmpm_S1S2) })
	registerALUOperation(0b11000111, 0b01000010, func() ALUOperation { return new(Or_SD) })
	registerALUOperation(0b11000111, 0b01000011, func() ALUOperation { return new(Eor_SD) })
	registerALUOperation(0b11000111, 0b01000110, func() ALUOperation { return new(And_SD) })
	registerALUOperation(0b11100111, 0b00100001, func() ALUOperation { return new(Adc) })
	registerALUOperation(0b11100111, 0b00100101, func() ALUOperation { return new(Sbc) })
	registerALUOperation(0b11110111, 0b00000010, func() ALUOperation { return new(Addr) })
	registerALUOperation(0b11110111, 0b00000011, func() ALUOperation { return new(Tst) })
	registerALUOperation(0b11110111, 0b00000110, func() ALUOperation { return new(Subr) })
	registerALUOperation(0b11110111, 0b00010001, func() ALUOperation { return new(Rnd) })
	registerALUOperation(0b11110111, 0b00010010, func() ALUOperation { return new(Addl) })
	registerALUOperation(0b11110111, 0b00010011, func() ALUOperation { return new(Clr) })
	registerALUOperation(0b11110111, 0b00010110, func() ALUOperation { return new(Subl) })
	registerALUOperation(0b11110111, 0b00010111, func() ALUOperation { return new(Not) })
	registerALUOperation(0b11110111, 0b00100010, func() ALUOperation { return new(Asr_D) })
	registerALUOperation(0b11110111, 0b00100011, func() ALUOperation { return new(Lsr_D) })
	registerALUOperation(0b11110111, 0b00100110, func() ALUOperation { return new(Abs) })
	registerALUOperation(0b11110111, 0b00100111, func() ALUOperation { return new(Ror) })
	registerALUOperation(0b11110111, 0b00110010, func() ALUOperation { return new(Asl_D) })
	registerALUOperation(0b11110111, 0b00110011, func() ALUOperation { return new(Lsl_D) })
	registerALUOperation(0b11110111, 0b00110110, func() ALUOperation { return new(Neg) })
	registerALUOperation(0b11110111, 0b00110111, func() ALUOperation { return new(Rol) })
	registerALUOperation(0b11111111, 0b00010101, func() ALUOperation { return new(Maxm) })
	registerALUOperation(0b11111111, 0b00011101, func() ALUOperation { return new(Max) })
	registerALUOperation(0b11111111, 0b00000000, func() ALUOperation { return new(AluNop) })
}

// Page 156
// Move S,D
// 00000000
type AluNop struct{}

func (*AluNop) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	return 0, true
}

func (*AluNop) Disassemble(w TokenWriter) error {
	return w.Write(MnemonicMove)
}

func bit8(byte uint8, shift int) uint32 {
	return uint32(byte>>shift) & 1
}

func mask8(word uint8, shift, width int) uint32 {
	mask := uint8(1<<width - 1)
	return uint32((word >> shift) & mask)
}

// ABS D
// Page 25
// 0010d110
type Abs struct {
	Destination Register
}

func (ins *Abs) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAbs,
		ColumnSeparator,
		ins.Destination,
	)
}

func (ins *Abs) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

// ADC S,D
// Page 26
// 001Jd001
type Adc struct {
	Source      Register
	Destination Register
}

func (ins *Adc) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	ins.Source = inputRegister(bit8(opcode, 4))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Adc) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAdc,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// ADD S,D
// Page 27
// 0JJJd000
type Add_SD struct {
	Source      Register
	Destination Register
}

func (ins *Add_SD) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = jjj(mask8(opcode, 4, 3), bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)

	if ins.Source == RegisterInvalid {
		return 0, false
	}

	return writes, true
}

func (ins *Add_SD) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAdd,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// ADDL S,D
// Page 29
// 001d010
type Addl struct {
	Source      Register
	Destination Register
}

func (ins *Addl) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = oppositeAccumulator(bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Addl) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAddl,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// ADDR S,D
// Page 30
// 0000d010
type Addr struct {
	Source      Register
	Destination Register
}

func (ins *Addr) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = oppositeAccumulator(bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Addr) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAddr,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// AND S,D
// Page 31
// 01JJd110
type And_SD struct {
	Source      Register
	Destination Register
}

func (ins *And_SD) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUSourceOperands(mask8(opcode, 4, 2))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *And_SD) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAnd,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// ASL D
// Page 35
// 0011d010
type Asl_D struct {
	Destination Register
}

func (ins *Asl_D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Asl_D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAsl,
		ColumnSeparator,
		ins.Destination,
	)
}

// ASR D
// Page 38
// 0010d010
type Asr_D struct {
	Destination Register
}

func (ins *Asr_D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Asr_D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicAsr,
		ColumnSeparator,
		ins.Destination,
	)
}

// CLR D
// Page 75
// 0001d011
type Clr struct {
	Destination Register
}

func (ins *Clr) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Clr) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicClr,
		ColumnSeparator,
		ins.Destination,
	)
}

// CMP S1, S2
// Page 76
// 0JJJd101
type Cmp_S1S2 struct {
	Source      Register
	Destination Register
}

func (ins *Cmp_S1S2) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUOperands3(mask8(opcode, 4, 3), bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	return writes, true
}

func (ins *Cmp_S1S2) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicCmp,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// CMPM S1, S2
// Page 78
// 0JJJd111
type Cmpm_S1S2 struct {
	Source      Register
	Destination Register
}

func (ins *Cmpm_S1S2) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUOperands3(mask8(opcode, 4, 3), bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	return writes, true
}

func (ins *Cmpm_S1S2) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicCmpm,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// EOR S,D
// Page 101
// 01JJd011
type Eor_SD struct {
	Source      Register
	Destination Register
}

func (ins *Eor_SD) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUSourceOperands(mask8(opcode, 4, 2))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Eor_SD) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicEor,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// LSL D
// Page 136
// 0011D011
type Lsl_D struct {
	Destination Register
}

func (ins *Lsl_D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Lsl_D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicLsl,
		ColumnSeparator,
		ins.Destination,
	)
}

// LSR D
// Page 139
// 0010D011
type Lsr_D struct {
	Destination Register
}

func (ins *Lsr_D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Lsr_D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicLsr,
		ColumnSeparator,
		ins.Destination,
	)
}

// MAC (+/-)S1,S2,D / MAC (+/-)S2,S1,D
// Page 144
// 1QQQdk10
type Mac_S1S2 struct {
	Source      RegisterPair
	Destination Register
	Sign        bool
}

func (ins *Mac_S1S2) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	ins.Sign = boolean(bit8(opcode, 2))
	ins.Source = multiplyPairs(mask8(opcode, 4, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Mac_S1S2) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMac,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source.First,
		OperandSeparator,
		ins.Source.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// MACR (+/-)S1,S2,D / MACR (+/-)S2,S1,D
// Page 148
// 1QQQdk11
type Macr_S1S2 struct {
	Source      RegisterPair
	Destination Register
	Sign        bool
}

func (ins *Macr_S1S2) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	ins.Sign = boolean(bit8(opcode, 2))
	ins.Source = multiplyPairs(mask8(opcode, 4, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Macr_S1S2) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMacr,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source.First,
		OperandSeparator,
		ins.Source.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// MAX A, B
// Page 152
// 00011101
type Max struct{}

func (ins *Max) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	writes.RecordWrite(RegisterB)
	return writes, true
}

func (ins *Max) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMax,
		ColumnSeparator,
		RegisterA,
		OperandSeparator,
		RegisterB,
	)
}

// MAXM A, B
// Page 153
// 00010101
type Maxm struct{}

func (ins *Maxm) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	writes.RecordWrite(RegisterB)
	return writes, true
}

func (ins *Maxm) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMaxm,
		ColumnSeparator,
		RegisterA,
		OperandSeparator,
		RegisterB,
	)
}

// MPY (+/-)S1,S2,D / MPY (+/-)S2,S1,D
// Page 186
// 1QQQdk00
type Mpy_S1S2D struct {
	Source      RegisterPair
	Destination Register
	Sign        bool
}

func (ins *Mpy_S1S2D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMpy,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source.First,
		OperandSeparator,
		ins.Source.Second,
		OperandSeparator,
		ins.Destination,
	)
}

func (ins *Mpy_S1S2D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	ins.Sign = boolean(bit8(opcode, 2))
	ins.Source = multiplyPairs(mask8(opcode, 4, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

// MPYR (+/-)S1,S2,D / MPYR (+/-)S2,S1,D
// Page 190
// 1QQQdk01
type Mpyr_S1S2D struct {
	Source      RegisterPair
	Destination Register
	Sign        bool
}

func (ins *Mpyr_S1S2D) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	ins.Sign = boolean(bit8(opcode, 2))
	ins.Source = multiplyPairs(mask8(opcode, 4, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Mpyr_S1S2D) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicMpyr,
		ColumnSeparator,
		ifThen(ins.Sign, Minus, EmptyToken),
		ins.Source.First,
		OperandSeparator,
		ins.Source.Second,
		OperandSeparator,
		ins.Destination,
	)
}

// NEG D
// Page 194
// 0011d110
type Neg struct {
	Destination Register
}

func (ins *Neg) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Neg) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicNeg,
		ColumnSeparator,
		ins.Destination,
	)
}

// NOT D
// Page 200
// 0001d111
type Not struct {
	Destination Register
}

func (ins *Not) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Not) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicNot,
		ColumnSeparator,
		ins.Destination,
	)
}

// OR S,D
// Page 201
// 01JJd010
type Or_SD struct {
	Source      Register
	Destination Register
}

func (ins *Or_SD) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUSourceOperands(mask8(opcode, 4, 2))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Or_SD) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicOr,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// RND D
// Page 214
// 0001d001
type Rnd struct {
	Destination Register
}

func (ins *Rnd) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Rnd) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicRnd,
		ColumnSeparator,
		ins.Destination,
	)
}

// ROL D
// Page 216
// 0011d111
type Rol struct {
	Destination Register
}

func (ins *Rol) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Rol) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicRol,
		ColumnSeparator,
		ins.Destination,
	)
}

// ROR D
// Page 218
// 0010d111
type Ror struct {
	Destination Register
}

func (ins *Ror) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordPartialWrite(ins.Destination)
	return writes, true
}

func (ins *Ror) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicRor,
		ColumnSeparator,
		ins.Destination,
	)
}

// SBC S,D
// Page 222
// 001Jd101
type Sbc struct {
	Source      Register
	Destination Register
}

func (ins *Sbc) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = inputRegister(bit8(opcode, 4))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Sbc) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicSbc,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// SUB S,D
// Page 225
// 0JJJd100
type Sub_SD struct {
	Source      Register
	Destination Register
}

func (ins *Sub_SD) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = jjj(mask8(opcode, 4, 3), bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)

	if ins.Source == RegisterInvalid {
		return 0, false
	}

	return writes, true
}

func (ins *Sub_SD) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicSub,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// SUBL S,D
// Page 227
// 0001d110
type Subl struct {
	Source      Register
	Destination Register
}

func (ins *Subl) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = oppositeAccumulator(bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Subl) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicSubl,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// SUBR S,D
// Page 229
// 0000d110
type Subr struct {
	Source      Register
	Destination Register
}

func (ins *Subr) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = oppositeAccumulator(bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Subr) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicSubr,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// TFR S,D
// Page 232
// 0JJJd001
type Tfr struct {
	Source      Register
	Destination Register
}

func (ins *Tfr) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Source = dataALUOperands3(mask8(opcode, 4, 3), bit8(opcode, 3))
	ins.Destination = accumulator(bit8(opcode, 3))
	writes.RecordWrite(ins.Destination)
	return writes, true
}

func (ins *Tfr) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicTfr,
		ColumnSeparator,
		ins.Source,
		OperandSeparator,
		ins.Destination,
	)
}

// TST S
// Page 235
// 0000d011
type Tst struct {
	Destination Register
}

func (ins *Tst) DecodeALU(opcode uint8) (AccumulatorWrites, bool) {
	writes := AccumulatorWrites(0)
	ins.Destination = accumulator(bit8(opcode, 3))
	return writes, true
}

func (ins *Tst) Disassemble(w TokenWriter) error {
	return w.Write(
		MnemonicTst,
		ColumnSeparator,
		ins.Destination,
	)
}
