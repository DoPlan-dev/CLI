package statistics

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Storage manages historical statistics data
type Storage struct {
	projectRoot string
	storagePath string
}

// NewStorage creates a new statistics storage
func NewStorage(projectRoot string) *Storage {
	storageDir := filepath.Join(projectRoot, ".doplan", "stats")
	storagePath := filepath.Join(storageDir, "statistics.json")

	return &Storage{
		projectRoot: projectRoot,
		storagePath: storagePath,
	}
}

// HistoricalData represents a single statistics snapshot
type HistoricalData struct {
	Timestamp time.Time          `json:"timestamp"`
	Metrics   *StatisticsMetrics `json:"metrics"`
	Data      *StatisticsData    `json:"data"`
}

// Save stores statistics data
func (s *Storage) Save(metrics *StatisticsMetrics, data *StatisticsData) error {
	// Ensure directory exists
	dir := filepath.Dir(s.storagePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create storage directory: %w", err)
	}

	// Load existing data
	historical, err := s.LoadAll()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load existing data: %w", err)
	}

	// Add new entry
	entry := &HistoricalData{
		Timestamp: time.Now(),
		Metrics:   metrics,
		Data:      data,
	}

	historical = append(historical, entry)

	// Keep only last 100 entries
	if len(historical) > 100 {
		historical = historical[len(historical)-100:]
	}

	// Save to file
	fileData, err := json.MarshalIndent(historical, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(s.storagePath, fileData, 0644); err != nil {
		return fmt.Errorf("failed to write storage file: %w", err)
	}

	return nil
}

// LoadAll loads all historical data
func (s *Storage) LoadAll() ([]*HistoricalData, error) {
	if _, err := os.Stat(s.storagePath); os.IsNotExist(err) {
		return []*HistoricalData{}, nil
	}

	data, err := os.ReadFile(s.storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read storage file: %w", err)
	}

	var historical []*HistoricalData
	if err := json.Unmarshal(data, &historical); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return historical, nil
}

// LoadSince loads historical data since a given time
func (s *Storage) LoadSince(since time.Time) ([]*HistoricalData, error) {
	all, err := s.LoadAll()
	if err != nil {
		return nil, err
	}

	filtered := []*HistoricalData{}
	for _, entry := range all {
		if entry.Timestamp.After(since) || entry.Timestamp.Equal(since) {
			filtered = append(filtered, entry)
		}
	}

	return filtered, nil
}

// LoadRange loads historical data for a date range
func (s *Storage) LoadRange(start, end time.Time) ([]*HistoricalData, error) {
	all, err := s.LoadAll()
	if err != nil {
		return nil, err
	}

	filtered := []*HistoricalData{}
	for _, entry := range all {
		if (entry.Timestamp.After(start) || entry.Timestamp.Equal(start)) &&
			(entry.Timestamp.Before(end) || entry.Timestamp.Equal(end)) {
			filtered = append(filtered, entry)
		}
	}

	return filtered, nil
}

// GetLatest returns the most recent statistics entry
func (s *Storage) GetLatest() (*HistoricalData, error) {
	all, err := s.LoadAll()
	if err != nil {
		return nil, err
	}

	if len(all) == 0 {
		return nil, fmt.Errorf("no historical data available")
	}

	return all[len(all)-1], nil
}

// Clear removes all historical data
func (s *Storage) Clear() error {
	if _, err := os.Stat(s.storagePath); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(s.storagePath)
}
