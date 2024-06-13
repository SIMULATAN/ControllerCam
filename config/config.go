package config

import "controllercontrol/mappings"

type Config struct {
	Camera     CameraConfig     `yaml:"camera"`
	Controller ControllerConfig `yaml:"controller"`
	CameraHost string           `yaml:"camera_host"`
	JoystickId int              `yaml:"joystick_id"`
}

type CameraConfig struct {
	PanSpeed  int16   `yaml:"pan_speed"`
	TiltSpeed int16   `yaml:"tilt_speed"`
	Step      float64 `yaml:"step"`
}

type ControllerConfig struct {
	PanTiltExponent float64 `yaml:"pan_tilt_exponent"`
	ZoomExponent    float64 `yaml:"zoom_exponent"`
	Deadzone        float64 `yaml:"deadzone"`
}

func NewConfig() *Config {
	return &Config{
		Camera: CameraConfig{
			PanSpeed:  25,
			TiltSpeed: 25,
			Step:      20,
		},
		Controller: ControllerConfig{
			PanTiltExponent: 1,
			ZoomExponent:    1,
			Deadzone:        mappings.XboxDeadzone,
		},
		JoystickId: 0,
	}
}
