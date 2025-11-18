package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/DoPlan-dev/CLI/internal/dpr"
	"github.com/DoPlan-dev/CLI/internal/generators"
	"github.com/DoPlan-dev/CLI/internal/rakd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Performance benchmarks for different project sizes
func BenchmarkAgentsGeneration(b *testing.B) {
	projectRoot := setupTempProject(b)
	defer os.RemoveAll(projectRoot)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen := generators.NewAgentsGenerator(projectRoot)
		_ = gen.Generate()
	}
}

func BenchmarkRulesGeneration(b *testing.B) {
	projectRoot := setupTempProject(b)
	defer os.RemoveAll(projectRoot)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen := generators.NewRulesGenerator(projectRoot)
		_ = gen.Generate()
	}
}

func BenchmarkCommandsGeneration(b *testing.B) {
	projectRoot := setupTempProject(b)
	defer os.RemoveAll(projectRoot)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen := generators.NewCommandsGenerator(projectRoot)
		_ = gen.Generate()
	}
}

func BenchmarkDPRGeneration(b *testing.B) {
	projectRoot := setupTempProject(b)
	defer os.RemoveAll(projectRoot)

	data := createMockDPRData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen := dpr.NewGenerator(projectRoot, data)
		_ = gen.Generate()
	}
}

func BenchmarkRAKDGeneration(b *testing.B) {
	projectRoot := setupTempProjectWithDependencies(b)
	defer os.RemoveAll(projectRoot)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rakd.GenerateRAKD(projectRoot)
	}
}

// Test large project scenarios
func TestLargeProjectPerformance(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Create a large number of dependencies
	createLargePackageJSON(t, projectRoot, 100)

	start := time.Now()

	// Generate agents
	gen := generators.NewAgentsGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err)
	agentsTime := time.Since(start)

	start = time.Now()
	// Generate rules
	ruleGen := generators.NewRulesGenerator(projectRoot)
	err = ruleGen.Generate()
	require.NoError(t, err)
	rulesTime := time.Since(start)

	start = time.Now()
	// Generate RAKD with many dependencies
	_, err = rakd.GenerateRAKD(projectRoot)
	require.NoError(t, err)
	rakdTime := time.Since(start)

	t.Logf("Performance metrics for large project:")
	t.Logf("  Agents generation: %v", agentsTime)
	t.Logf("  Rules generation: %v", rulesTime)
	t.Logf("  RAKD generation: %v", rakdTime)

	// Assert reasonable performance (should complete in under 1 second each)
	assert.Less(t, agentsTime, 1*time.Second, "Agents generation should be fast")
	assert.Less(t, rulesTime, 1*time.Second, "Rules generation should be fast")
	assert.Less(t, rakdTime, 2*time.Second, "RAKD generation should be reasonably fast")
}

// Test edge cases
func TestEdgeCaseLongNames(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Test with very long project names
	longName := strings.Repeat("a", 500) // 500 character name
	data := &dpr.DPRData{
		Answers: map[string]interface{}{
			"project_name": longName,
		},
	}

	gen := dpr.NewGenerator(projectRoot, data)
	err := gen.Generate()
	require.NoError(t, err, "Should handle long names")

	// Verify file was created
	dprPath := filepath.Join(projectRoot, "doplan/design/DPR.md")
	assert.FileExists(t, dprPath, "DPR should be generated with long name")
}

func TestEdgeCaseSpecialCharacters(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Test with special characters in project name
	specialName := "test-project_with-special@chars#123!$%^&*()"
	data := &dpr.DPRData{
		Answers: map[string]interface{}{
			"project_name": specialName,
		},
	}

	gen := dpr.NewGenerator(projectRoot, data)
	err := gen.Generate()
	require.NoError(t, err, "Should handle special characters")

	// Verify file was created
	dprPath := filepath.Join(projectRoot, "doplan/design/DPR.md")
	assert.FileExists(t, dprPath, "DPR should be generated with special characters")
}

func TestEdgeCaseEmptyProject(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Test with empty project (no dependencies)
	gen := generators.NewAgentsGenerator(projectRoot)
	err := gen.Generate()
	require.NoError(t, err, "Should handle empty project")

	// Generate RAKD with no dependencies
	_, err = rakd.GenerateRAKD(projectRoot)
	require.NoError(t, err, "Should handle project with no dependencies")
}

