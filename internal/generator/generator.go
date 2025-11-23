package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/doplan/cli/internal/utils"
	"github.com/doplan/cli/pkg/models"
)

// Generator represents a component that generates part of the project
type Generator interface {
	// Generate generates the component and returns an error if it fails
	Generate(request *models.ProjectRequest, projectPath string) error
	// Name returns the name of the generator for logging
	Name() string
}

// GenerationStep represents a step in the generation pipeline
type GenerationStep struct {
	Generator Generator
	Name      string
}

// GenerationContext holds the context for project generation
type GenerationContext struct {
	Request      *models.ProjectRequest
	ProjectPath  string
	CreatedDirs  []string
	CreatedFiles []string
}

// Orchestrate orchestrates the entire project generation process.
// It validates the request, creates the project directory, and runs all generators in order.
// Returns an error if generation fails, and attempts rollback on critical errors.
//
// Cross-platform considerations:
// - Uses filepath.Join for path construction (handles Windows/Unix differences)
// - Uses os.MkdirAll for directory creation (works on all platforms)
// - Atomic file writes ensure data integrity across platforms
// - Path sanitization prevents path traversal attacks
func Orchestrate(request *models.ProjectRequest) error {
	// Comprehensive input validation
	if err := request.Validate(); err != nil {
		return fmt.Errorf("invalid project request: %w", err)
	}

	// Additional validation: IDE must be supported
	if !models.IsValidIDE(request.IDE) {
		return fmt.Errorf("unsupported IDE: %s (supported: %v)", request.IDE, models.GetSupportedIDEs())
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Create project path (filepath.Join handles platform-specific separators)
	projectPath := filepath.Join(cwd, request.ProjectName)

	// Validate and sanitize project path
	sanitizedPath, err := utils.ValidatePath(projectPath)
	if err != nil {
		return fmt.Errorf("invalid project path: %w", err)
	}
	projectPath = sanitizedPath

	// Check if project directory already exists
	if utils.PathExists(projectPath) {
		return fmt.Errorf("directory '%s' already exists. Please choose a different name or remove the existing directory", request.ProjectName)
	}

	// Check permissions with detailed error messages
	if err := utils.CheckPermissions(projectPath); err != nil {
		return fmt.Errorf("insufficient permissions to create project at '%s': %w. Please check directory permissions and try again", projectPath, err)
	}

	// Create generation context
	ctx := &GenerationContext{
		Request:      request,
		ProjectPath:  projectPath,
		CreatedDirs:  []string{},
		CreatedFiles: []string{},
	}

	// Create project root directory
	if err := utils.CreateDirectory(projectPath); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}
	ctx.CreatedDirs = append(ctx.CreatedDirs, projectPath)

	// Define generation pipeline
	steps := []GenerationStep{
		{Generator: &AgentsGenerator{}, Name: "AI Agents"},
		{Generator: &CommandsGenerator{}, Name: "Commands"},
		{Generator: &RulesGenerator{}, Name: "Rules Library"},
		{Generator: &PlanGenerator{}, Name: ".plan Structure"},
		{Generator: &GitHubGenerator{}, Name: "GitHub Workflows"},
		{Generator: &IDEGenerator{}, Name: "IDE Configs"},
		{Generator: &BoilerplateGenerator{}, Name: "Boilerplate"},
		{Generator: &DocsGenerator{}, Name: "Documentation"},
		// {Generator: &DirectoryGenerator{}, Name: "Directory Structure"}, // Placeholder
	}

	// Track generation start time for performance monitoring
	startTime := time.Now()

	// Execute generation pipeline with progress tracking
	for i, step := range steps {
		if step.Generator == nil {
			// Skip placeholder steps
			continue
		}

		// Track step start time
		stepStartTime := time.Now()

		// Generate component with enhanced error context
		if err := step.Generator.Generate(request, projectPath); err != nil {
			// Provide detailed error context
			stepDuration := time.Since(stepStartTime)
			errorMsg := fmt.Errorf("generation failed at step %d/%d (%s) after %v: %w",
				i+1, len(steps), step.Name, stepDuration, err)

			// Rollback on error
			rollbackErr := rollback(ctx)
			if rollbackErr != nil {
				return fmt.Errorf("%v (rollback also failed: %v)", errorMsg, rollbackErr)
			}
			return errorMsg
		}

		// Log step completion (could be enhanced with actual logging)
		stepDuration := time.Since(stepStartTime)
		_ = stepDuration // Suppress unused variable warning for now
	}

	// Log total generation time
	totalDuration := time.Since(startTime)
	_ = totalDuration // Suppress unused variable warning for now

	// Performance check: warn if generation takes too long
	if totalDuration > 5*time.Second {
		// Could add logging here: fmt.Printf("Warning: Generation took %v (target: <5s)\n", totalDuration)
		_ = totalDuration
	}

	return nil
}

// rollback attempts to remove all created files and directories in reverse order
func rollback(ctx *GenerationContext) error {
	var errors []error

	// Remove files first (in reverse order)
	for i := len(ctx.CreatedFiles) - 1; i >= 0; i-- {
		if err := os.Remove(ctx.CreatedFiles[i]); err != nil && !os.IsNotExist(err) {
			errors = append(errors, fmt.Errorf("failed to remove file %s: %w", ctx.CreatedFiles[i], err))
		}
	}

	// Remove directories (in reverse order)
	for i := len(ctx.CreatedDirs) - 1; i >= 0; i-- {
		if err := os.RemoveAll(ctx.CreatedDirs[i]); err != nil && !os.IsNotExist(err) {
			errors = append(errors, fmt.Errorf("failed to remove directory %s: %w", ctx.CreatedDirs[i], err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("rollback completed with errors: %v", errors)
	}

	return nil
}

// trackCreatedDir adds a directory to the context's created directories list
func trackCreatedDir(ctx *GenerationContext, dir string) {
	ctx.CreatedDirs = append(ctx.CreatedDirs, dir)
}

// trackCreatedFile adds a file to the context's created files list
func trackCreatedFile(ctx *GenerationContext, file string) {
	ctx.CreatedFiles = append(ctx.CreatedFiles, file)
}
