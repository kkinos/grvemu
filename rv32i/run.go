package rv32i

import "fmt"

func Loop(cpu Cpu, memory Memory, debug bool) error {
	var inst uint32
	for {
		inst = uint32(memory.Memory[cpu.Pc]) | uint32(memory.Memory[cpu.Pc+1])<<8 | uint32(memory.Memory[cpu.Pc+2])<<16 | uint32(memory.Memory[cpu.Pc+3])<<24

		if debug {
			fmt.Printf("pc   : 0x%x\n", cpu.Pc)
			fmt.Printf("inst : 0x%x\n", inst)
			fmt.Print("------\n")
		}
		if inst == cpu.Exit {
			break
		}

		cpu = MovePc(cpu, 4)
	}
	return nil
}

func Run(binary []byte, debug bool) error {
	var cpu Cpu
	cpu = SetExit(cpu, 0x34333231)

	var memory Memory
	memory = RoadMemory(binary, memory)

	if err := Loop(cpu, memory, debug); err != nil {
		return err
	}

	return nil
}
