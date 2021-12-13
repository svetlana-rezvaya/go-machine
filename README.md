# go-machine

[![GoDoc](https://godoc.org/github.com/svetlana-rezvaya/go-machine?status.svg)](https://godoc.org/github.com/svetlana-rezvaya/go-machine)
[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-machine)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-machine)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/go-machine.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/go-machine)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/go-machine/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/go-machine)

The library implementing a simple virtual machine.

## Instruction Set

[Instruction Set](docs/)

## Installation

```
$ go get github.com/svetlana-rezvaya/go-machine
```

## Usage Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/svetlana-rezvaya/go-machine"
)

func isPrime(number int) (bool, error) {
	const (
		inputRegisterIndex = iota
		outputRegisterIndex
		utilityRegisterIndex
		counterRegisterIndex
		ipRegisterIndex
		registerCount
	)

	const (
		loopIterationLabel = iota*100 + 100
		successLabel
		failureLabel
		finishLabel
	)

	const (
		successMark = iota + 1
		failureMark
	)

	program := []int{
		// load the input
		machine.LoadConstantOpcode, number, inputRegisterIndex,

		// compare the input with 2
		machine.LoadConstantOpcode, 2, utilityRegisterIndex,
		machine.SubtractionOpcode, inputRegisterIndex, utilityRegisterIndex, utilityRegisterIndex,
		machine.JumpIfNegativeOpcode, utilityRegisterIndex, failureLabel,
		machine.JumpIfZeroOpcode, utilityRegisterIndex, successLabel,

		// start the check loop
		// initialize the counter
		machine.LoadConstantOpcode, 2, counterRegisterIndex,
		machine.JumpOpcode, loopIterationLabel,
		// check the loop break condition: the counter should be less than the input
		loopIterationLabel: machine.SubtractionOpcode, inputRegisterIndex, counterRegisterIndex, utilityRegisterIndex,
		machine.JumpIfNegativeOpcode, utilityRegisterIndex, successLabel,
		machine.JumpIfZeroOpcode, utilityRegisterIndex, successLabel,
		// check divisibility of the input by the counter
		machine.ModuloOpcode, inputRegisterIndex, counterRegisterIndex, utilityRegisterIndex,
		machine.JumpIfZeroOpcode, utilityRegisterIndex, failureLabel,
		// increase the counter
		machine.LoadConstantOpcode, 1, utilityRegisterIndex,
		machine.AdditionOpcode, counterRegisterIndex, utilityRegisterIndex, counterRegisterIndex,
		// go to the next iteration
		machine.JumpOpcode, loopIterationLabel,

		// set the output to success
		successLabel: machine.LoadConstantOpcode, successMark, outputRegisterIndex,
		machine.JumpOpcode, finishLabel,

		// set the output to failure
		failureLabel: machine.LoadConstantOpcode, failureMark, outputRegisterIndex,
		machine.JumpOpcode, finishLabel,

		// no operation
		finishLabel: machine.JumpOpcode, finishLabel + 2,
	}

	machineInstance := machine.Machine{
		Memory:          program,
		Registers:       make([]int, registerCount),
		IPRegisterIndex: ipRegisterIndex,
	}
	err := machine.ExecuteMachine(machineInstance)
	if err != nil {
		return false, fmt.Errorf("unable to execute the machine: %s", err)
	}

	isSuccessful := machineInstance.Registers[outputRegisterIndex] == successMark
	return isSuccessful, nil
}

func main() {
	for number := 0; number < 100; number++ {
		isSuccessful, err := isPrime(number)
		if err != nil {
			log.Fatalf("unable to check primality of number %d: %s", number, err)
		}

		if isSuccessful {
			fmt.Printf("number %d is prime\n", number)
		}
	}

	// Output:
	// number 2 is prime
	// number 3 is prime
	// number 5 is prime
	// number 7 is prime
	// number 11 is prime
	// number 13 is prime
	// number 17 is prime
	// number 19 is prime
	// number 23 is prime
	// number 29 is prime
	// number 31 is prime
	// number 37 is prime
	// number 41 is prime
	// number 43 is prime
	// number 47 is prime
	// number 53 is prime
	// number 59 is prime
	// number 61 is prime
	// number 67 is prime
	// number 71 is prime
	// number 73 is prime
	// number 79 is prime
	// number 83 is prime
	// number 89 is prime
	// number 97 is prime
}
```

## License

The MIT License (MIT)

Copyright &copy; 2021 svetlana-rezvaya
