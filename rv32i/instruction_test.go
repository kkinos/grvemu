package rv32i

import (
	"math/rand"
	"testing"
	"time"
)

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
			imm_s_11_5 := uint32(imm_s) & 0x00000FE0
			imm_s_4_0 := uint32(imm_s) & 0x0000001F
			bits = bits | imm_s_11_5<<20 | uint32(rs2<<20) | uint32(rs1<<15) | test.func3<<12 | imm_s_4_0<<7 | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_s != int(inst.Imm_s) {
				t.Errorf("Decode ImmS Error want %d got %d", imm_s, inst.Imm_s)
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

func TestDecodeBFormat(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		func3  uint32
		want   InstructionName
	}{
		{"BEQ", 99, 0, BEQ},
		{"BNE", 99, 1, BNE},
		{"BLT", 99, 4, BLT},
		{"BGE", 99, 5, BGE},
		{"BLTU", 99, 6, BLTU},
		{"BGEU", 99, 7, BGEU},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			imm_b := RandIntRange(-2048, 2047) << 1
			rs1 := rand.Intn(32)
			rs2 := rand.Intn(32)
			imm_b_12 := uint32(imm_b) & 0x00001000
			imm_b_10_5 := uint32(imm_b) & 0x000007E0
			imm_b_4_1 := uint32(imm_b) & 0x0000001E
			imm_b_11 := uint32(imm_b) & 0x00000800
			bits = bits | imm_b_12<<19 | imm_b_10_5<<20 | uint32(rs2<<20) | uint32(rs1<<15) | test.func3<<12 | imm_b_4_1<<7 | imm_b_11>>4 | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_b != int(inst.Imm_b) {
				t.Errorf("Decode ImmB Error want %d got %d", imm_b, inst.Imm_b)
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

func TestDecodeUFormat(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		want   InstructionName
	}{
		{"LUI", 55, LUI},
		{"AUIPC", 23, AUIPC},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			imm_u := RandIntRange(-524288, 524287)
			imm_u_31_12 := uint32(imm_u) & 0x000FFFFF
			rd := rand.Intn(32)

			bits = bits | imm_u_31_12<<12 | uint32(rd<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_u != int(inst.Imm_u) {
				t.Errorf("Decode ImmJ Error want %d got %d", imm_u, inst.Imm_u)
			}
			if rd != int(inst.Rd) {
				t.Errorf("Decode RD Error want %d got %d", rd, inst.Rd)
			}
		})
	}
}
func TestDecodeJFormat(t *testing.T) {
	tests := []struct {
		name   string
		opcode uint32
		want   InstructionName
	}{
		{"JAL", 111, JAL},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			var bits uint32
			imm_j := RandIntRange(-524288, 524287) << 1
			rd := rand.Intn(32)
			imm_j_20 := uint32(imm_j) & 0x00100000
			imm_j_10_1 := uint32(imm_j) & 0x000007FE
			imm_j_11 := uint32(imm_j) & 0x00000800
			imm_j_19_12 := uint32(imm_j) & 0x000FF000

			bits = bits | imm_j_20<<11 | imm_j_10_1<<20 | imm_j_11<<9 | imm_j_19_12 | uint32(rd<<7) | test.opcode
			inst := Decode(bits)
			if test.want != GetInstructionName(inst) {
				t.Errorf("Decode Instrcution Name Error want %s got %s", InstNameToString(test.want), InstNameToString(GetInstructionName(inst)))
			}
			if imm_j != int(inst.Imm_j) {
				t.Errorf("Decode ImmJ Error want %d got %d", imm_j, inst.Imm_j)
			}
			if rd != int(inst.Rd) {
				t.Errorf("Decode RD Error want %d got %d", rd, inst.Rd)
			}
		})
	}
}
