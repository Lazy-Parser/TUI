package service

import "context"

func (s *Service) getAllTokens() ResponseGetAllTokensMsg {
	t, err := s.tokenRepo.GetAll(context.Background())
	return ResponseGetAllTokensMsg{Tokens: t, Err: err}
}
