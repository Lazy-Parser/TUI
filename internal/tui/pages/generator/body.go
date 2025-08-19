package page_generator

import (
	"context"
	"fmt"

	"github.com/Lazy-Parser/Collector/api"
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/Collector/worker"
	tea "github.com/charmbracelet/bubbletea"
)
type FuturesMsg struct {
	futures []market.Token
	err     error
}

type mainView struct {
	cfg           *config.Config
	chainsService *chains.Chains
	start         bool
	futures       []market.Token
	steps         []Step
}

func NewMain(cfg *config.Config, chainsService *chains.Chains) tea.Model {
	return &mainView{cfg: cfg, chainsService: chainsService, steps: make([]Step, 1)}
}

func (m *mainView) Init() tea.Cmd {
	return nil
}

func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FuturesMsg:
		if msg.err != nil {
			m.steps[0].err = msg.err
			return m, nil
		}
		m.futures = msg.futures
		m.steps[0].isLoading = false
		return m, nil

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

	return m, nil
}

func (m *mainView) View() string {
	if !m.start {
		return "Press [Enter] to start!"
	} else {
		if m.steps[0].err != nil {
			return m.steps[0].err.Error()
		}
		if m.steps[0].isLoading {
			return "Loading..."
		}
		if len(m.futures) > 0 {
			var str string
			for _, token := range m.futures {
				str += fmt.Sprintf("Name: %s, Network: %s\n", token.Name, token.Network)
			}
			return str
		}

		return ""
	}
}

func (m *mainView) handleEnter() (*mainView, tea.Cmd) {
	if m.start {
		return m, nil
	}

	m.start = true
	// create first step
	m.steps[0] = Step{isLoading: true}

	return m, m.getFutures()
}

func (m *mainView) getFutures() tea.Cmd {
	return func() tea.Msg {
		// create mexcApi. TODO: create mexc api in the root, not there
		ctx := context.Background()
		api := api.NewMexcApi(m.cfg)
		mexc := worker.NewMexcWorker(api, m.chainsService)

		tokens, err := mexc.GetAllTokens(ctx)
		if err != nil {
			return FuturesMsg{futures: nil, err: fmt.Errorf("failed to fetch all tokens (Step #1) from Mexc exchange: %v", err)}
		}

		futures, err := mexc.GetAllFutures(ctx)
		if err != nil {
			return FuturesMsg{futures: nil, err: fmt.Errorf("failed to fetch futures (Step #2) from Mexc exchange: %v", err)}
		}

		var res []market.Token
		for _, token := range tokens {
			contract, ok := mexc.FindContractBySymbol(&futures, token.Coin)
			if !ok {
				// current token does not exist on futures
				continue
			}

			res = append(res, market.Token{
				Name:        token.Coin,
				Decimal:     0, // unknown, will find it later
				Network:     token.NetworkList[0].Network,
				Address:     token.NetworkList[0].Contract, // TODO: NetworkList can contain > 1 elems
				WithdrawFee: token.NetworkList[0].WithdrawFee,
				Image_url:   contract.ImageUrl,
				CreateTime:  contract.CreateTime,
			})
		}

		return FuturesMsg{futures: res, err: nil}
	}
}
