package dsp56k

import (
	"errors"
	"fmt"
	"io"
)

type Register int

const (
	RegisterInvalid Register = iota

	// Data ALU Registers
	RegisterA10 // Accumulator A = A1:A0 (48-bit)
	RegisterAB  // Accumulators A and B = A1:B1 (48-bit)
	RegisterBA  // Accumulators B and A = B1:A1 (48-bit)
	RegisterB10 // Accumulator B = B1:B0 (48-bit)
	RegisterA   // Accumulator A = A2:A1:A0 (56-bit)
	RegisterB   // Accumulator B = B2:B1:B0 (56-bit)
	RegisterX   // Input Register X = X1:X0 (48-bit)
	RegisterY   // Input Register Y = Y1:Y0 (48-bit)
	RegisterA0  // Accumulator Register (24-bit)
	RegisterA1  // Accumulator Register (24-bit)
	RegisterA2  // Accumulator Register (8-bit)
	RegisterB0  // Accumulator Register (24-bit)
	RegisterB1  // Accumulator Register (24-bit)
	RegisterB2  // Accumulator Register (8-bit)
	RegisterX0  // Input Register (24-bit)
	RegisterX1  // Input Register (24-bit)
	RegisterY0  // Input Register (24-bit)
	RegisterY1  // Input Register (24-bit)

	// Address ALU Regsiters
	RegisterR0 // Address Regsiter R0 (24-bit)
	RegisterR1 // Address Regsiter R1 (24-bit)
	RegisterR2 // Address Regsiter R2 (24-bit)
	RegisterR3 // Address Regsiter R3 (24-bit)
	RegisterR4 // Address Regsiter R4 (24-bit)
	RegisterR5 // Address Regsiter R5 (24-bit)
	RegisterR6 // Address Regsiter R6 (24-bit)
	RegisterR7 // Address Regsiter R7 (24-bit)
	RegisterN0 // Address Offset Regiser R0 (24-bit)
	RegisterN1 // Address Offset Regiser R1 (24-bit)
	RegisterN2 // Address Offset Regiser R2 (24-bit)
	RegisterN3 // Address Offset Regiser R3 (24-bit)
	RegisterN4 // Address Offset Regiser R4 (24-bit)
	RegisterN5 // Address Offset Regiser R5 (24-bit)
	RegisterN6 // Address Offset Regiser R6 (24-bit)
	RegisterN7 // Address Offset Regiser R7 (24-bit)
	RegisterM0 // Address Modifier Regsiter M0 (24-bit)
	RegisterM1 // Address Modifier Regsiter M1 (24-bit)
	RegisterM2 // Address Modifier Regsiter M2 (24-bit)
	RegisterM3 // Address Modifier Regsiter M3 (24-bit)
	RegisterM4 // Address Modifier Regsiter M4 (24-bit)
	RegisterM5 // Address Modifier Regsiter M5 (24-bit)
	RegisterM6 // Address Modifier Regsiter M6 (24-bit)
	RegisterM7 // Address Modifier Regsiter M7 (24-bit)

	// Program Control Unit Registers
	RegisterMR  // Mode Regsiter (8-bit)
	RegisterCCR // Mode Regsiter (8-bit)
	RegisterCOM // Chip Operating Mode Mode Regsiter (8-bit)
	RegisterEOM // Extended Chip Operating Mode Regsiter (8-bit)
	RegisterEP
	RegisterVBA // Vector Base Address Regsiter (24-bit). Used to determine interrupt addresses.
	RegisterSC  // System Stack Counter Regsiter (5-bit)
	RegisterSZ  // System Stack Size Regsiter (24-bit)
	RegisterSR  // Status Regsiter (24-bit). EMR:MR:CCR
	RegisterOMR // Operating Mode Register Regsiter (24-bit). SCS:EOM:COM
	RegisterSP  // System Stack Pointer Regsiter (24-bit)
	RegisterSSH // Upper Portiton of the Current Top of the Stack (24-bit)
	RegisterSSL // Lower Portiton of the Current Top of the Stack (24-bit)
	RegisterLA  // Hardware Loop Address Register (24-bit)
	RegisterLC  // Hardware Loop Counter Register (24-bit)
)

