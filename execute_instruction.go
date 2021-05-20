package main

func executeInstruction(computer machine, instructionInstance instruction) {
	switch instructionInstance.kind {
	case loadConstantOpcode:
		constant := instructionInstance.parameters[0]
		registerIndex := instructionInstance.parameters[1]

		computer.registers[registerIndex] = constant

	case loadMemoryOpcode:
		memoryIndex := instructionInstance.parameters[0]
		registerIndex := instructionInstance.parameters[1]

		computer.registers[registerIndex] = computer.memory[memoryIndex]

	case storeConstantOpcode:
		constant := instructionInstance.parameters[0]
		memoryIndex := instructionInstance.parameters[1]

		computer.memory[memoryIndex] = constant

	case storeMemoryOpcode:
		registerIndex := instructionInstance.parameters[0]
		memoryIndex := instructionInstance.parameters[1]

		computer.memory[memoryIndex] = computer.registers[registerIndex]

	case additionOpcode, subtractionOpcode,
		multiplicationOpcode, divisionOpcode, moduloOpcode:
		leftRegisterIndex := instructionInstance.parameters[0]
		leftOperand := computer.registers[leftRegisterIndex]

		rightRegisterIndex := instructionInstance.parameters[1]
		rightOperand := computer.registers[rightRegisterIndex]

		result := 0
		switch instructionInstance.kind {
		case additionOpcode:
			result = leftOperand + rightOperand
		case subtractionOpcode:
			result = leftOperand - rightOperand
		case multiplicationOpcode:
			result = leftOperand * rightOperand
		case divisionOpcode:
			result = leftOperand / rightOperand
		case moduloOpcode:
			result = leftOperand % rightOperand
		}

		resultRegisterIndex := instructionInstance.parameters[2]
		computer.registers[resultRegisterIndex] = result
	}
}
