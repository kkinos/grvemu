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
}

type InstructionType int

const (
	LW InstructionType = iota
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
	Unknown
)

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
	return inst
}

func GetInstructionType(inst Instruction) InstructionType {
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
	default:
		return Unknown
	}
}
