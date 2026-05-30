package ui

import (
	"github.com/charmbracelet/bubbletea"
)

// Model represents the main application state
type Model struct {
	currentView ViewType
	inputModel  *InputModel
	resultModel *ResultModel
}

// ViewType represents different views in the app
type ViewType int

const (
	InputView ViewType = iota
	ResultView
)

// NewModel creates a new model for the application
func NewModel() *Model {
	return &Model{
		currentView: InputView,
		inputModel:  NewInputModel(),
		resultModel: NewResultModel(),
	}
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// Switch between views
			if m.currentView == InputView {
				m.currentView = ResultView
			} else {
				m.currentView = InputView
			}
			return m, nil
		}
	}

	// Update the current view
	if m.currentView == InputView {
		input, cmd := m.inputModel.Update(msg)
		m.inputModel = input.(*InputModel)
		return m, cmd
	} else {
		result, cmd := m.resultModel.Update(msg)
		m.resultModel = result.(*ResultModel)
		return m, cmd
	}
}

// View renders the current view
func (m *Model) View() string {
	if m.currentView == InputView {
		return m.inputModel.View()
	} else {
		return m.resultModel.View()
	}
}
