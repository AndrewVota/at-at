package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the app.
type KeyMap struct {
	Quit              key.Binding
	ToggleActiveLeft  key.Binding
	ToggleActiveRight key.Binding
}

// DefaultKeyMap returns a set of default keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
		),
		ToggleActiveLeft: key.NewBinding(
			key.WithKeys("left", "shift+tab"),
		),
		ToggleActiveRight: key.NewBinding(
			key.WithKeys("right", "tab"),
		),
	}
}
