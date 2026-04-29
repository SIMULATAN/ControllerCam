package mappings

import (
	"controllercontrol/config"
	"controllercontrol/state"
)

func RemapStateIds(states state.States, config config.Remappings) {
	for key, value := range config {
		if button, ok := states.Buttons[key]; ok {
			button.Id = value.Id
			if value.Name != nil {
				button.Name = *value.Name
			}
			if value.IsImplicitStick != nil {
				button.IsImplicitStick = *value.IsImplicitStick
			}
		}
		if stick, ok := states.Sticks[key]; ok {
			stick.Id = value.Id
			if value.Name != nil {
				stick.Name = *value.Name
			}
			if value.Invert != nil {
				stick.Invert = *value.Invert
			}
		}
	}
}
