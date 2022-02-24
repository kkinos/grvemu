package rv32i

type Memory struct {
	Memory [1024 * 1024]uint8
}

func RoadMemory(binary []byte, memory Memory) Memory {
	copy(memory.Memory[:], []uint8(binary)[:])
	return memory
}
