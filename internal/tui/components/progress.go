package component

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ProgressMsg struct {
	Increment bool
	Id        int // because can be used multiple Progres models at the same time
}

type Progress struct {
	t       Timer
	current int
	max     int
	id      int
}

func (p Progress) Init() tea.Cmd {
	return p.t.Init()
}

func (p Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case ProgressMsg:
		// if not message for this component
		if p.id != msg.Id {
			return p, nil
		}

		// stop
		if p.max == p.current {
			return p, nil
		}

		return p.increment(), nil

	case tea.KeyMsg:
		switch msg.String() {
		// exit
		case "q":
			return p, tea.Quit
		}
		switch msg.Type {
		case tea.KeyCtrlC:
			return p, tea.Quit
		}
	}

	// if stopped, do not continue timer
	var cmd tea.Cmd
	if p.max != p.current {
		p.t, cmd = p.t.Update(msg)
	}

	return p, cmd
}

func (p Progress) View() string {
	// build string
	var str string

	if p.current < p.max {
		str += "["
		percent := int(float32(p.current) / float32(p.max) * 100.0) // ignore floating

		// progress bar has 10 segments: [1|2|3|4|5|6|7|8|9|10]
		// set '>' at the current segment

		segmentIdx := percent / 10

		for range segmentIdx - 1 {
			// fill all segments before '>' with '='
			str += "="
		}
		str += ">"
		for range 10 - segmentIdx {
			// fill all segments after '>' with space
			str += " "
		}
		str += "]"
	} else {
		str += "Done!"
		// stop timer.
	}

	str += " | "

	str += fmt.Sprintf("%d/%d", p.current, p.max)

	str += " | "

	str += p.t.View()

	return str
}

func (p Progress) increment() tea.Model {
	p.current++
	return p
}

// Important! Timer and Progress will have the same [id].
func NewProgress(max int, id int) tea.Model {
	return &Progress{
		max:     max,
		current: 0,
		id:      id,
		t:       NewTimer(id),
	}
}
