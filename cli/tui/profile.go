package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type profileModel struct {
	token  string
	width  int
	height int
}

func newProfileModel(token string) profileModel {
	return profileModel{token: token}
}

func (m profileModel) Init() tea.Cmd {
	return nil
}

func (m profileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m profileModel) View() string {
	titleStyle := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("62"))

	contentStyle := lipgloss.NewStyle().
		Width(m.width-4).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		MarginTop(2)

	title := titleStyle.Render("Profile")
	content := contentStyle.Render(fmt.Sprintf("You successfully logged in! Your token: %s", m.token))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, title+"\n\n"+content)
}
