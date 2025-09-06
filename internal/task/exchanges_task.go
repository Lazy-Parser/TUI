package task

import (
	"github.com/Lazy-Parser/TUI/internal/logic"
	tea "github.com/charmbracelet/bubbletea"
)

type NewGenerationTaskMsg struct{ Exchanges []string }

func NewGenerationTask(logic *logic.Logic, exchanges []string) Tasker {
	return NewTask(3, "Generate for exchanges", func(ch chan<- tea.Msg) {

	})
}
