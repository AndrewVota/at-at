package responses

func (m *Model) View() string {
	styles := DefaultStyles()
    m.TextArea.SetHeight(m.WindowHeight - 6)
	return styles.Focused.Base.Render(m.TextArea.View())
}
