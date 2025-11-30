package cli

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/internal/version"
)

// BackupType represents the type of backup
type BackupType string

const (
	BackupTypeProject     BackupType = "project"
	BackupTypePlan        BackupType = "plan"
	BackupTypeProjectPlan BackupType = "project-plan"
	BackupTypeFull        BackupType = "full"
)

// BackupManifest contains metadata about a backup
type BackupManifest struct {
	Timestamp      string     `json:"timestamp"`
	Description    string     `json:"description,omitempty"`
	BackupType     BackupType `json:"backup_type"`
	IncludedPaths  []string   `json:"included_paths"`
	ExcludedPaths  []string   `json:"excluded_paths"`
	FileCount      int        `json:"file_count"`
	CompressedSize int64      `json:"compressed_size"`
	DoplanVersion  string     `json:"doplan_version"`
	Purpose        string     `json:"purpose"`
	Checksum       string     `json:"checksum"`
	ProjectName    string     `json:"project_name,omitempty"`
}

// BackupFileInfo represents information about a backup file
type BackupFileInfo struct {
	Path        string
	Filename    string
	BackupType  BackupType
	Timestamp   time.Time
	Size        int64
	Description string
	Version     string
}

// CreateBackup creates a backup based on the specified type
func CreateBackup(projectPath string, backupType BackupType, description string) (string, error) {
	// Ensure backup directory exists
	backupDir := filepath.Join(projectPath, ".do", "backup")
	if err := utils.CreateDirectory(backupDir); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Determine what to backup
	includedPaths, excludedPaths, purpose := getBackupPaths(projectPath, backupType)

	// Collect files to backup
	var filesToBackup []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if path should be excluded
		relPath, _ := filepath.Rel(projectPath, path)
		if shouldExclude(relPath, excludedPaths) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if path should be included
		if !shouldInclude(relPath, includedPaths) {
			return nil
		}

		// Skip directories (we'll add files)
		if info.IsDir() {
			return nil
		}

		filesToBackup = append(filesToBackup, path)
		return nil
	}

	// Walk project directory
	if err := filepath.Walk(projectPath, walkFunc); err != nil {
		return "", fmt.Errorf("failed to walk project directory: %w", err)
	}

	// Get project name
	projectName := filepath.Base(projectPath)
	if projectName == "." || projectName == "" {
		projectName = "project"
	}

	// Create initial manifest (checksum will be added after archive creation)
	manifest := BackupManifest{
		Timestamp:     time.Now().Format(time.RFC3339),
		Description:   description,
		BackupType:    backupType,
		IncludedPaths: includedPaths,
		ExcludedPaths: excludedPaths,
		FileCount:     len(filesToBackup),
		DoplanVersion: version.GetVersion(),
		Purpose:       purpose,
		ProjectName:   projectName,
	}

	// Create backup with manifest
	return createBackupWithManifest(projectPath, backupType, description, filesToBackup, manifest)
}

