package service

import (
	"context"
	"fmt"
)

func (s *Service) savePoolWithTokens(toSave RequestSavePoolMsg) ResponseSavePoolMsg {
	// TODO: hardcoded context
	ctx := context.Background()

	// first save tokens
	baseId, err := s.tokenRepo.FindOrCreate(ctx, toSave.BaseToken)
	if err != nil {
		return ResponseSavePoolMsg{Err: fmt.Errorf("failed to save base token: %v", err)}
	}

	quoteId, err := s.tokenRepo.FindOrCreate(ctx, toSave.QuoteToken)
	if err != nil {
		return ResponseSavePoolMsg{Err: fmt.Errorf("failed to save quote token: %v", err)}
	}

	// then save pool and pass saved tokens ids
	_, err = s.poolRepo.FindOrCreate(ctx, toSave.Pool, baseId, quoteId)
	if err != nil {
		return ResponseSavePoolMsg{Err: fmt.Errorf("failed to save pool: %v", err)}
	}

	// without error - success
	return ResponseSavePoolMsg{}
}

func (s *Service) getAllPools() ResponseGetAllPoolsMsg {
	p, err := s.poolRepo.GetAll(context.Background())
	return ResponseGetAllPoolsMsg{Pools: p, Err: err}
}
