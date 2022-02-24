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
