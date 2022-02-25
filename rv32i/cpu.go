package rv32i

type Cpu struct {
	Register [32]int32
	Pc       uint32
	Exit     uint32
}

func MovePc(cpu Cpu, addr uint32) Cpu {
	cpu.Pc = cpu.Pc + addr
	return cpu
}

func SetExit(cpu Cpu, exit uint32) Cpu {
	cpu.Exit = exit
	return cpu
}

func Execute(inst Instruction, cpu Cpu) (uint32, error) {
	insttype := GetInstructionType(inst)
	switch insttype {
	case LW:
		addr := cpu.Register[inst.Rs1] + inst.Imm_i
		return uint32(addr), nil
	default:
		return 0, nil
	}
}

func WriteBack(data int32, inst Instruction, cpu Cpu) Cpu {
	insttype := GetInstructionType(inst)
	switch insttype {
	case LW:
		cpu.Register[inst.Rd] = data
		return cpu
	default:
		return cpu
	}
}
