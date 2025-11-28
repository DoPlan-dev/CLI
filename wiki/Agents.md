# Agents Documentation

Complete guide to the DoPlan CLI hierarchical AI agent system.

---

## Agent System Overview

DoPlan CLI uses a **hierarchical AI agency system** with 18 specialized agents. Each agent has a specific role, expertise, and responsibilities. Agents work together in a structured hierarchy to help you build your project.

### Key Features

- **18 Specialized Agents**: Each with unique expertise
- **Hierarchical Organization**: Clear reporting structure
- **Transparent Logic**: All agent definitions in markdown files
- **Customizable**: Modify agents to fit your needs
- **Collaborative**: Agents work together on tasks

---

## Agent Hierarchy

```
                    Project Orchestrator
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
   Product Manager   Engineering Lead   Design & UX Manager
        │                  │                  │
        │         ┌────────┼────────┐        │
        │         │        │        │        │
        │   System    Frontend  Backend   UI/UX
        │  Architect   Lead     Lead    Designer
        │                  │
        │            DevOps Engineer
        │                  │
        │         Security Lead
        │                  │
        │      Performance Engineer
        │
   QA & Reliability Manager
        │
     QA Engineer
        │
   Release & Growth Manager
        │
   ┌────┴────┐
   │         │
Release   Growth
Captain   Coach
        │
   Documentation Lead
        │
   Documentation Writer
```

---

## Individual Agent Roles

### Project Orchestrator

**Role**: CEO / Engineering Manager

**Responsibilities**:
- Strategic Vision: Define overall project vision, goals, and success metrics
- Resource Allocation: Allocate resources and prioritize work across all teams
- Decision Making: Make final decisions on architecture, features, and trade-offs
- Coordination: Ensure all teams are aligned
- Escalation: Handle escalations and resolve conflicts between teams
- Reporting: Report project status to stakeholders

**When Activated**: 
- `/tell` - Analyzes project idea
- `/good` - Approves and locks plan
- `/build` - Coordinates task execution
- `/finished` - Manages task completion

**File**: `.cursor/agents/project_orchestrator.md`

---

### Product Manager

**Role**: Product Strategy and Requirements

**Responsibilities**:
- Product Requirements: Define and manage product requirements
- Roadmap Planning: Create and maintain product roadmap
- Feature Prioritization: Prioritize features based on value
- User Research: Understand user needs and pain points
- Stakeholder Communication: Communicate with stakeholders

**When Activated**:
- `/tell` - Analyzes product idea
- `/improve` - Provides product insights
- `/write` - Creates PRD.md
- `/plan` - Generates product-related tasks

**File**: `.cursor/agents/product_manager.md`

---

### Engineering Lead

**Role**: Technical Leadership and Coordination

**Responsibilities**:
- Technical Strategy: Define technical direction
- Team Coordination: Coordinate frontend and backend teams
- Code Review: Review code quality and architecture
- Technical Decisions: Make technical decisions
- Resource Management: Manage engineering resources

**When Activated**:
- `/improve` - Provides technical insights
- `/write` - Creates ARCHITECTURE.md
- `/plan` - Generates technical tasks
- `/build` - Coordinates implementation

**File**: `.cursor/agents/engineering_lead.md`

---

### System Architect

**Role**: System Design and Architecture

**Responsibilities**:
- System Design: Design system architecture
- Technology Selection: Choose appropriate technologies
- Scalability: Ensure system scalability
- Performance: Design for performance
- Integration: Plan system integrations

**When Activated**:
- `/improve` - Provides architectural insights
- `/write` - Creates ARCHITECTURE.md
- `/build` - Designs system components

**File**: `.cursor/agents/system_architect.md`

---

### Frontend Lead

**Role**: Frontend Development Leadership

**Responsibilities**:
- Frontend Strategy: Define frontend architecture
- UI/UX Coordination: Work with design team
- Framework Selection: Choose frontend frameworks
- Performance: Optimize frontend performance
- Team Management: Lead frontend development

**When Activated**:
- `/build` - Implements frontend tasks
- Frontend-specific features

