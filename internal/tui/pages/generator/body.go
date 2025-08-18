package page_generator

import (
	"fmt"
	"math/rand/v2"
	"time"

	component "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct {
	progress1 tea.Model
	progress2 tea.Model
	progress3 tea.Model

	counter1 int
	counter2 int
	counter3 int
}

func NewMain() tea.Model {
	return &mainView{
		progress1: component.NewProgress(10, 0),
		progress2: component.NewProgress(10, 1),
		progress3: component.NewProgress(10, 2),
	}
}

func ticker(id int) tea.Cmd {
	return func() tea.Msg {
		// wait
		wait := time.Duration(random(0.1, 2.0) * float64(time.Second))
		time.Sleep(wait)

		return component.ProgressMsg{
			Increment: true,
			Id:        id,
		}
	}
}

func (m *mainView) Init() tea.Cmd {
	cmds := []tea.Cmd{ticker(0), ticker(1), ticker(2)}
	cmds = append(cmds, m.progress1.Init())
	cmds = append(cmds, m.progress2.Init())
	cmds = append(cmds, m.progress3.Init())

	return tea.Batch(cmds...)
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case component.ProgressMsg:

		// бля короче проблема в том что прогресс возвращает нил если происходит инкрементация и если не происходит
		m.progress1, _ = m.progress1.Update(msg)
		if msg.Id == 0 {
			if m.counter1 == 10 {
				return m, nil
			}
			m.counter1++
			return m, ticker(0)
		}
		m.progress2, _ = m.progress2.Update(msg)
		if msg.Id == 1 {
			if m.counter2 == 10 {
				return m, nil
			}
			m.counter2++
			return m, ticker(1)
		}
		m.progress3, _ = m.progress3.Update(msg)
		if msg.Id == 2 {
			if m.counter3 == 10 {
				return m, nil
			}
			m.counter3++
			return m, ticker(2)
		}

	case component.TickMsg:
		var (
			cmd  tea.Cmd
			cmds []tea.Cmd
		)

		m.progress1, cmd = m.progress1.Update(msg)
		cmds = append(cmds, cmd)
		m.progress2, cmd = m.progress2.Update(msg)
		cmds = append(cmds, cmd)
		m.progress3, cmd = m.progress3.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch msg.String() {
		// exit
		case "q":
			return m, tea.Quit
		}
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *mainView) View() string {
	return fmt.Sprintf("%s\n%s\n%s", m.progress1.View(), m.progress2.View(), m.progress3.View())
}

func random(min float64, max float64) float64 {
	return rand.Float64()*max - min + min
}
