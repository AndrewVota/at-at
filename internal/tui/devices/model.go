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

	CurrentDeviceMake  string
	CurrentDeviceModel string
	CurrentBaudeRate   int
	CurrentDataBits    int
	CurrentStopBits    float32
	CurrentParity      string
}

type Device struct {
	Make      string
	Model     string
	Type      string
	BaudeRate int
	DataBits  int
	StopBits  float32
	Parity    string
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
		items[i] = &Device{Make: d.Details.Make, Model: d.Details.Model, Type: d.Details.Type, BaudeRate: d.SerialSettings.BaudRate, DataBits: d.SerialSettings.DataBits, StopBits: d.SerialSettings.StopBits, Parity: d.SerialSettings.Parity}
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return &Model{
		WindowWidth:  0,
		WindowHeight: 0,

		List:    l,
		Focused: false,

		CurrentDeviceMake:  "",
		CurrentDeviceModel: "",
		CurrentBaudeRate:   0,
		CurrentDataBits:    0,
		CurrentStopBits:    0,
		CurrentParity:      "",
	}
}
