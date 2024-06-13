package mappings

import (
	"controllercontrol/config"
	"controllercontrol/utils"
	"github.com/0xcafed00d/joystick"
	"github.com/rs/zerolog/log"
)

type Controller interface {
	GetMapping(name string) *config.Input
}

func IsTriggered(c Controller, presetMapping *config.PresetMapping, state joystick.State) bool {
	mapping := c.GetMapping(presetMapping.Button)
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
