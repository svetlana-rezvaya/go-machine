package machine

// ExecuteInstruction ...
func ExecuteInstruction(computer Machine, instructionInstance Instruction) {
	switch instructionInstance.Opcode {
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
		switch instructionInstance.Opcode {
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

	case JumpOpcode:
		ExecuteJumpInstruction(computer, instructionInstance, 0)

	case JumpIfNegativeOpcode, JumpIfZeroOpcode, JumpIfPositiveOpcode:
		registerIndex := instructionInstance.Parameters[0]
		operand := computer.Registers[registerIndex]

		isSuccessful := false
		switch instructionInstance.Opcode {
		case JumpIfNegativeOpcode:
			isSuccessful = operand < 0
		case JumpIfZeroOpcode:
			isSuccessful = operand == 0
		case JumpIfPositiveOpcode:
			isSuccessful = operand > 0
		}

		if isSuccessful {
			ExecuteJumpInstruction(computer, instructionInstance, 1)
		}
	}
}

// ExecuteJumpInstruction ...
func ExecuteJumpInstruction(
	computer Machine,
	instructionInstance Instruction,
	addressParameterIndex int,
) {
	nextIPRegisterValue := instructionInstance.Parameters[addressParameterIndex]
	// subtract the instruction length to compensate for increasing
	// the instruction pointer register in the ExecuteMachine function
	nextIPRegisterValue = nextIPRegisterValue - instructionInstance.Len()

	computer.SetIPRegisterValue(nextIPRegisterValue)
}
