package page_generator

import (
	"context"
	"fmt"
	"log"

	"github.com/Lazy-Parser/Collector/api"
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/Collector/worker"
	tea "github.com/charmbracelet/bubbletea"
)

type logic struct {
	cfg           *config.Config
	chainsService *chains.Chains
}

func newLogic(cfg *config.Config, chainsService *chains.Chains) *logic {
	return &logic{
		cfg:           cfg,
		chainsService: chainsService,
	}
}

// FUTURES
// GetFutures fetches tokens and futures from Mexc
func (l *logic) GetFutures() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()

		// Initialize services
		api := api.NewMexcApi(l.cfg)
		mexc := worker.NewMexcWorker(api, l.chainsService)

		// Fetch data
		tokens, err := l.fetchTokens(ctx, mexc)
		if err != nil {
			return FuturesMsg{futures: nil, err: err}
		}

		futures, err := l.fetchFutures(ctx, mexc)
		if err != nil {
			return FuturesMsg{futures: nil, err: err}
		}

		// Process and combine data
		result := l.combineTokensWithFutures(tokens, futures, mexc)

		return FuturesMsg{futures: result, err: nil}
	}
}

// fetchTokens gets all tokens from Mexc
func (l *logic) fetchTokens(ctx context.Context, mexc worker.MexcWorker) ([]market.MexcAsset, error) {
	tokens, err := mexc.GetAllTokens(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tokens from Mexc: %v", err)
	}
	return tokens, nil
}

// fetchFutures gets all futures from Mexc
func (l *logic) fetchFutures(ctx context.Context, mexc worker.MexcWorker) ([]market.MexcContractDetail, error) {
	futures, err := mexc.GetAllFutures(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch futures from Mexc: %v", err)
	}
	return futures, nil
}

// combineTokensWithFutures processes tokens and matches them with futures
func (l *logic) combineTokensWithFutures(
	tokens []market.MexcAsset,
	futures []market.MexcContractDetail,
	mexc worker.MexcWorker,
) []market.Token {
	var result []market.Token

	for _, token := range tokens {
		// Find matching contract
		contract, ok := mexc.FindContractBySymbol(&futures, token.Coin)
		if !ok {
			// Skip tokens without futures
			continue
		}

		// Create market token
		marketToken := l.createMarketToken(token, contract)
		result = append(result, marketToken)
	}

	return result
}

// createMarketToken converts API data to market.Token
func (l *logic) createMarketToken(token market.MexcAsset, contract market.MexcContractDetail) market.Token {
	return market.Token{
		Name:        token.Coin,
		Decimal:     0, // Will be filled later
		Network:     token.NetworkList[0].Network,
		Address:     token.NetworkList[0].Contract,
		WithdrawFee: token.NetworkList[0].WithdrawFee,
		Image_url:   contract.ImageUrl,
		CreateTime:  contract.CreateTime,
	}
}

// ---- PAIRS ----
func (l *logic) GetPairs(tokens []market.Token) tea.Cmd {
	return func() tea.Msg {
		// create instance first
		dsApi := api.NewDexscreenerApi(l.cfg)
		dsWorker := worker.NewDexscreenerWorker(dsApi, l.chainsService)

		ctx := context.Background()
		pairs := make([]market.Pair, len(tokens))
		emptyPair := market.Pair{}
		for _, token := range tokens {
			pair, err := dsWorker.FetchPairByToken(ctx, token)
			if err != nil {
				log.Println(err)
				continue
			}
			if pair == emptyPair {
				// couldnot find pair for provided token
				continue
			}

			pairs = append(pairs, pair)
		}

		return DexscreenerMsg{pairs: pairs, err: nil}
	}
}
