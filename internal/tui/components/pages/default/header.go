package page_default

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"github.com/Lazy-Parser/TUI/internal/tui/components/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v4/cpu"
)

var ()

type modelHeader struct {
	cpuInfo  []cpu.InfoStat
	os       string
	timer    modelTimer
	someInfo string
	logo     string
	width    int
	height   int
}

func (m modelHeader) Init() tea.Cmd {
	return m.timer.Init()
}

func (m modelHeader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for borders (1 for left and 1 for right)
		m.height = msg.Height
	}

	updatedModel, cmd := m.timer.Update(msg)
	a, _ := updatedModel.(modelTimer)
	m.timer = a

	return m, cmd
}

func (m modelHeader) View() string {
	left := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Foreground(theme.MintColor).
		Bold(true).
		Render(m.logo)

	var rightStr strings.Builder
	rightStr.WriteString(fmt.Sprintf("Operation system: %s\n", m.os))
	rightStr.WriteString(fmt.Sprintf("CPU: %s  Cores: %d\n", m.cpuInfo[0].ModelName, m.cpuInfo[0].Cores))
	rightStr.WriteString(m.timer.View())
	rightStr.WriteString("\n" + m.someInfo)
	right := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Render(rightStr.String())

	gap := lipgloss.NewStyle().
		Width(m.width - lipgloss.Width(left) - lipgloss.Width(right)).
		Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, left, gap, right)
}

func newHeader() *modelHeader {
	logo := `__/\\\\\\\\\\\\\\\__/\\\________/\\\__/\\\\\\\\\\\_
 _\///////\\\/////__\/\\\_______\/\\\_\/////\\\///__
  _______\/\\\_______\/\\\_______\/\\\_____\/\\\_____
   _______\/\\\_______\/\\\_______\/\\\_____\/\\\_____
    _______\/\\\_______\/\\\_______\/\\\_____\/\\\_____
     _______\/\\\_______\/\\\_______\/\\\_____\/\\\_____
      _______\/\\\_______\//\\\______/\\\______\/\\\_____
       _______\/\\\________\///\\\\\\\\\/____/\\\\\\\\\\\_
        _______\///___________\/////////_____\///////////__`

	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("Failed to get cpu info: %v", err)
	}

	return &modelHeader{
		logo:    logo,
		os:      runtime.GOOS,
		timer:   newTimer(),
		cpuInfo: cpuInfo,
	}
}
