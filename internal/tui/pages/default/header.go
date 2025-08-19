package page_default

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	component "github.com/Lazy-Parser/TUI/internal/tui/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v4/cpu"
)

// TODO: Very important!!! This page updates always. So when other page selected, it update ticker and make update all layout and selected page too. Solve it, to not rerender when not selected.

type modelHeader struct {
	cpuInfo  []cpu.InfoStat
	os       string
	timer    component.Timer
	someInfo string
	logo     string
	width    int
	height   int
}

func (m modelHeader) Init() tea.Cmd {
	return m.timer.Init()
}

func (m modelHeader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for borders (1 for left and 1 for right)
		m.height = msg.Height
	}

	m.timer, cmd = m.timer.Update(msg)
	return m, cmd
}

func (m modelHeader) View() string {
	left := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#000")). // theme.MintColor
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
		logo:     logo,
		os:       runtime.GOOS,
		timer:    component.NewTimer(0),
		cpuInfo:  cpuInfo,
		someInfo: "Vlad pidor",
	}
}
