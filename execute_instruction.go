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
	} else if instructionInstance.kind == additionOpcode ||
		instructionInstance.kind == subtractionOpcode ||
		instructionInstance.kind == multiplicationOpcode ||
		instructionInstance.kind == divisionOpcode ||
		instructionInstance.kind == moduloOpcode {
		leftRegisterIndex := instructionInstance.parameters[0]
		leftOperand := computer.registers[leftRegisterIndex]

		rightRegisterIndex := instructionInstance.parameters[1]
		rightOperand := computer.registers[rightRegisterIndex]

		result := 0
		if instructionInstance.kind == additionOpcode {
			result = leftOperand + rightOperand
		} else if instructionInstance.kind == subtractionOpcode {
			result = leftOperand - rightOperand
		} else if instructionInstance.kind == multiplicationOpcode {
			result = leftOperand * rightOperand
		} else if instructionInstance.kind == divisionOpcode {
			result = leftOperand / rightOperand
		} else if instructionInstance.kind == moduloOpcode {
			result = leftOperand % rightOperand
		}

		resultRegisterIndex := instructionInstance.parameters[2]
		computer.registers[resultRegisterIndex] = result
	}
}
