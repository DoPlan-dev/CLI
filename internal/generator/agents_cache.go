package generator

import (
	"path/filepath"
	"sync"
	"time"
)

// AgentsCache provides lazy loading and caching for agents
type AgentsCache struct {
	agentsCache   map[string][]Agent   // keyed by project path, stores full agent list
	agentsTime    map[string]time.Time // timestamps for each project
	allAgents     []Agent              // cached GetAllAgents result
	allAgentsTime time.Time
	mux           sync.RWMutex
	ttl           time.Duration
	loadTime      time.Duration
	hits          int64
	misses        int64
}

// NewAgentsCache creates a new agents cache with specified TTL
func NewAgentsCache(ttl time.Duration) *AgentsCache {
	return &AgentsCache{
		agentsCache: make(map[string][]Agent),
		agentsTime:  make(map[string]time.Time),
		ttl:         ttl,
	}
}

// GetAgentsForProject retrieves agents for a specific project, loading lazily if not cached
func (ac *AgentsCache) GetAgentsForProject(projectPath string) ([]Agent, error) {
	start := time.Now()
	defer func() {
		ac.loadTime += time.Since(start)
	}()

	// Normalize project path
	normalizedPath := filepath.Clean(projectPath)
	agentsDir := filepath.Join(normalizedPath, ".do", "core", "agents")

	// Check cache first
	ac.mux.RLock()
	cachedAgents, exists := ac.agentsCache[normalizedPath]
	timestamp, timeExists := ac.agentsTime[normalizedPath]
	if exists && timeExists && time.Since(timestamp) < ac.ttl {
		// Cache hit - return cached copy
		agents := make([]Agent, len(cachedAgents))
		copy(agents, cachedAgents)
		ac.mux.RUnlock()
		ac.hits++
		return agents, nil
	}
	ac.mux.RUnlock()

	// Cache miss - load from directory
	ac.mux.Lock()
	defer ac.mux.Unlock()

	// Double-check after acquiring write lock
	if cachedAgents, exists := ac.agentsCache[normalizedPath]; exists {
		if timestamp, timeExists := ac.agentsTime[normalizedPath]; timeExists && time.Since(timestamp) < ac.ttl {
			agents := make([]Agent, len(cachedAgents))
			copy(agents, cachedAgents)
			ac.hits++
			return agents, nil
		}
	}

	// Load agents from directory
	agents, err := LoadAgentsFromDirectory(agentsDir)
	if err != nil {
		// Fallback to embedded agents
		agents, err = GetAllAgentsFileBased()
		if err != nil {
			ac.misses++
			return nil, err
		}
	}

	// Cache the full agent list
	ac.agentsCache[normalizedPath] = agents
	ac.agentsTime[normalizedPath] = time.Now()

	ac.misses++
	return agents, nil
}

// GetAllAgentsCached retrieves all agents with caching (for embedded agents)
func (ac *AgentsCache) GetAllAgentsCached() ([]Agent, error) {
	start := time.Now()
	defer func() {
		ac.loadTime += time.Since(start)
	}()

	// Check cache first
	ac.mux.RLock()
	if ac.allAgents != nil && time.Since(ac.allAgentsTime) < ac.ttl {
		agents := make([]Agent, len(ac.allAgents))
		copy(agents, ac.allAgents)
		ac.mux.RUnlock()
		ac.hits++
		return agents, nil
	}
	ac.mux.RUnlock()

	// Cache miss - load
	ac.mux.Lock()
	defer ac.mux.Unlock()

	// Double-check after acquiring write lock
	if ac.allAgents != nil && time.Since(ac.allAgentsTime) < ac.ttl {
		agents := make([]Agent, len(ac.allAgents))
		copy(agents, ac.allAgents)
		ac.hits++
		return agents, nil
	}

	// Load agents
	agents, err := GetAllAgentsFileBased()
	if err != nil {
		ac.misses++
		return nil, err
	}

	// Cache it
	ac.allAgents = agents
	ac.allAgentsTime = time.Now()
	ac.misses++
	return agents, nil
}

// InvalidateProject removes agents for a specific project from cache
func (ac *AgentsCache) InvalidateProject(projectPath string) {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	normalizedPath := filepath.Clean(projectPath)
	delete(ac.agentsCache, normalizedPath)
	delete(ac.agentsTime, normalizedPath)
}

// InvalidateAll clears the entire cache
func (ac *AgentsCache) InvalidateAll() {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	ac.agentsCache = make(map[string][]Agent)
	ac.agentsTime = make(map[string]time.Time)
	ac.allAgents = nil
	ac.allAgentsTime = time.Time{}
}

// CleanExpired removes expired entries from cache
func (ac *AgentsCache) CleanExpired() int {
	ac.mux.Lock()
	defer ac.mux.Unlock()

	now := time.Now()
	count := 0

	// Clean project-specific cache
	for path := range ac.agentsCache {
		if timestamp, exists := ac.agentsTime[path]; exists && now.Sub(timestamp) >= ac.ttl {
			delete(ac.agentsCache, path)
			delete(ac.agentsTime, path)
			count++
		}
	}

	// Clean all agents cache
	if ac.allAgents != nil && now.Sub(ac.allAgentsTime) >= ac.ttl {
		ac.allAgents = nil
		ac.allAgentsTime = time.Time{}
		count++
	}

	return count
}

// GetMetrics returns performance metrics
func (ac *AgentsCache) GetMetrics() AgentsMetrics {
	ac.mux.RLock()
	defer ac.mux.RUnlock()

	total := ac.hits + ac.misses
	hitRate := 0.0
	if total > 0 {
		hitRate = float64(ac.hits) / float64(total) * 100.0
	}

	return AgentsMetrics{
		CacheSize:       len(ac.agentsCache),
		AllAgentsCached: ac.allAgents != nil,
		CacheHits:       ac.hits,
		CacheMisses:     ac.misses,
		HitRate:         hitRate,
		TotalLoadTime:   ac.loadTime,
		AvgLoadTime:     ac.avgLoadTime(),
	}
}

// avgLoadTime calculates average load time
func (ac *AgentsCache) avgLoadTime() time.Duration {
	total := ac.hits + ac.misses
	if total == 0 {
		return 0
	}
	return ac.loadTime / time.Duration(total)
}

// AgentsMetrics represents performance metrics for agents cache
type AgentsMetrics struct {
	CacheSize       int
	AllAgentsCached bool
	CacheHits       int64
	CacheMisses     int64
	HitRate         float64
	TotalLoadTime   time.Duration
	AvgLoadTime     time.Duration
}

// DefaultAgentsCache is the global agents cache instance
var (
	defaultAgentsCache *AgentsCache
	agentsCacheOnce    sync.Once
)

// GetDefaultAgentsCache returns the default agents cache (lazy initialization)
func GetDefaultAgentsCache() *AgentsCache {
	agentsCacheOnce.Do(func() {
		// Default TTL: 5 minutes
		defaultAgentsCache = NewAgentsCache(5 * time.Minute)
	})
	return defaultAgentsCache
}

// StartCleanupRoutine starts a background routine to clean expired entries
func (ac *AgentsCache) StartCleanupRoutine(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			ac.CleanExpired()
		}
	}()
}
