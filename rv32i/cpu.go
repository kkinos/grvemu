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
	case SW:
		addr := cpu.Register[inst.Rs1] + inst.Imm_s
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
	case ADD:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] + cpu.Register[inst.Rs2]
		return cpu
	case SUB:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] - cpu.Register[inst.Rs2]
		return cpu
	case ADDI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] + inst.Imm_i
		return cpu
	case AND:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] & cpu.Register[inst.Rs2]
		return cpu
	case OR:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] | cpu.Register[inst.Rs2]
		return cpu
	case XOR:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] ^ cpu.Register[inst.Rs2]
		return cpu
	case ANDI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] & inst.Imm_i
		return cpu
	case ORI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] | inst.Imm_i
		return cpu
	case XORI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] ^ inst.Imm_i
		return cpu
	default:
		return cpu
	}
}
