package page_default

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// New page selection
type PageSelected struct{ title string }

func SelectPage(title string) tea.Cmd {
	return func() tea.Msg {
		return PageSelected{title: title}
	}
}

// terminal creation
type terminalResultMsg struct{ Err error }

func OpenNewTerminal() tea.Cmd {
	return func() tea.Msg {
		// get logs path
		wd, _ := os.Getwd()
		path := filepath.Join(wd, "logs", "debug.log")

		command := `tail -f ` + path
		script := fmt.Sprintf(`tell application "Terminal" to do script %q`, command)
		err := exec.Command("osascript", "-e", script).Run()

		return terminalResultMsg{Err: err}
	}
}
