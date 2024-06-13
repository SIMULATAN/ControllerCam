package mappings

import (
	"controllercontrol/config"
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
	}

	log.Warn().Msgf("Could not find any input mapping for %v", name)
	return nil
}
