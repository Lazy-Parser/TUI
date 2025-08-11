package tui

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func StartLogger() *os.File {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "logs", "debug.log")
	f, err := tea.LogToFile(path, "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	return f
}
