package page_generator

import (
	component "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

// First step
func (m *mainView) handleEnter() (*mainView, tea.Cmd) {
	if m.start {
		return m, nil
	}

	m.start = true
	// start first step
	// TaskFlow first task is -1, so to start first task, we need to use [nextTask]
	cmd := m.taskFlow.NextTask()
	m.taskFlow.SetCurrentStatus(component.InProgress)

	return m, tea.Batch(m.logic.GetFutures(), cmd)
}

// When first step end. It also init second step
func (m *mainView) handleFuturesMsg(msg FuturesMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		// m.steps[0].err = msg.err
		// return m, nil
		// do not start next task if error
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.futures = msg.futures

	// start next task (dexscreener)
	cmd = m.logic.GetPairs(m.futures)
	cmds = append(cmds, cmd)

	// next task vizually
	m.taskFlow.SetCurrentStatus(component.Done)
	cmd = m.taskFlow.NextTask()
	m.taskFlow.SetCurrentStatus(component.InProgress)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *mainView) handleDsMsg(msg DexscreenerMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		// do smth
	}

	m.pairs = msg.pairs

	// move to the next task
	m.taskFlow.SetCurrentStatus(component.Done)
	cmd := m.taskFlow.NextTask()
	m.taskFlow.SetCurrentStatus(component.InProgress)

	return m, cmd
}

// save pairs / tokens to the database
func (m *mainView) handleTasksEnd() (tea.Model, tea.Cmd) {
	return nil, nil
}
