package orio

import (
	"fmt"
	"os"
	"path/filepath"
)

func opiExportPin(pin Pin) error {
	path := filepath.Join(deviceDriversPath, "export")
	debug(fmt.Sprintf("Export pin %d (%s)", pin, path))
	f, err := os.OpenFile(path, os.O_WRONLY, 0200)
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("%d", pin))
	f.Close()
	return nil
}

func opiUnexportPin(pin Pin) error {
	path := filepath.Join(deviceDriversPath, "unexport")
	debug(fmt.Sprintf("Unxport pin %d (%s)", pin, path))
	f, err := os.OpenFile(path, os.O_WRONLY, 0200)
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("%d", pin))
	f.Close()
	return nil
}

func opiSetModePin(pin Pin, mode pinMode) error {
	path := filepath.Join(deviceDriversPath, fmt.Sprintf("gpio%d", pin), "direction")
	debug(fmt.Sprintf("Set mode pin %d (mode %d) (%s)", pin, mode, path))
	var err error
	f, err := os.OpenFile(path, os.O_WRONLY, 0200)
	if err != nil {
		return err
	}
	defer f.Close()
	switch mode {
	case Input:
		_, err = f.WriteString("in")
	case Output:
		_, err = f.WriteString("out")
	default:
		err = fmt.Errorf("not implementet yet")
	}
	if err != nil {
		return err
	}
	return nil
}

func opiWritePin(pin Pin, state pinState) error {
	path := filepath.Join(deviceDriversPath, fmt.Sprintf("gpio%d", pin), "value")
	debug(fmt.Sprintf("Write pin %d (value %d) (%s)", pin, state, path))
	f, err := os.OpenFile(path, os.O_WRONLY, 0200)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", state))
	return nil
}

func opiReadPin(pin Pin) (pinState, error) {
	path := filepath.Join(deviceDriversPath, fmt.Sprintf("gpio%d", pin), "value")
	debug(fmt.Sprintf("Read pin %d (%s)", pin, path))
	state, err := os.ReadFile(path)
	if err != nil {
		return Low, err
	}
	debug(fmt.Sprintf("Pin %d, state: %s", pin, string(state[0])))
	if state[0] == 0x31 {
	  	return High, nil
	}
	return Low, nil
}

// OrangePi name = pin
// pin = (position of letter in alphabet - 1) * 32 + number
const (
	PA0  pinNumber = 0
	PA1  pinNumber = 1
	PA2  pinNumber = 2
	PA3  pinNumber = 3
	PA4  pinNumber = 4
	PA5  pinNumber = 5
	PA6  pinNumber = 6
	PA7  pinNumber = 7
	PA8  pinNumber = 8
	PA9  pinNumber = 9
	PA10 pinNumber = 10
	PA11 pinNumber = 11
	PA12 pinNumber = 12
	PC0  pinNumber = 64
	PC1  pinNumber = 65
	PC2  pinNumber = 66
	PC3  pinNumber = 67
	PC4  pinNumber = 68
	PC5  pinNumber = 69
	PC6  pinNumber = 70
	PC7  pinNumber = 71
	PC8  pinNumber = 72
	PC9  pinNumber = 73
	PC10 pinNumber = 74
	PC11 pinNumber = 75
	PC12 pinNumber = 76
	PC13 pinNumber = 77
	PC14 pinNumber = 78
	PC15 pinNumber = 79
	PC16 pinNumber = 80
	PF0  pinNumber = 160
	PF1  pinNumber = 161
	PF2  pinNumber = 162
	PF3  pinNumber = 163
	PF4  pinNumber = 164
	PF5  pinNumber = 165
	PF6  pinNumber = 166
	PG0  pinNumber = 192
	PG1  pinNumber = 193
	PG2  pinNumber = 194
	PG3  pinNumber = 195
	PG4  pinNumber = 196
	PG5  pinNumber = 197
	PG6  pinNumber = 198
	PG7  pinNumber = 199
	PG8  pinNumber = 200
	PG9  pinNumber = 201
	PG10 pinNumber = 202
	PG11 pinNumber = 203
	PG12 pinNumber = 204
	PG13 pinNumber = 205
	PG14 pinNumber = 206
	PG15 pinNumber = 207
	PG16 pinNumber = 208
	PG17 pinNumber = 209
	PG18 pinNumber = 210
	PG19 pinNumber = 211
	PH0  pinNumber = 224
	PH1  pinNumber = 225
	PH2  pinNumber = 226
	PH3  pinNumber = 227
	PH4  pinNumber = 228
	PH5  pinNumber = 229
	PH6  pinNumber = 230
	PH7  pinNumber = 231
	PH8  pinNumber = 232
	PH9  pinNumber = 233
	PH10 pinNumber = 234
	PI0  pinNumber = 256
	PI1  pinNumber = 257
	PI2  pinNumber = 258
	PI3  pinNumber = 259
	PI4  pinNumber = 260
	PI5  pinNumber = 261
	PI6  pinNumber = 262
	PI7  pinNumber = 263
	PI8  pinNumber = 264
	PI9  pinNumber = 265
	PI10 pinNumber = 266
	PI11 pinNumber = 267
	PI12 pinNumber = 268
	PI13 pinNumber = 269
	PI14 pinNumber = 270
	PI15 pinNumber = 271
	PI16 pinNumber = 272
)
