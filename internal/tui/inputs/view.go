package inputs

func (m *Model) View() string {
	styles := DefaultStyles()

	return styles.Focused.Base.Render(m.TextInput.View())
}
