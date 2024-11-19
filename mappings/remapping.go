package mappings

import (
	"controllercontrol/config"
	"controllercontrol/state"
)

func RemapStateIds(states state.States, config config.Remappings) {
	for key, value := range config {
		if button, ok := states.Buttons[key]; ok {
			button.Id = value
		}
		if stick, ok := states.Sticks[key]; ok {
			stick.Id = value
		}
	}
}
