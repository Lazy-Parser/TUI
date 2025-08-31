package task

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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

type Task[T any] struct {
	id       int
	title    string
	started  time.Time
	finished time.Time

	// private
	data *T
	do   func(data *T, ch chan<- tea.Msg)
}

func (t *Task[T]) ID() int {
	return t.id
}

func (t *Task[T]) Title() string {
	return t.title
}

func (t *Task[T]) StartedAt() time.Time {
	return t.started
}

func (t *Task[T]) FinishedAt() time.Time {
	return t.finished
}

func (t *Task[T]) Finish() {
	t.finished = time.Now()
}

func (t *Task[T]) Run(ch chan<- tea.Msg) {
	if t.do != nil {
		t.do(t.data, ch)
	}
}

// TODO: implement!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (t *Task[T]) Cancel() {}

func NewTask[T any](id int, title string, instance *T, do func(data *T, ch chan<- tea.Msg)) *Task[T] {
	return &Task[T]{id: id, title: title, data: instance, do: do}
}
