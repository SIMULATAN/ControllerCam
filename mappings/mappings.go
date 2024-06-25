package mappings

import (
	"controllercontrol/config"
	"controllercontrol/state"
	"controllercontrol/utils"
	"github.com/0xcafed00d/joystick"
	"github.com/rs/zerolog/log"
)

type Controller interface {
	GetMapping(name string) *config.Input
}

func IsTriggered(c Controller, presetMapping config.Mapping, state joystick.State) bool {
	mapping := c.GetMapping(presetMapping.GetButton())
	if mapping == nil {
		return false
	}
	switch mapping.Type {
	case config.InputType_Button:
		return utils.GetButtonState(state.Buttons, uint32(mapping.Index))
	case config.InputType_StickAsButton:
		direction := mapping.Data.(int)
		axis := state.AxisData[mapping.Index]
		if direction == 1 {
			return axis > 0
		} else if direction == -1 {
			return axis < 0
		}
		log.Warn().Msgf("Could not handle StickAsButton with direction %v", direction)
		return false
	default:
		log.Warn().Msgf("Could not handle input of type %v", mapping.Type)
		return false
	}
}

func updateButtonStates(js joystick.State, buttons []state.ButtonState) {
	for i := range buttons {
		buttons[i].State = utils.GetButtonState(js.Buttons, uint32(buttons[i].Id))
	}
}

func updateStickStates(js joystick.State, sticks []state.StickState) {
	for i := range sticks {
		sticks[i].State = float64(js.AxisData[sticks[i].Id]) / float64(XboxStickMax)
	}
}
