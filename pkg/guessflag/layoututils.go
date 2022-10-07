package guessflag

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	Foreground = lipgloss.AdaptiveColor{Light: "#06283D", Dark: "#DFF6FF"}
	Background = lipgloss.AdaptiveColor{Light: "#DFF6FF", Dark: "#06283D"}
	Green      = lipgloss.Color("#B6E388")
	Red        = lipgloss.Color("#F96666")
)

func Divider(padding int) string {
	return lipgloss.NewStyle().
		Foreground(Foreground).
		SetString("•").
		Padding(0, padding).
		String()
}

func GetWidth() int {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	return physicalWidth
}
