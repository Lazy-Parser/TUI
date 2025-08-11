package main

import (
	"fmt"

	service "github.com/Lazy-Parser/Collector/config/service"
	"github.com/Lazy-Parser/TUI/internal/tui"
)

func main() {
	// put some test fields for tests
	cfgService, err := service.NewConfig("some filepath")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfgService.Coingecko.API.KEY)

	if err := tui.Run(); err != nil {
		panic(err)
	}
}
