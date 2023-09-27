package dsp56k

import (
	"fmt"
	"io"
)

type TokenWriter interface {
	Write(...Token) error
}

type Location struct {
	Region  Memory
	Address uint32
}

type SymbolTable interface {
	LookupSymbol(Memory, uint32) (string, bool)
}

type Options struct {
	SymbolTable
}

type Token interface {
	WriteOperand(w io.Writer, opts Options, pc uint32) (int, error)
}

func ifThen(test bool, whenTrue, whenFalse Token) Token {
	if test {
		return whenTrue
	}

	return whenFalse
}

type relativeAddrToken int32

func (r relativeAddrToken) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	if opts.SymbolTable != nil {
		sym, ok := opts.LookupSymbol(MemoryP, uint32(int32(pc)+int32(r)))
		if ok {
			return fmt.Fprint(w, sym)
		}
	}

	if r == 0 {
		return fmt.Fprint(w, "*")
	} else if r < 0 {
		return fmt.Fprintf(w, "*-$%x", -r)
	}

	return fmt.Fprintf(w, "*+$%x", r)
}

func relativeAddr(displacement int32) Token {
	return relativeAddrToken(displacement)
}

type absoluteAddrToken struct {
	Memory  Memory
	Address uint32
}

func (r absoluteAddrToken) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	if opts.SymbolTable != nil {
		sym, ok := opts.LookupSymbol(r.Memory, r.Address)
		if ok {
			return fmt.Fprint(w, sym)
		}
	}

	return fmt.Fprintf(w, "$%x", r.Address)
}

func absAddr(mem Memory, addr uint32) Token {
	return absoluteAddrToken{mem, addr}
}

type numberToken struct {
	value        int32
	showPositive bool
}

func (n numberToken) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	if n.value < 0 {
		return fmt.Fprintf(w, "-$%x", -n.value)
	} else if n.showPositive {
		return fmt.Fprintf(w, "+$%x", n.value)
	}

	return fmt.Fprintf(w, "$%x", n.value)
}

func immediate(v uint32) numberToken {
	return numberToken{int32(v), false}
}

func abs(v int32) uint32 {
	if v < 0 {
		return uint32(-v)
	}

	return uint32(v)
}

type stringToken string

func (s stringToken) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	return w.Write([]byte(s))
}

const (
	ForceIOShortAddressMode        = stringToken("<<")
	ForceShortAddressMode          = stringToken("<")
	ForceLongAddressMode           = stringToken(">")
	ImmediateAddressMode           = stringToken("#")
	ForceLongImmediateAddressMode  = stringToken("#>")
	ForceShortImmediateAddressMode = stringToken("#<")

	ColumnSeparator  = stringToken(" ")
	OperandSeparator = stringToken(",")
	EmptyToken       = stringToken("")
	UpdateCCRToken   = stringToken(".u")

	Plus       = stringToken("+")
	Minus      = stringToken("-")
	OpenBrace  = stringToken("(")
	CloseBrace = stringToken(")")

	ZeroToken = stringToken("0")
	OneToken  = stringToken("1")
)

// We're using a separate mnemonic token because this allows us to potentially
// rename them easily
type mnemonicToken string

func (s mnemonicToken) WriteOperand(w io.Writer, opts Options, pc uint32) (int, error) {
	return w.Write([]byte(s))
}

