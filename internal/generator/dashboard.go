package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// DashboardGenerator generates the DoPlan Dashboard
type DashboardGenerator struct{}

// Name returns the name of the generator
func (g *DashboardGenerator) Name() string {
	return "Dashboard"
}

// Generate creates the dashboard HTML files and structure
func (g *DashboardGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	dashboardDir := filepath.Join(projectPath, "dashboard")
	dataDir := filepath.Join(dashboardDir, "data")
	imagesDir := filepath.Join(dashboardDir, "images")

	// Create directories
	if err := utils.CreateDirectory(dashboardDir); err != nil {
		return fmt.Errorf("failed to create dashboard directory: %w", err)
	}
	if err := utils.CreateDirectory(dataDir); err != nil {
		return fmt.Errorf("failed to create dashboard data directory: %w", err)
	}
	if err := utils.CreateDirectory(imagesDir); err != nil {
		return fmt.Errorf("failed to create dashboard images directory: %w", err)
	}

	// Copy artwork images
	if err := copyDashboardImages(projectPath, imagesDir); err != nil {
		return fmt.Errorf("failed to copy dashboard images: %w", err)
	}

	// Generate all HTML pages
	if err := generateDashboardIndex(dashboardDir, request); err != nil {
		return fmt.Errorf("failed to generate index.html: %w", err)
	}
	if err := generateDashboardPlan(dashboardDir, request); err != nil {
		return fmt.Errorf("failed to generate plan.html: %w", err)
	}
	if err := generateDashboardMeetings(dashboardDir, request); err != nil {
		return fmt.Errorf("failed to generate meetings.html: %w", err)
	}
	if err := generateDashboardAchievements(dashboardDir, request); err != nil {
		return fmt.Errorf("failed to generate achievements.html: %w", err)
	}
	if err := generateDashboardSettings(dashboardDir, request); err != nil {
		return fmt.Errorf("failed to generate settings.html: %w", err)
	}

	// Generate initial project.json data file
	if err := generateDashboardProjectData(dataDir, request); err != nil {
		return fmt.Errorf("failed to generate project.json: %w", err)
	}

	return nil
}

// copyDashboardImages copies artwork images from docs to dashboard/images
func copyDashboardImages(projectPath string, imagesDir string) error {
	// Try multiple possible paths for the docs directory
	possiblePaths := []string{
		filepath.Join(projectPath, "..", "..", "docs", "doplan-dashboard", "src", "Images"),
		filepath.Join(projectPath, "docs", "doplan-dashboard", "src", "Images"),
		filepath.Join(projectPath, "..", "docs", "doplan-dashboard", "src", "Images"),
	}

	var docsImagesPath string
	for _, path := range possiblePaths {
		absPath, err := filepath.Abs(path)
		if err == nil && utils.PathExists(absPath) {
			docsImagesPath = absPath
			break
		}
	}

	// If no path found, don't fail - images might not exist yet
	if docsImagesPath == "" {
		return nil
	}

	images := []string{"chatbot.png", "loading.png", "version-control.png"}

	for _, img := range images {
		srcPath := filepath.Join(docsImagesPath, img)
		if utils.PathExists(srcPath) {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return fmt.Errorf("failed to read image %s: %w", img, err)
			}
			dstPath := filepath.Join(imagesDir, img)
			if err := utils.WriteFile(dstPath, data); err != nil {
				return fmt.Errorf("failed to write image %s: %w", img, err)
			}
		}
	}

	return nil
}

// generateNavbar generates the shared horizontal navigation bar
func generateNavbar(activePage string) string {
	navItems := []struct {
		name     string
		href     string
		icon     string
		isActive bool
	}{
		{"Dashboard", "index.html", "ti-dashboard", activePage == "index"},
		{"Plan", "plan.html", "ti-list-check", activePage == "plan"},
		{"Meetings", "meetings.html", "ti-calendar", activePage == "meetings"},
		{"Achievements", "achievements.html", "ti-trophy", activePage == "achievements"},
		{"Settings", "settings.html", "ti-settings", activePage == "settings"},
	}

	navLinks := ""
	for _, item := range navItems {
		activeClass := ""
		if item.isActive {
			activeClass = "text-[#206bc4] bg-[#206bc4]/10 font-semibold"
		} else {
			activeClass = "text-gray-600 hover:text-gray-900 hover:bg-gray-100"
		}
		navLinks += fmt.Sprintf(`
              <a href="%s" class="px-3 py-2 rounded-md text-sm font-medium %s transition-all flex items-center gap-2">
                 <i class="ti %s"></i>
                 %s
              </a>`, item.href, activeClass, item.icon, item.name)
	}

	return fmt.Sprintf(`<header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm" style="position: sticky; top: 0; z-index: 50; background: white; border-bottom: 1px solid #e5e7eb; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">
      <div class="container-xl">
        <div class="d-flex align-items-center justify-content-between" style="height: 64px;">
          <div class="d-flex align-items-center gap-4">
            <a href="index.html" class="d-flex align-items-center gap-2 text-decoration-none" style="opacity: 1; transition: opacity 0.2s;">
              <span style="font-size: 1.25rem; font-weight: 800; letter-spacing: -0.025em; color: #206bc4;">DoPlan.dev</span>
            </a>
            <nav class="d-none d-md-flex align-items-center gap-1">
%s
            </nav>
          </div>
          <div class="d-flex align-items-center">
            <button class="btn btn-link d-md-none p-2 text-gray-500" id="mobileMenuToggle">
              <i class="ti ti-menu-2" style="font-size: 1.5rem;"></i>
            </button>
          </div>
        </div>
      </div>
      <div class="d-md-none border-top border-gray-200 bg-white" id="mobileMenu" style="display: none;">
        <div class="px-2 pt-2 pb-3">
          <a href="index.html" class="d-block px-3 py-2 rounded-md text-base font-medium text-gray-600" style="text-decoration: none;">Dashboard</a>
          <a href="plan.html" class="d-block px-3 py-2 rounded-md text-base font-medium text-gray-600" style="text-decoration: none;">Plan</a>
          <a href="meetings.html" class="d-block px-3 py-2 rounded-md text-base font-medium text-gray-600" style="text-decoration: none;">Meetings</a>
          <a href="achievements.html" class="d-block px-3 py-2 rounded-md text-base font-medium text-gray-600" style="text-decoration: none;">Achievements</a>
          <a href="settings.html" class="d-block px-3 py-2 rounded-md text-base font-medium text-gray-600" style="text-decoration: none;">Settings</a>
        </div>
      </div>
    </header>
    <script>
      document.getElementById('mobileMenuToggle')?.addEventListener('click', function() {
        const menu = document.getElementById('mobileMenu');
        if (menu) menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
      });
    </script>`, navLinks)
}

