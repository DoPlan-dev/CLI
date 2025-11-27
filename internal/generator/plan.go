package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/DoPlan-dev/CLI/internal/utils"
	"github.com/DoPlan-dev/CLI/pkg/models"
)

// PlanGenerator generates the .do/ directory structure
type PlanGenerator struct{}

// Name returns the name of the generator
func (g *PlanGenerator) Name() string {
	return ".do Structure"
}

// Generate creates the .do/ directory structure with all planning documents
func (g *PlanGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// Create .do structure: core, system, plan
	coreDir := filepath.Join(projectPath, ".do", "core")
	systemDir := filepath.Join(projectPath, ".do", "system")
	planDir := filepath.Join(projectPath, ".do", "plan")

	// Create core directory (for templates)
	if err := utils.CreateDirectory(coreDir); err != nil {
		return fmt.Errorf("failed to create .do/core directory: %w", err)
	}

	// Create system directory (for system docs and history)
	if err := utils.CreateDirectory(systemDir); err != nil {
		return fmt.Errorf("failed to create .do/system directory: %w", err)
	}

	// Create history subdirectory in system
	historyDir := filepath.Join(systemDir, "history")
	if err := utils.CreateDirectory(historyDir); err != nil {
		return fmt.Errorf("failed to create .do/system/history directory: %w", err)
	}

	// Create content subdirectory in system for content creation
	contentDir := filepath.Join(systemDir, "content")
	if err := utils.CreateDirectory(contentDir); err != nil {
		return fmt.Errorf("failed to create .do/system/content directory: %w", err)
	}

	// Create plan directory (for feature folders, TASKS.md, active_state.json)
	if err := utils.CreateDirectory(planDir); err != nil {
		return fmt.Errorf("failed to create .do/plan directory: %w", err)
	}

	// Generate IDEA.md
	if err := generateIDEA(systemDir, request); err != nil {
		return fmt.Errorf("failed to generate IDEA.md: %w", err)
	}

	// Generate BRAINSTORM.md
	if err := generateBRAINSTORM(systemDir, request); err != nil {
		return fmt.Errorf("failed to generate BRAINSTORM.md: %w", err)
	}

	// Generate PRD.md
	if err := generatePRD(systemDir, request); err != nil {
		return fmt.Errorf("failed to generate PRD.md: %w", err)
	}

	// Generate ARCHITECTURE.md
	if err := generateARCHITECTURE(systemDir, request); err != nil {
		return fmt.Errorf("failed to generate ARCHITECTURE.md: %w", err)
	}

	// Generate DESIGN_SYSTEM.md
	if err := generateDESIGNSYSTEM(systemDir, request); err != nil {
		return fmt.Errorf("failed to generate DESIGN_SYSTEM.md: %w", err)
	}

	// Generate TASKS.md in plan directory
	if err := generateTASKS(planDir, request); err != nil {
		return fmt.Errorf("failed to generate TASKS.md: %w", err)
	}

	// Generate active_state.json in history directory (state tracking)
	if err := generateActiveState(historyDir, request); err != nil {
		return fmt.Errorf("failed to generate active_state.json: %w", err)
	}

	// Generate brainstorm templates in core/brainstorm/
	if err := generateBrainstormTemplates(coreDir, request); err != nil {
		return fmt.Errorf("failed to generate brainstorm templates: %w", err)
	}

	// Generate all templates in core/templates/
	if err := generateCoreTemplates(coreDir, request); err != nil {
		return fmt.Errorf("failed to generate core templates: %w", err)
	}

	// Generate content folder structure in system/content/
	if err := generateContentStructure(contentDir, request); err != nil {
		return fmt.Errorf("failed to generate content structure: %w", err)
	}

	// Move BRAINSTORM.md to core/ (it's a template/process document)
	// Note: This will be generated in system/ first, then we could move it, but for now keep it in system/
	// If you want it in core/, we can change generateBRAINSTORM to write to coreDir instead

	// Mirror key documents into docs/ hierarchy
	if err := mirrorPlanDocsToDocs(projectPath); err != nil {
		return fmt.Errorf("failed to mirror docs into docs/: %w", err)
	}

	return nil
}

// generateIDEA generates IDEA.md template
func generateIDEA(systemDir string, request *models.ProjectRequest) error {
	content := `# Project Idea

## Overview
Describe your project idea here.

## Goals
- Goal 1
- Goal 2
- Goal 3

## Target Users
- User persona 1
- User persona 2

## Key Features
- Feature 1
- Feature 2
- Feature 3

## Success Metrics
- Metric 1
- Metric 2

---
` + generateFooter(request.ProjectName, "Product Manager") + `
`
	path := filepath.Join(systemDir, "IDEA.md")
	return utils.WriteFile(path, []byte(content))
}

// generateBRAINSTORM generates BRAINSTORM.md template
func generateBRAINSTORM(systemDir string, request *models.ProjectRequest) error {
	content := `# Brainstorm Session

## Questions & Insights

### Product Manager
**Questions**:
- What problem are we solving?
- Who is our target audience?

**Insights**:
- Insight 1
- Insight 2

### Engineering Lead
**Questions**:
- What are the technical requirements?
- What are the constraints?

**Insights**:
- Insight 1
- Insight 2

### Design Manager
**Questions**:
- What is the user experience we want to create?
- What are the design constraints?

**Insights**:
- Insight 1
- Insight 2

---
` + generateFooter(request.ProjectName, "Product Manager") + `
`
	path := filepath.Join(systemDir, "BRAINSTORM.md")
	return utils.WriteFile(path, []byte(content))
}

// generatePRD generates PRD.md template
func generatePRD(systemDir string, request *models.ProjectRequest) error {
	content := `# Product Requirements Document

## Product Overview
**Product Name**: ` + request.ProjectName + `
**Version**: 1.0.0
**Status**: Draft

## User Personas
### Primary Persona
- **Name**: [Persona Name]
- **Role**: [Role]
- **Goals**: [Goals]
- **Pain Points**: [Pain Points]

## Features
### Feature 1
**Priority**: P0
**Description**: [Description]
**User Stories**:
- As a [user], I want [goal] so that [benefit]

**Acceptance Criteria**:
- [ ] Criterion 1
- [ ] Criterion 2

## Success Metrics
- Metric 1: [Target]
- Metric 2: [Target]

## Timeline
- Phase 1: [Dates]
- Phase 2: [Dates]

---
` + generateFooter(request.ProjectName, "Product Manager") + `
`
	path := filepath.Join(systemDir, "PRD.md")
	return utils.WriteFile(path, []byte(content))
}

