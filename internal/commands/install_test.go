package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/DoPlan-dev/CLI/test/helpers"
)

func TestNewInstallCommand(t *testing.T) {
	cmd := NewInstallCommand()
	assert.NotNil(t, cmd)
	assert.Equal(t, "install", cmd.Use)
	assert.NotEmpty(t, cmd.Short)
	assert.NotEmpty(t, cmd.Long)
}

func TestNewInstaller(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	ide := "cursor"

	installer := NewInstaller(projectRoot, ide)
	assert.NotNil(t, installer)
	assert.Equal(t, projectRoot, installer.projectRoot)
	assert.Equal(t, ide, installer.ide)
}

func TestInstaller_Install(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.Install()
	require.NoError(t, err)

	// Check directories were created
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "commands"))
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "rules"))
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "config"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "contracts"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "templates"))

	// Check config was created
	configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
	assert.FileExists(t, configPath)

	// Check README was created
	readmePath := filepath.Join(projectRoot, "README.md")
	assert.FileExists(t, readmePath)

	// Check dashboard was created
	dashboardPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, dashboardPath)
}

func TestInstaller_Install_DifferentIDEs(t *testing.T) {
	ides := []string{"cursor", "gemini", "claude", "codex", "opencode", "qwen"}

	for _, ide := range ides {
		t.Run(ide, func(t *testing.T) {
			projectRoot := helpers.CreateTempProject(t)
			installer := NewInstaller(projectRoot, ide)

			err := installer.Install()
			require.NoError(t, err)

			// Check IDE-specific directories
			switch ide {
			case "gemini":
				assert.DirExists(t, filepath.Join(projectRoot, ".gemini", "commands"))
			case "claude":
				assert.DirExists(t, filepath.Join(projectRoot, ".claude", "commands"))
			case "codex":
				assert.DirExists(t, filepath.Join(projectRoot, ".codex", "prompts"))
			case "opencode":
				assert.DirExists(t, filepath.Join(projectRoot, ".opencode", "command"))
			case "qwen":
				assert.DirExists(t, filepath.Join(projectRoot, ".qwen", "commands"))
			}
		})
	}
}

func TestInstaller_createDirectories(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.createDirectories()
	require.NoError(t, err)

	// Check common directories
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "commands"))
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "rules"))
	assert.DirExists(t, filepath.Join(projectRoot, ".cursor", "config"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "contracts"))
	assert.DirExists(t, filepath.Join(projectRoot, "doplan", "templates"))
}

func TestInstaller_installCommands(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	// Create directory first
	err := installer.createDirectories()
	require.NoError(t, err)

	err = installer.installCommands()
	require.NoError(t, err)

	// Check command files were created
	commandsDir := filepath.Join(projectRoot, ".cursor", "commands")
	assert.FileExists(t, filepath.Join(commandsDir, "discuss.md"))
	assert.FileExists(t, filepath.Join(commandsDir, "refine.md"))
	assert.FileExists(t, filepath.Join(commandsDir, "generate.md"))
	assert.FileExists(t, filepath.Join(commandsDir, "plan.md"))
	assert.FileExists(t, filepath.Join(commandsDir, "dashboard.md"))
}

func TestInstaller_installGeminiCommands(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "gemini")

	// Create directory first
	err := installer.createDirectories()
	require.NoError(t, err)

	err = installer.installGeminiCommands()
	require.NoError(t, err)

	commandsDir := filepath.Join(projectRoot, ".gemini", "commands")
	assert.FileExists(t, filepath.Join(commandsDir, "discuss.toml"))
	assert.FileExists(t, filepath.Join(commandsDir, "refine.toml"))
}

func TestInstaller_installClaudeCommands(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "claude")

	// Create directory first
	err := installer.createDirectories()
	require.NoError(t, err)

	err = installer.installClaudeCommands()
	require.NoError(t, err)

	commandsDir := filepath.Join(projectRoot, ".claude", "commands")
	assert.FileExists(t, filepath.Join(commandsDir, "discuss.md"))
	assert.FileExists(t, filepath.Join(commandsDir, "refine.md"))
}

func TestInstaller_createTemplates(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.createTemplates()
	require.NoError(t, err)

	templatesDir := filepath.Join(projectRoot, "doplan", "templates")
	assert.FileExists(t, filepath.Join(templatesDir, "plan-template.md"))
	assert.FileExists(t, filepath.Join(templatesDir, "design-template.md"))
	assert.FileExists(t, filepath.Join(templatesDir, "tasks-template.md"))
}

func TestInstaller_generateRules(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.generateRules()
	require.NoError(t, err)

	rulesDir := filepath.Join(projectRoot, ".cursor", "rules")
	assert.FileExists(t, filepath.Join(rulesDir, "workflow-rules.md"))
	assert.FileExists(t, filepath.Join(rulesDir, "github-rules.md"))
}

func TestInstaller_generateContext(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.generateContext()
	require.NoError(t, err)

	contextPath := filepath.Join(projectRoot, "CONTEXT.md")
	assert.FileExists(t, contextPath)
}

func TestInstaller_generateConfig(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.generateConfig()
	require.NoError(t, err)

	configPath := filepath.Join(projectRoot, ".cursor", "config", "doplan-config.json")
	assert.FileExists(t, configPath)
}

func TestInstaller_createREADME(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	err := installer.createREADME()
	require.NoError(t, err)

	readmePath := filepath.Join(projectRoot, "README.md")
	assert.FileExists(t, readmePath)

	content, err := os.ReadFile(readmePath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "DoPlan")
}

func TestInstaller_createDashboard(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	installer := NewInstaller(projectRoot, "cursor")

	// Create doplan directory first
	err := installer.createDirectories()
	require.NoError(t, err)

	err = installer.createDashboard()
	require.NoError(t, err)

	dashboardPath := filepath.Join(projectRoot, "doplan", "dashboard.md")
	assert.FileExists(t, dashboardPath)
}

