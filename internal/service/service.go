// This service works like an HTTP. Some component need to send some RequestMsg, and this service will answer with ResponseMsg. So no need to import in anywhere except of layout.go
package service

import tea "github.com/charmbracelet/bubbletea"

type Service struct {
	ch chan tea.Msg
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetMsgChannel(ch chan tea.Msg) {
	s.ch = ch
}

func (s *Service) HandleMsg(msg tea.Msg) {
	switch msg.(type) {
		case RequestExchangesListMsg:
			s.ch <- s.getAllExchanges()
			break
	}
}