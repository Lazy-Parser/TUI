package tui

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Lazy-Parser/Collector/market"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(tokenRepo market.TokenRepo) error {
	f := StartLogger()

	// try to create test tokens
	ctx := context.Background()
	tokens := []market.Token{
		market.Token{
			Name:        "Trump",
			Decimal:     8,
			Address:     "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			Image_url:   "https://dd.dexscreener.com/ds-data/tokens/solana/6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN/header.png?key=f02e9e",
			WithdrawFee: "10",
			CreateTime:  time.Now().UnixMilli(),
			Network:     "Solana",
		},
		market.Token{
			Name:        "USDT",
			Decimal:     16,
			Address:     "0xfde4C96c8593536E31F229EA8f37b2ADa2699bb2",
			Image_url:   "",
			WithdrawFee: "5",
			CreateTime:  time.Now().UnixMilli(),
			Network:     "Ethereum",
		},
		market.Token{
			Name:        "BR",
			Decimal:     10,
			Address:     "0xFf7d6A96ae471BbCD7713aF9CB1fEeB16cf56B41",
			Image_url:   "https://dd.dexscreener.com/ds-data/tokens/bsc/0xff7d6a96ae471bbcd7713af9cb1feeb16cf56b41/header.png?key=1a07a9",
			WithdrawFee: "20",
			CreateTime:  time.Now().UnixMilli(),
			Network:     "BSC",
		},
	}
	
	for _, token := range tokens {
		if err := tokenRepo.Save(ctx, token); err != nil {
			log.Println(err)
		}
	}

	p := tea.NewProgram(InitLayout(tokenRepo), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		f.Close()
		return fmt.Errorf("Alas, there's been an error: %v", err)
	}

	f.Close()
	return nil
}
