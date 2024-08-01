package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type splashModel struct {
	choices []string
	cursor  int
	width   int
	height  int
}

func initialModel() splashModel {
	return splashModel{
		choices: []string{"Login", "Exit"},
	}
}

func (m splashModel) Init() tea.Cmd {
	return nil
}

func (m splashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.choices[m.cursor] {
			case "Login":
				return initialLoginModel(m.width, m.height), nil
			case "Exit":
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m splashModel) View() string {
	trulyTerrifyingLogo := `
 +++                                +++                          ...................                
++++++        +++++++++++++        +++++                  ..........:--=-=-----:.............       
+++++++   +++++++++++++++++++++   +++++++         ............:=+=----------------==-............   
+++++++++++++++++++++++++++++++++++++++++       ......-++:.-=-----------------+++++---=+=---+:....  
+++++++++++++++           +++++++++++++++       ....=----==-=-......:+-----+:...  ..----=+=---=...  
++++++++++++                 ++++++++++++       ...*--%@*--+...     ..-----.::..    .:=---%=---...  
++++++++++                     ++++++++++       ...*--##--=.+@@*.     :----+@@#..   ..=---=---=...  
++++++++                         ++++++++       ....=-*----:*@**.     ----=.%@=.    ..=----++.....  
+++++++        +++++              +++++++       .....:=-----:...   ...=--=-+..    ..-------=:....   
++++++          ++++++             +++++           ..:=------=-:..:-=-+@@@@@-==--==---------=...    
++++++           +++++             ++++++         ...---------------=-:-++-::+--------------=...    
++++++            ++++++           ++++++        ....=--------------=:::-=::::=--------------...    
++++++           +++++++           ++++++        ....=----------------=.-:.+=----------------...    
++++++          ++++++++++         ++++++         ...-----------------=.--.+-----------------:..    
++++++         +++++ ++++++        +++++=          ..:--------------------------------------=:..    
 ++++++       +++++    +++++      ++++++           ..:=-------------------------------------=:..    
 +++++++                         ++++++            ..:=-------------------------------------=:..    
  +++++++                       ++++++          ......--------------------------------------=:..... 
    +++++++                   +++++++          .......--------------------------------------=:......
     +++++++++             +++++++++           ...=:::--------------------------------------=-::=...
       +++++++++++++++++++++++++++             ..:=::==--------------------------------------+::=:..
         +++++++++++++++++++++++                                                                    
             +++++++++++++++                                                                        
`

	titleStyle := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("62"))

	title := titleStyle.Render("Welcome to DDN Gopher")

	var options []string
	for i, choice := range m.choices {
		cursor := " " // No cursor
		if m.cursor == i {
			cursor = ">" // Cursor
		}
		options = append(options, cursor+" "+choice)
	}

	optionsStr := lipgloss.JoinVertical(lipgloss.Left, options...)

	content := lipgloss.JoinVertical(lipgloss.Center, trulyTerrifyingLogo, title, optionsStr)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
