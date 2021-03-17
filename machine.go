package main

type machine struct {
	memory          []int
	registers       []int
	ipRegisterIndex int // index of the instruction pointer register
}
