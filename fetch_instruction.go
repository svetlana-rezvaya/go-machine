package main

import "fmt"

func fetchInstruction(computer machine) (instruction, error) {
	ipRegister := computer.registers[computer.ipRegisterIndex]
	instructionKind := computer.memory[ipRegister]
	if !isOpcodeKnown(instructionKind) {
		return instruction{}, fmt.Errorf(
			"unknown instruction 0x%x at address 0x%x",
			instructionKind,
			ipRegister,
		)
	}

	fetchedInstruction := instruction{kind: instructionKind}
	parameterCount := getOpcodeParameterCount(instructionKind)
	for shift := 1; shift <= parameterCount; shift = shift + 1 {
		parameter := computer.memory[ipRegister+shift]
		fetchedInstruction.parameters =
			append(fetchedInstruction.parameters, parameter)
	}

	return fetchedInstruction, nil
}
