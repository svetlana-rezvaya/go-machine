package machine

import "fmt"

// FetchInstruction ...
func FetchInstruction(computer Machine) (Instruction, error) {
	ipRegister := computer.Registers[computer.IPRegisterIndex]
	instructionKind := computer.Memory[ipRegister]
	if !IsOpcodeKnown(instructionKind) {
		return Instruction{}, fmt.Errorf(
			"unknown instruction 0x%x at address 0x%x",
			instructionKind,
			ipRegister,
		)
	}

	fetchedInstruction := Instruction{Kind: instructionKind}
	parameterCount := GetOpcodeParameterCount(instructionKind)
	for shift := 1; shift <= parameterCount; shift = shift + 1 {
		parameter := computer.Memory[ipRegister+shift]
		fetchedInstruction.Parameters =
			append(fetchedInstruction.Parameters, parameter)
	}

	return fetchedInstruction, nil
}
