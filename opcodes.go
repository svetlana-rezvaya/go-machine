package machine

const loadConstantOpcode = 1
const loadMemoryOpcode = 2
const storeConstantOpcode = 3
const storeMemoryOpcode = 4
const additionOpcode = 5
const subtractionOpcode = 6
const multiplicationOpcode = 7
const divisionOpcode = 8
const moduloOpcode = 9

func isOpcodeKnown(opcode int) bool {
	return opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode ||
		opcode == storeMemoryOpcode ||
		opcode == additionOpcode ||
		opcode == subtractionOpcode ||
		opcode == multiplicationOpcode ||
		opcode == divisionOpcode ||
		opcode == moduloOpcode
}

func getOpcodeParameterCount(opcode int) int {
	switch opcode {
	case loadConstantOpcode, loadMemoryOpcode,
		storeConstantOpcode, storeMemoryOpcode:
		return 2
	case additionOpcode, subtractionOpcode,
		multiplicationOpcode, divisionOpcode, moduloOpcode:
		return 3
	default:
		return 0
	}
}
