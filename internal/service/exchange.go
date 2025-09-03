package service

func (s *Service) getAllExchanges() ResponseExchangesListMsg {
	return ResponseExchangesListMsg{
		Exchanges: []string{"Mexc"}, // make up to store exchanges somewhere
	}
}
