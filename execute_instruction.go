package main

func executeInstruction(computer machine, instructionInstance instruction) {
	if instructionInstance.kind == loadConstantOpcode {
		constant := instructionInstance.parameters[0]
		registerIndex := instructionInstance.parameters[1]

		computer.registers[registerIndex] = constant
	} else if instructionInstance.kind == loadMemoryOpcode {
		memoryIndex := instructionInstance.parameters[0]
		registerIndex := instructionInstance.parameters[1]

		computer.registers[registerIndex] = computer.memory[memoryIndex]
	} else if instructionInstance.kind == storeConstantOpcode {
		constant := instructionInstance.parameters[0]
		memoryIndex := instructionInstance.parameters[1]

		computer.memory[memoryIndex] = constant
	} else if instructionInstance.kind == storeMemoryOpcode {
		registerIndex := instructionInstance.parameters[0]
		memoryIndex := instructionInstance.parameters[1]

		computer.memory[memoryIndex] = computer.registers[registerIndex]
	}
}
