package page_default

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// New page selection
type PageSelected struct{ Title string }

func SelectPage(title string) tea.Cmd {
	return func() tea.Msg {
		return PageSelected{Title: title}
	}
}

// terminal creation
type terminalResultMsg struct{ Err error }

func OpenNewTerminal() tea.Cmd {
	return func() tea.Msg {
		var err error
		wd, _ := os.Getwd()
		path := filepath.Join(wd, "logs", "debug.log")

		// if macos
		if strings.Contains(runtime.GOOS, "darwin") {
			command := `tail -f ` + path
			script := fmt.Sprintf(`tell application "Terminal" to do script %q`, command)
			err = exec.Command("osascript", "-e", script).Run()
		}

		if strings.Contains(runtime.GOOS, "linux") {
			// i use ptyxis
			// ptyxis --tab -x "tail -f logs/debug.log"
			command := `tail -f ` + path
			err = exec.Command("ptyxis", "--tab", "-x", command).Run()
		}

		return terminalResultMsg{Err: err}
	}
}
