package mappings

import (
	"controllercontrol/state"
	"github.com/0xcafed00d/joystick"
)

const (
	ButtonX = "ButtonX"
	ButtonY = "ButtonY"
	ButtonA = "ButtonA"
	ButtonB = "ButtonB"

	StickLeftX = "LeftStickX"
	StickLeftY = "LeftStickY"

	RightStickX = "RightStickX"
	RightStickY = "RightStickY"

	DpadUp    = "DpadUp"
	DpadRight = "DpadRight"
	DpadDown  = "DpadDown"
	DpadLeft  = "DpadLeft"
)

type Controller interface {
	GetInitialStates() state.States
	GetButtonState(button *state.ButtonState, jsButtons uint32) bool
	// GetStickState returns the state of the stick with the given id in the range of -1 to 1
	GetStickState(stick *state.StickState, jsAxis []int) float64
	// GetButtonStateFromStick returns the state of the stick with the given id as a button
	GetButtonStateFromStick(button *state.ButtonState, jsAxis []int) bool
}

func UpdateStates(controller Controller, states *state.States, js joystick.State) {
	for _, value := range states.Buttons {
		if value.IsImplicitStick {
			value.State = controller.GetButtonStateFromStick(value, js.AxisData)
		} else {
			value.State = controller.GetButtonState(value, js.Buttons)
		}
	}
	for _, value := range states.Sticks {
		value.State = controller.GetStickState(value, js.AxisData)
	}
}
