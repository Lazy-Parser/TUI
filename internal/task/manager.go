package task

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/TUI/internal/logic"
	tea "github.com/charmbracelet/bubbletea"
)

type TaskManager struct {
	logic *logic.Logic

	// list of all tasks (new and old)
	tasks []Tasker
	// listen for new task to exec
	tasksCh chan Tasker

	quit chan struct{}
}

func CreateTaskManager() (*TaskManager, error) {
	l, err := logic.NewLogic()
	if err != nil {
		return nil, fmt.Errorf("failed to create logic service. reason: %v", err)
	}
	return &TaskManager{
		logic:   l,
		tasksCh: make(chan Tasker, 1024),
		quit:    make(chan struct{}, 1),
	}, nil
}

func (manager *TaskManager) Stop() {
	close(manager.quit)
	close(manager.tasksCh)
}

func (manager *TaskManager) Add(task Tasker) {
	// add to the list
	manager.tasks = append(manager.tasks, task)
	// push to the channel to exec
	manager.tasksCh <- task
}

func (manager *TaskManager) GetAllTasks() []Tasker {
	return manager.tasks
}

func (manager *TaskManager) Run(ch chan tea.Msg) {
	for {
		select {
		case t, ok := <-manager.tasksCh:
			if !ok {
				return // manager stopped
			}
			go func(task Tasker) {
				log.Printf("Start task '%s'", task.Title())
				task.Run(ch)
			}(t)
		case <-manager.quit:
			return
		}
	}
}

// creates a task for the corresponding message
func (manager *TaskManager) HandleMsg(msg tea.Msg) {
	switch msg := msg.(type) {
	case NewTimerTaskMsg:
		log.Println("Start timer task!")
		manager.Add(NewTimerTask(msg.Id))
	case NewFetchPoolTaskMsg:
		log.Println("Start fetch pool task!")
		manager.Add(NewFetchPoolTask(manager.logic, msg.Address, msg.Network))
	}
}
