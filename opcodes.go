package machine

// ...
const (
	LoadConstantOpcode = iota + 1
	LoadMemoryOpcode
	StoreConstantOpcode
	StoreMemoryOpcode
	AdditionOpcode
	SubtractionOpcode
	MultiplicationOpcode
	DivisionOpcode
	ModuloOpcode
	JumpOpcode
)

// IsOpcodeKnown ...
func IsOpcodeKnown(opcode int) bool {
	return opcode == LoadConstantOpcode ||
		opcode == LoadMemoryOpcode ||
		opcode == StoreConstantOpcode ||
		opcode == StoreMemoryOpcode ||
		opcode == AdditionOpcode ||
		opcode == SubtractionOpcode ||
		opcode == MultiplicationOpcode ||
		opcode == DivisionOpcode ||
		opcode == ModuloOpcode ||
		opcode == JumpOpcode
}

// GetOpcodeParameterCount ...
func GetOpcodeParameterCount(opcode int) int {
	count := 0
	switch opcode {
	case JumpOpcode:
		count = 1
	case LoadConstantOpcode, LoadMemoryOpcode,
		StoreConstantOpcode, StoreMemoryOpcode:
		count = 2
	case AdditionOpcode, SubtractionOpcode,
		MultiplicationOpcode, DivisionOpcode, ModuloOpcode:
		count = 3
	}

	return count
}
