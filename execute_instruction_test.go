package machine

import (
	"reflect"
	"testing"
)

func TestExecuteInstruction(test *testing.T) {
	type data struct {
		name                  string
		machineInstance       Machine
		wantedMachineInstance Machine
	}

	tests := []data{
		data{
			name: "loadConstantOpcode",
			machineInstance: Machine{
				Memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 5, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "loadMemoryOpcode",
			machineInstance: Machine{
				Memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 2,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
				Registers:       []int{2, 9, 4, 8, 13, 1, 5, 3},
				IPRegisterIndex: 2,
			},
		},
		data{
			name: "storeConstantOpcode",
			machineInstance: Machine{
				Memory:          []int{12, 3, 5, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{5, 3, 5, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "storeMemoryOpcode",
			machineInstance: Machine{
				Memory:          []int{12, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "additionOpcode",
			machineInstance: Machine{
				Memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{15, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "subtractionOpcode",
			machineInstance: Machine{
				Memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{-3, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "multiplicationOpcode",
			machineInstance: Machine{
				Memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{54, 9, 4, 8, 6, 0, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "divisionOpcode",
			machineInstance: Machine{
				Memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 27, 0, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{3, 9, 4, 8, 27, 0, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "moduloOpcode",
			machineInstance: Machine{
				Memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{2, 9, 4, 8, 21, 0, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				Registers:       []int{3, 9, 4, 8, 21, 0, 5, 3},
				IPRegisterIndex: 5,
			},
		},
		data{
			name: "jumpOpcode",
			machineInstance: Machine{
				Memory:          []int{5, 10, 4, 0, 10, 1, 7, 6},
				Registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				IPRegisterIndex: 5,
			},
			wantedMachineInstance: Machine{
				Memory:          []int{5, 10, 4, 0, 10, 1, 7, 6},
				Registers:       []int{2, 9, 4, 8, 6, 2, 5, 3},
				IPRegisterIndex: 5,
			},
		},
	}
	for _, testData := range tests {
		instructionInstance, err := FetchInstruction(testData.machineInstance)
		if err != nil {
			test.Logf("failed %q: %s", testData.name, err)
			test.FailNow()
		}

		ExecuteInstruction(testData.machineInstance, instructionInstance)

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
	}
}
