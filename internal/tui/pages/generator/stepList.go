package page_generator

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// messages
type AddStepMsg struct{ Title string }
type MakeStepDoneMsg struct{}
type TickLoadingMsg struct{} // used to make dots 'animation' in Loading...

// commands
// By default isLoading state in new Step is [True]
func (sl *StepList) AddStep(title string) (*StepList, tea.Cmd) {
	sl.steps = append(sl.steps, newStep(title))

	return sl, tick()
}

// Sets current step's state 'Done'
// 'Done' state is a state, when no error (err == nil) and 'isLoading' == false
func (sl *StepList) MakeStepDoneCmd() (*StepList, tea.Cmd) {
	sl.steps[sl.stepIdx].isLoading = false

	return sl, nil
}
func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return TickLoadingMsg{}
	})
}

type Step struct {
	title     string
	isLoading bool // default true
	err       error
	dots      int // from 0 to 3
}

type StepList struct {
	steps   []Step
	stepIdx int // current step. All previous steps have error or done state
}

func NewStepList() *StepList {
	return &StepList{}
}

func (pg *StepList) Init() tea.Cmd {
	return nil
}

func (pg *StepList) Update(msg tea.Msg) (*StepList, tea.Cmd) {
	switch msg := msg.(type) {

	case TickLoadingMsg:
		return pg.handleTick()

	case AddStepMsg:
		return pg.AddStep(msg.Title)

	case MakeStepDoneMsg:
		return pg.MakeStepDoneCmd()

	case tea.KeyMsg:
		switch msg.String() {
		// exit
		case "q":
			return pg, tea.Quit
		}
		switch msg.Type {
		case tea.KeyCtrlC:
			return pg, tea.Quit
		}
	}

	return pg, nil
}

func (pg *StepList) View() string {
	// ⏳ ✅ ❗
	var str string

	for _, step := range pg.steps {
		log.Printf("Title: %s", step.title)
		str += fmt.Sprintf("• %s:", step.title)

		if step.isStepLoading() {
			str += "\tLoading"
			// dots
			for range step.dots {
				str += "."
			}
			str += "\n"
		} else if step.isStepError() {
			str += "\t" + step.err.Error() + "\n"
		}
		str += "\n"
	}

	return str
}

// help methods
func (pg *StepList) handleTick() (*StepList, tea.Cmd) {
	var cmd tea.Cmd
	pg.steps[pg.stepIdx], cmd = pg.getCurrentStep().nextDot()
	return pg, cmd
}

func (sl *StepList) getCurrentStep() Step {
	return sl.steps[sl.stepIdx]
}

func (s Step) isStepDone() bool {
	return s.isLoading == false && s.err == nil
}

func (s Step) isStepError() bool {
	return s.err != nil
}

func (s Step) isStepLoading() bool {
	return s.isLoading
}

func (s Step) nextDot() (Step, tea.Cmd) {
	if s.dots == 3 {
		s.dots = 0
	} else {
		s.dots++
	}

	return s, tick()
}

func newStep(title string) Step {
	return Step{isLoading: true, title: title, dots: 0}
}
