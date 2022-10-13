package main

import (
	"time"
	"fmt"

	"github.com/mmalessa/orio"
)

func main() {
	orio.DebugMode = true
	defer orio.Close()

	btt := orio.Pin(orio.PH6)

	checkError(btt.Input())
	for i:=0; i<10; i++ {
		state := btt.State()
		fmt.Printf("State: %d", state)
		time.Sleep(500 * time.Millisecond)
	}
	
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}