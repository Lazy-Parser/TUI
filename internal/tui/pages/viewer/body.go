package page_viewer

import (
	"log"

	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/service"
	"github.com/Lazy-Parser/TUI/internal/tui/common"
	custom "github.com/Lazy-Parser/TUI/internal/tui/components"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	TableModeTokens Mode = iota
	TableModePools
)

type mainView struct {
	width, height int

	tokens []market.Token
	pools  []market.Pool

	table *custom.Table
	mode  Mode

	keys keymap
}

func NewMain() tea.Model {
	return &mainView{table: custom.NewModel(), mode: TableModeTokens, keys: keys}
}

func (m *mainView) Init() tea.Cmd {
	tableCmd := m.table.Init()
	load := m.loadTokensPools()

	return tea.Batch(tableCmd, load)
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case service.ResponseGetAllTokensMsg:
		return m.handleGetAllTokensResponse(msg)
	case service.ResponseGetAllPoolsMsg:
		return m.handleGetAllPoolsResponse(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for border (border goes to the left)
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.showPools):
			return m.showPools()
		case key.Matches(msg, m.keys.showTokens):
			return m.showTokens()

		case key.Matches(msg, m.keys.r):
			return m, m.loadTokensPools()
		case key.Matches(msg, m.keys.q):
			return m, tea.Quit
		}

	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m *mainView) View() string {
	return m.table.View()
}

// methods
func (m *mainView) loadTokensPools() tea.Cmd {
	loadTokensCmd := common.CmdHandler(service.RequestGetAllTokensMsg{})
	loadPoolsCmd := common.CmdHandler(service.RequestGetAllPoolsMsg{})
	return tea.Batch(loadTokensCmd, loadPoolsCmd)
}

func (m *mainView) handleGetAllTokensResponse(msg service.ResponseGetAllTokensMsg) (tea.Model, tea.Cmd) {
	if msg.Err != nil {
		log.Println(msg.Err)
		return m, nil
	}
	m.tokens = msg.Tokens

	// update if current table is tokens
	if m.mode == TableModeTokens {
		m.table.SetRowsTokens(m.tokens)
	}

	return m, nil
}

func (m *mainView) handleGetAllPoolsResponse(msg service.ResponseGetAllPoolsMsg) (tea.Model, tea.Cmd) {
	if msg.Err != nil {
		log.Println(msg.Err)
		return m, nil
	}
	m.pools = msg.Pools

	// update if current table is pools
	if m.mode == TableModePools {
		m.table.SetRowsPools(m.pools)
	}

	return m, nil
}

func (m *mainView) showPools() (tea.Model, tea.Cmd) {
	m.table.ShowPoolsCol()
	m.table.SetRowsPools(m.pools)

	return m, nil
}

func (m *mainView) showTokens() (tea.Model, tea.Cmd) {
	m.table.ShowTokensCol()
	m.table.SetRowsTokens(m.tokens)

	return m, nil
}
