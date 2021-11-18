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
	JumpIfNegativeOpcode
	JumpIfZeroOpcode
	JumpIfPositiveOpcode
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
		opcode == JumpOpcode ||
		opcode == JumpIfNegativeOpcode ||
		opcode == JumpIfZeroOpcode ||
		opcode == JumpIfPositiveOpcode
}

// GetOpcodeParameterCount ...
func GetOpcodeParameterCount(opcode int) int {
	count := 0
	switch opcode {
	case JumpOpcode:
		count = 1
	case LoadConstantOpcode, LoadMemoryOpcode,
		StoreConstantOpcode, StoreMemoryOpcode,
		JumpIfNegativeOpcode, JumpIfZeroOpcode, JumpIfPositiveOpcode:
		count = 2
	case AdditionOpcode, SubtractionOpcode,
		MultiplicationOpcode, DivisionOpcode, ModuloOpcode:
		count = 3
	}

	return count
}
