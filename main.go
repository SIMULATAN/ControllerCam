package main

import (
	"controllercontrol/camera"
	"controllercontrol/config"
	"controllercontrol/gui"
	"controllercontrol/mappings"
	"controllercontrol/state"
	"controllercontrol/utils"
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
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	js, err := joystick.Open(cfg.JoystickId)

	if err != nil {
		return err
	}
	fmt.Printf("Joystick Name: %s\n", js.Name())
	fmt.Printf("   Axis Count: %d\n", js.AxisCount())
	fmt.Printf(" Button Count: %d\n", js.ButtonCount())

	states := state.NewStates(mappings.Buttons, mappings.Sticks)

	go handleLoop(cfg, js, &states)

	err = gui.RunGui(&states)
	return err
}

func handleLoop(cfg *config.Config, js joystick.Joystick, states *state.States) {
	handler, err := camera.NewProtocolHandler(cfg.Cameras, &mappings.XboxController{})
	if err != nil {
		log.Fatalln("Error creating ProtocolHandler!", err)
	}

	for {
		controllerState, err := js.Read()
		if err != nil {
			log.Println("Error reading from Joystick!", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Printf("Axis Data: %v\n", controllerState.AxisData)
		fmt.Printf("Button Data: %v\n", controllerState.Buttons)
		fmt.Printf("Button States: [%v, %v, %v, %v]\n",
			utils.GetButtonState(controllerState.Buttons, mappings.XboxOneA),
			utils.GetButtonState(controllerState.Buttons, mappings.XboxOneB),
			utils.GetButtonState(controllerState.Buttons, mappings.XboxOneX),
			utils.GetButtonState(controllerState.Buttons, mappings.XboxOneY),
		)

		mappings.UpdateStates(controllerState)
		states.UpdateCallback()

		go camera.HandleJoystickInputs(*handler, controllerState, cfg)

		time.Sleep(50 * time.Millisecond)
	}
	js.Close()
}
