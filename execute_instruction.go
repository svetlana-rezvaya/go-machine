package machine

// ExecuteInstruction ...
func ExecuteInstruction(computer Machine, instructionInstance Instruction) {
	switch instructionInstance.Kind {
	case LoadConstantOpcode:
		constant := instructionInstance.Parameters[0]
		registerIndex := instructionInstance.Parameters[1]

		computer.Registers[registerIndex] = constant

	case LoadMemoryOpcode:
		memoryIndex := instructionInstance.Parameters[0]
		registerIndex := instructionInstance.Parameters[1]

		computer.Registers[registerIndex] = computer.Memory[memoryIndex]

	case StoreConstantOpcode:
		constant := instructionInstance.Parameters[0]
		memoryIndex := instructionInstance.Parameters[1]

		computer.Memory[memoryIndex] = constant

	case StoreMemoryOpcode:
		registerIndex := instructionInstance.Parameters[0]
		memoryIndex := instructionInstance.Parameters[1]

		computer.Memory[memoryIndex] = computer.Registers[registerIndex]

	case AdditionOpcode, SubtractionOpcode,
		MultiplicationOpcode, DivisionOpcode, ModuloOpcode:
		leftRegisterIndex := instructionInstance.Parameters[0]
		leftOperand := computer.Registers[leftRegisterIndex]

		rightRegisterIndex := instructionInstance.Parameters[1]
		rightOperand := computer.Registers[rightRegisterIndex]

		result := 0
		switch instructionInstance.Kind {
		case AdditionOpcode:
			result = leftOperand + rightOperand
		case SubtractionOpcode:
			result = leftOperand - rightOperand
		case MultiplicationOpcode:
			result = leftOperand * rightOperand
		case DivisionOpcode:
			result = leftOperand / rightOperand
		case ModuloOpcode:
			result = leftOperand % rightOperand
		}

		resultRegisterIndex := instructionInstance.Parameters[2]
		computer.Registers[resultRegisterIndex] = result
	}
}
