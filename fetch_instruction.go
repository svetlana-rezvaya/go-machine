package main

import "fmt"

func fetchInstruction(computer machine) (instruction, error) {
	ipRegister := computer.registers[computer.ipRegisterIndex]
	instructionKind := computer.memory[ipRegister]

	fetchedInstruction := instruction{kind: instructionKind}
	if instructionKind == loadConstantOpcode {
		constant := computer.memory[ipRegister+1]
		registerIndex := computer.memory[ipRegister+2]

		fetchedInstruction.parameters = []int{constant, registerIndex}
	} else if instructionKind == loadMemoryOpcode {
		memoryIndex := computer.memory[ipRegister+1]
		registerIndex := computer.memory[ipRegister+2]

		fetchedInstruction.parameters = []int{memoryIndex, registerIndex}
	} else {
		return instruction{}, fmt.Errorf(
			"unknown instruction 0x%x at address 0x%x",
			instructionKind,
			ipRegister,
		)
	}

	return fetchedInstruction, nil
}
