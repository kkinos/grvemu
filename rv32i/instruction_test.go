package rv32i

import (
	"math/rand"
	"testing"
	"time"
)

func RandIntRange(a int, b int) int {
	rand.Seed(time.Now().UnixNano())
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
		{"JALR", 103, 0, JALR},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
