# Templates

This directory contains template files used by the DoPlan CLI generator system. Templates are embedded into the binary and used to generate project documentation and structure.

## Structure

- `agents/` - Templates for rendering agent markdown files
- `commands/` - Templates for rendering command markdown files
- `documents/` - Templates for document generation
  - `brainstorm/` - Brainstorm meeting phase templates
  - `01_strategy/` - Strategy and planning templates
  - `02_architecture_design/` - Architecture and design templates
  - `03_delivery_execution/` - Delivery and execution templates
  - `04_quality_testing/` - Quality assurance and testing templates
  - `05_operations_support/` - Operations and support templates
  - `06_governance_compliance/` - Governance and compliance templates
  - `07_business_finance/` - Business and finance templates
  - `08_people_process/` - People and process templates

## Template Types

### Agent Templates
- `agents/agent.md.tmpl` - Template for rendering agent definitions to markdown

### Command Templates
- `commands/command.md.tmpl` - Template for rendering command definitions to markdown

### Document Templates

#### Brainstorm Templates
- `documents/brainstorm/phase-01-vision.md` - Phase 01: Vision & Outcomes
- `documents/brainstorm/phase-02-audience.md` - Phase 02: Audience & Differentiation
- `documents/brainstorm/phase-03-experience.md` - Phase 03: Experience, UI/UX & Tech
- `documents/brainstorm/phase-04-content.md` - Phase 04: Content & SEO
- `documents/brainstorm/phase-05-marketing.md` - Phase 05: Marketing & Growth
- `documents/brainstorm/phase-06-delivery.md` - Phase 06: Delivery, Ops & Risks
- `documents/brainstorm/CONFIRMATION_TEMPLATE.md` - Template for meeting confirmation display
- `documents/brainstorm/TEMPLATE_BRAINSTORM.md` - Template for final BRAINSTORM.md output

#### Strategy Templates (01_strategy/)
- `TEMPLATE_PRD.md` - Product Requirements Document template
- `TEMPLATE_ROADMAP.md` - Product roadmap template
- `TEMPLATE_KPI_TRACKER.md` - KPI tracking template
- `TEMPLATE_STAKEHOLDER_MAP.md` - Stakeholder mapping template
- `TEMPLATE_OPPORTUNITY_ASSESSMENT.md` - Opportunity assessment template
- `TEMPLATE_DISCOVERY_NOTES.md` - Discovery notes template

#### Architecture & Design Templates (02_architecture_design/)
- `TEMPLATE_SYSTEM_ARCHITECTURE.md` - System architecture template
- `TEMPLATE_DATA_MODEL.md` - Data model template
- `TEMPLATE_API_CONTRACT.yaml` - API contract template (YAML)
- `TEMPLATE_DESIGN_SYSTEM_SPEC.md` - Design system specification template
- `TEMPLATE_INTEGRATION_MAP.md` - Integration mapping template
- `TEMPLATE_INFRASTRUCTURE_CHECKLIST.md` - Infrastructure checklist template
- `TEMPLATE_ACCESSIBILITY_AUDIT.md` - Accessibility audit template

#### Delivery & Execution Templates (03_delivery_execution/)
- `TEMPLATE_SPRINT_PLAN.md` - Sprint planning template
- `TEMPLATE_RACI_MATRIX.md` - RACI matrix template
- `TEMPLATE_DEPENDENCY_MAP.md` - Dependency mapping template
- `TEMPLATE_RISK_REGISTER.md` - Risk register template
- `TEMPLATE_DECISION_LOG.md` - Decision log template
- `TEMPLATE_SERVICE_CHARTER.md` - Service charter template

#### Quality & Testing Templates (04_quality_testing/)
- `TEMPLATE_TEST_STRATEGY.md` - Test strategy template
- `TEMPLATE_TEST_MATRIX.md` - Test matrix template
- `TEMPLATE_AUTOMATION_PLAN.md` - Test automation plan template
- `TEMPLATE_NON_FUNCTIONAL_CHECKLIST.md` - Non-functional requirements checklist
- `TEMPLATE_RELEASE_READINESS.md` - Release readiness template

#### Operations & Support Templates (05_operations_support/)
- `TEMPLATE_RUNBOOK.md` - Runbook template
- `TEMPLATE_INCIDENT_RESPONSE.md` - Incident response template
- `TEMPLATE_POSTMORTEM.md` - Postmortem template
- `TEMPLATE_MONITORING_DASHBOARD.md` - Monitoring dashboard template
- `TEMPLATE_ROLLBACK_PLAN.md` - Rollback plan template
- `TEMPLATE_SUPPORT_HANDOFF.md` - Support handoff template

#### Governance & Compliance Templates (06_governance_compliance/)
- `TEMPLATE_SECURITY_REVIEW.md` - Security review template
- `TEMPLATE_PRIVACY_ASSESSMENT.md` - Privacy assessment template
- `TEMPLATE_COMPLIANCE_LOG.md` - Compliance log template
- `TEMPLATE_AUDIT_PREP.md` - Audit preparation template
- `TEMPLATE_VENDOR_ASSESSMENT.md` - Vendor assessment template

#### Business & Finance Templates (07_business_finance/)
- `TEMPLATE_BUDGET_TRACKER.md` - Budget tracking template
- `TEMPLATE_COST_BENEFIT.md` - Cost-benefit analysis template
- `TEMPLATE_CONTRACT_SUMMARY.md` - Contract summary template
- `TEMPLATE_PROCUREMENT_CHECKLIST.md` - Procurement checklist template

#### People & Process Templates (08_people_process/)
- `TEMPLATE_TEAM_CHARTER.md` - Team charter template
- `TEMPLATE_HIRING_BRIEF.md` - Hiring brief template
- `TEMPLATE_ONBOARDING_CHECKLIST.md` - Onboarding checklist template
- `TEMPLATE_TRAINING_PLAN.md` - Training plan template
- `TEMPLATE_MEETING_AGENDA.md` - Meeting agenda template

## Usage

Templates are loaded from embedded files and used by generator functions to create markdown files in generated projects. The templates are embedded at build time using Go's `embed` directive in `internal/content/content.go`.

### Embedding

Templates are embedded with the following pattern:
```go
//go:embed templates/agents/*.tmpl templates/commands/*.tmpl templates/documents/**/*.md
var embeddedTemplates embed.FS
```

This ensures all templates are included in the compiled binary and available at runtime without requiring external files.

