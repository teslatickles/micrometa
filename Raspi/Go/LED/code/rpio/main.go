package main

import (
	"fmt"
	"log"

	term "github.com/nsf/termbox-go"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	flashLed()
}

func reset() {
	err := term.Sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Press Up Arrow Key to turn LED on || Press Down Arrow Key to turn LED off")
}

func flashLed() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("unable to open gpio")
		log.Fatal(err)
	}
	defer rpio.Close()

	// physical pin 12
	pin := rpio.Pin(18)
	pin.Output()

	err = term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				reset()
				pin.High()
				currentState := pin.Read()
				fmt.Printf("\nLED pin set to %d - (HIGH)", currentState)
			case term.KeyArrowDown:
				reset()
				pin.Low()
				currentState := pin.Read()
				fmt.Printf("\nLED pin set to %d - (LOW)", currentState)
			default:
				// we only want to read a single character or one key pressed event
				reset()
				fmt.Println("ASCII : ", ev.Key)
			}
		case term.EventError:
			pin.Low()
			panic(ev.Err)
		}
	}
}
