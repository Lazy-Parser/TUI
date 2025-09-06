package page_default

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/Lazy-Parser/TUI/internal/logic"
	"github.com/Lazy-Parser/TUI/internal/task"
	"github.com/Lazy-Parser/TUI/internal/tui/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v4/cpu"
)

// TODO: Very important!!! This page updates always. So when other page selected, it update ticker and make update all layout and selected page too. Solve it, to not rerender when not selected.
// TODO: when i will make the server, this app will be like an admin panel. Here, in header, i want to add a string like "Server health check: Running"
type modelHeader struct {
	cpuInfo  []cpu.InfoStat
	os       string
	time     time.Duration
	someInfo string
	logo     string
	width    int
	height   int
}

func (m *modelHeader) Init() tea.Cmd {
	// send msg to create a timer task
	return common.CmdHandler(task.NewTimerTaskMsg{Id: 0})
}

func (m *modelHeader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case logic.TimerTickMsg: // from the background task
		if msg.Id == 0 {
			m.time = msg.T
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for borders (1 for left and 1 for right)
		m.height = msg.Height
	}

	// m.timer, cmd = m.timer.Update(msg)
	return m, cmd
}

func (m *modelHeader) View() string {
	left := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#000")). // theme.MintColor
		Bold(true).
		Render(m.logo)

	var rightStr strings.Builder
	rightStr.WriteString(fmt.Sprintf("Operation system: %s\n", m.os))
	rightStr.WriteString(fmt.Sprintf("CPU: %s  Cores: %d\n", m.cpuInfo[0].ModelName, m.cpuInfo[0].Cores))
	rightStr.WriteString(timeToString(&m.time))
	rightStr.WriteString("\n" + m.someInfo)
	right := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Render(rightStr.String())

	gap := lipgloss.NewStyle().
		Width(m.width - lipgloss.Width(left) - lipgloss.Width(right)).
		Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, left, gap, right)
}

func timeToString(t *time.Duration) string {
	total := int(t.Seconds())

	d := total / 86400
	h := (total % 86400) / 3600
	m := (total % 3600) / 60
	s := total % 60

	var strBuilder strings.Builder
	strBuilder.WriteString("Running: ")
	if d > 0 {
		strBuilder.WriteString(fmt.Sprintf("%dd, ", d))
	}
	if h > 0 {
		strBuilder.WriteString(fmt.Sprintf("%02dh, ", h))
	}
	strBuilder.WriteString(fmt.Sprintf("%02dm, ", m))
	strBuilder.WriteString(fmt.Sprintf("%02ds", s))

	return strBuilder.String()
}

func newHeader() *modelHeader {
	logo := `
   █████████   ███████████   █████   ████ █████   █████   █████████   ██████   ██████
  ███░░░░░███ ░░███░░░░░███ ░░███   ███░ ░░███   ░░███   ███░░░░░███ ░░██████ ██████
 ░███    ░███  ░███    ░███  ░███  ███    ░███    ░███  ░███    ░███  ░███░█████░███
 ░███████████  ░██████████   ░███████     ░███████████  ░███████████  ░███░░███ ░███
 ░███░░░░░███  ░███░░░░░███  ░███░░███    ░███░░░░░███  ░███░░░░░███  ░███ ░░░  ░███
 ░███    ░███  ░███    ░███  ░███ ░░███   ░███    ░███  ░███    ░███  ░███      ░███
 █████   █████ █████   █████ █████ ░░████ █████   █████ █████   █████ █████     █████
░░░░░   ░░░░░ ░░░░░   ░░░░░ ░░░░░   ░░░░ ░░░░░   ░░░░░ ░░░░░   ░░░░░ ░░░░░     ░░░░░`

	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("Failed to get cpu info: %v", err)
	}

	return &modelHeader{
		logo:    logo,
		os:      runtime.GOOS,
		cpuInfo: cpuInfo,
	}
}
