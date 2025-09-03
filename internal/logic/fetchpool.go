package logic

import (
	"context"

	"github.com/Lazy-Parser/Collector/market"
	tea "github.com/charmbracelet/bubbletea"
)

type FetchPoolResultMsg struct {
	Pool market.Pool
	Err  error
}

// todo: save in bd
func (core *Logic) FetchPoolByToken(ch chan<- tea.Msg, address string, network string) {
	pool, err := core.DiscoveryService.MetaByAddress(context.Background(), network, address)
	ch <- FetchPoolResultMsg{Pool: pool, Err: err}
}
