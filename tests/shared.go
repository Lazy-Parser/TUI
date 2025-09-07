package shared_test

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func StartLogger(path string) *os.File {
	f, err := tea.LogToFile(path, "test")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	return f
}

func OpenTerminalLogs(path string) {
	command := `echo -n -e "\033]0;LOGS\007"; tail -f ` + path
	script := fmt.Sprintf(`tell application "Terminal" to do script %q`, command)
	err := exec.Command("osascript", "-e", script).Run()
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}

func CloseTerminalLogs() {
	script := `
tell application "Terminal"
    repeat with w in windows
        if name of w contains "LOGS" then
            close w
            exit repeat
        end if
    end repeat
end tell
`
	err := exec.Command("osascript", "-e", script).Run()
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
}
