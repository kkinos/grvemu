package rv32i

import "errors"

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
		return addr, nil
	case SW:
		addr := cpu.Register[inst.Rs1] + uint32(inst.Imm_s)
		return addr, nil
	case ADD:
		res := cpu.Register[inst.Rs1] + cpu.Register[inst.Rs2]
		return res, nil
	case SUB:
		res := cpu.Register[inst.Rs1] - cpu.Register[inst.Rs2]
		return res, nil
	case ADDI:
		res := cpu.Register[inst.Rs1] + uint32(inst.Imm_i)
		return res, nil
	case AND:
		res := cpu.Register[inst.Rs1] & cpu.Register[inst.Rs2]
		return res, nil
	case OR:
		res := cpu.Register[inst.Rs1] | cpu.Register[inst.Rs2]
		return res, nil
	case XOR:
		res := cpu.Register[inst.Rs1] ^ cpu.Register[inst.Rs2]
		return res, nil
	case ANDI:
		res := cpu.Register[inst.Rs1] & uint32(inst.Imm_i)
		return res, nil
	case ORI:
		res := cpu.Register[inst.Rs1] | uint32(inst.Imm_i)
		return res, nil
	case XORI:
		res := cpu.Register[inst.Rs1] ^ uint32(inst.Imm_i)
		return res, nil
	case SLL:
		res := cpu.Register[inst.Rs1] << (cpu.Register[inst.Rs2] & 0x1F)
		return res, nil
	case SRL:
		res := cpu.Register[inst.Rs1] >> (cpu.Register[inst.Rs2] & 0x1F)
		return res, nil
	case SRA:
		res := uint32(int32(cpu.Register[inst.Rs1]) >> (cpu.Register[inst.Rs2] & 0x1F))
		return res, nil
	case SLLI:
		res := cpu.Register[inst.Rs1] << (inst.Imm_i & 0x1F)
		return res, nil
	case SRLI:
		res := cpu.Register[inst.Rs1] >> (inst.Imm_i & 0x1F)
		return res, nil
	case SRAI:
		res := uint32(int32(cpu.Register[inst.Rs1]) >> (inst.Imm_i & 0x1F))
		return res, nil
	case SLT:
		if int32(cpu.Register[inst.Rs1]) < int32(cpu.Register[inst.Rs2]) {
			return 1, nil
		} else {
			return 0, nil
		}
	case SLTU:
		if cpu.Register[inst.Rs1] < cpu.Register[inst.Rs2] {
			return 1, nil
		} else {
			return 0, nil
		}
	case SLTI:
		if int32(cpu.Register[inst.Rs1]) < inst.Imm_i {
			return 1, nil
		} else {
			return 0, nil
		}
	case SLTIU:
		if cpu.Register[inst.Rs1] < uint32(inst.Imm_i) {
			return 1, nil
		} else {
			return 0, nil
		}

	default:
		return 0, errors.New("unknown instruction")
	}
}

func WriteBack(data uint32, inst Instruction, cpu Cpu) Cpu {
	insttype := GetInstructionType(inst)
	switch insttype {
	case SW:
		return cpu
	default:
		cpu.Register[inst.Rd] = data
		return cpu
	}
}
