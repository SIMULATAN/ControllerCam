package camera

import (
	"cmp"
	"controllercontrol/config"
	"controllercontrol/mappings"
	"controllercontrol/state"
	"controllercontrol/utils"
	"math"
)

func HandleJoystickInputs(handler ProtocolHandler, states state.States, config *config.Config) {
	HandleCameraSelection(handler, states, config)
	if handler.GetActiveCamera() == nil {
		return
	}
	HandlePanTilt(handler, states, config)
	HandleZoom(handler, states, config)
	HandlePresets(handler, states, config)
}

func HandleCameraSelection(handler ProtocolHandler, states state.States, cfg *config.Config) {
	for _, item := range cfg.Mappings.Cameras {
		if val, err := states.GetButtonState(item.Button); err == nil && val {
			camera := handler.GetCameraByName(item.Camera)
			if camera != nil && camera != handler.GetActiveCamera() {
				handler.SetActiveCamera(camera)
			}
		}
	}
}

func HandlePanTilt(handler ProtocolHandler, states state.States, cfg *config.Config) {
	x, err := states.GetStickState(mappings.StickLeftX)
	if err != nil {
		return
	}
	y, err := states.GetStickState(mappings.StickLeftY)
	if err != nil {
		return
	}

	camera := handler.GetActiveCamera()
	props := camera.Config.Properties
	defaultCam := cfg.DefaultCameraProperties
	step := *cmp.Or(props.Step, defaultCam.Step, &config.DefaultStep)

	panPercent := utils.CalculateExponentialValue(x, cfg.Controller.PanTiltExponent)
	panOffset := int16(math.RoundToEven(panPercent * step))
	PanSpeed := *cmp.Or(props.PanSpeed, defaultCam.PanSpeed, &config.DefaultSpeed)
	panSpeed := int16(float64(PanSpeed) * panPercent)

	tiltPercent := utils.CalculateExponentialValue(y, cfg.Controller.PanTiltExponent)
	tiltOffset := -1 * int16(math.RoundToEven(tiltPercent*step))
	TiltSpeed := *cmp.Or(props.TiltSpeed, defaultCam.TiltSpeed, &config.DefaultSpeed)
	tiltSpeed := int16(float64(TiltSpeed) * tiltPercent)

	byteSlice := camera.Model.PanTilt(
		panSpeed,
		tiltSpeed,
		panOffset,
		tiltOffset,
		true,
	)

	if (tiltOffset != 0 && math.Abs(y) > mappings.XboxDeadzone) || (panOffset != 0 && math.Abs(x) > mappings.XboxDeadzone) {
		camera.SendPacketYolo(byteSlice)
	}
}

var lastZoomOffset int8

func HandleZoom(handler ProtocolHandler, states state.States, config *config.Config) {
	measurement, err := states.GetStickState(mappings.RightStickY)
	if err != nil {
		return
	}
	zoom := -1 * measurement

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
		handler.GetActiveCamera().SendPacketYolo(byteSlice)
	}
	lastZoomOffset = zoomOffset
}

func HandlePresets(handler ProtocolHandler, states state.States, cfg *config.Config) {
	for _, item := range cfg.Mappings.Presets {
		if val, err := states.GetButtonState(item.Button); err == nil && val {
			camera := handler.GetActiveCamera()
			camera.SendPacketYolo(camera.Model.RecallPreset(item.Preset))
		}
	}
}
