package page_generator

import (
	"fmt"

	"github.com/Lazy-Parser/Collector/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func makeBold(str string) string {
	return lipgloss.NewStyle().Bold(true).Render(str)
}

type header struct {
	cfg *config.Config
}

func NewHeader(cfg *config.Config) tea.Model {
	return &header{cfg: cfg}
}

func (h *header) Init() tea.Cmd                           { return nil }
func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return h, nil }
func (h *header) View() string {
	var str string

	str += "Mexc:\n"
	str += fmt.Sprintf("  access   %s (no limit)\n", makeBold(h.cfg.Mexc.ACCESS_TOKEN))
	str += fmt.Sprintf("  private  %s (no limit)\n", makeBold(h.cfg.Mexc.PRIVATE_TOKEN))
	str += "\n"
	str += fmt.Sprintf("Coingecko: %s (watch limits / usage on https://www.coingecko.com/developers/dashboard)", makeBold(h.cfg.Coingecko.API.KEY))

	return lipgloss.NewStyle().Align(lipgloss.Left).Render(str)
}
