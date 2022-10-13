package orio

import "fmt"

var pins []Pin
var deviceDriversPath string = "/sys/class/gpio/"

type pinMode uint8
type pinNumber uint16
type pinState uint8
type pinEdge uint8

func debug(text string) {
	if !DebugMode {
		return
	}
	fmt.Println(text)
}

func registerPin(pin Pin, mode pinMode) error {
	pins = append(pins, pin)
	debug(fmt.Sprintf("Set pin %d to mode %d", pin, mode))

	if err := opiExportPin(pin); err != nil {
		fmt.Println(err)
		return err
	}
	if err := opiSetModePin(pin, mode); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func unregisterPin(pin Pin) error {
	return opiUnexportPin(pin)
}
