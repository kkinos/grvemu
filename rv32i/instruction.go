package rv32i

type Instruction struct {
	Opcode uint8
	Rs1    uint8
	Rs2    uint8
	Rd     uint8
	Func3  uint8
	Func7  uint8
	Imm_i  int32
	Imm_s  int32
	Imm_b  int32
	Imm_j  int32
	Imm_u  int32
}

type InstructionName int

const (
	LW InstructionName = iota
	SW
	ADD
	SUB
	ADDI
	AND
	OR
	XOR
	ANDI
	ORI
	XORI
	SLL
	SRL
	SRA
	SLLI
	SRLI
	SRAI
	SLT
	SLTU
	SLTI
	SLTIU
	BEQ
	BNE
	BLT
	BGE
	BLTU
	BGEU
	JAL
	JALR
	LUI
	AUIPC
	Unknown
)

func InstNameToString(instname InstructionName) string {
	switch instname {
	case LW:
		return "LW"
	case SW:
		return "SW"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case ADDI:
		return "ADDI"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	case ANDI:
		return "ANDI"
	case ORI:
		return "ORI"
	case XORI:
		return "XORI"
	case SLL:
		return "SLL"
	case SRL:
		return "SRL"
	case SRA:
		return "SRA"
	case SLLI:
		return "SLLI"
	case SRLI:
		return "SRL"
	case SRAI:
		return "SRAI"
	case SLT:
		return "SLT"
	case SLTU:
		return "SLTU"
	case SLTI:
		return "SLTI"
	case SLTIU:
		return "SLTIU"
	case BEQ:
		return "BEQ"
	case BNE:
		return "BNE"
	case BLT:
		return "BLT"
	case BGE:
		return "BGE"
	case BLTU:
		return "BLTU"
	case BGEU:
		return "BGEU"
	case JAL:
		return "JAL"
	case JALR:
		return "JALR"
	case LUI:
		return "LUI"
	case AUIPC:
		return "AUIPC"
	default:
		return "Unknown"
	}
}

func Decode(bits uint32) Instruction {
	var inst Instruction
	inst.Opcode = uint8(bits & 0x0000007F)
	inst.Rs1 = uint8((bits & 0x000F8000) >> 15)
	inst.Rs2 = uint8((bits & 0x01F00000) >> 20)
	inst.Rd = uint8((bits & 0x00000F80) >> 7)
	inst.Func3 = uint8((bits & 0x00007000) >> 12)
	inst.Func7 = uint8((bits & 0xFE000000) >> 25)
	inst.Imm_i = int32(bits&0xFFF00000) >> 20
	inst.Imm_s = (int32(bits&0x00000F80) >> 7) | (int32(bits&0xFE000000) >> 20)
	inst.Imm_b = (int32(bits&0x80000000) >> 19) | int32(bits&0x00000080<<4) | (int32((bits & 0x7E000000) >> 20)) | int32((bits&0x00000F00)>>7)
	inst.Imm_j = (int32(bits&0x80000000) >> 11) | int32(bits&0x000FF000) | int32((bits&0x00100000)>>9) | int32((bits&0x7FE00000)>>20)
	inst.Imm_u = int32(bits&0xFFFFF000) >> 12
	return inst
}

func GetInstructionName(inst Instruction) InstructionName {
	switch inst.Opcode {
	case 3:
		switch inst.Func3 {
		case 2:
			return LW
		default:
			return Unknown
		}
	case 19:
		switch inst.Func3 {
		case 0:
			return ADDI
		case 1:
			return SLLI
		case 2:
			return SLTI
		case 3:
			return SLTIU
		case 4:
			return XORI
		case 5:
			switch inst.Func7 {
			case 0:
				return SRLI
			case 32:
				return SRAI
			default:
				return Unknown
			}
		case 6:
			return ORI
		case 7:
			return ANDI
		default:
			return Unknown
		}
	case 23:
		return AUIPC
	case 35:
		switch inst.Func3 {
		case 2:
			return SW
		default:
			return Unknown
		}
	case 51:
		switch inst.Func3 {
		case 0:
			switch inst.Func7 {
			case 0:
				return ADD
			case 32:
				return SUB
			default:
				return Unknown
			}
		case 1:
			return SLL
		case 2:
			return SLT
		case 3:
			return SLTU
		case 5:
			switch inst.Func7 {
			case 0:
				return SRL
			case 32:
				return SRA
			default:
				return Unknown
			}
		case 4:
			return XOR
		case 6:
			return OR
		case 7:
			return AND
		default:
			return Unknown
		}
	case 55:
		return LUI
	case 99:
		switch inst.Func3 {
		case 0:
			return BEQ
		case 1:
			return BNE
		case 4:
			return BLT
		case 5:
			return BGE
		case 6:
			return BLTU
		case 7:
			return BGEU
		default:
			return Unknown
		}
	case 103:
		switch inst.Func3 {
		case 0:
			return JALR
		default:
			return Unknown
		}
	case 111:
		return JAL
	default:
		return Unknown
	}
}
