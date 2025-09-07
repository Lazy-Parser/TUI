package main

import (
	"log"
	"os"
	"path/filepath"

	page_generator "github.com/Lazy-Parser/TUI/internal/tui/pages/generator"
	shared_test "github.com/Lazy-Parser/TUI/tests"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	border = lipgloss.
		NewStyle().
		Border(lipgloss.NormalBorder(), true).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
)

func initLogs() *os.File {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "..", "test.log")
	f := shared_test.StartLogger(path)
	shared_test.OpenTerminalLogs(path)

	return f
}

func main() {
	f := initLogs()

	model := page_generator.NewAddPool()
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		f.Close()
		shared_test.CloseTerminalLogs()
		log.Printf("Alas, there's been an error: %v", err)
	}

	shared_test.CloseTerminalLogs()
	f.Close()
}
