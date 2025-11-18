package tui

// CommandExecutor defines the interface for executing commands from the TUI
// This avoids import cycles by having the TUI define the interface
// and the root command wire up the actual implementations
type CommandExecutor interface {
	RunDevServer() error
	UndoLastAction() error
	CreateNewProject() error
	DeployProject() error
	PublishPackage() error
	RunSecurityScan() error
	AutoFix() error
	DiscussIdea() error
	GenerateDocuments() error
	CreatePlan() error
	UpdateProgress() error
	ManageAPIKeys() error
	ApplyDesign() error
	SetupIntegration() error
}

// DefaultExecutor is a no-op executor that returns "not implemented" errors
type DefaultExecutor struct{}

func (e *DefaultExecutor) RunDevServer() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) UndoLastAction() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) CreateNewProject() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) DeployProject() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) PublishPackage() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) RunSecurityScan() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) AutoFix() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) DiscussIdea() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) GenerateDocuments() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) CreatePlan() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) UpdateProgress() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) ManageAPIKeys() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) ApplyDesign() error {
	return nil // Will be implemented
}

func (e *DefaultExecutor) SetupIntegration() error {
	return nil // Will be implemented
}

