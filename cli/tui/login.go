package tui

import (
	"fmt"
	"log"

	"dominguezdev.com/cli/auth"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type loginState int

const (
	loginScreen loginState = iota
	profileScreen
)

type loginModel struct {
	username textinput.Model
	password textinput.Model
	err      error
	token    string
	state    loginState
	width    int
	height   int
}

func initialLoginModel(width, height int) loginModel {
	m := loginModel{
		state:  loginScreen,
		width:  width,
		height: height,
	}

	m.username = textinput.New()
	m.username.Placeholder = "Username"
	m.username.Focus()

	m.password = textinput.New()
	m.password.Placeholder = "Password"
	m.password.EchoMode = textinput.EchoPassword

	return m
}

func (m loginModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tea.EnterAltScreen)
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			if m.state == loginScreen {
				if m.username.Focused() {
					m.username.Blur()
					m.password.Focus()
				} else {
					m.password.Blur()
					m.username.Focus()
				}
			}
		case tea.KeyEnter:
			if m.state == loginScreen {
				log.Printf("Username: %s, Password: %s\n", m.username.Value(), m.password.Value())
				var token string
				token, m.err = auth.Login(m.username.Value(), m.password.Value())
				if m.err == nil {
					m.token = token
					return newProfileModel(token).Update(msg)
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.username, cmd = m.username.Update(msg)
	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m loginModel) View() string {
	switch m.state {
	case loginScreen:
		return m.loginView()
	case profileScreen:
		return newProfileModel(m.token).View()
	default:
		return "Unknown state!"
	}
}

func (m loginModel) loginView() string {
	titleStyle := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("62"))

	preambleStyle := lipgloss.NewStyle().
		Width(m.width).
		Padding(1, 0, 1, 0).
		Align(lipgloss.Center)

	boxStyle := lipgloss.NewStyle().
		Width(m.width-4).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		MarginTop(2)

	title := titleStyle.Render("Welcome to DDN Gopher")
	preamble := preambleStyle.Render("Please enter your credentials to log in:")
	inputs := m.username.View() + "\n" + m.password.View()

	content := lipgloss.JoinVertical(lipgloss.Center, title, preamble, boxStyle.Render(inputs))

	if m.err != nil {
		content = fmt.Sprintf("Error: %v\n\n", m.err) + content
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
