package service

import "github.com/Lazy-Parser/Collector/market"

// get list of supported exchanges
type RequestExchangesListMsg struct{}
type ResponseExchangesListMsg struct{ Exchanges []string }

// database operations
type RequestSavePoolMsg struct {
	Pool       market.Pool
	BaseToken  market.Token
	QuoteToken market.Token
}

// idk if i should make responses for database operations. But maybe better to make it, to handle errors / or show saved info. But i will make it later
type ResponseSavePoolMsg struct{ Err error }

type RequestGetAllTokensMsg struct{}
type ResponseGetAllTokensMsg struct {
	Tokens []market.Token
	Err    error
}

type RequestGetAllPoolsMsg struct{}
type ResponseGetAllPoolsMsg struct {
	Pools []market.Pool
	Err   error
}
