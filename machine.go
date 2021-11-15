package machine

// Machine ...
type Machine struct {
	Memory          []int
	Registers       []int
	IPRegisterIndex int // index of the instruction pointer register
}

// IPRegisterValue ...
func (computer Machine) IPRegisterValue() int {
	return computer.Registers[computer.IPRegisterIndex]
}