const (
	MnemonicAbs        mnemonicToken = "abs"
	MnemonicAdc        mnemonicToken = "adc"
	MnemonicAdd        mnemonicToken = "add"
	MnemonicAddl       mnemonicToken = "addl"
	MnemonicAddr       mnemonicToken = "addr"
	MnemonicAnd        mnemonicToken = "and"
	MnemonicAndi       mnemonicToken = "andi"
	MnemonicAsl        mnemonicToken = "asl"
	MnemonicAsr        mnemonicToken = "asr"
	MnemonicBchg       mnemonicToken = "bchg"
	MnemonicBclr       mnemonicToken = "bclr"
	MnemonicBset       mnemonicToken = "bset"
	MnemonicBtst       mnemonicToken = "btst"
	MnemonicClr        mnemonicToken = "clr"
	MnemonicCmp        mnemonicToken = "cmp"
	MnemonicCmpm       mnemonicToken = "cmpm"
	MnemonicDiv        mnemonicToken = "div"
	MnemonicDo         mnemonicToken = "do"
	MnemonicDoForever  mnemonicToken = "do forever"
	MnemonicEnddo      mnemonicToken = "enddo"
	MnemonicEor        mnemonicToken = "eor"
	MnemonicJclr       mnemonicToken = "jclr"
	MnemonicJmp        mnemonicToken = "jmp"
	MnemonicJs         mnemonicToken = "js"
	MnemonicJsclr      mnemonicToken = "jsclr"
	MnemonicJset       mnemonicToken = "jset"
	MnemonicJsr        mnemonicToken = "jsr"
	MnemonicJsset      mnemonicToken = "jsset"
	MnemonicLsl        mnemonicToken = "lsl"
	MnemonicLsr        mnemonicToken = "lsr"
	MnemonicLua        mnemonicToken = "lua"
	MnemonicMac        mnemonicToken = "mac"
	MnemonicMaci       mnemonicToken = "maci"
	MnemonicMacr       mnemonicToken = "macr"
	MnemonicMacri      mnemonicToken = "macri"
	MnemonicMove       mnemonicToken = "move"
	MnemonicMovec      mnemonicToken = "movec"
	MnemonicMovep      mnemonicToken = "movep"
	MnemonicMpy        mnemonicToken = "mpy"
	MnemonicMpyi       mnemonicToken = "mpyi"
	MnemonicMpyr       mnemonicToken = "mpyr"
	MnemonicMpyri      mnemonicToken = "mpyri"
	MnemonicNeg        mnemonicToken = "neg"
	MnemonicNop        mnemonicToken = "nop"
	MnemonicNorm       mnemonicToken = "norm"
	MnemonicNot        mnemonicToken = "not"
	MnemonicOr         mnemonicToken = "or"
	MnemonicOri        mnemonicToken = "ori"
	MnemonicRep        mnemonicToken = "rep"
	MnemonicReset      mnemonicToken = "reset"
	MnemonicRnd        mnemonicToken = "rnd"
	MnemonicRol        mnemonicToken = "rol"
	MnemonicRor        mnemonicToken = "ror"
	MnemonicRti        mnemonicToken = "rti"
	MnemonicRts        mnemonicToken = "rts"
	MnemonicSbc        mnemonicToken = "sbc"
	MnemonicStart      mnemonicToken = "start"
	MnemonicStop       mnemonicToken = "stop"
	MnemonicSub        mnemonicToken = "sub"
	MnemonicSubl       mnemonicToken = "subl"
	MnemonicSubr       mnemonicToken = "subr"
	MnemonicSwi        mnemonicToken = "swi"
	MnemonicTfr        mnemonicToken = "tfr"
	MnemonicTst        mnemonicToken = "tst"
	MnemonicWait       mnemonicToken = "wait"
	MnemonicIllegal    mnemonicToken = "illegal"
	MnemonicInc        mnemonicToken = "inc"
	MnemonicDec        mnemonicToken = "dec"
	MnemonicDebug      mnemonicToken = "debug"
	MnemonicBsset      mnemonicToken = "bsset"
	MnemonicBsclr      mnemonicToken = "bsclr"
	MnemonicBrset      mnemonicToken = "brset"
	MnemonicBrclr      mnemonicToken = "brclr"
	MnemonicBs         mnemonicToken = "bs"
	MnemonicBsr        mnemonicToken = "bsr"
	MnemonicBra        mnemonicToken = "bra"
	MnemonicDor        mnemonicToken = "dor"
	MnemonicDorForever mnemonicToken = "dor forever"
	MnemonicMerge      mnemonicToken = "merge"
	MnemonicClb        mnemonicToken = "clb"
	MnemonicIf         mnemonicToken = "if"
	MnemonicLra        mnemonicToken = "lra"
	MnemonicNormf      mnemonicToken = "normf"
	MnemonicInsert     mnemonicToken = "insert"
	MnemonicMax        mnemonicToken = "max"
	MnemonicMaxm       mnemonicToken = "maxm"
	MnemonicPflush     mnemonicToken = "pflush"
	MnemonicPfree      mnemonicToken = "pfree"
	MnemonicPlock      mnemonicToken = "plock"
	MnemonicPlockr     mnemonicToken = "plockr"
	MnemonicPunlock    mnemonicToken = "punlock"
	MnemonicPunlockr   mnemonicToken = "punlockr"
	MnemonicExtractu   mnemonicToken = "extractu"
	MnemonicExtract    mnemonicToken = "extract"
	MnemonicTrap       mnemonicToken = "trap"
	MnemonicDmac       mnemonicToken = "dmac"
	MnemonicBrk        mnemonicToken = "brk"
	MnemonicCmpu       mnemonicToken = "cmpu"
	MnemonicPflushun   mnemonicToken = "pflushun"
	MnemonicVsl        mnemonicToken = "vsl"
	MnemonicJcc        mnemonicToken = "j"
	MnemonicTcc        mnemonicToken = "t"
	MnemonicBcc        mnemonicToken = "b"
)
