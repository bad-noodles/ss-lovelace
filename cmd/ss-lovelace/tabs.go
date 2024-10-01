package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bad-noodles/ss-lovelace/pkg/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var tabStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(theme.Foreground).
	Background(theme.Background).
	PaddingLeft(2).
	PaddingRight(2)
var selectedTabStyle = tabStyle.Background(lipgloss.Color("#7D56F4"))

var contentStyle = lipgloss.NewStyle()

type Tab struct {
	Title   string
	Content tea.Model
}
type Tabs struct {
	tabs     []Tab
	selected int
}

func NewTabs(tabs []Tab) (t Tabs) {
	t.tabs = tabs
	t.selected = 0
	return
}

func (t Tabs) Init() tea.Cmd {
	inits := []tea.Cmd{}

	for _, tab := range t.tabs {
		inits = append(inits, tab.Content.Init())
	}
	return tea.Batch(inits...)
}

func (t Tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		msg.Height = msg.Height - tabStyle.GetHeight() - 1
		contentStyle = contentStyle.Height(msg.Height).Width(msg.Width)
	case tea.KeyMsg:
		num, err := strconv.Atoi(msg.String())

		if err == nil && num <= len(t.tabs) && num > 0 {
			t.selected = num - 1

			return t, nil
		}
	}

	cmds := []tea.Cmd{}
	for i, tab := range t.tabs {
		content, cmd := tab.Content.Update(msg)

		cmds = append(cmds, cmd)
		t.tabs[i].Content = content
	}

	return t, tea.Batch(cmds...)
}

func (t Tabs) View() string {
	var output strings.Builder

	output.WriteString(contentStyle.Render(t.tabs[t.selected].Content.View()))
	output.WriteString("\n")

	for i, tab := range t.tabs {
		s := tabStyle

		if i == t.selected {
			s = selectedTabStyle
		}

		if i != 0 {
			output.WriteString("|")
		}

		output.WriteString(s.Render(fmt.Sprintf("(%d) %s", i+1, tab.Title)))

	}

	return output.String()
}
