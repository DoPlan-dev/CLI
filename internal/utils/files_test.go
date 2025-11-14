package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteJSON(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	outputPath := filepath.Join(projectRoot, "test", "data.json")

	data := map[string]interface{}{
		"key":    "value",
		"number": 42,
	}

	err := WriteJSON(outputPath, data)
	require.NoError(t, err)

	// Verify file exists
	assert.FileExists(t, outputPath)

	// Verify content
	fileData, err := os.ReadFile(outputPath)
	require.NoError(t, err)

	var loaded map[string]interface{}
	err = json.Unmarshal(fileData, &loaded)
	require.NoError(t, err)
	assert.Equal(t, "value", loaded["key"])
}

func TestEnsureDir(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	dirPath := filepath.Join(projectRoot, "nested", "deep", "directory")

	err := EnsureDir(dirPath)
	require.NoError(t, err)

	// Verify directory exists
	assert.DirExists(t, dirPath)
}

func TestEnsureDir_Existing(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	dirPath := filepath.Join(projectRoot, "existing")

	err := os.MkdirAll(dirPath, 0755)
	require.NoError(t, err)

	err = EnsureDir(dirPath)
	require.NoError(t, err)
	assert.DirExists(t, dirPath)
}

func TestBatchWriteJSON_Success(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	writes := []struct {
		Path string
		Data interface{}
	}{
		{
			Path: filepath.Join(projectRoot, "dir1", "file1.json"),
			Data: map[string]interface{}{"key1": "value1"},
		},
		{
			Path: filepath.Join(projectRoot, "dir1", "file2.json"),
			Data: map[string]interface{}{"key2": "value2"},
		},
		{
			Path: filepath.Join(projectRoot, "dir2", "file3.json"),
			Data: map[string]interface{}{"key3": "value3"},
		},
	}

	err := BatchWriteJSON(writes)
	require.NoError(t, err)

	// Verify all files were created
	for _, write := range writes {
		assert.FileExists(t, write.Path)

		// Verify content
		fileData, err := os.ReadFile(write.Path)
		require.NoError(t, err)

		var loaded map[string]interface{}
		err = json.Unmarshal(fileData, &loaded)
		require.NoError(t, err)
		assert.Equal(t, write.Data, loaded)
	}
}

func TestBatchWriteJSON_SameDirectory(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	dir := filepath.Join(projectRoot, "same-dir")

	writes := []struct {
		Path string
		Data interface{}
	}{
		{
			Path: filepath.Join(dir, "file1.json"),
			Data: map[string]int{"count": 1},
		},
		{
			Path: filepath.Join(dir, "file2.json"),
			Data: map[string]int{"count": 2},
		},
		{
			Path: filepath.Join(dir, "file3.json"),
			Data: map[string]int{"count": 3},
		},
	}

	err := BatchWriteJSON(writes)
	require.NoError(t, err)

	// Verify all files exist
	for _, write := range writes {
		assert.FileExists(t, write.Path)
	}
}

func TestBatchWriteJSON_EmptyList(t *testing.T) {
	writes := []struct {
		Path string
		Data interface{}
	}{}

	err := BatchWriteJSON(writes)
	// Should succeed with empty list
	assert.NoError(t, err)
}

func TestBatchWriteJSON_MultipleDirectories(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	writes := []struct {
		Path string
		Data interface{}
	}{
		{
			Path: filepath.Join(projectRoot, "a", "b", "c", "file1.json"),
			Data: map[string]string{"path": "a/b/c"},
		},
		{
			Path: filepath.Join(projectRoot, "x", "y", "file2.json"),
			Data: map[string]string{"path": "x/y"},
		},
		{
			Path: filepath.Join(projectRoot, "root", "file3.json"),
			Data: map[string]string{"path": "root"},
		},
	}

	err := BatchWriteJSON(writes)
	require.NoError(t, err)

	// Verify all directories were created
	for _, write := range writes {
		assert.FileExists(t, write.Path)
		assert.DirExists(t, filepath.Dir(write.Path))
	}
}

func TestBatchWriteJSON_ComplexData(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)

	writes := []struct {
		Path string
		Data interface{}
	}{
		{
			Path: filepath.Join(projectRoot, "complex1.json"),
			Data: map[string]interface{}{
				"string": "value",
				"number": 42,
				"bool":   true,
				"array":  []string{"a", "b", "c"},
				"nested": map[string]interface{}{
					"key": "value",
				},
			},
		},
		{
			Path: filepath.Join(projectRoot, "complex2.json"),
			Data: []map[string]interface{}{
				{"id": 1, "name": "first"},
				{"id": 2, "name": "second"},
			},
		},
	}

	err := BatchWriteJSON(writes)
	require.NoError(t, err)

	// Verify files and content
	for _, write := range writes {
		assert.FileExists(t, write.Path)

		fileData, err := os.ReadFile(write.Path)
		require.NoError(t, err)

		var loaded interface{}
		err = json.Unmarshal(fileData, &loaded)
		require.NoError(t, err)

		// Verify structure matches
		expectedJSON, _ := json.Marshal(write.Data)
		actualJSON, _ := json.Marshal(loaded)
		assert.Equal(t, expectedJSON, actualJSON)
	}
}