func (r Register) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	switch r {
	case RegisterAB:
		return fmt.Fprint(w, "ab")
	case RegisterBA:
		return fmt.Fprint(w, "ba")
	case RegisterA10:
		return fmt.Fprint(w, "a10")
	case RegisterB10:
		return fmt.Fprint(w, "b10")
	case RegisterA:
		return fmt.Fprint(w, "a")
	case RegisterA0:
		return fmt.Fprint(w, "a0")
	case RegisterA1:
		return fmt.Fprint(w, "a1")
	case RegisterA2:
		return fmt.Fprint(w, "a2")
	case RegisterB:
		return fmt.Fprint(w, "b")
	case RegisterB0:
		return fmt.Fprint(w, "b0")
	case RegisterB1:
		return fmt.Fprint(w, "b1")
	case RegisterB2:
		return fmt.Fprint(w, "b2")
	case RegisterX:
		return fmt.Fprint(w, "x")
	case RegisterX0:
		return fmt.Fprint(w, "x0")
	case RegisterX1:
		return fmt.Fprint(w, "x1")
	case RegisterY:
		return fmt.Fprint(w, "y")
	case RegisterY0:
		return fmt.Fprint(w, "y0")
	case RegisterY1:
		return fmt.Fprint(w, "y1")
	case RegisterR0:
		return fmt.Fprint(w, "r0")
	case RegisterR1:
		return fmt.Fprint(w, "r1")
	case RegisterR2:
		return fmt.Fprint(w, "r2")
	case RegisterR3:
		return fmt.Fprint(w, "r3")
	case RegisterR4:
		return fmt.Fprint(w, "r4")
	case RegisterR5:
		return fmt.Fprint(w, "r5")
	case RegisterR6:
		return fmt.Fprint(w, "r6")
	case RegisterR7:
		return fmt.Fprint(w, "r7")
	case RegisterN0:
		return fmt.Fprint(w, "n0")
	case RegisterN1:
		return fmt.Fprint(w, "n1")
	case RegisterN2:
		return fmt.Fprint(w, "n2")
	case RegisterN3:
		return fmt.Fprint(w, "n3")
	case RegisterN4:
		return fmt.Fprint(w, "n4")
	case RegisterN5:
		return fmt.Fprint(w, "n5")
	case RegisterN6:
		return fmt.Fprint(w, "n6")
	case RegisterN7:
		return fmt.Fprint(w, "n7")
	case RegisterM0:
		return fmt.Fprint(w, "m0")
	case RegisterM1:
		return fmt.Fprint(w, "m1")
	case RegisterM2:
		return fmt.Fprint(w, "m2")
	case RegisterM3:
		return fmt.Fprint(w, "m3")
	case RegisterM4:
		return fmt.Fprint(w, "m4")
	case RegisterM5:
		return fmt.Fprint(w, "m5")
	case RegisterM6:
		return fmt.Fprint(w, "m6")
	case RegisterM7:
		return fmt.Fprint(w, "m7")
	case RegisterMR:
		return fmt.Fprint(w, "mr")
	case RegisterCCR:
		return fmt.Fprint(w, "ccr")
	case RegisterCOM:
		return fmt.Fprint(w, "omr")
	case RegisterEOM:
		return fmt.Fprint(w, "eom")
	case RegisterEP:
		return fmt.Fprint(w, "ep")
	case RegisterVBA:
		return fmt.Fprint(w, "vba")
	case RegisterSC:
		return fmt.Fprint(w, "sc")
	case RegisterSZ:
		return fmt.Fprint(w, "sz")
	case RegisterSR:
		return fmt.Fprint(w, "sr")
	case RegisterOMR:
		return fmt.Fprint(w, "omr")
	case RegisterSP:
		return fmt.Fprint(w, "sp")
	case RegisterSSH:
		return fmt.Fprint(w, "ssh")
	case RegisterSSL:
		return fmt.Fprint(w, "ssl")
	case RegisterLA:
		return fmt.Fprint(w, "la")
	case RegisterLC:
		return fmt.Fprint(w, "lc")
	}

	return 0, errors.New("invalid register")
}
