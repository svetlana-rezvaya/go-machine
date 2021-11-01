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
		opcode == ModuloOpcode
}

// GetOpcodeParameterCount ...
func GetOpcodeParameterCount(opcode int) int {
	switch opcode {
	case LoadConstantOpcode, LoadMemoryOpcode,
		StoreConstantOpcode, StoreMemoryOpcode:
		return 2
	case AdditionOpcode, SubtractionOpcode,
		MultiplicationOpcode, DivisionOpcode, ModuloOpcode:
		return 3
	default:
		return 0
	}
}
