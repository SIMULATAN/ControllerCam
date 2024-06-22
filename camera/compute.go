package camera

import (
	"cmp"
	"controllercontrol/config"
	"controllercontrol/mappings"
	"controllercontrol/utils"
	"fmt"
	"github.com/0xcafed00d/joystick"
	"math"
)

func HandleJoystickInputs(handler ProtocolHandler, state joystick.State, config *config.Config) {
	HandleCameraSelection(handler, state, config)
	if activeCamera == nil {
		fmt.Println("No active camera")
		return
	}
	HandlePanTilt(state, config)
	HandleZoom(state, config)
	HandlePresets(handler, state, config)
}

var activeCamera *Camera

func HandleCameraSelection(handler ProtocolHandler, state joystick.State, cfg *config.Config) {
	for _, item := range cfg.Mappings.Cameras {
		if mappings.IsTriggered(handler.controller, &item, state) {
			camera := handler.GetCameraByName(item.Camera)
			if camera != nil && camera != activeCamera {
				activeCamera = camera
			}
		}
	}
}

func HandlePanTilt(state joystick.State, cfg *config.Config) {
	cam := activeCamera.config.Properties
	defaultCam := cfg.DefaultCameraProperties
	step := *cmp.Or(cam.Step, defaultCam.Step, &config.DefaultStep)

	panPercentRaw := float64(state.AxisData[mappings.XboxLeftStickX]) / mappings.XboxStickMax
	panPercent := utils.CalculateExponentialValue(panPercentRaw, cfg.Controller.PanTiltExponent)
	panOffset := int16(math.RoundToEven(panPercent * step))
	PanSpeed := *cmp.Or(cam.PanSpeed, defaultCam.PanSpeed, &config.DefaultSpeed)
	panSpeed := int16(float64(PanSpeed) * panPercent)

	tiltPercentRaw := float64(state.AxisData[mappings.XboxLeftStickY]) / mappings.XboxStickMax
	tiltPercent := utils.CalculateExponentialValue(tiltPercentRaw, cfg.Controller.PanTiltExponent)
	tiltOffset := -1 * int16(math.RoundToEven(tiltPercent*step))
	TiltSpeed := *cmp.Or(cam.TiltSpeed, defaultCam.TiltSpeed, &config.DefaultSpeed)
	tiltSpeed := int16(float64(TiltSpeed) * tiltPercent)

	byteSlice := PanTilt(
		panSpeed,
		tiltSpeed,
		panOffset,
		tiltOffset,
		true,
	)

	if (tiltOffset != 0 && math.Abs(tiltPercentRaw) > mappings.XboxDeadzone) || (panOffset != 0 && math.Abs(panPercentRaw) > mappings.XboxDeadzone) {
		activeCamera.SendPacketYolo(byteSlice)
	}
}

var lastZoomOffset int8

func HandleZoom(state joystick.State, config *config.Config) {
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
		activeCamera.SendPacketYolo(byteSlice)
	}
	lastZoomOffset = zoomOffset
}

func HandlePresets(handler ProtocolHandler, state joystick.State, cfg *config.Config) {
	for _, item := range cfg.Mappings.Presets {
		if mappings.IsTriggered(handler.controller, &item, state) {
			activeCamera.SendPacketYolo(RecallPreset(item.Preset))
		}
	}
}
