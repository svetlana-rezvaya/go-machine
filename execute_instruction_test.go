package machine

import (
	"reflect"
	"testing"
)

func Test_executeInstruction(test *testing.T) {
	type data struct {
		name                  string
		machineInstance       machine
		wantedMachineInstance machine
	}

	tests := []data{
		data{
			name: "loadConstantOpcode",
			machineInstance: machine{
				memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 5, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "loadMemoryOpcode",
			machineInstance: machine{
				memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 2,
			},
			wantedMachineInstance: machine{
				memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
				registers:       []int{2, 9, 4, 8, 13, 1, 5, 3},
				ipRegisterIndex: 2,
			},
		},
		data{
			name: "storeConstantOpcode",
			machineInstance: machine{
				memory:          []int{12, 3, 5, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{5, 3, 5, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "storeMemoryOpcode",
			machineInstance: machine{
				memory:          []int{12, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "additionOpcode",
			machineInstance: machine{
				memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{15, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "subtractionOpcode",
			machineInstance: machine{
				memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{-3, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "multiplicationOpcode",
			machineInstance: machine{
				memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{54, 9, 4, 8, 6, 0, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "divisionOpcode",
			machineInstance: machine{
				memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 27, 0, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{3, 9, 4, 8, 27, 0, 5, 3},
				ipRegisterIndex: 5,
			},
		},
		data{
			name: "moduloOpcode",
			machineInstance: machine{
				memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{2, 9, 4, 8, 21, 0, 5, 3},
				ipRegisterIndex: 5,
			},
			wantedMachineInstance: machine{
				memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
				registers:       []int{3, 9, 4, 8, 21, 0, 5, 3},
				ipRegisterIndex: 5,
			},
		},
	}
	for _, testData := range tests {
		instructionInstance, err := fetchInstruction(testData.machineInstance)
		if err != nil {
			test.Logf("failed %q: %s", testData.name, err)
			test.FailNow()
		}

		executeInstruction(testData.machineInstance, instructionInstance)

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
