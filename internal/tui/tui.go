package tui

import (
	"fmt"

	"github.com/Lazy-Parser/TUI/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: add cancelation function for each task
func Run() error {
	f := StartLogger()

	manager := task.CreateTaskManager()
	p := tea.NewProgram(InitLayout(manager), tea.WithAltScreen())
	ch := runManager(p, manager)

	if _, err := p.Run(); err != nil {
		manager.Stop()
		close(ch)
		f.Close()
		return fmt.Errorf("Alas, there's been an error: %v", err)
	}

	manager.Stop()
	close(ch)
	f.Close()
	return nil
}

func runManager(p *tea.Program, manager *task.TaskManager) chan tea.Msg {
	ch := make(chan tea.Msg, 1024)

	// execute all
	go manager.Run(ch)
	go func() {
		for msg := range ch {
			p.Send(msg)
		}
	}()

	return ch
}
