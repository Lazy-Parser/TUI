package main

import (
	"tui/internal/tui"
	"github.com/Lazy-Parser/Collector/market"
)

func main() {
	// put some test fields for tests

	
	if err := tui.Run(); err != nil {
		panic(err)
	}
}
