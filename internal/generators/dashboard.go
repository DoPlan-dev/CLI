package generators

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/dashboard"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// DashboardGenerator generates dashboard files
type DashboardGenerator struct {
	projectRoot string
	state       *models.State
	githubData  *models.GitHubData
}

// NewDashboardGenerator creates a new dashboard generator
func NewDashboardGenerator(projectRoot string, state *models.State, githubData *models.GitHubData) *DashboardGenerator {
	return &DashboardGenerator{
		projectRoot: projectRoot,
		state:       state,
		githubData:  githubData,
	}
}

// Generate creates JSON, markdown and HTML dashboards
func (g *DashboardGenerator) Generate() error {
	// Generate JSON dashboard first (machine-readable)
	if err := g.GenerateJSON(); err != nil {
		return fmt.Errorf("failed to generate JSON dashboard: %w", err)
	}

	doplanDir := filepath.Join(g.projectRoot, "doplan")

	// Ensure directory exists
	if err := os.MkdirAll(doplanDir, 0755); err != nil {
		return fmt.Errorf("failed to create doplan directory: %w", err)
	}

	// Generate markdown dashboard (for backward compatibility)
	mdContent := g.generateMarkdown()
	mdPath := filepath.Join(doplanDir, "dashboard.md")
	if err := os.WriteFile(mdPath, []byte(mdContent), 0644); err != nil {
		return err
	}

	// Generate HTML dashboard
	htmlContent := g.generateHTML()
	htmlPath := filepath.Join(doplanDir, "dashboard.html")
	if err := os.WriteFile(htmlPath, []byte(htmlContent), 0644); err != nil {
		return err
	}

	return nil
}

// GenerateJSON generates the machine-readable dashboard.json
func (g *DashboardGenerator) GenerateJSON() error {
	doplanDir := filepath.Join(g.projectRoot, ".doplan")
	if err := os.MkdirAll(doplanDir, 0755); err != nil {
		return fmt.Errorf("failed to create .doplan directory: %w", err)
	}

	dashboard := g.buildDashboardJSON()
	
	data, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal dashboard JSON: %w", err)
	}

	jsonPath := filepath.Join(doplanDir, "dashboard.json")
	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write dashboard JSON: %w", err)
	}

	return nil
}

// DashboardJSON represents the dashboard JSON structure
type DashboardJSON struct {
	Version   string                 `json:"version"`
	Generated string                 `json:"generated"`
	Project   ProjectJSON            `json:"project"`
	GitHub    GitHubJSON             `json:"github"`
	Phases    []PhaseJSON            `json:"phases"`
	Summary   SummaryJSON            `json:"summary"`
	Activity  ActivityJSON           `json:"activity"`
	APIKeys   APIKeysJSON            `json:"apiKeys"`
	Velocity  VelocityJSON           `json:"velocity"`
}

type ProjectJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Progress    int    `json:"progress"`
	Status      string `json:"status"`
	StartDate   string `json:"startDate"`
	TargetDate  string `json:"targetDate"`
}

type GitHubJSON struct {
	Repository   string   `json:"repository"`
	Branch       string   `json:"branch"`
	Commits      int      `json:"commits"`
	Contributors []string `json:"contributors"`
	LastCommit   string   `json:"lastCommit"`
}

type PhaseJSON struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	Progress    int          `json:"progress"`
	StartDate   string       `json:"startDate"`
	TargetDate  string       `json:"targetDate"`
	Features    []FeatureJSON `json:"features"`
	Stats       PhaseStatsJSON `json:"stats"`
}

type FeatureJSON struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Progress     int    `json:"progress"`
	Branch       string `json:"branch"`
	PR           *PRJSON `json:"pr"`
	Commits      int    `json:"commits"`
	LastActivity string `json:"lastActivity"`
	Tasks        []TaskJSON `json:"tasks"`
}

