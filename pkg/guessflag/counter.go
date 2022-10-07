package guessflag

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	successCounter = lipgloss.NewStyle().
			Bold(true).
			Foreground(Green)
	errorCounter = lipgloss.NewStyle().
			Bold(true).
			Foreground(Red)
)

func RenderCounter(correct, incorrect, total int) {
	width := GetWidth()
	successCounter = successCounter.SetString(fmt.Sprintf("Correct: %d/%d", correct, total))
	errorCounter = errorCounter.SetString(fmt.Sprintf("Incorrect: %d/%d", incorrect, total))

	counterContainer := lipgloss.JoinHorizontal(
		lipgloss.Center,
		successCounter.String(),
		Divider(2),
		errorCounter.String())

	counterContainer = lipgloss.PlaceHorizontal(
		width,
		lipgloss.Center,
		counterContainer)

	fmt.Println(counterContainer + "\n\n")
}
