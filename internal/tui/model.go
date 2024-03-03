package tui

import "github.com/andrewvota/at-at/internal/tui/ports"

type Model struct {
	// Global
	Keys         KeyMap
	WindowWidth  int
	WindowHeight int
	// State messages.State

	// Components
	PortsComponent *ports.Model
}

func New() *Model {
	return &Model{
		Keys:         DefaultKeyMap(),
		WindowWidth:  0,
		WindowHeight: 0,

		PortsComponent: ports.New(),
	}
}