type TaskJSON struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type PRJSON struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type PhaseStatsJSON struct {
	TotalFeatures   int `json:"totalFeatures"`
	Completed       int `json:"completed"`
	InProgress      int `json:"inProgress"`
	Todo            int `json:"todo"`
	TotalTasks      int `json:"totalTasks"`
	CompletedTasks  int `json:"completedTasks"`
}

type SummaryJSON struct {
	TotalPhases      int `json:"totalPhases"`
	Completed        int `json:"completed"`
	InProgress       int `json:"inProgress"`
	Todo             int `json:"todo"`
	TotalFeatures    int `json:"totalFeatures"`
	TotalTasks       int `json:"totalTasks"`
	CompletedTasks   int `json:"completedTasks"`
}

type ActivityJSON struct {
	Last24Hours ActivityPeriodJSON `json:"last24Hours"`
	Last7Days   ActivityPeriodJSON `json:"last7Days"`
	RecentActivity []ActivityItemJSON `json:"recentActivity"`
}

type ActivityPeriodJSON struct {
	Commits       int `json:"commits"`
	TasksCompleted int `json:"tasksCompleted"`
	FilesChanged  int `json:"filesChanged"`
}

type ActivityItemJSON struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type APIKeysJSON struct {
	Total      int `json:"total"`
	Configured int `json:"configured"`
	Pending    int `json:"pending"`
	Optional   int `json:"optional"`
	Completion int `json:"completion"`
}

type VelocityJSON struct {
	TasksPerDay        float64 `json:"tasksPerDay"`
	CommitsPerDay      float64 `json:"commitsPerDay"`
	EstimatedCompletion string `json:"estimatedCompletion"`
	DaysToLaunch       int     `json:"daysToLaunch"`
}

