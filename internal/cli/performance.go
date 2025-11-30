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
// Returns nil for new projects (fast path)
func getOrCreateEngagementOrchestrator(projectPath string) (*EngagementOrchestrator, error) {
	if shouldUseFastPath(projectPath) {
		return nil, nil // Fast path: skip engagement system
	}

	// For existing projects, initialize engagement system
	return NewEngagementOrchestrator()
}
