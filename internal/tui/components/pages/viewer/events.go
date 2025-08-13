package page_viewer

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func onTokensMsg(m *mainView, msg *tokensMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		log.Println(fmt.Errorf("failed to load tokens: %w", msg.err))
		return m, nil
	}

	m.tokens = msg.tokens
	cmd := setTokenAmount(len(msg.tokens))

	// tablePayload := make([]table.Row, len(msg.tokens))
	// for _, token := range msg.tokens {
	// 	tablePayload = append(tablePayload, table.Row{token.Name, string(token.Decimal), token.Address, token.Image_url, string(token.WithdrawFee), time.UnixMilli(token.CreateTime).String(), token.Network})
	// }
	// updateTable := func() tea.Cmd {
	// 	return func() tea.Msg {
	// 		m.table.SetRows(tablePayload)
	// 		return nil
	// 	}
	// }

	return m, tea.Batch(cmd)
}
