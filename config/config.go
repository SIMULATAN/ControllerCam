package config

type Config struct {
	Camera     CameraConfig `yaml:"camera"`
	CameraHost string       `yaml:"camera_host"`
	JoystickId int          `yaml:"joystick_id"`
}

type CameraConfig struct {
	PanSpeed  int16   `yaml:"pan_speed"`
	TiltSpeed int16   `yaml:"tilt_speed"`
	Step      float64 `yaml:"step"`
}

func NewConfig() *Config {
	return &Config{
		Camera: CameraConfig{
			PanSpeed:  25,
			TiltSpeed: 25,
			Step:      20,
		},
	}
}