// generateARCHITECTURE generates ARCHITECTURE.md template
func generateARCHITECTURE(systemDir string, request *models.ProjectRequest) error {
	content := `# Technical Architecture

## System Overview
**Project**: ` + request.ProjectName + `
**Version**: 1.0.0
**Status**: Draft

## Technology Stack
### Frontend
- Framework: [Framework]
- UI Library: [Library]

### Backend
- Language: [Language]
- Framework: [Framework]

### Database
- Type: [Type]
- Database: [Database]

## System Design
### Architecture Pattern
[Pattern description]

### Components
1. Component 1
2. Component 2
3. Component 3

## API Design
### Endpoints
- GET /api/...
- POST /api/...

## Data Models
### Model 1
` + "```" + `
{
  "field1": "type",
  "field2": "type"
}
` + "```" + `

## Infrastructure
### Deployment
- Platform: [Platform]
- Environment: [Environment]

---
` + generateFooter(request.ProjectName, "System Architect") + `
`
	path := filepath.Join(systemDir, "ARCHITECTURE.md")
	return utils.WriteFile(path, []byte(content))
}

// generateDESIGNSYSTEM generates DESIGN_SYSTEM.md template
func generateDESIGNSYSTEM(systemDir string, request *models.ProjectRequest) error {
	content := `# Design System

## Design Principles
1. Principle 1
2. Principle 2
3. Principle 3

## Color Palette
### Primary Colors
- Primary: #000000
- Secondary: #FFFFFF

### Semantic Colors
- Success: #00FF00
- Error: #FF0000
- Warning: #FFFF00

## Typography
### Headings
- H1: [Font], [Size], [Weight]
- H2: [Font], [Size], [Weight]

### Body
- Body: [Font], [Size], [Weight]

## Components
### Button
- Primary Button
- Secondary Button

### Form Elements
- Input Field
- Select Dropdown

## Spacing
- Base Unit: 8px
- Scale: 8, 16, 24, 32, 40, 48

---
` + generateFooter(request.ProjectName, "Design & UX Manager") + `
`
	path := filepath.Join(systemDir, "DESIGN_SYSTEM.md")
	return utils.WriteFile(path, []byte(content))
}

// generateTASKS generates TASKS.md template
func generateTASKS(planDir string, request *models.ProjectRequest) error {
	content := `# Implementation Tasks

## Phase 1: Foundation
**Status**: ⏳ Pending

### Task 1.1
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: None
**Effort**: [Hours]

## Phase 2: Core Features
**Status**: ⏳ Pending

### Task 2.1
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: Phase 1
**Effort**: [Hours]

---
` + generateFooter(request.ProjectName, "Engineering Lead") + `

*Use /plan command to generate detailed tasks from approved plan*
`
	path := filepath.Join(planDir, "TASKS.md")
	return utils.WriteFile(path, []byte(content))
}

// generateActiveState generates active_state.json with initial state
func generateActiveState(historyDir string, request *models.ProjectRequest) error {
	state := map[string]interface{}{
		"phase":       "idea",
		"active_task": nil,
		"completed":   []string{},
		"locked":      false,
	}

	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal active_state.json: %w", err)
	}

	// Store active_state.json in history/ directory (current state tracking)
	path := filepath.Join(historyDir, "active_state.json")
	return utils.WriteFile(path, jsonData)
}

// generateBrainstormTemplates creates all brainstorm phase templates in core/brainstorm/
// Uses file-based templates if available, falls back to hardcoded templates
func generateBrainstormTemplates(coreDir string, request *models.ProjectRequest) error {
	brainstormDir := filepath.Join(coreDir, "brainstorm")
	if err := utils.CreateDirectory(brainstormDir); err != nil {
		return fmt.Errorf("failed to create brainstorm directory: %w", err)
	}

	// Try to load templates from embedded files first
	templates := make(map[string]string)

	// Load phase templates
	phases := map[string]string{
		"phase-01.md": "01",
		"phase-02.md": "02",
		"phase-03.md": "03",
		"phase-04.md": "04",
		"phase-05.md": "05",
		"phase-06.md": "06",
	}

	for filename, phaseNum := range phases {
		content, err := LoadBrainstormPhaseTemplate(phaseNum)
		if err == nil {
			templates[filename] = content
		} else {
			// Fallback to hardcoded template (already defined in original function)
			// This will be handled below if templates map is incomplete
		}
	}

	// Load confirmation template
	confirmationContent, err := LoadBrainstormConfirmationTemplate()
	if err == nil {
		templates["CONFIRMATION_TEMPLATE.md"] = confirmationContent
	}

	// Load output template
	outputContent, err := LoadBrainstormOutputTemplate()
	if err == nil {
		// Add footer to output template
		outputContent += "\n" + generateFooter(request.ProjectName, "Documentation Lead") + "\n"
		templates["TEMPLATE_BRAINSTORM.md"] = outputContent
	}

	// If we couldn't load all templates from files, fall back to hardcoded
	if len(templates) < 8 {
		// Fallback to original hardcoded templates
		return generateBrainstormTemplatesFallback(brainstormDir, request)
	}

	// Write all templates to disk
	for filename, content := range templates {
		path := filepath.Join(brainstormDir, filename)
		if err := utils.WriteFile(path, []byte(content)); err != nil {
			return fmt.Errorf("failed to write template %s: %w", filename, err)
		}
	}

	return nil
}

