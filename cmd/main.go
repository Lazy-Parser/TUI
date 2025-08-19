package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Lazy-Parser/TUI/internal/tui"

	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/database"
)

func main() {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "storage", "storage.db")

	db, err := database.Start(path)
	if err != nil {
		panic(err)
	}
	tokenRepo := database.NewTokenRepo(db)

	// config
	path = filepath.Join(wd, ".env")
	cfg, err := config.NewConfig(path)
	if err != nil {
		err = fmt.Errorf("solve: for running this app you need to place '.env' file in the folder of the app. %v", err)
		panic(err)
	}

	// chains
	path = filepath.Join(wd, "chains.json")
	chainsService, err := chains.NewChains(path)
	if err != nil {
		err = fmt.Errorf("failed to create chains service, the possible problem is that you passed wrong path to 'chains.json'. %v", err)
		panic(err)
	}

	if err := tui.Run(tokenRepo, cfg, chainsService); err != nil {
		panic(err)
	}
}
