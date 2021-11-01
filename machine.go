package machine

// Machine ...
type Machine struct {
	Memory          []int
	Registers       []int
	IPRegisterIndex int // index of the instruction pointer register
}
