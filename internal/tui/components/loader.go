package component

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Loader struct {
	dots    int
	text    string
	working bool
}

// {text}... .By default "Loading"
func NewLoader(text string) *Loader {
	return &Loader{dots: 0, text: text, working: true}
}

type incrementDotsMsg struct{}

func incrementDotsCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second)
		return incrementDotsMsg{}
	}
}

func (model *Loader) Init() tea.Cmd {
	return incrementDotsCmd()
}

func (model *Loader) Update(msg tea.Msg) (*Loader, tea.Cmd) {
	switch msg.(type) {
	case incrementDotsMsg:
		return model.handleIncrementDots()
	}

	return model, nil
}

func (model *Loader) View() string {
	var dotsStr string
	for range model.dots {
		dotsStr += "."
	}

	return model.text + dotsStr
}

// methods public
//
// closing loader animation
func (model *Loader) StopLoader() {
	model.working = false
}

// methods private
func (model *Loader) handleIncrementDots() (*Loader, tea.Cmd) {
	if model.dots == 3 {
		model.dots = 0
	} else {
		model.dots++
	}

	var cmd tea.Cmd
	if model.working {
		cmd = incrementDotsCmd()
	}

	return model, cmd
}
