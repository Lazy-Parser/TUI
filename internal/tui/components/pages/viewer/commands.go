package page_viewer

import (
	"context"

	"github.com/Lazy-Parser/Collector/market"
	tea "github.com/charmbracelet/bubbletea"
)

type tokensMsg struct {
	tokens []market.Token
	err    error
}

func loadAllTokens(ctx context.Context, tokenRepo market.TokenRepo) tea.Cmd {
	return func() tea.Msg {
		tokens, err := tokenRepo.GetAll(ctx)
		return tokensMsg{tokens: tokens, err: err}
	}
}
