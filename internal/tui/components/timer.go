package component

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct {
	id int
	t  time.Time
}

func tick(id int) tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg{t: t, id: id}
	})
}

type Timer struct {
	start   time.Time
	elapsed time.Duration // how many time passed
	id      int
}

func (t Timer) Init() tea.Cmd {
	return tick(t.id)
}

func (t Timer) Update(msg tea.Msg) (Timer, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		if msg.id != t.id {
			return t, nil
		}

		t.elapsed = time.Since(t.start).Round(time.Second)
		return t, tick(t.id)
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return t, tea.Quit
		}
	}

	return t, nil
}

func (t Timer) View() string {
	total := int(t.elapsed.Seconds())

	d := total / 86400
	h := (total % 86400) / 3600
	m := (total % 3600) / 60
	s := total % 60

	var strBuilder strings.Builder
	strBuilder.WriteString("Time: ")
	if d > 0 {
		strBuilder.WriteString(fmt.Sprintf("%dd, ", d))
	}
	if h > 0 {
		strBuilder.WriteString(fmt.Sprintf("%02dh, ", h))
	}
	strBuilder.WriteString(fmt.Sprintf("%02dm, ", m))
	strBuilder.WriteString(fmt.Sprintf("%02ds", s))
	strBuilder.WriteString(" ⏰")

	return strBuilder.String()
}

func NewTimer(id int) Timer {
	return Timer{
		start:   time.Now(),
		elapsed: 0,
		id:      id,
	}
}
