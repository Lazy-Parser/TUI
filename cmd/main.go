package main

import (
	"github.com/Lazy-Parser/TUI/internal/tui"
)

func main() {
	// put some test fields for tests

	
	if err := tui.Run(); err != nil {
		panic(err)
	}
}