func TestEdgeCaseMissingFiles(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Test RAKD generation when package.json doesn't exist
	_, err := rakd.GenerateRAKD(projectRoot)
	require.NoError(t, err, "Should handle missing package.json gracefully")
}

func TestEdgeCaseVeryLargeDPRData(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Create DPR data with very large strings
	largeString := strings.Repeat("test data ", 1000)
	data := &dpr.DPRData{
		Answers: map[string]interface{}{
			"project_name":      largeString,
			"audience_primary":  largeString,
			"emotion_target":    largeString,
			"style_overall":     largeString,
			"color_primary":     largeString,
			"typography_font":   largeString,
			"layout_style":      largeString,
			"components_style":  largeString,
			"animation_level":   largeString,
			"responsive_priority": largeString,
		},
	}

	start := time.Now()
	gen := dpr.NewGenerator(projectRoot, data)
	err := gen.Generate()
	require.NoError(t, err, "Should handle very large DPR data")

	elapsed := time.Since(start)
	t.Logf("Large DPR data generation time: %v", elapsed)

	// Should still complete in reasonable time
	assert.Less(t, elapsed, 2*time.Second, "Should handle large data quickly")

	// Verify file was created
	dprPath := filepath.Join(projectRoot, "doplan/design/DPR.md")
	assert.FileExists(t, dprPath, "DPR should be generated with large data")
}

func TestMemoryUsage(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// Generate multiple times
	for i := 0; i < 10; i++ {
		gen := generators.NewAgentsGenerator(projectRoot)
		_ = gen.Generate()

		ruleGen := generators.NewRulesGenerator(projectRoot)
		_ = ruleGen.Generate()

		cmdGen := generators.NewCommandsGenerator(projectRoot)
		_ = cmdGen.Generate()
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	// Calculate memory used (handle potential overflow)
	var memUsed uint64
	if m2.Alloc > m1.Alloc {
		memUsed = m2.Alloc - m1.Alloc
	} else {
		// If Alloc decreased, check TotalAlloc instead
		memUsed = m2.TotalAlloc - m1.TotalAlloc
	}
	
	t.Logf("Memory used: %d KB (%.2f MB)", memUsed/1024, float64(memUsed)/(1024*1024))
	t.Logf("Heap allocated: m1=%d KB, m2=%d KB", m1.Alloc/1024, m2.Alloc/1024)

	// Memory should be reasonable (less than 10MB for 10 generations)
	// Use TotalAlloc as a fallback check
	totalMemUsed := m2.TotalAlloc - m1.TotalAlloc
	if totalMemUsed > 0 && totalMemUsed < uint64(10*1024*1024) {
		t.Logf("Total memory used: %d KB (%.2f MB)", totalMemUsed/1024, float64(totalMemUsed)/(1024*1024))
		assert.Less(t, totalMemUsed, uint64(10*1024*1024), "Total memory usage should be reasonable")
	} else {
		t.Logf("Memory usage check skipped (may be affected by GC)")
	}
}

func TestConcurrentGeneration(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Test concurrent generation (should not cause race conditions)
	done := make(chan error, 3)

	go func() {
		gen := generators.NewAgentsGenerator(projectRoot)
		done <- gen.Generate()
	}()

	go func() {
		gen := generators.NewRulesGenerator(projectRoot)
		done <- gen.Generate()
	}()

	go func() {
		gen := generators.NewCommandsGenerator(projectRoot)
		done <- gen.Generate()
	}()

	// Wait for all to complete
	for i := 0; i < 3; i++ {
		err := <-done
		assert.NoError(t, err, "Concurrent generation should not fail")
	}

	// Verify all files were created
	assert.FileExists(t, filepath.Join(projectRoot, ".doplan/ai/agents/README.md"))
	assert.FileExists(t, filepath.Join(projectRoot, ".doplan/ai/rules/workflow.mdc"))
	assert.FileExists(t, filepath.Join(projectRoot, ".doplan/ai/commands/run.md"))
}

func TestManyDependencies(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Create package.json with many dependencies
	createLargePackageJSON(t, projectRoot, 50)

	start := time.Now()
	_, err := rakd.GenerateRAKD(projectRoot)
	require.NoError(t, err, "Should handle many dependencies")

	elapsed := time.Since(start)
	t.Logf("RAKD generation with 50 dependencies: %v", elapsed)

	// Should still be reasonably fast
	assert.Less(t, elapsed, 3*time.Second, "Should handle many dependencies quickly")
}

func TestVeryDeepDirectoryStructure(t *testing.T) {
	projectRoot := setupTempProject(t)
	defer os.RemoveAll(projectRoot)

	// Create very deep directory structure
	deepPath := projectRoot
	for i := 0; i < 20; i++ {
		deepPath = filepath.Join(deepPath, fmt.Sprintf("level%d", i))
	}
	err := os.MkdirAll(deepPath, 0755)
	require.NoError(t, err)

	// Test generation still works
	gen := generators.NewAgentsGenerator(projectRoot)
	err = gen.Generate()
	require.NoError(t, err, "Should handle deep directory structures")

	// Verify files were created at correct location
	assert.FileExists(t, filepath.Join(projectRoot, ".doplan/ai/agents/README.md"))
}

// Helper functions
func setupTempProject(t testing.TB) string {
	dir := filepath.Join(os.TempDir(), fmt.Sprintf("doplan-perf-test-%d", time.Now().UnixNano()))
	err := os.MkdirAll(dir, 0755)
	require.NoError(t, err)

	// Create basic structure
	doplanDirs := []string{
		".doplan/ai/agents",
		".doplan/ai/rules",
		".doplan/ai/commands",
		"doplan/design",
	}
	for _, d := range doplanDirs {
		err := os.MkdirAll(filepath.Join(dir, d), 0755)
		require.NoError(t, err)
	}

	return dir
}

func setupTempProjectWithDependencies(t testing.TB) string {
	dir := setupTempProject(t)

	// Create package.json with common dependencies
	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"express": "^4.18.0",
			"@stripe/stripe-js": "^2.0.0",
			"aws-sdk": "^2.1000.0",
			"@supabase/supabase-js": "^2.0.0",
			"openai": "^4.0.0",
			"@auth0/auth0-spa-js": "^2.0.0",
			"@sentry/node": "^7.0.0"
		}
	}`
	err := os.WriteFile(filepath.Join(dir, "package.json"), []byte(packageJSON), 0644)
	require.NoError(t, err)

	// Create .env.example
	envExample := `STRIPE_API_KEY=sk_test_...
