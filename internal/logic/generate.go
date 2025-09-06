package logic

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GenerationTaskResultMsg struct{ Err error }

// TODO: make progress in future
// type GenerationTaskProgresstMsg struct{ Exchanges []string }

func (core *Logic) Generation(ch chan<- tea.Msg, exchanges []string) {
	// ctx := context.Background()

	// TODO: now there is only Mexc, but i need to make smth for others exchanges
}
