package main

// A simple program that queries and displays the window-size.

import (
	"log"
	"wadjet_client/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(models.NewAppModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
