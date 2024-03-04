package commands

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	Focused   BaseStyle
	Unfocused BaseStyle
}

type BaseStyle struct {
	Base                  lipgloss.Style
	Title                 lipgloss.Style
	TitleBar              lipgloss.Style
	Selected              lipgloss.Style
	SelectedDescription   lipgloss.Style
	Unselected            lipgloss.Style
	UnselectedDescription lipgloss.Style
	ItemDelegate          list.ItemDelegate
}

func DefaultStyles() Style {
	var style = Style{
		Focused: BaseStyle{
			Base:                  lipgloss.NewStyle().MaxWidth(72).Margin(1, 2),
			Title:                 lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF")),
			TitleBar:              lipgloss.NewStyle().Width(70).Margin(0, 1, 1, 1).Background(lipgloss.Color("#64708D")),
			Selected:              lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			SelectedDescription:   lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			Unselected:            lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("241")),
			UnselectedDescription: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("241")),
			ItemDelegate:          nil,
		},
		Unfocused: BaseStyle{
			Base:                  lipgloss.NewStyle().MaxWidth(72).Margin(1, 2),
			Title:                 lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF")),
			TitleBar:              lipgloss.NewStyle().Width(70).Margin(0, 1, 1, 1).Background(lipgloss.Color("#373B41")),
			Selected:              lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			SelectedDescription:   lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#AFBEE1")),
			Unselected:            lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("237")),
			UnselectedDescription: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("237")),
			ItemDelegate:          nil,
		},
	}

	// Define the styles for the focused item delegate
	var focusedItemDelegate = list.NewDefaultDelegate()
	focusedItemDelegate.ShowDescription = true
	focusedItemDelegate.Styles.NormalTitle = style.Focused.Unselected
	focusedItemDelegate.Styles.NormalDesc = style.Focused.UnselectedDescription
	focusedItemDelegate.Styles.SelectedTitle = style.Focused.Selected
	focusedItemDelegate.Styles.SelectedDesc = style.Focused.SelectedDescription

	// Define the styles for the unfocused item delegate
	var unfocusedItemDelegate = list.NewDefaultDelegate()
	unfocusedItemDelegate.ShowDescription = true
	unfocusedItemDelegate.Styles.NormalTitle = style.Unfocused.Unselected
	unfocusedItemDelegate.Styles.NormalDesc = style.Unfocused.UnselectedDescription
	unfocusedItemDelegate.Styles.SelectedTitle = style.Unfocused.Selected
	unfocusedItemDelegate.Styles.SelectedDesc = style.Unfocused.SelectedDescription

	// Set the item delegate for the focused and unfocused styles
	style.Focused.ItemDelegate = focusedItemDelegate
	style.Unfocused.ItemDelegate = unfocusedItemDelegate

	return style
}
