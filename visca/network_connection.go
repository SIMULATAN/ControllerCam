package visca

import (
	"controllercontrol/utils"
	"github.com/josh23french/visca"
	"github.com/rs/zerolog/log"
	"net"
	"time"
)

// NetworkConnection implements the Iface interface without sending packets anywhere
type NetworkConnection struct {
	hostPort     string
	proto        string
	conn         net.Conn
	scanner      *Scanner
	receiveQueue chan *[]byte
	quit         chan struct{}
}

// Connect creates a new connection with the specified protocol and connects
func Connect(proto string, hostPort string) *NetworkConnection {
	conn := NetworkConnection{
		hostPort:     hostPort,
		proto:        proto,
		conn:         nil,
		scanner:      nil,
		receiveQueue: nil,
		quit:         make(chan struct{}),
	}
	start := func() {
		if err := conn.Start(); err != nil {
			log.Warn().Err(err).Msgf("Error starting %v connection to %v", proto, hostPort)
		}
	}
	go start()

	return &conn
}

func (i *NetworkConnection) IsConnected() bool {
	return i.conn != nil
}

// Start the interface
func (i *NetworkConnection) Start() error {
	conn, err := net.DialTimeout(i.proto, i.hostPort, time.Second)
	if err != nil {
		return err
	}

	i.conn = conn

	i.scanner = NewScanner(i.conn)

	go i.scanner.Scan(i.receiveQueue, i.quit)
	log.Info().Msgf("Started read loop from %v connection to %v", i.proto, i.hostPort)

	return nil
}

// Stop the interface
func (i *NetworkConnection) Stop() {
	if i.conn == nil {
		log.Warn().Msg("Never Started")
		return
	}

	// Stop the receive goroutine first
	close(i.quit)
	err := i.conn.Close()
	if err != nil {
		log.Warn().Err(err).Msgf("Error stopping %v connection to %v", i.proto, i.hostPort)
	}
	return
}

// Send a packet
func (i *NetworkConnection) Send(pkt []byte) error {
	if i.conn == nil {
		log.Warn().Msg("not started")
		return visca.ErrNotStarted
	}
	log.Debug().Msgf("Sending packet %v to %v", utils.DumpByteSlice(pkt), i.hostPort)
	written, err := i.conn.Write(pkt)
	if err != nil {
		return err
	}

	if written != len(pkt) {
		return visca.ErrIncompletePacketSent
	}

	return nil
}

// SetReceiveQueue for received packets
func (i *NetworkConnection) SetReceiveQueue(q chan *[]byte) {
	i.receiveQueue = q
	return
}