// generateBrainstormTemplatesFallback provides hardcoded templates as fallback
func generateBrainstormTemplatesFallback(brainstormDir string, request *models.ProjectRequest) error {
	// Phase 01: Vision & Outcomes
	phase01 := `# Phase 01: Vision & Outcomes
**Led by: Product Manager**

## Purpose
Define the core vision, success metrics, and desired outcomes for the project.

## Key Questions

1. **What is the primary problem we're solving?**
   - What pain point does this address?
   - Why is this problem important?

2. **What is the vision for this product?**
   - What does success look like in 6 months? 1 year?
   - What impact do we want to make?

3. **What are the key success metrics?**
   - How will we measure success?
   - What KPIs matter most?

4. **What are the must-have outcomes?**
   - What absolutely must be achieved?
   - What would make this project a failure if missing?

5. **What is the unique value proposition?**
   - What makes this different from existing solutions?
   - Why would users choose this over alternatives?
`

	// Phase 02: Audience & Differentiation
	phase02 := `# Phase 02: Audience & Differentiation
**Led by: Product Manager + Design Manager**

## Purpose
Identify target users, their needs, and how we differentiate from competitors.

## Key Questions

1. **Who is the primary target audience?**
   - Demographics, psychographics, behaviors
   - Who will use this most?

2. **Who are the secondary audiences?**
   - Other user groups that matter
   - How do their needs differ?

3. **What are the user personas?**
   - Create 2-3 detailed personas
   - What are their goals, frustrations, motivations?

4. **Who are the main competitors?**
   - Direct and indirect competitors
   - What do they do well? What are their weaknesses?

5. **How do we differentiate?**
   - What unique features or approaches set us apart?
   - What can we do better than competitors?

6. **What are the user's current alternatives?**
   - How do they solve this problem today?
   - What workarounds exist?
`

	// Phase 03: Experience, UI/UX & Tech
	phase03 := `# Phase 03: Experience, UI/UX & Tech
**Led by: Design Manager + Engineering Lead**

## Purpose
Define the user experience, design approach, and technical foundation.

## Key Questions

1. **What is the ideal user experience?**
   - Walk through the user journey from start to finish
   - What emotions should users feel?

2. **What are the key user flows?**
   - Primary flows (e.g., create profile, view analytics)
   - Secondary flows (e.g., share profile, customize QR)

3. **What is the design aesthetic?**
   - Modern, minimalist, bold, professional?
   - What design inspiration or references?

4. **What are the technical requirements?**
   - Performance requirements (load time, responsiveness)
   - Scalability needs
   - Integration requirements

5. **What technology stack should we use?**
   - Frontend framework preferences?
   - Backend architecture?
   - Database choices?

6. **What are the design constraints?**
   - Platform limitations?
   - Browser support requirements?
   - Accessibility standards?
`

	// Phase 04: Content & SEO
	phase04 := `# Phase 04: Content & SEO
**Led by: Content Strategist + SEO Specialist**

## Purpose
Define content strategy, messaging, and SEO approach.

## Key Questions

1. **What is the core messaging?**
   - What should users understand immediately?
   - What is the elevator pitch?

2. **What content is needed?**
   - Landing page copy
   - Product descriptions
   - Help documentation
   - Marketing materials

3. **What is the SEO strategy?**
   - Target keywords
   - Content optimization approach
   - Link building strategy

4. **What is the tone of voice?**
   - Professional, friendly, technical, casual?
   - Brand personality?

5. **What content formats?**
   - Blog posts, videos, guides, tutorials?
   - User-generated content?
`

	// Phase 05: Marketing & Growth
	phase05 := `# Phase 05: Marketing & Growth
**Led by: Marketing Manager**

## Purpose
Define marketing strategy, growth channels, and user acquisition.

## Key Questions

1. **What are the primary marketing channels?**
   - Social media, content marketing, paid ads, partnerships?
   - Which channels align with target audience?

2. **What is the growth strategy?**
   - Organic growth tactics
   - Viral/word-of-mouth mechanisms
   - Referral programs

3. **What is the launch strategy?**
   - Pre-launch activities
   - Launch day plan
   - Post-launch follow-up

4. **What are the key metrics to track?**
   - User acquisition cost
   - Conversion rates
   - Retention metrics

5. **What partnerships or collaborations?**
   - Influencer partnerships?
   - Integration partnerships?
   - Community building?
`

	// Phase 06: Delivery, Ops & Risks
	phase06 := `# Phase 06: Delivery, Ops & Risks
**Led by: Engineering Lead + Project Orchestrator**

## Purpose
Define delivery approach, operations, and risk mitigation.

## Key Questions

1. **What is the delivery timeline?**
   - MVP timeline
   - Full launch timeline
   - Milestone dates

2. **What are the operational requirements?**
   - Hosting and infrastructure
   - Monitoring and logging
   - Backup and disaster recovery

3. **What are the main risks?**
   - Technical risks
   - Market risks
   - Resource risks

4. **How will we mitigate risks?**
   - Risk mitigation strategies
   - Contingency plans

5. **What are the success criteria for launch?**
   - Minimum viable features
   - Performance benchmarks
   - Quality standards

6. **What is the maintenance plan?**
   - Update frequency
   - Support structure
   - Long-term roadmap
`

	// CONFIRMATION_TEMPLATE
	confirmationTemplate := `# Brainstorm Session Summary

## Review & Confirm

Please review the following summary of our brainstorming session. You can:
1. **Save it** - Confirm and save to BRAINSTORM.md
2. **Revise a phase** - Request changes to specific phase(s)
3. **Add information** - Add additional details to any phase
4. **Start over** - Restart the brainstorming process

---

## Phase 01: Vision & Outcomes ✅
[Summary of answers]

## Phase 02: Audience & Differentiation ✅
[Summary of answers]

## Phase 03: Experience, UI/UX & Tech ✅
[Summary of answers]

## Phase 04: Content & SEO ✅
[Summary of answers]

## Phase 05: Marketing & Growth ✅
[Summary of answers]

## Phase 06: Delivery, Ops & Risks ✅
[Summary of answers]

---

## Next Steps

Once confirmed, this summary will be saved to .do/system/BRAINSTORM.md and you can proceed with /write to generate PRD, Architecture, and Design System documents.
`

	// TEMPLATE_BRAINSTORM
	templateBrainstorm := `# Brainstorm Session

*Generated: [DATE]*

## Overview
This document captures the comprehensive brainstorming session for the project, organized by phase.

---

## Phase 01: Vision & Outcomes
**Led by: Product Manager**

### Questions & Answers
[Organized Q&A from Phase 01]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 02: Audience & Differentiation
**Led by: Product Manager + Design Manager**

### Questions & Answers
[Organized Q&A from Phase 02]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 03: Experience, UI/UX & Tech
**Led by: Design Manager + Engineering Lead**

### Questions & Answers
[Organized Q&A from Phase 03]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 04: Content & SEO
**Led by: Content Strategist + SEO Specialist**

### Questions & Answers
[Organized Q&A from Phase 04]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 05: Marketing & Growth
**Led by: Marketing Manager**

### Questions & Answers
[Organized Q&A from Phase 05]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 06: Delivery, Ops & Risks
**Led by: Engineering Lead + Project Orchestrator**

### Questions & Answers
[Organized Q&A from Phase 06]

### Key Insights
- [Insight 1]
- [Insight 2]

---

` + generateFooter(request.ProjectName, "Documentation Lead") + `
`

	templates := map[string]string{
		"phase-01.md":              phase01,
		"phase-02.md":              phase02,
		"phase-03.md":              phase03,
		"phase-04.md":              phase04,
		"phase-05.md":              phase05,
		"phase-06.md":              phase06,
		"CONFIRMATION_TEMPLATE.md": confirmationTemplate,
		"TEMPLATE_BRAINSTORM.md":   templateBrainstorm,
	}

	for filename, content := range templates {
		path := filepath.Join(brainstormDir, filename)
		if err := utils.WriteFile(path, []byte(content)); err != nil {
			return fmt.Errorf("failed to write template %s: %w", filename, err)
		}
	}

	return nil
}

