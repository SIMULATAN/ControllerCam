package mappings

import (
	"controllercontrol/state"
	"fmt"
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

func (c *XboxController) GetButtonState(input *state.ButtonState, jsButtons uint32) bool {
	return jsButtons&(1<<input.Id) != 0
}

func (c *XboxController) GetStickState(input *state.StickState, jsAxis []int) float64 {
	return c.getStickState(jsAxis, input.Id)
}

func (c *XboxController) getStickState(jsAxis []int, id uint) float64 {
	return float64(jsAxis[id]) / float64(XboxStickMax)
}

func (c *XboxController) GetButtonStateFromStick(button *state.ButtonState, jsAxis []int) bool {
	value := c.getStickState(jsAxis, button.Id)
	if button.Name == DpadLeft || button.Name == DpadUp {
		return value < 0
	} else if button.Name == DpadRight || button.Name == DpadDown {
		return value > 0
	} else {
		fmt.Println("Don't know how to get button state from stick", button.Name)
		return false
	}
}

func (c *XboxController) GetInitialStates() state.States {
	return state.NewStates([]state.ButtonState{
		{
			Id:   XboxOneA,
			Name: ButtonA,
		},
		{
			Id:   XboxOneB,
			Name: ButtonB,
		},
		{
			Id:   XboxOneX,
			Name: ButtonX,
		},
		{
			Id:   XboxOneY,
			Name: ButtonY,
		},
		{
			Id:              XboxDpadHorizontal,
			Name:            DpadLeft,
			IsImplicitStick: true,
		},
		{
			Id:              XboxDpadHorizontal,
			Name:            DpadRight,
			IsImplicitStick: true,
		},
		{
			Id:              XboxDpadVertical,
			Name:            DpadUp,
			IsImplicitStick: true,
		},
		{
			Id:              XboxDpadVertical,
			Name:            DpadDown,
			IsImplicitStick: true,
		},
	}, []state.StickState{
		{
			Id:   XboxLeftStickX,
			Name: StickLeftX,
		},
		{
			Id:   XboxLeftStickY,
			Name: StickLeftY,
		},
		{
			Id:   XboxRightStickX,
			Name: RightStickX,
		},
		{
			Id:   XboxRightStickY,
			Name: RightStickY,
		},
	})
}
