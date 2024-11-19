package camera

import (
	models "controllercontrol/camera/models"
	"controllercontrol/config"
	"errors"
)

func CreateFromCameraConfig(config config.CameraConfig) (models.CameraModel, error) {
	switch config.Type {
	case "rgblink":
		return &models.RGBLink{}, nil
	case "marshall":
		return &models.Marshall{}, nil
	}

	return nil, errors.New("Unknown camera type: " + config.Type)
}
