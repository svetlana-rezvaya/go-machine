package machine

// Instruction ...
type Instruction struct {
	Opcode     int
	Parameters []int
}

// Len ...
func (instructionInstance Instruction) Len() int {
	return GetOpcodeParameterCount(instructionInstance.Opcode) + 1
}
