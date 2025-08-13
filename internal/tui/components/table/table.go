package custom

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

const (
	minWidth  = 30
	minHeight = 8

	// Add a fixed margin to account for description & instructions
	fixedVerticalMargin = 4
)

type Table struct {
	table table.Model

	width  int
	height int

	// Table dimensions
	horizontalMargin int
	verticalMargin   int
}

func (ct Table) Init() tea.Cmd {
	return nil
}

func (ct Table) Update(msg tea.Msg) (Table, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	ct.table, cmd = ct.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			cmds = append(cmds, tea.Quit)
		}

		switch msg.Type {
		case tea.KeyCtrlC:
			cmds = append(cmds, tea.Quit)
		}

	case tea.WindowSizeMsg:
		ct.width = msg.Width
		ct.height = msg.Height

		ct.recalculateTable()
	}

	return ct, tea.Batch(cmds...)
}

func (ct Table) View() string {
	return ct.table.View() + "\n"
}

func (ct Table) recalculateTable() {
	ct.table = ct.table.
		WithTargetWidth(ct.calculateWidth()).
		WithMinimumHeight(ct.calculateHeight())
}
func (ct Table) calculateWidth() int {
	return ct.width - ct.horizontalMargin
}
func (ct Table) calculateHeight() int {
	return ct.height - ct.verticalMargin - fixedVerticalMargin
}

func NewModel() Table {
	return Table{
		table: table.New([]table.Column{
			table.NewColumn("name", "Name", 10),
			table.NewFlexColumn("network", "Network", 1),
			table.NewFlexColumn("decimal", "Decimal", 3),
		}).WithRows([]table.Row{
			table.NewRow(table.RowData{
				"name":    "BTC",
				"network": "Ethereum",
				"decimal": "8",
			}),
			table.NewRow(table.RowData{
				"name":    "ETH",
				"network": "Solana",
				"decimal": "18",
			}),
		}).WithStaticFooter("A footer!"),

		horizontalMargin: 4,
		verticalMargin:   2,
	}
}
