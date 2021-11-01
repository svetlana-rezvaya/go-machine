package machine

import (
	"reflect"
	"testing"
)

func TestFetchInstruction_success(test *testing.T) {
	machineInstance := Machine{
		Memory:          []int{12, 19, 14, 18, 2, 17, 15, 13},
		Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		IPRegisterIndex: 2,
	}
	instructionInstance, err := FetchInstruction(machineInstance)

	wantedInstructionInstance := Instruction{
		Kind:       2,
		Parameters: []int{17, 15},
	}
	if !reflect.DeepEqual(instructionInstance, wantedInstructionInstance) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFetchInstruction_error(test *testing.T) {
	machineInstance := Machine{
		Memory:          []int{12, 19, 14, 18, 16, 17, 15, 13},
		Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		IPRegisterIndex: 2,
	}
	instructionInstance, err := FetchInstruction(machineInstance)

	wantedInstructionInstance := Instruction{}
	if !reflect.DeepEqual(instructionInstance, wantedInstructionInstance) {
		test.Fail()
	}

	wantedErrStr := "unknown instruction 0x10 at address 0x4"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
