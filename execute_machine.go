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

		ipRegisterValue := computer.IPRegisterValue()
		nextIPRegisterValue := ipRegisterValue + instructionInstance.Len()
		if nextIPRegisterValue > len(computer.Memory)-1 {
			break
		}

		computer.SetIPRegisterValue(nextIPRegisterValue)
	}

	return nil
}
