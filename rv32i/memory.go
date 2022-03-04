package rv32i

type Memory struct {
	Memory [1024 * 1024]uint8
}

func RoadMemory(binary []byte, memory Memory) Memory {
	copy(memory.Memory[:], []uint8(binary)[:])
	return memory
}

func MemoryAccess(addr uint32, inst Instruction, cpu Cpu, memory Memory) (uint32, Cpu, Memory) {
	insttype := GetInstructionType(inst)
	switch insttype {
	case LW:
		data := uint32(memory.Memory[addr]) | (uint32(memory.Memory[addr+1]) << 8) | (uint32(memory.Memory[addr+2]) << 16) | (uint32(memory.Memory[addr+3]) << 24)
		return data, cpu, memory
	case SW:
		memory.Memory[addr] = uint8((cpu.Register[inst.Rs2] & 0x000000FF))
		memory.Memory[addr+1] = uint8((cpu.Register[inst.Rs2] >> 8) & 0x000000FF)
		memory.Memory[addr+2] = uint8((cpu.Register[inst.Rs2] >> 16) & 0x000000FF)
		memory.Memory[addr+3] = uint8((cpu.Register[inst.Rs2] >> 24) & 0x000000FF)
		data := uint32(memory.Memory[addr]) | (uint32(memory.Memory[addr+1]) << 8) | (uint32(memory.Memory[addr+2]) << 16) | (uint32(memory.Memory[addr+3]) << 24)
		return data, cpu, memory
	default:
		return 0, cpu, memory
	}
}
