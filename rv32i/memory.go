package rv32i

type Memory struct {
	Memory [1024 * 1024]uint8
}

func RoadMemory(binary []byte, memory Memory) Memory {
	copy(memory.Memory[:], []uint8(binary)[:])
	return memory
}

func MemoryAccess(addr uint32, inst Instruction, cpu Cpu, memory Memory) (int32, Cpu, Memory) {
	insttype := GetInstructionType(inst)
	switch insttype {
	case LW:
		data := int32(uint32(memory.Memory[addr]) | uint32(memory.Memory[addr+1])<<8 | uint32(memory.Memory[addr+2])<<16 | uint32(memory.Memory[addr+3])<<24)
		return data, cpu, memory
	default:
		return 0, cpu, memory
	}
}
