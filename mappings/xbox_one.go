package mappings

import (
	"controllercontrol/config"
	"controllercontrol/state"
	"github.com/0xcafed00d/joystick"
	"github.com/rs/zerolog/log"
)

const (
	XboxOneA = 0
	XboxOneB = 1
	XboxOneX = 2
	XboxOneY = 3

	XboxStickMax = 32767

	XboxLeftStickX = 0
	XboxLeftStickY = 1

	XboxLeftTrigger = 2

	XboxRightStickX = 3
	XboxRightStickY = 4

	XboxRightTrigger = 5

	XboxDpadHorizontal = 6
	XboxDpadVertical   = 7

	XboxDeadzone = 0.1
)

type XboxController struct{}

func (c *XboxController) GetMapping(name string) *config.Input {
	switch name {
	case "DpadUp":
		return &config.Input{
			Type:  config.InputType_StickAsButton,
			Index: XboxDpadVertical,
			Data:  -1,
		}
	case "DpadDown":
		return &config.Input{
			Type:  config.InputType_StickAsButton,
			Index: XboxDpadVertical,
			Data:  1,
		}
	case "DpadLeft":
		return &config.Input{
			Type:  config.InputType_StickAsButton,
			Index: XboxDpadHorizontal,
			Data:  -1,
		}
	case "DpadRight":
		return &config.Input{
			Type:  config.InputType_StickAsButton,
			Index: XboxDpadHorizontal,
			Data:  1,
		}
	case "ButtonA":
		return &config.Input{
			Type:  config.InputType_Button,
			Index: XboxOneA,
		}
	case "ButtonB":
		return &config.Input{
			Type:  config.InputType_Button,
			Index: XboxOneB,
		}
	case "ButtonX":
		return &config.Input{
			Type:  config.InputType_Button,
			Index: XboxOneX,
		}
	case "ButtonY":
		return &config.Input{
			Type:  config.InputType_Button,
			Index: XboxOneY,
		}
	}

	log.Warn().Msgf("Could not find any input mapping for %v", name)
	return nil
}

var ButtonA = state.ButtonState{
	Id:    XboxOneA,
	Name:  "A",
	State: false,
}
var ButtonB = state.ButtonState{
	Id:    XboxOneB,
	Name:  "B",
	State: false,
}
var ButtonX = state.ButtonState{
	Id:    XboxOneX,
	Name:  "X",
	State: false,
}
var ButtonY = state.ButtonState{
	Id:    XboxOneY,
	Name:  "Y",
	State: false,
}

var StickLeftX = state.StickState{
	Id:    XboxLeftStickX,
	Name:  "Left X",
	State: 0,
}
var StickLeftY = state.StickState{
	Id:    XboxLeftStickY,
	Name:  "Left Y",
	State: 0,
}

var StickRightX = state.StickState{
	Id:    XboxRightStickX,
	Name:  "Right X",
	State: 0,
}
var StickRightY = state.StickState{
	Id:    XboxRightStickY,
	Name:  "Right Y",
	State: 0,
}

var Buttons = []state.ButtonState{ButtonA, ButtonB, ButtonX, ButtonY}
var Sticks = []state.StickState{StickLeftX, StickLeftY, StickRightX, StickRightY}

func UpdateIndexes(config config.RemappingsConfig) {
	for _, btn := range Buttons {
		remapping, ok := config.Remapping[btn.Name]
		if ok {
			btn.Id = remapping
		}
	}
	for _, stick := range Sticks {
		remapping, ok := config.Remapping[stick.Name]
		if ok {
			stick.Id = remapping
		}
	}
}

func UpdateStates(js joystick.State) {
	updateButtonStates(js, Buttons)
	updateStickStates(js, Sticks)
}
