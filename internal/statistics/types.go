package statistics

import "time"

// StatisticsData contains all collected data
type StatisticsData struct {
	State       *StateData       `json:"state"`
	GitHub      *GitHubStats     `json:"github"`
	Checkpoints *CheckpointStats `json:"checkpoints"`
	Progress    *ProgressHistory `json:"progress"`
	Tasks       *TaskStats       `json:"tasks"`
	Testing     *TestingStats    `json:"testing"`
	CollectedAt time.Time        `json:"collectedAt"`
}

// StateData contains state-related statistics
type StateData struct {
	TotalPhases        int `json:"totalPhases"`
	TotalFeatures      int `json:"totalFeatures"`
	CompletedPhases    int `json:"completedPhases"`
	CompletedFeatures  int `json:"completedFeatures"`
	InProgressFeatures int `json:"inProgressFeatures"`
	PendingFeatures    int `json:"pendingFeatures"`
}

// GitHubStats contains GitHub-related statistics
type GitHubStats struct {
	TotalBranches  int `json:"totalBranches"`
	TotalCommits   int `json:"totalCommits"`
	TotalPRs       int `json:"totalPRs"`
	MergedPRs      int `json:"mergedPRs"`
	OpenPRs        int `json:"openPRs"`
	ClosedPRs      int `json:"closedPRs"`
	ActiveBranches int `json:"activeBranches"`
}

// CheckpointStats contains checkpoint-related statistics
type CheckpointStats struct {
	TotalCheckpoints   int       `json:"totalCheckpoints"`
	ManualCheckpoints  int       `json:"manualCheckpoints"`
	FeatureCheckpoints int       `json:"featureCheckpoints"`
	PhaseCheckpoints   int       `json:"phaseCheckpoints"`
	LastCheckpoint     time.Time `json:"lastCheckpoint,omitempty"`
}

// ProgressHistory contains progress tracking data
type ProgressHistory struct {
	OverallProgress int            `json:"overallProgress"`
	PhaseProgress   map[string]int `json:"phaseProgress"`
	FeatureProgress map[string]int `json:"featureProgress"`
	LastUpdated     time.Time      `json:"lastUpdated"`
}

// TaskStats contains task-related statistics
type TaskStats struct {
	TotalTasks     int `json:"totalTasks"`
	CompletedTasks int `json:"completedTasks"`
	PendingTasks   int `json:"pendingTasks"`
	CompletionRate int `json:"completionRate"` // percentage
}

// TestingStats contains code coverage statistics
type TestingStats struct {
	TotalStatements   int                              `json:"totalStatements"`
	CoveredStatements int                              `json:"coveredStatements"`
	PackageStats      map[string]*PackageCoverageStats `json:"packageStats"`
}

// PackageCoverageStats tracks coverage totals per package
type PackageCoverageStats struct {
	Name              string `json:"name"`
	Statements        int    `json:"statements"`
	CoveredStatements int    `json:"coveredStatements"`
}

// StatisticsMetrics contains all calculated metrics
type StatisticsMetrics struct {
	Velocity     *VelocityMetrics `json:"velocity"`
	Completion   *CompletionRates `json:"completion"`
	Time         *TimeMetrics     `json:"time"`
	Quality      *QualityMetrics  `json:"quality"`
	Testing      *TestingMetrics  `json:"testing,omitempty"`
	Trends       *Trends          `json:"trends,omitempty"`
	CalculatedAt time.Time        `json:"calculatedAt"`
}

// VelocityMetrics tracks development velocity
type VelocityMetrics struct {
	FeaturesPerDay  float64 `json:"featuresPerDay"`
	FeaturesPerWeek float64 `json:"featuresPerWeek"`
	CommitsPerDay   float64 `json:"commitsPerDay"`
	CommitsPerWeek  float64 `json:"commitsPerWeek"`
	TasksPerDay     float64 `json:"tasksPerDay"`
	PRsPerWeek      float64 `json:"prsPerWeek"`
}

// CompletionRates tracks completion percentages
type CompletionRates struct {
	Overall  int            `json:"overall"`
	Phases   map[string]int `json:"phases"`
	Features map[string]int `json:"features"`
	Tasks    int            `json:"tasks"`
}

// TimeMetrics tracks time-related statistics
type TimeMetrics struct {
	ProjectStartDate    time.Time `json:"projectStartDate"`
	DaysSinceStart      int       `json:"daysSinceStart"`
	AvgFeatureTime      float64   `json:"avgFeatureTime"` // days
	AvgPhaseTime        float64   `json:"avgPhaseTime"`   // days
	EstimatedCompletion time.Time `json:"estimatedCompletion,omitempty"`
}

// QualityMetrics tracks code quality indicators
type QualityMetrics struct {
	AvgPRReviewTime     float64 `json:"avgPRReviewTime"`     // hours
	PRMergeRate         float64 `json:"prMergeRate"`         // percentage
	AvgBranchLifetime   float64 `json:"avgBranchLifetime"`   // days
	CheckpointFrequency float64 `json:"checkpointFrequency"` // per week
}

// Trends tracks changes over time
type Trends struct {
	VelocityTrend    string  `json:"velocityTrend"`    // "improving", "declining", "stable"
	CompletionTrend  string  `json:"completionTrend"`  // "improving", "declining", "stable"
	QualityTrend     string  `json:"qualityTrend"`     // "improving", "declining", "stable"
	VelocityChange   float64 `json:"velocityChange"`   // percentage
	CompletionChange float64 `json:"completionChange"` // percentage
}

// TestingMetrics exposes testing/coverage metrics
type TestingMetrics struct {
	OverallCoverage float64                 `json:"overallCoverage"`
	Packages        []PackageCoverageMetric `json:"packages"`
}

// PackageCoverageMetric represents per-package coverage percentage
type PackageCoverageMetric struct {
	Name     string  `json:"name"`
	Coverage float64 `json:"coverage"`
}