**File**: `.cursor/agents/frontend_lead.md`

---

### Backend Lead

**Role**: Backend Development Leadership

**Responsibilities**:
- Backend Strategy: Define backend architecture
- API Design: Design APIs and endpoints
- Database Design: Design database schemas
- Security: Ensure backend security
- Team Management: Lead backend development

**When Activated**:
- `/build` - Implements backend tasks
- Backend-specific features

**File**: `.cursor/agents/backend_lead.md`

---

### DevOps Engineer

**Role**: Infrastructure and Deployment

**Responsibilities**:
- Infrastructure: Set up and manage infrastructure
- CI/CD: Configure CI/CD pipelines
- Deployment: Manage deployments
- Monitoring: Set up monitoring and logging
- Automation: Automate operations

**When Activated**:
- `/build` - Implements DevOps tasks
- Infrastructure setup
- `/cheap` - Cost optimization

**File**: `.cursor/agents/devops_engineer.md`

---

### Security Lead

**Role**: Security and Compliance

**Responsibilities**:
- Security Strategy: Define security strategy
- Security Review: Review code for security issues
- Compliance: Ensure compliance with standards
- Vulnerability Management: Manage vulnerabilities
- Security Training: Provide security guidance

**When Activated**:
- `/safe` - Security audit
- `/build` - Security-related tasks

**File**: `.cursor/agents/security_lead.md`

---

### Performance Engineer

**Role**: Performance Optimization

**Responsibilities**:
- Performance Analysis: Analyze system performance
- Optimization: Optimize performance bottlenecks
- Monitoring: Monitor performance metrics
- Load Testing: Conduct load testing
- Recommendations: Provide performance recommendations

**When Activated**:
- `/build` - Performance-related tasks
- `/cheap` - Cost optimization

**File**: `.cursor/agents/performance_engineer.md`

---

### Design & UX Manager

**Role**: Design System Management

**Responsibilities**:
- Design Strategy: Define design strategy
- Design System: Create and maintain design system
- UX Guidelines: Establish UX guidelines
- Team Coordination: Coordinate design team
- Quality Assurance: Ensure design quality

**When Activated**:
- `/improve` - Provides design insights
- `/write` - Creates DESIGN_SYSTEM.md

**File**: `.cursor/agents/design_manager.md`

---

### UI/UX Designer

**Role**: User Interface and Experience Design

**Responsibilities**:
- UI Design: Create user interfaces
- UX Design: Design user experiences
- Prototyping: Create prototypes
- User Testing: Conduct user testing
- Design Implementation: Implement designs

**When Activated**:
- `/write` - Creates DESIGN_SYSTEM.md
- `/build` - Implements UI/UX tasks

**File**: `.cursor/agents/ui_ux_designer.md`

---

### QA & Reliability Manager

**Role**: Quality Assurance Management

**Responsibilities**:
- QA Strategy: Define QA strategy
- Test Planning: Plan testing approach
- Quality Standards: Establish quality standards
- Team Coordination: Coordinate QA team
- Quality Metrics: Track quality metrics

**When Activated**:
- `/improve` - Provides QA insights
- Quality assurance tasks

**File**: `.cursor/agents/qa_manager.md`

---

### QA Engineer

**Role**: Testing and Validation

**Responsibilities**:
- Test Execution: Execute tests
- Bug Reporting: Report and track bugs
- Test Automation: Automate tests
- Quality Validation: Validate quality
- Test Coverage: Ensure test coverage

**When Activated**:
- `/build` - Testing tasks
- Code review

**File**: `.cursor/agents/qa_engineer.md`

---

### Release & Growth Manager

**Role**: Release and Growth Management

**Responsibilities**:
- Release Strategy: Define release strategy
- Growth Strategy: Define growth strategy
- Release Planning: Plan releases
- Metrics Tracking: Track growth metrics
- Team Coordination: Coordinate release and growth teams

**When Activated**:
- `/improve` - Provides release and growth insights
- Release planning

**File**: `.cursor/agents/release_manager.md`

---

