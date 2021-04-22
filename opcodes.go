package main

const loadConstantOpcode = 1
const loadMemoryOpcode = 2
const storeConstantOpcode = 3
const storeMemoryOpcode = 4
const additionOpcode = 5
const subtractionOpcode = 6

func isOpcodeKnown(opcode int) bool {
	return opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode ||
		opcode == storeMemoryOpcode ||
		opcode == additionOpcode ||
		opcode == subtractionOpcode
}

func getOpcodeParameterCount(opcode int) int {
	if opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode ||
		opcode == storeMemoryOpcode {
		return 2
	}
	if opcode == additionOpcode || opcode == subtractionOpcode {
		return 3
	}

	return 0
}
