package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the main UI application
type App struct {
	model *Model
	done  bool
}

// NewApp creates a new UI application
func NewApp() *App {
	return &App{
		model: NewModel(),
		done:  false,
	}
}

// Run starts the UI application
func (a *App) Run() error {
	p := tea.NewProgram(a.model, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// GetModel returns the current model
func (a *App) GetModel() *Model {
	return a.model
}