// createBackupWithManifest creates backup with manifest included
func createBackupWithManifest(projectPath string, backupType BackupType, description string, files []string, manifest BackupManifest) (string, error) {
	backupDir := filepath.Join(projectPath, ".do", "backup")
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("backup-%s-%s.doplan", backupType, timestamp)
	backupPath := filepath.Join(backupDir, filename)

	// Create temporary file first for checksum calculation
	tempPath := backupPath + ".tmp"
	file, err := os.Create(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to create backup file: %w", err)
	}

	// Create hash writer to calculate checksum while writing
	hash := sha256.New()
	multiWriter := io.MultiWriter(file, hash)

	gzipWriter := gzip.NewWriter(multiWriter)
	tarWriter := tar.NewWriter(gzipWriter)

	// Add manifest as first file
	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		file.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to marshal manifest: %w", err)
	}

	manifestHeader := &tar.Header{
		Name:    "BACKUP_MANIFEST.json",
		Size:    int64(len(manifestData)),
		Mode:    0644,
		ModTime: time.Now(),
	}

	if err := tarWriter.WriteHeader(manifestHeader); err != nil {
		tarWriter.Close()
		gzipWriter.Close()
		file.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to write manifest header: %w", err)
	}

	if _, err := tarWriter.Write(manifestData); err != nil {
		tarWriter.Close()
		gzipWriter.Close()
		file.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to write manifest: %w", err)
	}

	// Add all files
	for _, filePath := range files {
		relPath, _ := filepath.Rel(projectPath, filePath)
		if err := addFileToArchive(tarWriter, filePath, relPath); err != nil {
			tarWriter.Close()
			gzipWriter.Close()
			file.Close()
			os.Remove(tempPath)
			return "", fmt.Errorf("failed to add file: %w", err)
		}
	}

	// Close writers
	if err := tarWriter.Close(); err != nil {
		gzipWriter.Close()
		file.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to close tar writer: %w", err)
	}

	if err := gzipWriter.Close(); err != nil {
		file.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to close gzip writer: %w", err)
	}

	// Get final size and checksum
	fileInfo, _ := file.Stat()
	file.Close()

	checksum := fmt.Sprintf("%x", hash.Sum(nil))

	// Update manifest with checksum and final size
	manifest.CompressedSize = fileInfo.Size()
	manifest.Checksum = checksum

	// Rewrite archive with updated manifest
	// Read temp file
	tempFile, err := os.Open(tempPath)
	if err != nil {
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to reopen temp file: %w", err)
	}

	// Create final file
	finalFile, err := os.Create(backupPath)
	if err != nil {
		tempFile.Close()
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to create final backup file: %w", err)
	}

	// Copy data and update manifest
	gzipReader, _ := gzip.NewReader(tempFile)
	tarReader := tar.NewReader(gzipReader)

	gzipWriter2 := gzip.NewWriter(finalFile)
	tarWriter2 := tar.NewWriter(gzipWriter2)

	// Add updated manifest first
	updatedManifestData, _ := json.MarshalIndent(manifest, "", "  ")
	tarWriter2.WriteHeader(&tar.Header{
		Name:    "BACKUP_MANIFEST.json",
		Size:    int64(len(updatedManifestData)),
		Mode:    0644,
		ModTime: time.Now(),
	})
	tarWriter2.Write(updatedManifestData)

	// Copy all other files
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			tempFile.Close()
			finalFile.Close()
			os.Remove(tempPath)
			os.Remove(backupPath)
			return "", fmt.Errorf("failed to read archive: %w", err)
		}

		if header.Name == "BACKUP_MANIFEST.json" {
			continue // Skip old manifest
		}

		if err := tarWriter2.WriteHeader(header); err != nil {
			tempFile.Close()
			finalFile.Close()
			os.Remove(tempPath)
			os.Remove(backupPath)
			return "", fmt.Errorf("failed to write header: %w", err)
		}

		if _, err := io.Copy(tarWriter2, tarReader); err != nil {
			tempFile.Close()
			finalFile.Close()
			os.Remove(tempPath)
			os.Remove(backupPath)
			return "", fmt.Errorf("failed to copy file: %w", err)
		}
	}

	tarWriter2.Close()
	gzipWriter2.Close()
	tempFile.Close()
	finalFile.Close()

	// Remove temp file
	os.Remove(tempPath)

	return backupPath, nil
}

// addFileToArchive adds a file to the tar archive
func addFileToArchive(tarWriter *tar.Writer, filePath, relPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	header.Name = relPath

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	return err
}

// getBackupPaths determines what paths to include/exclude based on backup type
func getBackupPaths(projectPath string, backupType BackupType) (included []string, excluded []string, purpose string) {
	baseExcludes := []string{".git", "node_modules", ".do/backup"}

	switch backupType {
	case BackupTypeProject:
		included = []string{"src", "docs", "package.json", "go.mod", "*.config.*", "README.md"}
		excluded = append(baseExcludes, ".do", ".cursor")
		purpose = "Project backup for DoPlan system updates"

	case BackupTypePlan:
		included = []string{".do/plan"}
		excluded = baseExcludes
		purpose = "Preserve planning structure separately"

	case BackupTypeProjectPlan:
		included = []string{"src", "docs", ".do/plan", "package.json", "go.mod", "*.config.*", "README.md"}
		excluded = append(baseExcludes, ".do/core", ".do/system", ".cursor")
		purpose = "Project work + planning, ready for DoPlan update"

	case BackupTypeFull:
		included = []string{} // Include everything
		excluded = baseExcludes
		purpose = "Complete disaster recovery backup"

	default:
		included = []string{}
		excluded = baseExcludes
		purpose = "Unknown backup type"
	}

	return included, excluded, purpose
}

// shouldExclude checks if a path should be excluded
func shouldExclude(relPath string, excludedPaths []string) bool {
	normalized := filepath.ToSlash(relPath)
	for _, exclude := range excludedPaths {
		excludeNormalized := filepath.ToSlash(exclude)
		if strings.HasPrefix(normalized, excludeNormalized) || strings.HasPrefix(normalized+"/", excludeNormalized+"/") {
			return true
		}
	}
	return false
}

// shouldInclude checks if a path should be included
func shouldInclude(relPath string, includedPaths []string) bool {
	// If no includes specified, include everything (after exclusions)
	if len(includedPaths) == 0 {
		return true
	}

	normalized := filepath.ToSlash(relPath)
	for _, include := range includedPaths {
		includeNormalized := filepath.ToSlash(include)
		if strings.HasPrefix(normalized, includeNormalized) || strings.HasPrefix(normalized+"/", includeNormalized+"/") {
			return true
		}
		// Support glob patterns
		matched, _ := filepath.Match(include, filepath.Base(relPath))
		if matched {
			return true
		}
	}
	return false
}

