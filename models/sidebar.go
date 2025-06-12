package models

import (
	"wadjet_client/models/ui"
	"wadjet_client/shared/colors"
	"wadjet_client/shared/format"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SidebarModel struct {
	options       []string
	selected      int
	styleNormal   lipgloss.Style
	styleSelected lipgloss.Style
	height        int
	focus         bool
}

func NewSidebar() SidebarModel {
	styleNormal := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colors.GOLD))
	styleSelected := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colors.BLUE))
	return SidebarModel{
		options:       []string{"Info", "Listening", "Register", "Exit"},
		selected:      0,
		styleNormal:   styleNormal,
		styleSelected: styleSelected,
		focus:         true,
	}
}

func (m SidebarModel) Init() tea.Cmd {
	return nil
}

func (m SidebarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			if m.focus {
				if m.selected > 0 {
					m.selected--
				}
			}
		case "down", "s":
			// debug_log.LogToFile(fmt.Sprintf("focus:%t", m.focus))
			if m.focus {
				if m.selected < len(m.options)-1 {
					m.selected++
				}
			}
		case "enter":
			if m.focus {
				m.focus = false
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height - 2
	}

	return m, nil
}

func (m SidebarModel) View() string {
	BorderRight := ui.BorderRight(m.height)
	menuOptions := ""
	for i, opt := range m.options {
		cursor := " "
		if i == m.selected {
			if m.focus {
				cursor = ">"
			}
			menuOptions += m.styleSelected.Render(cursor + opt)
		} else {
			menuOptions += m.styleNormal.Render(cursor + opt)
		}
		menuOptions += "\n"
	}
	content := format.CenterVertical(menuOptions, m.height)
	return BorderRight.Render(content)
}
