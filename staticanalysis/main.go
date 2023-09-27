package staticanalysis

import "github.com/jamesbehr/dsp56k"

func relativeAddress(pc uint32, displacement int32) uint32 {
	return uint32(int32(pc) + displacement)
}

// Halts returns true when the instruction changes the program counter always,
// meaning the instruction after this one is never executed.
func Halts(ins dsp56k.Instruction) bool {
	switch ins.(type) {
	case *dsp56k.Rti, *dsp56k.Rts, *dsp56k.Bra_xxxx, *dsp56k.Bra_xxx, *dsp56k.Bra_Rn,
		*dsp56k.Jmp_ea_Abs, *dsp56k.Jmp_xxx, *dsp56k.Jmp_ea:
		return true
	}
	return false
}

type RefType int

const (
	RefBranch RefType = iota
	RefPeripheral
	RefLoop
	RefMemory
)

type AddressRef struct {
	Type    RefType
	Region  dsp56k.Memory
	Address uint32
}

func AbsoluteAddressReferences(ins dsp56k.Instruction, pc uint32) []AddressRef {
	switch ins := ins.(type) {
	case *dsp56k.ParallelInstruction:
		switch move := ins.Move.(type) {
		case *dsp56k.Movex_aa:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryX, move.Address},
			}
		case *dsp56k.Movey_aa:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryY, move.Address},
			}
		case *dsp56k.Movex_ea_Abs:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryX, move.Address},
			}
		case *dsp56k.Movey_ea_Abs:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryY, move.Address},
			}
		case *dsp56k.Movexr_ea_Abs:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryX, move.Address},
			}
		case *dsp56k.Moveyr_ea_Abs:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryY, move.Address},
			}
		case *dsp56k.Movel_aa:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryL, move.Address},
			}
		case *dsp56k.Movel_ea_Abs:
			return []AddressRef{
				{RefMemory, dsp56k.MemoryL, move.Address},
			}
		}
	case *dsp56k.Do_ea:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, ins.LoopAddress + 1},
		}
	case *dsp56k.Do_xxx:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, ins.LoopAddress + 1},
		}
	case *dsp56k.Do_S:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, ins.LoopAddress + 1},
		}
	case *dsp56k.DoForever:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, ins.LoopAddress + 1},
		}
	case *dsp56k.Dor_ea:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, relativeAddress(pc, ins.LoopDisplacement+1)},
		}
	case *dsp56k.Dor_xxx:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, relativeAddress(pc, ins.LoopDisplacement+1)},
		}
	case *dsp56k.Dor_S:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, relativeAddress(pc, ins.LoopDisplacement+1)},
		}
	case *dsp56k.DorForever:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, relativeAddress(pc, ins.LoopDisplacement+1)},
		}
	case *dsp56k.Bchg_ea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bchg_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bclr_ea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bclr_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bset_ea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bset_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Btst_ea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Btst_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Btst_pp:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Btst_qq:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bset_pp:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bset_qq:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bclr_pp:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bclr_qq:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bchg_pp:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bchg_qq:
		return []AddressRef{
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Lra_xxxx:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Do_aa:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, ins.LoopAddress + 1},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Dor_aa:
		return []AddressRef{
			{RefLoop, dsp56k.MemoryP, relativeAddress(pc, ins.LoopDisplacement+1)},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Movec_ea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Movec_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Movem_ea_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
		}
	case *dsp56k.Movem_aa:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
		}
	case *dsp56k.Movep_ppea:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_ppea_Abs:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Movep_ppea_Imm:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Xqqea:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryX, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Xqqea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
			{RefPeripheral, dsp56k.MemoryX, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Xqqea_Imm:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryX, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Yqqea:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Yqqea_Abs:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Yqqea_Imm:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eapp:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eapp_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eapp_Imm:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eaqq:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eaqq_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_eaqq_Imm:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_Spp:
		return []AddressRef{
			{RefPeripheral, ins.PeripheralMemory, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_SXqq:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryX, ins.PeripheralAddress},
		}
	case *dsp56k.Movep_SYqq:
		return []AddressRef{
			{RefPeripheral, dsp56k.MemoryY, ins.PeripheralAddress},
		}
	case *dsp56k.Plock_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
		}
	case *dsp56k.Plockr:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Punlock_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, ins.Address},
		}
	case *dsp56k.Punlockr:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Rep_aa:
		return []AddressRef{
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Vsl_Abs:
		return []AddressRef{
			{RefMemory, dsp56k.MemoryL, ins.Address},
		}
	case *dsp56k.Brclr_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Brclr_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Brclr_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Brset_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Brset_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Brset_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bsclr_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bsclr_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bsclr_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bsset_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Bsset_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bsset_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jclr_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Jclr_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jclr_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jsclr_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Jsclr_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jsclr_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jset_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Jset_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jset_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jsset_aa:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefMemory, ins.Memory, ins.Address},
		}
	case *dsp56k.Jsset_pp:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Jsset_qq:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
			{RefPeripheral, ins.Memory, ins.PeripheralAddress},
		}
	case *dsp56k.Bsclr_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Brclr_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Brclr_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bcc_xxxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bcc_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bra_xxxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bra_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Brset_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Brset_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.BScc_xxxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.BScc_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bsclr_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bsr_xxxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bsr_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bsset_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Bsset_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, relativeAddress(pc, ins.Displacement)},
		}
	case *dsp56k.Jcc_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jcc_ea_Abs:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jclr_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jclr_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jmp_ea_Abs:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jmp_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jscc_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jscc_ea_Abs:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsclr_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsclr_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jset_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jset_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsr_ea_Abs:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsr_xxx:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsset_ea:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	case *dsp56k.Jsset_S:
		return []AddressRef{
			{RefBranch, dsp56k.MemoryP, ins.JumpAddress},
		}
	}

	return []AddressRef{}
}
