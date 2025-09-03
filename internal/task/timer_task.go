package task

import (
	"time"

	"github.com/Lazy-Parser/TUI/internal/logic"
	tea "github.com/charmbracelet/bubbletea"
)

// creation
type NewTimerTaskMsg struct{ Id int }

func NewTimerTask(id int) Tasker {
	t := logic.NewTimer(id, time.Now())
	return NewTask(id, "Main timer", func(ch chan<- tea.Msg) {
		t.Run(ch)
	})
}
