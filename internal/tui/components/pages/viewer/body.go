package page_viewer

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Lazy-Parser/Collector/market"
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct {
	tokenRepo market.TokenRepo
	tokens    []market.Token
}

func NewMain(tokenRepo market.TokenRepo) tea.Model {
	return &mainView{tokenRepo: tokenRepo}
}

func (m *mainView) Init() tea.Cmd {
	return loadAllTokens(context.Background(), m.tokenRepo)
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			log.Println("Reload")
			return m, loadAllTokens(context.Background(), m.tokenRepo)

		case "q":
			log.Println("Quit")
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case tokensMsg:
		log.Println("New message!")
		if msg.err != nil {
			log.Println(fmt.Errorf("failed to load tokens: %w", msg.err))
			return m, nil
		}
		if len(msg.tokens) == 0 {
			log.Println("Not tokens loaded...")
			m.tokens = nil
		} else {
			log.Println("Tokens exist!")
			m.tokens = msg.tokens
		}
	}

	return m, nil
}

func (m *mainView) View() string {
	if len(m.tokens) == 0 {
		return "Empty..."
	}

	var strBuilder strings.Builder
	for _, token := range m.tokens {
		strBuilder.WriteString(token.Name)
		strBuilder.WriteString("\n")
	}

	return strBuilder.String()
}
