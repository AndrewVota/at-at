package ports

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	Focused   BaseStyle
	Unfocused BaseStyle
}

type BaseStyle struct {
	Base         lipgloss.Style
	Title        lipgloss.Style
	TitleBar     lipgloss.Style
	Selected     lipgloss.Style
	Unselected   lipgloss.Style
	ItemDelegate list.ItemDelegate
}

func DefaultStyles() Style {
	var style = Style{
		Focused: BaseStyle{
			Base:         lipgloss.NewStyle().Width(22).Margin(1, 2),
			Title:        lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF")),
			TitleBar:     lipgloss.NewStyle().Width(20).Margin(0, 1, 1, 1).Background(lipgloss.Color("#64708D")),
			Selected:     lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			Unselected:   lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("241")),
			ItemDelegate: nil,
		},
		Unfocused: BaseStyle{
			Base:         lipgloss.NewStyle().Width(22).Margin(1, 2),
			Title:        lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF")),
			TitleBar:     lipgloss.NewStyle().Width(20).Margin(0, 1, 1, 1).Background(lipgloss.Color("#373B41")),
			Selected:     lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			Unselected:   lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("237")),
			ItemDelegate: nil,
		},
	}

	// Define the styles for the focused item delegate
	var focsuedItemDelegate = list.NewDefaultDelegate()
	focsuedItemDelegate.ShowDescription = false
	focsuedItemDelegate.Styles.NormalTitle = style.Focused.Unselected
	focsuedItemDelegate.Styles.SelectedTitle = style.Focused.Selected

	// Define the styles for the unfocused item delegate
	var unfocusedItemDelegate = list.NewDefaultDelegate()
	unfocusedItemDelegate.ShowDescription = false
	unfocusedItemDelegate.Styles.NormalTitle = style.Unfocused.Unselected
	unfocusedItemDelegate.Styles.SelectedTitle = style.Unfocused.Selected

	// Set the item delegate for the focused and unfocused styles
	style.Focused.ItemDelegate = focsuedItemDelegate
	style.Unfocused.ItemDelegate = unfocusedItemDelegate

	return style
}
