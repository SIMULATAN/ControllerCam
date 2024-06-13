package config

type Config struct {
	Camera CameraConfig
}

type CameraConfig struct {
	PanSpeed  int16
	TiltSpeed int16
	Step      float64
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
