package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func RunTUI() {
	// Set up logging to a file
	f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Starting TUI")

	p := tea.NewProgram(initialModel())
	_, err = p.Run()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}
