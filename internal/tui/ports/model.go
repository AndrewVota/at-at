package ports

import (
	"github.com/andrewvota/at-at/internal/serial"
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	WindowWidth  int
	WindowHeight int

	List    list.Model
	Focused bool
}

type Port struct {
	Name string
}

func (p *Port) Title() string       { return p.Name }
func (p *Port) Description() string { return "" }
func (p *Port) FilterValue() string { return p.Name }

func New() *Model {
	ports, err := serial.GetSerialPorts()
	if err != nil {
		panic(err)
	}

	items := make([]list.Item, len(ports))
	for i, port := range ports {
		items[i] = &Port{Name: port}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return &Model{
		List:    l,
		Focused: true,
	}
}
