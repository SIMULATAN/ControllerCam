package camera

import (
	"controllercontrol/config"
	"controllercontrol/mappings"
	"controllercontrol/utils"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"github.com/josh23french/visca"
)

func RunLogQueue(queue chan *visca.Packet) {
	for {
		packet := <-queue
		fmt.Println(InterpretResponse(packet))
	}
}

func BuildPacket(bytes []byte) (*visca.Packet, error) {
	packet, err := visca.NewPacket(0, 1, bytes)
	if packet != nil {
		fmt.Println(utils.DumpByteSlice(packet.Bytes()))
	}
	return packet, err
}

func InterpretResponse(packet *visca.Packet) string {
	message := packet.Message
	if utils.BytesEqual(message, RETURN_ACK) {
		return "[SUC] ACK"
	} else if utils.BytesEqual(message, RETURN_Completion) {
		return "[SUC] Completion"
	} else if utils.BytesEqual(message, RETURN_SyntaxError) {
		return "[ERR] Syntax Error"
	} else if utils.BytesEqual(message, RETURN_CommandNotExecutable) {
		return "[ERR] Command not executable"
	}

	return utils.DumpByteSlice(packet.Message)
}

type Camera struct {
	Conn   *visca.Connection
	Config config.CameraConfig
}

type ProtocolHandler struct {
	controller   mappings.Controller
	cameras      []Camera
	ActiveCamera binding.Untyped
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

func (p *ProtocolHandler) GetCameras() []Camera {
	return p.cameras
}

func (p *ProtocolHandler) GetCameraByName(name string) *Camera {
	for _, camera := range p.cameras {
		if camera.Config.Name == name {
			return &camera
		}
	}
	return nil
}

func NewProtocolHandler(cameraConfigs []config.CameraConfig, controller mappings.Controller) (*ProtocolHandler, error) {
	cameras := make([]Camera, len(cameraConfigs))
	for i, camera := range cameraConfigs {
		conn, err := visca.NewConnectionFromString(camera.Host)
		if err != nil {
			return nil, err
		}

		queue := make(chan *visca.Packet)
		conn.SetReceiveQueue(queue)
		go RunLogQueue(queue)
		err = conn.Start()

		var connection *visca.Connection = nil
		if err == nil {
			connection = &conn
		}
		cameras[i] = Camera{
			Conn:   connection,
			Config: camera,
		}
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
	packet, err := BuildPacket(bytes)
	if err != nil {
		return err
	}
	if c.Conn == nil {
		return fmt.Errorf("camera connection isn't open")
	}
	err = (*c.Conn).Send(packet)
	return err
}

// SendPacketYolo same as SendPacket but just logs the error
func (c *Camera) SendPacketYolo(bytes []byte) {
	err := c.SendPacket(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
