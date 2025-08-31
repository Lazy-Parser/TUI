package task

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// creation
type NewTimerTaskMsg struct { Id int }

// tick
type TimerTickMsg struct {
	Id int
	T  time.Duration
}

// Its a common (shared) task
type Timer struct {
	Id int
	start time.Time
}

func tick(data *Timer, ch chan<- tea.Msg) {
	for {
		time.Sleep(time.Second)
		ch <- TimerTickMsg{
			Id: data.Id,
			T:  time.Since(data.start).Round(time.Second),
		}
	}
}

func NewTimerTask(id int) Tasker {
	t := &Timer{start: time.Now()}
	return NewTask[Timer](id, "Main timer", t, tick)
}