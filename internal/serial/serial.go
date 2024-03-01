package serial

import (
	"time"

	"go.bug.st/serial"
)

type Serial struct {
	Port   serial.Port
	Config SerialConfig
}

type SerialConfig struct {
	Name     string
	BaudRate int
	DataBits int
	StopBits serial.StopBits
	Parity   serial.Parity
}

func New(portName string, baudRate int, dataBits int, stopBits float32, parity string) *Serial {
	return &Serial{
		Port: nil,
		Config: SerialConfig{
			Name:     portName,
			BaudRate: baudRate,
			DataBits: dataBits,
			StopBits: convertStopBits(stopBits),
			Parity:   convertParity(parity),
		},
	}
}

func convertStopBits(stopBits float32) serial.StopBits {
	switch stopBits {
	case 1:
		return serial.OneStopBit
	case 1.5:
		return serial.OnePointFiveStopBits
	case 2:
		return serial.TwoStopBits
	default:
		return serial.OneStopBit
	}
}

func convertParity(parity string) serial.Parity {
	switch parity {
	case "none":
		return serial.NoParity
	case "odd":
		return serial.OddParity
	case "even":
		return serial.EvenParity
	case "mark":
		return serial.MarkParity
	case "space":
		return serial.SpaceParity
	default:
		return serial.NoParity
	}
}

func (s *Serial) Open() error {
	port, err := serial.Open(s.Config.Name, &serial.Mode{
		BaudRate: s.Config.BaudRate,
		DataBits: s.Config.DataBits,
		StopBits: s.Config.StopBits,
		Parity:   s.Config.Parity,
	})
	if err != nil {
		return err
	}

	s.Port = port
	return nil
}

func (s *Serial) Close() error {
	if s.Port == nil {
		return nil
	}

	return s.Port.Close()
}

func (s *Serial) SendCommand(command string) (string, error) {
	if s.Port == nil {
		return "", nil
	}

	_, err := s.Port.Write([]byte(command))
	if err != nil {
		return "", err
	}

	time.Sleep(3 * time.Second)

	buf := make([]byte, 128)
	n, err := s.Port.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func GetSerialPorts() ([]string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	return ports, nil
}
