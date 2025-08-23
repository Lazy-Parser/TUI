package page_generator

import (
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/market"
	component "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FuturesMsg struct { futures []market.Token; err error }
type DexscreenerMsg struct { pairs []market.Pair; err error }
type DecimalsMsg struct { futures []market.Token; err error }

type mainView struct {
	logic    *logic
	start    bool
	futures  []market.Token
	pairs []market.Pair
	taskFlow *component.TaskFlow
}

func NewMain(cfg *config.Config, chainsService *chains.Chains) tea.Model {
	return &mainView{
		logic: newLogic(cfg, chainsService),
		taskFlow: component.NewTaskFlow(
			component.NewTask("Fetch futures from Mexc", ""),
			component.NewTask("Fetch Pairs from DS ", ""),
			// component.NewTask("Fetch decimals from CG", ""),
		),
	}
}

func (m *mainView) Init() tea.Cmd {
	return m.taskFlow.Init()
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FuturesMsg:
		return m.handleFuturesMsg(msg)

	case DexscreenerMsg:
		return m.handleDsMsg(msg)
		
	case component.TaskEndMsg:
		return m.handleTasksEnd()
		
	case tea.KeyMsg:
		switch msg.String() {
		// exit
		case "q":
			return m, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEnter:
			return m.handleEnter()
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.taskFlow, cmd = m.taskFlow.Update(msg)

	return m, cmd
}

func (m *mainView) View() string {
	if !m.start {
		return "Press [Enter] to start!"
	} else {
		return lipgloss.NewStyle().Align(lipgloss.Left).Render(m.taskFlow.View())
	}
}