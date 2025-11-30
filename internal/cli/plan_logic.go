package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/internal/utils"
)

// planLogic handles the /plan command workflow
type planLogic struct {
	projectPath string
	systemDir   string
	planDir     string
	historyDir  string
}

// newPlanLogic creates a new planLogic instance
func newPlanLogic(projectPath string) (*planLogic, error) {
	systemDir := filepath.Join(projectPath, ".do", "system")
	planDir := filepath.Join(projectPath, ".do", "plan")
	historyDir := filepath.Join(projectPath, ".do", "system", "history")

	// Ensure directories exist
	if err := utils.CreateDirectory(systemDir); err != nil {
		return nil, fmt.Errorf("failed to create system directory: %w", err)
	}
	if err := utils.CreateDirectory(planDir); err != nil {
		return nil, fmt.Errorf("failed to create plan directory: %w", err)
	}
	if err := utils.CreateDirectory(historyDir); err != nil {
		return nil, fmt.Errorf("failed to create history directory: %w", err)
	}

	return &planLogic{
		projectPath: projectPath,
		systemDir:   systemDir,
		planDir:     planDir,
		historyDir:  historyDir,
	}, nil
}

// runPlanGeneration generates execution plan from IDEA.md and BRAINSTORM.md
func (p *planLogic) runPlanGeneration(out io.Writer, memoryCard *MemoryCard) error {
	// Read IDEA.md and BRAINSTORM.md
	ideaPath := filepath.Join(p.systemDir, "IDEA.md")
	brainstormPath := filepath.Join(p.systemDir, "BRAINSTORM.md")

	ideaContent, err := os.ReadFile(ideaPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("IDEA.md not found. Please run /do first to capture your project idea")
		}
		return fmt.Errorf("failed to read IDEA.md: %w", err)
	}

	brainstormContent := ""
	if utils.PathExists(brainstormPath) {
		data, err := os.ReadFile(brainstormPath)
		if err == nil {
			brainstormContent = string(data)
		}
	}

	// Analyze content and generate plan
	fmt.Fprintln(out, "📋 Analyzing your project idea and discovery meeting...")
	fmt.Fprintln(out, "")

	// Extract key information
	ideaInfo := p.extractIdeaInfo(string(ideaContent))
	brainstormInfo := p.extractBrainstormInfo(brainstormContent)

	// Determine planning style based on memory card
	planningStyle := "comprehensive"
	if memoryCard != nil {
		if memoryCard.WorkStyle == "fast" {
			planningStyle = "quick"
		} else if memoryCard.WorkStyle == "thoughtful" {
			planningStyle = "detailed"
		}
		// Learn from user preferences
		memoryCard.UpdateFromConversation(map[string]string{
			"planning_style": planningStyle,
			"project_type":   ideaInfo.ProjectType,
		})
		SaveMemoryCard(memoryCard)
	}

	// Generate TASKS.md
	tasksPath := filepath.Join(p.planDir, "TASKS.md")
	if err := p.generateTasksFile(tasksPath, ideaInfo, brainstormInfo, planningStyle, out); err != nil {
		return fmt.Errorf("failed to generate TASKS.md: %w", err)
	}

	// Scaffold the hierarchy
	fmt.Fprintln(out, "🏗️  Scaffolding plan hierarchy...")
	if err := generator.ScaffoldPlanHierarchy(p.projectPath); err != nil {
		return fmt.Errorf("failed to scaffold plan hierarchy: %w", err)
	}

	// Sync documentation
	if err := generator.SyncPlanDocumentation(p.projectPath); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Failed to sync plan documentation: %v\n", err)
	}

	// Update active state
	if err := p.updateActiveState("tasks", false); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Could not update active state: %v\n", err)
	}

	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "✅ Execution plan generated and saved to %s\n", tasksPath)
	fmt.Fprintln(out, "   Plan hierarchy scaffolded in .do/plan/")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "📋 Next steps:")
	fmt.Fprintln(out, "   • Review TASKS.md to see your execution plan")
	fmt.Fprintln(out, "   • Type /dev to start development on a feature")

	return nil
}

type ideaInfo struct {
	Overview       string
	Goals          []string
	TargetUsers    []string
	KeyFeatures    []string
	SuccessMetrics []string
	ProjectType    string
}

type brainstormInfo struct {
	Phases       map[string]map[string]string
	MeetingSpeed string
	TechStack    []string
	Requirements []string
}

