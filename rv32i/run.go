package rv32i

import "fmt"

func Run(binary []byte, end uint32, debug bool, test bool) error {
	var cpu Cpu
	cpu = SetExit(cpu, end)
	cpu.Register[2] = MEMORY_SIZE // sp
	var memory Memory
	memory = RoadMemory(binary, memory)

	cpu, memory, err := Loop(cpu, memory, debug, test)
	if err != nil {
		return err
	}

	return nil
}

func Loop(cpu Cpu, memory Memory, debug bool, test bool) (Cpu, Memory, error) {
	var bits uint32
	var pc uint32
	for {
		// IF Stage
		bits = uint32(memory.Memory[cpu.Pc]) | (uint32(memory.Memory[cpu.Pc+1]) << 8) | (uint32(memory.Memory[cpu.Pc+2]) << 16) | (uint32(memory.Memory[cpu.Pc+3]) << 24)
		pc = cpu.Pc
		// ID Stage
		inst := Decode(bits)
		if debug {
			fmt.Print("------\n")
			fmt.Printf("pc   	  : 0x%x\n", pc)
			fmt.Printf("bits 	  : 0x%x\n", bits)
			fmt.Printf("inst 	  : %s\n", InstNameToString(GetInstructionName(inst)))
			fmt.Printf("rs1  	  : %d\n", inst.Rs1)
			fmt.Printf("rs2  	  : %d\n", inst.Rs2)
			fmt.Printf("rd  	  : %d\n", inst.Rd)
			fmt.Printf("rs1_data  : 0x%x\n", cpu.Register[inst.Rs1])
			fmt.Printf("rs2_data  : 0x%x\n", cpu.Register[inst.Rs2])
		}

		// EX Stage
		pcchanged, res, err := Execute(inst, cpu)
		if err != nil && !debug {
			return cpu, memory, err
		}
		if pcchanged {
			cpu = MovePc(cpu, res)
		} else {
			cpu = AddPc(cpu, 4)
		}

		// MEM Stage
		var data uint32
		data, cpu, memory = MemoryAccess(res, inst, cpu, memory)

		// WB Stage
		if pcchanged {
			cpu = WriteBack(pc+4, inst, cpu)
		} else {
			cpu = WriteBack(data, inst, cpu)
		}

		if debug {
			fmt.Printf("wb_data   : 0x%x\n", cpu.Register[inst.Rd])
		}
		if test {
			fmt.Printf("gp  	  : %d\n", cpu.Register[3])
		}

		if bits == cpu.Exit {
			break
		}

	}
	return cpu, memory, nil
}
