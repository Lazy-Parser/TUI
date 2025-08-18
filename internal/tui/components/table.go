package component

import (
	"fmt"
	"time"

	"github.com/Lazy-Parser/Collector/market"
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

type SetContentTokensMsg struct {
	Tokens []market.Token
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
	case SetContentTokensMsg:
		ct.table = ct.setRows(msg.Tokens)
		return ct, nil

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
			table.NewFlexColumn("name", "Name", 3),
			table.NewFlexColumn("network", "Network", 3),
			table.NewFlexColumn("decimal", "Decimal", 1),
			table.NewFlexColumn("address", "Address", 5),
			table.NewFlexColumn("createdAt", "Created At", 3),
			table.NewFlexColumn("fee", "fee", 1),
		}).WithStaticFooter("Waiting...").Focused(true).SelectableRows(true).BorderRounded(),

		horizontalMargin: 4,
		verticalMargin:   2,
	}
}

func (ct Table) setRows(tokens []market.Token) table.Model {
	rows := make([]table.Row, len(tokens))
	for _, token := range tokens {
		rowData := table.RowData{
			"name":      token.Name,
			"network":   token.Network,
			"decimal":   token.Network,
			"address":   token.Address,
			"createdAt": milliToDate(token.CreateTime),
			"fee":       token.WithdrawFee,
		}

		rows = append(rows, table.NewRow(rowData))
	}

	ct.table.WithRows(rows)
	return ct.table
}

func milliToDate(milliseconds int64) string {
	y, m, d := time.UnixMilli(milliseconds).Date()

	return fmt.Sprintf("%s %d, %d", m.String(), d, y)
}
