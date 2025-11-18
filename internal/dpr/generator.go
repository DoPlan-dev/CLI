package dpr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Generator generates DPR.md document from questionnaire data
type Generator struct {
	projectRoot string
	data        *DPRData
}

// NewGenerator creates a new DPR generator
func NewGenerator(projectRoot string, data *DPRData) *Generator {
	return &Generator{
		projectRoot: projectRoot,
		data:        data,
	}
}

// Generate creates the DPR.md file
func (g *Generator) Generate() error {
	dprPath := filepath.Join(g.projectRoot, "doplan", "design", "DPR.md")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(dprPath), 0755); err != nil {
		return fmt.Errorf("failed to create design directory: %w", err)
	}

	content := g.generateDPRContent()

	return os.WriteFile(dprPath, []byte(content), 0644)
}

func (g *Generator) generateDPRContent() string {
	var content strings.Builder

	// Header
	content.WriteString("# Design Preferences & Requirements (DPR)\n\n")
	content.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("January 2, 2006")))
	content.WriteString("---\n\n")

	// Executive Summary
	content.WriteString("## Executive Summary\n\n")
	content.WriteString(g.generateExecutiveSummary())
	content.WriteString("\n\n")

	// Audience Analysis
	content.WriteString("## Audience Analysis\n\n")
	content.WriteString(g.generateAudienceAnalysis())
	content.WriteString("\n\n")

	// Design Principles
	content.WriteString("## Design Principles\n\n")
	content.WriteString(g.generateDesignPrinciples())
	content.WriteString("\n\n")

	// Visual Identity
	content.WriteString("## Visual Identity\n\n")
	content.WriteString(g.generateVisualIdentity())
	content.WriteString("\n\n")

	// Layout Guidelines
	content.WriteString("## Layout Guidelines\n\n")
	content.WriteString(g.generateLayoutGuidelines())
	content.WriteString("\n\n")

	// Component Library
	content.WriteString("## Component Library\n\n")
	content.WriteString(g.generateComponentLibrary())
	content.WriteString("\n\n")

	// Animation Guidelines
	content.WriteString("## Animation Guidelines\n\n")
	content.WriteString(g.generateAnimationGuidelines())
	content.WriteString("\n\n")

	// Accessibility
	content.WriteString("## Accessibility\n\n")
	content.WriteString(g.generateAccessibility())
	content.WriteString("\n\n")

	// Responsive Design
	content.WriteString("## Responsive Design\n\n")
	content.WriteString(g.generateResponsiveDesign())
	content.WriteString("\n\n")

	// Wireframes Section
	content.WriteString("## Wireframes & Mockups\n\n")
	content.WriteString(g.generateWireframes())
	content.WriteString("\n\n")

	// Implementation Guidelines
	content.WriteString("## Implementation Guidelines\n\n")
	content.WriteString(g.generateImplementationGuidelines())
	content.WriteString("\n\n")

	return content.String()
}

func (g *Generator) generateExecutiveSummary() string {
	style := g.getAnswer("style_overall", "Modern")
	emotion := g.getAnswer("emotion_primary", "Professional")
	
	return fmt.Sprintf(`This design system is built around a **%s** aesthetic with a focus on evoking **%s** emotions. The design prioritizes clarity, usability, and consistency across all interfaces.`, style, emotion)
}

func (g *Generator) generateAudienceAnalysis() string {
	primary := g.getAnswer("audience_primary", "Users")
	age := g.getAnswer("audience_age", "26-35")
	tech := g.getAnswer("audience_tech_level", "Intermediate")
	
	return fmt.Sprintf(`### Primary Audience
- **Target Group:** %s
- **Age Range:** %s
- **Technical Level:** %s

### User Needs
Based on the target audience, the design should prioritize:
- Clear information hierarchy
- Intuitive navigation
- Accessible interactions
- Consistent visual language`, primary, age, tech)
}

func (g *Generator) generateDesignPrinciples() string {
	emotion := g.getAnswer("emotion_primary", "Professional")
	style := g.getAnswer("style_overall", "Modern")
	
	return fmt.Sprintf(`1. **Emotional Resonance:** Design should evoke %s feelings
2. **Visual Consistency:** Maintain %s aesthetic throughout
3. **User-Centered:** Prioritize user needs and accessibility
4. **Clarity:** Clear information hierarchy and visual communication
5. **Efficiency:** Streamlined interactions and workflows`, emotion, style)
}

func (g *Generator) generateVisualIdentity() string {
	primary := g.getAnswer("color_primary", "#667eea")
	secondary := g.getAnswer("color_secondary", "#764ba2")
	scheme := g.getAnswer("color_scheme", "Both")
	
	return fmt.Sprintf(`### Color Palette
- **Primary Color:** %s
- **Secondary Color:** %s
- **Color Scheme:** %s

### Typography
- **Style:** %s
- **Importance Level:** %v/5

### Visual Style
- **Overall Style:** %s
- **Component Style:** %s`, 
		primary, secondary, scheme,
		g.getAnswer("typography_style", "Sans-serif"),
		g.getAnswer("typography_importance", 3),
		g.getAnswer("style_overall", "Modern"),
		g.getAnswer("components_style", "Elevated"))
}

