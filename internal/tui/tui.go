package tui

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/tui/components"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(tokenRepo market.TokenRepo) error {
	f := StartLogger()

	log.Println("Start programm")

	p := tea.NewProgram(components.InitLayout(tokenRepo), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		f.Close()
		return fmt.Errorf("Alas, there's been an error: %v", err)
	}

	f.Close()
	return nil
}
