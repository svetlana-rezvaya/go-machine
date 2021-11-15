package machine

import "fmt"

// FetchInstruction ...
func FetchInstruction(computer Machine) (Instruction, error) {
	ipRegisterValue := computer.IPRegisterValue()
	instructionOpcode := computer.Memory[ipRegisterValue]
	if !IsOpcodeKnown(instructionOpcode) {
		return Instruction{}, fmt.Errorf(
			"unknown instruction 0x%x at address 0x%x",
			instructionOpcode,
			ipRegisterValue,
		)
	}

	fetchedInstruction := Instruction{Opcode: instructionOpcode}
	parameterCount := GetOpcodeParameterCount(instructionOpcode)
	for shift := 1; shift <= parameterCount; shift = shift + 1 {
		parameter := computer.Memory[ipRegisterValue+shift]
		fetchedInstruction.Parameters =
			append(fetchedInstruction.Parameters, parameter)
	}

	return fetchedInstruction, nil
}
