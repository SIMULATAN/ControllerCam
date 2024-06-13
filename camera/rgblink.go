package camera

var (
	SYS_Menu_ON          = []byte{0x01, 0x04, 0x06, 0x06, 0x02}
	CAM_Picture_Flip_ON  = []byte{0x01, 0x04, 0x66, 0x02}
	CAM_Picture_Flip_OFF = []byte{0x01, 0x04, 0x66, 0x03}

	RETURN_ACK                  = []byte{0x41}
	RETURN_Completion           = []byte{0x51}
	RETURN_SyntaxError          = []byte{0x60, 0x02}
	RETURN_CommandNotExecutable = []byte{0x61, 0x41}
)

func CAM_Brightness(p byte, q byte) []byte {
	return []byte{0x01, 0x04, 0xA1, 0x00, 0x00, p, q}
}

func RecallPreset(number uint8) []byte {
	return []byte{0x01, 0x04, 0x3F, 0x02, number}
}

type RGBLink struct {
	deviceId int
}