func (p *planLogic) extractIdeaInfo(content string) ideaInfo {
	info := ideaInfo{
		Goals:          []string{},
		TargetUsers:    []string{},
		KeyFeatures:    []string{},
		SuccessMetrics: []string{},
	}

	lines := strings.Split(content, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "##") {
			currentSection = strings.ToLower(line)
		}

		switch {
		case strings.Contains(currentSection, "overview"):
			if !strings.HasPrefix(line, "#") && line != "" {
				info.Overview += line + " "
			}
		case strings.Contains(currentSection, "goals"):
			if strings.HasPrefix(line, "-") {
				goal := strings.TrimPrefix(line, "-")
				goal = strings.TrimSpace(goal)
				if goal != "" && !strings.Contains(goal, "[To be") {
					info.Goals = append(info.Goals, goal)
				}
			}
		case strings.Contains(currentSection, "target users") || strings.Contains(currentSection, "users"):
			if strings.HasPrefix(line, "-") {
				user := strings.TrimPrefix(line, "-")
				user = strings.TrimSpace(user)
				if user != "" && !strings.Contains(user, "[To be") {
					info.TargetUsers = append(info.TargetUsers, user)
				}
			}
		case strings.Contains(currentSection, "features"):
			if strings.HasPrefix(line, "-") {
				feature := strings.TrimPrefix(line, "-")
				feature = strings.TrimSpace(feature)
				if feature != "" && !strings.Contains(feature, "[To be") {
					info.KeyFeatures = append(info.KeyFeatures, feature)
				}
			}
		case strings.Contains(currentSection, "metrics"):
			if strings.HasPrefix(line, "-") {
				metric := strings.TrimPrefix(line, "-")
				metric = strings.TrimSpace(metric)
				if metric != "" && !strings.Contains(metric, "[To be") {
					info.SuccessMetrics = append(info.SuccessMetrics, metric)
				}
			}
		}
	}

	info.Overview = strings.TrimSpace(info.Overview)

	// Detect project type from overview
	overviewLower := strings.ToLower(info.Overview)
	switch {
	case strings.Contains(overviewLower, "saas") || strings.Contains(overviewLower, "software as a service"):
		info.ProjectType = "SaaS"
	case strings.Contains(overviewLower, "mobile app") || strings.Contains(overviewLower, "ios") || strings.Contains(overviewLower, "android"):
		info.ProjectType = "Mobile App"
	case strings.Contains(overviewLower, "website") || strings.Contains(overviewLower, "web app"):
		info.ProjectType = "Web App"
	case strings.Contains(overviewLower, "cli") || strings.Contains(overviewLower, "command line"):
		info.ProjectType = "CLI Tool"
	default:
		info.ProjectType = "Fullstack"
	}

	return info
}

func (p *planLogic) extractBrainstormInfo(content string) brainstormInfo {
	info := brainstormInfo{
		Phases:       make(map[string]map[string]string),
		TechStack:    []string{},
		Requirements: []string{},
	}

	if content == "" {
		return info
	}

	lines := strings.Split(content, "\n")
	var currentPhase string
	var currentPhaseData map[string]string

	rePhase := regexp.MustCompile(`## Phase (\d+):`)
	reQuestion := regexp.MustCompile(`### (.+)`)

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if matches := rePhase.FindStringSubmatch(line); len(matches) == 2 {
			// Save previous phase
			if currentPhase != "" && currentPhaseData != nil {
				info.Phases[currentPhase] = currentPhaseData
			}
			currentPhase = matches[1]
			currentPhaseData = make(map[string]string)
		}

		if matches := reQuestion.FindStringSubmatch(line); len(matches) == 2 && currentPhase != "" {
			question := matches[1]
			// Get answer from next line
			if i+1 < len(lines) {
				answer := strings.TrimSpace(lines[i+1])
				if answer != "" && !strings.HasPrefix(answer, "#") {
					currentPhaseData[question] = answer
				}
			}
		}

		// Extract tech stack mentions
		techKeywords := []string{"react", "next.js", "vue", "angular", "node.js", "python", "go", "typescript", "javascript", "postgresql", "mongodb", "redis"}
		lineLower := strings.ToLower(line)
		for _, keyword := range techKeywords {
			if strings.Contains(lineLower, keyword) {
				info.TechStack = append(info.TechStack, keyword)
			}
		}
	}

	// Save last phase
	if currentPhase != "" && currentPhaseData != nil {
		info.Phases[currentPhase] = currentPhaseData
	}

	// Extract meeting speed
	if strings.Contains(strings.ToLower(content), "quick start") {
		info.MeetingSpeed = "quick"
	} else if strings.Contains(strings.ToLower(content), "comprehensive") {
		info.MeetingSpeed = "comprehensive"
	} else {
		info.MeetingSpeed = "standard"
	}

	return info
}

