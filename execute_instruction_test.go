package main

import (
	"reflect"
	"testing"
)

func Test_executeInstruction_withLoadConstantOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{12, 1, 5, 3, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 5, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withLoadMemoryOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 2,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{12, 1, 5, 3, 2, 7, 4, 13},
		registers:       []int{2, 9, 4, 8, 13, 1, 5, 3},
		ipRegisterIndex: 2,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withStoreConstantOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 3, 5, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{5, 3, 5, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withStoreMemoryOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{12, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 1, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withAdditionOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{5, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{15, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withSubtractionOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{6, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{-3, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withMultiplicationOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{7, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{54, 9, 4, 8, 6, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withDivisionOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 27, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{8, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{3, 9, 4, 8, 27, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}

func Test_executeInstruction_withModuloOpcode(test *testing.T) {
	machineInstance := machine{
		memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{2, 9, 4, 8, 21, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	instructionInstance, err := fetchInstruction(machineInstance)
	if err != nil {
		test.FailNow()
	}

	executeInstruction(machineInstance, instructionInstance)

	wantedMachineInstance := machine{
		memory:          []int{9, 4, 1, 0, 2, 17, 15, 13},
		registers:       []int{3, 9, 4, 8, 21, 0, 5, 3},
		ipRegisterIndex: 5,
	}
	if !reflect.DeepEqual(machineInstance, wantedMachineInstance) {
		test.Fail()
	}
}
