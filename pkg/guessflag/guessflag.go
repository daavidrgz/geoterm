package guessflag

import (
	"os"
	"os/exec"

	"fmt"

	ct "geoterm/internal/country"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput *textinput.Model
	Incorrect bool
}

func restoreDefaultModel() {
	m.textInput.TextStyle = lipgloss.NewStyle().Foreground(Foreground)
	m.textInput.PromptStyle = lipgloss.NewStyle().Foreground(Foreground)
	m.textInput.SetValue("")
	m.Incorrect = false
}

func setIncorrectModel() {
	m.textInput.TextStyle = lipgloss.NewStyle().Background(Red)
	m.textInput.PromptStyle = lipgloss.NewStyle().Foreground(Red)
	m.Incorrect = true
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Guess the country"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.TextStyle = lipgloss.NewStyle().Foreground(Foreground)

	return model{textInput: &ti, Incorrect: false}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(0)
		}
	}

	*m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.textInput.View(),
		"Press Esc to quit",
	)
}

var incorrect = 0
var correct = 0
var m = initialModel()

func LaunchGame() {
	totalCountries := InitFlagSystem()

	for len(countries) > 0 {
		hideCursor()
		clearScreen()

		RenderTitle()
		RenderCounter(correct, incorrect, totalCountries)
		ShowFlag()

		tea.NewProgram(m).Start()

		if ct.MatchesName(GetCurrentCountry(), m.textInput.Value()) {
			if !m.Incorrect {
				correct++
			}
			restoreDefaultModel()
			NextCountry()
		} else {
			if !m.Incorrect {
				incorrect++
				setIncorrectModel()
			}
		}
	}
}

func hideCursor() {
	fmt.Printf("\x1b[?25l")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
