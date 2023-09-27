package dsp56k

import (
	"errors"
	"fmt"
	"io"
)

type Condition int

const (
	ConditionInvalid Condition = iota
	ConditionCC                // (HS) carry clear (higher or same) C=0
	ConditionCS                // (LO) carry set (lower) C=1
	ConditionEC                // extension clear E=0
	ConditionEQ                // equal Z=1
	ConditionES                // extension set E=1
	ConditionGE                // greater than or equal N ⊕ V=0
	ConditionGT                // greater than Z+(N ⊕ V)=0
	ConditionLC                // limit clear L=0
	ConditionLE                // less than or equal Z+(N ⊕ V)=1
	ConditionLS                // limit set L=1
	ConditionLT                // less than N ⊕ V=1
	ConditionMI                // minus N=1
	ConditionNE                // not equal Z=0
	ConditionNR                // normalized Z+(U•E)=1
	ConditionPL                // plus N=0
	ConditionNN                // not normalized Z+(U•E)=
)

func (c Condition) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	switch c {
	case ConditionCC:
		return fmt.Fprint(w, "cc")
	case ConditionCS:
		return fmt.Fprint(w, "cs")
	case ConditionEC:
		return fmt.Fprint(w, "ec")
	case ConditionEQ:
		return fmt.Fprint(w, "eq")
	case ConditionES:
		return fmt.Fprint(w, "es")
	case ConditionGE:
		return fmt.Fprint(w, "ge")
	case ConditionGT:
		return fmt.Fprint(w, "gt")
	case ConditionLC:
		return fmt.Fprint(w, "lc")
	case ConditionLE:
		return fmt.Fprint(w, "le")
	case ConditionLS:
		return fmt.Fprint(w, "ls")
	case ConditionLT:
		return fmt.Fprint(w, "lt")
	case ConditionMI:
		return fmt.Fprint(w, "mi")
	case ConditionNE:
		return fmt.Fprint(w, "ne")
	case ConditionNR:
		return fmt.Fprint(w, "nr")
	case ConditionPL:
		return fmt.Fprint(w, "pl")
	case ConditionNN:
		return fmt.Fprint(w, "nn")
	}

	return 0, errors.New("invalid condition")
}
