package rv32i

import (
	"math/rand"
	"testing"
	"time"
)

func RandIntRange(a int, b int) int {
	return a + rand.Intn(b-a+1)
}
func TestDecodeIFormat1(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		func3  uint32
		want   InstructionName
	}{
		{"LW", 3, 2, LW},
		{"ADDI", 19, 0, ADDI},
		{"ANDI", 19, 7, ANDI},
		{"ORI", 19, 6, ORI},
		{"XORI", 19, 4, XORI},
		{"SLTI", 19, 2, SLTI},
		{"SLTIU", 19, 3, SLTIU},
		{"JALR", 103, 0, JALR},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			imm_i := RandIntRange(-2048, 2047)
			rs1 := rand.Intn(32)
			rd := rand.Intn(32)
			bits = bits | uint32(imm_i<<20) | uint32(rs1<<15) | test.func3<<12 | uint32(rd<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_i != int(inst.Imm_i) {
				t.Errorf("Decode ImmI Error want %d got %d", imm_i, inst.Imm_i)
			}
			if rs1 != int(inst.Rs1) {
				t.Errorf("Decode RS1 Error want %d got %d", rs1, inst.Rs1)
			}
			if rd != int(inst.Rd) {
				t.Errorf("Decode RD Error want %d got %d", rd, inst.Rd)
			}
		})
	}
}
func TestDecodeIFormat2(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		func3  uint32
		func7  uint32
		want   InstructionName
	}{
		{"SLLI", 19, 1, 0, SLLI},
		{"SRLI", 19, 5, 0, SRLI},
		{"SRAI", 19, 5, 32, SRAI},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			rs1 := rand.Intn(32)
			shamt := rand.Intn(32)
			rd := rand.Intn(32)
			bits = bits | test.func7<<25 | uint32(shamt<<20) | uint32(rs1<<15) | test.func3<<12 | uint32(rd<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if rs1 != int(inst.Rs1) {
				t.Errorf("Decode RS1 Error want %d got %d", rs1, inst.Rs1)
			}
			if shamt != int(inst.Imm_i&0x1F) {
				t.Errorf("Decode RS2 Error want %d got %d", shamt, inst.Imm_i&0x1F)
			}
			if rd != int(inst.Rd) {
				t.Errorf("Decode RD Error want %d got %d", rd, inst.Rd)
			}
		})
	}
}

func TestDecodeSFormat(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		func3  uint32
		want   InstructionName
	}{
		{"SW", 35, 2, SW},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			imm_s := RandIntRange(-2048, 2047)
			rs1 := rand.Intn(32)
			rs2 := rand.Intn(32)
			imm_s_1 := uint32(imm_s) & 0x00000FE0
			imm_s_2 := uint32(imm_s) & 0x0000001F
			bits = bits | uint32(imm_s_1<<20) | uint32(rs2<<20) | uint32(rs1<<15) | test.func3<<12 | uint32(imm_s_2<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_s != int(inst.Imm_s) {
				t.Errorf("Decode ImmI Error want %d got %d", imm_s, inst.Imm_i)
			}
			if rs1 != int(inst.Rs1) {
				t.Errorf("Decode RS1 Error want %d got %d", rs1, inst.Rs1)
			}
			if rs2 != int(inst.Rs2) {
				t.Errorf("Decode RS2 Error want %d got %d", rs2, inst.Rs2)
			}
		})
	}
}

func TestDecodeRFormat(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		func3  uint32
		func7  uint32
		want   InstructionName
	}{
		{"ADD", 51, 0, 0, ADD},
		{"SUB", 51, 0, 32, SUB},
		{"AND", 51, 7, 0, AND},
		{"OR", 51, 6, 0, OR},
		{"XOR", 51, 4, 0, XOR},
		{"SLL", 51, 1, 0, SLL},
		{"SRL", 51, 5, 0, SRL},
		{"SRA", 51, 5, 32, SRA},
		{"SLT", 51, 2, 0, SLT},
		{"SLTU", 51, 3, 0, SLTU},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			rs1 := rand.Intn(32)
			rs2 := rand.Intn(32)
			rd := rand.Intn(32)
			bits = bits | test.func7<<25 | uint32(rs2<<20) | uint32(rs1<<15) | test.func3<<12 | uint32(rd<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if rs1 != int(inst.Rs1) {
				t.Errorf("Decode RS1 Error want %d got %d", rs1, inst.Rs1)
			}
			if rs2 != int(inst.Rs2) {
				t.Errorf("Decode RS2 Error want %d got %d", rs2, inst.Rs2)
			}
			if rd != int(inst.Rd) {
				t.Errorf("Decode RD Error want %d got %d", rd, inst.Rd)
			}
		})
	}
}
