package inputs

func (m *Model) View() string {
	styles := DefaultStyles()

	switch m.Spinning {
	case true:
		return styles.Base.Render(m.Spinner.View() + "Sending command and awaiting response...")

	case false:
		return styles.Base.Render(m.TextInput.View())
	}

	return "Inputs component view..."
}
