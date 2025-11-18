package dpr

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CursorRulesGenerator generates design_rules.mdc for AI agents
type CursorRulesGenerator struct {
	projectRoot string
	data        *DPRData
}

// NewCursorRulesGenerator creates a new cursor rules generator
func NewCursorRulesGenerator(projectRoot string, data *DPRData) *CursorRulesGenerator {
	return &CursorRulesGenerator{
		projectRoot: projectRoot,
		data:        data,
	}
}

// Generate creates the design_rules.mdc file
func (crg *CursorRulesGenerator) Generate() error {
	rulesPath := filepath.Join(crg.projectRoot, ".doplan", "ai", "rules", "design_rules.mdc")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(rulesPath), 0755); err != nil {
		return fmt.Errorf("failed to create ai/rules directory: %w", err)
	}

	content := crg.generateRulesContent()

	return os.WriteFile(rulesPath, []byte(content), 0644)
}

func (crg *CursorRulesGenerator) generateRulesContent() string {
	var content strings.Builder

	// Header
	content.WriteString("# Design Rules\n\n")
	content.WriteString("This file defines the design system rules that all AI agents MUST follow when working on this project.\n\n")
	content.WriteString("**Source:** Generated from DPR (Design Preferences & Requirements) questionnaire\n\n")
	content.WriteString("---\n\n")

	// Color Usage Rules
	content.WriteString("## Color Usage Rules\n\n")
	content.WriteString(crg.generateColorRules())
	content.WriteString("\n\n")

	// Typography Rules
	content.WriteString("## Typography Rules\n\n")
	content.WriteString(crg.generateTypographyRules())
	content.WriteString("\n\n")

	// Spacing Rules
	content.WriteString("## Spacing Rules\n\n")
	content.WriteString(crg.generateSpacingRules())
	content.WriteString("\n\n")

	// Component Guidelines
	content.WriteString("## Component Guidelines\n\n")
	content.WriteString(crg.generateComponentGuidelines())
	content.WriteString("\n\n")

	// Responsive Rules
	content.WriteString("## Responsive Rules\n\n")
	content.WriteString(crg.generateResponsiveRules())
	content.WriteString("\n\n")

	// Accessibility Requirements
	content.WriteString("## Accessibility Requirements\n\n")
	content.WriteString(crg.generateAccessibilityRules())
	content.WriteString("\n\n")

	// Code Style
	content.WriteString("## Code Style\n\n")
	content.WriteString(crg.generateCodeStyleRules())
	content.WriteString("\n\n")

	return content.String()
}

func (crg *CursorRulesGenerator) generateColorRules() string {
	primary := crg.getAnswerString("color_primary", "#667eea")
	secondary := crg.getAnswerString("color_secondary", "#764ba2")

	return fmt.Sprintf(`### Primary Colors
- **Primary Color:** %s - Use for primary actions, links, and brand elements
- **Secondary Color:** %s - Use for secondary actions and accents
- **DO NOT** use colors outside the design system
- **DO NOT** create new colors without updating the design tokens

### Color Usage
- Use semantic colors (success, warning, error, info) from design tokens
- Ensure sufficient contrast ratios (WCAG AA minimum)
- Test colors in both light and dark modes if applicable
- Use neutral colors for backgrounds and text`, primary, secondary)
}

func (crg *CursorRulesGenerator) generateTypographyRules() string {
	style := crg.getAnswerString("typography_style", "Sans-serif")
	importance := crg.getAnswerInt("typography_importance", 3)

	rules := fmt.Sprintf(`### Typography System
- **Font Style:** %s
- **Importance Level:** %d/5

### Type Scale
Use the following font sizes from design tokens:
- xs: 0.75rem (12px)
- sm: 0.875rem (14px)
- base: 1rem (16px)
- lg: 1.125rem (18px)
- xl: 1.25rem (20px)
- 2xl: 1.5rem (24px)
- 3xl: 1.875rem (30px)
- 4xl: 2.25rem (36px)
- 5xl: 3rem (48px)

### Line Heights
- Use line-height values from design tokens
- Headings: tight (1.25)
- Body text: normal (1.5)
- Long-form content: relaxed (1.625)

### Font Weights
- Use font weights from design tokens
- Headings: semibold (600) or bold (700)
- Body text: normal (400) or medium (500)`, style, importance)

	return rules
}

func (crg *CursorRulesGenerator) generateSpacingRules() string {
	spacing := crg.getAnswerString("layout_spacing", "Moderate")

	return fmt.Sprintf(`### Spacing System
- **Spacing Preference:** %s
- **Base Unit:** Use spacing values from design-tokens.json
- **DO NOT** use arbitrary spacing values
- **DO NOT** use magic numbers (e.g., margin: 17px)

### Spacing Guidelines
- Use spacing scale: 0, 1, 2, 3, 4, 5, 6, 8, 10, 12, 16, 20, 24, 32, 40, 48, 56, 64
- Maintain consistent spacing rhythm
- Use larger spacing for section separation
- Use smaller spacing for related elements

### Tailwind Utilities
When using Tailwind CSS, use spacing utilities:
- p-4, m-4, gap-4, space-y-4, etc.
- DO NOT use arbitrary values like p-[17px]`, spacing)
}

