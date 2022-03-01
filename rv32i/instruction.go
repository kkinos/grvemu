package rv32i

type Instruction struct {
	Opcode uint8
	Rs1    uint8
	Rs2    uint8
	Rd     uint8
	Func3  uint8
	Imm_i  int32
	Imm_s  int32
}

type InstructionType int

const (
	LW InstructionType = iota
	SW
	Unknown
)

func Decode(bits uint32) Instruction {
	var inst Instruction
	inst.Opcode = uint8(bits & 0x0000007F)
	inst.Rs1 = uint8((bits & 0x000F8000) >> 15)
	inst.Rs2 = uint8((bits & 0x01F00000) >> 20)
	inst.Rd = uint8((bits & 0x00000F80) >> 7)
	inst.Func3 = uint8((bits & 0x00007000) >> 12)
	inst.Imm_i = int32((bits & 0xFFF00000) >> 20)
	inst.Imm_s = int32((bits&0x00000F80)>>7 | (bits&0xFE000000)>>20)
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
	case 35:
		switch inst.Func3 {
		case 2:
			return SW
		default:
			return Unknown
		}
	default:
		return Unknown
	}
}