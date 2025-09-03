package logic

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Lazy-Parser/Collector/api"
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/discovery"
)

type Logic struct {
	cfg              *config.Config
	chainsService    *chains.Chains
	DiscoveryService discovery.Discovery
}

func NewLogic() (*Logic, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config, reason: %v", err)
	}

	chainsService, err := createChainsService()
	if err != nil {
		return nil, fmt.Errorf("failed to create chains service, reason: %v", err)
	}

	discoveryService := createDiscoveryService(cfg, chainsService)

	return &Logic{
		cfg:              cfg,
		chainsService:    chainsService,
		DiscoveryService: discoveryService,
	}, nil
}

func loadConfig() (*config.Config, error) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, ".env")

	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func createChainsService() (*chains.Chains, error) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "chains.json")

	chains, err := chains.NewChains(path)
	if err != nil {
		return nil, err
	}

	return chains, nil
}

func createDiscoveryService(cfg *config.Config, chainsService *chains.Chains) discovery.Discovery {
	dsApi := api.NewDexscreenerApi(cfg)
	cgApi := api.NewCoingeckoApi(cfg)

	return discovery.NewDiscovery(dsApi, cgApi, chainsService)
}
