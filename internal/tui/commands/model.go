package commands

import (
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	WindowWidth  int
	WindowHeight int

	List    list.Model
	Focused bool

	CurrentPortName       string
	CurrentDeviceMake     string
	CurrentDeviceModel    string
	CurrentDeviceBaudRate int
	CurrentDeviceDataBits int
	CurrentDeviceStopBits float32
	CurrentDeviceParity   string
}

type Command struct {
	Command string
	Details string
}

func (c *Command) Title() string       { return c.Command }
func (c *Command) Description() string { return c.Details }
func (c *Command) FilterValue() string { return c.Command + " " + c.Details }

func New() *Model {
	items := []list.Item{}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return &Model{
		WindowWidth:  0,
		WindowHeight: 0,

		List:    l,
		Focused: false,

		CurrentPortName:       "",
		CurrentDeviceMake:     "",
		CurrentDeviceModel:    "",
		CurrentDeviceBaudRate: 0,
		CurrentDeviceDataBits: 0,
		CurrentDeviceStopBits: 0,
		CurrentDeviceParity:   "",
	}
}
