package camera

import (
	"controllercontrol/config"
	"controllercontrol/mappings"
	"controllercontrol/utils"
	"github.com/0xcafed00d/joystick"
	"math"
)

func HandleJoystickInputs(handler ProtocolHandler, state joystick.State, config *config.Config) {
	HandlePanTilt(handler, state, config)
	HandleZoom(handler, state, config)
	HandlerPresets(handler, state, config)
}

func HandlePanTilt(handler ProtocolHandler, state joystick.State, cfg *config.Config) {
	step := cfg.Camera.Step

	pan := float64(state.AxisData[mappings.XboxLeftStickX]) / mappings.XboxStickMax
	pan = utils.CalculateExponentialValue(pan, cfg.Controller.PanTiltExponent)
	panOffset := int16(pan * step)
	tilt := float64(state.AxisData[mappings.XboxLeftStickY]) / mappings.XboxStickMax
	tilt = utils.CalculateExponentialValue(tilt, cfg.Controller.PanTiltExponent)
	tiltOffset := -1 * int16(tilt*step)

	byteSlice := PanTilt(
		cfg.Camera.PanSpeed*(1+int16(pan)*4),
		cfg.Camera.TiltSpeed*(1+int16(tilt)*4),
		panOffset,
		tiltOffset,
		true,
	)

	if (tiltOffset != 0 && math.Abs(tilt) > mappings.XboxDeadzone) || (panOffset != 0 && math.Abs(pan) > mappings.XboxDeadzone) {
		handler.SendPacketYolo(byteSlice)
	}
}

var lastZoomOffset int8

func HandleZoom(handler ProtocolHandler, state joystick.State, config *config.Config) {
	zoom := -1 * float64(state.AxisData[mappings.XboxRightStickY]) / mappings.XboxStickMax

	// if within deadzone, send zoom = 0
	if math.Abs(zoom) < mappings.XboxDeadzone {
		zoom = 0
	}

	// apply zoom factor
	zoom = utils.CalculateExponentialValue(zoom, config.Controller.ZoomExponent)
	zoom = math.Min(zoom+mappings.XboxDeadzone, 1)
	zoomOffset := int8(zoom * 7)

	// if zoom is the same as last time, don't send the packet
	if zoomOffset != lastZoomOffset {
		byteSlice := Zoom(zoomOffset)
		handler.SendPacketYolo(byteSlice)
	}
	lastZoomOffset = zoomOffset
}

func HandlerPresets(handler ProtocolHandler, state joystick.State, cfg *config.Config) {
	for _, item := range cfg.Mappings.Presets {
		if mappings.IsTriggered(handler.controller, &item, state) {
			handler.SendPacketYolo(RecallPreset(item.Preset))
		}
	}
}
