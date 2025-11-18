package dashboard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ProgressData represents progress data from a progress.json file
type ProgressData struct {
	FeatureID   string         `json:"featureID,omitempty"`
	FeatureName string         `json:"featureName,omitempty"`
	PhaseID     string         `json:"phaseID,omitempty"`
	PhaseName   string         `json:"phaseName,omitempty"`
	Status      string         `json:"status"`
	Progress    int            `json:"progress"`
	Branch      string         `json:"branch,omitempty"`
	PR          *PRData        `json:"pr,omitempty"`
	Tasks       []TaskProgress `json:"tasks,omitempty"`
	LastUpdated time.Time      `json:"lastUpdated,omitempty"`
}

// PRData represents PR information in progress.json
type PRData struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// TaskProgress represents task completion data
type TaskProgress struct {
	Name        string    `json:"name"`
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
}

// ProgressParser parses progress.json files from the project
type ProgressParser struct {
	projectRoot string
}

// NewProgressParser creates a new progress parser
func NewProgressParser(projectRoot string) *ProgressParser {
	return &ProgressParser{
		projectRoot: projectRoot,
	}
}

// ReadProgressFiles reads all progress.json files from feature directories
func (p *ProgressParser) ReadProgressFiles() (map[string]*ProgressData, error) {
	progressMap := make(map[string]*ProgressData)
	doplanDir := filepath.Join(p.projectRoot, "doplan")

	// Check if doplan directory exists
	if _, err := os.Stat(doplanDir); os.IsNotExist(err) {
		return progressMap, nil
	}

	// Walk through doplan directory
	err := filepath.Walk(doplanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors, continue walking
		}

		// Look for progress.json files
		if info.Name() == "progress.json" {
			data, err := p.readProgressFile(path)
			if err != nil {
				return nil // Skip invalid files, continue
			}
			if data != nil {
				// Use featureID or path as key
				key := data.FeatureID
				if key == "" {
					key = path
				}
				progressMap[key] = data
			}
		}

		return nil
	})

	return progressMap, err
}

// readProgressFile reads and parses a single progress.json file
func (p *ProgressParser) readProgressFile(path string) (*ProgressData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var progress ProgressData
	if err := json.Unmarshal(data, &progress); err != nil {
		return nil, fmt.Errorf("failed to parse progress.json: %w", err)
	}

	// Set last updated time from file modification time
	if info, err := os.Stat(path); err == nil {
		progress.LastUpdated = info.ModTime()
	}

	// Extract phase/feature info from path
	relPath, err := filepath.Rel(filepath.Join(p.projectRoot, "doplan"), path)
	if err == nil {
		parts := filepath.SplitList(relPath)
		if len(parts) >= 2 {
			progress.PhaseID = parts[0]
			if len(parts) >= 3 {
				progress.FeatureID = parts[1]
			}
		}
	}

	return &progress, nil
}

// ReadPhaseProgressFiles reads all phase-progress.json files
func (p *ProgressParser) ReadPhaseProgressFiles() (map[string]*ProgressData, error) {
	progressMap := make(map[string]*ProgressData)
	doplanDir := filepath.Join(p.projectRoot, "doplan")

	if _, err := os.Stat(doplanDir); os.IsNotExist(err) {
		return progressMap, nil
	}

	err := filepath.Walk(doplanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.Name() == "phase-progress.json" {
			data, err := p.readProgressFile(path)
			if err != nil {
				return nil
			}
			if data != nil {
				key := data.PhaseID
				if key == "" {
					key = path
				}
				progressMap[key] = data
			}
		}

		return nil
	})

	return progressMap, err
}

// GetTaskCompletionHistory extracts task completion history from tasks.md files
func (p *ProgressParser) GetTaskCompletionHistory() ([]TaskProgress, error) {
	var tasks []TaskProgress
	doplanDir := filepath.Join(p.projectRoot, "doplan")

	if _, err := os.Stat(doplanDir); os.IsNotExist(err) {
		return tasks, nil
	}

	err := filepath.Walk(doplanDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.Name() == "tasks.md" {
			taskData, err := p.parseTasksFile(path)
			if err == nil {
				tasks = append(tasks, taskData...)
			}
		}

		return nil
	})

	return tasks, err
}

// parseTasksFile parses a tasks.md file to extract task completion data
func (p *ProgressParser) parseTasksFile(path string) ([]TaskProgress, error) {
	_, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []TaskProgress
	
	// Simple parsing: look for - [x] or - [X] for completed tasks
	// This is a simplified parser - can be enhanced later
	// For now, we'll extract from the state model which has better task data
	
	return tasks, nil
}