// generateCoreTemplates creates all core templates in core/templates/
func generateCoreTemplates(coreDir string, request *models.ProjectRequest) error {
	templatesDir := filepath.Join(coreDir, "templates")
	if err := utils.CreateDirectory(templatesDir); err != nil {
		return fmt.Errorf("failed to create templates directory: %w", err)
	}

	// PRD Template
	prdTemplate := `# Product Requirements Document

## Product Overview
**Product Name**: [Product Name]
**Version**: 1.0.0
**Status**: Draft

## Goals & Success Metrics
| Goal | KPI | Target |
| --- | --- | --- |
| [Goal 1] | [Metric] | [Target] |
| [Goal 2] | [Metric] | [Target] |

## User Personas
### Persona 1
- **Name**: [Name]
- **Role**: [Role]
- **Goals**: [Goals]
- **Pain Points**: [Pain Points]
- **Needs**: [Needs]

## User Journey
1. [Step 1]
2. [Step 2]
3. [Step 3]

## Feature Requirements

### Feature 1 (P0)
**Description**: [Description]

**User Stories**:
- As a [user], I want [goal] so that [benefit]

**Acceptance Criteria**:
- [ ] Criterion 1
- [ ] Criterion 2

## Non-Functional Requirements
- **Performance**: [Requirements]
- **Security**: [Requirements]
- **Accessibility**: [Requirements]

## Timeline & Milestones
- **Phase 1**: [Dates] - [Description]
- **Phase 2**: [Dates] - [Description]

---
*Template - Customize as needed*
`

	// Architecture Template
	architectureTemplate := `# Technical Architecture

## System Overview
**Project**: [Project Name]
**Version**: 1.0.0
**Status**: Draft

## Technology Stack
| Layer | Technology | Rationale |
| --- | --- | --- |
| Frontend | [Framework] | [Reason] |
| Backend | [Framework] | [Reason] |
| Database | [Database] | [Reason] |
| Infrastructure | [Platform] | [Reason] |

## System Design
### Architecture Pattern
[Pattern description - e.g., Client-Driven Modular Monolith, Microservices, etc.]

### Core Components
1. **Component 1**
   - Purpose: [Purpose]
   - Responsibilities: [Responsibilities]

2. **Component 2**
   - Purpose: [Purpose]
   - Responsibilities: [Responsibilities]

## Request Flow
1. [Step 1]
2. [Step 2]
3. [Step 3]

## API Design
### Endpoints
- GET /api/[resource] - [Description]
- POST /api/[resource] - [Description]

## Data Models
### Model 1
{
  "id": "uuid",
  "field1": "type",
  "field2": "type"
}

## Infrastructure
### Deployment
- **Platform**: [Platform]
- **Environments**: [Environments]

### Security & Compliance
- [Security measures]
- [Compliance requirements]

---
*Template - Customize as needed*
`

	// Design System Template
	designSystemTemplate := `# Design System

## Design Principles
1. **Principle 1** - [Description]
2. **Principle 2** - [Description]
3. **Principle 3** - [Description]

## Color Palette
| Token | Hex | Usage |
| --- | --- | --- |
| --color-primary | #000000 | Primary actions, CTAs |
| --color-secondary | #000000 | Secondary actions |
| --color-bg | #000000 | Background surfaces |

### Semantic Colors
- **Success**: #22C55E
- **Warning**: #FACC15
- **Error**: #F87171
- **Info**: #38BDF8

## Typography
| Role | Font | Size | Weight | Notes |
| --- | --- | --- | --- | --- |
| H1 | [Font] | 48px | 600 | Headings |
| Body | [Font] | 16px | 400 | Body text |

## Layout & Spacing
- **Base unit**: 8px
- **Section spacing**: 32px
- **Card padding**: 24px

## Components
### Buttons
- **Primary Button**: [Specifications]
- **Secondary Button**: [Specifications]

### Form Elements
- **Input**: [Specifications]
- **Select**: [Specifications]

## Motion
- **Transitions**: 200ms ease-out
- **Animations**: [Guidelines]

## Accessibility
- Color contrast ratio >= 4.5:1
- Keyboard navigation support
- Screen reader labels

---
*Template - Customize as needed*
`

	// IDE Template
	ideTemplate := `# IDE Configuration

## Overview
This document describes IDE-specific configurations and setup for the project.

## Supported IDEs
- [ ] Cursor
- [ ] Claude Code
- [ ] Windsurf
- [ ] Cline
- [ ] Antigravity
- [ ] OpenCode

## IDE-Specific Settings

### Cursor
- **Rules**: Located in .cursor/rules/
- **Commands**: Located in .do/core/commands/ (symlinked to .cursor/commands/)
- **Agents**: Located in .do/core/agents/ (symlinked to .cursor/agents/)

### Claude Code
- **Config**: docs/CLAUDE.md
- **Rules**: Located in .claude/rules/

## Recommended Extensions
- [Extension 1]
- [Extension 2]

## Workspace Settings
{
  "settings": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "[formatter]"
  }
}

---
*Template - Customize as needed*
`

	// Brainstorm Template
	brainstormTemplate := `# Brainstorm Session

*Date: [DATE]*

## Overview
This document captures the brainstorming session for the project.

## Phase 01: Vision & Outcomes
**Led by: Product Manager**

### Questions & Answers
[Q&A from Phase 01]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 02: Audience & Differentiation
**Led by: Product Manager + Design Manager**

### Questions & Answers
[Q&A from Phase 02]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 03: Experience, UI/UX & Tech
**Led by: Design Manager + Engineering Lead**

### Questions & Answers
[Q&A from Phase 03]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 04: Content & SEO
**Led by: Content Strategist + SEO Specialist**

### Questions & Answers
[Q&A from Phase 04]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 05: Marketing & Growth
**Led by: Marketing Manager**

### Questions & Answers
[Q&A from Phase 05]

### Key Insights
- [Insight 1]
- [Insight 2]

---

## Phase 06: Delivery, Ops & Risks
**Led by: Engineering Lead + Project Orchestrator**

### Questions & Answers
[Q&A from Phase 06]

### Key Insights
- [Insight 1]
- [Insight 2]

---
*Template - Customize as needed*
`

	// Contracts Template
	contractsTemplate := `# Contracts

## Overview
This document defines API contracts, data schemas, and integration agreements.

## API Contracts

### Endpoint: GET /api/[resource]
**Description**: [Description]

**Request**:
{
  "param1": "type",
  "param2": "type"
}

**Response**:
{
  "data": {},
  "status": "success"
}

**Status Codes**:
- 200 - Success
- 400 - Bad Request
- 404 - Not Found

## Data Schemas

### Schema: [Schema Name]
interface SchemaName {
  id: string;
  field1: string;
  field2: number;
}

## Integration Contracts
- [Integration 1]: [Description]
- [Integration 2]: [Description]

---
*Template - Customize as needed*
`

	// Design Template
	designTemplate := `# Design: [Feature Name]

**Task ID**: [ID]
**Status**: Draft

## Overview
[Feature description]

## Design Decisions
- [ ] Decision 1
- [ ] Decision 2
- [ ] Decision 3

## UI/UX Considerations
- **User Flow**: [Description]
- **Accessibility**: [Requirements]
- **Responsive Design**: [Breakpoints]

## Visual Mockups
Add mockups, wireframes, or design references here.

## Design Tokens
- Colors: [Reference to design system]
- Typography: [Reference to design system]
- Spacing: [Reference to design system]

## Components Needed
- [ ] Component 1
- [ ] Component 2

---
*Template - Customize as needed*
`

	// Plan Template
	planTemplate := `# Plan: [Feature Name]

**Task ID**: [ID]
**Status**: Draft

## Objectives
- [ ] Objective 1
- [ ] Objective 2
- [ ] Objective 3

## Approach
[Implementation approach description]

## Scope
- **In Scope**: [Items]
- **Out of Scope**: [Items]

## Dependencies
- Dependency 1
- Dependency 2

## Timeline
- **Start**: [Date]
- **Target Completion**: [Date]
- **Milestones**: [List]

## Risks & Mitigations
- **Risk 1**: [Description] → **Mitigation**: [Strategy]
- **Risk 2**: [Description] → **Mitigation**: [Strategy]

---
*Template - Customize as needed*
`

	// Prompts Template
	promptsTemplate := `# Prompts: [Feature Name]

**Task ID**: [ID]
**Status**: Draft

## AI Prompts Used
This file tracks prompts and AI interactions for this feature.

### Prompt 1
[Prompt text]

**Response**: [Summary of response]

**Date**: [Date]

### Prompt 2
[Prompt text]

**Response**: [Summary of response]

**Date**: [Date]

## Key Learnings
- [Learning 1]
- [Learning 2]

---
*Template - Customize as needed*
`

	// Feature Tasks Template
	featureTasksTemplate := `# Feature Tasks: [Feature Name]

**Feature ID**: [ID]
**Status**: Draft

## Checklist
- [ ] Task 1
- [ ] Task 2
- [ ] Task 3

## Subtasks

### Subtask 1
- [ ] Step 1
- [ ] Step 2
- [ ] Step 3

### Subtask 2
- [ ] Step 1
- [ ] Step 2

## Dependencies
- [Dependency 1]
- [Dependency 2]

## Notes
[Implementation notes]

---
*Template - Customize as needed*
`

	// Phase Tasks Template
	phaseTasksTemplate := `# Phase Tasks: [Phase Name]

**Phase**: [Phase Number]
**Status**: Draft

## Phase Overview
[Phase description]

## Tasks

### Task 1.1
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: None
**Effort**: [Hours]

### Task 1.2
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: Task 1.1
**Effort**: [Hours]

## Phase Milestones
- [ ] Milestone 1
- [ ] Milestone 2

## Phase Dependencies
- [Dependency 1]
- [Dependency 2]

---
*Template - Customize as needed*
`

	// Project Tasks Template
	projectTasksTemplate := `# Project Tasks

**Project**: [Project Name]
**Status**: Draft

## Implementation Tasks

## Phase 1: Foundation
**Objective**: [Objective]
**Status**: ⏳ Pending

### Task 1.1
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: None
**Effort**: [Hours]

### Task 1.2
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: Task 1.1
**Effort**: [Hours]

---

## Phase 2: Core Features
**Objective**: [Objective]
**Status**: ⏳ Pending

### Task 2.1
**Description**: [Description]
**Status**: ⏳ Pending
**Dependencies**: Phase 1
**Effort**: [Hours]

---

## Progress Tracking
- **Total Tasks**: [Number]
- **Completed**: [Number]
- **In Progress**: [Number]
- **Pending**: [Number]

---
*Template - Customize as needed*
`

	// README Template
	readmeTemplate := `# [Project Name]

## Overview
[Project description]

## Quick Start

# Install dependencies
npm install

# Run development server
npm run dev

## Project Structure

project/
  ├── src/
  ├── .do/
  ├── docs/
  └── README.md

## Commands

- /tell - Capture your idea
- /meeting - Discovery meeting with adaptive speed options
- /write - Generate documents
- /plan - Generate execution plan
- /build - Start coding

## Documentation

- [PRD](.do/system/PRD.md)
- [Architecture](.do/system/ARCHITECTURE.md)
- [Design System](.do/system/DESIGN_SYSTEM.md)

## Contributing

[Contributing guidelines]

---
*Template - Customize as needed*
`

	templates := map[string]string{
		"PRD.md":           prdTemplate,
		"ARCHITECTURE.md":  architectureTemplate,
		"DESIGN_SYSTEM.md": designSystemTemplate,
		"IDE.md":           ideTemplate,
		"BRAINSTORM.md":    brainstormTemplate,
		"contracts.md":     contractsTemplate,
		"design.md":        designTemplate,
		"plan.md":          planTemplate,
		"prompts.md":       promptsTemplate,
		"feature-tasks.md": featureTasksTemplate,
		"phase-tasks.md":   phaseTasksTemplate,
		"project-tasks.md": projectTasksTemplate,
		"README.md":        readmeTemplate,
	}

	for filename, content := range templates {
		path := filepath.Join(templatesDir, filename)
		if err := utils.WriteFile(path, []byte(content)); err != nil {
			return fmt.Errorf("failed to write template %s: %w", filename, err)
		}
	}

	return nil
}