func (p *planLogic) generateTasksFile(tasksPath string, idea ideaInfo, brainstorm brainstormInfo, style string, out io.Writer) error {
	var sb strings.Builder

	sb.WriteString("# Implementation Tasks\n")
	sb.WriteString(fmt.Sprintf("**Project**: %s\n", idea.ProjectType))
	sb.WriteString(fmt.Sprintf("**Generated**: %s\n", time.Now().Format("January 2, 2006")))
	sb.WriteString(fmt.Sprintf("**Planning Style**: %s\n", strings.Title(style)))
	sb.WriteString("\n---\n\n")

	// Generate phases based on project type and content
	phases := p.generatePhases(idea, brainstorm, style)

	phaseNum := 1
	for _, phase := range phases {
		sb.WriteString(fmt.Sprintf("## Phase %d: %s\n", phaseNum, phase.Name))
		sb.WriteString(fmt.Sprintf("**Goal**: %s\n", phase.Goal))
		sb.WriteString("**Status**: ⏳ Pending\n\n")

		taskNum := 1
		for _, task := range phase.Tasks {
			sb.WriteString(fmt.Sprintf("### %d.%d %s\n", phaseNum, taskNum, task.Title))
			sb.WriteString(fmt.Sprintf("**Description**: %s\n", task.Description))
			sb.WriteString("**Status**: ⏳ Pending\n")
			sb.WriteString(fmt.Sprintf("**Dependencies**: %s\n", task.Dependencies))
			sb.WriteString(fmt.Sprintf("**Effort**: %s\n\n", task.Effort))
			taskNum++
		}
		sb.WriteString("\n")
		phaseNum++
	}

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("*Generated: %s*\n", time.Now().Format(time.RFC3339)))
	sb.WriteString("*Command: /plan*\n")

	return utils.WriteFile(tasksPath, []byte(sb.String()))
}

type phase struct {
	Name  string
	Goal  string
	Tasks []task
}

type task struct {
	Title        string
	Description  string
	Dependencies string
	Effort       string
}

func (p *planLogic) generatePhases(idea ideaInfo, brainstorm brainstormInfo, style string) []phase {
	phases := []phase{}

	// Phase 1: Foundation
	phase1 := phase{
		Name: "Foundation",
		Goal: "Set up core infrastructure and project structure",
		Tasks: []task{
			{
				Title:        "Project Setup",
				Description:  "Initialize project structure, dependencies, and configuration",
				Dependencies: "None",
				Effort:       "2-4 hours",
			},
			{
				Title:        "Development Environment",
				Description:  "Set up development environment, tooling, and IDE configuration",
				Dependencies: "1.1",
				Effort:       "1-2 hours",
			},
		},
	}

	// Add database setup if mentioned
	if strings.Contains(strings.ToLower(idea.Overview), "database") || len(brainstorm.TechStack) > 0 {
		phase1.Tasks = append(phase1.Tasks, task{
			Title:        "Database Setup",
			Description:  "Design and set up database schema and migrations",
			Dependencies: "1.1",
			Effort:       "2-4 hours",
		})
	}

	phases = append(phases, phase1)

	// Phase 2: Core Features
	phase2 := phase{
		Name:  "Core Features",
		Goal:  "Implement main functionality and user-facing features",
		Tasks: []task{},
	}

	// Generate tasks from key features
	for i, feature := range idea.KeyFeatures {
		if i >= 5 && style == "quick" {
			break // Limit features for quick planning
		}
		phase2.Tasks = append(phase2.Tasks, task{
			Title:        feature,
			Description:  fmt.Sprintf("Implement %s feature", feature),
			Dependencies: "Phase 1",
			Effort:       "4-8 hours",
		})
	}

	// If no features extracted, add generic tasks
	if len(phase2.Tasks) == 0 {
		phase2.Tasks = append(phase2.Tasks, task{
			Title:        "Main Feature Implementation",
			Description:  "Implement the core functionality based on project requirements",
			Dependencies: "Phase 1",
			Effort:       "8-16 hours",
		})
	}

	phases = append(phases, phase2)

	// Phase 3: Enhancement (only for comprehensive style)
	if style == "comprehensive" || style == "detailed" {
		phase3 := phase{
			Name: "Enhancement",
			Goal: "Add polish, optimizations, and additional features",
			Tasks: []task{
				{
					Title:        "Performance Optimization",
					Description:  "Optimize performance, caching, and resource usage",
					Dependencies: "Phase 2",
					Effort:       "4-8 hours",
				},
				{
					Title:        "Testing & QA",
					Description:  "Write tests, perform QA, and fix bugs",
					Dependencies: "Phase 2",
					Effort:       "4-8 hours",
				},
			},
		}
		phases = append(phases, phase3)
	}

	return phases
}

func (p *planLogic) updateActiveState(phase string, locked bool) error {
	statePath := filepath.Join(p.historyDir, "active_state.json")

	state := map[string]interface{}{
		"phase":  phase,
		"locked": locked,
	}

	// Try to read existing state
	if data, err := os.ReadFile(statePath); err == nil {
		var existing map[string]interface{}
		if err := json.Unmarshal(data, &existing); err == nil {
			// Preserve existing fields
			if activeTask, ok := existing["active_task"]; ok {
				state["active_task"] = activeTask
			}
			if completed, ok := existing["completed"]; ok {
				state["completed"] = completed
			}
		}
	}

	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return utils.WriteFile(statePath, jsonData)
}
