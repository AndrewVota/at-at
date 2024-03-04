package tui

import (
	"github.com/andrewvota/at-at/internal/tui/commands"
	"github.com/andrewvota/at-at/internal/tui/devices"
	"github.com/andrewvota/at-at/internal/tui/messages"
	"github.com/andrewvota/at-at/internal/tui/ports"
)

type Model struct {
	// Global
	Keys         KeyMap
	WindowWidth  int
	WindowHeight int
	State        messages.State

	// Components
	PortsComponent    *ports.Model
	DevicesComponent  *devices.Model
	CommandsComponent *commands.Model
}

func New() *Model {
	return &Model{
		Keys:         DefaultKeyMap(),
		WindowWidth:  0,
		WindowHeight: 0,
		State:        messages.SelectingPort,

		PortsComponent:    ports.New(),
		DevicesComponent:  devices.New(),
		CommandsComponent: commands.New(),
	}
}
