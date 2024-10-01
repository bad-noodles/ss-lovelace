package main

import (
	"fmt"

	"github.com/bad-noodles/ss-lovelace/pkg/ship/modules"
	"github.com/bad-noodles/ss-lovelace/pkg/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dashboard struct {
	modulesChannel chan []modules.ModuleDescriptor
	modules        []modules.ModuleDescriptor
}

func NewDashboard(modulesChannel chan []modules.ModuleDescriptor) Dashboard {
	return Dashboard{modulesChannel, []modules.ModuleDescriptor{}}
}

func (d Dashboard) listen() tea.Cmd {
	return func() tea.Msg {
		return <-d.modulesChannel
	}
}

func (d Dashboard) Init() tea.Cmd {
	return d.listen()
}

func (d Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case []modules.ModuleDescriptor:
		d.modules = msg
		return d, d.listen()
	}
	return d, nil
}

var errorModuleStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Width(40).
	Height(4).
	Padding(1, 2).
	Margin(1, 2).
	BorderForeground(theme.Error)
var successModuleStyle = errorModuleStyle.BorderForeground(theme.Success)

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Width(11)
	successStatus = lipgloss.NewStyle().Foreground(theme.Success)
	errorStatus   = lipgloss.NewStyle().Foreground(theme.Error)
)

func (d Dashboard) View() string {
	mods := []string{}

	for _, m := range d.modules {
		modStyle := errorModuleStyle
		status := errorStatus.Render("Disconnected")
		if m.Connected {
			status = errorStatus.Render("Unhealthy")
		}
		if m.Health {
			modStyle = successModuleStyle
			status = successStatus.Render("Healthy")
		}
		mods = append(mods,
			modStyle.Render(lipgloss.JoinVertical(
				lipgloss.Left,
				fmt.Sprintf("%s%s", titleStyle.Render("Module"), m.Name),
				fmt.Sprintf("%s%d", titleStyle.Render("Port"), m.Port),
				fmt.Sprintf("%s%s", titleStyle.Render("Status"), status),
			),
			))
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, mods...)
}
