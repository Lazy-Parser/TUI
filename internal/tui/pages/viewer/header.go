package page_viewer

import (
	"fmt"

	"github.com/Lazy-Parser/TUI/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type header struct {
	tokensAmount int
	poolsAmount  int
}

func NewHeader() tea.Model { return &header{tokensAmount: -1, poolsAmount: -1} }

func (h *header) Init() tea.Cmd { return nil }

func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case service.ResponseGetAllTokensMsg:
		return h.handleAllTokensResponse(msg)
	case service.ResponseGetAllPoolsMsg:
		return h.handleAllPoolsResponse(msg)
	}

	return h, nil
}

func (h *header) View() string {
	var str string

	str += "Tokens: "
	if h.tokensAmount == -1 {
		str += "?"
	} else if h.tokensAmount == 0 {
		str += "No tokens"
	} else {
		str += fmt.Sprintf("%d", h.tokensAmount)
	}

	str += "\n"
	str += "Pairs: "
	if h.poolsAmount == -1 {
		str += "?"
	} else if h.poolsAmount == 0 {
		str += "No Pools"
	} else {
		str += fmt.Sprintf("%d", h.poolsAmount)
	}

	return str
}

// methods
func (h *header) handleAllTokensResponse(msg service.ResponseGetAllTokensMsg) (tea.Model, tea.Cmd) {
	if msg.Err == nil {
		h.tokensAmount = len(msg.Tokens)
	}

	return h, nil
}

func (h *header) handleAllPoolsResponse(msg service.ResponseGetAllPoolsMsg) (tea.Model, tea.Cmd) {
	if msg.Err == nil {
		h.poolsAmount = len(msg.Pools)
	}

	return h, nil
}
