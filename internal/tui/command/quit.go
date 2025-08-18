package command

import (
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func clearLogs() tea.Cmd {
	return func() tea.Msg {
		wd, _ := os.Getwd()
		path := filepath.Join(wd, "logs", "debug.log")

		// clear before using. TODO: clear logs on closing, not opening logs.
		if err := os.Truncate(path, 0); err != nil {
			log.Printf("Failed to clear logs: %s", path)
		}

		return nil
	}
}

func OnQuit() tea.Cmd {
	return tea.Sequence(clearLogs(), tea.Quit)
}
