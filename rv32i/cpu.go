package rv32i

type Cpu struct {
	Register [32]uint32
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
		addr := cpu.Register[inst.Rs1] + uint32(inst.Imm_i)
		return uint32(addr), nil
	case SW:
		addr := cpu.Register[inst.Rs1] + uint32(inst.Imm_s)
		return uint32(addr), nil
	default:
		return 0, nil
	}
}

func WriteBack(data uint32, inst Instruction, cpu Cpu) Cpu {
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
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] + uint32(inst.Imm_i)
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
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] & uint32(inst.Imm_i)
		return cpu
	case ORI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] | uint32(inst.Imm_i)
		return cpu
	case XORI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] ^ uint32(inst.Imm_i)
		return cpu
	case SLL:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] << (cpu.Register[inst.Rs2] & 0x1F)
		return cpu
	case SRL:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] >> (cpu.Register[inst.Rs2] & 0x1F)
		return cpu
	case SRA:
		cpu.Register[inst.Rd] = uint32(int32(cpu.Register[inst.Rs1]) >> (cpu.Register[inst.Rs2] & 0x1F))
		return cpu
	case SLLI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] << (inst.Imm_i & 0x1F)
		return cpu
	case SRLI:
		cpu.Register[inst.Rd] = cpu.Register[inst.Rs1] >> (inst.Imm_i & 0x1F)
		return cpu
	case SRAI:
		cpu.Register[inst.Rd] = uint32(int32(cpu.Register[inst.Rs1]) >> (inst.Imm_i & 0x1F))
		return cpu
	default:
		return cpu
	}
}
