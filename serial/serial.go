package serial

import "go.bug.st/serial"

type Serial struct {
	port serial.Port
}

type Mode struct {
	serial.Mode
}

func GetPortsList() ([]string, error) {
	return serial.GetPortsList()
}

func Open(name string, mode *Mode) (*Serial, error) {
	port, err := serial.Open(name, &mode.Mode)
	if err != nil {
		return nil, err
	}
	return &Serial{port: port}, nil
}

func (s *Serial) Close() error {
	return s.port.Close()
}

func (s *Serial) Read(p []byte) (n int, err error) {
	return s.port.Read(p)
}

func (s *Serial) Write(p []byte) (n int, err error) {
	return s.port.Write(p)
}
