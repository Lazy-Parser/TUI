package page_viewer

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/Collector/market"
	component "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

func onTokensMsg(m *mainView, msg *tokensMsg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if msg.err != nil {
		log.Println(fmt.Errorf("failed to load tokens: %w", msg.err))
		return m, nil
	}

	m.tokens = msg.tokens

	// set token amount in header
	cmd := setTokenAmount(len(msg.tokens))
	cmds = append(cmds, cmd)

	// set tokens in table
	cmd = sendTokens(m.tokens)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func sendTokens(tokens []market.Token) tea.Cmd {
	return func() tea.Msg {
		return component.SetContentTokensMsg{Tokens: tokens}
	}
}
