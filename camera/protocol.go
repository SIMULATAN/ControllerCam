package camera

import (
	"controllercontrol/config"
	"controllercontrol/mappings"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
)

type ProtocolHandler struct {
	controller   mappings.Controller
	cameras      []*Camera
	ActiveCamera binding.Untyped
}

func RunLogQueue(camera *Camera, queue chan *[]byte) {
	for {
		packet := <-queue
		fmt.Println(camera.Model.InterpretResponse(*packet))
	}
}

func (p *ProtocolHandler) GetActiveCamera() *Camera {
	cam, err := p.ActiveCamera.Get()
	if err == nil {
		camAsCamera, worked := cam.(*Camera)
		if worked {
			return camAsCamera
		}
	}

	return nil
}

func (p *ProtocolHandler) SetActiveCamera(camera *Camera) {
	err := p.ActiveCamera.Set(camera)
	if err != nil {
		fmt.Println("Error setting active camera", err)
	}
}

func (p *ProtocolHandler) GetCameras() []*Camera {
	return p.cameras
}

func (p *ProtocolHandler) GetCameraByName(name string) *Camera {
	for _, camera := range p.cameras {
		if camera.Config.Name == name {
			return camera
		}
	}
	return nil
}

func NewProtocolHandler(cameraConfigs []config.CameraConfig, controller mappings.Controller) (*ProtocolHandler, error) {
	cameras := make([]*Camera, len(cameraConfigs))
	for i, camera := range cameraConfigs {
		model, err := CreateFromCameraConfig(camera)
		if err != nil {
			return nil, err
		}

		conn := model.Connect(camera)
		queue := make(chan *[]byte)
		conn.SetReceiveQueue(queue)

		camera := Camera{
			Model:  model,
			Config: camera,
			Conn:   conn,
		}
		cameras[i] = &camera

		go RunLogQueue(&camera, queue)
	}

	if len(cameras) == 0 {
		return nil, fmt.Errorf("no cameras found")
	}

	fmt.Printf("Created %v cameras: %+v\n", len(cameras), cameras)

	return &ProtocolHandler{
		cameras:      cameras,
		controller:   controller,
		ActiveCamera: binding.NewUntyped(),
	}, nil
}

func (c *Camera) SendPacket(bytes []byte) error {
	// Send the packet we created
	packet := c.Model.CreatePacket(bytes)
	err := c.Conn.Send(packet)
	return err
}

// SendPacketYolo same as SendPacket but just logs the error
func (c *Camera) SendPacketYolo(bytes []byte) {
	err := c.SendPacket(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