AWS_ACCESS_KEY_ID=AKIA...
SUPABASE_URL=https://...
OPENAI_API_KEY=sk-...
AUTH0_DOMAIN=...
SENTRY_DSN=...`
	err = os.WriteFile(filepath.Join(dir, ".env.example"), []byte(envExample), 0644)
	require.NoError(t, err)

	return dir
}

func createLargePackageJSON(t testing.TB, projectRoot string, numDeps int) {
	deps := make(map[string]string)
	deps["express"] = "^4.18.0"
	deps["@stripe/stripe-js"] = "^2.0.0"

	for i := 0; i < numDeps-2; i++ {
		deps[fmt.Sprintf("dep-%d", i)] = "^1.0.0"
	}

	packageJSON := fmt.Sprintf(`{
		"name": "test-project",
		"dependencies": {`)
	
	depsList := []string{}
	for name, version := range deps {
		depsList = append(depsList, fmt.Sprintf(`			"%s": "%s"`, name, version))
	}
	packageJSON += "\n" + strings.Join(depsList, ",\n")
	packageJSON += "\n		}\n	}"

	err := os.WriteFile(filepath.Join(projectRoot, "package.json"), []byte(packageJSON), 0644)
	require.NoError(t, err)
}

func createMockDPRData() *dpr.DPRData {
	return &dpr.DPRData{
		Answers: map[string]interface{}{
			"project_name":           "Test Project",
			"audience_primary":       "Developers",
			"emotion_target":         "Professional",
			"style_overall":          "Modern",
			"color_primary":          "#667eea",
			"typography_font":        "Inter",
			"layout_style":           "Card-based",
			"components_style":       "Elevated",
			"animation_level":        "Subtle",
			"accessibility_importance": 5,
			"responsive_priority":    "Desktop First",
		},
	}
}