func (g *Generator) generateLayoutGuidelines() string {
	layout := g.getAnswer("layout_style", "Centered")
	spacing := g.getAnswer("layout_spacing", "Moderate")
	
	return fmt.Sprintf(`### Layout Structure
- **Layout Style:** %s
- **Spacing Preference:** %s

### Grid System
- Use consistent spacing units
- Maintain visual rhythm
- Ensure responsive breakpoints`, layout, spacing)
}

func (g *Generator) generateComponentLibrary() string {
	style := g.getAnswer("components_style", "Elevated")
	interactivity := g.getAnswer("components_interactivity", 3)
	
	return fmt.Sprintf(`### Component Style
- **Style:** %s
- **Interactivity Level:** %v/5

### Component Guidelines
- Consistent styling across all components
- Clear states (default, hover, active, disabled)
- Accessible interactions
- Responsive behavior`, style, interactivity)
}

func (g *Generator) generateAnimationGuidelines() string {
	level := g.getAnswer("animation_level", "Subtle")
	style := g.getAnswer("animation_style", "Smooth")
	
	return fmt.Sprintf(`### Animation Principles
- **Level:** %s
- **Style:** %s

### Usage Guidelines
- Use animations to enhance understanding
- Keep animations subtle and purposeful
- Respect user preferences (reduced motion)
- Ensure performance optimization`, level, style)
}

func (g *Generator) generateAccessibility() string {
	importance := g.getAnswer("accessibility_importance", 4)
	requirements := g.getAnswer("accessibility_requirements", "")
	
	content := fmt.Sprintf(`### Accessibility Priority
- **Importance Level:** %v/5

### Requirements
- WCAG 2.1 AA compliance minimum
- Keyboard navigation support
- Screen reader compatibility
- Color contrast ratios
- Focus indicators`, importance)
	
	if requirements != "" && requirements != nil {
		content += fmt.Sprintf("\n- **Specific Requirements:** %v", requirements)
	}
	
	return content
}

func (g *Generator) generateResponsiveDesign() string {
	priority := g.getAnswer("responsive_priority", "Mobile, Desktop")
	approach := g.getAnswerString("responsive_approach", "Mobile-first")
	
	return fmt.Sprintf(`### Device Priority
- **Priority Devices:** %v

### Responsive Strategy
- **Approach:** %s
- Breakpoints: Mobile (320px), Tablet (768px), Desktop (1024px), Large (1440px)
- Fluid layouts where possible
- Touch-friendly interactions on mobile

### Breakpoint Guidelines
- **Mobile (< 768px):** Optimize for touch, single column layouts, simplified navigation
- **Tablet (768px - 1024px):** Two-column layouts, enhanced interactions
- **Desktop (1024px+):** Full feature set, multi-column layouts, hover states
- **Large (1440px+):** Maximum content width, enhanced spacing`, priority, approach)
}

func (g *Generator) generateWireframes() string {
	return `### Wireframe Guidelines
- Create wireframes for all major screens and user flows
- Focus on layout structure and information hierarchy
- Include annotations for interactions and behaviors
- Consider responsive breakpoints in wireframes

### Mockup Requirements
- High-fidelity mockups for key screens
- Show visual design and branding
- Include all states (default, hover, active, error)
- Provide design specifications for developers

### Design Tools
- Use design tools like Figma, Sketch, or Adobe XD
- Maintain design system components
- Keep designs in sync with code implementation
- Document design decisions and rationale`
}

func (g *Generator) generateImplementationGuidelines() string {
	return fmt.Sprintf(`### Development Workflow
1. **Design Review:** Review DPR.md and design tokens before implementation
2. **Component Development:** Build components following design system
3. **Design Token Usage:** Use design-tokens.json for all styling values
4. **AI Agent Rules:** Follow .doplan/ai/rules/design_rules.mdc for AI-generated code
5. **Testing:** Verify implementation matches design specifications

### Code Standards
- Use design tokens for all colors, spacing, typography
- Follow component guidelines from DPR
- Maintain consistency across all implementations
- Document any deviations from design system

### Design System Maintenance
- Update design tokens as design evolves
- Keep DPR.md current with project changes
- Review and update design rules regularly
- Ensure all team members follow design system`)
}

func (g *Generator) getAnswer(key string, defaultValue interface{}) interface{} {
	if val, ok := g.data.Answers[key]; ok {
		return val
	}
	return defaultValue
}

func (g *Generator) getAnswerString(key string, defaultValue string) string {
	if val, ok := g.data.Answers[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", val)
	}
	return defaultValue
}

