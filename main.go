package main

import (
	"controllercontrol/mappings"
	"fmt"
	"log"
	"time"

	"github.com/0xcafed00d/joystick"
)

func main() {
	js, err := joystick.Open(1)
	if err != nil {
		log.Fatalln("Error opening Joystick", err)
	}

	fmt.Printf("Joystick Name: %s\n", js.Name())
	fmt.Printf("   Axis Count: %d\n", js.AxisCount())
	fmt.Printf(" Button Count: %d\n", js.ButtonCount())

	for {
		state, err := js.Read()
		if err != nil {
			log.Println("Error reading from Joystick!", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Printf("Axis Data: %v\n", state.AxisData)
		fmt.Printf("Button Data: %v\n", state.Buttons)
		fmt.Printf("Button States: [%v, %v, %v, %v]\n",
			GetButtonState(state.Buttons, mappings.XboxOneA),
			GetButtonState(state.Buttons, mappings.XboxOneB),
			GetButtonState(state.Buttons, mappings.XboxOneX),
			GetButtonState(state.Buttons, mappings.XboxOneY),
		)
		time.Sleep(100 * time.Millisecond)
	}
	js.Close()
}
