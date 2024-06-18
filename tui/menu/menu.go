package menu

import (
	"github.com/andrewvota/at-at/serial"
	"github.com/andrewvota/at-at/tui/messages"
	"github.com/andrewvota/at-at/tui/selector"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	FocusNextSelector key.Binding
	FocusPrevSelector key.Binding
	Submit            key.Binding
}

var DefaultKeyMap = KeyMap{
	FocusNextSelector: key.NewBinding(key.WithKeys("tab", "down", "j")),
	FocusPrevSelector: key.NewBinding(key.WithKeys("shift+tab", "up", "k")),
	Submit:            key.NewBinding(key.WithKeys("enter")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool
	width  int
	height int

	// State
	activeSelector ActiveSelectors

	// Components
	nameSelector     selector.Model
	baudSelector     selector.Model
	paritySelector   selector.Model
	dataBitsSelector selector.Model
	stopBitsSelector selector.Model
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  true,

		activeSelector: NameSelectorActive,

		nameSelector:     selector.New(),
		baudSelector:     selector.New(),
		paritySelector:   selector.New(),
		dataBitsSelector: selector.New(),
		stopBitsSelector: selector.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return loadInitialData
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case InitMsg:
		m.nameSelector.AddChoices(msg.Ports...)
		m.baudSelector.AddChoices(msg.Bauds...)
		m.paritySelector.AddChoices(msg.Parity...)
		m.dataBitsSelector.AddChoices(msg.DataBits...)
		m.stopBitsSelector.AddChoices(msg.StopBits...)
		m.nameSelector.Focus()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.FocusNextSelector):
			m.activeSelector = (m.activeSelector + 1) % 5
		case key.Matches(msg, m.KeyMap.FocusPrevSelector):
			m.activeSelector = (m.activeSelector + 4) % 5
		case key.Matches(msg, m.KeyMap.Submit):
			return m, tea.Batch(m.sendSerialConfigMsg, messages.ChangeStateTo(messages.StateRepl))
		}
	}

	switch m.activeSelector {
	case NameSelectorActive:
		m.BlurAllComponents()
		m.nameSelector.Focus()
		m.nameSelector, _ = m.nameSelector.Update(msg)
	case BaudSelectorActive:
		m.BlurAllComponents()
		m.baudSelector.Focus()
		m.baudSelector, _ = m.baudSelector.Update(msg)
	case ParitySelectorActive:
		m.BlurAllComponents()
		m.paritySelector.Focus()
		m.paritySelector, _ = m.paritySelector.Update(msg)
	case DataBitsSelectorActive:
		m.BlurAllComponents()
		m.dataBitsSelector.Focus()
		m.dataBitsSelector, _ = m.dataBitsSelector.Update(msg)
	case StopBitsSelectorActive:
		m.BlurAllComponents()
		m.stopBitsSelector.Focus()
		m.stopBitsSelector, _ = m.stopBitsSelector.Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var view = m.nameSelector.View() + "\n" + m.baudSelector.View() + "\n" + m.paritySelector.View() + "\n" + m.dataBitsSelector.View() + "\n" + m.stopBitsSelector.View()

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view)
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

type InitMsg struct {
	Ports    []string
	Bauds    []string
	Parity   []string
	DataBits []string
	StopBits []string
}

func loadInitialData() tea.Msg {
	ports, err := serial.GetPortsList()
	if err != nil {
		ports = []string{"NO SERIAL PORTS FOUND"}
	}

	var bauds = []string{"4800 Baud", "9600 Baud", "19200 Baud", "38400 Baud", "57600 Baud", "115200 Baud", "230400 Baud", "460800 Baud", "921600 Baud"}
	var parity = []string{"No Parity", "Odd Parity", "Even Parity", "Mark Parity", "Space Parity"}
	var dataBits = []string{"5 Data Bits", "6 Data Bits", "7 Data Bits", "8 Data Bits"}
	var stopBits = []string{"1 Stop Bit", "1.5 Stop Bits", "2 Stop Bits"}

	return InitMsg{
		Ports:    ports,
		Bauds:    bauds,
		Parity:   parity,
		DataBits: dataBits,
		StopBits: stopBits,
	}

}

type ActiveSelectors int

const (
	NameSelectorActive ActiveSelectors = iota
	BaudSelectorActive
	ParitySelectorActive
	DataBitsSelectorActive
	StopBitsSelectorActive
)

func (m *Model) BlurAllComponents() {
	m.nameSelector.Blur()
	m.baudSelector.Blur()
	m.paritySelector.Blur()
	m.dataBitsSelector.Blur()
	m.stopBitsSelector.Blur()
}

type SerialConfigMsg struct {
	PortName string
	BuadRate string
	Parity   string
	DataBits string
	StopBits string
}

func (m *Model) sendSerialConfigMsg() tea.Msg {
	return SerialConfigMsg{
		PortName: m.nameSelector.Value(),
		BuadRate: m.baudSelector.Value(),
		Parity:   m.paritySelector.Value(),
		DataBits: m.dataBitsSelector.Value(),
		StopBits: m.stopBitsSelector.Value(),
	}
}
