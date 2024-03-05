package commands

func (m *Model) View() string {
	styles := DefaultStyles()
	m.List.SetShowHelp(false)
	m.List.Title = "Commands"

	switch m.Focused {
	case true:
		m.List.SetWidth(m.WindowWidth)
		m.List.SetHeight(m.WindowHeight)
		m.List.Styles.Title = styles.Focused.Title
		m.List.Styles.TitleBar = styles.Focused.TitleBar
		m.List.SetDelegate(styles.Focused.ItemDelegate)
		return m.List.View()

	case false:
		m.List.SetWidth(m.WindowWidth)
		m.List.SetHeight(m.WindowHeight)
		m.List.Styles.Title = styles.Unfocused.Title
		m.List.Styles.TitleBar = styles.Unfocused.TitleBar
		m.List.SetDelegate(styles.Unfocused.ItemDelegate)
		return m.List.View()
	}

	return "commands component view..."
}
