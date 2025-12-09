# Design System

## Design Principles
1. **Clarity** - Interface should be clear and easy to understand
2. **Consistency** - Design patterns should be consistent across the application
3. **Accessibility** - Design should be accessible to all users
4. **Performance** - Design should not compromise application performance
5. **Scalability** - Design system should scale with the product

## Color Palette

### Primary Colors
| Token | Hex | Usage | RGB |
| --- | --- | --- | --- |
| --color-primary | #3B82F6 | Primary actions, CTAs, links | rgb(59, 130, 246) |
| --color-primary-dark | #2563EB | Hover states, active states | rgb(37, 99, 235) |
| --color-primary-light | #60A5FA | Light backgrounds, highlights | rgb(96, 165, 250) |

### Secondary Colors
| Token | Hex | Usage | RGB |
| --- | --- | --- | --- |
| --color-secondary | #8B5CF6 | Secondary actions, accents | rgb(139, 92, 246) |
| --color-secondary-dark | #7C3AED | Hover states | rgb(124, 58, 237) |
| --color-secondary-light | #A78BFA | Light backgrounds | rgb(167, 139, 250) |

### Neutral Colors
| Token | Hex | Usage | RGB |
| --- | --- | --- | --- |
| --color-bg | #FFFFFF | Primary background | rgb(255, 255, 255) |
| --color-bg-secondary | #F9FAFB | Secondary background | rgb(249, 250, 251) |
| --color-text | #111827 | Primary text | rgb(17, 24, 39) |
| --color-text-secondary | #6B7280 | Secondary text | rgb(107, 114, 128) |
| --color-border | #E5E7EB | Borders, dividers | rgb(229, 231, 235) |

### Semantic Colors
| Token | Hex | Usage | RGB |
| --- | --- | --- | --- |
| --color-success | #22C55E | Success messages, positive actions | rgb(34, 197, 94) |
| --color-warning | #FACC15 | Warning messages, caution | rgb(250, 204, 21) |
| --color-error | #F87171 | Error messages, destructive actions | rgb(248, 113, 113) |
| --color-info | #38BDF8 | Informational messages | rgb(56, 189, 248) |

### Dark Mode Support (Optional)
| Token | Hex | Usage |
| --- | --- | --- |
| --color-bg-dark | #1F2937 | Dark mode background |
| --color-text-dark | #F9FAFB | Dark mode text |
| --color-border-dark | #374151 | Dark mode borders |

## Typography

### Font Family
- **Primary Font**: System font stack (San Francisco, Segoe UI, Roboto, sans-serif)
- **Monospace Font**: 'Courier New', Courier, monospace (for code)

### Type Scale
| Role | Font Size | Line Height | Weight | Usage |
| --- | --- | --- | --- | --- |
| H1 | 48px (3rem) | 1.2 | 700 | Page titles, hero headings |
| H2 | 36px (2.25rem) | 1.3 | 600 | Section headings |
| H3 | 30px (1.875rem) | 1.4 | 600 | Subsection headings |
| H4 | 24px (1.5rem) | 1.5 | 600 | Card titles, small headings |
| H5 | 20px (1.25rem) | 1.5 | 600 | Small headings |
| H6 | 18px (1.125rem) | 1.5 | 600 | Smallest headings |
| Body Large | 18px (1.125rem) | 1.6 | 400 | Large body text |
| Body | 16px (1rem) | 1.6 | 400 | Standard body text |
| Body Small | 14px (0.875rem) | 1.5 | 400 | Small body text, captions |
| Label | 14px (0.875rem) | 1.4 | 500 | Form labels |
| Button | 16px (1rem) | 1.5 | 500 | Button text |
| Caption | 12px (0.75rem) | 1.4 | 400 | Captions, fine print |

## Layout & Spacing

### Spacing Scale
- **Base unit**: 4px
- **Scale**: 4, 8, 12, 16, 20, 24, 32, 40, 48, 64, 80, 96, 128

### Spacing Tokens
| Token | Value | Usage |
| --- | --- | --- |
| --spacing-xs | 4px | Tight spacing, icon padding |
| --spacing-sm | 8px | Small gaps, compact layouts |
| --spacing-md | 16px | Standard spacing |
| --spacing-lg | 24px | Section spacing, card padding |
| --spacing-xl | 32px | Large section spacing |
| --spacing-2xl | 48px | Extra large spacing |
| --spacing-3xl | 64px | Hero section spacing |

### Layout Guidelines
- **Container Max Width**: 1280px for desktop
- **Section Padding**: 32px (mobile), 48px (desktop)
- **Card Padding**: 24px
- **Grid Gap**: 24px (desktop), 16px (mobile)
- **Border Radius**: 8px (standard), 12px (cards), 4px (inputs)

## Components

### Buttons

#### Primary Button
- **Background**: --color-primary
- **Text Color**: White
- **Padding**: 12px 24px
- **Border Radius**: 8px
- **Font Weight**: 500
- **Hover**: --color-primary-dark
- **Disabled**: 50% opacity
- **States**: Default, Hover, Active, Disabled, Loading

#### Secondary Button
- **Background**: Transparent
- **Text Color**: --color-primary
- **Border**: 1px solid --color-primary
- **Padding**: 12px 24px
- **Border Radius**: 8px
- **Hover**: --color-primary-light background
- **States**: Default, Hover, Active, Disabled

#### Text Button
- **Background**: Transparent
- **Text Color**: --color-primary
- **Padding**: 8px 16px
- **Hover**: --color-bg-secondary background
- **States**: Default, Hover, Active

