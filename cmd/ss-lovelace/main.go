package main

import (
	"github.com/bad-noodles/ss-lovelace/pkg/game"
	tea "github.com/charmbracelet/bubbletea"
)

type root struct {
	tabs tea.Model
}

func newRoot(g *game.Game) root {
	return root{
		tabs: NewTabs([]Tab{
			{Title: "S-Mail", Content: newMessages(g.MessageChannel)},
			{Title: "Dashboard", Content: NewDashboard(g.ModulesChannel)},
		}),
	}
}

func (r root) Init() tea.Cmd {
	return tea.Batch(r.tabs.Init())
}

func (r root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d":
			return r, tea.Quit
		}
	}

	tabs, cmd := r.tabs.Update(msg)
	r.tabs = tabs

	return r, cmd
}

func (r root) View() string {
	return r.tabs.View()
}

func main() {
	g := game.NewGame()
	p := tea.NewProgram(newRoot(g), tea.WithAltScreen())

	g.Start()
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
