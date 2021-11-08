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
	}

	return nil
}
