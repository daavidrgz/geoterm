package guessflag

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	titleContent = ` ██████╗ ██╗   ██╗███████╗███████╗███████╗    ████████╗██╗  ██╗███████╗    ███████╗██╗      █████╗  ██████╗ 
██╔════╝ ██║   ██║██╔════╝██╔════╝██╔════╝    ╚══██╔══╝██║  ██║██╔════╝    ██╔════╝██║     ██╔══██╗██╔════╝ 
██║  ███╗██║   ██║█████╗  ███████╗███████╗       ██║   ███████║█████╗      █████╗  ██║     ███████║██║  ███╗
██║   ██║██║   ██║██╔══╝  ╚════██║╚════██║       ██║   ██╔══██║██╔══╝      ██╔══╝  ██║     ██╔══██║██║   ██║
╚██████╔╝╚██████╔╝███████╗███████║███████║       ██║   ██║  ██║███████╗    ██║     ███████╗██║  ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚══════╝╚══════╝╚══════╝       ╚═╝   ╚═╝  ╚═╝╚══════╝    ╚═╝     ╚══════╝╚═╝  ╚═╝ ╚═════╝`
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(0, 2, 1, 2).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(Foreground).
		Foreground(Foreground).
		SetString(titleContent)
)

func RenderTitle() {
	width := GetWidth()
	title := lipgloss.PlaceHorizontal(width, lipgloss.Center, titleStyle.String())
	fmt.Println(title + "\n\n")
}