### Release Captain

**Role**: Release Coordination

**Responsibilities**:
- Release Execution: Execute releases
- Version Management: Manage versions
- Deployment: Coordinate deployments
- Release Notes: Create release notes
- Post-Release: Monitor post-release

**When Activated**:
- `/ship` - Release management
- `/finished` - Release coordination

**File**: `.cursor/agents/release_captain.md`

---

### Growth Coach

**Role**: User Growth and Engagement

**Responsibilities**:
- Growth Strategy: Define growth strategy
- User Acquisition: Plan user acquisition
- User Engagement: Improve user engagement
- Analytics: Analyze growth metrics
- Optimization: Optimize growth channels

**When Activated**:
- Growth-related tasks
- Analytics and optimization

**File**: `.cursor/agents/growth_coach.md`

---

### Documentation Lead

**Role**: Documentation Management

**Responsibilities**:
- Documentation Strategy: Define documentation strategy
- Documentation Standards: Establish standards
- Team Coordination: Coordinate documentation team
- Quality Assurance: Ensure documentation quality
- Maintenance: Maintain documentation

**When Activated**:
- `/improve` - Provides documentation insights
- Documentation tasks

**File**: `.cursor/agents/documentation_lead.md`

---

### Documentation Writer

**Role**: Technical Writing

**Responsibilities**:
- Writing: Write documentation
- Editing: Edit documentation
- Formatting: Format documentation
- Review: Review documentation
- Updates: Update documentation

**When Activated**:
- Documentation tasks
- Content creation

**File**: `.cursor/agents/documentation_writer.md`

---

### Content Strategist

**Role**: Brand Messaging & Content Systems

**Responsibilities**:
- Define tone of voice, messaging pillars, and slogans
- Audit existing copy/assets and identify gaps
- Provide structured briefs for feature/page copy
- Coordinate content approvals with legal and stakeholders

**When Activated**:
- `/improve` content & storytelling phase
- `/plan` when tasks require net-new copy
- Landing pages, marketing sites, onboarding flows

**File**: `.cursor/agents/content_strategist.md`

---

### SEO Specialist

**Role**: Search Optimization & Content Performance

**Responsibilities**:
- Build keyword strategy and map intents to pages
- Capture metadata/structured data requirements
- Define measurement plan for organic KPIs
- Review copy for SEO compliance and readability

**When Activated**:
- `/improve` SEO phase
- `/plan` when execution tasks include SEO workstreams
- Marketing/documentation updates that impact organic traffic

**File**: `.cursor/agents/seo_specialist.md`

---

### Marketing Manager

**Role**: Go-To-Market & Growth Strategy

**Responsibilities**:
- Map acquisition → activation → retention funnels
- Define campaign timelines, channels, and budgets
- Establish KPIs plus instrumentation plan
- Coordinate collateral with Content & SEO

**When Activated**:
- `/improve` marketing & funnel phase
- `/write` when GTM requirements influence PRD/ARCH
- `/plan` to ensure tasks align with GTM goals

**File**: `.cursor/agents/marketing_manager.md`

---

## How Agents Work Together

### Command-Agent Mapping

| Command | Primary Agents | Supporting Agents |
|---------|---------------|-------------------|
| `/tell` | Project Orchestrator, Product Manager | - |
| `/improve` | Product Manager, Engineering Lead, Design Manager, Marketing Manager | Content Strategist, SEO Specialist, QA Manager, Release Manager, Documentation Lead |
| `/write` | Product Manager, Engineering Lead, System Architect, Design Manager, UI/UX Designer | Project Orchestrator |
| `/change` | Project Orchestrator | Relevant agents based on document |
| `/good` | Project Orchestrator | - |
| `/plan` | Project Orchestrator, Engineering Lead, Product Manager | - |
| `/build` | Engineering Lead, Project Orchestrator | Frontend Lead, Backend Lead, etc. (based on task) |
| `/finished` | Project Orchestrator, Release Captain | - |
| `/ship` | Release Captain, Release Manager | - |
| `/safe` | Security Lead | - |
| `/cheap` | DevOps Engineer, Performance Engineer | - |

