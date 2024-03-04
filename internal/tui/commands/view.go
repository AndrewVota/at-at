package commands

func (m *Model) View() string {
	styles := DefaultStyles()
	m.List.SetShowHelp(false)

	switch m.Focused {
	case true:
		m.List.SetWidth(m.WindowWidth - styles.Focused.Base.GetHorizontalFrameSize())
		m.List.SetHeight(m.WindowHeight - styles.Focused.Base.GetVerticalFrameSize())
		m.List.Styles.Title = styles.Focused.Title
		m.List.Styles.TitleBar = styles.Focused.TitleBar
		m.List.SetDelegate(styles.Focused.ItemDelegate)
		return styles.Focused.Base.Render(m.List.View())

	case false:
		m.List.SetWidth(m.WindowWidth - styles.Unfocused.Base.GetHorizontalFrameSize())
		m.List.SetHeight(m.WindowHeight - styles.Unfocused.Base.GetVerticalFrameSize())
		m.List.Styles.Title = styles.Unfocused.Title
		m.List.Styles.TitleBar = styles.Unfocused.TitleBar
		m.List.SetDelegate(styles.Unfocused.ItemDelegate)
		return styles.Unfocused.Base.Render(m.List.View())
	}

	return "Device component view..."
}