// buildDashboardJSON builds the dashboard JSON structure
func (g *DashboardGenerator) buildDashboardJSON() *DashboardJSON {
	overallProgress := g.calculateOverallProgress()
	
	// Read progress data early for use in feature building
	progressParser := dashboard.NewProgressParser(g.projectRoot)
	progressData, err := progressParser.ReadProgressFiles()
	if err != nil {
		progressData = make(map[string]*dashboard.ProgressData)
	}
	
	// Build phases
	phases := []PhaseJSON{}
	for _, phase := range g.state.Phases {
		phaseProgress := g.calculatePhaseProgress(phase.ID)
		features := []FeatureJSON{}
		
		for _, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature != nil {
				tasks := []TaskJSON{}
				for _, taskPhase := range feature.TaskPhases {
					for _, task := range taskPhase.Tasks {
						tasks = append(tasks, TaskJSON{
							Name:      task.Name,
							Completed: task.Completed,
						})
					}
				}
				
				var pr *PRJSON
				if feature.PR != nil {
					pr = &PRJSON{
						Number: feature.PR.Number,
						Title:  feature.PR.Title,
						URL:    feature.PR.URL,
						Status: feature.PR.Status,
					}
				}
				
				// Calculate commits for this feature from branch name
				commits := 0
				if feature.Branch != "" && g.githubData != nil {
					for _, commit := range g.githubData.Commits {
						if commit.Branch == feature.Branch {
							commits++
						}
					}
				}

				// Get last activity from progress data or commits
				lastActivity := feature.StartDate
				if progressData != nil {
					if pd, ok := progressData[feature.ID]; ok && !pd.LastUpdated.IsZero() {
						lastActivity = pd.LastUpdated.Format(time.RFC3339)
					}
				}
				// Also check commits for this feature
				if commits > 0 && g.githubData != nil {
					for _, commit := range g.githubData.Commits {
						if commit.Branch == feature.Branch {
							lastActivity = commit.Date
							break
						}
					}
				}

				features = append(features, FeatureJSON{
					ID:           feature.ID,
					Name:         feature.Name,
					Status:       feature.Status,
					Progress:     feature.Progress,
					Branch:       feature.Branch,
					PR:           pr,
					Commits:      commits,
					LastActivity: lastActivity,
					Tasks:        tasks,
				})
			}
		}
		
		// Calculate phase stats
		completed := 0
		inProgress := 0
		todo := 0
		totalTasks := 0
		completedTasks := 0
		
		for _, feature := range features {
			switch feature.Status {
			case "complete":
				completed++
			case "in-progress":
				inProgress++
			default:
				todo++
			}
			for _, task := range feature.Tasks {
				totalTasks++
				if task.Completed {
					completedTasks++
				}
			}
		}
		
		phases = append(phases, PhaseJSON{
			ID:          phase.ID,
			Name:        phase.Name,
			Description: phase.Description,
			Status:      phase.Status,
			Progress:    phaseProgress,
			StartDate:   phase.StartDate,
			TargetDate:  phase.TargetDate,
			Features:    features,
			Stats: PhaseStatsJSON{
				TotalFeatures:  len(features),
				Completed:      completed,
				InProgress:     inProgress,
				Todo:           todo,
				TotalTasks:     totalTasks,
				CompletedTasks: completedTasks,
			},
		})
	}
	
	// Calculate summary
	completedPhases := 0
	inProgressPhases := 0
	todoPhases := 0
	totalFeatures := 0
	totalTasks := 0
	completedTasks := 0
	
	for _, phase := range phases {
		switch phase.Status {
		case "complete":
			completedPhases++
		case "in-progress":
			inProgressPhases++
		default:
			todoPhases++
		}
		totalFeatures += phase.Stats.TotalFeatures
		totalTasks += phase.Stats.TotalTasks
		completedTasks += phase.Stats.CompletedTasks
	}
	
	// Build GitHub data
	githubJSON := GitHubJSON{
		Repository:   "", // Will be populated from config
		Branch:       "main",
		Commits:      len(g.githubData.Commits),
		Contributors: []string{},
		LastCommit:   "",
	}
	if len(g.githubData.Commits) > 0 {
		githubJSON.LastCommit = g.githubData.Commits[0].Date
		// Extract unique contributors
		contributorMap := make(map[string]bool)
		for _, commit := range g.githubData.Commits {
			if commit.Author != "" {
				contributorMap[commit.Author] = true
			}
		}
		for contributor := range contributorMap {
			githubJSON.Contributors = append(githubJSON.Contributors, contributor)
		}
	}
	
	// Generate comprehensive activity feed (progressData already loaded above)
	activityGen := dashboard.NewActivityGenerator(g.state, g.githubData, progressData)
	activityData := activityGen.GenerateActivityFeed()
	
	// Convert to ActivityJSON
	activity := ActivityJSON{
		Last24Hours: ActivityPeriodJSON{
			Commits:       activityData.Last24Hours.Commits,
			TasksCompleted: activityData.Last24Hours.TasksCompleted,
			FilesChanged:  activityData.Last24Hours.FilesChanged,
		},
		Last7Days: ActivityPeriodJSON{
			Commits:       activityData.Last7Days.Commits,
			TasksCompleted: activityData.Last7Days.TasksCompleted,
			FilesChanged:  activityData.Last7Days.FilesChanged,
		},
		RecentActivity: []ActivityItemJSON{},
	}
	
	for _, item := range activityData.RecentActivity {
		activity.RecentActivity = append(activity.RecentActivity, ActivityItemJSON{
			Type:      item.Type,
			Message:   item.Message,
			Timestamp: item.Timestamp,
		})
	}
	
	return &DashboardJSON{
		Version:   "1.0",
		Generated: time.Now().Format(time.RFC3339),
		Project: ProjectJSON{
			Name:        "", // Will be populated from config
			Description: "",
			Version:     "1.0.0",
			Progress:    overallProgress,
			Status:      "in-progress",
			StartDate:   "",
			TargetDate:  "",
		},
		GitHub:   githubJSON,
		Phases:   phases,
		Summary: SummaryJSON{
			TotalPhases:    len(phases),
			Completed:      completedPhases,
			InProgress:     inProgressPhases,
			Todo:           todoPhases,
			TotalFeatures:  totalFeatures,
			TotalTasks:     totalTasks,
			CompletedTasks: completedTasks,
		},
		Activity: activity,
		APIKeys: APIKeysJSON{
			Total:      0,
			Configured: 0,
			Pending:    0,
			Optional:   0,
			Completion: 0,
		},
		Velocity: g.calculateVelocity(),
	}
}

