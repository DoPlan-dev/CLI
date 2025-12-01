# DoPlan Website & Documentation Project

## Project Overview

Build a modern, production-ready website and documentation site for DoPlan CLI using **Astro** and **HeroUI**. This project will create two interconnected sites:

1. **Main Website** (doplan.dev) - Marketing and product showcase
2. **Documentation Site** (docs.doplan.dev) - Comprehensive documentation portal

---

## Technology Stack

- **Framework**: Astro (latest stable version)
- **UI Components**: HeroUI (https://heroui.com)
- **Styling**: Tailwind CSS (via HeroUI)
- **Content**: Markdown with MDX support
- **Deployment**: Static site generation (SSG) ready for Vercel, Netlify, or GitHub Pages

---

## Project Structure

```
website/
├── src/
│   ├── components/          # Reusable components
│   │   ├── ui/             # HeroUI components
│   │   ├── layout/         # Layout components (Header, Footer, Nav)
│   │   └── sections/      # Page sections (Hero, Features, etc.)
│   ├── layouts/            # Astro layouts
│   │   ├── MainLayout.astro
│   │   └── DocsLayout.astro
│   ├── pages/              # Astro pages (file-based routing)
│   │   ├── index.astro     # Home page
│   │   ├── features.astro  # Features page
│   │   └── docs/           # Documentation pages
│   ├── content/            # Content collections
│   │   ├── docs/           # Documentation markdown files
│   │   └── config.ts       # Content collection config
│   ├── styles/             # Global styles
│   └── utils/              # Utility functions
├── public/                  # Static assets
│   ├── images/
│   └── favicon.ico
├── astro.config.mjs        # Astro configuration
├── package.json
└── tsconfig.json
```

---

## Main Website (doplan.dev)

### Navigation Structure

**Left Side:**
- Home
- Features

**Right Side:**
- GitHub (external link to https://github.com/DoPlan-dev/CLI)
- Docs (link to docs.doplan.dev)

### Pages

#### 1. Home Page (`/`)
- **Hero Section**
  - Large headline: "Idea to A Real Product" or "Zero-install AI Project Director"
  - Subheadline: "Bootstrap production-ready projects with a complete hierarchical AI agency system in seconds"
  - Interactive terminal component showing installation command: `npx @doplan-dev/cli`
  - Copy-to-clipboard functionality
  - CTA buttons: "Get Started" and "View Docs"

- **Key Features Grid**
  - Display 6-8 main features with icons:
    - ⚡ Zero-Install
    - 🚀 Lightning Fast (80-90% faster)
    - 🤖 18 AI Agents
    - 📚 1000+ Rules Library
    - 🎨 Interactive TUI
    - 🔌 IDE-Agnostic (6 IDEs)
    - 🧠 Memory & Brain
    - 🏆 Engagement System

- **Workflow Visualization**
  - Visual representation of: `/hey` → `/do` → `/plan` → `/dev` → `/sys`
  - Interactive cards showing each step with descriptions:
    - `/hey` - Welcome & Tutorial
    - `/do` - Capture Idea & Discovery Meeting
    - `/plan` - Generate Planning Documents & Tasks
    - `/dev` - Start Development (Auto-Detects Completion)
    - `/sys` - System Management
  - Show that `/dev` automatically detects task completion and handles commits/pushes

- **Command Overview**
  - Display the 5 core commands clearly:
    - `/hey` - Onboarding & Tutorial
    - `/do` - Idea Capture & Discovery
    - `/plan` - Generate Execution Plan
    - `/dev` - Start Development (with auto-completion detection)
    - `/sys` - System Control Panel

- **Stats/KPIs Section**
  - NPM Downloads
  - GitHub Stars
  - Version badge
  - Performance metrics

- **Footer**
  - Copyright: © 2025 DoPlan, Inc.
  - Links to GitHub, Docs, Releases
  - Social links (if applicable)

#### 2. Features Page (`/features`)
- Detailed feature descriptions
- Feature comparisons
- Use cases
- Screenshots/demos (if available)
- Feature categories:
  - Core Features
  - Engagement System
  - Personalization
  - Automation
  - Learning Support

### Design Requirements

- **Theme**: Dark theme by default with light mode toggle
- **Colors**:
  - Primary: Indigo (#6366f1)
  - Secondary: Purple (#8b5cf6)
  - Accent: Based on HeroUI theme
- **Typography**: Modern, clean fonts (HeroUI defaults or Poppins/JetBrains Mono)
- **Responsive**: Mobile-first, fully responsive
- **Animations**: Smooth, subtle animations using HeroUI/Framer Motion
- **Accessibility**: WCAG 2.1 AA compliant

---

## Documentation Site (docs.doplan.dev)

### Content Source

Primary content from:
- `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/wiki/` directory (9 main sections, 52+ markdown files)
- `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/docs/` directory (design docs, development guides, etc.)

### Documentation Structure

Based on existing wiki structure:

1. **Getting Started**
   - Welcome to DoPlan
   - Installation
   - First Project
   - Quick Tour

2. **Commands**
   - Command Overview (5 core commands: /hey, /do, /plan, /dev, /sys)
   - Workflow Commands (/hey, /do, /plan, /dev)
   - System Commands (/sys and subcommands)
   - Command Reference (complete syntax and examples)

3. **Engagement System**
   - Overview
   - Achievements (200+)
   - Challenges (30+)
   - Score System
   - Reward System
   - Engagement Dashboard

4. **Memory and Brain**
   - Overview
   - Memory Card
   - Brain System
   - Personalization
   - Relationship Building

5. **Workflow**
   - Complete Workflow
   - Phase-by-Phase Guide
   - Best Practices
   - Common Patterns

6. **Features**
   - Time Tracking
   - State Management
   - Git Automation
   - Learning Support
   - Personalization
   - Backup and Restore

7. **Learning & Education**
   - For Beginners
   - For Intermediate
   - For Advanced
   - Learning Goals
   - Tech Exploration

8. **Advanced Topics**
   - System Control
   - Customization
   - Integration
   - Troubleshooting
   - Contributing

9. **Reference**
   - Command Quick Reference
   - File Structure
   - State Management
   - Time Tracker Format
   - Memory Card Schema

### Documentation Features

- **Sidebar Navigation**: Hierarchical navigation with collapsible sections
- **Search**: Full-text search across all documentation
- **Code Blocks**: Syntax highlighting for code examples
- **Command Examples**: Interactive command examples
- **Table of Contents**: Auto-generated TOC for each page
- **Breadcrumbs**: Navigation breadcrumbs
- **Version Badge**: Show current CLI version
- **Edit on GitHub**: Link to edit each page on GitHub
- **Dark/Light Mode**: Theme toggle
- **Mobile Responsive**: Mobile-friendly navigation

### Documentation Layout

- **Header**: Same as main site (Home, Features, GitHub, Docs)
- **Sidebar**: Documentation navigation (sticky on scroll)
- **Content Area**: Markdown content with proper styling
- **Footer**: Links to related docs, GitHub, etc.

---

## Implementation Requirements

### Phase 1: Setup & Configuration
1. Initialize Astro project with TypeScript
2. Install and configure HeroUI
3. Set up Tailwind CSS
4. Configure content collections for docs
5. Set up project structure

### Phase 2: Main Website
1. Create MainLayout component
2. Build Header component with navigation
3. Build Footer component
4. Create Home page with all sections
5. Create Features page
6. Implement theme toggle (dark/light)
7. Add interactive terminal component
8. Implement copy-to-clipboard

### Phase 3: Documentation Site
1. Create DocsLayout component
2. Set up content collections from `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/wiki/` and `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/docs/`
3. Build sidebar navigation component
4. Create documentation pages with routing
5. Implement search functionality
6. Add table of contents generation
7. Style markdown content
8. Add code syntax highlighting

### Phase 4: Content Migration
1. Review and update wiki content from `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/wiki/`
2. Review and update docs content from `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/docs/`
3. Migrate markdown files to content collections
4. Ensure all links work correctly
5. Add frontmatter to all content files
6. Generate proper metadata

### Phase 5: Polish & Optimization
1. Add animations and transitions
2. Optimize images and assets
3. Implement SEO (meta tags, Open Graph)
4. Add analytics (if needed)
5. Test accessibility
6. Performance optimization
7. Mobile responsiveness testing

### Phase 6: Deployment Configuration
1. Configure build output
2. Set up deployment scripts
3. Configure domain routing (doplan.dev and docs.doplan.dev)
4. Add redirects if needed
5. Set up CI/CD for automatic deployments

---

## Key Components to Build

### Main Website Components
- `Header.astro` - Navigation header
- `Footer.astro` - Site footer
- `Hero.astro` - Hero section with terminal
- `FeatureCard.astro` - Feature display card
- `WorkflowVisualization.astro` - Workflow steps
- `StatsSection.astro` - KPI/stats display
- `Terminal.astro` - Interactive terminal component
- `ThemeToggle.astro` - Dark/light mode toggle

### Documentation Components
- `DocsHeader.astro` - Docs-specific header
- `DocsSidebar.astro` - Navigation sidebar
- `DocsContent.astro` - Content wrapper
- `TableOfContents.astro` - Auto-generated TOC
- `SearchBar.astro` - Documentation search
- `CodeBlock.astro` - Enhanced code blocks
- `CommandExample.astro` - Command examples
- `Breadcrumbs.astro` - Navigation breadcrumbs

---

## Content Requirements

### From Existing Sources

**Wiki Content** (`/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/wiki/`):
- 9 main sections
- 52+ markdown files
- Complete workflow documentation
- Command references
- Feature explanations

**Docs Content** (`/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/docs/`):
- Design documents
- Development guides
- API references
- Architecture docs

### Content Status

**✅ Content is ready for migration:**
- All wiki content has been cleaned and updated
- Only 5 commands referenced: `/hey`, `/do`, `/plan`, `/dev`, `/sys`
- All old command references removed
- Workflow documentation updated to reflect auto-completion detection
- Content is consistent and accurate

**Before migration:**
1. ✅ Wiki content reviewed and cleaned (completed)
2. ✅ All old commands removed (completed)
3. ✅ Workflow updated to new structure (completed)
4. Add proper frontmatter (title, description, date, etc.) to each markdown file
5. Fix any internal wiki links to match new structure
6. Ensure code examples use current commands only
7. Verify all command references are correct

---

## Design Guidelines

### Visual Style
- **Modern & Clean**: Minimalist design with focus on content
- **Professional**: Enterprise-ready appearance
- **Engaging**: Interactive elements that showcase DoPlan's capabilities
- **Consistent**: Unified design language across both sites

### HeroUI Integration
- Use HeroUI components for:
  - Buttons, Cards, Modals
  - Navigation components
  - Form elements
  - Icons
  - Theme system

### Branding
- **Logo**: DoPlan.dev branding
- **Colors**: Indigo/Purple theme
- **Typography**: Modern, readable fonts
- **Icons**: Consistent icon set (HeroUI icons or Font Awesome)

---

## Technical Requirements

### Performance
- Lighthouse score: 90+ on all metrics
- Fast page loads (< 2s)
- Optimized images (WebP, lazy loading)
- Code splitting
- Minimal JavaScript

### SEO
- Proper meta tags
- Open Graph tags
- Structured data (JSON-LD)
- Sitemap generation
- robots.txt

### Accessibility
- WCAG 2.1 AA compliance
- Keyboard navigation
- Screen reader support
- Proper ARIA labels
- Focus indicators

### Browser Support
- Modern browsers (Chrome, Firefox, Safari, Edge)
- Last 2 versions
- Mobile browsers

---

## Deployment

### Main Site (doplan.dev)
- Deploy to Vercel/Netlify
- Custom domain: doplan.dev
- HTTPS enabled
- CDN distribution

### Documentation Site (docs.doplan.dev)
- Deploy to Vercel/Netlify
- Custom domain: docs.doplan.dev
- HTTPS enabled
- CDN distribution

### Build Process
- Astro builds static HTML
- Optimized assets
- Pre-rendered pages
- Fast deployment

---

## Success Criteria

1. ✅ Both sites built with Astro + HeroUI
2. ✅ Main site navigation: Home, Features (left), GitHub, Docs (right)
3. ✅ Documentation site with full wiki content
4. ✅ Responsive design (mobile, tablet, desktop)
5. ✅ Dark/light theme toggle
6. ✅ Fast load times (< 2s)
7. ✅ SEO optimized
8. ✅ Accessible (WCAG 2.1 AA)
9. ✅ All content migrated and updated
10. ✅ Ready for production deployment

---

## Reference Materials

### Content Sources (Verified & Clean)

- **Wiki Content**: `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/wiki/` 
  - 9 main sections, 52+ markdown files
  - ✅ All cleaned - only references `/hey`, `/do`, `/plan`, `/dev`, `/sys`
  - Structure: 01-Getting-Started, 02-Commands, 03-Engagement-System, 04-Memory-and-Brain, 05-Workflow, 06-Features, 07-Learning-Education, 08-Advanced, 09-Reference

- **Documentation**: `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/docs/`
  - Design documents, development guides, architecture docs
  - ✅ Foundation docs cleaned and updated
  - Key files: `foundation/the-guide.md`, `foundation/prompt.md`, `WIKI_MINDMAP.md`

- **Existing Website Reference**: `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/website/` 
  - Reference design for visual style
  - HTML/CSS/JS structure to understand layout
  - Color scheme: Indigo (#6366f1), Purple (#8b5cf6)
  - Workflow: Hey → Do → Plan → Dev

- **Website Reference**: `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/test/docs/references/website/`
  - Example HTML structure
  - Design patterns and components
  - Terminal component reference

### Technical References

- **HeroUI Docs**: https://heroui.com
- **Astro Docs**: https://docs.astro.build
- **Astro Content Collections**: https://docs.astro.build/en/guides/content-collections/

---

## Important Notes

### Command System (Updated)
- **Only 5 commands exist**: `/hey`, `/do`, `/plan`, `/dev`, `/sys`
- `/dev` automatically detects task completion - no separate `/done` command
- Progress checking is done via `/sys status` (not `/status`)
- All old commands have been removed from documentation

### Workflow (Updated)
The workflow is now:
1. `/hey` - Welcome & Tutorial
2. `/do` - Capture Idea & Discovery Meeting
3. `/plan` - Generate Planning Documents & Tasks
4. `/dev` - Start Development (auto-detects completion, auto-commits/pushes)
5. `/sys` - System Management (status, performance, backup, etc.)

### Content Status
- ✅ Wiki content cleaned and verified
- ✅ Only current commands referenced
- ✅ Workflow documentation updated
- ✅ Ready for migration to Astro content collections

### Design Reference
- Use `/Users/Dorgham/Documents/Work/Devleopment/DoPlan/Go-Agents/doplan/GoPlan-CLI/test/docs/references/website/` as visual reference
- Maintain Indigo/Purple color scheme
- Keep workflow visualization: Hey → Do → Plan → Dev

---

**Project Status**: ✅ Ready to start - All content cleaned and verified
**Priority**: High
**Estimated Timeline**: 2-3 weeks for complete implementation

