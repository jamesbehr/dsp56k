package staticanalysis

import (
	"testing"

	"github.com/jamesbehr/dsp56k"

	"github.com/stretchr/testify/require"
)

func TestAbsolute(t *testing.T) {
	type SymbolTestCase struct {
		Instruction dsp56k.Instruction
		Refs        []AddressRef
	}

	programCounter := 0x1234
	displacement := int32(0xcafe - programCounter)

	branch := AddressRef{RefBranch, dsp56k.MemoryP, 0xcafe}
	loop := AddressRef{RefLoop, dsp56k.MemoryP, 0xcafe}
	yperipheral := AddressRef{RefPeripheral, dsp56k.MemoryY, 0xcafe}
	yperipheral2 := AddressRef{RefPeripheral, dsp56k.MemoryY, 0xdead}
	xperipheral := AddressRef{RefPeripheral, dsp56k.MemoryX, 0xcafe}
	pmem := AddressRef{RefMemory, dsp56k.MemoryP, 0xcafe}
	pmem2 := AddressRef{RefMemory, dsp56k.MemoryP, 0xdead}
	xmem := AddressRef{RefMemory, dsp56k.MemoryX, 0xcafe}
	ymem := AddressRef{RefMemory, dsp56k.MemoryY, 0xcafe}
	ymem2 := AddressRef{RefMemory, dsp56k.MemoryY, 0xdead}
	lmem := AddressRef{RefMemory, dsp56k.MemoryL, 0xcafe}

	tests := []SymbolTestCase{
		{&dsp56k.Bcc_xxxx{
			Condition:    dsp56k.ConditionGE,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bcc_xxx{
			Condition:    dsp56k.ConditionGE,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bchg_ea_Abs{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bchg_aa{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bchg_pp{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bchg_qq{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bclr_ea_Abs{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bclr_aa{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bclr_pp{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bclr_qq{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bra_xxxx{
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bra_xxx{
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Brclr_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Brclr_aa{
			Address:      0xdead,
			Memory:       dsp56k.MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Brclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Brclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Brclr_S{
			Source:       dsp56k.RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Brset_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Brset_aa{
			Address:      0xdead,
			Memory:       dsp56k.MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Brset_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Brset_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Brset_S{
			Source:       dsp56k.RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.BScc_xxxx{
			Condition:    dsp56k.ConditionGE,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.BScc_xxx{
			Condition:    dsp56k.ConditionGE,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bsclr_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bsclr_aa{
			Address:      0xdead,
			Memory:       dsp56k.MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Bsclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Bsclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Bsclr_S{
			Source:       dsp56k.RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bset_ea_Abs{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bset_aa{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Bset_pp{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bset_qq{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Bsr_xxxx{
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bsr_xxx{
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bsset_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			Displacement:     displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Bsset_aa{
			Address:      0xdead,
			Memory:       dsp56k.MemoryY,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Bsset_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Bsset_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			Displacement:      displacement,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Bsset_S{
			Source:       dsp56k.RegisterB,
			BitNumber:    12,
			Displacement: displacement,
		}, []AddressRef{branch}},
		{&dsp56k.Btst_ea_Abs{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Btst_aa{
			Address:   0xcafe,
			Memory:    dsp56k.MemoryY,
			BitNumber: 12,
		}, []AddressRef{ymem}},
		{&dsp56k.Btst_pp{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Btst_qq{
			PeripheralAddress: 0xcafe,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Do_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			LoopAddress:      0xcafd,
		}, []AddressRef{loop}},
		{&dsp56k.Do_aa{
			Address:     0xdead,
			Memory:      dsp56k.MemoryY,
			LoopAddress: 0xcafd,
		}, []AddressRef{ymem2, loop}},
		{&dsp56k.Do_xxx{
			Immediate:   0x1234,
			LoopAddress: 0xcafd,
		}, []AddressRef{loop}},
		{&dsp56k.Do_S{
			Source:      dsp56k.RegisterB,
			LoopAddress: 0xcafd,
		}, []AddressRef{loop}},
		{&dsp56k.DoForever{
			LoopAddress: 0xcafd,
		}, []AddressRef{loop}},
		{&dsp56k.Dor_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			LoopDisplacement: displacement - 1,
		}, []AddressRef{loop}},
		{&dsp56k.Dor_aa{
			Address:          0xdead,
			Memory:           dsp56k.MemoryY,
			LoopDisplacement: displacement - 1,
		}, []AddressRef{ymem2, loop}},
		{&dsp56k.Dor_xxx{
			Immediate:        0x1234,
			LoopDisplacement: displacement - 1,
		}, []AddressRef{loop}},
		{&dsp56k.Dor_S{
			Source:           dsp56k.RegisterB,
			LoopDisplacement: displacement - 1,
		}, []AddressRef{loop}},
		{&dsp56k.DorForever{
			LoopDisplacement: displacement - 1,
		}, []AddressRef{loop}},
		{&dsp56k.Jcc_xxx{
			Condition:   dsp56k.ConditionGE,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jcc_ea_Abs{
			JumpAddress: 0xcafe,
			Condition:   dsp56k.ConditionGE,
		}, []AddressRef{branch}},
		{&dsp56k.Jclr_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jclr_aa{
			Address:     0xdead,
			Memory:      dsp56k.MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Jclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jclr_S{
			Source:      dsp56k.RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jmp_ea_Abs{
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jmp_xxx{
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jscc_xxx{
			Condition:   dsp56k.ConditionGE,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jscc_ea_Abs{
			JumpAddress: 0xcafe,
			Condition:   dsp56k.ConditionGE,
		}, []AddressRef{branch}},
		{&dsp56k.Jsclr_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jsclr_aa{
			Address:     0xdead,
			Memory:      dsp56k.MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Jsclr_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jsclr_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jsclr_S{
			Source:      dsp56k.RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jset_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jset_aa{
			Address:     0xdead,
			Memory:      dsp56k.MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Jset_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jset_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jset_S{
			Source:      dsp56k.RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jsr_ea_Abs{
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jsr_xxx{
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jsset_ea{
			EffectiveAddress: dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:           dsp56k.MemoryY,
			BitNumber:        12,
			JumpAddress:      0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Jsset_aa{
			Address:     0xdead,
			Memory:      dsp56k.MemoryY,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{ymem2, branch}},
		{&dsp56k.Jsset_pp{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jsset_qq{
			PeripheralAddress: 0xdead,
			Memory:            dsp56k.MemoryY,
			BitNumber:         12,
			JumpAddress:       0xcafe,
		}, []AddressRef{yperipheral2, branch}},
		{&dsp56k.Jsset_S{
			Source:      dsp56k.RegisterB,
			BitNumber:   12,
			JumpAddress: 0xcafe,
		}, []AddressRef{branch}},
		{&dsp56k.Lra_xxxx{
			DestinationAddress: dsp56k.RegisterB,
			Displacement:       displacement,
		}, []AddressRef{pmem}},
		{&dsp56k.Movec_ea_Abs{
			IsWrite:           false,
			Memory:            dsp56k.MemoryY,
			Address:           0xcafe,
			ProgramController: dsp56k.RegisterB,
		}, []AddressRef{ymem}},
		{&dsp56k.Movec_aa{
			IsWrite:           false,
			Address:           0xcafe,
			Memory:            dsp56k.MemoryY,
			ProgramController: dsp56k.RegisterB,
		}, []AddressRef{ymem}},
		{&dsp56k.Movem_ea_Abs{
			IsWrite:             false,
			Address:             0xcafe,
			SourceOrDestination: dsp56k.RegisterB,
		}, []AddressRef{pmem}},
		{&dsp56k.Movem_aa{
			IsWrite:             false,
			Address:             0xcafe,
			SourceOrDestination: dsp56k.RegisterB,
		}, []AddressRef{pmem}},
		{&dsp56k.Movep_ppea{
			PeripheralMemory:  dsp56k.MemoryY,
			IsWrite:           false,
			EffectiveAddress:  dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_ppea_Abs{
			PeripheralMemory:  dsp56k.MemoryY,
			IsWrite:           false,
			Address:           0xdead,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral, ymem2}},
		{&dsp56k.Movep_ppea_Imm{
			PeripheralMemory:  dsp56k.MemoryY,
			Immediate:         0x1234,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_Xqqea{
			IsWrite:           false,
			EffectiveAddress:  dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{xperipheral}},
		{&dsp56k.Movep_Xqqea_Abs{
			IsWrite:           false,
			Address:           0xdead,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{xperipheral, ymem2}},
		{&dsp56k.Movep_Xqqea_Imm{
			Immediate:         0x1234,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{xperipheral}},
		{&dsp56k.Movep_Yqqea{
			IsWrite:           false,
			EffectiveAddress:  dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_Yqqea_Abs{
			IsWrite:           false,
			Address:           0xdead,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral, ymem2}},
		{&dsp56k.Movep_Yqqea_Imm{
			Immediate:         0x1234,
			Memory:            dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_eapp{
			PeripheralMemory:  dsp56k.MemoryY,
			IsWrite:           false,
			EffectiveAddress:  dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_eapp_Abs{
			PeripheralMemory:  dsp56k.MemoryY,
			IsWrite:           false,
			Address:           0xdead,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral, pmem2}},
		{&dsp56k.Movep_eapp_Imm{
			PeripheralMemory:  dsp56k.MemoryY,
			Immediate:         0x1234,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_eaqq{
			IsWrite:           false,
			EffectiveAddress:  dsp56k.EA(dsp56k.AddressModePostIncrement, 5),
			PeripheralMemory:  dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_eaqq_Abs{
			IsWrite:           false,
			Address:           0xdead,
			PeripheralMemory:  dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral, pmem2}},
		{&dsp56k.Movep_eaqq_Imm{
			Immediate:         0x1234,
			PeripheralMemory:  dsp56k.MemoryY,
			PeripheralAddress: 0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_Spp{
			PeripheralMemory:    dsp56k.MemoryY,
			IsWrite:             false,
			SourceOrDestination: dsp56k.RegisterB,
			PeripheralAddress:   0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Movep_SXqq{
			IsWrite:             false,
			SourceOrDestination: dsp56k.RegisterB,
			PeripheralAddress:   0xcafe,
		}, []AddressRef{xperipheral}},
		{&dsp56k.Movep_SYqq{
			IsWrite:             false,
			SourceOrDestination: dsp56k.RegisterB,
			PeripheralAddress:   0xcafe,
		}, []AddressRef{yperipheral}},
		{&dsp56k.Plock_Abs{
			Address: 0xcafe,
		}, []AddressRef{pmem}},
		{&dsp56k.Plockr{
			Displacement: displacement,
		}, []AddressRef{pmem}},
		{&dsp56k.Punlock_Abs{
			Address: 0xcafe,
		}, []AddressRef{pmem}},
		{&dsp56k.Punlockr{
			Displacement: displacement,
		}, []AddressRef{pmem}},
		{&dsp56k.Rep_aa{
			Address: 0xcafe,
			Memory:  dsp56k.MemoryY,
		}, []AddressRef{ymem}},
		{&dsp56k.Vsl_Abs{
			Source:  dsp56k.RegisterB,
			Address: 0xcafe,
			Bit:     12,
		}, []AddressRef{lmem}},

		// Parallel Instructions
		{dsp56k.Parallel(&dsp56k.Movex_aa{
			SourceOrDestination: dsp56k.RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &dsp56k.Max{}), []AddressRef{xmem}},
		{dsp56k.Parallel(&dsp56k.Movey_aa{
			SourceOrDestination: dsp56k.RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &dsp56k.Max{}), []AddressRef{ymem}},
		{dsp56k.Parallel(&dsp56k.Movex_ea_Abs{
			SourceOrDestination: dsp56k.RegisterA,
			IsWrite:             false,
			Address:             0xcafe,
		}, &dsp56k.Max{}), []AddressRef{xmem}},
		{dsp56k.Parallel(&dsp56k.Movey_ea_Abs{
			SourceOrDestination: dsp56k.RegisterA,
			IsWrite:             false,
			Address:             0xcafe,
		}, &dsp56k.Max{}), []AddressRef{ymem}},
		{dsp56k.Parallel(&dsp56k.Movexr_ea_Abs{
			Address:             0xcafe,
			IsWrite:             false,
			Source:              dsp56k.RegisterX,
			Destination:         dsp56k.RegisterY,
			SourceOrDestination: dsp56k.RegisterA,
		}, &dsp56k.Max{}), []AddressRef{xmem}},
		{dsp56k.Parallel(&dsp56k.Moveyr_ea_Abs{
			Address:             0xcafe,
			IsWrite:             false,
			Source:              dsp56k.RegisterX,
			Destination:         dsp56k.RegisterY,
			SourceOrDestination: dsp56k.RegisterA,
		}, &dsp56k.Max{}), []AddressRef{ymem}},
		{dsp56k.Parallel(&dsp56k.Movel_aa{
			Address:             0xcafe,
			SourceOrDestination: dsp56k.RegisterA,
			IsWrite:             false,
		}, &dsp56k.Max{}), []AddressRef{lmem}},
		{dsp56k.Parallel(&dsp56k.Movel_ea_Abs{
			SourceOrDestination: dsp56k.RegisterA,
			Address:             0xcafe,
			IsWrite:             false,
		}, &dsp56k.Max{}), []AddressRef{lmem}},
	}

	for _, test := range tests {
		got := AbsoluteAddressReferences(test.Instruction, uint32(programCounter))
		require.ElementsMatch(t, got, test.Refs, "%#v", test.Instruction)
	}
}
