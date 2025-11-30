package cli

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/internal/rules"
)

// PerformanceMonitor tracks performance metrics across the CLI
type PerformanceMonitor struct {
	// Rules metrics
	RulesCache *rules.RulesCache

	// Agents metrics
	AgentsCache *generator.AgentsCache

	// Command execution metrics
	CommandExecutions map[string]*CommandMetrics
	commandMux        sync.RWMutex

	// Overall statistics
	StartTime time.Time
}

// CommandMetrics tracks metrics for a specific command
type CommandMetrics struct {
	Count         int64
	TotalDuration time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
	Errors        int64
	LastExecuted  time.Time
}

// NewPerformanceMonitor creates a new performance monitor
func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		RulesCache:        rules.GetDefaultCache(),
		AgentsCache:       generator.GetDefaultAgentsCache(),
		CommandExecutions: make(map[string]*CommandMetrics),
		StartTime:         time.Now(),
	}
}

// RecordCommandExecution records metrics for a command execution
func (pm *PerformanceMonitor) RecordCommandExecution(command string, duration time.Duration, err error) {
	pm.commandMux.Lock()
	defer pm.commandMux.Unlock()

	metrics, exists := pm.CommandExecutions[command]
	if !exists {
		metrics = &CommandMetrics{
			MinDuration: duration,
			MaxDuration: duration,
		}
		pm.CommandExecutions[command] = metrics
	}

	metrics.Count++
	metrics.TotalDuration += duration
	metrics.LastExecuted = time.Now()

	if duration < metrics.MinDuration || metrics.MinDuration == 0 {
		metrics.MinDuration = duration
	}
	if duration > metrics.MaxDuration {
		metrics.MaxDuration = duration
	}

	if err != nil {
		metrics.Errors++
	}
}

// GetCommandMetrics returns metrics for a specific command
func (pm *PerformanceMonitor) GetCommandMetrics(command string) *CommandMetrics {
	pm.commandMux.RLock()
	defer pm.commandMux.RUnlock()
	return pm.CommandExecutions[command]
}

// GetAverageCommandDuration returns the average execution time for a command
func (pm *PerformanceMonitor) GetAverageCommandDuration(command string) time.Duration {
	metrics := pm.GetCommandMetrics(command)
	if metrics == nil || metrics.Count == 0 {
		return 0
	}
	return metrics.TotalDuration / time.Duration(metrics.Count)
}

// GetReport returns a comprehensive performance report
func (pm *PerformanceMonitor) GetReport() Report {
	pm.commandMux.RLock()
	defer pm.commandMux.RUnlock()

	rulesMetrics := pm.RulesCache.GetMetrics()
	agentsMetrics := pm.AgentsCache.GetMetrics()

	// Aggregate command metrics
	totalCommands := int64(0)
	totalDuration := time.Duration(0)
	for _, metrics := range pm.CommandExecutions {
		totalCommands += metrics.Count
		totalDuration += metrics.TotalDuration
	}

	avgCommandDuration := time.Duration(0)
	if totalCommands > 0 {
		avgCommandDuration = totalDuration / time.Duration(totalCommands)
	}

	return Report{
		Uptime:               time.Since(pm.StartTime),
		RulesMetrics:         rulesMetrics,
		AgentsMetrics:        agentsMetrics,
		TotalCommands:        totalCommands,
		TotalCommandDuration: totalDuration,
		AvgCommandDuration:   avgCommandDuration,
		CommandMetrics:       pm.copyCommandMetrics(),
	}
}

// copyCommandMetrics creates a copy of command metrics for reporting
func (pm *PerformanceMonitor) copyCommandMetrics() map[string]*CommandMetrics {
	result := make(map[string]*CommandMetrics)
	for cmd, metrics := range pm.CommandExecutions {
		result[cmd] = &CommandMetrics{
			Count:         metrics.Count,
			TotalDuration: metrics.TotalDuration,
			MinDuration:   metrics.MinDuration,
			MaxDuration:   metrics.MaxDuration,
			Errors:        metrics.Errors,
			LastExecuted:  metrics.LastExecuted,
		}
	}
	return result
}

// Report represents a comprehensive performance report
type Report struct {
	Uptime               time.Duration
	RulesMetrics         rules.Metrics
	AgentsMetrics        generator.AgentsMetrics
	TotalCommands        int64
	TotalCommandDuration time.Duration
	AvgCommandDuration   time.Duration
	CommandMetrics       map[string]*CommandMetrics
}