func mirrorPlanDocsToDocs(projectPath string) error {
	docsRoot := filepath.Join(projectPath, "docs")
	if _, err := os.Stat(docsRoot); err != nil {
		// Docs directory might not exist yet; skip quietly
		return nil
	}

	foundationDir := filepath.Join(docsRoot, "foundation")
	if err := utils.CreateDirectory(foundationDir); err != nil {
		return err
	}

	systemDir := filepath.Join(projectPath, ".do", "system")
	foundationDocs := []string{"IDEA.md", "BRAINSTORM.md", "PRD.md", "ARCHITECTURE.md", "DESIGN_SYSTEM.md"}
	for _, name := range foundationDocs {
		src := filepath.Join(systemDir, name)
		data, err := os.ReadFile(src)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		dest := filepath.Join(foundationDir, name)
		if err := utils.WriteFile(dest, data); err != nil {
			return err
		}
	}

	// Mirror TASKS.md into docs/features/_plan/TASKS.md for quick access
	tasksSrc := filepath.Join(projectPath, ".do", "plan", "TASKS.md")
	if data, err := os.ReadFile(tasksSrc); err == nil {
		destDir := filepath.Join(docsRoot, "features", "_plan")
		if err := utils.CreateDirectory(destDir); err != nil {
			return err
		}
		if err := utils.WriteFile(filepath.Join(destDir, "TASKS.md"), data); err != nil {
			return err
		}
	}

	if err := scaffoldFeaturePhaseDocs(projectPath, docsRoot); err != nil {
		return err
	}

	return nil
}

