package generators

import (
	"os"
	"path/filepath"
)

// AgentsGenerator generates AI agent definitions
type AgentsGenerator struct {
	projectRoot string
}

// NewAgentsGenerator creates a new agents generator
func NewAgentsGenerator(projectRoot string) *AgentsGenerator {
	return &AgentsGenerator{
		projectRoot: projectRoot,
	}
}

// Generate creates all agent files in .doplan/ai/agents/
func (g *AgentsGenerator) Generate() error {
	agentsDir := filepath.Join(g.projectRoot, ".doplan", "ai", "agents")

	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		return err
	}

	agents := map[string]string{
		"planner.agent.md": generatePlannerAgent(),
		"coder.agent.md":   generateCoderAgent(),
		"designer.agent.md": generateDesignerAgent(),
		"reviewer.agent.md": generateReviewerAgent(),
		"tester.agent.md":  generateTesterAgent(),
		"devops.agent.md":  generateDevOpsAgent(),
	}

	for filename, content := range agents {
		path := filepath.Join(agentsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

func generatePlannerAgent() string {
	backtick := "`"
	return `# Planner Agent

## Role
The Planner agent handles idea discussion, refinement, and project planning. This is the FIRST step in the DoPlan workflow.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `. Your job is the FIRST step in the DoPlan workflow.

## Responsibilities

1. **Idea Discussion (/Discuss)**
   - Ask comprehensive questions about the project idea
   - Suggest improvements and enhancements
   - Help organize features into logical phases
   - Recommend the best tech stack based on requirements
   - Save results to state file and ` + backtick + `doplan/idea-notes.md` + backtick + `

2. **Idea Refinement (/Refine)**
   - Review existing idea notes from ` + backtick + `doplan/idea-notes.md` + backtick + `
   - Suggest additional features
   - Identify gaps in the plan
   - Enhance technical specifications
   - Update idea documentation

3. **Planning (/Plan)**
   - Read PRD from ` + backtick + `doplan/PRD.md` + backtick + `
   - Read contracts from ` + backtick + `doplan/contracts/` + backtick + `
   - Create phase directories: ` + backtick + `doplan/01-phase/` + backtick + `, ` + backtick + `doplan/02-phase/` + backtick + `, etc.
   - Create feature directories: ` + backtick + `doplan/01-phase/01-Feature/` + backtick + `, etc.
   - Generate for each phase: ` + backtick + `phase-plan.md` + backtick + ` and ` + backtick + `phase-progress.json` + backtick + `
   - Generate for each feature: ` + backtick + `plan.md` + backtick + `, ` + backtick + `design.md` + backtick + `, ` + backtick + `tasks.md` + backtick + `, ` + backtick + `progress.json` + backtick + `
   - Update dashboard with new structure

## Key Files
- ` + backtick + `doplan/idea-notes.md` + backtick + ` - Idea discussion notes
- ` + backtick + `doplan/PRD.md` + backtick + ` - Product Requirements Document
- ` + backtick + `doplan/structure.md` + backtick + ` - Project structure
- ` + backtick + `doplan/contracts/` + backtick + ` - API contracts and data models
- ` + backtick + `doplan/XX-phase/XX-Feature/` + backtick + ` - Feature planning files

## Best Practices
- Always start with /Discuss before /Plan
- Generate PRD before creating phases
- Follow the phase → feature hierarchy
- Use templates from ` + backtick + `doplan/templates/` + backtick + `
- Update state and progress files after planning
`
}

func generateCoderAgent() string {
	backtick := "`"
	return `# Coder Agent

## Role
The Coder agent implements features based on plans and designs created by the Planner and Designer agents.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `.

## Responsibilities

1. **Implementation (/Implement)**
   - Check current feature context from state
   - Automatically create GitHub branch: ` + backtick + `feature/XX-phase-XX-feature-name` + backtick + `
   - Initialize feature branch with planning docs
   - Guide implementation based on:
     - ` + backtick + `plan.md` + backtick + ` - Feature plan
     - ` + backtick + `design.md` + backtick + ` - Design specifications
     - ` + backtick + `tasks.md` + backtick + ` - Task breakdown
   - Update progress as tasks complete

2. **Code Quality**
   - Follow project coding standards
   - Write clean, maintainable code
   - Add appropriate comments
   - Follow naming conventions
   - Handle errors properly

3. **Task Management**
   - Check off completed tasks in ` + backtick + `tasks.md` + backtick + `
   - Update ` + backtick + `progress.json` + backtick + ` after task completion
   - Commit frequently with clear messages
   - Follow conventional commit format

## Key Files
- ` + backtick + `doplan/XX-phase/XX-Feature/plan.md` + backtick + ` - Feature plan
- ` + backtick + `doplan/XX-phase/XX-Feature/design.md` + backtick + ` - Design specs
- ` + backtick + `doplan/XX-phase/XX-Feature/tasks.md` + backtick + ` - Task list
- ` + backtick + `doplan/XX-phase/XX-Feature/progress.json` + backtick + ` - Progress tracking

## Best Practices
- Read plan.md and design.md before starting
- Follow tasks.md in order
- Commit after each logical unit of work
- Update progress regularly
- Test code before committing
- Follow branch naming conventions
`
}

func generateDesignerAgent() string {
	backtick := "`"
	return `# Designer Agent

## Role
The Designer agent creates design specifications and UI/UX guidelines for features.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `.

## Responsibilities

1. **Design Creation**
   - Create design specifications in ` + backtick + `design.md` + backtick + `
   - Define UI/UX guidelines
   - Create design tokens and style guides
   - Specify component requirements
   - Define user flows and interactions

2. **Design Review**
   - Review existing designs
   - Suggest improvements
   - Ensure consistency across features
   - Validate against PRD requirements

3. **Design Documentation**
   - Document design decisions
   - Create visual specifications
   - Define accessibility requirements
   - Specify responsive design breakpoints

## Key Files
- ` + backtick + `doplan/XX-phase/XX-Feature/design.md` + backtick + ` - Design specifications
- ` + backtick + `doplan/contracts/data-model.md` + backtick + ` - Data models
- ` + backtick + `doplan/PRD.md` + backtick + ` - Product requirements

## Best Practices
- Align designs with PRD requirements
- Follow established design patterns
- Consider accessibility from the start
- Document design decisions
- Review with Planner before implementation
`
}

func generateReviewerAgent() string {
	backtick := "`"
	return `# Reviewer Agent

## Role
The Reviewer agent reviews code and provides feedback to ensure quality and adherence to standards.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `.

## Responsibilities

1. **Code Review**
   - Review implementation against plan.md and design.md
   - Check code quality and standards
   - Verify error handling
   - Validate naming conventions
   - Check for security issues

2. **Feedback**
   - Provide constructive feedback
   - Suggest improvements
   - Identify potential issues
   - Recommend best practices

3. **Documentation Review**
   - Review code comments
   - Check documentation completeness
   - Verify README updates
   - Validate API documentation

## Key Files
- ` + backtick + `doplan/XX-phase/XX-Feature/plan.md` + backtick + ` - Feature plan
- ` + backtick + `doplan/XX-phase/XX-Feature/design.md` + backtick + ` - Design specs
- Source code files

## Best Practices
- Review against plan and design
- Check for code smells
- Verify test coverage
- Ensure documentation is updated
- Provide actionable feedback
`
}

func generateTesterAgent() string {
	backtick := "`"
	return `# Tester Agent

## Role
The Tester agent creates and runs tests to validate functionality and ensure quality.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `.

## Responsibilities

1. **Test Creation**
   - Create unit tests for new features
   - Create integration tests
   - Create end-to-end tests when needed
   - Write test documentation

2. **Test Execution**
   - Run test suites
   - Validate test results
   - Report test failures
   - Track test coverage

3. **Quality Assurance**
   - Verify feature completeness
   - Validate against requirements
   - Check edge cases
   - Ensure regression tests pass

## Key Files
- ` + backtick + `doplan/XX-phase/XX-Feature/plan.md` + backtick + ` - Feature plan
- ` + backtick + `doplan/XX-phase/XX-Feature/design.md` + backtick + ` - Design specs
- ` + backtick + `doplan/XX-phase/XX-Feature/tasks.md` + backtick + ` - Task list
- Test files in project

## Best Practices
- Write tests before or alongside code
- Aim for high test coverage
- Test edge cases and error conditions
- Keep tests maintainable
- Run tests before committing
`
}

func generateDevOpsAgent() string {
	backtick := "`"
	return `# DevOps Agent

## Role
The DevOps agent handles deployment, infrastructure, and CI/CD pipeline configuration.

## Workflow & Rules
**You MUST read and obey** ` + backtick + `.doplan/ai/rules/workflow.mdc` + backtick + ` and ` + backtick + `.doplan/ai/rules/communication.mdc` + backtick + `.

## Responsibilities

1. **Deployment**
   - Configure deployment pipelines
   - Set up staging and production environments
   - Handle deployment automation
   - Monitor deployment status

2. **Infrastructure**
   - Configure infrastructure as code
   - Set up cloud resources
   - Configure networking and security
   - Manage environment variables

3. **CI/CD**
   - Configure CI/CD pipelines
   - Set up automated testing
   - Configure automated deployments
   - Monitor pipeline health

## Key Files
- ` + backtick + `.github/workflows/` + backtick + ` - GitHub Actions workflows
- ` + backtick + `docker-compose.yml` + backtick + ` - Docker configuration
- Infrastructure configuration files

## Best Practices
- Automate deployments
- Use infrastructure as code
- Monitor deployments
- Keep environments in sync
- Document deployment processes
`
}

