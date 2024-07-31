package tui

import (
	"fmt"

	"dominguezdev.com/cli/auth"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	username textinput.Model
	password textinput.Model
	err      error
	success  bool
}

func initialModel() model {
	m := model{}

	m.username = textinput.New()
	m.username.Placeholder = "Username"
	m.username.Focus()

	m.password = textinput.New()
	m.password.Placeholder = "Password"
	m.password.EchoMode = textinput.EchoPassword

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			if m.username.Focused() {
				m.username.Blur()
				m.password.Focus()
			} else {
				m.password.Blur()
				m.username.Focus()
			}
		case tea.KeyEnter:
			m.err = auth.Login(m.username.Value(), m.password.Value())
			if m.err == nil {
				m.success = true
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.username, cmd = m.username.Update(msg)
	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m model) View() string {
	if m.success {
		return "You successfully logged in!"
	}
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\n", m.err) + m.username.View() + "\n" + m.password.View()
	}
	return m.username.View() + "\n" + m.password.View()
}

func RunTUI() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
