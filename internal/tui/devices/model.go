package devices

import (
	"github.com/andrewvota/at-at/internal/device"
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	WindowWidth  int
	WindowHeight int

	List    list.Model
	Focused bool
}

type Device struct {
	Make     string
	Model    string
	Type     string
    BaudRate int
    DataBits int
    StopBits float32
    Parity   string
	Commands []Command
}

type Command struct {
	Command string
	Details string
}

func (p *Device) Title() string       { return p.Make + " " + p.Model }
func (p *Device) Description() string { return p.Type }
func (p *Device) FilterValue() string { return p.Make + " " + p.Model }

func New() *Model {
	devices, err := device.LoadDeviceConfigs()
	if err != nil {
		panic(err)
	}

	items := make([]list.Item, len(devices))
	for i, d := range devices {
		Commands := make([]Command, len(d.Commands))
		for j, c := range d.Commands {
			Commands[j] = Command{Command: c.Command, Details: c.Details}
		}
		items[i] = &Device{
            Make:     d.Details.Make,
            Model:    d.Details.Model,
            Type:     d.Details.Type,
            BaudRate: d.SerialSettings.BaudRate,
            DataBits: d.SerialSettings.DataBits,
            StopBits: d.SerialSettings.StopBits,
            Parity:   d.SerialSettings.Parity,
            Commands: Commands,
        }
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Devices"

	return &Model{
		List:    l,
		Focused: false,
	}
}