func scaffoldFeaturePhaseDocs(projectPath, docsRoot string) error {
	tasksPath := filepath.Join(projectPath, ".do", "plan", "TASKS.md")
	data, err := os.ReadFile(tasksPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	rePhase := regexp.MustCompile(`^##\s+Phase\s+(\d+):\s+(.*)$`)
	lines := strings.Split(string(data), "\n")
	seen := make(map[string]bool)

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		matches := rePhase.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}
		numStr := matches[1]
		title := strings.TrimSpace(matches[2])
		slug := buildFeatureSlug(numStr, title)
		if seen[slug] {
			continue
		}
		seen[slug] = true

		dir := filepath.Join(docsRoot, "features", slug)
		if err := utils.CreateDirectory(dir); err != nil {
			return err
		}
		readmePath := filepath.Join(dir, "README.md")
		content := fmt.Sprintf(`# Phase %s · %s

This folder holds specs, prompts, and history for **Phase %s – %s**.

- Source checklist: `+"`docs/features/_plan/TASKS.md`"+`
- Original plan section: `+"`.do/plan/TASKS.md`"+` (Phase %s)

Add discovery notes, prompt logs, and feature-specific contracts here so the docs/features tree mirrors the task plan.
`, padPhaseNumber(numStr), title, padPhaseNumber(numStr), title, padPhaseNumber(numStr))
		if err := utils.WriteFile(readmePath, []byte(content)); err != nil {
			return err
		}
	}

	return nil
}

func buildFeatureSlug(numStr, title string) string {
	num := padPhaseNumber(numStr)
	name := strings.ToLower(title)
	name = strings.ReplaceAll(name, " ", "_")
	var b strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			b.WriteRune(r)
		}
	}
	sanitized := b.String()
	if sanitized == "" {
		sanitized = "feature"
	}
	return fmt.Sprintf("%s_%s", num, sanitized)
}

func padPhaseNumber(numStr string) string {
	if n, err := strconv.Atoi(numStr); err == nil {
		// Add 1 to phase number so Phase 0 becomes 01, Phase 1 becomes 02, etc.
		return fmt.Sprintf("%02d", n+1)
	}
	return numStr
}

// ScaffoldPlanHierarchy scaffolds the phase/feature hierarchy from TASKS.md
// This is called by the /plan command to create the structured planning folders
func ScaffoldPlanHierarchy(projectPath string) error {
	tasksPath := filepath.Join(projectPath, ".do", "plan", "TASKS.md")
	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return fmt.Errorf("failed to read TASKS.md: %w", err)
	}

	planDir := filepath.Join(projectPath, ".do", "plan")
	phases, err := parsePhasesFromTasks(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse phases: %w", err)
	}

	for _, phase := range phases {
		phaseDir := filepath.Join(planDir, phase.FolderName)
		if err := utils.CreateDirectory(phaseDir); err != nil {
			return fmt.Errorf("failed to create phase directory %s: %w", phaseDir, err)
		}

		// Create contracts directory for shared schemas
		contractsDir := filepath.Join(phaseDir, "_contracts")
		if err := utils.CreateDirectory(contractsDir); err != nil {
			return fmt.Errorf("failed to create contracts directory: %w", err)
		}
		// Add README to contracts
		contractsReadme := `# Contracts

This directory contains shared API/data schemas for **` + phase.Name + `**.

## Usage
- API schemas: Define request/response structures
- Database schemas: Define data models and relationships
- Type definitions: Shared TypeScript/Go types

## Files
Add your contract files here (e.g., ` + "`api.json`" + `, ` + "`schema.sql`" + `, ` + "`types.ts`" + `).

---
` + generateFooter(filepath.Base(projectPath), "System Architect") + `
`
		if err := utils.WriteFile(filepath.Join(contractsDir, "README.md"), []byte(contractsReadme)); err != nil {
			return fmt.Errorf("failed to create contracts README: %w", err)
		}

		// Scaffold feature folders for each task in the phase
		for _, feature := range phase.Features {
			featureDir := filepath.Join(phaseDir, feature.FolderName)
			if err := utils.CreateDirectory(featureDir); err != nil {
				return fmt.Errorf("failed to create feature directory %s: %w", featureDir, err)
			}

			// Generate template files
			// Extract project name from project path
			projectName := filepath.Base(projectPath)
			if err := generateFeatureTemplates(featureDir, feature, projectName); err != nil {
				return fmt.Errorf("failed to generate feature templates: %w", err)
			}
		}
	}

	return nil
}

