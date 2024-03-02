package commands

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the app.
type KeyMap struct {
	Submit key.Binding
}

// DefaultKeyMap returns a set of default keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Submit: key.NewBinding(
			key.WithKeys("enter"),
		),
	}
}
