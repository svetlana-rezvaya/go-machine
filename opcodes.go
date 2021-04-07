package main

const loadConstantOpcode = 1
const loadMemoryOpcode = 2
const storeConstantOpcode = 3

func isOpcodeKnown(opcode int) bool {
	return opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode
}

func getOpcodeParameterCount(opcode int) int {
	if opcode == loadConstantOpcode ||
		opcode == loadMemoryOpcode ||
		opcode == storeConstantOpcode {
		return 2
	}

	return 0
}
