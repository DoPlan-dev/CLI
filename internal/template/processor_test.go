package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DoPlan-dev/CLI/pkg/models"
	"github.com/DoPlan-dev/CLI/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProcessor(t *testing.T) {
	projectRoot := helpers.CreateTempProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	processor := NewProcessor(templatesDir)

	assert.NotNil(t, processor)
	assert.Equal(t, templatesDir, processor.templatesDir)
}

func TestProcessor_LoadTemplate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := "Hello {{.Feature.Name}}"
	err = os.WriteFile(filepath.Join(templatesDir, "test.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)
	err = processor.LoadTemplate("test.tmpl")
	require.NoError(t, err)
}

func TestProcessor_ProcessTemplate(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := "Feature: {{.Feature.Name}}"
	err = os.WriteFile(filepath.Join(templatesDir, "test.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)
	data := TemplateData{
		Feature: &models.Feature{
			Name: "Test Feature",
		},
	}

	result, err := processor.ProcessTemplate("test.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Test Feature")
}

func TestProcessor_ProcessTemplate_formatList(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	// Test formatList with items
	content := `{{formatList .Project.Items}}`
	err = os.WriteFile(filepath.Join(templatesDir, "list.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)
	data := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{"Item 1", "Item 2", "Item 3"},
		},
	}

	result, err := processor.ProcessTemplate("list.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Item 1")
	assert.Contains(t, result, "Item 2")
	assert.Contains(t, result, "Item 3")
	assert.Contains(t, result, "- ")

	// Test formatList with empty list
	data2 := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{},
		},
	}
	result2, err := processor.ProcessTemplate("list.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "None")
}

func TestProcessor_ProcessTemplate_formatChecklist(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{formatChecklist .Project.Items}}`
	err = os.WriteFile(filepath.Join(templatesDir, "checklist.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)
	data := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{"Task 1", "Task 2"},
		},
	}

	result, err := processor.ProcessTemplate("checklist.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "[ ]")
	assert.Contains(t, result, "Task 1")
	assert.Contains(t, result, "Task 2")

	// Test with empty list
	data2 := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{},
		},
	}
	result2, err := processor.ProcessTemplate("checklist.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "None")
}

func TestProcessor_ProcessTemplate_progressBar(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{progressBar .Project.Progress 20}}`
	err = os.WriteFile(filepath.Join(templatesDir, "progress.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)

	// Test with 50% progress
	data := TemplateData{
		Project: map[string]interface{}{
			"Progress": 50,
		},
	}
	result, err := processor.ProcessTemplate("progress.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "50%")
	assert.Contains(t, result, "[")

	// Test with 0% progress
	data2 := TemplateData{
		Project: map[string]interface{}{
			"Progress": 0,
		},
	}
	result2, err := processor.ProcessTemplate("progress.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "0%")

	// Test with 100% progress
	data3 := TemplateData{
		Project: map[string]interface{}{
			"Progress": 100,
		},
	}
	result3, err := processor.ProcessTemplate("progress.tmpl", data3)
	require.NoError(t, err)
	assert.Contains(t, result3, "100%")
}

func TestProcessor_ProcessTemplate_default(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{default .Project.Value "default"}}`
	err = os.WriteFile(filepath.Join(templatesDir, "default.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)

	// Test with empty value
	data := TemplateData{
		Project: map[string]interface{}{
			"Value": "",
		},
	}
	result, err := processor.ProcessTemplate("default.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "default")

	// Test with non-empty value
	data2 := TemplateData{
		Project: map[string]interface{}{
			"Value": "actual",
		},
	}
	result2, err := processor.ProcessTemplate("default.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "actual")
	assert.NotContains(t, result2, "default")
}

func TestProcessor_ProcessTemplate_hasBranch(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{if hasBranch .Feature}}Has Branch{{else}}No Branch{{end}}`
	err = os.WriteFile(filepath.Join(templatesDir, "branch.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)

	// Test with branch
	data := TemplateData{
		Feature: &models.Feature{
			Branch: "feature/test",
		},
	}
	result, err := processor.ProcessTemplate("branch.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Has Branch")
	assert.NotContains(t, result, "No Branch")

	// Test without branch
	data2 := TemplateData{
		Feature: &models.Feature{
			Branch: "",
		},
	}
	result2, err := processor.ProcessTemplate("branch.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "No Branch")

	// Test with nil feature
	data3 := TemplateData{
		Feature: nil,
	}
	result3, err := processor.ProcessTemplate("branch.tmpl", data3)
	require.NoError(t, err)
	assert.Contains(t, result3, "No Branch")
}

func TestProcessor_ProcessTemplate_hasPR(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{if hasPR .Feature}}Has PR{{else}}No PR{{end}}`
	err = os.WriteFile(filepath.Join(templatesDir, "pr.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)

	// Test with PR
	data := TemplateData{
		Feature: &models.Feature{
			PR: &models.PullRequest{
				URL: "https://github.com/test/pr/1",
			},
		},
	}
	result, err := processor.ProcessTemplate("pr.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Has PR")

	// Test without PR
	data2 := TemplateData{
		Feature: &models.Feature{
			PR: nil,
		},
	}
	result2, err := processor.ProcessTemplate("pr.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "No PR")

	// Test with nil feature
	data3 := TemplateData{
		Feature: nil,
	}
	result3, err := processor.ProcessTemplate("pr.tmpl", data3)
	require.NoError(t, err)
	assert.Contains(t, result3, "No PR")
}

func TestProcessor_ProcessTemplate_isEmpty(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{if isEmpty .Project.Value}}Empty{{else}}Not Empty{{end}}`
	err = os.WriteFile(filepath.Join(templatesDir, "empty.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)

	// Test with empty string
	data := TemplateData{
		Project: map[string]interface{}{
			"Value": "",
		},
	}
	result, err := processor.ProcessTemplate("empty.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Empty")

	// Test with non-empty string
	data2 := TemplateData{
		Project: map[string]interface{}{
			"Value": "not empty",
		},
	}
	result2, err := processor.ProcessTemplate("empty.tmpl", data2)
	require.NoError(t, err)
	assert.Contains(t, result2, "Not Empty")
}

func TestProcessor_ProcessTemplate_join(t *testing.T) {
	projectRoot := helpers.SetupTestProject(t)
	templatesDir := filepath.Join(projectRoot, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	content := `{{join .Project.Items ", "}}`
	err = os.WriteFile(filepath.Join(templatesDir, "join.tmpl"), []byte(content), 0644)
	require.NoError(t, err)

	processor := NewProcessor(templatesDir)
	data := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{"Item 1", "Item 2", "Item 3"},
		},
	}

	result, err := processor.ProcessTemplate("join.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result, "Item 1, Item 2, Item 3")

	// Test with different separator
	content2 := `{{join .Project.Items " | "}}`
	err = os.WriteFile(filepath.Join(templatesDir, "join2.tmpl"), []byte(content2), 0644)
	require.NoError(t, err)

	result2, err := processor.ProcessTemplate("join2.tmpl", data)
	require.NoError(t, err)
	assert.Contains(t, result2, "Item 1 | Item 2 | Item 3")

	// Test with single item
	data3 := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{"Single"},
		},
	}
	result3, err := processor.ProcessTemplate("join.tmpl", data3)
	require.NoError(t, err)
	assert.Contains(t, result3, "Single")

	// Test with empty list
	data4 := TemplateData{
		Project: map[string]interface{}{
			"Items": []string{},
		},
	}
	result4, err := processor.ProcessTemplate("join.tmpl", data4)
	require.NoError(t, err)
	assert.Empty(t, result4)
}

func TestRepeatString(t *testing.T) {
	// Test with positive count
	result := repeatString("a", 3)
	assert.Equal(t, "aaa", result)

	// Test with count 0
	result2 := repeatString("a", 0)
	assert.Equal(t, "", result2)

	// Test with negative count
	result3 := repeatString("a", -1)
	assert.Equal(t, "", result3)

	// Test with count 1
	result4 := repeatString("test", 1)
	assert.Equal(t, "test", result4)

	// Test with empty string
	result5 := repeatString("", 5)
	assert.Equal(t, "", result5)
}
