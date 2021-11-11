package machine

import (
	"reflect"
	"testing"
)

func TestExecuteMachine(test *testing.T) {
	type data struct {
		name                  string
		machineInstance       Machine
		wantedMachineInstance Machine
		wantedErrStr          string
	}

	tests := []data{
		data{
			name: "success",
			machineInstance: Machine{
				Memory:          []int{1, 2, 0, 1, 3, 1, 5, 0, 1, 2},
				Registers:       []int{0, 0, 0, 0},
				IPRegisterIndex: 3,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{1, 2, 0, 1, 3, 1, 5, 0, 1, 2},
				Registers:       []int{2, 3, 5, 6},
				IPRegisterIndex: 3,
			},
			wantedErrStr: "",
		},
		data{
			name: "error",
			machineInstance: Machine{
				Memory:          []int{1, 2, 0, 23, 3, 1, 5, 0, 1, 2},
				Registers:       []int{0, 0, 0, 0},
				IPRegisterIndex: 3,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{1, 2, 0, 23, 3, 1, 5, 0, 1, 2},
				Registers:       []int{2, 0, 0, 3},
				IPRegisterIndex: 3,
			},
			wantedErrStr: "unable to fetch the instruction: " +
				"unknown instruction 0x17 at address 0x3",
		},
	}
	for _, testData := range tests {
		err := ExecuteMachine(testData.machineInstance)

		if !reflect.DeepEqual(
			testData.machineInstance,
			testData.wantedMachineInstance,
		) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedMachineInstance,
				testData.machineInstance,
			)
			test.Fail()
		}

		wantedErr := testData.wantedErrStr != ""
		if !wantedErr && err != nil ||
			wantedErr && (err == nil || err.Error() != testData.wantedErrStr) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedErrStr,
				err,
			)
			test.Fail()
		}
	}
}
