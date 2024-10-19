package progress

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// NewProgram initializes and returns a new Bubble Tea program with the progress model.
func NewProgram(total int) *tea.Program {
	model := model{
		progressBar: progress.New(progress.WithDefaultGradient()),
		total:       total,
		current:     0,
	}
	return tea.NewProgram(model)
}

type ProgressMsg struct{}

// Model represents the state of the progress bar.
type model struct {
	progressBar progress.Model
	total       int
	current     int
}

// Init initializes the model. No initialization required for this example.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// Allow quitting by pressing any key
		return m, tea.Quit
	default:
		// Continue updating the progress bar until the total is reached
		if m.current < m.total {
			m.current++
			progressPercent := float64(m.current) / float64(m.total)
			return m, m.progressBar.SetPercent(progressPercent)
		}
	}
	return m, nil
}

// View renders the current state of the progress bar.
func (m model) View() string {
	return fmt.Sprintf("\n%s\n\nPress any key to exit.", m.progressBar.View())
}

// IncrementProgress triggers the increment by sending an incrementMsg
func IncrementProgress() tea.Cmd {
	return func() tea.Msg {
		return ProgressMsg{}
	}
}
