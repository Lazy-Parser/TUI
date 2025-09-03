package page_generator

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/logic"
	"github.com/Lazy-Parser/TUI/internal/task"
	"github.com/Lazy-Parser/TUI/internal/tui/common"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	internalFieldsEmpty = "Warning in Page -> Generator -> Generation mode -> Init(): cannot send msg to start 'Fetch pool' task, because internal fields 'Address' or 'Network' is / are empty"
)

type GeneratationMode struct {
	Address string
	Network string

	info      string
	pool      market.Pool
	isLoading bool
}

func NewGenetationMode() *GeneratationMode {
	return &GeneratationMode{isLoading: true}
}

func (model *GeneratationMode) Init() tea.Cmd {
	if model.Address == "" || model.Network == "" {
		log.Print(internalFieldsEmpty)
		return nil
	}

	// start task
	return common.CmdHandler(task.NewFetchPoolTaskMsg{
		Address: model.Address,
		Network: model.Network,
	})
}

func (model *GeneratationMode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case logic.FetchPoolResultMsg:
		return model.handleFetchPoolResult(msg)
	}

	return model, nil
}

func (model *GeneratationMode) View() string {
	var str string

	if model.isLoading {
		str += "Loading..."
	} else {
		str += "Info " + model.info + "\n"
		str += fmt.Sprintf("Pool: %+v\n", model.pool)
	}

	return str
}

// methods
func (mode *GeneratationMode) Set(address string, network string) {
	mode.Address = address
	mode.Network = network
}

func (mode *GeneratationMode) handleFetchPoolResult(msg logic.FetchPoolResultMsg) (tea.Model, tea.Cmd) {
	if msg.Err != nil {
		mode.info = msg.Err.Error()
	}
	mode.pool = msg.Pool
	mode.isLoading = false

	return mode, nil
}
