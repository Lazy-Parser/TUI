package page_default

import (
	"strings"
	"tui/internal/tui/components/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorEmpty    = "   "
	cursorSelected = "👉 "
)

type option struct {
	title       string
	description string
}

type modelBody struct {
	options  []option
	selected int
}

func newBody() modelBody {
	return modelBody{
		options: []option{
			{title: "Generate", description: "Create pairs and tokens from MEXC / Dexscreener"},
			{title: "Listen prices", description: "Make money button 💸"},
			{title: "View database", description: "View generated pairs and tokens"},
			{title: "View logs", description: "For nerds 🤓 (MacOS only)"},
			{title: "Exit", description: "Are you sure?"},
		},
	}
}

func (m modelBody) Init() tea.Cmd {
	return nil
}

func (m modelBody) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		// move up
		case tea.KeyUp:
			if m.selected > 0 {
				m.selected--
			} else {
				m.selected = len(m.options) - 1
			}

		// move down
		case tea.KeyDown:
			if m.selected < len(m.options)-1 {
				m.selected++
			} else {
				m.selected = 0
			}

		// select option
		case tea.KeyEnter:
			switch m.selected {
			case 0:
				// Generate
			case 1:
				// Listen prices
			case 2:
				// View database
				return m, SelectPage("Database")
			case 3:
				// View logs
				return m, OpenNewTerminal()
			case 4:
				// Exit
				return m, tea.Quit
			}
		}
	}

	// select option
	return m, nil
}

func (m modelBody) View() string {
	var strBuilder strings.Builder
	for i, option := range m.options {
		cursor := cursorEmpty
		if m.selected == i {
			cursor = cursorSelected
		}

		// if selected
		strBuilder.WriteString(cursor)
		if cursor != cursorEmpty {
			// if exit selected
			if m.selected == len(m.options)-1 {
				strBuilder.WriteString(theme.SelectedDangerOptionTextStyle.Render(option.title))
			} else {
				strBuilder.WriteString(theme.SelectedOptionTextStyle.Render(option.title))
			}
		} else {
			strBuilder.WriteString(theme.UnselectedOptionTextStyle.Render(option.title))
		}
		strBuilder.WriteString("\n")
		strBuilder.WriteString(cursorEmpty + theme.FadedTextStyle.Render(option.description))
		strBuilder.WriteString("\n\n")
	}

	list := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Render(strBuilder.String())

	return list
}
