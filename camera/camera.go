package camera

import (
	camera_models "controllercontrol/camera/models"
	"controllercontrol/config"
	"controllercontrol/visca"
)

type Camera struct {
	Model  camera_models.CameraModel
	Conn   *visca.NetworkConnection
	Config config.CameraConfig
}
