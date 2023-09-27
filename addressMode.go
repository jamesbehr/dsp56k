package dsp56k

import (
	"errors"
	"fmt"
	"io"
)

type EffectiveAddress struct {
	Mode            AddressMode
	AddressRegister uint32
}

func EA(mode AddressMode, reg uint32) EffectiveAddress {
	return EffectiveAddress{mode, reg}
}

type AddressMode int

const (
	AddressModeInvalid AddressMode = iota
	AddressModePostDecrementOffset
	AddressModePostIncrementOffset
	AddressModePostDecrement
	AddressModePostIncrement
	AddressModeNoUpdate
	AddressModeIndexed
	AddressModePreDecrement
)

func (ea EffectiveAddress) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	if ea.AddressRegister > 7 {
		return 0, errors.New("invalid address register")
	}

	switch ea.Mode {
	case AddressModePostDecrementOffset:
		return fmt.Fprintf(w, "(r%[1]d)-n%[1]d", ea.AddressRegister)
	case AddressModePostIncrementOffset:
		return fmt.Fprintf(w, "(r%[1]d)+n%[1]d", ea.AddressRegister)
	case AddressModePostDecrement:
		return fmt.Fprintf(w, "(r%d)-", ea.AddressRegister)
	case AddressModePostIncrement:
		return fmt.Fprintf(w, "(r%d)+", ea.AddressRegister)
	case AddressModeNoUpdate:
		return fmt.Fprintf(w, "(r%d)", ea.AddressRegister)
	case AddressModeIndexed:
		return fmt.Fprintf(w, "(r%[1]d+n%[1]d)", ea.AddressRegister)
	case AddressModePreDecrement:
		return fmt.Fprintf(w, "-(r%d)", ea.AddressRegister)
	}

	return 0, errors.New("invalid effective address mode")
}
