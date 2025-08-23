package component

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	bold = lipgloss.NewStyle().Bold(true)
)

// Messages. TODO: error
// Move to the next task. If there is no next task, then all tasks are done
type NextTaskMsg struct{}

// Set status for the current task
type TaskStatusMsg struct{ Status TaskStatus }
type TaskTimerTickMsg struct {
	id   int
	time time.Time
}
// msg, that signals end of all tasks
type TaskEndMsg struct {}

// Commands
// Command for timer
func taskTimerTick(id int) tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TaskTimerTickMsg{
			id:   id,
			time: t,
		}
	})

}

// Task
type TaskStatus int

const (
	InProgress TaskStatus = iota // 0
	Done                         // 1
	Error                        // 2
	Waiting                      // 3. waiting for the previous task
)

type Task struct {
	Status      TaskStatus
	Err         error
	Title       string
	Description string
	Id          int

	// timer
	startTime time.Time
	elapsed   time.Duration
}

func NewTask(title string, description string) Task {
	return Task{
		Status:      Waiting,
		Err:         nil,
		Title:       title,
		Description: description,
	}
}

func (t *Task) UpdateTimer(id int) tea.Cmd {
	// if running
	if t.Status == InProgress {
		t.elapsed = time.Since(t.startTime).Round(time.Second)
		return taskTimerTick(id)
	}
	return nil
}

func (t Task) GetElapsedTime() string {
	var duration time.Duration
	if t.Status == InProgress { // is running
		duration = time.Since(t.startTime)
	} else {
		duration = t.elapsed
	}

	seconds := int(duration.Seconds())
	minutes := seconds / 60
	if seconds < 60 {
		seconds = seconds % 60
	}

	return fmt.Sprintf("%dm %ds", minutes, seconds)
}

// ● = reached
// ○ = not started
// ✔ = done
// ⟳ = in progress
func (t Task) RenderTask(i int) string {
	var str, circle, status, time string

	// circle + status
	switch t.Status {
	case InProgress:
		circle = "●"
		status = "⟳ In progress"
		time = fmt.Sprintf("(%s elapsed)", t.GetElapsedTime())
	case Done:
		circle = "●"
		status = "✔ Completed"
		time = fmt.Sprintf("(took %s)", t.GetElapsedTime())
	case Waiting:
		circle = "○"
		status = "... waiting"
		time = ""
	}

	// main info
	str += fmt.Sprintf("%s Task %d: %s\t%s %s", circle, i, t.Title, status, time)

	return str
}

// TaskFlow
type TaskFlow struct {
	tasks []Task
	note  string // set note when all tasks end or some error occure
	idx   int    // current task
}

func (p TaskFlow) Init() tea.Cmd {
	return nil
}

func (p *TaskFlow) Update(msg tea.Msg) (*TaskFlow, tea.Cmd) {
	switch msg := msg.(type) {

	case TaskTimerTickMsg:
		return p.handleTimerMsg(msg)

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

	case NextTaskMsg:
		// if no next
		if p.idx == len(p.tasks)-1 {
			p.note = "Complete succesful!"
		} else {
			p.idx++
		}

		// stop timer for current task ...
		return p, nil

	case TaskStatusMsg:
		p.tasks[p.idx].Status = msg.Status
		return p, nil
	}

	return p, nil
}

func (p TaskFlow) View() string {
	var str string

	for i, t := range p.tasks {
		str += t.RenderTask(i + 1)

		if i != len(p.tasks)-1 {
			str += "\n│\n"
		}
	}

	if p.note != "" {
		str += "\n"
		str += bold.Render(p.note)
		str += "\n"
	}

	return str
}

// methods
// Important! Each [Id] in [Task] must be unique!
func NewTaskFlow(tasks ...Task) *TaskFlow {
	return &TaskFlow{tasks: tasks, idx: -1}
}

func (tf *TaskFlow) handleTimerMsg(msg TaskTimerTickMsg) (*TaskFlow, tea.Cmd) {
	return tf, tf.tasks[msg.id].UpdateTimer(msg.id)
}

// Public methods to change state
func (tf *TaskFlow) SetCurrentStatus(status TaskStatus) {
	tf.tasks[tf.idx].Status = status
}

func (tf *TaskFlow) NextTask() tea.Cmd {
	// if no next tasks
	if tf.idx == len(tf.tasks)-1 {
		tf.note = "Complete succesful!"
		return sendTaskEndMsg()
	}

	// move to the next task
	tf.idx++
	tf.tasks[tf.idx].startTime = time.Now()
	// start timer for this task
	return taskTimerTick(tf.idx)
}

func sendTaskEndMsg() tea.Cmd {
	return func() tea.Msg {
		return TaskEndMsg{}
	}
}