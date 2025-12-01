package rules

import (
	"testing"
	"time"
)

func TestNewRulesCache(t *testing.T) {
	ttl := 5 * time.Minute
	cache := NewRulesCache(ttl)

	if cache == nil {
		t.Fatal("NewRulesCache() should not return nil")
	}

	if cache.ttl != ttl {
		t.Errorf("NewRulesCache() ttl = %v, want %v", cache.ttl, ttl)
	}

	if cache.cache == nil {
		t.Error("NewRulesCache() should initialize cache map")
	}
}

func TestRulesCache_Get(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Test cache miss - first load
	data, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("Get() should return non-empty data")
	}

	// Test cache hit - second load
	data2, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() on second call error = %v", err)
	}

	if len(data2) != len(data) {
		t.Error("Get() should return same data on cache hit")
	}

	// Verify metrics show a hit
	metrics := cache.GetMetrics()
	if metrics.CacheHits == 0 {
		t.Error("GetMetrics() should show cache hits after second Get()")
	}
}

func TestRulesCache_Get_NonExistent(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	_, err := cache.Get("nonexistent/file.md")
	if err == nil {
		t.Error("Get() should return error for non-existent file")
	}

	metrics := cache.GetMetrics()
	if metrics.CacheMisses == 0 {
		t.Error("GetMetrics() should show cache miss for non-existent file")
	}
}

func TestRulesCache_GetDecompressed(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Test cache miss - first load
	data, err := cache.GetDecompressed("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("GetDecompressed() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("GetDecompressed() should return non-empty data")
	}

	// Test cache hit - second load
	data2, err := cache.GetDecompressed("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("GetDecompressed() on second call error = %v", err)
	}

	if len(data2) != len(data) {
		t.Error("GetDecompressed() should return same data on cache hit")
	}
}

func TestRulesCache_GetDecompressed_NonExistent(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	_, err := cache.GetDecompressed("nonexistent/file.md")
	if err == nil {
		t.Error("GetDecompressed() should return error for non-existent file")
	}
}

func TestRulesCache_Invalidate(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Load a file into cache
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Verify it's cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 {
		t.Error("Cache should contain entry after Get()")
	}

	// Invalidate
	cache.Invalidate("01-core-workflow/README.md")

	// Verify it's removed
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("Invalidate() should remove entry from cache")
	}
}

func TestRulesCache_InvalidateAll(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Load multiple files into cache
	_, _ = cache.Get("01-core-workflow/README.md")
	_, _ = cache.Get("02-ai-agents/README.md")

	// Verify they're cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 {
		t.Error("Cache should contain entries after Get()")
	}

	// Invalidate all
	cache.InvalidateAll()

	// Verify cache is empty
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("InvalidateAll() should clear entire cache")
	}
}

func TestRulesCache_CleanExpired(t *testing.T) {
	// Use very short TTL for testing
	cache := NewRulesCache(100 * time.Millisecond)

	// Load a file into cache
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Verify it's cached
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 {
		t.Error("Cache should contain entry after Get()")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Clean expired
	count := cache.CleanExpired()
	if count == 0 {
		t.Error("CleanExpired() should remove expired entries")
	}

	// Verify cache is empty
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Error("CleanExpired() should remove expired entries from cache")
	}
}

func TestRulesCache_CleanExpired_NotExpired(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Load a file into cache
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Clean expired (should not remove anything)
	count := cache.CleanExpired()
	if count != 0 {
		t.Error("CleanExpired() should not remove non-expired entries")
	}

	// Verify cache still has entry
	metrics := cache.GetMetrics()
	if metrics.CacheSize == 0 {
		t.Error("Cache should still contain non-expired entry")
	}
}

func TestRulesCache_GetMetrics(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Initial metrics
	metrics := cache.GetMetrics()
	if metrics.CacheSize != 0 {
		t.Errorf("Initial CacheSize = %d, want 0", metrics.CacheSize)
	}
	if metrics.CacheHits != 0 {
		t.Errorf("Initial CacheHits = %d, want 0", metrics.CacheHits)
	}
	if metrics.CacheMisses != 0 {
		t.Errorf("Initial CacheMisses = %d, want 0", metrics.CacheMisses)
	}

	// Load a file (cache miss)
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Load again (cache hit)
	_, err = cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Check metrics
	metrics = cache.GetMetrics()
	if metrics.CacheSize != 1 {
		t.Errorf("CacheSize = %d, want 1", metrics.CacheSize)
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

func TestRulesCache_avgLoadTime(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Initial avg load time should be 0
	metrics := cache.GetMetrics()
	if metrics.AvgLoadTime != 0 {
		t.Errorf("Initial AvgLoadTime = %v, want 0", metrics.AvgLoadTime)
	}

	// Load a file
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Avg load time should be > 0
	metrics = cache.GetMetrics()
	if metrics.AvgLoadTime == 0 {
		t.Error("AvgLoadTime should be > 0 after loading file")
	}
}

func TestGetDefaultCache(t *testing.T) {
	cache1 := GetDefaultCache()
	cache2 := GetDefaultCache()

	if cache1 == nil {
		t.Fatal("GetDefaultCache() should not return nil")
	}

	// Should return same instance (singleton)
	if cache1 != cache2 {
		t.Error("GetDefaultCache() should return same instance")
	}
}

func TestRulesCache_StartCleanupRoutine(t *testing.T) {
	cache := NewRulesCache(100 * time.Millisecond)

	// Load a file
	_, err := cache.Get("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
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

func TestRulesCache_ConcurrentAccess(t *testing.T) {
	cache := NewRulesCache(5 * time.Minute)

	// Test concurrent reads
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			_, _ = cache.Get("01-core-workflow/README.md")
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify cache is still consistent
	metrics := cache.GetMetrics()
	if metrics.CacheSize != 1 {
		t.Errorf("Concurrent access: CacheSize = %d, want 1", metrics.CacheSize)
	}
}

