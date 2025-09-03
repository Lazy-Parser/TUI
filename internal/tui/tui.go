package tui

import (
	"fmt"

	"github.com/Lazy-Parser/TUI/internal/service"
	"github.com/Lazy-Parser/TUI/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: add cancelation function for each task
func Run() error {
	f := StartLogger()

	manager, err := task.CreateTaskManager()
	if err != nil {
		return fmt.Errorf("failed to create task manager: %v", err)
	}
	service := service.NewService()
	p := tea.NewProgram(InitLayout(manager, service), tea.WithAltScreen())
	ch := setup(p, manager, service)

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

func setup(p *tea.Program, manager *task.TaskManager, service *service.Service) chan tea.Msg {
	ch := make(chan tea.Msg, 1024)

	service.SetMsgChannel(ch)
	go manager.Run(ch)

	go func() {
		for msg := range ch {
			p.Send(msg)
		}
	}()

	return ch
}
