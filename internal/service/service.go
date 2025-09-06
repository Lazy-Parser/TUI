// This service works like an HTTP. Some component need to send some RequestMsg, and this service will answer with ResponseMsg. So no need to import in anywhere except of layout.go
package service

import (
	"github.com/Lazy-Parser/Collector/market"
	tea "github.com/charmbracelet/bubbletea"
)

type Service struct {
	ch        chan tea.Msg
	tokenRepo market.TokenRepo
	poolRepo  market.PoolRepo
}

func NewService(tokenRepo market.TokenRepo, poolRepo market.PoolRepo) *Service {
	return &Service{
		tokenRepo: tokenRepo,
		poolRepo:  poolRepo,
	}
}

func (s *Service) SetMsgChannel(ch chan tea.Msg) {
	s.ch = ch
}

func (s *Service) HandleMsg(msg tea.Msg) {
	switch msg := msg.(type) {
	case RequestExchangesListMsg:
		s.ch <- s.getAllExchanges()
	case RequestSavePoolMsg:
		s.ch <- s.savePoolWithTokens(msg)
	case RequestGetAllTokensMsg:
		s.ch <- s.getAllTokens()
	case RequestGetAllPoolsMsg:
		s.ch <- s.getAllPools()
	}
}
