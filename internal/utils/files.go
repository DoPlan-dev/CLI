package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// WriteJSON writes data to a JSON file
func WriteJSON(path string, data interface{}) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// EnsureDir creates directory if it doesn't exist
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// BatchWriteJSON writes multiple JSON files in parallel, grouped by directory
func BatchWriteJSON(writes []struct {
	Path string
	Data interface{}
}) error {
	// Group writes by directory
	dirGroups := make(map[string][]struct {
		Path string
		Data interface{}
	})

	for _, write := range writes {
		dir := filepath.Dir(write.Path)
		dirGroups[dir] = append(dirGroups[dir], write)
	}

	// Ensure all directories exist
	for dir := range dirGroups {
		if err := EnsureDir(dir); err != nil {
			return err
		}
	}

	// Write files in parallel
	var wg sync.WaitGroup
	errChan := make(chan error, len(writes))

	for _, write := range writes {
		wg.Add(1)
		go func(w struct {
			Path string
			Data interface{}
		}) {
			defer wg.Done()
			if err := WriteJSON(w.Path, w.Data); err != nil {
				errChan <- err
			}
		}(write)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
