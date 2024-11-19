package config

type Config struct {
	DefaultCameraProperties CameraProperties `yaml:"default_camera_properties"`
	Controller              ControllerConfig `yaml:"controller"`
	Cameras                 []CameraConfig   `yaml:"cameras"`
	JoystickId              int              `yaml:"joystick_id"`
	Mappings                MappingsConfig   `yaml:"mappings"`
	Remappings              Remappings       `yaml:"remappings"`
}

type CameraConfig struct {
	Name       string           `yaml:"name"`
	Host       string           `yaml:"host"`
	Type       string           `yaml:"type"`
	Properties CameraProperties `yaml:"properties"`
}

type CameraProperties struct {
	PanSpeed  *int16   `yaml:"pan_speed"`
	TiltSpeed *int16   `yaml:"tilt_speed"`
	Step      *float64 `yaml:"step"`
}

type ControllerConfig struct {
	PanTiltExponent float64 `yaml:"pan_tilt_exponent"`
	ZoomExponent    float64 `yaml:"zoom_exponent"`
	Deadzone        float64 `yaml:"deadzone"`
}

var DefaultSpeed int16 = 25
var DefaultStep float64 = 50

func NewConfig() *Config {
	return &Config{
		DefaultCameraProperties: CameraProperties{
			PanSpeed:  &DefaultSpeed,
			TiltSpeed: &DefaultSpeed,
			Step:      &DefaultStep,
		},
		Controller: ControllerConfig{
			PanTiltExponent: 1,
			ZoomExponent:    1,
			Deadzone:        0.1,
		},
		JoystickId: 0,
	}
}
