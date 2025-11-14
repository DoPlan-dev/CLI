package models

import "time"

// Config represents the DoPlan configuration
type Config struct {
	IDE         string          `json:"ide"`
	Installed   bool            `json:"installed"`
	InstalledAt time.Time       `json:"installedAt"`
	Version     string          `json:"version"`
	GitHub      GitHubConfig    `json:"github"`
	Checkpoint  CheckpointConfig `json:"checkpoint"`
	State       StateConfig     `json:"state"`
}

// GitHubConfig contains GitHub-related settings
type GitHubConfig struct {
	Enabled    bool `json:"enabled"`
	AutoBranch bool `json:"autoBranch"`
	AutoPR     bool `json:"autoPR"`
}

// CheckpointConfig contains checkpoint-related settings
type CheckpointConfig struct {
	AutoFeature bool `json:"autoFeature"` // Auto-create checkpoint when feature starts
	AutoPhase   bool `json:"autoPhase"`   // Auto-create checkpoint when phase starts
	AutoComplete bool `json:"autoComplete"` // Auto-create checkpoint when feature/phase completes
}

// StateConfig contains current workflow state
type StateConfig struct {
	CurrentPhase   string `json:"currentPhase"`
	CurrentFeature string `json:"currentFeature"`
	IdeaCaptured   bool   `json:"ideaCaptured"`
	PRDGenerated   bool   `json:"prdGenerated"`
	PlanGenerated  bool   `json:"planGenerated"`
}

// State represents the full project state
type State struct {
	Idea     *Idea     `json:"idea"`
	Phases   []Phase   `json:"phases"`
	Features []Feature `json:"features"`
	Progress Progress  `json:"progress"`
}

// Idea contains idea discussion data
type Idea struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	ProblemStatement string   `json:"problemStatement"`
	Solution         string   `json:"solution"`
	TargetUsers      []string `json:"targetUsers"`
	TechStack        []string `json:"techStack"`
}

// Phase represents a project phase
type Phase struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Objectives  []string `json:"objectives"`
	Features    []string `json:"features"`
	StartDate   string   `json:"startDate"`
	TargetDate  string   `json:"targetDate"`
	Duration    string   `json:"duration"`
}

// Feature represents a feature within a phase
type Feature struct {
	ID             string       `json:"id"`
	Phase          string       `json:"phase"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Status         string       `json:"status"`
	Progress       int          `json:"progress"`
	Branch         string       `json:"branch"`
	PR             *PullRequest `json:"pr"`
	CheckpointID   string       `json:"checkpointId"` // Latest checkpoint ID
	Objectives     []string     `json:"objectives"`
	Requirements   []string     `json:"requirements"`
	Dependencies   []string     `json:"dependencies"`
	DesignOverview string       `json:"designOverview"`
	Architecture   string       `json:"architecture"`
	UserFlow       string       `json:"userFlow"`
	TechnicalSpecs string       `json:"technicalSpecs"`
	TaskPhases     []TaskPhase  `json:"taskPhases"`
	StartDate      string       `json:"startDate"`
	TargetDate     string       `json:"targetDate"`
	Duration       string       `json:"duration"`
}

// TaskPhase represents a phase of tasks
type TaskPhase struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

// Task represents a single task
type Task struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// Progress tracks project progress
type Progress struct {
	Overall int            `json:"overall"`
	Phases  map[string]int `json:"phases"`
}

// PullRequest represents a GitHub pull request
type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// GitHubData contains GitHub activity data
type GitHubData struct {
	Branches []Branch      `json:"branches"`
	Commits  []Commit      `json:"commits"`
	PRs      []PullRequest `json:"prs"`
	Pushes   []Push        `json:"pushes"`
}

// Branch represents a Git branch
type Branch struct {
	Name        string  `json:"name"`
	Status      string  `json:"status"`
	AheadCount  int     `json:"aheadCount"`
	BehindCount int     `json:"behindCount"`
	CommitCount int     `json:"commitCount"`
	LastCommit  *Commit `json:"lastCommit"`
	HasPR       bool    `json:"hasPR"`
	PRURL       string  `json:"prUrl"`
}

// Commit represents a Git commit
type Commit struct {
	Hash    string `json:"hash"`
	Message string `json:"message"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Branch  string `json:"branch"`
	PRURL   string `json:"prUrl"`
}

// Push represents a Git push
type Push struct {
	Branch      string `json:"branch"`
	Status      string `json:"status"`
	CommitCount int    `json:"commitCount"`
	Timestamp   string `json:"timestamp"`
}
