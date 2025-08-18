package page_viewer

import (
	"context"

	"github.com/Lazy-Parser/Collector/market"
	custom "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct {
	width, height int

	tokenRepo market.TokenRepo
	tokens    []market.Token

	table custom.Table
}

func NewMain(tokenRepo market.TokenRepo) tea.Model {
	return &mainView{tokenRepo: tokenRepo, table: custom.NewModel()}
}

func (m *mainView) Init() tea.Cmd {
	cmd := loadAllTokens(context.Background(), m.tokenRepo)
	cmd2 := m.table.Init()

	return tea.Batch(cmd, cmd2)
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for border (border goes to the left)
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, loadAllTokens(context.Background(), m.tokenRepo)

		case "q":
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case tokensMsg:
		return onTokensMsg(m, &msg)

		// do the same for pairsMsg
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m *mainView) View() string {
	return m.table.View()
}
