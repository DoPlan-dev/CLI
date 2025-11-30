package rules

import (
	"sync"
	"time"
)

// CacheEntry represents a cached rule file
type CacheEntry struct {
	Data      []byte
	Timestamp time.Time
}

// RulesCache provides lazy loading and caching for rules
type RulesCache struct {
	cache    map[string]*CacheEntry
	mux      sync.RWMutex
	ttl      time.Duration
	loadTime time.Duration
	hits     int64
	misses   int64
}

// NewRulesCache creates a new rules cache with specified TTL
func NewRulesCache(ttl time.Duration) *RulesCache {
	return &RulesCache{
		cache: make(map[string]*CacheEntry),
		ttl:   ttl,
	}
}

// Get retrieves a rule file, loading it lazily if not cached
func (rc *RulesCache) Get(name string) ([]byte, error) {
	start := time.Now()
	defer func() {
		rc.loadTime += time.Since(start)
	}()

	// Check cache first
	rc.mux.RLock()
	entry, exists := rc.cache[name]
	if exists && time.Since(entry.Timestamp) < rc.ttl {
		// Cache hit
		data := make([]byte, len(entry.Data))
		copy(data, entry.Data)
		rc.mux.RUnlock()
		rc.hits++
		return data, nil
	}
	rc.mux.RUnlock()

	// Cache miss - load from embedded FS
	rc.mux.Lock()
	defer rc.mux.Unlock()

	// Double-check after acquiring write lock
	if entry, exists := rc.cache[name]; exists && time.Since(entry.Timestamp) < rc.ttl {
		data := make([]byte, len(entry.Data))
		copy(data, entry.Data)
		rc.hits++
		return data, nil
	}

	// Load from embedded filesystem
	data, err := ReadFile(name)
	if err != nil {
		rc.misses++
		return nil, err
	}

	// Cache it
	rc.cache[name] = &CacheEntry{
		Data:      data,
		Timestamp: time.Now(),
	}
	rc.misses++
	return data, nil
}

// GetDecompressed retrieves a decompressed rule file with caching
func (rc *RulesCache) GetDecompressed(name string) ([]byte, error) {
	start := time.Now()
	defer func() {
		rc.loadTime += time.Since(start)
	}()

	// Check cache first
	rc.mux.RLock()
	entry, exists := rc.cache[name]
	if exists && time.Since(entry.Timestamp) < rc.ttl {
		data := make([]byte, len(entry.Data))
		copy(data, entry.Data)
		rc.mux.RUnlock()
		rc.hits++
		return data, nil
	}
	rc.mux.RUnlock()

	// Cache miss - load and decompress
	rc.mux.Lock()
	defer rc.mux.Unlock()

	// Double-check after acquiring write lock
	if entry, exists := rc.cache[name]; exists && time.Since(entry.Timestamp) < rc.ttl {
		data := make([]byte, len(entry.Data))
		copy(data, entry.Data)
		rc.hits++
		return data, nil
	}

	// Load and decompress
	data, err := ReadFileDecompressed(name)
	if err != nil {
		rc.misses++
		return nil, err
	}

	// Cache it
	rc.cache[name] = &CacheEntry{
		Data:      data,
		Timestamp: time.Now(),
	}
	rc.misses++
	return data, nil
}

// Invalidate removes a specific rule from cache
func (rc *RulesCache) Invalidate(name string) {
	rc.mux.Lock()
	defer rc.mux.Unlock()
	delete(rc.cache, name)
}

// InvalidateAll clears the entire cache
func (rc *RulesCache) InvalidateAll() {
	rc.mux.Lock()
	defer rc.mux.Unlock()
	rc.cache = make(map[string]*CacheEntry)
}

// CleanExpired removes expired entries from cache
func (rc *RulesCache) CleanExpired() int {
	rc.mux.Lock()
	defer rc.mux.Unlock()

	now := time.Now()
	count := 0
	for name, entry := range rc.cache {
		if now.Sub(entry.Timestamp) >= rc.ttl {
			delete(rc.cache, name)
			count++
		}
	}
	return count
}

// GetMetrics returns performance metrics
func (rc *RulesCache) GetMetrics() Metrics {
	rc.mux.RLock()
	defer rc.mux.RUnlock()

	total := rc.hits + rc.misses
	hitRate := 0.0
	if total > 0 {
		hitRate = float64(rc.hits) / float64(total) * 100.0
	}

	return Metrics{
		CacheSize:     len(rc.cache),
		CacheHits:     rc.hits,
		CacheMisses:   rc.misses,
		HitRate:       hitRate,
		TotalLoadTime: rc.loadTime,
		AvgLoadTime:   rc.avgLoadTime(),
	}
}

// avgLoadTime calculates average load time
func (rc *RulesCache) avgLoadTime() time.Duration {
	total := rc.hits + rc.misses
	if total == 0 {
		return 0
	}
	return rc.loadTime / time.Duration(total)
}

// Metrics represents performance metrics for rules cache
type Metrics struct {
	CacheSize     int
	CacheHits     int64
	CacheMisses   int64
	HitRate       float64
	TotalLoadTime time.Duration
	AvgLoadTime   time.Duration
}

// DefaultRulesCache is the global rules cache instance
var (
	defaultRulesCache *RulesCache
	rulesCacheOnce    sync.Once
)

// GetDefaultCache returns the default rules cache (lazy initialization)
func GetDefaultCache() *RulesCache {
	rulesCacheOnce.Do(func() {
		// Default TTL: 5 minutes
		defaultRulesCache = NewRulesCache(5 * time.Minute)
	})
	return defaultRulesCache
}

// StartCleanupRoutine starts a background routine to clean expired entries
func (rc *RulesCache) StartCleanupRoutine(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			rc.CleanExpired()
		}
	}()
}
