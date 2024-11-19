package main

import (
	"controllercontrol/camera"
	"controllercontrol/config"
	"controllercontrol/gui"
	"controllercontrol/mappings"
	"controllercontrol/state"
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

	var controller mappings.Controller = &mappings.XboxController{}
	states := controller.GetInitialStates()

	mappings.RemapStateIds(states, cfg.Remappings)

	handler, err := camera.NewProtocolHandler(cfg.Cameras, controller)
	if err != nil {
		log.Fatalln("Error creating ProtocolHandler!", err)
	}

	go handleLoop(cfg, controller, js, &states, handler)

	err = gui.RunGui(&states, handler)
	return err
}

func handleLoop(
	cfg *config.Config,
	controller mappings.Controller,
	js joystick.Joystick,
	states *state.States,
	handler *camera.ProtocolHandler,
) {
	for {
		controllerState, err := js.Read()
		if err != nil {
			log.Println("Error reading from Joystick!", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		mappings.UpdateStates(controller, states, controllerState)
		states.UpdateCallback()

		go camera.HandleJoystickInputs(*handler, *states, cfg)

		time.Sleep(50 * time.Millisecond)
	}
	js.Close()
}
