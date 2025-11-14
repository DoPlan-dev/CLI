package github

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// GitHubSync syncs GitHub data
type GitHubSync struct {
	repoPath string
	cache    *githubCache
}

type githubCache struct {
	data      *models.GitHubData
	timestamp time.Time
	ttl       time.Duration
	mu        sync.RWMutex
}

// NewGitHubSync creates a new GitHub sync
func NewGitHubSync(repoPath string) *GitHubSync {
	return &GitHubSync{
		repoPath: repoPath,
		cache: &githubCache{
			ttl: 5 * time.Minute,
		},
	}
}

// Sync fetches GitHub data and updates github-data.json (with caching and parallel fetching)
func (gs *GitHubSync) Sync() (*models.GitHubData, error) {
	// Check cache first
	gs.cache.mu.RLock()
	if gs.cache.data != nil && time.Since(gs.cache.timestamp) < gs.cache.ttl {
		cached := gs.cache.data
		gs.cache.mu.RUnlock()
		return cached, nil
	}
	gs.cache.mu.RUnlock()

	data := &models.GitHubData{
		Branches: []models.Branch{},
		Commits:  []models.Commit{},
		PRs:      []models.PullRequest{},
		Pushes:   []models.Push{},
	}

	// Parallel fetching using goroutines
	var wg sync.WaitGroup

	// Fetch branches
	wg.Add(1)
	go func() {
		defer wg.Done()
		branches, err := gs.fetchBranches()
		if err == nil {
			data.Branches = branches
		}
	}()

	// Fetch commits
	wg.Add(1)
	go func() {
		defer wg.Done()
		commits, err := gs.fetchCommits()
		if err == nil {
			data.Commits = commits
		}
	}()

	// Fetch PRs
	wg.Add(1)
	go func() {
		defer wg.Done()
		prs, err := gs.fetchPRs()
		if err == nil {
			data.PRs = prs
		}
	}()

	// Wait for all fetches to complete
	wg.Wait()

	// Save to file
	if err := gs.saveData(data); err != nil {
		return nil, err
	}

	// Update cache
	gs.cache.mu.Lock()
	gs.cache.data = data
	gs.cache.timestamp = time.Now()
	gs.cache.mu.Unlock()

	return data, nil
}

func (gs *GitHubSync) fetchBranches() ([]models.Branch, error) {
	cmd := exec.Command("git", "branch", "-r")
	cmd.Dir = gs.repoPath
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	
	var branches []models.Branch
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "HEAD") {
			continue
		}
		
		// Remove "origin/" prefix
		branchName := strings.TrimPrefix(line, "origin/")
		
		branch := models.Branch{
			Name:        branchName,
			Status:      "active",
			CommitCount: 0,
		}
		
		branches = append(branches, branch)
	}
	
	return branches, nil
}

func (gs *GitHubSync) fetchCommits() ([]models.Commit, error) {
	cmd := exec.Command("git", "log", "--pretty=format:%H|%s|%an|%ad", "--date=iso", "-20")
	cmd.Dir = gs.repoPath
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	
	var commits []models.Commit
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}
		
		commit := models.Commit{
			Hash:    parts[0],
			Message: parts[1],
			Author:  parts[2],
			Date:    parts[3],
		}
		
		commits = append(commits, commit)
	}
	
	return commits, nil
}

func (gs *GitHubSync) fetchPRs() ([]models.PullRequest, error) {
	// Try to use GitHub CLI
	cmd := exec.Command("gh", "pr", "list", "--json", "number,title,url,state")
	cmd.Dir = gs.repoPath
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		// GitHub CLI not available or not authenticated
		return []models.PullRequest{}, nil
	}
	
	var prs []models.PullRequest
	
	// Parse JSON output
	var ghPRs []struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
		URL    string `json:"url"`
		State  string `json:"state"`
	}
	
	if err := json.Unmarshal(output, &ghPRs); err != nil {
		return []models.PullRequest{}, nil
	}
	
	for _, ghPR := range ghPRs {
		pr := models.PullRequest{
			Number: ghPR.Number,
			Title:  ghPR.Title,
			URL:    ghPR.URL,
			Status: ghPR.State,
		}
		prs = append(prs, pr)
	}
	
	return prs, nil
}

func (gs *GitHubSync) saveData(data *models.GitHubData) error {
	dataPath := filepath.Join(gs.repoPath, "doplan", "github-data.json")
	
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(dataPath), 0755); err != nil {
		return err
	}
	
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(dataPath, fileData, 0644)
}

// LoadData loads GitHub data from file
func (gs *GitHubSync) LoadData() (*models.GitHubData, error) {
	dataPath := filepath.Join(gs.repoPath, "doplan", "github-data.json")
	
	data, err := os.ReadFile(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.GitHubData{
				Branches: []models.Branch{},
				Commits:  []models.Commit{},
				PRs:      []models.PullRequest{},
				Pushes:   []models.Push{},
			}, nil
		}
		return nil, err
	}
	
	var githubData models.GitHubData
	if err := json.Unmarshal(data, &githubData); err != nil {
		return nil, err
	}
	
	return &githubData, nil
}

