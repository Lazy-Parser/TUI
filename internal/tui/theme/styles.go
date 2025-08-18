package theme

import "github.com/charmbracelet/lipgloss"

var (
	SelectedOptionTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000")).
				Bold(true)

	SelectedDangerOptionTextStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#ff383f")).
					Bold(true)

	UnselectedOptionTextStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#4A4A4A")).
					Bold(false)

	// for description in selection menu
	FadedTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#777777"))
)
