package rv32i

type Memory struct {
	Memory [1024 * 1024]uint8
}

func RoadMemory(binary []byte, memory Memory) Memory {
	copy(memory.Memory[:], []uint8(binary)[:])
	return memory
}

func MemoryAccess(res uint32, inst Instruction, cpu Cpu, memory Memory) (uint32, Cpu, Memory) {
	instname := GetInstructionName(inst)
	switch instname {
	case LW:
		data := uint32(memory.Memory[res]) | (uint32(memory.Memory[res+1]) << 8) | (uint32(memory.Memory[res+2]) << 16) | (uint32(memory.Memory[res+3]) << 24)
		return data, cpu, memory
	case SW:
		memory.Memory[res] = uint8((cpu.Register[inst.Rs2] & 0x000000FF))
		memory.Memory[res+1] = uint8((cpu.Register[inst.Rs2] >> 8) & 0x000000FF)
		memory.Memory[res+2] = uint8((cpu.Register[inst.Rs2] >> 16) & 0x000000FF)
		memory.Memory[res+3] = uint8((cpu.Register[inst.Rs2] >> 24) & 0x000000FF)
		data := uint32(memory.Memory[res]) | (uint32(memory.Memory[res+1]) << 8) | (uint32(memory.Memory[res+2]) << 16) | (uint32(memory.Memory[res+3]) << 24)
		return data, cpu, memory
	default:
		return res, cpu, memory
	}
}
