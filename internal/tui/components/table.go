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
	maxHeight = 30

	// Add a fixed margin to account for description & instructions
	fixedVerticalMargin = 4
)

var (
	tokensColumns = []table.Column{
		table.NewFlexColumn("name", "Name", 3),
		table.NewFlexColumn("network", "Network", 3),
		table.NewFlexColumn("decimal", "Decimal", 1),
		table.NewFlexColumn("address", "Address", 5),
		table.NewFlexColumn("createdAt", "Created At", 3),
	}
	poolsColumns = []table.Column{
		table.NewFlexColumn("base_name", "Base", 2),
		table.NewFlexColumn("quote_name", "Quote", 2),
		table.NewFlexColumn("network", "Network", 2),
		table.NewFlexColumn("pool", "Pool", 3),
		table.NewFlexColumn("address", "Address", 4),
	}
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

func (ct *Table) Init() tea.Cmd {
	return nil
}

func (ct *Table) Update(msg tea.Msg) (*Table, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	ct.table, cmd = ct.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case SetContentTokensMsg:
		ct.table = ct.setTokensRows(msg.Tokens)
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

func (ct *Table) View() string {
	return ct.table.View() + "\n"
}

func (ct *Table) recalculateTable() {
	ct.table = ct.table.
		WithTargetWidth(ct.calculateWidth()).
		WithMinimumHeight(20)
}
func (ct *Table) calculateWidth() int {
	return ct.width - ct.horizontalMargin
}
func (ct *Table) calculateHeight() int {
	return ct.height - ct.verticalMargin - fixedVerticalMargin
}

func NewModel() *Table {
	return &Table{
		table: table.New(tokensColumns).
			WithStaticFooter("Waiting...").
			Focused(true).
			SelectableRows(true).
			BorderRounded().
			WithPageSize(20),

		horizontalMargin: 4,
		verticalMargin:   2,
	}
}

func (ct *Table) setTokensRows(tokens []market.Token) table.Model {
	rows := make([]table.Row, len(tokens))
	for i, token := range tokens {
		rowData := table.RowData{
			"name":      token.Name,
			"network":   token.Network,
			"decimal":   token.Decimal,
			"address":   token.Address,
			"createdAt": milliToDate(token.CreateTime),
			"fee":       token.WithdrawFee,
		}

		rows[i] = table.NewRow(rowData)
	}

	newTable := ct.table.WithRows(rows).WithStaticFooter("")
	return newTable
}

func (ct *Table) setPoolsRows(pools []market.Pool) table.Model {
	rows := make([]table.Row, len(pools))
	for i, pool := range pools {
		rowData := table.RowData{
			"base_name":  tokenWithAddress(pool.Pair.BaseToken),
			"quote_name": tokenWithAddress(pool.Pair.QuoteToken),
			"network":    pool.Network,
			"pool":       pool.Pool,
			"address":    pool.Address,
		}

		rows[i] = table.NewRow(rowData)
	}

	newTable := ct.table.WithRows(rows).WithStaticFooter("")
	return newTable
}
func tokenWithAddress(token market.Token) string {
	var str string
	str += token.Name

	addrLen := len(token.Address)
	if addrLen >= 9 {
		// first 5 + last 4
		str += fmt.Sprintf(" (%s...%s)", token.Address[0:6], token.Address[addrLen-4:addrLen])
	}
	return str
}

func milliToDate(milliseconds int64) string {
	y, m, d := time.UnixMilli(milliseconds).Date()

	return fmt.Sprintf("%s %d, %d", m.String(), d, y)
}

func (ct *Table) SetRowsTokens(tokens []market.Token) {
	ct.table = ct.setTokensRows(tokens)
}
func (ct *Table) ShowTokensCol() {
	ct.table = ct.table.WithColumns(tokensColumns)
}

func (ct *Table) SetRowsPools(pools []market.Pool) {
	ct.table = ct.setPoolsRows(pools)
}
func (ct *Table) ShowPoolsCol() {

	ct.table = ct.table.WithColumns(poolsColumns)
}