// generateDashboardIndex generates index.html (Dashboard page)
func generateDashboardIndex(dashboardDir string, request *models.ProjectRequest) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
	<title>Dashboard - ` + request.ProjectName + `</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css" rel="stylesheet">
	<style>
		:root {
			--tblr-primary: #206bc4;
			--tblr-info: #206bc4;
		}
		body { background: #f9fafb; color: #1f2937; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; }
		.card { background: #ffffff; border: 1px solid #e5e7eb; border-radius: 0.75rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05); }
		.card:hover { border-color: rgba(32, 107, 196, 0.5); transition: border-color 0.2s; }
		.card-header { background: #ffffff; border-bottom: 1px solid #e5e7eb; color: #1f2937; }
		.progress { background: #f3f4f6; height: 6px; border-radius: 9999px; }
		.progress-bar { background: #206bc4; border-radius: 9999px; }
		.text-primary { color: #206bc4 !important; }
		.bg-primary { background-color: #206bc4 !important; }
		.btn-primary { background-color: #206bc4; border-color: #206bc4; }
		.btn-primary:hover { background-color: #1a5c98; border-color: #1a5c98; }
		.badge { background-color: #206bc4; color: white; }
		.space-y-8 > * + * { margin-top: 2rem !important; }
		.space-y-6 > * + * { margin-top: 1.5rem !important; }
		.space-y-4 > * + * { margin-top: 1rem !important; }
	</style>
</head>
<body>
	` + generateNavbar("index") + `
	<div class="page-body" style="padding: 2rem 0;">
		<div class="container-xl">
			<div class="space-y-8 mb-4">
				<!-- Welcome Hero -->
				<div class="card position-relative overflow-hidden d-flex flex-column flex-md-row align-items-center justify-content-between mb-4" style="border-radius: 0.75rem; padding: 3rem 2rem;">
					<div class="position-relative mb-4 mb-md-0 me-md-4" style="z-index: 10; max-width: 42rem; flex: 1;">
						<h2 class="mb-3" style="font-size: 1.875rem; font-weight: 700; color: #1f2937; line-height: 1.2;">Welcome back!</h2>
						<p class="mb-4" style="color: #6b7280; font-size: 1rem; line-height: 1.6;">You have completed <span id="completedTasksCount" style="color: #206bc4; font-weight: 700;">0 tasks</span> in ` + request.ProjectName + `. Keep up the great work!</p>
						<div class="d-flex flex-wrap gap-3">
							<a href="plan.html" class="btn btn-primary px-5 py-2" style="font-weight: 700; border-radius: 0.5rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">Go to Plan</a>
							<button class="btn btn-light px-5 py-2" style="font-weight: 500; border-radius: 0.5rem; border: 1px solid #d1d5db; color: #374151;">View Reports</button>
						</div>
					</div>
					<div class="position-relative d-none d-md-block" style="z-index: 10; flex-shrink: 0; width: 256px; height: 192px;">
						<img src="images/chatbot.png" alt="AI Assistant" class="w-100 h-100" style="object-fit: contain; max-width: 100%; max-height: 100%;">
					</div>
				</div>

				<!-- Stats Grid -->
				<div class="row g-4">
					<div class="col-12 col-sm-6 col-lg-3">
						<div class="card p-4">
							<div class="d-flex align-items-center justify-content-between mb-3">
								<span style="color: #6b7280; font-size: 0.875rem; font-weight: 500;">Sprint Progress</span>
								<span class="badge" style="background: #dcfce7; color: #16a34a; font-size: 0.75rem; padding: 0.25rem 0.5rem; border-radius: 0.25rem;">+12%</span>
							</div>
							<div class="h2 mb-2" style="font-size: 1.875rem; font-weight: 700; color: #1f2937;" id="progressPercentage">0%</div>
							<div class="progress">
								<div class="progress-bar" id="progressBar" style="width: 0%;"></div>
							</div>
						</div>
					</div>
					<div class="col-12 col-sm-6 col-lg-3">
						<div class="card p-4">
							<div class="d-flex align-items-center justify-content-between mb-3">
								<span style="color: #6b7280; font-size: 0.875rem; font-weight: 500;">Pending Tasks</span>
								<span class="badge" style="background: #fef3c7; color: #d97706; font-size: 0.75rem; padding: 0.25rem 0.5rem; border-radius: 0.25rem;">-2%</span>
							</div>
							<div class="h2 mb-1" style="font-size: 1.875rem; font-weight: 700; color: #1f2937;" id="pendingTasks">-</div>
							<div style="font-size: 0.75rem; color: #6b7280;">
								<span style="display: inline-block; width: 0.5rem; height: 0.5rem; border-radius: 9999px; background: #eab308; margin-right: 0.25rem;"></span>
								Needs attention
							</div>
						</div>
					</div>
					<div class="col-12 col-sm-6 col-lg-3">
						<div class="card p-4">
							<div class="d-flex align-items-center justify-content-between mb-3">
								<span style="color: #6b7280; font-size: 0.875rem; font-weight: 500;">Current Phase</span>
							</div>
							<div class="h2 mb-1" style="font-size: 1.875rem; font-weight: 700; color: #1f2937;" id="currentPhase">-</div>
							<div style="font-size: 0.75rem; color: #6b7280;" id="phaseStatus">Getting Started</div>
						</div>
					</div>
					<div class="col-12 col-sm-6 col-lg-3">
						<div class="card p-4">
							<div class="d-flex align-items-center justify-content-between mb-2">
								<span style="color: #6b7280; font-size: 0.875rem; font-weight: 500;">Tasks Completed</span>
							</div>
							<div class="h2 mb-1" style="font-size: 1.875rem; font-weight: 700; color: #1f2937;" id="tasksCompleted">-</div>
							<div style="font-size: 0.75rem; color: #6b7280;">of <span id="totalTasks">-</span> total</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js"></script>
	<script>
		async function loadProgress() {
			try {
				const response = await fetch('../.do/system/history/active_state.json');
				if (!response.ok) throw new Error('No progress data');
				const state = await response.json();
				
				const phase = state.phase || 'idea';
				const activeTask = state.active_task || null;
				
				document.getElementById('currentPhase').textContent = phase.charAt(0).toUpperCase() + phase.slice(1);
				document.getElementById('phaseStatus').textContent = activeTask ? 'In Progress' : 'Planning';
				
				try {
					const tasksResponse = await fetch('../.do/plan/TASKS.md');
					if (tasksResponse.ok) {
						const tasksText = await tasksResponse.text();
						const taskMatches = tasksText.match(/### \d+\.\d+/g) || [];
						const completedMatches = tasksText.match(/Status.*✅ Complete/g) || [];
						
						const total = taskMatches.length;
						const completed = completedMatches.length;
						const pending = total - completed;
						const progress = total > 0 ? Math.round((completed / total) * 100) : 0;
						
						document.getElementById('totalTasks').textContent = total;
						document.getElementById('tasksCompleted').textContent = completed;
						document.getElementById('pendingTasks').textContent = pending;
						document.getElementById('progressPercentage').textContent = progress + '%';
						document.getElementById('progressBar').style.width = progress + '%';
						document.getElementById('completedTasksCount').textContent = completed + ' tasks';
					}
				} catch (e) {
					// No tasks yet
				}
			} catch (e) {
				document.getElementById('currentPhase').textContent = 'Idea';
				document.getElementById('phaseStatus').textContent = 'Getting Started';
			}
		}
		loadProgress();
	</script>
</body>
</html>`

	return utils.WriteFile(filepath.Join(dashboardDir, "index.html"), []byte(content))
}

// generateDashboardPlan generates plan.html with Kanban board layout
func generateDashboardPlan(dashboardDir string, request *models.ProjectRequest) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
	<title>Plan - ` + request.ProjectName + `</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css" rel="stylesheet">
	<script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
	<style>
		body { background: #f9fafb; color: #1f2937; }
		.card { background: #ffffff; border: 1px solid #e5e7eb; border-radius: 0.75rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05); }
		.kanban-column { background: rgba(243, 244, 246, 0.8); border: 1px solid #e5e7eb; border-radius: 0.75rem; padding: 1rem; }
		.task-card { background: white; padding: 1rem; border-radius: 0.5rem; border: 1px solid #e5e7eb; cursor: move; transition: all 0.2s; }
		.task-card:hover { box-shadow: 0 4px 6px -1px rgba(0,0,0,0.1); border-color: rgba(32, 107, 196, 0.5); }
		.task-card.active { border-left: 4px solid #206bc4; }
		.custom-scrollbar::-webkit-scrollbar { width: 6px; }
		.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
		.custom-scrollbar::-webkit-scrollbar-thumb { background: #d1d5db; border-radius: 3px; }
	</style>
</head>
<body>
	` + generateNavbar("plan") + `
	<div class="page-body" style="padding: 2rem 0; height: calc(100vh - 64px); display: flex; flex-direction: column;">
		<div class="container-xl flex-grow-1 d-flex flex-column">
			<div class="d-flex flex-column flex-md-row justify-content-between align-items-start align-items-md-center mb-4 gap-3">
				<div>
					<h2 style="font-size: 1.875rem; font-weight: 700; color: #1f2937; margin-bottom: 0.5rem;">Project Plan</h2>
					<p style="color: #6b7280; font-size: 0.875rem; margin: 0;">Manage tasks and track progress across phases.</p>
				</div>
				<div class="d-flex align-items-center gap-3 w-100 w-md-auto">
					<div class="position-relative flex-grow-1 flex-md-grow-0" style="max-width: 256px;">
						<input type="text" id="searchInput" placeholder="Search tasks..." class="form-control" style="padding-left: 2.5rem; border-radius: 0.5rem; border: 1px solid #d1d5db; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">
						<i class="ti ti-search position-absolute" style="left: 0.75rem; top: 50%; transform: translateY(-50%); color: #9ca3af; font-size: 1rem;"></i>
					</div>
					<button class="btn btn-primary px-4" style="font-weight: 700; border-radius: 0.5rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">+ New Task</button>
				</div>
			</div>
			<div class="flex-grow-1 overflow-x-auto pb-4" style="min-width: 1000px;">
				<div class="d-flex gap-4 h-100">
					<div class="flex-fill d-flex flex-column kanban-column">
						<div class="d-flex justify-content-between align-items-center mb-3">
							<h3 class="mb-0 d-flex align-items-center" style="font-weight: 600; color: #374151;">
								<span style="display: inline-block; width: 0.5rem; height: 0.5rem; border-radius: 9999px; background: #6b7280; margin-right: 0.5rem;"></span>
								To Do
							</h3>
							<span class="badge" style="background: white; border: 1px solid #e5e7eb; color: #6b7280; font-size: 0.75rem; padding: 0.125rem 0.5rem; border-radius: 9999px; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);" id="todoCount">0</span>
						</div>
						<div class="flex-grow-1 overflow-y-auto custom-scrollbar" id="todoTasks" style="max-height: calc(100vh - 300px);">
							<p class="text-muted text-center p-4">Loading tasks...</p>
						</div>
					</div>
					<div class="flex-fill d-flex flex-column kanban-column">
						<div class="d-flex justify-content-between align-items-center mb-3">
							<h3 class="mb-0 d-flex align-items-center" style="font-weight: 600; color: #206bc4;">
								<span style="display: inline-block; width: 0.5rem; height: 0.5rem; border-radius: 9999px; background: #206bc4; margin-right: 0.5rem;"></span>
								In Progress
							</h3>
							<span class="badge" style="background: rgba(32, 107, 196, 0.1); border: 1px solid rgba(32, 107, 196, 0.2); color: #206bc4; font-size: 0.75rem; padding: 0.125rem 0.5rem; border-radius: 9999px; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);" id="inProgressCount">0</span>
						</div>
						<div class="flex-grow-1 overflow-y-auto custom-scrollbar" id="inProgressTasks" style="max-height: calc(100vh - 300px);">
							<p class="text-muted text-center p-4">No tasks in progress</p>
						</div>
					</div>
					<div class="flex-fill d-flex flex-column kanban-column">
						<div class="d-flex justify-content-between align-items-center mb-3">
							<h3 class="mb-0 d-flex align-items-center" style="font-weight: 600; color: #16a34a;">
								<span style="display: inline-block; width: 0.5rem; height: 0.5rem; border-radius: 9999px; background: #16a34a; margin-right: 0.5rem;"></span>
								Done
							</h3>
							<span class="badge" style="background: #dcfce7; border: 1px solid #86efac; color: #16a34a; font-size: 0.75rem; padding: 0.125rem 0.5rem; border-radius: 9999px; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);" id="doneCount">0</span>
						</div>
						<div class="flex-grow-1 overflow-y-auto custom-scrollbar" id="doneTasks" style="max-height: calc(100vh - 300px);">
							<p class="text-muted text-center p-4">No completed tasks</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js"></script>
	<script>
		async function loadTasks() {
			try {
				const response = await fetch('../.do/plan/TASKS.md');
				if (!response.ok) throw new Error('No tasks file');
				const tasksText = await response.text();
				
				// Parse tasks into columns
				const taskRegex = /### (\d+\.\d+)\s+(.+?)\n\*\*ID\*:.*?\n\*\*Status\*: (.+?)\n/gs;
				let match;
				const tasks = { todo: [], inProgress: [], done: [] };
				
				while ((match = taskRegex.exec(tasksText)) !== null) {
					const id = match[1];
					const title = match[2].trim();
					const status = match[3].trim();
					
					const task = { id, title, status };
					if (status.includes('Complete') || status.includes('✅')) {
						tasks.done.push(task);
					} else if (status.includes('Progress') || status.includes('⏳')) {
						tasks.inProgress.push(task);
					} else {
						tasks.todo.push(task);
					}
				}
				
				function renderTask(task, status) {
					const tagClass = status === 'done' ? 'bg-success' : status === 'inProgress' ? 'bg-primary' : 'bg-info';
					const cardClass = status === 'inProgress' ? 'task-card active' : 'task-card';
					const titleStyle = status === 'done' ? 'text-decoration: line-through; color: #6b7280;' : 'color: #1f2937;';
					
					return '<div class="' + cardClass + ' mb-3">' +
						'<div class="d-flex justify-content-between align-items-start mb-2">' +
						'<span class="badge ' + tagClass + '" style="font-size: 0.75rem; padding: 0.125rem 0.5rem; border-radius: 0.25rem;">' + task.id + '</span>' +
						'</div>' +
						'<h4 style="font-size: 0.875rem; font-weight: 500; margin-bottom: 0.75rem; ' + titleStyle + '">' + task.title + '</h4>' +
						(status === 'inProgress' ? '<div class="progress mb-3" style="height: 4px;"><div class="progress-bar" style="width: 66%;"></div></div>' : '') +
						'</div>';
				}
				
				document.getElementById('todoCount').textContent = tasks.todo.length;
				document.getElementById('inProgressCount').textContent = tasks.inProgress.length;
				document.getElementById('doneCount').textContent = tasks.done.length;
				
				document.getElementById('todoTasks').innerHTML = tasks.todo.length > 0 
					? tasks.todo.map(t => renderTask(t, 'todo')).join('')
					: '<p class="text-muted text-center p-4">No tasks to do</p>';
				document.getElementById('inProgressTasks').innerHTML = tasks.inProgress.length > 0
					? tasks.inProgress.map(t => renderTask(t, 'inProgress')).join('')
					: '<p class="text-muted text-center p-4">No tasks in progress</p>';
				document.getElementById('doneTasks').innerHTML = tasks.done.length > 0
					? tasks.done.map(t => renderTask(t, 'done')).join('')
					: '<p class="text-muted text-center p-4">No completed tasks</p>';
			} catch (e) {
				document.getElementById('todoTasks').innerHTML = '<p class="text-muted text-center p-4">No tasks file available yet. Run /plan to generate your execution plan.</p>';
			}
		}
		
		document.getElementById('searchInput')?.addEventListener('input', function(e) {
			const searchTerm = e.target.value.toLowerCase();
			// Simple search - could be enhanced
		});
		
		loadTasks();
	</script>
</body>
</html>`

	return utils.WriteFile(filepath.Join(dashboardDir, "plan.html"), []byte(content))
}

// generateDashboardMeetings generates meetings.html
func generateDashboardMeetings(dashboardDir string, request *models.ProjectRequest) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
	<title>Meetings - ` + request.ProjectName + `</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css" rel="stylesheet">
	<style>
		body { background: #f9fafb; color: #1f2937; }
		.card { background: #ffffff; border: 1px solid #e5e7eb; border-radius: 0.75rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05); }
		.card:hover { border-color: #206bc4; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.1); transition: all 0.2s; }
		.date-badge { width: 64px; height: 64px; background: #f9fafb; border: 1px solid #e5e7eb; border-radius: 0.5rem; display: flex; flex-direction: column; align-items: center; justify-content: center; text-align: center; }
		.date-badge:hover { background: rgba(32, 107, 196, 0.05); border-color: rgba(32, 107, 196, 0.2); transition: all 0.2s; }
	</style>
</head>
<body>
	` + generateNavbar("meetings") + `
	<div class="page-body" style="padding: 2rem 0;">
		<div class="container-xl">
			<div class="space-y-6 mb-4">
				<div class="d-flex justify-content-between align-items-center mb-4">
					<h2 style="font-size: 1.875rem; font-weight: 700; color: #1f2937;">Meetings</h2>
					<button class="btn btn-primary px-4" style="font-weight: 700; border-radius: 0.5rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">+ Schedule Meeting</button>
				</div>
				<div class="row g-4">
					<div class="col-12 col-lg-8">
						<h3 style="font-size: 1.125rem; font-weight: 600; color: #374151; margin-bottom: 1rem;">Upcoming</h3>
						<div id="upcomingMeetings" class="space-y-4 mb-4"></div>
					</div>
					<div class="col-12 col-lg-4 space-y-4 mb-4">
						<div class="card p-4">
							<h3 style="font-weight: 600; color: #1f2937; margin-bottom: 1rem;">Quick Calendar</h3>
							<div class="d-grid" style="grid-template-columns: repeat(7, 1fr); gap: 0.5rem; text-align: center; font-size: 0.875rem; margin-bottom: 0.5rem;">
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">S</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">M</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">T</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">W</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">T</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">F</span>
								<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #9ca3af;">S</span>
							</div>
							<div class="d-grid" style="grid-template-columns: repeat(7, 1fr); gap: 0.5rem; text-align: center; font-size: 0.875rem; font-weight: 500;">
								<span style="color: #9ca3af; padding: 0.25rem;">29</span>
								<span style="color: #9ca3af; padding: 0.25rem;">30</span>
								<span style="color: #374151; padding: 0.25rem; border-radius: 0.25rem; cursor: pointer;">1</span>
								<span style="color: #374151; padding: 0.25rem; border-radius: 0.25rem; cursor: pointer;">2</span>
								<span style="background: #206bc4; color: white; padding: 0.25rem; border-radius: 0.25rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">3</span>
								<span style="color: #374151; padding: 0.25rem; border-radius: 0.25rem; cursor: pointer;">4</span>
								<span style="color: #374151; padding: 0.25rem; border-radius: 0.25rem; cursor: pointer;">5</span>
							</div>
						</div>
						<div class="card p-4 position-relative overflow-hidden">
							<div class="position-absolute" style="top: 0; right: 0; padding: 1rem; opacity: 0.1;">
								<i class="ti ti-info-circle" style="font-size: 6rem; color: #206bc4;"></i>
							</div>
							<h3 style="font-weight: 600; color: #1f2937; margin-bottom: 0.5rem; position: relative; z-index: 10;">Meeting Stats</h3>
							<div class="d-flex align-items-center justify-content-between mb-2 position-relative" style="z-index: 10;">
								<span style="color: #6b7280; font-size: 0.875rem;">Hours this week</span>
								<span style="color: #206bc4; font-weight: 700; font-size: 1.25rem;">12.5h</span>
							</div>
							<div class="progress mb-3 position-relative" style="z-index: 10;">
								<div class="progress-bar" style="background: #a855f7; width: 60%;"></div>
							</div>
							<p style="font-size: 0.75rem; color: #6b7280; position: relative; z-index: 10;">You spend 20% less time in meetings compared to last week.</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js"></script>
	<script>
		function formatDateMonth(iso) {
			return new Date(iso).toLocaleString('default', { month: 'short' });
		}
		function formatDateDay(iso) {
			return new Date(iso).getDate().toString();
		}
		
		async function loadMeetings() {
			try {
				const response = await fetch('../.do/system/meeting_session.json');
				if (!response.ok) throw new Error('No meeting data');
				const meeting = await response.json();
				
				const startTime = meeting.start_time || '';
				const speed = meeting.speed || 'standard';
				const projectType = meeting.project_type || '';
				
				if (startTime) {
					const meetingDate = new Date(startTime);
					const isUpcoming = meetingDate > new Date();
					
					if (isUpcoming) {
						document.getElementById('upcomingMeetings').innerHTML = 
							'<div class="card p-4 d-flex align-items-center">' +
							'<div class="date-badge me-4 flex-shrink-0">' +
							'<span style="font-size: 0.75rem; text-transform: uppercase; font-weight: 700; color: #6b7280;">' + formatDateMonth(startTime) + '</span>' +
							'<span style="font-size: 1.25rem; font-weight: 700; color: #1f2937;">' + formatDateDay(startTime) + '</span>' +
							'</div>' +
							'<div class="flex-grow-1">' +
							'<div class="d-flex justify-content-between align-items-start mb-2">' +
							'<h4 style="font-size: 1.125rem; font-weight: 600; color: #1f2937;">Discovery Meeting</h4>' +
							'<span style="font-size: 0.875rem; font-family: monospace; color: #6b7280; background: #f3f4f6; padding: 0.25rem 0.5rem; border-radius: 0.25rem;">' + meetingDate.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit' }) + '</span>' +
							'</div>' +
							'<div style="font-size: 0.875rem; color: #6b7280;">' +
							'<span class="badge" style="background: #206bc4; color: white; margin-right: 0.5rem;">' + speed + '</span>' +
							'<span>' + projectType + '</span>' +
							'</div>' +
							'</div>' +
							'</div>';
					}
				}
			} catch (e) {
				document.getElementById('upcomingMeetings').innerHTML = '<p class="text-muted">No meeting data available yet. Run /do meeting to start a discovery session.</p>';
			}
		}
		loadMeetings();
	</script>
</body>
</html>`

	return utils.WriteFile(filepath.Join(dashboardDir, "meetings.html"), []byte(content))
}

// generateDashboardAchievements generates achievements.html
func generateDashboardAchievements(dashboardDir string, request *models.ProjectRequest) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
	<title>Achievements - ` + request.ProjectName + `</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css" rel="stylesheet">
	<script src="https://cdn.jsdelivr.net/npm/canvas-confetti@1.9.3/dist/confetti.browser.min.js"></script>
	<style>
		body { background: #f9fafb; color: #1f2937; }
		.card { background: #ffffff; border: 1px solid #e5e7eb; border-radius: 0.75rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05); transition: all 0.3s; }
		.card:hover { transform: translateY(-4px); box-shadow: 0 10px 15px -3px rgba(0,0,0,0.1); }
		.achievement-locked { opacity: 0.6; filter: grayscale(100%); }
		.achievement-unlocked { border-color: rgba(32, 107, 196, 0.2); }
		.level-circle { width: 128px; height: 128px; position: relative; }
	</style>
</head>
<body>
	` + generateNavbar("achievements") + `
	<div class="page-body" style="padding: 2rem 0;">
		<div class="container-xl">
			<div class="space-y-8 mb-4">
				<!-- Hero Section -->
				<div class="card p-8 position-relative overflow-hidden d-flex flex-column flex-md-row align-items-center mb-4" style="border-radius: 1rem;">
					<div class="position-absolute" style="right: 0; top: 0; height: 100%; width: 33%; background: linear-gradient(to left, rgba(239, 246, 255, 1), transparent); pointer-events: none;"></div>
					<div class="position-relative d-flex flex-column flex-md-row align-items-center w-100" style="z-index: 10;">
						<div class="level-circle mb-4 mb-md-0 me-md-4" style="flex-shrink: 0;">
							<svg class="w-100 h-100" style="transform: rotate(-90deg);">
								<circle cx="64" cy="64" r="58" stroke="#f1f5f9" stroke-width="8" fill="none" />
								<circle cx="64" cy="64" r="58" stroke="#206bc4" stroke-width="8" fill="none" stroke-dasharray="364" stroke-dashoffset="90" stroke-linecap="round" id="levelCircle" />
							</svg>
							<div class="position-absolute" style="top: 50%; left: 50%; transform: translate(-50%, -50%); text-align: center;">
								<span style="font-size: 1.875rem; font-weight: 700; color: #1f2937;" id="levelDisplay">Lvl 1</span>
								<span style="display: block; font-size: 0.75rem; color: #6b7280;" id="xpDisplay">XP 0/100</span>
							</div>
						</div>
						<div class="text-center text-md-start flex-grow-1">
							<h2 style="font-size: 1.875rem; font-weight: 700; color: #1f2937; margin-bottom: 0.5rem;">Achievement Master</h2>
							<p style="color: #6b7280; max-width: 32rem; margin-bottom: 1rem;" id="achievementMessage">Keep using DoPlan to unlock achievements!</p>
							<div class="d-flex flex-wrap gap-2 justify-content-center justify-content-md-start">
								<span class="badge px-3 py-1" style="background: #fef3c7; color: #d97706; border: 1px solid #fde68a; border-radius: 9999px; font-size: 0.75rem; font-weight: 600;">
									🔥 <span id="streakBadge">0</span> Day Streak
								</span>
							</div>
						</div>
					</div>
				</div>
				<!-- Achievements Grid -->
				<div>
					<div class="d-flex align-items-center justify-content-between mb-4">
						<h3 style="font-size: 1.25rem; font-weight: 700; color: #1f2937;">Badges & Milestones</h3>
						<div class="d-flex gap-2">
							<button class="btn btn-sm px-3 py-1" style="font-size: 0.875rem; border-radius: 9999px; background: rgba(32, 107, 196, 0.1); color: #206bc4; font-weight: 500; border: none;">All</button>
							<button class="btn btn-sm px-3 py-1" style="font-size: 0.875rem; border-radius: 9999px; color: #6b7280; border: none; background: transparent;">Locked</button>
							<button class="btn btn-sm px-3 py-1" style="font-size: 0.875rem; border-radius: 9999px; color: #6b7280; border: none; background: transparent;">Unlocked</button>
						</div>
					</div>
					<div class="row g-4" id="achievementsGrid">
						<p class="text-muted">Loading achievements...</p>
					</div>
				</div>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js"></script>
	<script>
		async function loadAchievements() {
			try {
				let response = await fetch('../.do/system/memory_card.json');
				if (!response.ok) response = await fetch('../../.doplan/memory_card.json');
				if (!response.ok) throw new Error('No memory card');
				
				const memoryCard = await response.json();
				const score = memoryCard.score || 0;
				const achievements = memoryCard.achievements || [];
				const streak = memoryCard.streak || 0;
				
				const level = Math.floor(score / 100) + 1;
				const xpInLevel = score % 100;
				const xpNeeded = 100;
				const progress = (xpInLevel / xpNeeded) * 100;
				const circumference = 2 * Math.PI * 58;
				const offset = circumference - (progress / 100) * circumference;
				
				document.getElementById('levelDisplay').textContent = 'Lvl ' + level;
				document.getElementById('xpDisplay').textContent = 'XP ' + score + '/1000';
				document.getElementById('streakBadge').textContent = streak;
				document.getElementById('levelCircle').style.strokeDashoffset = offset;
				
				const achievementList = [
					{id: 'first_steps', name: 'First Steps', description: 'Complete your first project', xp: 50, rarity: 'common', icon: '🎉'},
					{id: 'getting_started', name: 'Getting Started', description: 'Reach 100 points', xp: 10, rarity: 'common', icon: '🎯'},
					{id: 'on_the_rise', name: 'On the Rise', description: 'Reach 250 points', xp: 25, rarity: 'common', icon: '📈'},
				];
				
				let html = '';
				achievementList.forEach(ach => {
					const unlocked = achievements.includes(ach.id);
					const borderClass = unlocked ? 'border-blue-200' : 'border-gray-200';
					const opacityClass = unlocked ? '' : 'achievement-locked';
					const bgClass = unlocked ? 'bg-blue-50' : 'bg-gray-100';
					const textClass = unlocked ? 'text-[#206bc4]' : '';
					
					html += '<div class="col-12 col-md-6 col-lg-3">' +
						'<div class="card p-4 d-flex flex-column align-items-center text-center ' + opacityClass + '" style="border-color: ' + (unlocked ? 'rgba(32, 107, 196, 0.2)' : '#e5e7eb') + ';" onclick="' + (unlocked ? 'confetti({particleCount: 50, spread: 70});' : '') + '">' +
						'<div style="width: 64px; height: 64px; border-radius: 9999px; margin-bottom: 1rem; display: flex; align-items: center; justify-content: center; font-size: 2rem; transition: all 0.5s; background: ' + (unlocked ? 'rgba(32, 107, 196, 0.05)' : '#f3f4f6') + '; color: ' + (unlocked ? '#206bc4' : '#6b7280') + '; transform: ' + (unlocked ? 'scale(1.1)' : 'scale(1)') + ';">' + ach.icon + '</div>' +
						'<h4 style="color: #1f2937; font-weight: 700; margin-bottom: 0.25rem;">' + ach.name + '</h4>' +
						'<p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 1rem; min-height: 2.5rem;">' + ach.description + '</p>' +
						'<div class="w-100 bg-gray-100" style="height: 8px; border-radius: 9999px; overflow: hidden; margin-top: auto;">' +
						'<div class="bg-[#206bc4]" style="height: 100%; transition: width 1s; width: ' + (unlocked ? '100' : '0') + '%; border-radius: 9999px;"></div>' +
						'</div>' +
						'<span style="font-size: 0.75rem; color: #9ca3af; margin-top: 0.5rem; display: block;">' + (unlocked ? '100' : '0') + '% Complete</span>' +
						'</div>' +
						'</div>';
				});
				document.getElementById('achievementsGrid').innerHTML = html;
			} catch (e) {
				document.getElementById('achievementsGrid').innerHTML = '<p class="text-muted">No achievement data available yet.</p>';
			}
		}
		loadAchievements();
	</script>
</body>
</html>`

	return utils.WriteFile(filepath.Join(dashboardDir, "achievements.html"), []byte(content))
}

// generateDashboardSettings generates settings.html
func generateDashboardSettings(dashboardDir string, request *models.ProjectRequest) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
	<title>Settings - ` + request.ProjectName + `</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css" rel="stylesheet">
	<style>
		body { background: #f9fafb; color: #1f2937; }
		.card { background: #ffffff; border: 1px solid #e5e7eb; border-radius: 0.75rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05); }
		.tab-button { padding: 0.75rem 1.5rem; color: #6b7280; border: none; background: transparent; border-bottom: 2px solid transparent; transition: color 0.2s; }
		.tab-button.active { color: #206bc4; border-bottom-color: #206bc4; font-weight: 500; }
		.tab-button:hover { color: #1f2937; }
		.toggle-switch { width: 44px; height: 24px; background: #e5e7eb; border-radius: 9999px; position: relative; cursor: pointer; transition: background 0.2s; }
		.toggle-switch.active { background: #206bc4; }
		.toggle-switch::after { content: ''; position: absolute; top: 2px; left: 2px; width: 20px; height: 20px; background: white; border-radius: 9999px; transition: transform 0.2s; }
		.toggle-switch.active::after { transform: translateX(20px); }
	</style>
</head>
<body>
	` + generateNavbar("settings") + `
	<div class="page-body" style="padding: 2rem 0;">
		<div class="container-xl" style="max-width: 896px;">
			<h2 style="font-size: 1.875rem; font-weight: 700; color: #1f2937; margin-bottom: 2rem;">Settings</h2>
			<div class="d-flex border-bottom border-gray-200 mb-4 overflow-x-auto">
				<button class="tab-button active" data-tab="general">General</button>
				<button class="tab-button" data-tab="agents">Agents</button>
				<button class="tab-button" data-tab="notifications">Notifications</button>
				<button class="tab-button" data-tab="integrations">Integrations</button>
			</div>
			<div class="space-y-8 mb-4">
				<section class="card p-4 mb-4">
					<h3 style="font-size: 1.125rem; font-weight: 600; color: #1f2937; margin-bottom: 1rem;">Profile Information</h3>
					<div class="d-flex flex-column flex-md-row align-items-start">
						<div style="width: 80px; height: 80px; border-radius: 9999px; border: 1px solid #e5e7eb; margin-right: 1.5rem; margin-bottom: 1rem; margin-bottom-md: 0; background: #f3f4f6; display: flex; align-items: center; justify-content: center; font-size: 2rem; color: #206bc4; font-weight: 700;">` + string([]rune(request.ProjectName)[0]) + `</div>
						<div class="flex-grow-1" style="max-width: 448px;">
							<div class="row g-3 mb-3">
								<div class="col-md-6">
									<label style="display: block; font-size: 0.75rem; color: #6b7280; margin-bottom: 0.25rem; font-weight: 500;">Project Name</label>
									<input type="text" value="` + request.ProjectName + `" class="form-control" readonly style="border-radius: 0.25rem; border: 1px solid #d1d5db;">
								</div>
								<div class="col-md-6">
									<label style="display: block; font-size: 0.75rem; color: #6b7280; margin-bottom: 0.25rem; font-weight: 500;">Project Type</label>
									<input type="text" value="` + request.ProjectType + `" class="form-control" readonly style="border-radius: 0.25rem; border: 1px solid #d1d5db;">
								</div>
							</div>
							<div class="mb-3">
								<label style="display: block; font-size: 0.75rem; color: #6b7280; margin-bottom: 0.25rem; font-weight: 500;">IDE</label>
								<input type="text" value="` + getIDEsList(request.IDEs) + `" class="form-control" readonly style="border-radius: 0.25rem; border: 1px solid #d1d5db;">
							</div>
						</div>
					</div>
				</section>
				<section class="card p-4 mb-4">
					<h3 style="font-size: 1.125rem; font-weight: 600; color: #1f2937; margin-bottom: 1rem;">AI Agents</h3>
					<div class="table-responsive">
						<table class="table table-striped">
							<thead>
								<tr>
									<th>Agent</th>
									<th>Role</th>
									<th>Status</th>
								</tr>
							</thead>
							<tbody id="agentsList">
								<tr><td colspan="3" class="text-muted">Loading agents...</td></tr>
							</tbody>
						</table>
					</div>
				</section>
				<section class="card p-4" style="border-color: #fecaca; background: #fef2f2;">
					<h3 style="color: #dc2626; font-weight: 600; margin-bottom: 0.5rem;">Danger Zone</h3>
					<p style="font-size: 0.875rem; color: #4b5563; margin-bottom: 1rem;">Once you delete a project, there is no going back. Please be certain.</p>
					<button class="btn" style="background: white; color: #dc2626; border: 1px solid #fecaca; border-radius: 0.25rem; font-size: 0.875rem; font-weight: 500; padding: 0.5rem 1rem; box-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);">Delete Project</button>
				</section>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js"></script>
	<script>
		document.querySelectorAll('.tab-button').forEach(btn => {
			btn.addEventListener('click', function() {
				document.querySelectorAll('.tab-button').forEach(b => b.classList.remove('active'));
				this.classList.add('active');
			});
		});
		
		async function loadAgents() {
			try {
				const agentFiles = ['project_orchestrator.md', 'product_manager.md', 'engineering_lead.md', 'system_architect.md', 'frontend_lead.md', 'backend_lead.md'];
				let html = '';
				for (const file of agentFiles) {
					try {
						const response = await fetch('../.cursor/agents/' + file);
						if (response.ok) {
							const content = await response.text();
							const nameMatch = content.match(/# (.+)/);
							const roleMatch = content.match(/Role.*?:\s*(.+)/);
							const name = nameMatch ? nameMatch[1] : file.replace('.md', '').replace('_', ' ');
							const role = roleMatch ? roleMatch[1] : 'Agent';
							html += '<tr><td>' + name + '</td><td>' + role + '</td><td><span class="badge bg-success">Active</span></td></tr>';
						}
					} catch (e) {}
				}
				document.getElementById('agentsList').innerHTML = html || '<tr><td colspan="3" class="text-muted">No agents data available yet.</td></tr>';
			} catch (e) {
				document.getElementById('agentsList').innerHTML = '<tr><td colspan="3" class="text-muted">Error loading agents.</td></tr>';
			}
		}
		loadAgents();
	</script>
</body>
</html>`

	return utils.WriteFile(filepath.Join(dashboardDir, "settings.html"), []byte(content))
}

// generateDashboardProjectData generates initial project.json data file
func generateDashboardProjectData(dataDir string, request *models.ProjectRequest) error {
	projectData := fmt.Sprintf(`{
  "name": "%s",
  "type": "%s",
  "created": "%s",
  "ide": "%s"
}`, request.ProjectName, request.ProjectType, time.Now().Format("2006-01-02"), getIDEsList(request.IDEs))

	return utils.WriteFile(filepath.Join(dataDir, "project.json"), []byte(projectData))
}

// Helper function to format IDE list (uses existing formatIDEList from docs.go)
func getIDEsList(ides []string) string {
	return formatIDEList(ides)
}
