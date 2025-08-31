package task

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type TaskManager struct {
	// list of all tasks (new and old)
	tasks []Tasker
	// listen for new task to exec
	tasksCh chan Tasker

	quit chan struct{}
}

func CreateTaskManager() *TaskManager {
	return &TaskManager{
		tasksCh: make(chan Tasker, 1024),
		quit:    make(chan struct{}, 1),
	}
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
func (manager *TaskManager) ConsumeMsg(msg tea.Msg) {
	switch msg := msg.(type) {
	case NewTimerTaskMsg:
		manager.Add(NewTimerTask(msg.Id))
	}
}
