package checkpoint

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/internal/config"
	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/fatih/color"
)

// Checkpoint represents a project checkpoint
type Checkpoint struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"` // "feature", "phase", "manual"
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	State       *models.State  `json:"state"`
	Config      *models.Config `json:"config"`
	FilePath    string         `json:"filePath"`
}

// CheckpointManager manages checkpoints
type CheckpointManager struct {
	projectRoot    string
	checkpointsDir string
}

// NewCheckpointManager creates a new checkpoint manager
func NewCheckpointManager(projectRoot string) *CheckpointManager {
	return &CheckpointManager{
		projectRoot:    projectRoot,
		checkpointsDir: filepath.Join(projectRoot, ".doplan", "checkpoints"),
	}
}

// CreateCheckpoint creates a checkpoint
func (cm *CheckpointManager) CreateCheckpoint(checkpointType, name, description string) (*Checkpoint, error) {
	// Load current state and config
	cfgMgr := config.NewManager(cm.projectRoot)
	state, err := cfgMgr.LoadState()
	if err != nil {
		return nil, fmt.Errorf("failed to load state: %w", err)
	}

	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Create checkpoint
	checkpoint := &Checkpoint{
		ID:          generateCheckpointID(),
		Type:        checkpointType,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		State:       state,
		Config:      cfg,
	}

	// Ensure directories exist
	if err := os.MkdirAll(filepath.Join(cm.checkpointsDir, "metadata"), 0755); err != nil {
		return nil, fmt.Errorf("failed to create metadata directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(cm.checkpointsDir, "archives"), 0755); err != nil {
		return nil, fmt.Errorf("failed to create archives directory: %w", err)
	}

	// Create archive
	archivePath, err := cm.createArchive(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create archive: %w", err)
	}

	checkpoint.FilePath = archivePath

	// Save checkpoint metadata
	if err := cm.saveCheckpointMetadata(checkpoint); err != nil {
		return nil, fmt.Errorf("failed to save metadata: %w", err)
	}

	color.Green("✅ Checkpoint created: %s (%s)\n", checkpoint.ID, checkpoint.Name)

	return checkpoint, nil
}

// AutoCreateFeatureCheckpoint automatically creates checkpoint for feature
func (cm *CheckpointManager) AutoCreateFeatureCheckpoint(feature *models.Feature) error {
	cfgMgr := config.NewManager(cm.projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		return err
	}

	// Check if auto-checkpoint is enabled
	if !cfg.Checkpoint.AutoFeature {
		return nil
	}

	checkpoint, err := cm.CreateCheckpoint(
		"feature",
		fmt.Sprintf("Feature: %s", feature.Name),
		fmt.Sprintf("Auto-checkpoint for feature %s in phase %s", feature.Name, feature.Phase),
	)

	if err != nil {
		return err
	}

	// Store checkpoint reference in feature
	feature.CheckpointID = checkpoint.ID

	return nil
}

// AutoCreatePhaseCheckpoint automatically creates checkpoint for phase
func (cm *CheckpointManager) AutoCreatePhaseCheckpoint(phase *models.Phase) error {
	cfgMgr := config.NewManager(cm.projectRoot)
	cfg, err := cfgMgr.LoadConfig()
	if err != nil {
		return err
	}

	// Check if auto-checkpoint is enabled
	if !cfg.Checkpoint.AutoPhase {
		return nil
	}

	_, err = cm.CreateCheckpoint(
		"phase",
		fmt.Sprintf("Phase: %s", phase.Name),
		fmt.Sprintf("Auto-checkpoint for phase %s", phase.Name),
	)

	return err
}

// ListCheckpoints lists all checkpoints
func (cm *CheckpointManager) ListCheckpoints() ([]*Checkpoint, error) {
	metadataDir := filepath.Join(cm.checkpointsDir, "metadata")

	if _, err := os.Stat(metadataDir); os.IsNotExist(err) {
		return []*Checkpoint{}, nil
	}

	entries, err := os.ReadDir(metadataDir)
	if err != nil {
		return nil, err
	}

	var checkpoints []*Checkpoint

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		metadataPath := filepath.Join(metadataDir, entry.Name())
		checkpoint, err := cm.loadCheckpointMetadata(metadataPath)
		if err != nil {
			continue
		}

		checkpoints = append(checkpoints, checkpoint)
	}

	return checkpoints, nil
}

// RestoreCheckpoint restores a checkpoint
func (cm *CheckpointManager) RestoreCheckpoint(checkpointID string) error {
	checkpoint, err := cm.loadCheckpointByID(checkpointID)
	if err != nil {
		return fmt.Errorf("failed to load checkpoint: %w", err)
	}

	// Extract archive
	if err := cm.extractArchive(checkpoint.FilePath); err != nil {
		return fmt.Errorf("failed to extract archive: %w", err)
	}

	// Restore state
	cfgMgr := config.NewManager(cm.projectRoot)
	if err := cfgMgr.SaveState(checkpoint.State); err != nil {
		return fmt.Errorf("failed to restore state: %w", err)
	}

	color.Green("✅ Checkpoint restored: %s\n", checkpoint.Name)

	return nil
}

func (cm *CheckpointManager) createArchive(checkpoint *Checkpoint) (string, error) {
	archivePath := filepath.Join(cm.checkpointsDir, "archives", fmt.Sprintf("%s.tar.gz", checkpoint.ID))

	file, err := os.Create(archivePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// Archive doplan directory
	doplanDir := filepath.Join(cm.projectRoot, "doplan")
	if _, err := os.Stat(doplanDir); err == nil {
		if err := cm.addDirectoryToArchive(tarWriter, doplanDir, "doplan"); err != nil {
			return "", err
		}
	}

	// Archive config
	configDir := filepath.Join(cm.projectRoot, ".cursor", "config")
	if _, err := os.Stat(configDir); err == nil {
		if err := cm.addDirectoryToArchive(tarWriter, configDir, ".cursor/config"); err != nil {
			return "", err
		}
	}

	return archivePath, nil
}

func (cm *CheckpointManager) addDirectoryToArchive(tw *tar.Writer, dirPath, archivePath string) error {
	return filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dirPath, filePath)
		if err != nil {
			return err
		}

		header.Name = filepath.Join(archivePath, relPath)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	})
}

func (cm *CheckpointManager) extractArchive(archivePath string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(cm.projectRoot, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()

			if err := os.Chmod(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (cm *CheckpointManager) saveCheckpointMetadata(checkpoint *Checkpoint) error {
	metadataDir := filepath.Join(cm.checkpointsDir, "metadata")

	metadataPath := filepath.Join(metadataDir, fmt.Sprintf("%s.json", checkpoint.ID))

	data, err := json.MarshalIndent(checkpoint, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(metadataPath, data, 0644)
}

func (cm *CheckpointManager) loadCheckpointMetadata(path string) (*Checkpoint, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var checkpoint Checkpoint
	if err := json.Unmarshal(data, &checkpoint); err != nil {
		return nil, err
	}

	return &checkpoint, nil
}

func (cm *CheckpointManager) loadCheckpointByID(checkpointID string) (*Checkpoint, error) {
	metadataPath := filepath.Join(cm.checkpointsDir, "metadata", fmt.Sprintf("%s.json", checkpointID))
	return cm.loadCheckpointMetadata(metadataPath)
}

func generateCheckpointID() string {
	return fmt.Sprintf("cp-%d", time.Now().Unix())
}

// IsCommandAvailable checks if a command is available
func IsCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
