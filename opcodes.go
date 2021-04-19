package main

const loadConstantOpcode = 1
const loadMemoryOpcode = 2
const storeConstantOpcode = 3
const storeMemoryOpcode = 4

func isOpcodeKnown(opcode int) bool {
	return opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode ||
		opcode == storeMemoryOpcode
}

func getOpcodeParameterCount(opcode int) int {
	if opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode ||
		opcode == storeMemoryOpcode {
		return 2
	}

	return 0
}
