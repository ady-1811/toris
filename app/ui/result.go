package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/toris/ai"
)

// ResultModel represents the result view showing command suggestions
type ResultModel struct {
	result     *ai.CommandResponse
	loading    bool
	error      string
	osName     string
	useCommand bool
}

// NewResultModel creates a new result model
func NewResultModel() *ResultModel {
	return &ResultModel{
		loading:    false,
		error:      "",
		useCommand: false,
	}
}

// Init initializes the result model
func (m *ResultModel) Init() tea.Cmd {
	return nil
}

// Update handles result view events
func (m *ResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.useCommand = true
			return m, nil
		case "r":
			// Reset and go back to input
			m.result = nil
			m.error = ""
			m.useCommand = false
			return m, nil
		}
	}
	return m, nil
}

// View renders the result view
func (m *ResultModel) View() string {
	var view strings.Builder

	view.WriteString("\n╔═══════════════════════════════════════════════════════════╗\n")
	view.WriteString("║  TORIS - Command Result                                    ║\n")
	view.WriteString("╚═══════════════════════════════════════════════════════════╝\n\n")

	if m.loading {
		view.WriteString("⏳ Loading... Please wait\n")
		return view.String()
	}

	if m.error != "" {
		view.WriteString("❌ Error: " + m.error + "\n")
		view.WriteString("\n🔑 Press 'r' to try again\n")
		return view.String()
	}

	if m.result == nil {
		view.WriteString("No results yet. Switch back to input and submit a command.\n")
		view.WriteString("\n🔑 Press 'Tab' to go back to input\n")
		return view.String()
	}

	// Display OS info
	view.WriteString(fmt.Sprintf("💻 OS: %s\n\n", m.osName))

	// Display suggested command
	view.WriteString("📋 Suggested Command:\n")
	view.WriteString("┌─────────────────────────────────────────────────────────────┐\n")
	view.WriteString("│ " + m.result.Command + "\n")
	view.WriteString("└─────────────────────────────────────────────────────────────┘\n\n")

	// Display confidence
	view.WriteString(fmt.Sprintf("🎯 Confidence: %.1f%%\n\n", m.result.Confidence*100))

	// Display instructions
	if len(m.result.Instruction) > 0 {
		view.WriteString("📝 Instructions:\n")
		for i, instruction := range m.result.Instruction {
			view.WriteString(fmt.Sprintf("  %d. %s\n", i+1, instruction))
		}
		view.WriteString("\n")
	}

	// Display action buttons
	view.WriteString("🔑 Keybinds:\n")
	view.WriteString("  • Enter    - Execute command\n")
	view.WriteString("  • Tab      - Back to input\n")
	view.WriteString("  • r        - Try again\n")
	view.WriteString("  • Ctrl+C   - Quit\n")

	if m.useCommand {
		view.WriteString("\n✅ Command executed!\n")
	}

	return view.String()
}

// SetResult sets the command result
func (m *ResultModel) SetResult(result *ai.CommandResponse, osName string) {
	m.result = result
	m.osName = osName
	m.error = ""
	m.loading = false
}

// SetError sets an error message
func (m *ResultModel) SetError(err string) {
	m.error = err
	m.result = nil
	m.loading = false
}

// SetLoading sets the loading state
func (m *ResultModel) SetLoading(loading bool) {
	m.loading = loading
}

// GetResult returns the current result
func (m *ResultModel) GetResult() *ai.CommandResponse {
	return m.result
}
