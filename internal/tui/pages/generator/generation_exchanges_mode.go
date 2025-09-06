package page_generator

import (
	"log"

	component "github.com/Lazy-Parser/TUI/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type ModeExchangeGenerator struct {
	exchanges []string
	loading bool
	error     error

	loader *component.Loader
}

func NewModeExchangeGenerator() *ModeExchangeGenerator {
	return &ModeExchangeGenerator{
		loading: true, 
		loader: component.NewLoader("Generating"),
	}
}

func (model *ModeExchangeGenerator) Init() tea.Cmd {
	if len(model.exchanges) == 0 {
		log.Println("Generator -> Generate exchanges -> Init(): failed to init, because cannot send empty list of exchanges!")
		return nil
	}

	// start generation task

	// do not forget about loader init()
	return model.loader.Init()
}

func (model *ModeExchangeGenerator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch msg := msg.(type) {
	// 	case
	// }

	// update loader for animation
	var cmd tea.Cmd
	model.loader, cmd = model.loader.Update(msg)

	return model, cmd
}

func (model *ModeExchangeGenerator) View() string {
	if len(model.exchanges) == 0 {
		return "Something went wrong! Why exchanges did not selected?"
	}

	if model.isLoading() {
		return model.loader.View()
	}
	if model.isError() {
		return model.error.Error()
	}
	if model.isDone() {
		return "Done!"
	}
	
	return ""
}

// methods
func (model *ModeExchangeGenerator) SetExchanges(arr []string) {
	model.exchanges = arr
}

func (model *ModeExchangeGenerator) handleSomeTaskEnd() (tea.Model, tea.Cmd) {
	// do not forger to stop loader
	model.loader.StopLoader()

	// do smth

	return model, nil
}

func (model *ModeExchangeGenerator) isLoading() bool {
	return model.loading
}

func (model *ModeExchangeGenerator) isError() bool {
	return model.error != nil
}

func (model *ModeExchangeGenerator) isDone() bool {
	return model.loading == false && model.error == nil
}
