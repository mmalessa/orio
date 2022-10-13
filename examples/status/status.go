package main

import (
	"time"

	"github.com/mmalessa/orio"
)

func main() {
	orio.DebugMode = true
	// defer orio.Close()

	btt := orio.Pin(orio.PH6)

	checkError(btt.Input())
	for i:=0; i<10; i++ {
		_ = btt.State()
		time.Sleep(500 * time.Millisecond)
	}
	
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}