// GenerateInitialDashboard creates the initial dashboard content
func GenerateInitialDashboard() string {
	return `# Project Dashboard

*Last Updated: ` + time.Now().Format(time.RFC3339) + `*

## Overall Progress

` + "```" + `
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0%
` + "```" + `

## Phase Progress

_No phases created yet. Use ` + "`/Plan`" + ` command in Cursor to generate your project structure._

## Active Pull Requests

_No active pull requests._

## Next Actions

1. Use **/Discuss** command in Cursor to start idea discussion
2. Use **/Generate** command to generate PRD and contracts
3. Use **/Plan** command to create phase structure

---

*Generated by DoPlan - Project Workflow Manager*
`
}

func (g *DashboardGenerator) generateMarkdown() string {
	var sb strings.Builder

	sb.WriteString("# Project Dashboard\n\n")
	sb.WriteString(fmt.Sprintf("*Last Updated: %s*\n\n", time.Now().Format(time.RFC3339)))

	// Overall Progress
	overallProgress := g.calculateOverallProgress()
	sb.WriteString("## Overall Progress\n\n")
	sb.WriteString(g.renderProgressBar(overallProgress))
	sb.WriteString(fmt.Sprintf(" **%d%%**\n\n", overallProgress))

	// Phase Progress
	sb.WriteString("## Phase Progress\n\n")
	if len(g.state.Phases) == 0 {
		sb.WriteString("_No phases created yet. Use `/Plan` command in Cursor to generate your project structure._\n\n")
	} else {
		for i, phase := range g.state.Phases {
			phaseProgress := g.calculatePhaseProgress(phase.ID)
			sb.WriteString(fmt.Sprintf("### Phase %d: %s\n", i+1, phase.Name))
			sb.WriteString(fmt.Sprintf("Status: **%s**\n", phase.Status))
			sb.WriteString(g.renderProgressBar(phaseProgress))
			sb.WriteString(fmt.Sprintf(" **%d%%**\n\n", phaseProgress))

			// Features in this phase
			for _, featureID := range phase.Features {
				feature := g.findFeature(featureID)
				if feature != nil {
					sb.WriteString(fmt.Sprintf("- **%s**: %d%% - %s\n", feature.Name, feature.Progress, feature.Status))
				}
			}
			sb.WriteString("\n")
		}
	}

	// GitHub Activity
	sb.WriteString("## GitHub Activity\n\n")
	if g.githubData != nil && (len(g.githubData.Branches) > 0 || len(g.githubData.Commits) > 0 || len(g.githubData.PRs) > 0) {
		// Active Branches
		if len(g.githubData.Branches) > 0 {
			sb.WriteString("### Active Branches\n\n")
			sb.WriteString("| Branch | Status | Commits | PR |\n")
			sb.WriteString("|--------|--------|---------|----|\n")
			for _, branch := range g.githubData.Branches {
				prStatus := "No"
				if branch.HasPR {
					prStatus = fmt.Sprintf("[Yes](%s)", branch.PRURL)
				}
				sb.WriteString(fmt.Sprintf("| %s | %s | %d | %s |\n", branch.Name, branch.Status, branch.CommitCount, prStatus))
			}
			sb.WriteString("\n")
		}

		// Recent Commits
		if len(g.githubData.Commits) > 0 {
			sb.WriteString("### Recent Commits\n\n")
			sb.WriteString("| Hash | Message | Author | Date |\n")
			sb.WriteString("|------|---------|--------|------|\n")
			maxCommits := 10
			if len(g.githubData.Commits) < maxCommits {
				maxCommits = len(g.githubData.Commits)
			}
			for i := 0; i < maxCommits; i++ {
				commit := g.githubData.Commits[i]
				hashDisplay := commit.Hash
				if len(commit.Hash) > 8 {
					hashDisplay = commit.Hash[:8]
				}
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					hashDisplay, commit.Message, commit.Author, commit.Date))
			}
			sb.WriteString("\n")
		}

		// Pull Requests
		if len(g.githubData.PRs) > 0 {
			sb.WriteString("### Pull Requests\n\n")
			sb.WriteString("| # | Title | Status |\n")
			sb.WriteString("|---|-------|--------|\n")
			for _, pr := range g.githubData.PRs {
				sb.WriteString(fmt.Sprintf("| [#%d](%s) | %s | %s |\n", pr.Number, pr.URL, pr.Title, pr.Status))
			}
			sb.WriteString("\n")
		}
	} else {
		sb.WriteString("_No GitHub activity yet._\n\n")
	}

	// Next Actions
	sb.WriteString("## Next Actions\n\n")
	nextActions := g.getNextActions()
	for i, action := range nextActions {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, action))
	}
	sb.WriteString("\n")

	sb.WriteString("---\n\n")
	sb.WriteString("*Generated by DoPlan - Project Workflow Manager*\n")

	return sb.String()
}

