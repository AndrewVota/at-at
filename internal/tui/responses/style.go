package responses

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Focused   BaseStyle
	Unfocused BaseStyle
}

type BaseStyle struct {
	Base lipgloss.Style
}

func DefaultStyles() Style {
	var style = Style{
		Focused: BaseStyle{
			Base: lipgloss.NewStyle().Width(72).Height(2).Margin(1, 2),
		},
		Unfocused: BaseStyle{
			Base: lipgloss.NewStyle().Width(72).Height(2).Margin(1, 2),
		},
	}

	return style
}