// String returns a formatted string representation of the report
func (r *Report) String() string {
	var b strings.Builder

	b.WriteString("=== Performance Report ===\n\n")
	b.WriteString(fmt.Sprintf("Uptime: %s\n\n", r.Uptime.Round(time.Second)))

	// Rules metrics
	b.WriteString("Rules Cache:\n")
	b.WriteString(fmt.Sprintf("  Cache Size: %d entries\n", r.RulesMetrics.CacheSize))
	b.WriteString(fmt.Sprintf("  Hits: %d\n", r.RulesMetrics.CacheHits))
	b.WriteString(fmt.Sprintf("  Misses: %d\n", r.RulesMetrics.CacheMisses))
	b.WriteString(fmt.Sprintf("  Hit Rate: %.2f%%\n", r.RulesMetrics.HitRate))
	b.WriteString(fmt.Sprintf("  Avg Load Time: %s\n", r.RulesMetrics.AvgLoadTime))
	b.WriteString(fmt.Sprintf("  Total Load Time: %s\n\n", r.RulesMetrics.TotalLoadTime))

	// Agents metrics
	b.WriteString("Agents Cache:\n")
	b.WriteString(fmt.Sprintf("  Cache Size: %d projects\n", r.AgentsMetrics.CacheSize))
	b.WriteString(fmt.Sprintf("  All Agents Cached: %v\n", r.AgentsMetrics.AllAgentsCached))
	b.WriteString(fmt.Sprintf("  Hits: %d\n", r.AgentsMetrics.CacheHits))
	b.WriteString(fmt.Sprintf("  Misses: %d\n", r.AgentsMetrics.CacheMisses))
	b.WriteString(fmt.Sprintf("  Hit Rate: %.2f%%\n", r.AgentsMetrics.HitRate))
	b.WriteString(fmt.Sprintf("  Avg Load Time: %s\n", r.AgentsMetrics.AvgLoadTime))
	b.WriteString(fmt.Sprintf("  Total Load Time: %s\n\n", r.AgentsMetrics.TotalLoadTime))

	// Command metrics
	b.WriteString("Commands:\n")
	b.WriteString(fmt.Sprintf("  Total Executions: %d\n", r.TotalCommands))
	b.WriteString(fmt.Sprintf("  Total Duration: %s\n", r.TotalCommandDuration))
	b.WriteString(fmt.Sprintf("  Avg Duration: %s\n\n", r.AvgCommandDuration))

	// Per-command breakdown
	if len(r.CommandMetrics) > 0 {
		b.WriteString("Per-Command Metrics:\n")
		for cmd, metrics := range r.CommandMetrics {
			avg := time.Duration(0)
			if metrics.Count > 0 {
				avg = metrics.TotalDuration / time.Duration(metrics.Count)
			}
			b.WriteString(fmt.Sprintf("  %s:\n", cmd))
			b.WriteString(fmt.Sprintf("    Executions: %d\n", metrics.Count))
			b.WriteString(fmt.Sprintf("    Avg Duration: %s\n", avg))
			b.WriteString(fmt.Sprintf("    Min Duration: %s\n", metrics.MinDuration))
			b.WriteString(fmt.Sprintf("    Max Duration: %s\n", metrics.MaxDuration))
			b.WriteString(fmt.Sprintf("    Errors: %d\n", metrics.Errors))
			b.WriteString(fmt.Sprintf("    Last Executed: %s\n", metrics.LastExecuted.Format(time.RFC3339)))
		}
	}

	return b.String()
}

// DefaultPerformanceMonitor is the global performance monitor instance
var (
	defaultPerformanceMonitor *PerformanceMonitor
	performanceMonitorOnce    sync.Once
)

// GetDefaultPerformanceMonitor returns the default performance monitor (lazy initialization)
func GetDefaultPerformanceMonitor() *PerformanceMonitor {
	performanceMonitorOnce.Do(func() {
		defaultPerformanceMonitor = NewPerformanceMonitor()
		// Start cleanup routines
		defaultPerformanceMonitor.RulesCache.StartCleanupRoutine(1 * time.Minute)
		defaultPerformanceMonitor.AgentsCache.StartCleanupRoutine(1 * time.Minute)
	})
	return defaultPerformanceMonitor
}
