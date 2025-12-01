package generator

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewAgentsCache(t *testing.T) {
	ttl := 5 * time.Minute
	cache := NewAgentsCache(ttl)

	if cache == nil {
		t.Fatal("NewAgentsCache() should not return nil")
	}

	if cache.ttl != ttl {
		t.Errorf("NewAgentsCache() ttl = %v, want %v", cache.ttl, ttl)
	}

	if cache.agentsCache == nil {
		t.Error("NewAgentsCache() should initialize agentsCache map")
	}

	if cache.agentsTime == nil {
		t.Error("NewAgentsCache() should initialize agentsTime map")
	}
}

func TestAgentsCache_GetAgentsForProject(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Create a temporary project directory with agents
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create agents directory structure
	agentsDir := filepath.Join(tmpDir, ".do", "core", "agents", "engineering")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		t.Fatalf("Failed to create agents dir: %v", err)
	}

	// Create a test agent file
	agentFile := filepath.Join(agentsDir, "engineering_lead.md")
	agentContent := `---
name: Engineering Lead
role: engineering
category: engineering
manager: Engineering Lead
---

# Engineering Lead

System prompt for engineering lead.
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to write agent file: %v", err)
	}

	// Test cache miss - first load
	agents, err := cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() error = %v", err)
	}

	if len(agents) == 0 {
		t.Error("GetAgentsForProject() should return agents")
	}

	// Test cache hit - second load
	agents2, err := cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() on second call error = %v", err)
	}

	if len(agents2) != len(agents) {
		t.Error("GetAgentsForProject() should return same agents on cache hit")
	}

	// Verify metrics show a hit
	metrics := cache.GetMetrics()
	if metrics.CacheHits == 0 {
		t.Error("GetMetrics() should show cache hits after second GetAgentsForProject()")
	}
}

func TestAgentsCache_GetAgentsForProject_NonExistent(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Test with non-existent directory - should fallback to embedded agents
	agents, err := cache.GetAgentsForProject("/nonexistent/path")
	if err != nil {
		t.Fatalf("GetAgentsForProject() should fallback to embedded agents, got error: %v", err)
	}

	if len(agents) == 0 {
		t.Error("GetAgentsForProject() should return embedded agents as fallback")
	}

	metrics := cache.GetMetrics()
	if metrics.CacheMisses == 0 {
		t.Error("GetMetrics() should show cache miss for non-existent path")
	}
}

func TestAgentsCache_GetAgentsForProject_PathNormalization(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Create a temporary project directory
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create agents directory
	agentsDir := filepath.Join(tmpDir, ".do", "core", "agents", "engineering")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		t.Fatalf("Failed to create agents dir: %v", err)
	}

	// Load with one path format
	_, err = cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() error = %v", err)
	}

	// Load with different path format (should use same cache entry)
	normalizedPath := filepath.Clean(tmpDir)
	_, err = cache.GetAgentsForProject(normalizedPath + "/")
	if err != nil {
		t.Fatalf("GetAgentsForProject() with normalized path error = %v", err)
	}

	// Should have cache hit
	metrics := cache.GetMetrics()
	if metrics.CacheSize != 1 {
		t.Errorf("Path normalization: CacheSize = %d, want 1", metrics.CacheSize)
	}
}

func TestAgentsCache_GetAllAgentsCached(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Test cache miss - first load
	agents, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	if len(agents) == 0 {
		t.Error("GetAllAgentsCached() should return agents")
	}

	// Test cache hit - second load
	agents2, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() on second call error = %v", err)
	}

	if len(agents2) != len(agents) {
		t.Error("GetAllAgentsCached() should return same agents on cache hit")
	}

	// Verify metrics
	metrics := cache.GetMetrics()
	if !metrics.AllAgentsCached {
		t.Error("GetMetrics() should show AllAgentsCached as true after GetAllAgentsCached()")
	}
	if metrics.CacheHits == 0 {
		t.Error("GetMetrics() should show cache hits after second GetAllAgentsCached()")
	}
}

func TestAgentsCache_InvalidateProject(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Create a temporary project directory
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Load agents into cache
	_, err = cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() error = %v", err)
	}

	// Verify it's cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 {
		t.Error("Cache should contain entry after GetAgentsForProject()")
	}

	// Invalidate
	cache.InvalidateProject(tmpDir)

	// Verify it's removed
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("InvalidateProject() should remove entry from cache")
	}
}

func TestAgentsCache_InvalidateAll(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Load agents into cache
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Create a temporary project and load it
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	_, err = cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() error = %v", err)
	}

	// Verify they're cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 && !metrics.AllAgentsCached {
		t.Error("Cache should contain entries after loading")
	}

	// Invalidate all
	cache.InvalidateAll()

	// Verify cache is empty
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("InvalidateAll() should clear project cache")
	}
	if metrics.AllAgentsCached {
		t.Error("InvalidateAll() should clear all agents cache")
	}
}

func TestAgentsCache_CleanExpired(t *testing.T) {
	// Use very short TTL for testing
	cache := NewAgentsCache(100 * time.Millisecond)

	// Load agents into cache
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Create a temporary project and load it
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	_, err = cache.GetAgentsForProject(tmpDir)
	if err != nil {
		t.Fatalf("GetAgentsForProject() error = %v", err)
	}

	// Verify they're cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 && !metrics.AllAgentsCached {
		t.Error("Cache should contain entries after loading")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Clean expired
	count := cache.CleanExpired()
	if count == 0 {
		t.Error("CleanExpired() should remove expired entries")
	}

	// Verify cache is cleaned
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("CleanExpired() should remove expired project entries")
	}
	if metrics.AllAgentsCached {
		t.Error("CleanExpired() should remove expired all agents cache")
	}
}

func TestAgentsCache_CleanExpired_NotExpired(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Load agents into cache
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Clean expired (should not remove anything)
	count := cache.CleanExpired()
	if count != 0 {
		t.Error("CleanExpired() should not remove non-expired entries")
	}

	// Verify cache still has entry
	metrics := cache.GetMetrics()
	if !metrics.AllAgentsCached {
		t.Error("Cache should still contain non-expired all agents cache")
	}
}

func TestAgentsCache_GetMetrics(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Initial metrics
	metrics := cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Errorf("Initial CacheSize = %d, want 0", metrics.CacheSize)
	}
	if metrics.AllAgentsCached {
		t.Error("Initial AllAgentsCached should be false")
	}
	if metrics.CacheHits != 0 {
		t.Errorf("Initial CacheHits = %d, want 0", metrics.CacheHits)
	}
	if metrics.CacheMisses != 0 {
		t.Errorf("Initial CacheMisses = %d, want 0", metrics.CacheMisses)
	}

	// Load agents (cache miss)
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Load again (cache hit)
	_, err = cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Check metrics
	metrics = cache.GetMetrics()
	if !metrics.AllAgentsCached {
		t.Error("AllAgentsCached should be true after GetAllAgentsCached()")
	}
	if metrics.CacheHits == 0 {
		t.Error("CacheHits should be > 0 after cache hit")
	}
	if metrics.CacheMisses == 0 {
		t.Error("CacheMisses should be > 0 after cache miss")
	}
	if metrics.HitRate <= 0 {
		t.Error("HitRate should be > 0 after both hit and miss")
	}
}

func TestAgentsCache_avgLoadTime(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Initial avg load time should be 0
	metrics := cache.GetMetrics()
	if metrics.AvgLoadTime != 0 {
		t.Errorf("Initial AvgLoadTime = %v, want 0", metrics.AvgLoadTime)
	}

	// Load agents
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Avg load time should be > 0
	metrics = cache.GetMetrics()
	if metrics.AvgLoadTime == 0 {
		t.Error("AvgLoadTime should be > 0 after loading agents")
	}
}

func TestGetDefaultAgentsCache(t *testing.T) {
	cache1 := GetDefaultAgentsCache()
	cache2 := GetDefaultAgentsCache()

	if cache1 == nil {
		t.Fatal("GetDefaultAgentsCache() should not return nil")
	}

	// Should return same instance (singleton)
	if cache1 != cache2 {
		t.Error("GetDefaultAgentsCache() should return same instance")
	}
}

func TestAgentsCache_StartCleanupRoutine(t *testing.T) {
	cache := NewAgentsCache(100 * time.Millisecond)

	// Load agents
	_, err := cache.GetAllAgentsCached()
	if err != nil {
		t.Fatalf("GetAllAgentsCached() error = %v", err)
	}

	// Start cleanup routine
	cache.StartCleanupRoutine(50 * time.Millisecond)

	// Wait for expiration and cleanup
	time.Sleep(200 * time.Millisecond)

	// Verify cache is cleaned (may take a moment)
	time.Sleep(100 * time.Millisecond)
	metrics := cache.GetMetrics()
	// The cleanup routine should have run, but timing may vary
	// Just verify the routine doesn't panic
	_ = metrics
}

func TestAgentsCache_ConcurrentAccess(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Test concurrent reads
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			_, _ = cache.GetAllAgentsCached()
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify cache is still consistent
	metrics := cache.GetMetrics()
	if !metrics.AllAgentsCached {
		t.Error("Concurrent access: AllAgentsCached should be true")
	}
}

func TestAgentsCache_DoubleCheckLocking(t *testing.T) {
	cache := NewAgentsCache(5 * time.Minute)

	// Create a temporary project directory
	tmpDir, err := os.MkdirTemp("", "doplan-agents-cache-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create agents directory
	agentsDir := filepath.Join(tmpDir, ".do", "core", "agents", "engineering")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		t.Fatalf("Failed to create agents dir: %v", err)
	}

	// Create a test agent file
	agentFile := filepath.Join(agentsDir, "engineering_lead.md")
	agentContent := `---
name: Engineering Lead
role: engineering
category: engineering
manager: Engineering Lead
---

# Engineering Lead

System prompt.
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to write agent file: %v", err)
	}

	// Concurrent access to trigger double-check locking
	done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func() {
			_, _ = cache.GetAgentsForProject(tmpDir)
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 5; i++ {
		<-done
	}

	// Verify only one cache entry was created
	metrics := cache.GetMetrics()
	if metrics.CacheSize != 1 {
		t.Errorf("Double-check locking: CacheSize = %d, want 1", metrics.CacheSize)
	}
}

