package main

import (
	"os"
	"path/filepath"

	"github.com/Lazy-Parser/TUI/internal/tui"

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

	if err := tui.Run(tokenRepo); err != nil {
		panic(err)
	}
}
