package tui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Lazy-Parser/Collector/database"
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/service"
	"github.com/Lazy-Parser/TUI/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: add cancelation function for each task
func Run() error {
	f := StartLogger()

	tokenRepo, poolRepo, err := startDatabase()
	if err != nil {
		return err
	}
	tables, err := poolRepo.GetAllRepos()
	if err != nil {
		return err
	}
	log.Printf("Tables: %+v", tables)

	manager, err := task.CreateTaskManager()
	if err != nil {
		return fmt.Errorf("failed to create task manager: %v", err)
	}
	service := service.NewService(tokenRepo, poolRepo)
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

func startDatabase() (market.TokenRepo, market.PoolRepo, error) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "storage", "storage.db")
	db, err := database.Start(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start database: %v", err)
	}

	tokenRepo := database.NewTokenRepo(db)
	poolRepo := database.NewPoolRepo(db)

	return tokenRepo, poolRepo, nil
}
