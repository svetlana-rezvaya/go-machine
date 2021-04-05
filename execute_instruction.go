package main

func executeInstruction(computer machine, instructionInstance instruction) {
	if instructionInstance.kind == loadConstantOpcode {
		constant := instructionInstance.parameters[0]
		registerIndex := instructionInstance.parameters[1]

		computer.registers[registerIndex] = constant
	} else if instructionInstance.kind == loadMemoryOpcode {

	}
}
