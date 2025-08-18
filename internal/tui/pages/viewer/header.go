package page_viewer

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type header struct {
	tokensAmount int
	pairsAmount  int
}

func NewHeader() tea.Model { return &header{tokensAmount: 0, pairsAmount: 0} }

func (h *header) Init() tea.Cmd { return nil }

func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tokenAmountMsg:
		h.tokensAmount = msg.amount
	case pairAmountMsg:
		h.pairsAmount = msg.amount
	}

	return h, nil
}

func (h *header) View() string {
	var str string

	str += "Tokens: "
	if h.tokensAmount == 0 {
		str += "No tokens"
	} else {
		str += fmt.Sprintf("%d", h.tokensAmount)
	}

	str += "\n"
	str += "Pairs: "
	if h.pairsAmount == 0 {
		str += "?"
	} else {
		str += fmt.Sprintf("%d", h.pairsAmount)
	}

	return str
}
