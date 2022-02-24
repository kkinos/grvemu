package rv32i

type Cpu struct {
	Register [32]uint32
	Pc       uint32
	Exit     uint32
}

type Instruction struct {
	Rs1 uint8
	Rs2 uint8
	Rd  uint8
}

func MovePc(cpu Cpu, addr uint32) Cpu {
	cpu.Pc = cpu.Pc + addr
	return cpu
}

func SetExit(cpu Cpu, exit uint32) Cpu {
	cpu.Exit = exit
	return cpu
}

func Decode(bits uint32) Instruction {
	var inst Instruction
	inst.Rs1 = uint8((bits & 0x000F8000) >> 15)
	inst.Rs2 = uint8((bits & 0x01F00000) >> 20)
	inst.Rd = uint8((bits & 0x00000F80) >> 7)

	return inst
}
