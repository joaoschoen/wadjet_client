package ui

import (
	"wadjet_client/shared/colors"

	"github.com/charmbracelet/lipgloss"
)

func BorderRight(height int) lipgloss.Style {
	doubleRightBorder := lipgloss.Border{
		Right: "║", // double vertical line
	}

	style := lipgloss.NewStyle().
		Border(doubleRightBorder).
		BorderForeground(lipgloss.Color(colors.GOLD)).
		BorderTop(false).BorderBottom(false).BorderLeft(false).
		Height(height)
	return style
}