func (crg *CursorRulesGenerator) generateComponentGuidelines() string {
	style := crg.getAnswerString("components_style", "Elevated")
	interactivity := crg.getAnswerInt("components_interactivity", 3)

	return fmt.Sprintf(`### Component Style
- **Style:** %s
- **Interactivity Level:** %d/5

### Component Rules
- All components MUST follow the design system
- Use design tokens for colors, spacing, typography
- Maintain consistent styling across components
- Ensure all components have proper states:
  - Default
  - Hover
  - Active/Focus
  - Disabled
  - Error (if applicable)

### Component Structure
- Use consistent border radius from tokens
- Apply shadows according to component style (%s)
- Ensure proper focus indicators for accessibility
- Use semantic HTML elements`, style, interactivity, style)
}

func (crg *CursorRulesGenerator) generateResponsiveRules() string {
	priority := crg.getAnswerString("responsive_priority", "Mobile, Desktop")
	approach := crg.getAnswerString("responsive_approach", "Mobile-first")

	return fmt.Sprintf(`### Responsive Strategy
- **Device Priority:** %s
- **Approach:** %s

### Breakpoints
Use breakpoints from design tokens:
- Mobile: 320px
- Tablet: 768px
- Desktop: 1024px
- Large: 1440px

### Responsive Guidelines
- Design %s
- Test on all priority devices
- Use fluid layouts where possible
- Ensure touch-friendly interactions on mobile
- Maintain usability across all screen sizes

### Tailwind Breakpoints
When using Tailwind CSS:
- sm: 640px (use sparingly)
- md: 768px (tablet)
- lg: 1024px (desktop)
- xl: 1280px
- 2xl: 1536px (large)`, priority, approach, approach)
}

func (crg *CursorRulesGenerator) generateAccessibilityRules() string {
	importance := crg.getAnswerInt("accessibility_importance", 4)
	requirements := crg.getAnswerString("accessibility_requirements", "")

	rules := fmt.Sprintf(`### Accessibility Priority
- **Importance Level:** %d/5

### WCAG Compliance
- **Minimum:** WCAG 2.1 AA compliance
- **Target:** WCAG 2.1 AAA where possible

### Accessibility Requirements
- **Keyboard Navigation:** All interactive elements must be keyboard accessible
- **Screen Readers:** Use semantic HTML and ARIA attributes appropriately
- **Color Contrast:** Minimum 4.5:1 for normal text, 3:1 for large text
- **Focus Indicators:** Visible focus indicators on all interactive elements
- **Alt Text:** Provide meaningful alt text for images
- **Form Labels:** All form inputs must have associated labels
- **Error Messages:** Clear, accessible error messages
- **Skip Links:** Provide skip navigation links for keyboard users
- **ARIA Labels:** Use ARIA labels for complex interactions
- **Reduced Motion:** Respect prefers-reduced-motion media query

### Semantic HTML
- Use proper heading hierarchy (h1, h2, h3, etc.)
- Use semantic elements (nav, main, article, section, etc.)
- Use form elements correctly (input, label, fieldset, legend)
- Use button for actions, not div or span

### Testing
- Test with keyboard navigation (Tab, Enter, Space, Arrow keys)
- Test with screen readers (NVDA, JAWS, VoiceOver)
- Verify color contrast ratios using tools
- Test focus indicators visibility
- Test with zoom up to 200%%
- Verify all content is accessible without mouse`, importance)

	if requirements != "" && requirements != "nil" {
		rules += fmt.Sprintf("\n\n### Specific Requirements\n%s", requirements)
	}

	return rules
}

func (crg *CursorRulesGenerator) generateCodeStyleRules() string {
	return `### Tailwind CSS Utilities
When using Tailwind CSS, follow these patterns:

#### Colors
- Use color utilities from design tokens: text-primary, bg-secondary
- Use semantic colors: text-success, bg-error, etc.
- DO NOT use arbitrary colors: text-[#ff0000]

#### Spacing
- Use spacing scale: p-4, m-6, gap-8
- DO NOT use arbitrary spacing: p-[17px]

#### Typography
- Use font size utilities: text-sm, text-lg, text-2xl
- Use font weight utilities: font-medium, font-bold
- Use line height utilities: leading-tight, leading-normal

#### Components
- Use border radius from tokens: rounded-md, rounded-lg
- Use shadows: shadow-sm, shadow-md, shadow-lg
- Use consistent component patterns

### Code Organization
- Group related styles together
- Use consistent naming conventions
- Comment complex styling decisions
- Reference design tokens in comments when helpful`
}

func (crg *CursorRulesGenerator) getAnswerString(key string, defaultValue string) string {
	if val, ok := crg.data.Answers[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", val)
	}
	return defaultValue
}

func (crg *CursorRulesGenerator) getAnswerInt(key string, defaultValue int) int {
	if val, ok := crg.data.Answers[key]; ok {
		if num, ok := val.(int); ok {
			return num
		}
		if str, ok := val.(string); ok {
			if num, err := strconv.Atoi(str); err == nil {
				return num
			}
		}
	}
	return defaultValue
}
