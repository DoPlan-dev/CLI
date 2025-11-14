package statistics

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/internal/github"
)

// Collector gathers data from various sources
type Collector struct {
	projectRoot string
	configMgr   *config.Manager
	githubSync  *github.GitHubSync
}

// NewCollector creates a new statistics collector
func NewCollector(projectRoot string) *Collector {
	return &Collector{
		projectRoot: projectRoot,
		configMgr:   config.NewManager(projectRoot),
		githubSync:  github.NewGitHubSync(projectRoot),
	}
}

// Collect gathers all statistics data (parallel collection)
func (c *Collector) Collect() (*StatisticsData, error) {
	data := &StatisticsData{
		CollectedAt: time.Now(),
	}

	// Use goroutines for parallel collection
	type result struct {
		name string
		err  error
	}

	var wg sync.WaitGroup
	results := make(chan result, 6)

	// Collect state data
	wg.Add(1)
	go func() {
		defer wg.Done()
		stateData, err := c.CollectState()
		if err == nil {
			data.State = stateData
		}
		results <- result{name: "state", err: err}
	}()

	// Collect GitHub data
	wg.Add(1)
	go func() {
		defer wg.Done()
		githubData, err := c.CollectGitHub()
		if err == nil {
			data.GitHub = githubData
		}
		results <- result{name: "github", err: err}
	}()

	// Collect checkpoint data
	wg.Add(1)
	go func() {
		defer wg.Done()
		checkpointData, err := c.CollectCheckpoints()
		if err == nil {
			data.Checkpoints = checkpointData
		}
		results <- result{name: "checkpoint", err: err}
	}()

	// Collect progress data
	wg.Add(1)
	go func() {
		defer wg.Done()
		progressData, err := c.CollectProgress()
		if err == nil {
			data.Progress = progressData
		}
		results <- result{name: "progress", err: err}
	}()

	// Collect task data
	wg.Add(1)
	go func() {
		defer wg.Done()
		taskData, err := c.CollectTasks()
		if err == nil {
			data.Tasks = taskData
		}
		results <- result{name: "tasks", err: err}
	}()

	// Collect testing data
	wg.Add(1)
	go func() {
		defer wg.Done()
		testingData, err := c.CollectTesting()
		if err == nil && testingData != nil {
			data.Testing = testingData
		}
		results <- result{name: "testing", err: err}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
	close(results)

	return data, nil
}

// CollectState collects state-related statistics
func (c *Collector) CollectState() (*StateData, error) {
	state, err := c.configMgr.LoadState()
	if err != nil {
		return nil, err
	}

	data := &StateData{
		TotalPhases:   len(state.Phases),
		TotalFeatures: len(state.Features),
	}

	for _, phase := range state.Phases {
		if phase.Status == "complete" {
			data.CompletedPhases++
		}
	}

	for _, feature := range state.Features {
		switch feature.Status {
		case "complete":
			data.CompletedFeatures++
		case "in-progress":
			data.InProgressFeatures++
		case "pending":
			data.PendingFeatures++
		}
	}

	return data, nil
}

// CollectGitHub collects GitHub-related statistics
func (c *Collector) CollectGitHub() (*GitHubStats, error) {
	githubData, err := c.githubSync.LoadData()
	if err != nil {
		return nil, err
	}

	stats := &GitHubStats{
		TotalBranches: len(githubData.Branches),
		TotalCommits:  len(githubData.Commits),
		TotalPRs:      len(githubData.PRs),
	}

	for _, pr := range githubData.PRs {
		switch pr.Status {
		case "merged":
			stats.MergedPRs++
		case "open":
			stats.OpenPRs++
		case "closed":
			stats.ClosedPRs++
		}
	}

	// Count active branches (branches with recent commits)
	for _, branch := range githubData.Branches {
		if branch.CommitCount > 0 {
			stats.ActiveBranches++
		}
	}

	return stats, nil
}

// CollectCheckpoints collects checkpoint-related statistics
func (c *Collector) CollectCheckpoints() (*CheckpointStats, error) {
	checkpointDir := filepath.Join(c.projectRoot, ".doplan", "checkpoints")
	if _, err := os.Stat(checkpointDir); os.IsNotExist(err) {
		return &CheckpointStats{}, nil
	}

	entries, err := os.ReadDir(checkpointDir)
	if err != nil {
		return nil, err
	}

	stats := &CheckpointStats{
		TotalCheckpoints: len(entries),
	}

	var lastCheckpoint time.Time
	for _, entry := range entries {
		if entry.IsDir() {
			// Check checkpoint metadata
			metadataPath := filepath.Join(checkpointDir, entry.Name(), "metadata.json")
			if data, err := os.ReadFile(metadataPath); err == nil {
				var metadata struct {
					Type      string    `json:"type"`
					CreatedAt time.Time `json:"createdAt"`
				}
				if err := json.Unmarshal(data, &metadata); err == nil {
					switch metadata.Type {
					case "manual":
						stats.ManualCheckpoints++
					case "feature":
						stats.FeatureCheckpoints++
					case "phase":
						stats.PhaseCheckpoints++
					}
					if metadata.CreatedAt.After(lastCheckpoint) {
						lastCheckpoint = metadata.CreatedAt
					}
				}
			}
		}
	}

	if !lastCheckpoint.IsZero() {
		stats.LastCheckpoint = lastCheckpoint
	}

	return stats, nil
}

// CollectProgress collects progress-related statistics
func (c *Collector) CollectProgress() (*ProgressHistory, error) {
	state, err := c.configMgr.LoadState()
	if err != nil {
		return nil, err
	}

	history := &ProgressHistory{
		PhaseProgress:   make(map[string]int),
		FeatureProgress: make(map[string]int),
		LastUpdated:     time.Now(),
	}

	history.OverallProgress = state.Progress.Overall

	for phaseID, progress := range state.Progress.Phases {
		history.PhaseProgress[phaseID] = progress
	}

	for _, feature := range state.Features {
		history.FeatureProgress[feature.ID] = feature.Progress
	}

	return history, nil
}

// CollectTasks collects task-related statistics
func (c *Collector) CollectTasks() (*TaskStats, error) {
	state, err := c.configMgr.LoadState()
	if err != nil {
		return nil, err
	}

	stats := &TaskStats{}

	// Count tasks from all features
	for _, feature := range state.Features {
		for _, taskPhase := range feature.TaskPhases {
			for _, task := range taskPhase.Tasks {
				stats.TotalTasks++
				if task.Completed {
					stats.CompletedTasks++
				} else {
					stats.PendingTasks++
				}
			}
		}
	}

	// Calculate completion rate
	if stats.TotalTasks > 0 {
		stats.CompletionRate = (stats.CompletedTasks * 100) / stats.TotalTasks
	}

	return stats, nil
}

// CollectTesting collects test coverage metrics from coverage profiles
func (c *Collector) CollectTesting() (*TestingStats, error) {
	coveragePaths := []string{
		filepath.Join(c.projectRoot, "coverage.out"),
		filepath.Join(c.projectRoot, "coverage", "coverage.out"),
		filepath.Join(c.projectRoot, ".coverage", "coverage.out"),
	}

	var coveragePath string
	for _, candidate := range coveragePaths {
		if _, err := os.Stat(candidate); err == nil {
			coveragePath = candidate
			break
		}
	}

	if coveragePath == "" {
		// No coverage profile available
		return nil, nil
	}

	file, err := os.Open(coveragePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &TestingStats{
		PackageStats: make(map[string]*PackageCoverageStats),
	}

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		// Skip coverage mode declaration
		if line == 0 && strings.HasPrefix(text, "mode:") {
			line++
			continue
		}

		fields := strings.Fields(text)
		if len(fields) < 3 {
			continue
		}

		statements, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		count, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		stats.TotalStatements += statements
		if count > 0 {
			stats.CoveredStatements += statements
		}

		filePath := fields[0]
		if idx := strings.Index(filePath, ":"); idx != -1 {
			filePath = filePath[:idx]
		}
		pkgPath := filepath.Dir(filePath)
		pkgStats, ok := stats.PackageStats[pkgPath]
		if !ok {
			pkgStats = &PackageCoverageStats{Name: pkgPath}
			stats.PackageStats[pkgPath] = pkgStats
		}
		pkgStats.Statements += statements
		if count > 0 {
			pkgStats.CoveredStatements += statements
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if stats.TotalStatements == 0 {
		return nil, nil
	}

	return stats, nil
}
