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

	panPercentRaw := float64(state.AxisData[mappings.XboxLeftStickX]) / mappings.XboxStickMax
	panPercent := utils.CalculateExponentialValue(panPercentRaw, cfg.Controller.PanTiltExponent)
	panOffset := int16(math.RoundToEven(panPercent * step))
	panSpeed := int16(float64(cfg.Camera.PanSpeed) * panPercent)

	tiltPercentRaw := float64(state.AxisData[mappings.XboxLeftStickY]) / mappings.XboxStickMax
	tiltPercent := utils.CalculateExponentialValue(tiltPercentRaw, cfg.Controller.PanTiltExponent)
	tiltOffset := -1 * int16(math.RoundToEven(tiltPercent*step))
	tiltSpeed := int16(float64(cfg.Camera.PanSpeed) * tiltPercent)

	byteSlice := PanTilt(
		panSpeed,
		tiltSpeed,
		panOffset,
		tiltOffset,
		true,
	)

	if (tiltOffset != 0 && math.Abs(tiltPercentRaw) > mappings.XboxDeadzone) || (panOffset != 0 && math.Abs(panPercentRaw) > mappings.XboxDeadzone) {
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
