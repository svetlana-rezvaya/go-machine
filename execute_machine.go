package machine

import "fmt"

// ExecuteMachine ...
func ExecuteMachine(computer Machine) error {
	for {
		instructionInstance, err := FetchInstruction(computer)
		if err != nil {
			return fmt.Errorf("unable to fetch the instruction: %v", err)
		}

		ExecuteInstruction(computer, instructionInstance)

		ipRegisterValue := computer.Registers[computer.IPRegisterIndex]
		nextIPRegisterValue :=
			ipRegisterValue + (GetOpcodeParameterCount(instructionInstance.Opcode) + 1)
		if nextIPRegisterValue > len(computer.Memory)-1 {
			break
		}

		computer.Registers[computer.IPRegisterIndex] = nextIPRegisterValue
	}

	return nil
}
