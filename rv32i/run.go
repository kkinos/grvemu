package rv32i

import "fmt"

func Run(binary []byte, debug bool) error {
	var cpu Cpu
	cpu = SetExit(cpu, 0x00602823)

	var memory Memory
	memory = RoadMemory(binary, memory)

	if err := Loop(cpu, memory, debug); err != nil {
		return err
	}

	return nil
}

func Loop(cpu Cpu, memory Memory, debug bool) error {
	var bits uint32
	for {
		// IF Stage
		bits = uint32(memory.Memory[cpu.Pc]) | uint32(memory.Memory[cpu.Pc+1])<<8 | uint32(memory.Memory[cpu.Pc+2])<<16 | uint32(memory.Memory[cpu.Pc+3])<<24

		// ID Stage
		inst := Decode(bits)

		// EX Stage
		addr, err := Execute(inst, cpu)
		if err != nil {
			return err
		}

		// MEM Stage
		var data int32
		data, cpu, memory = MemoryAccess(addr, inst, cpu, memory)

		// WB Stage
		cpu = WriteBack(data, inst, cpu)

		if debug {
			fmt.Printf("pc   	  : 0x%x\n", cpu.Pc)
			fmt.Printf("bits 	  : 0x%x\n", bits)
			fmt.Printf("inst 	  : %s\n", InstTypeToString(GetInstructionType(inst)))
			fmt.Printf("rs1  	  : %d\n", inst.Rs1)
			fmt.Printf("rs2  	  : %d\n", inst.Rs2)
			fmt.Printf("rd  	  : %d\n", inst.Rd)
			fmt.Printf("rs1_data  : 0x%x\n", cpu.Register[inst.Rs1])
			fmt.Printf("rs2_data  : 0x%x\n", cpu.Register[inst.Rs2])
			fmt.Printf("tgt_data  : 0x%x\n", data)
			fmt.Printf("dmem_addr : %d\n", addr)
			fmt.Print("------\n")
		}
		if bits == cpu.Exit {
			break
		}

		cpu = MovePc(cpu, 4)
	}
	return nil
}
