package inputs

import (
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	Base lipgloss.Style
}

func DefaultStyles() Style {
	var style = Style{
		Base: lipgloss.NewStyle().Width(72).Margin(1, 2),
	}

	return style
}
