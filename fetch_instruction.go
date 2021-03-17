package main

func fetchInstruction(computer machine) (instruction, error) {
	ipRegister := computer.registers[computer.ipRegisterIndex]
	instructionKind := computer.memory[ipRegister]
	fetchedInstruction := instruction{kind: instructionKind}
	return fetchedInstruction, nil
}
