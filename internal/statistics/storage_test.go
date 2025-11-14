package statistics

import (
	"os"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStorage(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	storage := NewStorage(projectRoot)

	assert.NotNil(t, storage)
	assert.Equal(t, projectRoot, storage.projectRoot)
	assert.Contains(t, storage.storagePath, ".doplan/stats/statistics.json")
}

func TestSaveAndLoad(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	metrics := &StatisticsMetrics{
		Velocity: &VelocityMetrics{
			FeaturesPerDay: 0.5,
		},
		Completion: &CompletionRates{
			Overall: 65,
		},
		CalculatedAt: time.Now(),
	}

	data := &StatisticsData{
		State: &StateData{
			TotalFeatures: 10,
		},
		CollectedAt: time.Now(),
	}

	// Save
	err := storage.Save(metrics, data)
	require.NoError(t, err)

	// Load all
	historical, err := storage.LoadAll()
	require.NoError(t, err)
	require.Len(t, historical, 1)

	assert.Equal(t, metrics.Velocity.FeaturesPerDay, historical[0].Metrics.Velocity.FeaturesPerDay)
	assert.Equal(t, data.State.TotalFeatures, historical[0].Data.State.TotalFeatures)
}

func TestLoadSince(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	now := time.Now()

	// Save multiple entries
	for i := 0; i < 5; i++ {
		metrics := &StatisticsMetrics{
			CalculatedAt: now.Add(time.Duration(i) * time.Hour),
		}
		data := &StatisticsData{
			CollectedAt: now.Add(time.Duration(i) * time.Hour),
		}
		require.NoError(t, storage.Save(metrics, data))
	}

	// Load since 2 hours ago
	since := now.Add(-2 * time.Hour)
	historical, err := storage.LoadSince(since)
	require.NoError(t, err)

	// Should have entries from the last 2 hours
	assert.GreaterOrEqual(t, len(historical), 2)
}

func TestLoadRange(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	baseTime := time.Now().Add(-10 * time.Hour) // Start in the past

	// Save multiple entries - storage.Save uses time.Now() for Timestamp
	// So we need to save them with small delays to ensure different timestamps
	for i := 0; i < 5; i++ {
		metrics := &StatisticsMetrics{
			CalculatedAt: baseTime.Add(time.Duration(i) * time.Hour),
		}
		data := &StatisticsData{
			CollectedAt: baseTime.Add(time.Duration(i) * time.Hour),
		}
		require.NoError(t, storage.Save(metrics, data))
		// Small delay to ensure different timestamps
		time.Sleep(10 * time.Millisecond)
	}

	// Load all to see what we have
	all, err := storage.LoadAll()
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(all), 3, "Need at least 3 entries for range test")

	// Use timestamps from actual saved data
	start := all[1].Timestamp
	end := all[3].Timestamp
	historical, err := storage.LoadRange(start, end)
	require.NoError(t, err)

	// Should have entries in range (inclusive)
	assert.GreaterOrEqual(t, len(historical), 2)
}

func TestGetLatest(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	// No data yet
	_, err := storage.GetLatest()
	assert.Error(t, err)

	// Save entry
	metrics := &StatisticsMetrics{
		CalculatedAt: time.Now(),
	}
	data := &StatisticsData{
		CollectedAt: time.Now(),
	}
	require.NoError(t, storage.Save(metrics, data))

	// Get latest
	latest, err := storage.GetLatest()
	require.NoError(t, err)
	assert.NotNil(t, latest)
}

func TestClear(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	// Save entry
	metrics := &StatisticsMetrics{CalculatedAt: time.Now()}
	data := &StatisticsData{CollectedAt: time.Now()}
	require.NoError(t, storage.Save(metrics, data))

	// Verify file exists
	_, err := os.Stat(storage.storagePath)
	require.NoError(t, err)

	// Clear
	require.NoError(t, storage.Clear())

	// Verify file removed
	_, err = os.Stat(storage.storagePath)
	assert.True(t, os.IsNotExist(err))
}

func TestSave_MaxEntries(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	storage := NewStorage(projectRoot)

	// Save more than 100 entries
	for i := 0; i < 150; i++ {
		metrics := &StatisticsMetrics{
			CalculatedAt: time.Now().Add(time.Duration(i) * time.Minute),
		}
		data := &StatisticsData{
			CollectedAt: time.Now().Add(time.Duration(i) * time.Minute),
		}
		require.NoError(t, storage.Save(metrics, data))
	}

	// Load all
	historical, err := storage.LoadAll()
	require.NoError(t, err)

	// Should only keep last 100
	assert.LessOrEqual(t, len(historical), 100)
}