type Phase struct {
	Number     string
	Name       string
	FolderName string
	Features   []Feature
}

type Feature struct {
	ID          string
	Title       string
	FolderName  string
	Description string
}

func parsePhasesFromTasks(content string) ([]Phase, error) {
	var phases []Phase
	lines := strings.Split(content, "\n")

	rePhase := regexp.MustCompile(`^##\s+Phase\s+(\d+):\s+(.+?)(?:\s*\(|$)`)
	reTask := regexp.MustCompile(`^###\s+(\d+\.\d+)\s+(.+)$`)

	var currentPhase *Phase

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for phase header
		if matches := rePhase.FindStringSubmatch(trimmed); len(matches) == 3 {
			// Save previous phase if exists
			if currentPhase != nil {
				phases = append(phases, *currentPhase)
			}

			numStr := matches[1]
			name := strings.TrimSpace(matches[2])
			currentPhase = &Phase{
				Number:     numStr,
				Name:       name,
				FolderName: buildPhaseFolderName(numStr, name),
				Features:   []Feature{},
			}
			continue
		}

		// Check for task header (only if we're in a phase)
		if currentPhase != nil {
			if matches := reTask.FindStringSubmatch(trimmed); len(matches) == 3 {
				taskID := matches[1]
				taskTitle := strings.TrimSpace(matches[2])

				// Try to find description in next few lines
				description := ""
				for j := i + 1; j < len(lines) && j < i+5; j++ {
					nextLine := strings.TrimSpace(lines[j])
					if strings.HasPrefix(nextLine, "**Description**") {
						parts := strings.SplitN(nextLine, ":", 2)
						if len(parts) == 2 {
							description = strings.TrimSpace(parts[1])
						}
						break
					}
				}

				feature := Feature{
					ID:          taskID,
					Title:       taskTitle,
					FolderName:  buildFeatureFolderName(taskID, taskTitle),
					Description: description,
				}
				currentPhase.Features = append(currentPhase.Features, feature)
			}
		}
	}

	// Don't forget the last phase
	if currentPhase != nil {
		phases = append(phases, *currentPhase)
	}

	return phases, nil
}

func buildPhaseFolderName(numStr, name string) string {
	num := padPhaseNumber(numStr)
	slug := sanitizeForFolder(name)
	return fmt.Sprintf("%s-%s", num, slug)
}

func buildFeatureFolderName(taskID, title string) string {
	// Extract the second number from task ID (e.g., "1.1" -> "01")
	parts := strings.Split(taskID, ".")
	if len(parts) >= 2 {
		if n, err := strconv.Atoi(parts[1]); err == nil {
			taskNum := fmt.Sprintf("%02d", n)
			slug := sanitizeForFolder(title)
			return fmt.Sprintf("%s-%s", taskNum, slug)
		}
	}
	// Fallback
	slug := sanitizeForFolder(title)
	return fmt.Sprintf("%s-%s", taskID, slug)
}

func sanitizeForFolder(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "&", "and")
	var b strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			b.WriteRune(r)
		}
	}
	result := b.String()
	// Remove multiple underscores
	result = regexp.MustCompile(`_+`).ReplaceAllString(result, "_")
	result = strings.Trim(result, "_")
	if result == "" {
		result = "feature"
	}
	return result
}

func generateFeatureTemplates(featureDir string, feature Feature, projectName string) error {
	templates := map[string]string{
		"design.md":  generateDesignTemplate(feature, projectName),
		"plan.md":    generatePlanTemplate(feature, projectName),
		"tasks.md":   generateTasksTemplate(feature, projectName),
		"prompts.md": generatePromptsTemplate(feature, projectName),
		"github.md":  generateGithubTemplate(feature, projectName),
	}

	for filename, content := range templates {
		path := filepath.Join(featureDir, filename)
		if err := utils.WriteFile(path, []byte(content)); err != nil {
			return fmt.Errorf("failed to write %s: %w", filename, err)
		}
	}

	return nil
}

func generateDesignTemplate(feature Feature, projectName string) string {
	return `# Design: ` + feature.Title + `

**Task ID**: ` + feature.ID + `
**Status**: Draft

## Overview
` + feature.Description + `

## Design Decisions
- [ ] Design decision 1
- [ ] Design decision 2

## UI/UX Considerations
- [ ] User flow
- [ ] Accessibility
- [ ] Responsive design

## Visual Mockups
Add mockups or wireframes here.

---
` + generateFooter(projectName, "Design & UX Manager") + `
`
}

func generatePlanTemplate(feature Feature, projectName string) string {
	return `# Plan: ` + feature.Title + `

**Task ID**: ` + feature.ID + `
**Status**: Draft

## Objectives
- [ ] Objective 1
- [ ] Objective 2

## Approach
Describe the implementation approach here.

## Dependencies
- Dependency 1
- Dependency 2

## Timeline
- Start: [Date]
- Target completion: [Date]

---
` + generateFooter(projectName, "Design & UX Manager") + `
`
}

func generateTasksTemplate(feature Feature, projectName string) string {
	return `# Tasks: ` + feature.Title + `

**Task ID**: ` + feature.ID + `
**Status**: Draft

## Checklist
- [ ] Subtask 1
- [ ] Subtask 2
- [ ] Subtask 3

## Notes
Add implementation notes here.

---
` + generateFooter(projectName, "Engineering Lead") + `
`
}