### Form Elements

#### Input Field
- **Height**: 44px
- **Padding**: 12px 16px
- **Border**: 1px solid --color-border
- **Border Radius**: 8px
- **Font Size**: 16px
- **Focus**: 2px solid --color-primary outline
- **Error State**: Red border, error message below
- **Disabled**: Gray background, reduced opacity

#### Select Dropdown
- **Height**: 44px
- **Padding**: 12px 16px
- **Border**: 1px solid --color-border
- **Border Radius**: 8px
- **Background**: White
- **Focus**: 2px solid --color-primary outline

#### Checkbox
- **Size**: 20px × 20px
- **Border**: 2px solid --color-border
- **Border Radius**: 4px
- **Checked**: --color-primary background
- **Focus**: Outline ring

#### Radio Button
- **Size**: 20px × 20px
- **Border**: 2px solid --color-border
- **Checked**: --color-primary center dot
- **Focus**: Outline ring

### Cards
- **Background**: White
- **Padding**: 24px
- **Border Radius**: 12px
- **Shadow**: 0 1px 3px rgba(0, 0, 0, 0.1)
- **Hover**: Elevated shadow (optional)

### Navigation
- **Height**: 64px (desktop), 56px (mobile)
- **Background**: White or --color-bg-secondary
- **Border**: 1px solid --color-border (bottom)
- **Padding**: 0 24px
- **Link Spacing**: 24px

### Modals/Dialogs
- **Background**: White
- **Border Radius**: 12px
- **Padding**: 32px
- **Max Width**: 500px (standard), 800px (large)
- **Overlay**: rgba(0, 0, 0, 0.5)
- **Shadow**: Large shadow for depth

## Motion & Animation

### Transitions
- **Standard**: 200ms ease-out
- **Fast**: 150ms ease-out (hover states)
- **Slow**: 300ms ease-out (page transitions)

### Animation Guidelines
- **Micro-interactions**: Subtle animations for feedback
- **Page Transitions**: Smooth fade or slide transitions
- **Loading States**: Skeleton screens or spinners
- **Error States**: Gentle shake or highlight
- **Success States**: Checkmark animation or toast notification

### Animation Tokens
| Token | Value | Usage |
| --- | --- | --- |
| --duration-fast | 150ms | Hover effects |
| --duration-normal | 200ms | Standard transitions |
| --duration-slow | 300ms | Page transitions |
| --easing-standard | ease-out | Most animations |
| --easing-bounce | cubic-bezier(0.68, -0.55, 0.265, 1.55) | Playful animations |

## Responsive Design

### Breakpoints
| Breakpoint | Width | Usage |
| --- | --- | --- |
| Mobile | < 640px | Small phones |
| Tablet | 640px - 1024px | Tablets, large phones |
| Desktop | > 1024px | Desktop, laptops |
| Large Desktop | > 1280px | Large screens |

### Responsive Guidelines
- **Mobile First**: Design for mobile, enhance for larger screens
- **Touch Targets**: Minimum 44px × 44px for touch interactions
- **Typography**: Scale down on mobile (reduce by 2-4px)
- **Spacing**: Reduce padding/margins on mobile
- **Navigation**: Collapsible menu on mobile

## Accessibility

### Color Contrast
- **Text on Background**: Minimum 4.5:1 ratio (WCAG AA)
- **Large Text**: Minimum 3:1 ratio
- **Interactive Elements**: Clear focus states

### Keyboard Navigation
- **Tab Order**: Logical tab sequence
- **Focus Indicators**: Visible focus rings (2px solid outline)
- **Skip Links**: Skip to main content links
- **Keyboard Shortcuts**: Documented shortcuts

### Screen Reader Support
- **Semantic HTML**: Use proper HTML elements
- **ARIA Labels**: Descriptive labels for interactive elements
- **Alt Text**: Descriptive alt text for images
- **Landmarks**: Proper use of ARIA landmarks

### Other Accessibility Features
- **Error Messages**: Clear, actionable error messages
- **Form Labels**: All inputs have associated labels
- **Loading States**: Announce loading states to screen readers
- **Color Independence**: Don't rely solely on color to convey information

## Design Tokens (CSS Variables)

```css
:root {
  /* Colors */
  --color-primary: #3B82F6;
  --color-primary-dark: #2563EB;
  --color-primary-light: #60A5FA;
  --color-secondary: #8B5CF6;
  --color-bg: #FFFFFF;
  --color-text: #111827;
  --color-text-secondary: #6B7280;
  --color-border: #E5E7EB;
  --color-success: #22C55E;
  --color-warning: #FACC15;
  --color-error: #F87171;
  --color-info: #38BDF8;

  /* Spacing */
  --spacing-xs: 4px;
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --spacing-lg: 24px;
  --spacing-xl: 32px;
  --spacing-2xl: 48px;
  --spacing-3xl: 64px;

  /* Typography */
  --font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  --font-size-base: 16px;
  --line-height-base: 1.6;

  /* Transitions */
  --duration-fast: 150ms;
  --duration-normal: 200ms;
  --duration-slow: 300ms;
  --easing-standard: ease-out;
}
```

## Component Library Structure

### Recommended Organization
```
components/
├── atoms/          # Basic building blocks (Button, Input)
├── molecules/      # Simple combinations (Form, Card)
├── organisms/      # Complex components (Header, Footer)
└── templates/      # Page layouts
```

---

**Generated by**: DoPlan CLI vlatest  
**Sup-Agent**: Design & UX Manager  
**Date**: 2025-01-27  
**Status**: Draft - Ready for review and customization
