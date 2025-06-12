package ui

import (
	"wadjet_client/shared/ascii"

	tea "github.com/charmbracelet/bubbletea"
)

type EyeModel struct {
}

func NewEye() EyeModel {
	return EyeModel{}
}

func (m EyeModel) Init() tea.Cmd {
	return nil
}

func (m EyeModel) Update(msg tea.Msg) (EyeModel, tea.Cmd) {
	return m, nil
}

func (m EyeModel) View() string {
	return ascii.Eye()
}
