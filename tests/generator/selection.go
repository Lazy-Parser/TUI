package main

import (
	"log"
	"os"
	"path/filepath"

	page_generator "github.com/Lazy-Parser/TUI/internal/tui/pages/generator"
	shared_test "github.com/Lazy-Parser/TUI/tests"
	tea "github.com/charmbracelet/bubbletea"
)

func initLogs1() *os.File {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "..", "test.log")
	f := shared_test.StartLogger(path)
	shared_test.OpenTerminalLogs(path)

	return f
}

func main1() {
	f := initLogs1()

	model := page_generator.NewSelection(
	// page_generator.NewOption("Mexc", "Mexc"),
	// page_generator.NewOption("Bitget", "Bitget"),
	// page_generator.NewOption("Kukoin", "Kukoin"),
	// page_generator.NewOption("Gate.io", "Gate.io"),
	)
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		f.Close()
		shared_test.CloseTerminalLogs()
		log.Printf("Alas, there's been an error: %v", err)
	}

	shared_test.CloseTerminalLogs()
	f.Close()
}
