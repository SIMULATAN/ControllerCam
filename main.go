package main

import (
	"controllercontrol/camera"
	"controllercontrol/config"
	"controllercontrol/mappings"
	"fmt"
	"log"
	"time"

	"github.com/0xcafed00d/joystick"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalln("Error running program!", err)
	}
}

func run() error {
	js, err := joystick.Open(0)

	if err != nil {
		return err
	}
	fmt.Printf("Joystick Name: %s\n", js.Name())
	fmt.Printf("   Axis Count: %d\n", js.AxisCount())
	fmt.Printf(" Button Count: %d\n", js.ButtonCount())

	cfg := config.NewConfig()

	handler, err := camera.NewProtocolHandler("tcp://172.17.52.113:1259")
	if err != nil {
		return err
	}

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

		go camera.HandleJoystickInputs(*handler, state, cfg)

		time.Sleep(50 * time.Millisecond)
	}
	js.Close()

	return nil
}
