package main

const loadConstantOpcode = 1
const loadMemoryOpcode = 2

func isOpcodeKnown(opcode int) bool {
	return opcode == loadConstantOpcode || opcode == loadMemoryOpcode
}

func getOpcodeParameterCount(opcode int) int {
	if opcode == loadConstantOpcode || opcode == loadMemoryOpcode {
		return 2
	}

	return 0
}
