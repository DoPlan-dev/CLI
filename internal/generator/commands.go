package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/doplan/cli/internal/utils"
	"github.com/doplan/cli/pkg/models"
)

// Command represents a slash command in the AI agency system
type Command struct {
	// Basic Information
	Name        string // Command name (e.g., "tell", "build")
	Trigger     string // Trigger pattern (e.g., "/tell or /tell <idea>")
	Description string // Brief description
	Action      string // Detailed action description

	// Agent Involvement
	AgentInvolvement []string // List of agents involved

	// Files
	FilesRead    []string // Files read by this command
	FilesModified []string // Files modified by this command

	// Additional Information
	Examples      []string // Example usage
	GitHubAutomation string // GitHub automation details (if applicable)
}

// GetAllCommands returns all core and squad commands
func GetAllCommands() []Command {
	return []Command{
		// Core Commands (11)
		{
			Name:        "tell",
			Trigger:     "/tell or /tell <idea>",
			Description: "Capture project idea",
			Action: `When user types /tell or /tell <idea>:

1. **Capture the idea**: If idea is provided inline, save it. Otherwise, prompt user for their project idea.
2. **Save to IDEA.md**: Write the idea to .plan/00_System/IDEA.md
3. **Activate Project Orchestrator**: The Project Orchestrator analyzes the idea and activates appropriate agents.
4. **Response**: "Idea captured! Your project idea has been saved. Type /improve to brainstorm with the team."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Product Manager",
			},
			FilesRead:    []string{},
			FilesModified: []string{
				".plan/00_System/IDEA.md",
				".plan/active_state.json",
			},
			Examples: []string{
				"/tell",
				"/tell Build a todo app",
			},
		},
		{
			Name:        "improve",
			Trigger:     "/improve",
			Description: "Team brainstorm session",
			Action: `When user types /improve:

1. **Activate All Level 1 Managers**: Product Manager, Engineering Lead, Design & UX Manager, QA & Reliability Manager, Release & Growth Manager, Documentation Lead
2. **Brainstorm Session**: Each manager provides ideas and improvements
3. **Update BRAINSTORM.md**: Write all brainstormed ideas to .plan/00_System/BRAINSTORM.md
4. **Response**: "Brainstorm complete! Review BRAINSTORM.md. Type /write to generate planning documents."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"Design & UX Manager",
				"QA & Reliability Manager",
				"Release & Growth Manager",
				"Documentation Lead",
			},
			FilesRead:    []string{".plan/00_System/IDEA.md"},
			FilesModified: []string{
				".plan/00_System/BRAINSTORM.md",
			},
			Examples: []string{
				"/improve",
			},
		},
		{
			Name:        "team",
			Trigger:     "/team",
			Description: "Show active agents and hierarchy",
			Action: `When user types /team:

1. **Load Agent Definitions**: Read all agent files from .cursor/agents/
2. **Display Hierarchy**: Show the hierarchical structure of all agents
3. **Show Roles**: Display each agent's role and responsibilities
4. **Response**: Display formatted agent hierarchy and roles`,
			AgentInvolvement: []string{},
			FilesRead: []string{
				".cursor/agents/*.md",
			},
			FilesModified: []string{},
			Examples: []string{
				"/team",
			},
		},
		{
			Name:        "write",
			Trigger:     "/write",
			Description: "Generate PRD + ARCHITECTURE + DESIGN_SYSTEM",
			Action: `When user types /write:

1. **Generate PRD.md**: Product Manager creates comprehensive Product Requirements Document
2. **Generate ARCHITECTURE.md**: Engineering Lead and System Architect create technical architecture
3. **Generate DESIGN_SYSTEM.md**: Design Manager and UI/UX Designer create design system
4. **Save All Files**: Write to .plan/00_System/
5. **Response**: "Documents generated! Review PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md. Type /change to edit any document, or /good to approve."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Engineering Lead",
				"System Architect",
				"Design & UX Manager",
				"UI/UX Designer",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/IDEA.md",
				".plan/00_System/BRAINSTORM.md",
			},
			FilesModified: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/active_state.json",
			},
			Examples: []string{
				"/write",
			},
		},
		{
			Name:        "change",
			Trigger:     "/change <document> <change>",
			Description: "Edit any document",
			Action: `When user types /change <document> <change>:

1. **Parse Command**: Extract document name and change description
2. **Load Document**: Read the specified document from .plan/00_System/
3. **Apply Changes**: Update the document with the requested changes
4. **Save Document**: Write updated document back to file
5. **Response**: "Document updated! Changes saved to [document].md"`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/*.md",
			},
			FilesModified: []string{
				".plan/00_System/*.md",
			},
			Examples: []string{
				"/change prd Add dark mode",
				"/change architecture Use PostgreSQL instead of MySQL",
			},
		},
		{
			Name:        "good",
			Trigger:     "/good",
			Description: "Approve & lock plan",
			Action: `When user types /good:

1. **Validate Documents**: Ensure PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md exist
2. **Lock Plan**: Set locked: true in .plan/active_state.json
3. **Update Phase**: Set phase to "approved" in active_state.json
4. **Response**: "Plan approved and locked! Type /tasks to generate implementation tasks."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/active_state.json",
			},
			Examples: []string{
				"/good",
			},
		},
		{
			Name:        "tasks",
			Trigger:     "/tasks",
			Description: "Generate TASKS.md",
			Action: `When user types /tasks:

1. **Read Approved Plan**: Load PRD.md, ARCHITECTURE.md, and DESIGN_SYSTEM.md
2. **Generate Tasks**: Create implementation tasks organized by phases
3. **Create TASKS.md**: Write tasks to .plan/TASKS.md
4. **Update State**: Set phase to "tasks" in active_state.json
5. **Response**: "Tasks generated! Review .plan/TASKS.md. Type /build to start coding."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Engineering Lead",
				"Product Manager",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/DESIGN_SYSTEM.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			Examples: []string{
				"/tasks",
			},
		},
		{
			Name:        "load",
			Trigger:     "/load <path>",
			Description: "Inject context into AI agents",
			Action: `When user types /load <path>:

1. **Parse Path**: Extract file or directory path
2. **Load Content**: Read the specified file or files from directory
3. **Inject Context**: Add content to agent context for current session
4. **Response**: "Context loaded! [path] is now available to agents."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".cursor/rules/library/**",
				".plan/**",
			},
			FilesModified: []string{},
			Examples: []string{
				"/load @library/04-frameworks/frontend/nextjs.md",
				"/load .plan/00_System/PRD.md",
			},
		},
		{
			Name:        "build",
			Trigger:     "/build or /build <task_id>",
			Description: "Start coding next task",
			Action: `When user types /build or /build <task_id>:

1. **Determine Task**: 
   - If task_id provided, load that task
   - Otherwise, find next uncompleted task from TASKS.md
2. **Load Task Context**: Read task details, dependencies, and related code
3. **Activate Relevant Agents**: Activate agents needed for the task (Frontend Lead, Backend Lead, etc.)
4. **Start Implementation**: Begin coding the task with full context
5. **Update State**: Set active_task in .plan/active_state.json
6. **Response**: "Building task [ID]: [Description]. Focus on this task only."`,
			AgentInvolvement: []string{
				"Engineering Lead",
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/active_state.json",
			},
			Examples: []string{
				"/build",
				"/build 1.2",
				"/build 3",
			},
			GitHubAutomation: `After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md`,
		},
		{
			Name:        "progress",
			Trigger:     "/progress",
			Description: "Show current progress",
			Action: `When user types /progress:

1. **Read TASKS.md**: Load all tasks
2. **Read active_state.json**: Get completed tasks and current phase
3. **Calculate Progress**: 
   - Total tasks
   - Completed tasks
   - In progress tasks
   - Percentage complete
4. **Display Progress**: Show formatted progress report:
   - Phase: [current phase]
   - Tasks: X/Y completed (Z%)
   - Current task: [active task]
   - Next up: [next task]
5. **Response**: Display progress summary`,
			AgentInvolvement: []string{
				"Project Orchestrator",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			FilesModified: []string{},
			Examples: []string{
				"/progress",
			},
		},
		{
			Name:        "finished",
			Trigger:     "/finished",
			Description: "Mark current task done",
			Action: `When user types /finished:

1. **Mark Task Complete**: Update task status in TASKS.md
2. **Update State**: Remove active_task and add to completed in active_state.json
3. **Auto-commit**: Commit changes with conventional commit format
4. **Auto-push**: Push to current branch
5. **Update CHANGELOG**: Update CHANGELOG.md if significant changes
6. **Response**: "Task completed! Changes committed and pushed. Type /build to start next task."`,
			AgentInvolvement: []string{
				"Project Orchestrator",
				"Release Captain",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
			},
			FilesModified: []string{
				".plan/TASKS.md",
				".plan/active_state.json",
				"CHANGELOG.md",
			},
			Examples: []string{
				"/finished",
			},
			GitHubAutomation: `After task completion, the system will:
- Auto-commit changes with conventional commit format
- Auto-push to current branch (feature/bugfix/hotfix)
- Update CHANGELOG.md if significant changes
- Follow branching strategy from @library/01-core-workflow/github-workflow-automation.md`,
		},
		// Squad Commands
		{
			Name:        "secure",
			Trigger:     "/secure",
			Description: "Security review",
			Action: `When user types /secure:

1. **Security Review**: Security Lead conducts security review
2. **Vulnerability Assessment**: Identify and document security vulnerabilities
3. **Generate Report**: Create security review report
4. **Response**: "Security review complete! Review security findings."`,
			AgentInvolvement: []string{
				"Security Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/SECURITY.md",
			},
			Examples: []string{
				"/secure",
			},
		},
		{
			Name:        "roles",
			Trigger:     "/roles",
			Description: "Design RBAC system",
			Action: `When user types /roles:

1. **Design RBAC**: Security Lead and Backend Lead design role-based access control
2. **Define Roles**: Create role definitions and permissions
3. **Generate Documentation**: Document RBAC system
4. **Response**: "RBAC system designed! Review role definitions."`,
			AgentInvolvement: []string{
				"Security Lead",
				"Backend Lead",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/RBAC.md",
			},
			Examples: []string{
				"/roles",
			},
		},
		{
			Name:        "money",
			Trigger:     "/money",
			Description: "Billing & payment setup",
			Action: `When user types /money:

1. **Payment System Design**: Design billing and payment system
2. **Integration Planning**: Plan payment gateway integration
3. **Generate Documentation**: Document payment system architecture
4. **Response**: "Payment system designed! Review billing architecture."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Backend Lead",
			},
			FilesRead: []string{
				".plan/00_System/PRD.md",
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/BILLING.md",
			},
			Examples: []string{
				"/money",
			},
		},
		{
			Name:        "pretty",
			Trigger:     "/pretty",
			Description: "UI/UX improvements",
			Action: `When user types /pretty:

1. **UI/UX Review**: Design Manager and UI/UX Designer review interface
2. **Improvement Suggestions**: Provide UI/UX improvement recommendations
3. **Update Design System**: Update DESIGN_SYSTEM.md with improvements
4. **Response**: "UI/UX improvements suggested! Review design updates."`,
			AgentInvolvement: []string{
				"Design & UX Manager",
				"UI/UX Designer",
			},
			FilesRead: []string{
				".plan/00_System/DESIGN_SYSTEM.md",
				"src/**",
			},
			FilesModified: []string{
				".plan/00_System/DESIGN_SYSTEM.md",
			},
			Examples: []string{
				"/pretty",
			},
		},
		{
			Name:        "seo",
			Trigger:     "/seo",
			Description: "SEO optimization",
			Action: `When user types /seo:

1. **SEO Analysis**: Analyze current SEO implementation
2. **Optimization Recommendations**: Provide SEO optimization suggestions
3. **Generate SEO Plan**: Create SEO optimization plan
4. **Response**: "SEO analysis complete! Review SEO recommendations."`,
			AgentInvolvement: []string{
				"Product Manager",
				"Frontend Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/PRD.md",
			},
			FilesModified: []string{
				".plan/00_System/SEO.md",
			},
			Examples: []string{
				"/seo",
			},
		},
		{
			Name:        "ship",
			Trigger:     "/ship",
			Description: "Release management",
			Action: `When user types /ship:

1. **Release Planning**: Release Captain plans the release
2. **Version Management**: Manage version numbers and semantic versioning
3. **Release Notes**: Generate release notes
4. **Deployment Planning**: Plan deployment strategy
5. **Response**: "Release planned! Review release notes and deployment plan."`,
			AgentInvolvement: []string{
				"Release Captain",
				"Release & Growth Manager",
			},
			FilesRead: []string{
				".plan/TASKS.md",
				"CHANGELOG.md",
			},
			FilesModified: []string{
				"CHANGELOG.md",
				".plan/00_System/RELEASE.md",
			},
			Examples: []string{
				"/ship",
			},
		},
		{
			Name:        "safe",
			Trigger:     "/safe",
			Description: "Security audit",
			Action: `When user types /safe:

1. **Security Audit**: Security Lead conducts comprehensive security audit
2. **Vulnerability Scanning**: Scan for security vulnerabilities
3. **Compliance Check**: Verify compliance with security standards
4. **Generate Audit Report**: Create security audit report
5. **Response**: "Security audit complete! Review audit findings."`,
			AgentInvolvement: []string{
				"Security Lead",
			},
			FilesRead: []string{
				"src/**",
				".plan/00_System/ARCHITECTURE.md",
				".plan/00_System/SECURITY.md",
			},
			FilesModified: []string{
				".plan/00_System/SECURITY_AUDIT.md",
			},
			Examples: []string{
				"/safe",
			},
		},
		{
			Name:        "cheap",
			Trigger:     "/cheap",
			Description: "Cost optimization",
			Action: `When user types /cheap:

1. **Cost Analysis**: Analyze current infrastructure and service costs
2. **Optimization Recommendations**: Provide cost optimization suggestions
3. **Generate Cost Plan**: Create cost optimization plan
4. **Response**: "Cost analysis complete! Review optimization recommendations."`,
			AgentInvolvement: []string{
				"DevOps Engineer",
				"Performance Engineer",
			},
			FilesRead: []string{
				".plan/00_System/ARCHITECTURE.md",
			},
			FilesModified: []string{
				".plan/00_System/COST_OPTIMIZATION.md",
			},
			Examples: []string{
				"/cheap",
			},
		},
	}
}

// GetCommandByName returns a command by name, or nil if not found
func GetCommandByName(name string) *Command {
	commands := GetAllCommands()
	for i := range commands {
		if commands[i].Name == name {
			return &commands[i]
		}
	}
	return nil
}

// GetCoreCommands returns only the 11 core commands
func GetCoreCommands() []Command {
	allCommands := GetAllCommands()
	return allCommands[:11] // First 11 are core commands
}

// GetSquadCommands returns only the squad-specific commands
func GetSquadCommands() []Command {
	allCommands := GetAllCommands()
	return allCommands[11:] // Remaining are squad commands
}

// commandTemplate is the markdown template for command files
const commandTemplate = `# /{{.Name}}

## Trigger
{{.Trigger}}{{if .Examples}}

## Examples
{{range .Examples}}- {{.}}
{{end}}{{end}}

## Action
{{.Action}}

## Agent Involvement
{{range .AgentInvolvement}}- **{{.}}**
{{end}}{{if .FilesRead}}
## Files Read
{{range .FilesRead}}- {{.}}
{{end}}{{end}}{{if .FilesModified}}
## Files Modified
{{range .FilesModified}}- {{.}}
{{end}}{{end}}{{if .GitHubAutomation}}
## GitHub Automation
{{.GitHubAutomation}}{{end}}
`

// RenderCommandMarkdown renders a command to markdown format
func RenderCommandMarkdown(cmd *Command) (string, error) {
	if cmd == nil {
		return "", fmt.Errorf("command cannot be nil")
	}

	tmpl, err := template.New("command").Parse(commandTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cmd); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// CommandsGenerator generates the command markdown files.
type CommandsGenerator struct{}

// Name returns the name of the generator.
func (g *CommandsGenerator) Name() string {
	return "Commands"
}

// Generate creates the .cursor/commands directory and all command markdown files.
func (g *CommandsGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	commandsDir := filepath.Join(projectPath, ".cursor", "commands")

	// Create .cursor/commands directory
	if err := utils.CreateDirectory(commandsDir); err != nil {
		return fmt.Errorf("failed to create .cursor/commands directory: %w", err)
	}

	// Get all commands
	commands := GetAllCommands()

	// Generate each command file
	for _, cmd := range commands {
		// Render command markdown
		markdown, err := RenderCommandMarkdown(&cmd)
		if err != nil {
			return fmt.Errorf("failed to render command %s: %w", cmd.Name, err)
		}

		// Write command file
		commandPath := filepath.Join(commandsDir, cmd.Name+".md")
		if err := utils.WriteFile(commandPath, []byte(markdown)); err != nil {
			return fmt.Errorf("failed to write command file %s: %w", cmd.Name+".md", err)
		}
	}

	return nil
}

// GenerateCommands is a convenience function that creates a CommandsGenerator and generates commands
func GenerateCommands(request *models.ProjectRequest, projectPath string) error {
	generator := &CommandsGenerator{}
	return generator.Generate(request, projectPath)
}

