package visca

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"io"
)

// Scanner represents a
type Scanner struct {
	scanner *bufio.Scanner
	buffer  io.Reader
}

func splitPackets(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		// If we found a terminator byte, we have a packet
		if data[i] == 0xFF {
			// Include that FF in the packet!
			return i + 1, data[:i+1], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

// NewScanner constructs a Scanner
func NewScanner(buffer io.Reader) *Scanner {
	scanner := bufio.NewScanner(buffer)
	scanner.Buffer([]byte{}, 32)
	scanner.Split(splitPackets)

	return &Scanner{
		scanner,
		buffer,
	}
}

// Scan sends packets to the given channel
func (s *Scanner) Scan(c chan *[]byte, quit chan struct{}) {
loop:
	for {
		select {
		case <-quit:
			break loop
		default:
			ok := s.scanner.Scan()
			if !ok {
				log.Err(s.scanner.Err()).Msg("Scan stopped")
				break loop
			}
			packetBytes := s.scanner.Bytes()
			if len(packetBytes) == 0 {
				log.Warn().Msg("Packet is empty!!!")
				continue
			}
			if len(packetBytes) <= 2 {
				log.Warn().Msg("Packet is smaller than 3 bytes!")
				// Skip runts (scanner returns a zero-length slice last in the common case)
				continue
			}
			c <- &packetBytes
		}
	}
	if c != nil {
		close(c)
	}
}
