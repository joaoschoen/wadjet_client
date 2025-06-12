package modules

import (
	"fmt"
	"wadjet_client/models/ui"
	"wadjet_client/shared/debug_log"
	"wadjet_client/shared/format"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InfoModel struct {
	eye    ui.EyeModel
	width  int
	height int
}

func NewInfo() InfoModel {
	return InfoModel{
		eye: ui.NewEye(),
	}
}

func (m InfoModel) Init() tea.Cmd {
	return nil
}

func (m InfoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			// return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width - 14
		m.height = msg.Height - 2
		debug_log.LogToFile(fmt.Sprintf("app width:%d", msg.Width))
		debug_log.LogToFile(fmt.Sprintf("content max width:%d", m.width))
		// debug_log.LogToFile(fmt.Sprintf("app height:%d", msg.Height))
	}

	return m, nil
}

func (m InfoModel) View() string {
	eye := m.eye.View()

	// debug_log.LogToFile(fmt.Sprintf("eye width:%d", ))

	// eye =
	// 	debug_log.LogToFile("\n" + eye)
	// if lipgloss.Width(eye) > m.width && m.width > 0 {
	// }
	content := lipgloss.JoinVertical(lipgloss.Center, format.TruncateHorizontal(eye, m.width), "test")
	return content
}
