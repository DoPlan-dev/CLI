# Dashboard Generator Testing Guide

This guide explains how to test the Dashboard Generator in multiple ways.

## Automated Testing

### Run All Dashboard Tests

```bash
go test -v ./internal/generator -run TestDashboardGenerator
```

### Run Specific Test

```bash
# Test basic generation
go test -v ./internal/generator -run TestDashboardGenerator_Generate

# Test content validation
go test -v ./internal/generator -run TestDashboardGenerator_Generate_IndexContent

# Test integration with full orchestration
go test -v ./internal/generator -run TestDashboardGenerator_Integration
```

### Run All Generator Tests (Including Dashboard)

```bash
go test -v ./internal/generator
```

### Run Integration Tests (Full Project Generation)

```bash
go test -v ./internal/generator -run TestFullProjectGeneration
```

## Manual Testing

### Method 1: Generate a Test Project

1. **Run the CLI wizard:**
   ```bash
   go run . 
   # or if installed:
   doplan
   ```

2. **Complete the wizard** with any project details

3. **Navigate to the generated project:**
   ```bash
   cd <project-name>
   ```

4. **Open the dashboard:**
   ```bash
   # Open index.html in your browser
   open dashboard/index.html
   # or on Linux:
   xdg-open dashboard/index.html
   # or on Windows:
   start dashboard/index.html
   ```

5. **Test all pages:**
   - `dashboard/index.html` - Main dashboard
   - `dashboard/plan.html` - Tasks viewer
   - `dashboard/meetings.html` - Meetings history
   - `dashboard/achievements.html` - Achievements display
   - `dashboard/settings.html` - Project settings

### Method 2: Test Dashboard Generator Directly

Create a simple test script:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/generator"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

func main() {
	// Create a test project request
	request := &models.ProjectRequest{
		ProjectName: "test-dashboard",
		IDE:         "Cursor",
		IDEs:        []string{"Cursor", "VSCode"},
		ProjectType: "Fullstack",
	}

	// Create a temporary directory
	tmpDir := filepath.Join(".", "test-dashboard-output")
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	// Generate dashboard
	dashboardGen := &generator.DashboardGenerator{}
	if err := dashboardGen.Generate(request, tmpDir); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Dashboard generated successfully in %s\n", tmpDir)
	fmt.Printf("Open: %s\n", filepath.Join(tmpDir, "dashboard", "index.html"))
}
```

Run it:
```bash
go run test_dashboard.go
```

### Method 3: Test with a Local HTTP Server

Since the dashboard uses JavaScript to fetch data, you may need a local server:

```bash
# Navigate to your project
cd <project-name>

# Start a simple HTTP server (Python 3)
python3 -m http.server 8000

# Or with Node.js
npx http-server -p 8000

# Or with Go
go run -m http.server 8000
```

Then open: `http://localhost:8000/dashboard/index.html`

## What to Test

### Visual Testing

1. **Layout & Styling:**
   - ✅ Dark theme (#0A0E27 background)
   - ✅ Cyan accent color (#00D9FF)
   - ✅ Sidebar navigation
   - ✅ Responsive cards
   - ✅ Tabler icons display correctly

2. **Navigation:**
   - ✅ All sidebar links work
   - ✅ Active page is highlighted
   - ✅ Navigation persists across pages

3. **Dashboard Page (index.html):**
   - ✅ Project overview card shows project name and type
   - ✅ Progress bar displays (if data available)
   - ✅ Task count displays
   - ✅ Current phase displays

4. **Plan Page (plan.html):**
   - ✅ TASKS.md content loads and renders
   - ✅ Search functionality works
   - ✅ Markdown renders correctly

5. **Meetings Page (meetings.html):**
   - ✅ Meeting history displays
   - ✅ Upcoming meetings show (if any)

6. **Achievements Page (achievements.html):**
   - ✅ XP, Level, and Streak display
   - ✅ Achievement cards render
   - ✅ Confetti animation works on click

7. **Settings Page (settings.html):**
   - ✅ Project settings table displays
   - ✅ Agents list loads
   - ✅ Memory insights display

### Functional Testing

1. **Data Loading:**
   - Test with existing `.do/system/history/active_state.json`
   - Test with existing `.do/plan/TASKS.md`
   - Test with existing `.do/system/memory_card.json`
   - Test with missing data (should show fallback messages)

2. **Error Handling:**
   - Test with missing files (should show "No data available" messages)
   - Test with invalid JSON (should handle gracefully)

3. **Browser Compatibility:**
   - Test in Chrome/Edge
   - Test in Firefox
   - Test in Safari
   - Test on mobile devices

## Testing Checklist

- [ ] All HTML pages generate correctly
- [ ] All pages have proper navigation
- [ ] CSS and JavaScript load from CDN
- [ ] Dashboard displays project information
- [ ] Progress data loads from active_state.json
- [ ] Tasks display from TASKS.md
- [ ] Achievements load from memory_card.json
- [ ] Settings page shows project configuration
- [ ] All pages are responsive
- [ ] No console errors in browser
- [ ] All links work correctly

## Debugging Tips

1. **Check browser console** for JavaScript errors
2. **Check network tab** to see if CDN resources load
3. **Verify file paths** - dashboard expects data in `../.do/` relative to dashboard folder
4. **Check CORS** - if testing locally, ensure proper file:// or http:// protocol
5. **Validate HTML** - use browser dev tools to check structure

## Common Issues

### Issue: "No progress data available"
- **Solution:** Ensure `.do/system/history/active_state.json` exists in the project root

### Issue: "No tasks file available"
- **Solution:** Ensure `.do/plan/TASKS.md` exists in the project root

### Issue: CDN resources not loading
- **Solution:** Check internet connection, or use local copies of Bootstrap/Tabler

### Issue: JavaScript errors
- **Solution:** Check browser console, ensure all required data files exist

## Performance Testing

```bash
# Test generation speed
time go test -v ./internal/generator -run TestDashboardGenerator_Generate

# Test with large project data
# (Create a project with many tasks/achievements and test dashboard load time)
```

## Continuous Integration

The dashboard generator is automatically tested in CI when you run:

```bash
go test ./...
```

All tests should pass before merging code.
