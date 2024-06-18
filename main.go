package main

import (
	"github.com/andrewvota/at-at/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cmd.Execute()
}
