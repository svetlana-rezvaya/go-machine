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
	} else if instructionInstance.kind == additionOpcode {
		leftRegisterIndex := instructionInstance.parameters[0]
		rightRegisterIndex := instructionInstance.parameters[1]
		resultRegisterIndex := instructionInstance.parameters[2]

		computer.registers[resultRegisterIndex] =
			computer.registers[leftRegisterIndex] +
				computer.registers[rightRegisterIndex]
	} else if instructionInstance.kind == subtractionOpcode {
		leftRegisterIndex := instructionInstance.parameters[0]
		rightRegisterIndex := instructionInstance.parameters[1]
		resultRegisterIndex := instructionInstance.parameters[2]

		computer.registers[resultRegisterIndex] =
			computer.registers[leftRegisterIndex] -
				computer.registers[rightRegisterIndex]
	} else if instructionInstance.kind == multiplicationOpcode {
		leftRegisterIndex := instructionInstance.parameters[0]
		rightRegisterIndex := instructionInstance.parameters[1]
		resultRegisterIndex := instructionInstance.parameters[2]

		computer.registers[resultRegisterIndex] =
			computer.registers[leftRegisterIndex] *
				computer.registers[rightRegisterIndex]
	} else if instructionInstance.kind == divisionOpcode {
		leftRegisterIndex := instructionInstance.parameters[0]
		rightRegisterIndex := instructionInstance.parameters[1]
		resultRegisterIndex := instructionInstance.parameters[2]

		computer.registers[resultRegisterIndex] =
			computer.registers[leftRegisterIndex] /
				computer.registers[rightRegisterIndex]
	}
}
