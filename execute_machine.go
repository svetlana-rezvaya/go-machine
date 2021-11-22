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
		// if the instruction pointer register points at the next index
		// after the last one, then stop
		if nextIPRegisterValue == len(computer.Memory) {
			break
		}

		computer.SetIPRegisterValue(nextIPRegisterValue)
	}

	return nil
}
