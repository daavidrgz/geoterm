package guessflag

import (
	"os"
	"os/exec"

	"fmt"
	"log"

	ct "geoterm/internal/country"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput *textinput.Model
	exit      *bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Guess the country"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: &ti,
		exit:      new(bool),
	}
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
			return m, tea.Batch(tea.Quit)
		case tea.KeyCtrlC, tea.KeyEsc:
			*m.exit = true
			return m, tea.Batch(tea.Quit)
		}
	}

	*m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.textInput.View(),
		"Press q to quit",
	)
}

func LaunchGame() {
	InitFlagSystem()

	for len(countries) > 0 {
		fmt.Printf("\x1b[?25l")
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("\x1b[?25l")

		ShowFlag()

		model := initialModel()

		p := tea.NewProgram(model)
		if err := p.Start(); err != nil {
			log.Fatal(err)
		}

		if *model.exit {
			break
		}

		if ct.MatchesName(GetCurrentCountry(), model.textInput.Value()) {
			fmt.Println("Correct!")
			NextCountry()
		} else {
			os.Exit(1)
			fmt.Println("Incorrect!")
		}
	}
}
