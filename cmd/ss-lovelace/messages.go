package main

import (
	gameMsg "github.com/bad-noodles/ss-lovelace/pkg/message"
	"github.com/bad-noodles/ss-lovelace/pkg/theme"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type message struct {
	Subject string
	Body    string
	Read    bool
}

type messages struct {
	messageChannel  chan gameMsg.Message
	messages        []message
	cursor          int
	viewport        viewport.Model
	viewportFocused bool
}

func newMessages(messageChannel chan gameMsg.Message) messages {
	return messages{
		messageChannel: messageChannel,
		messages:       []message{},
		viewport:       viewport.New(10, 10),
	}
}

func (m messages) listen() tea.Cmd {
	return func() tea.Msg {
		return <-m.messageChannel
	}
}

func (m messages) Init() tea.Cmd {
	return m.listen()
}

func (m messages) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case gameMsg.Message:
		m.messages = append(m.messages, message{
			Subject: msg.Subject,
			Body:    msg.Body,
			Read:    false,
		})
		if len(m.messages) == 1 {
			m.viewport.SetContent(msg.Body)
			m.messages[0].Read = true
		}
		return m, m.listen()
	case tea.WindowSizeMsg:
		contentWidth = msg.Width - itemStyle.GetWidth() - 8
		if contentWidth > maxContendWidth {
			contentWidth = maxContendWidth
		}

		listStyle = listStyle.Height(msg.Height - listStyle.GetBorderTopSize() - listStyle.GetBorderBottomSize() - 1)
		bodyStyle = bodyStyle.Height(msg.Height - bodyStyle.GetBorderTopSize() - bodyStyle.GetBorderBottomSize() - 1)

		m.viewport.Width = contentWidth
		m.viewport.Height = bodyStyle.GetHeight()

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "h", "left", "l", "right":
			m.viewportFocused = !m.viewportFocused
			return m, nil
		}
		if !m.viewportFocused {
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.messages)-1 {
					m.cursor++
				}
			}
			m.viewport.SetContent(m.messages[m.cursor].Body)
			m.messages[m.cursor].Read = true
			return m, nil
		}
	}

	vp, cmd := m.viewport.Update(msg)
	m.viewport = vp

	return m, cmd
}

const maxContendWidth = 150

var (
	contentWidth = maxContendWidth
	bodyStyle    = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 2)
	listStyle    = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	itemStyle    = lipgloss.NewStyle().
			Foreground(theme.Foreground).
			PaddingLeft(2).
			PaddingRight(2).
			Width(50)
	selectedItemStyle = itemStyle.Background(theme.Accent)
	header            = selectedItemStyle.
				Padding(1).
				Align(lipgloss.Center).
				Border(lipgloss.NormalBorder(), false, false, true).
				Bold(true).
				Render("📨 Space Mail")
)

func (m messages) View() string {
	list := []string{header}
	for i, message := range m.messages {
		s := itemStyle

		if i == m.cursor {
			s = selectedItemStyle
		}

		if !message.Read {
			s = s.Bold(true)
		}

		list = append(list, s.Render(message.Subject))
	}

	if m.viewportFocused {
		bodyStyle = bodyStyle.BorderForeground(theme.Accent)
		listStyle = listStyle.BorderForeground(theme.Foreground)
	} else {
		bodyStyle = bodyStyle.BorderForeground(theme.Foreground)
		listStyle = listStyle.BorderForeground(theme.Accent)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		listStyle.Render(lipgloss.JoinVertical(lipgloss.Left, list...)),
		bodyStyle.Render(m.viewport.View()),
	)
}
