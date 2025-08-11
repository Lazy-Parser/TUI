package tui

import (
	"fmt"
	"tui/internal/tui/components"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	f := StartLogger()

	p := tea.NewProgram(components.InitLayout(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		f.Close()
		return fmt.Errorf("Alas, there's been an error: %v", err)
	}

	f.Close()
	return nil
}
