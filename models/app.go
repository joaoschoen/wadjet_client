package models

import (
	"wadjet_client/models/modules"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	sidebar SidebarModel
	content tea.Model
	width   int
	height  int
}

func NewAppModel() AppModel {
	return AppModel{
		sidebar: NewSidebar(),
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Wadjet")
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.sidebar.focus = true
			// return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// debug_log.LogToFile(fmt.Sprintf("app width:%d", msg.Width))
		// debug_log.LogToFile(fmt.Sprintf("app height:%d", msg.Height))
	}

	model, cmd := m.sidebar.Update(msg)
	m.sidebar = model.(SidebarModel)
	switch m.sidebar.selected {
	case 0:
		m.content = modules.NewInfo()
	}
	if m.content != nil {
		model, cmd := m.content.Update(msg)
		m.content = model
		return m, cmd
	}
	return m, cmd
}

func (m AppModel) View() string {
	// Create main content area style
	// mainStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	pos := lipgloss.Position(m.width / 2)
	appBorder := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Width(m.width - 2).Height(m.height - 2).
		BorderForeground(lipgloss.Color("#FFD700")).
		AlignHorizontal(pos)
	// content := mainStyle.Render("wadjet lives")
	var content string
	if m.content != nil {
		content = lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), m.content.View())
	} else {
		content = lipgloss.JoinHorizontal(lipgloss.Top, m.sidebar.View(), "loading")
	}

	// Join them horizontally
	return appBorder.Render(content)
}
