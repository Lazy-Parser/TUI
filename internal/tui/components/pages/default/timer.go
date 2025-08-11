package page_default

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type modelTimer struct {
	start   time.Time
	elapsed time.Duration // how many time passed
}

func (t modelTimer) Init() tea.Cmd {
	return tick()
}

func (t modelTimer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		t.elapsed = time.Since(t.start).Round(time.Second)
		return t, tick()
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return t, tea.Quit
		}
	}

	return t, nil
}

func (t modelTimer) View() string {
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

func newTimer() modelTimer {
	return modelTimer{
		start:   time.Now(),
		elapsed: 0,
	}
}