func (g *DashboardGenerator) generateHTML() string {
	overallProgress := g.calculateOverallProgress()

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DoPlan Dashboard</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            padding: 20px;
            min-height: 100vh;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 12px;
            padding: 30px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
        }
        .last-updated {
            color: #666;
            font-size: 14px;
            margin-bottom: 30px;
        }
        .section {
            margin-bottom: 40px;
        }
        .section h2 {
            color: #667eea;
            margin-bottom: 20px;
            padding-bottom: 10px;
            border-bottom: 2px solid #667eea;
        }
        .progress-bar {
            background: #e0e0e0;
            border-radius: 10px;
            height: 30px;
            margin: 10px 0;
            overflow: hidden;
            position: relative;
        }
        .progress-fill {
            background: linear-gradient(90deg, #667eea 0%%, #764ba2 100%%);
            height: 100%%;
            transition: width 0.3s ease;
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-weight: bold;
            font-size: 12px;
        }
        .phase {
            background: #f8f9fa;
            border-radius: 8px;
            padding: 20px;
            margin-bottom: 20px;
        }
        .phase h3 {
            color: #333;
            margin-bottom: 10px;
        }
        .feature {
            padding: 10px;
            margin: 5px 0;
            background: white;
            border-left: 4px solid #667eea;
            border-radius: 4px;
        }
        table {
            width: 100%%;
            border-collapse: collapse;
            margin-top: 10px;
        }
        th, td {
            padding: 12px;
            text-align: left;
            border-bottom: 1px solid #e0e0e0;
        }
        th {
            background: #f8f9fa;
            font-weight: 600;
            color: #333;
        }
        tr:hover {
            background: #f8f9fa;
        }
        .status {
            padding: 4px 8px;
            border-radius: 4px;
            font-size: 12px;
            font-weight: 600;
        }
        .status.active { background: #4caf50; color: white; }
        .status.complete { background: #2196f3; color: white; }
        .status.in-progress { background: #ff9800; color: white; }
        .next-actions {
            list-style: none;
        }
        .next-actions li {
            padding: 10px;
            margin: 5px 0;
            background: #f8f9fa;
            border-radius: 4px;
            border-left: 4px solid #667eea;
        }
        a {
            color: #667eea;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Project Dashboard</h1>
        <div class="last-updated">Last Updated: %s</div>
        
        <div class="section">
            <h2>Overall Progress</h2>
            <div class="progress-bar">
                <div class="progress-fill" style="width: %d%%">%d%%</div>
            </div>
        </div>
        
        <div class="section">
            <h2>Phase Progress</h2>
            %s
        </div>
        
        <div class="section">
            <h2>GitHub Activity</h2>
            %s
        </div>
        
        <div class="section">
            <h2>Next Actions</h2>
            <ul class="next-actions">
                %s
            </ul>
        </div>
    </div>
</body>
</html>`,
		time.Now().Format(time.RFC3339),
		overallProgress,
		overallProgress,
		g.generatePhaseHTML(),
		g.generateGitHubHTML(),
		g.generateNextActionsHTML(),
	)
}

func (g *DashboardGenerator) calculateOverallProgress() int {
	if len(g.state.Phases) == 0 {
		return 0
	}

	totalProgress := 0
	for _, phase := range g.state.Phases {
		totalProgress += g.calculatePhaseProgress(phase.ID)
	}

	return totalProgress / len(g.state.Phases)
}

func (g *DashboardGenerator) calculatePhaseProgress(phaseID string) int {
	phase := g.findPhase(phaseID)
	if phase == nil {
		return 0
	}

	if len(phase.Features) == 0 {
		return 0
	}

	totalProgress := 0
	for _, featureID := range phase.Features {
		feature := g.findFeature(featureID)
		if feature != nil {
			totalProgress += feature.Progress
		}
	}

	return totalProgress / len(phase.Features)
}

// calculateVelocity calculates velocity metrics
func (g *DashboardGenerator) calculateVelocity() VelocityJSON {
	velocity := VelocityJSON{
		TasksPerDay:        0.0,
		CommitsPerDay:      0.0,
		EstimatedCompletion: "",
		DaysToLaunch:       0,
	}

	// Calculate commits per day (last 7 days)
	if g.githubData != nil && len(g.githubData.Commits) > 0 {
		now := time.Now()
		sevenDaysAgo := now.AddDate(0, 0, -7)
		commitsCount := 0

		for _, commit := range g.githubData.Commits {
			commitTime, err := time.Parse(time.RFC3339, commit.Date)
			if err != nil {
				// Try other formats
				commitTime, err = time.Parse("2006-01-02 15:04:05", commit.Date)
				if err != nil {
					continue
				}
			}
			if commitTime.After(sevenDaysAgo) {
				commitsCount++
			}
		}

		velocity.CommitsPerDay = float64(commitsCount) / 7.0
	}

	// Calculate tasks per day from state
	if g.state != nil {
		totalTasks := 0
		completedTasks := 0
		for _, feature := range g.state.Features {
			for _, taskPhase := range feature.TaskPhases {
				for _, task := range taskPhase.Tasks {
					totalTasks++
					if task.Completed {
						completedTasks++
					}
				}
			}
		}

		// Estimate based on project start (simplified)
		if totalTasks > 0 {
			// Assume project started when first feature was created
			// For now, use a simple estimate
			velocity.TasksPerDay = float64(completedTasks) / 30.0 // Rough estimate
		}
	}

	// Calculate estimated completion (simplified)
	if velocity.TasksPerDay > 0 && g.state != nil {
		remainingTasks := 0
		for _, feature := range g.state.Features {
			for _, taskPhase := range feature.TaskPhases {
				for _, task := range taskPhase.Tasks {
					if !task.Completed {
						remainingTasks++
					}
				}
			}
		}

		if remainingTasks > 0 && velocity.TasksPerDay > 0 {
			daysRemaining := float64(remainingTasks) / velocity.TasksPerDay
			completionDate := time.Now().AddDate(0, 0, int(daysRemaining))
			velocity.EstimatedCompletion = completionDate.Format(time.RFC3339)
			velocity.DaysToLaunch = int(daysRemaining)
		}
	}

	return velocity
}

func (g *DashboardGenerator) findPhase(phaseID string) *models.Phase {
	for _, phase := range g.state.Phases {
		if phase.ID == phaseID {
			return &phase
		}
	}
	return nil
}

func (g *DashboardGenerator) findFeature(featureID string) *models.Feature {
	for _, feature := range g.state.Features {
		if feature.ID == featureID {
			return &feature
		}
	}
	return nil
}

func (g *DashboardGenerator) renderProgressBar(percentage int) string {
	filled := percentage / 2
	empty := 50 - filled

	filledBar := strings.Repeat("█", filled)
	emptyBar := strings.Repeat("░", empty)

	codeBlock := "```"
	return fmt.Sprintf("%s\n%s%s %d%%\n%s\n", codeBlock, filledBar, emptyBar, percentage, codeBlock)
}

func (g *DashboardGenerator) getNextActions() []string {
	actions := []string{}

	if g.state.Idea == nil {
		actions = append(actions, "Use **/Discuss** command in Cursor to start idea discussion")
		return actions
	}

	if len(g.state.Phases) == 0 {
		actions = append(actions, "Use **/Plan** command to create phase structure")
		return actions
	}

	// Find incomplete features
	for _, feature := range g.state.Features {
		if feature.Status != "complete" && feature.Progress < 100 {
			actions = append(actions, fmt.Sprintf("Continue work on **%s** (%d%% complete)", feature.Name, feature.Progress))
			if len(actions) >= 5 {
				break
			}
		}
	}

	if len(actions) == 0 {
		actions = append(actions, "All features complete! Great work!")
	}

	return actions
}

func (g *DashboardGenerator) generatePhaseHTML() string {
	if len(g.state.Phases) == 0 {
		return "<p><em>No phases created yet. Use `/Plan` command in Cursor to generate your project structure.</em></p>"
	}

	var sb strings.Builder
	for i, phase := range g.state.Phases {
		phaseProgress := g.calculatePhaseProgress(phase.ID)
		sb.WriteString(fmt.Sprintf(`<div class="phase">
            <h3>Phase %d: %s</h3>
            <div class="progress-bar">
                <div class="progress-fill" style="width: %d%%">%d%%</div>
            </div>
            <p><strong>Status:</strong> <span class="status %s">%s</span></p>
        `, i+1, phase.Name, phaseProgress, phaseProgress, strings.ToLower(phase.Status), phase.Status))

		// Features
		for _, featureID := range phase.Features {
			feature := g.findFeature(featureID)
			if feature != nil {
				sb.WriteString(fmt.Sprintf(`<div class="feature">
                    <strong>%s</strong> - %d%% - <span class="status %s">%s</span>
                </div>`, feature.Name, feature.Progress, strings.ToLower(feature.Status), feature.Status))
			}
		}

		sb.WriteString("</div>")
	}

	return sb.String()
}

func (g *DashboardGenerator) generateGitHubHTML() string {
	if g.githubData == nil || (len(g.githubData.Branches) == 0 && len(g.githubData.Commits) == 0 && len(g.githubData.PRs) == 0) {
		return "<p><em>No GitHub activity yet.</em></p>"
	}

	var sb strings.Builder

	// Branches
	if len(g.githubData.Branches) > 0 {
		sb.WriteString("<h3>Active Branches</h3><table><tr><th>Branch</th><th>Status</th><th>Commits</th><th>PR</th></tr>")
		for _, branch := range g.githubData.Branches {
			prCell := "No"
			if branch.HasPR {
				prCell = fmt.Sprintf(`<a href="%s">Yes</a>`, branch.PRURL)
			}
			sb.WriteString(fmt.Sprintf("<tr><td>%s</td><td><span class=\"status %s\">%s</span></td><td>%d</td><td>%s</td></tr>",
				branch.Name, strings.ToLower(branch.Status), branch.Status, branch.CommitCount, prCell))
		}
		sb.WriteString("</table>")
	}

	// PRs
	if len(g.githubData.PRs) > 0 {
		sb.WriteString("<h3>Pull Requests</h3><table><tr><th>#</th><th>Title</th><th>Status</th></tr>")
		for _, pr := range g.githubData.PRs {
			sb.WriteString(fmt.Sprintf("<tr><td><a href=\"%s\">#%d</a></td><td>%s</td><td><span class=\"status %s\">%s</span></td></tr>",
				pr.URL, pr.Number, pr.Title, strings.ToLower(pr.Status), pr.Status))
		}
		sb.WriteString("</table>")
	}

	return sb.String()
}

func (g *DashboardGenerator) generateNextActionsHTML() string {
	actions := g.getNextActions()
	var sb strings.Builder
	for _, action := range actions {
		sb.WriteString(fmt.Sprintf("<li>%s</li>", action))
	}
	return sb.String()
}
