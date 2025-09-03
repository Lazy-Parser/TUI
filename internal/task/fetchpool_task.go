package task

import (
	"github.com/Lazy-Parser/TUI/internal/logic"
	tea "github.com/charmbracelet/bubbletea"
)

type NewFetchPoolTaskMsg struct {
	Address string
	Network string
}

func NewFetchPoolTask(l *logic.Logic, address string, network string) Tasker {
	return NewTask(2, "Fetch pool", func(ch chan<- tea.Msg) {
		l.FetchPoolByToken(ch, address, network)
	})
}
