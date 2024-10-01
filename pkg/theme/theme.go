package theme

import "github.com/charmbracelet/lipgloss"

const (
	Accent     = lipgloss.Color("#7D56F4")
	Foreground = lipgloss.Color("#FAFAFA")
	Background = lipgloss.Color("#333333")
	Error      = lipgloss.Color("#a9597a")
	Success    = lipgloss.Color("#68c5b5")
)

var Bold = lipgloss.NewStyle().Bold(true)
