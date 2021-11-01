package machine

const LoadConstantOpcode = 1
const LoadMemoryOpcode = 2
const StoreConstantOpcode = 3
const StoreMemoryOpcode = 4
const AdditionOpcode = 5
const SubtractionOpcode = 6
const MultiplicationOpcode = 7
const DivisionOpcode = 8
const ModuloOpcode = 9

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
