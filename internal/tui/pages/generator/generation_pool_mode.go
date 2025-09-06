package page_generator

import (
	"fmt"
	"log"

	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/logic"
	"github.com/Lazy-Parser/TUI/internal/service"
	"github.com/Lazy-Parser/TUI/internal/task"
	"github.com/Lazy-Parser/TUI/internal/tui/common"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	internalFieldsEmpty = "Warning in Page -> Generator -> Generation mode -> Init(): cannot send msg to start 'Fetch pool' task, because internal fields 'Address' or 'Network' is / are empty"
)

type GeneratationPoolMode struct {
	Address string
	Network string

	info      string
	pool      market.Pool
	isLoading bool
}

func NewModeGenerationAddPool() *GeneratationPoolMode {
	return &GeneratationPoolMode{isLoading: true}
}

func (model *GeneratationPoolMode) Init() tea.Cmd {
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

func (model *GeneratationPoolMode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case logic.FetchPoolResultMsg:
		return model.handleFetchPoolResult(msg)
	case service.ResponseSavePoolMsg:
		return model.handleSavePoolResponse(msg)
	}

	return model, nil
}

func (model *GeneratationPoolMode) View() string {
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
func (mode *GeneratationPoolMode) Set(address string, network string) {
	mode.Address = address
	mode.Network = network
}

func (mode *GeneratationPoolMode) handleFetchPoolResult(msg logic.FetchPoolResultMsg) (tea.Model, tea.Cmd) {
	if msg.Err != nil {
		mode.info = msg.Err.Error()
	}
	mode.pool = msg.Pool
	mode.isLoading = false

	// try to save to database. TODO: ask if user want to save pool to the db
	savePool := common.CmdHandler(service.RequestSavePoolMsg{
		Pool:       msg.Pool,
		BaseToken:  msg.Pool.Pair.BaseToken,
		QuoteToken: msg.Pool.Pair.QuoteToken,
	})

	return mode, savePool
}

func (mode *GeneratationPoolMode) handleSavePoolResponse(msg service.ResponseSavePoolMsg) (tea.Model, tea.Cmd) {
	if msg.Err == nil {
		mode.info = "saved to database!"
	} else {
		mode.info = msg.Err.Error()
	}

	return mode, nil
}
