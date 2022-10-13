/*
Package orio provides GPIO access on the OrangePi (Zero 2)
(c) Marcin Malessa https://github.com/mmalessa/orio
*/
package orio

import "fmt"

var DebugMode bool = false

type Pin uint16

func (pin Pin) Output() error {
	return registerPin(pin, Output)
}
func (pin Pin) Input() error {
	return registerPin(pin, Input)
}

func (pin Pin) Clock() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) Pwm() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) Close() error {
	return unregisterPin(pin)
}

func (pin Pin) High() error {
	return opiWritePin(pin, High)
}

func (pin Pin) Low() error {
	return opiWritePin(pin, Low)
}

func (pin Pin) Toggle() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) State() pinState {
	state, err := opiReadPin(pin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("State %d\n", state)
	return state
}

func (pin Pin) Pull() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) PullOff() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) PullUp() {

}

func (pin Pin) PullDown() {

}

func (pin Pin) ReadPull() error {
	return fmt.Errorf("not implemented yet")
}

func (pin Pin) Detect() {

}

func (pin Pin) EdgeDetected() {

}

func Close() error {
	for _, pin := range pins {
		if err := pin.Close(); err != nil {
			return err
		}
	}
	pins = pins[:0]
	return nil
}

const (
	Input pinMode = iota
	Output
	Clock
	Pwm
	Spi
	Alt0
	Alt1
	Alt2
	Alt3
	Alt4
	Alt5
)

const (
	Low  pinState = 0
	High pinState = 1
)
