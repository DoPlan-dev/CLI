package cli

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
)

// Performance optimizations for command execution

var (
	// memoryCardCache caches the memory card to avoid repeated file I/O
	memoryCardCache     *MemoryCard
	memoryCardCacheTime time.Time
	memoryCardCacheMux  sync.RWMutex
	cacheExpiry         = 5 * time.Second // Cache expires after 5 seconds
)

// isNewProject checks if this is a new project (no memory card exists)
// This enables fast path for commands that don't need engagement systems
func isNewProject(projectPath string) bool {
	memoryCardPath := filepath.Join(projectPath, ".do", "system", "memory_card.json")
	return !utils.PathExists(memoryCardPath)
}

// loadMemoryCardCached loads memory card with caching to reduce file I/O
func loadMemoryCardCached() (*MemoryCard, error) {
	memoryCardCacheMux.RLock()
	// Check if cache is valid
	if memoryCardCache != nil && time.Since(memoryCardCacheTime) < cacheExpiry {
		mc := memoryCardCache
		memoryCardCacheMux.RUnlock()
		return mc, nil
	}
	memoryCardCacheMux.RUnlock()

	// Load from disk
	memoryCardCacheMux.Lock()
	defer memoryCardCacheMux.Unlock()

	// Double-check after acquiring write lock
	if memoryCardCache != nil && time.Since(memoryCardCacheTime) < cacheExpiry {
		return memoryCardCache, nil
	}

	// Load fresh
	mc, err := LoadMemoryCard()
	if err != nil {
		return nil, err
	}

	memoryCardCache = mc
	memoryCardCacheTime = time.Now()
	return mc, nil
}

// invalidateMemoryCardCache invalidates the memory card cache
// Call this after saving the memory card
func invalidateMemoryCardCache() {
	memoryCardCacheMux.Lock()
	defer memoryCardCacheMux.Unlock()
	memoryCardCache = nil
	memoryCardCacheTime = time.Time{}
}

// shouldUseFastPath determines if we should use fast path (skip engagement system)
// Fast path is used for new projects where engagement systems aren't needed
func shouldUseFastPath(projectPath string) bool {
	return isNewProject(projectPath)
}

// getOrCreateEngagementOrchestrator gets or creates engagement orchestrator
// Always initializes engagement system to ensure achievements work from the start
func getOrCreateEngagementOrchestrator(projectPath string) (*EngagementOrchestrator, error) {
	// Always initialize engagement system - it will create memory card if needed
	orchestrator, err := NewEngagementOrchestrator()
	if err != nil {
		// If memory card doesn't exist, create a new one and try again
		if _, loadErr := LoadMemoryCard(); loadErr != nil {
			// Create new memory card for new projects
			mc := &MemoryCard{
				FirstMet:          time.Now(),
				ProjectsCount:     0,
				CommandUsage:      make(map[string]int),
				Score:             0,
				Achievements:      []Achievement{},
				RelationshipLevel: 0,
				TrustLevel:        0,
				EngagementScore:   0.0,
			}
			if saveErr := SaveMemoryCard(mc); saveErr == nil {
				// Retry initialization after creating memory card
				orchestrator, err = NewEngagementOrchestrator()
			}
		}
	}
	return orchestrator, err
}
