package progress

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	bubbletea "github.com/charmbracelet/bubbletea"
)

func NewProgram(endpoints []string) *bubbletea.Program {
	model := model{
		progressBar: progress.New(progress.WithScaledGradient("#FF7F50", "#00FF7F")),
		endpoints:   endpoints,
		current:     0,
	}
	return bubbletea.NewProgram(model)
}

type model struct {
	progressBar progress.Model
	endpoints   []string
	current     int
}

func (m model) Init() bubbletea.Cmd {
	return nil
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	// Handle user input and update progress bar
	if m.current < len(m.endpoints) {
		m.progressBar.SetPercent(float64(m.current) / float64(len(m.endpoints)))
		m.current++
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Processing %s\n%s", m.endpoints[m.current], m.progressBar.View())
}
