package ui

import (
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// InputModel represents the input view where users describe commands
type InputModel struct {
	input        string
	cursor       int
	focused      bool
	submitted    bool
	lastInput    string
}

// NewInputModel creates a new input model
func NewInputModel() *InputModel {
	return &InputModel{
		input:     "",
		cursor:    0,
		focused:   true,
		submitted: false,
	}
}

// Init initializes the input model
func (m *InputModel) Init() tea.Cmd {
	return nil
}

// Update handles input events
func (m *InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.submitted = true
			m.lastInput = strings.TrimSpace(m.input)
			return m, nil
		case tea.KeyBackspace:
			if m.cursor > 0 {
				m.input = m.input[:m.cursor-1] + m.input[m.cursor:]
				m.cursor--
			}
		case tea.KeyLeft:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight:
			if m.cursor < len(m.input) {
				m.cursor++
			}
		case tea.KeyHome:
			m.cursor = 0
		case tea.KeyEnd:
			m.cursor = len(m.input)
		default:
			if msg.Type == tea.KeyRunes {
				m.input = m.input[:m.cursor] + string(msg.Runes) + m.input[m.cursor:]
				m.cursor += len(msg.Runes)
			}
		}
	}
	return m, nil
}

// View renders the input view
func (m *InputModel) View() string {
	var view strings.Builder

	view.WriteString("\n╔═══════════════════════════════════════════════════════════╗\n")
	view.WriteString("║  TORIS - Terminal Organized and Rational IntelliSense    ║\n")
	view.WriteString("╚═══════════════════════════════════════════════════════════╝\n\n")

	view.WriteString("📝 Describe what you want to do:\n\n")

	// Render input box
	view.WriteString("┌─────────────────────────────────────────────────────────────┐\n")
	view.WriteString("│ ")
	view.WriteString(m.input)
	if m.cursor == len(m.input) {
		view.WriteString("█")
	}
	padding := 59 - len(m.input)
	if padding > 0 {
		view.WriteString(strings.Repeat(" ", padding))
	}
	view.WriteString("│\n")
	view.WriteString("└─────────────────────────────────────────────────────────────┘\n\n")

	view.WriteString("🔑 Keybinds:\n")
	view.WriteString("  • Enter    - Submit command\n")
	view.WriteString("  • Tab      - Switch to results\n")
	view.WriteString("  • Ctrl+C   - Quit\n")

	if m.submitted && m.lastInput != "" {
		view.WriteString("\n✅ Submitted: " + m.lastInput + "\n")
	}

	return view.String()
}

// IsSubmitted returns whether the input has been submitted
func (m *InputModel) IsSubmitted() bool {
	return m.submitted
}

// GetInput returns the current input value
func (m *InputModel) GetInput() string {
	return m.lastInput
}

// Reset clears the input
func (m *InputModel) Reset() {
	m.input = ""
	m.cursor = 0
	m.submitted = false
	m.lastInput = ""
}
