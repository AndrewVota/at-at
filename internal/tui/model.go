package tui

import (
	"github.com/andrewvota/at-at/internal/tui/commands"
	"github.com/andrewvota/at-at/internal/tui/devices"
	"github.com/andrewvota/at-at/internal/tui/inputs"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/andrewvota/at-at/internal/tui/ports"
	"github.com/andrewvota/at-at/internal/tui/responses"
)

type Model struct {
	// Global
	Keys         KeyMap
	WindowWidth  int
	WindowHeight int
	State        messages.State

	// Components
	PortsComponent     *ports.Model
	DevicesComponent   *devices.Model
	CommandsComponent  *commands.Model
	ResponsesComponent *responses.Model
	InputsComponent    *inputs.Model
}

func New() *Model {
	return &Model{
		// Global
		Keys:         DefaultKeyMap(),
		WindowWidth:  0,
		WindowHeight: 0,

		// Components
		PortsComponent:     ports.New(),
		DevicesComponent:   devices.New(),
		CommandsComponent:  commands.New(),
		ResponsesComponent: responses.New(),
		InputsComponent:    inputs.New(),
	}
}
