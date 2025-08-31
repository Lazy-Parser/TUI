package page_generator

import "github.com/Lazy-Parser/Collector/market"

// set modes
type SelectionModeMsg struct{}
type AddPoolModeMsg struct{}
type GenerationMsg struct{}

// result msg for every mode state are located in each mode file accordinaly.  

// for internal logic
type FuturesMsg struct {
	futures []market.Token
	err     error
}
type DexscreenerMsg struct {
	pairs []market.Pair
	err   error
}
type DecimalsMsg struct {
	futures []market.Token
	err     error
}
