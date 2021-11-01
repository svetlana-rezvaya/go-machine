package machine

import (
	"reflect"
	"testing"
)

func Test_fetchInstruction_success(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 19, 14, 18, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 2,
	}
	instructionInstance, err := fetchInstruction(machineInstance)

	wantedInstructionInstance := instruction{
		kind:       2,
		parameters: []int{17, 15},
	}
	if !reflect.DeepEqual(instructionInstance, wantedInstructionInstance) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_fetchInstruction_error(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 19, 14, 18, 16, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 2,
	}
	instructionInstance, err := fetchInstruction(machineInstance)

	wantedInstructionInstance := instruction{}
	if !reflect.DeepEqual(instructionInstance, wantedInstructionInstance) {
		test.Fail()
	}

	wantedErrStr := "unknown instruction 0x10 at address 0x4"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