func generatePromptsTemplate(feature Feature, projectName string) string {
	return `# Prompts: ` + feature.Title + `

**Task ID**: ` + feature.ID + `
**Status**: Draft

## AI Prompts Used
This file tracks prompts and AI interactions for this feature.

### Prompt 1
` + "```" + `
[Prompt text]
` + "```" + `

**Response**: [Summary]

### Prompt 2
` + "```" + `
[Prompt text]
` + "```" + `

**Response**: [Summary]

---
` + generateFooter(projectName, "Documentation Lead") + `
`
}

func generateGithubTemplate(feature Feature, projectName string) string {
	branchName := "task/" + strings.ReplaceAll(feature.ID, ".", "-")
	return `# GitHub: ` + feature.Title + `

**Task ID**: ` + feature.ID + `
**Status**: Draft

## Issues
- [ ] Create GitHub issue
- [ ] Link to milestone

## Pull Requests
- [ ] Create PR template
- [ ] Add reviewers

## Branch Strategy
- Branch name: ` + "`" + branchName + "`" + `
- Base branch: ` + "`main`" + `

## Commits
Track commits related to this feature:
- ` + "`feat(" + feature.ID + "): [description]`" + `

---
` + generateFooter(projectName, "Design & UX Manager") + `
`
}

// generateContentStructure creates the content folder structure for content creation
func generateContentStructure(contentDir string, request *models.ProjectRequest) error {
	// Content folders organized by type
	contentFolders := []string{
		"app-pages",     // Application/web pages content
		"legal",         // Legal pages (Privacy Policy, Terms of Service, etc.)
		"social-media",  // Social media posts and content
		"blog",          // Blog posts and articles
		"documentation", // User documentation and guides
		"marketing",     // Marketing copy and campaigns
		"email",         // Email templates and campaigns
		"seo",           // SEO-optimized content and meta descriptions
	}

	// Create each folder with README
	for _, folder := range contentFolders {
		folderPath := filepath.Join(contentDir, folder)
		if err := utils.CreateDirectory(folderPath); err != nil {
			return fmt.Errorf("failed to create content folder %s: %w", folder, err)
		}

		// Create README for each folder
		readmePath := filepath.Join(folderPath, "README.md")
		readmeContent := generateContentFolderREADME(folder, request.ProjectName)
		if err := utils.WriteFile(readmePath, []byte(readmeContent)); err != nil {
			return fmt.Errorf("failed to create README for %s: %w", folder, err)
		}
	}

	// Create main content README
	mainReadmePath := filepath.Join(contentDir, "README.md")
	mainReadmeContent := `# Content Creation

This directory contains all content created for your project, organized by type.

## Content Types

### App Pages
Application and web page content including landing pages, feature pages, about pages, etc.

### Legal
Legal documents including Privacy Policy, Terms of Service, Cookie Policy, etc.

### Social Media
Social media posts, captions, and content for various platforms.

### Blog
Blog posts, articles, and long-form content.

### Documentation
User documentation, guides, tutorials, and help content.

### Marketing
Marketing copy, campaigns, advertisements, and promotional content.

### Email
Email templates, newsletters, and email campaign content.

### SEO
SEO-optimized content, meta descriptions, alt text, and keyword-optimized copy.

## Content Creation Process

1. **During Meeting**: Specify which content types you need
2. **Keyword Input**: Provide keywords or let LLM generate them
3. **Generation**: Content is created with SEO best practices
4. **Review**: Review and refine generated content
5. **Deployment**: Content is ready for use in your application

## SEO Guidelines

All content in this directory follows SEO best practices:
- Keyword optimization
- Meta descriptions
- Alt text for images
- Proper heading structure
- Internal linking opportunities
- Mobile-friendly formatting

## Content Ownership

You can either:
- **Full LLM Generation**: Let the AI create all content automatically
- **Keyword-Guided**: Provide keywords and let LLM create content around them
- **Hybrid**: Provide initial draft, let LLM refine and optimize

---
` + generateFooter(request.ProjectName, "Content Strategist") + `
`
	return utils.WriteFile(mainReadmePath, []byte(mainReadmeContent))
}

// generateContentFolderREADME generates README content for each content folder
func generateContentFolderREADME(folderType string, projectName string) string {
	descriptions := map[string]string{
		"app-pages": `# App Pages Content

Content for application and web pages including:
- Landing pages
- Feature pages
- About pages
- Product pages
- Service pages
- Contact pages

All content is SEO-optimized and ready for deployment.`,
		"legal": `# Legal Content

Legal documents and pages including:
- Privacy Policy
- Terms of Service
- Cookie Policy
- Disclaimer
- Refund Policy
- User Agreement

Note: Legal content should be reviewed by a legal professional before use.`,
		"social-media": `# Social Media Content

Social media posts and content for:
- Twitter/X
- LinkedIn
- Facebook
- Instagram
- TikTok
- Other platforms

Content includes captions, hashtags, and platform-specific formatting.`,
		"blog": `# Blog Content

Blog posts and articles including:
- How-to guides
- Industry insights
- Product updates
- Company news
- Educational content

All posts are SEO-optimized with proper meta descriptions and keywords.`,
		"documentation": `# Documentation Content

User documentation and guides including:
- User manuals
- API documentation
- Tutorials
- FAQs
- Getting started guides

Content is structured for easy navigation and search.`,
		"marketing": `# Marketing Content

Marketing copy and campaigns including:
- Ad copy
- Landing page content
- Promotional materials
- Sales pages
- Campaign messaging

All content is conversion-optimized.`,
		"email": `# Email Content

Email templates and campaigns including:
- Welcome emails
- Newsletter content
- Transactional emails
- Marketing emails
- Follow-up sequences

Templates are mobile-responsive and include plain text versions.`,
		"seo": `# SEO Content

SEO-optimized content including:
- Meta descriptions
- Alt text for images
- Keyword-optimized copy
- Schema markup suggestions
- Internal linking strategies

All content follows current SEO best practices.`,
	}

	desc, exists := descriptions[folderType]
	if !exists {
		desc = "Content for " + folderType
	}

	return desc + `

---
` + generateFooter(projectName, "Design & UX Manager") + `
`
}

// GeneratePlan is a convenience function that creates a PlanGenerator and generates plan structure
func GeneratePlan(request *models.ProjectRequest, projectPath string) error {
	generator := &PlanGenerator{}
	return generator.Generate(request, projectPath)
}