// ListBackups lists all available backup files
func ListBackups(projectPath string) ([]BackupFileInfo, error) {
	backupDir := filepath.Join(projectPath, ".do", "backup")
	if !utils.PathExists(backupDir) {
		return []BackupFileInfo{}, nil
	}

	var backups []BackupFileInfo

	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".doplan") {
			continue
		}

		backupPath := filepath.Join(backupDir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}

		// Try to extract manifest for detailed info
		manifest, err := ExtractManifest(backupPath)
		if err != nil {
			// If we can't read manifest, create basic info
			backupFileInfo := BackupFileInfo{
				Path:      backupPath,
				Filename:  entry.Name(),
				Size:      info.Size(),
				Timestamp: info.ModTime(),
			}
			backups = append(backups, backupFileInfo)
			continue
		}

		timestamp, _ := time.Parse(time.RFC3339, manifest.Timestamp)

		backupFileInfo := BackupFileInfo{
			Path:        backupPath,
			Filename:    entry.Name(),
			BackupType:  manifest.BackupType,
			Timestamp:   timestamp,
			Size:        manifest.CompressedSize,
			Description: manifest.Description,
			Version:     manifest.DoplanVersion,
		}

		backups = append(backups, backupFileInfo)
	}

	// Sort by timestamp (newest first)
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Timestamp.After(backups[j].Timestamp)
	})

	return backups, nil
}

// ExtractManifest extracts and parses the backup manifest
func ExtractManifest(backupPath string) (*BackupManifest, error) {
	file, err := os.Open(backupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open backup file: %w", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read tar header: %w", err)
		}

		if header.Name == "BACKUP_MANIFEST.json" {
			data, err := io.ReadAll(tarReader)
			if err != nil {
				return nil, fmt.Errorf("failed to read manifest: %w", err)
			}

			var manifest BackupManifest
			if err := json.Unmarshal(data, &manifest); err != nil {
				return nil, fmt.Errorf("failed to parse manifest: %w", err)
			}

			return &manifest, nil
		}
	}

	return nil, fmt.Errorf("manifest not found in backup")
}

// VerifyBackup verifies backup integrity
func VerifyBackup(backupPath string) error {
	// Check if file exists
	if !utils.PathExists(backupPath) {
		return fmt.Errorf("backup file not found: %s", backupPath)
	}

	// Extract manifest
	manifest, err := ExtractManifest(backupPath)
	if err != nil {
		return fmt.Errorf("failed to extract manifest: %w", err)
	}

	// Verify checksum
	file, err := os.Open(backupPath)
	if err != nil {
		return fmt.Errorf("failed to open backup file: %w", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}

	calculatedChecksum := fmt.Sprintf("%x", hash.Sum(nil))
	if calculatedChecksum != manifest.Checksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", manifest.Checksum, calculatedChecksum)
	}

	// Verify file count and structure
	file.Seek(0, 0)
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	fileCount := 0

	for {
		_, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read archive: %w", err)
		}
		fileCount++
	}

	// Manifest is also counted, so subtract 1
	fileCount--
	if fileCount != manifest.FileCount {
		return fmt.Errorf("file count mismatch: expected %d, found %d", manifest.FileCount, fileCount)
	}

	return nil
}

// CheckVersionCompatibility checks if backup version is compatible with current version
func CheckVersionCompatibility(backupVersion, currentVersion string) (bool, string) {
	// Simple version comparison (can be enhanced)
	if backupVersion == currentVersion {
		return true, ""
	}

	if backupVersion == "dev" || currentVersion == "dev" {
		return true, "One version is 'dev' - compatibility cannot be verified"
	}

	// Parse semantic versions (simplified)
	backupParts := strings.Split(backupVersion, ".")
	currentParts := strings.Split(currentVersion, ".")

	if len(backupParts) >= 2 && len(currentParts) >= 2 {
		// Major version compatibility check
		if backupParts[0] != currentParts[0] {
			return false, fmt.Sprintf("Major version mismatch: backup is v%s, current is v%s. Major version changes may require migration.", backupParts[0], currentParts[0])
		}
	}

	// Minor/patch differences are generally compatible
	if len(backupParts) >= 2 && len(currentParts) >= 2 {
		backupMinor := backupParts[1]
		currentMinor := currentParts[1]
		if backupMinor != currentMinor {
			return true, fmt.Sprintf("Version difference: backup is v%s, current is v%s. Minor version differences are generally compatible.", backupVersion, currentVersion)
		}
	}

	return true, ""
}
