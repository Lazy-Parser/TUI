package task

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// требования для задач
// время выполнения задачи
// задача выполнена / выполняеться / ошибка
//

type Tasker interface {
	ID() int
	Title() string
	StartedAt() time.Time
	FinishedAt() time.Time
	// call in 'Run' at the end
	Finish()
	Run(ch chan<- tea.Msg)
	Cancel()
	// Status() - running / error / done
}

type Task struct {
	id       int
	title    string
	started  time.Time
	finished time.Time

	// private
	do func(ch chan<- tea.Msg)
}

func (t *Task) ID() int {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) StartedAt() time.Time {
	return t.started
}

func (t *Task) FinishedAt() time.Time {
	return t.finished
}

func (t *Task) Finish() {
	t.finished = time.Now()
}

func (t *Task) Run(ch chan<- tea.Msg) {
	if t.do != nil {
		t.do(ch)
	}
}

// TODO: implement!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (t *Task) Cancel() {}

func NewTask(id int, title string, do func(ch chan<- tea.Msg)) *Task {
	return &Task{id: id, title: title, do: do}
}
