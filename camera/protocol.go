package camera

import (
	"controllercontrol/utils"
	"fmt"
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

type ProtocolHandler struct {
	conn *visca.Connection
}

func NewProtocolHandler(connectionString string) (*ProtocolHandler, error) {
	conn, err := visca.NewConnectionFromString(connectionString)
	if err != nil {
		return nil, err
	}

	queue := make(chan *visca.Packet)
	conn.SetReceiveQueue(queue)
	go RunLogQueue(queue)
	err = conn.Start()
	if err != nil {
		return nil, err
	}

	return &ProtocolHandler{
		conn: &conn,
	}, nil
}

func (p *ProtocolHandler) SendPacket(bytes []byte) error {
	// Send the packet we created
	packet, err := BuildPacket(bytes)
	if err != nil {
		return err
	}
	err = (*p.conn).Send(packet)
	return err
}

// SendPacketYolo same as SendPacket but just logs the error
func (p *ProtocolHandler) SendPacketYolo(bytes []byte) {
	err := p.SendPacket(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
