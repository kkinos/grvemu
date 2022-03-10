package rv32i

import "errors"

type Cpu struct {
	Register [32]uint32
	Pc       uint32
	Exit     uint32
}

func AddPc(cpu Cpu, addr uint32) Cpu {
	cpu.Pc = cpu.Pc + addr
	return cpu
}

func MovePc(cpu Cpu, addr uint32) Cpu {
	cpu.Pc = addr
	return cpu
}

func SetExit(cpu Cpu, exit uint32) Cpu {
	cpu.Exit = exit
	return cpu
}

func Execute(inst Instruction, cpu Cpu) (bool, uint32, error) {
	insttype := GetInstructionType(inst)
	switch insttype {
	case LW:
		addr := cpu.Register[inst.Rs1] + uint32(inst.Imm_i)
		return false, addr, nil
	case SW:
		addr := cpu.Register[inst.Rs1] + uint32(inst.Imm_s)
		return false, addr, nil
	case ADD:
		res := cpu.Register[inst.Rs1] + cpu.Register[inst.Rs2]
		return false, res, nil
	case SUB:
		res := cpu.Register[inst.Rs1] - cpu.Register[inst.Rs2]
		return false, res, nil
	case ADDI:
		res := cpu.Register[inst.Rs1] + uint32(inst.Imm_i)
		return false, res, nil
	case AND:
		res := cpu.Register[inst.Rs1] & cpu.Register[inst.Rs2]
		return false, res, nil
	case OR:
		res := cpu.Register[inst.Rs1] | cpu.Register[inst.Rs2]
		return false, res, nil
	case XOR:
		res := cpu.Register[inst.Rs1] ^ cpu.Register[inst.Rs2]
		return false, res, nil
	case ANDI:
		res := cpu.Register[inst.Rs1] & uint32(inst.Imm_i)
		return false, res, nil
	case ORI:
		res := cpu.Register[inst.Rs1] | uint32(inst.Imm_i)
		return false, res, nil
	case XORI:
		res := cpu.Register[inst.Rs1] ^ uint32(inst.Imm_i)
		return false, res, nil
	case SLL:
		res := cpu.Register[inst.Rs1] << (cpu.Register[inst.Rs2] & 0x1F)
		return false, res, nil
	case SRL:
		res := cpu.Register[inst.Rs1] >> (cpu.Register[inst.Rs2] & 0x1F)
		return false, res, nil
	case SRA:
		res := uint32(int32(cpu.Register[inst.Rs1]) >> (cpu.Register[inst.Rs2] & 0x1F))
		return false, res, nil
	case SLLI:
		res := cpu.Register[inst.Rs1] << (inst.Imm_i & 0x1F)
		return false, res, nil
	case SRLI:
		res := cpu.Register[inst.Rs1] >> (inst.Imm_i & 0x1F)
		return false, res, nil
	case SRAI:
		res := uint32(int32(cpu.Register[inst.Rs1]) >> (inst.Imm_i & 0x1F))
		return false, res, nil
	case SLT:
		if int32(cpu.Register[inst.Rs1]) < int32(cpu.Register[inst.Rs2]) {
			return false, 1, nil
		} else {
			return false, 0, nil
		}
	case SLTU:
		if cpu.Register[inst.Rs1] < cpu.Register[inst.Rs2] {
			return false, 1, nil
		} else {
			return false, 0, nil
		}
	case SLTI:
		if int32(cpu.Register[inst.Rs1]) < inst.Imm_i {
			return false, 1, nil
		} else {
			return false, 0, nil
		}
	case SLTIU:
		if cpu.Register[inst.Rs1] < uint32(inst.Imm_i) {
			return false, 1, nil
		} else {
			return false, 0, nil
		}
	case BEQ:
		if cpu.Register[inst.Rs1] == cpu.Register[inst.Rs2] {
			return true, cpu.Pc + uint32(inst.Imm_b), nil
		} else {
			return false, 0, nil
		}
	case BNE:
		if cpu.Register[inst.Rs1] != cpu.Register[inst.Rs2] {
			res := cpu.Pc + uint32(inst.Imm_b)
			return true, res, nil
		} else {
			return false, 0, nil
		}
	case BLT:
		if int32(cpu.Register[inst.Rs1]) < int32(cpu.Register[inst.Rs2]) {
			res := cpu.Pc + uint32(inst.Imm_b)
			return true, res, nil
		} else {
			return false, 0, nil
		}
	case BGE:
		if int32(cpu.Register[inst.Rs1]) >= int32(cpu.Register[inst.Rs2]) {
			res := cpu.Pc + uint32(inst.Imm_b)
			return true, res, nil
		} else {
			return false, 0, nil
		}
	case BLTU:
		if cpu.Register[inst.Rs1] < cpu.Register[inst.Rs2] {
			res := cpu.Pc + uint32(inst.Imm_b)
			return true, res, nil
		} else {
			return false, 0, nil
		}
	case BGEU:
		if cpu.Register[inst.Rs1] >= cpu.Register[inst.Rs2] {
			res := cpu.Pc + uint32(inst.Imm_b)
			return true, res, nil
		} else {
			return false, 0, nil
		}
	case JAL:
		res := cpu.Pc + uint32(inst.Imm_j)
		return true, res, nil
	case JALR:
		res := (cpu.Register[inst.Rs1] + uint32(inst.Imm_j)) & ^uint32(1)
		return true, res, nil

	default:
		return false, 0, errors.New("unknown instruction")
	}
}

func WriteBack(data uint32, inst Instruction, cpu Cpu) Cpu {
	insttype := GetInstructionType(inst)
	switch insttype {
	case SW, BEQ, BNE, BLT, BGE, BLTU, BGEU:
		return cpu
	default:
		cpu.Register[inst.Rd] = data
		return cpu
	}
}
