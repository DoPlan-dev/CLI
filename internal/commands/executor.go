package commands

import (
	"os"

	"github.com/DoPlan-dev/CLI/internal/tui"
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// TUICommandExecutor implements the CommandExecutor interface for the TUI
type TUICommandExecutor struct{}

// NewTUICommandExecutor creates a new command executor for the TUI
func NewTUICommandExecutor() tui.CommandExecutor {
	return &TUICommandExecutor{}
}

func (e *TUICommandExecutor) RunDevServer() error {
	return RunDevServer()
}

func (e *TUICommandExecutor) UndoLastAction() error {
	return UndoLastAction()
}

func (e *TUICommandExecutor) CreateNewProject() error {
	return CreateNewProject()
}

func (e *TUICommandExecutor) DeployProject() error {
	return DeployProject()
}

func (e *TUICommandExecutor) PublishPackage() error {
	return PublishPackage()
}

func (e *TUICommandExecutor) RunSecurityScan() error {
	return RunSecurityScan()
}

func (e *TUICommandExecutor) AutoFix() error {
	return AutoFix()
}

func (e *TUICommandExecutor) DiscussIdea() error {
	return DiscussIdea()
}

func (e *TUICommandExecutor) GenerateDocuments() error {
	return GenerateDocuments()
}

func (e *TUICommandExecutor) CreatePlan() error {
	return CreatePlan()
}

func (e *TUICommandExecutor) UpdateProgress() error {
	return UpdateProgress()
}

func (e *TUICommandExecutor) ManageAPIKeys() error {
	return ManageAPIKeys()
}

func (e *TUICommandExecutor) ApplyDesign() error {
	return ApplyDesign()
}

func (e *TUICommandExecutor) SetupIntegration() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}
	return wizard.RunIntegrationWizard(projectRoot)
}
