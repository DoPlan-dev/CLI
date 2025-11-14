package config

import (
	"sync"
	"time"

	"github.com/DoPlan-dev/CLI/pkg/models"
)

// CacheManager manages in-memory caching for config and state
type CacheManager struct {
	configCache      *cacheEntry[*models.Config]
	stateCache       *cacheEntry[*models.State]
	configCacheMutex sync.RWMutex
	stateCacheMutex  sync.RWMutex
}

type cacheEntry[T any] struct {
	data      T
	timestamp time.Time
	ttl       time.Duration
}

// NewCacheManager creates a new cache manager
func NewCacheManager() *CacheManager {
	return &CacheManager{
		configCache: &cacheEntry[*models.Config]{
			ttl: 5 * time.Minute,
		},
		stateCache: &cacheEntry[*models.State]{
			ttl: 1 * time.Minute,
		},
	}
}

// GetConfig retrieves config from cache if valid, otherwise returns nil
func (cm *CacheManager) GetConfig() *models.Config {
	cm.configCacheMutex.RLock()
	defer cm.configCacheMutex.RUnlock()

	if cm.configCache.data == nil {
		return nil
	}

	if time.Since(cm.configCache.timestamp) > cm.configCache.ttl {
		return nil // Cache expired
	}

	return cm.configCache.data
}

// SetConfig stores config in cache
func (cm *CacheManager) SetConfig(config *models.Config) {
	cm.configCacheMutex.Lock()
	defer cm.configCacheMutex.Unlock()

	cm.configCache.data = config
	cm.configCache.timestamp = time.Now()
}

// InvalidateConfig clears the config cache
func (cm *CacheManager) InvalidateConfig() {
	cm.configCacheMutex.Lock()
	defer cm.configCacheMutex.Unlock()

	cm.configCache.data = nil
	cm.configCache.timestamp = time.Time{}
}

// GetState retrieves state from cache if valid, otherwise returns nil
func (cm *CacheManager) GetState() *models.State {
	cm.stateCacheMutex.RLock()
	defer cm.stateCacheMutex.RUnlock()

	if cm.stateCache.data == nil {
		return nil
	}

	if time.Since(cm.stateCache.timestamp) > cm.stateCache.ttl {
		return nil // Cache expired
	}

	return cm.stateCache.data
}

// SetState stores state in cache
func (cm *CacheManager) SetState(state *models.State) {
	cm.stateCacheMutex.Lock()
	defer cm.stateCacheMutex.Unlock()

	cm.stateCache.data = state
	cm.stateCache.timestamp = time.Now()
}

// InvalidateState clears the state cache
func (cm *CacheManager) InvalidateState() {
	cm.stateCacheMutex.Lock()
	defer cm.stateCacheMutex.Unlock()

	cm.stateCache.data = nil
	cm.stateCache.timestamp = time.Time{}
}

// InvalidateAll clears all caches
func (cm *CacheManager) InvalidateAll() {
	cm.InvalidateConfig()
	cm.InvalidateState()
}