### Agent Collaboration Patterns

#### Pattern 1: Planning Collaboration

When using `/write`:
1. **Product Manager** creates PRD
2. **Engineering Lead** and **System Architect** create Architecture
3. **Design Manager** and **UI/UX Designer** create Design System
4. **Project Orchestrator** coordinates all

#### Pattern 2: Development Collaboration

When using `/build`:
1. **Engineering Lead** coordinates
2. **Frontend Lead** or **Backend Lead** implements (based on task)
3. **QA Engineer** reviews
4. **Security Lead** checks security (if needed)
5. **Project Orchestrator** monitors

#### Pattern 3: Brainstorming Collaboration

When using `/improve`:
1. Core panel:
   - Product Manager
   - Engineering Lead
   - Design Manager
   - Marketing Manager
2. Specialist breakouts:
   - Content Strategist (tone, messaging, assets)
   - SEO Specialist (keywords, metadata, performance)
   - QA Manager
   - Release Manager
   - Documentation Lead
3. Each provides insights from their expertise before the summary is confirmed

---

## Customizing Agents

### Modifying Agent Definitions

All agent definitions are in `.cursor/agents/` as markdown files. You can:

1. **Edit Agent Files**: Modify agent personas directly
2. **Add Custom Agents**: Create new agent files
3. **Remove Agents**: Delete agent files (not recommended)
4. **Change Hierarchy**: Update agent relationships

### Example: Customizing Product Manager

Edit `.cursor/agents/product_manager.md`:

```markdown
# Product Manager

## Role
Product Strategy and Requirements

## System Prompt
You are the Product Manager. Your focus is on [your custom focus].

## Responsibilities
- [Your custom responsibilities]
```

### Best Practices for Customization

1. **Keep Core Structure**: Maintain the basic structure
2. **Preserve Hierarchy**: Don't break the hierarchy
3. **Test Changes**: Test after modifications
4. **Document Changes**: Document your customizations
5. **Version Control**: Commit changes to git

---

## Agent Best Practices

### 1. Trust the Agents

Let agents do their work:
- Don't override agent decisions unnecessarily
- Review agent outputs before approving
- Use agents for their expertise

### 2. Use the Right Command

Match commands to agents:
- `/write` for planning (Product, Engineering, Design)
- `/build` for development (Engineering, Frontend, Backend)
- `/safe` for security (Security Lead)
- `/ship` for releases (Release Captain)

### 3. Review Agent Outputs

Always review:
- Planning documents before `/good`
- Code before `/finished`
- Releases before deploying

### 4. Leverage Agent Expertise

Use specialized agents:
- Security Lead for security reviews
- Performance Engineer for optimization
- QA Engineer for testing

### 5. Customize When Needed

Modify agents to fit your needs:
- Add domain-specific knowledge
- Adjust responsibilities
- Create custom agents

---

## Examples of Agent Interactions

### Example 1: Planning a Feature

```
User: /tell Add user authentication

1. Project Orchestrator: Analyzes request
2. Product Manager: Defines requirements
3. Engineering Lead: Plans technical approach
4. System Architect: Designs architecture
5. Security Lead: Reviews security implications
6. Design Manager: Plans UI/UX
```

### Example 2: Implementing a Feature

```
User: /build 1.2

1. Engineering Lead: Coordinates
2. Backend Lead: Implements API
3. Frontend Lead: Implements UI
4. QA Engineer: Tests
5. Security Lead: Reviews security
6. Project Orchestrator: Monitors progress
```

### Example 3: Releasing

```
User: /ship

1. Release Captain: Coordinates release
2. Release Manager: Plans release
3. QA Engineer: Final testing
4. DevOps Engineer: Deployment
5. Growth Coach: Growth strategy
```

---

## Related Pages

- [Workflow Guide](Workflow) - How agents work in the workflow
- [Commands Reference](Commands) - Commands that activate agents
- [Configuration Guide](Configuration) - Customizing agents
- [Rules Library](Rules) - Rules that guide agents
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

