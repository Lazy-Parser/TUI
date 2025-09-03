// this package imports 'Collector' lib and create logic of this app.
package logic

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TimerTickMsg struct {
	Id int
	T  time.Duration
}

type Timer struct {
	Id    int
	Start time.Time
}

func NewTimer(id int, start time.Time) *Timer {
	return &Timer{Id: id, Start: start}
}

func (t *Timer) Run(ch chan<- tea.Msg) {
	for {
		time.Sleep(time.Second)
		ch <- TimerTickMsg{
			Id: t.Id,
			T:  time.Since(t.Start).Round(time.Second),
		}
	}
}
