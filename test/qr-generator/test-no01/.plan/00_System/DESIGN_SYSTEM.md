# Design System
## QR Code Generator API - Micro SaaS

**Version**: 1.0  
**Date**: 2024-12-19  
**Authors**: Design & UX Manager, UI/UX Designer  
**Status**: Draft

---

## 1. Design Philosophy

### 1.1 Core Principles

**"Less is More" - Minimalist Excellence**

1. **Clarity Over Cleverness**: Every element serves a purpose
2. **Speed Over Features**: Fast, instant, responsive
3. **Developer-First**: Built for developers, loved by all
4. **Transparency**: Show, don't hide - build trust through openness
5. **Accessibility**: Usable by everyone, everywhere

### 1.2 Design Goals

- **Instant Gratification**: Users see results immediately
- **Zero Friction**: Minimal steps to achieve goals
- **Visual Hierarchy**: Clear focus on primary actions
- **Consistent Experience**: Predictable, reliable interactions
- **Beautiful Simplicity**: Modern, clean, professional

---

## 2. Visual Identity

### 2.1 Color Palette

#### Primary Colors
```css
/* Primary Blue - Trust, Technology, Professional */
--color-primary: #3B82F6;        /* rgb(59, 130, 246) */
--color-primary-dark: #2563EB;  /* rgb(37, 99, 235) */
--color-primary-light: #60A5FA;  /* rgb(96, 165, 250) */

/* Alternative: Modern Purple */
--color-primary-alt: #6366F1;    /* rgb(99, 102, 241) */
--color-primary-alt-dark: #4F46E5; /* rgb(79, 70, 229) */
```

#### Accent Colors
```css
/* Success/Positive Actions */
--color-success: #10B981;        /* rgb(16, 185, 129) */
--color-success-light: #34D399;  /* rgb(52, 211, 153) */

/* Warning/Attention */
--color-warning: #F59E0B;        /* rgb(245, 158, 11) */

/* Error/Destructive */
--color-error: #EF4444;          /* rgb(239, 68, 68) */
```

#### Neutral Colors
```css
/* Text Colors */
--color-text-primary: #111827;   /* rgb(17, 24, 39) */
--color-text-secondary: #6B7280; /* rgb(107, 114, 128) */
--color-text-tertiary: #9CA3AF;  /* rgb(156, 163, 175) */

/* Background Colors */
--color-bg-primary: #FFFFFF;     /* Pure white */
--color-bg-secondary: #FAFAFA;   /* Very light gray */
--color-bg-tertiary: #F3F4F6;    /* Light gray */

/* Border Colors */
--color-border: #E5E7EB;         /* rgb(229, 231, 235) */
--color-border-dark: #D1D5DB;    /* rgb(209, 213, 219) */
```

#### QR Code Colors
```css
/* Default QR Code (Classic) */
--qr-foreground: #000000;        /* Black */
--qr-background: #FFFFFF;        /* White */

/* Future: Customizable colors */
--qr-foreground-custom: var(--color-primary);
--qr-background-custom: var(--color-bg-primary);
```

### 2.2 Typography

#### Font Family
```css
/* Primary Font - System Font Stack for Performance */
--font-primary: -apple-system, BlinkMacSystemFont, 'Segoe UI', 
                'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 
                'Fira Sans', 'Droid Sans', 'Helvetica Neue', 
                sans-serif;

/* Monospace Font - Code/API Examples */
--font-mono: 'SF Mono', 'Monaco', 'Inconsolata', 'Fira Code', 
             'Fira Mono', 'Droid Sans Mono', 'Source Code Pro', 
             monospace;
```

#### Font Sizes
```css
/* Heading Sizes */
--text-4xl: 2.25rem;    /* 36px - Hero title */
--text-3xl: 1.875rem;   /* 30px - Section titles */
--text-2xl: 1.5rem;     /* 24px - Subsection titles */
--text-xl: 1.25rem;     /* 20px - Large text */
--text-lg: 1.125rem;    /* 18px - Body large */
--text-base: 1rem;      /* 16px - Body default */
--text-sm: 0.875rem;    /* 14px - Small text */
--text-xs: 0.75rem;     /* 12px - Extra small */
```

#### Font Weights
```css
--font-light: 300;
--font-normal: 400;
--font-medium: 500;
--font-semibold: 600;
--font-bold: 700;
```

#### Line Heights
```css
--leading-tight: 1.25;
--leading-normal: 1.5;
--leading-relaxed: 1.75;
```

### 2.3 Spacing System

```css
/* Spacing Scale (8px base unit) */
--space-1: 0.25rem;   /* 4px */
--space-2: 0.5rem;    /* 8px */
--space-3: 0.75rem;   /* 12px */
--space-4: 1rem;     /* 16px */
--space-5: 1.25rem;  /* 20px */
--space-6: 1.5rem;   /* 24px */
--space-8: 2rem;     /* 32px */
--space-10: 2.5rem;  /* 40px */
--space-12: 3rem;    /* 48px */
--space-16: 4rem;    /* 64px */
--space-20: 5rem;    /* 80px */
--space-24: 6rem;    /* 96px */
```

### 2.4 Border Radius

```css
--radius-sm: 0.25rem;   /* 4px - Small elements */
--radius-md: 0.5rem;    /* 8px - Buttons, inputs */
--radius-lg: 0.75rem;   /* 12px - Cards */
--radius-xl: 1rem;      /* 16px - Large cards */
--radius-full: 9999px;  /* Fully rounded */
```

### 2.5 Shadows

```css
/* Elevation System */
--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 
             0 2px 4px -1px rgba(0, 0, 0, 0.06);
--shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 
             0 4px 6px -2px rgba(0, 0, 0, 0.05);
--shadow-xl: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 
             0 10px 10px -5px rgba(0, 0, 0, 0.04);
```

---

## 3. Component Library

### 3.1 Input Field

**Purpose**: Primary text input for QR code generation

**Design Specifications**:
```css
.input-primary {
  width: 100%;
  max-width: 600px;
  padding: var(--space-4) var(--space-6);
  font-size: var(--text-lg);
  font-weight: var(--font-normal);
  color: var(--color-text-primary);
  background: var(--color-bg-primary);
  border: 2px solid var(--color-border);
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
}

.input-primary:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}
```

**States**:
- **Default**: Light border, white background
- **Focus**: Primary blue border, subtle shadow
- **Error**: Red border, error message below
- **Disabled**: Grayed out, no interaction

**Accessibility**:
- Proper label association
- ARIA labels for screen readers
- Keyboard navigation support
- Focus indicators visible

### 3.2 QR Preview Container

**Purpose**: Display generated QR code with smooth animations

**Design Specifications**:
```css
.qr-preview {
  width: 100%;
  max-width: 400px;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--color-bg-primary);
  border: 2px dashed var(--color-border);
  border-radius: var(--radius-lg);
  padding: var(--space-6);
  transition: all 0.3s ease;
}

.qr-preview.has-qr {
  border-color: var(--color-primary);
  border-style: solid;
  box-shadow: var(--shadow-md);
}

.qr-preview img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  animation: fadeIn 0.3s ease-in;
}

@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}
```

**States**:
- **Empty**: Dashed border, placeholder text
- **Loading**: Subtle pulse animation (if needed)
- **Loaded**: Solid border, QR code visible, smooth fade-in
- **Error**: Error message displayed

### 3.3 Button Components

#### Primary Button
```css
.btn-primary {
  padding: var(--space-3) var(--space-6);
  font-size: var(--text-base);
  font-weight: var(--font-semibold);
  color: var(--color-bg-primary);
  background: var(--color-primary);
  border: none;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary:hover {
  background: var(--color-primary-dark);
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

.btn-primary:active {
  transform: translateY(0);
}
```

#### Secondary Button
```css
.btn-secondary {
  padding: var(--space-3) var(--space-6);
  font-size: var(--text-base);
  font-weight: var(--font-medium);
  color: var(--color-primary);
  background: transparent;
  border: 2px solid var(--color-primary);
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-secondary:hover {
  background: var(--color-primary);
  color: var(--color-bg-primary);
}
```

#### Icon Button
```css
.btn-icon {
  width: 40px;
  height: 40px;
  padding: var(--space-2);
  background: transparent;
  border: none;
  border-radius: var(--radius-md);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.btn-icon:hover {
  background: var(--color-bg-tertiary);
}
```

### 3.4 Format Toggle

**Purpose**: Switch between PNG and SVG formats

**Design Specifications**:
```css
.format-toggle {
  display: flex;
  gap: var(--space-2);
  padding: var(--space-1);
  background: var(--color-bg-tertiary);
  border-radius: var(--radius-md);
}

.format-option {
  padding: var(--space-2) var(--space-4);
  font-size: var(--text-sm);
  font-weight: var(--font-medium);
  color: var(--color-text-secondary);
  background: transparent;
  border: none;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.2s ease;
}

.format-option.active {
  color: var(--color-text-primary);
  background: var(--color-bg-primary);
  box-shadow: var(--shadow-sm);
}
```

### 3.5 Size Slider

**Purpose**: Adjust QR code size (50-2000px)

**Design Specifications**:
```css
.size-slider {
  width: 100%;
  max-width: 300px;
  height: 6px;
  background: var(--color-bg-tertiary);
  border-radius: var(--radius-full);
  outline: none;
  -webkit-appearance: none;
}

.size-slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  width: 20px;
  height: 20px;
  background: var(--color-primary);
  border-radius: var(--radius-full);
  cursor: pointer;
  box-shadow: var(--shadow-sm);
}

.size-slider::-moz-range-thumb {
  width: 20px;
  height: 20px;
  background: var(--color-primary);
  border-radius: var(--radius-full);
  cursor: pointer;
  border: none;
  box-shadow: var(--shadow-sm);
}
```

### 3.6 Download Actions

**Purpose**: Download QR code in various formats

**Design Specifications**:
```css
.download-actions {
  display: flex;
  gap: var(--space-3);
  flex-wrap: wrap;
}

.download-btn {
  display: flex;
  align-items: center;
  gap: var(--space-2);
  padding: var(--space-3) var(--space-5);
  font-size: var(--text-sm);
  font-weight: var(--font-medium);
  color: var(--color-text-primary);
  background: var(--color-bg-primary);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s ease;
}

.download-btn:hover {
  border-color: var(--color-primary);
  color: var(--color-primary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}
```

### 3.7 Analytics Display

**Purpose**: Show usage statistics and build trust

**Design Specifications**:
```css
.analytics-display {
  display: flex;
  align-items: center;
  gap: var(--space-4);
  padding: var(--space-4);
  background: var(--color-bg-secondary);
  border-radius: var(--radius-lg);
}

.stat-item {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.stat-value {
  font-size: var(--text-2xl);
  font-weight: var(--font-bold);
  color: var(--color-primary);
}

.stat-label {
  font-size: var(--text-xs);
  color: var(--color-text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.status-indicator {
  width: 12px;
  height: 12px;
  border-radius: var(--radius-full);
  background: var(--color-success);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
```

### 3.8 API Playground

**Purpose**: Interactive API testing interface

**Design Specifications**:
```css
.api-playground {
  background: var(--color-bg-secondary);
  border-radius: var(--radius-lg);
  padding: var(--space-6);
  margin-top: var(--space-8);
}

.code-block {
  background: var(--color-text-primary);
  color: var(--color-bg-primary);
  padding: var(--space-4);
  border-radius: var(--radius-md);
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  line-height: var(--leading-relaxed);
  overflow-x: auto;
}

.code-block pre {
  margin: 0;
}

.copy-button {
  position: absolute;
  top: var(--space-2);
  right: var(--space-2);
  padding: var(--space-2);
  background: rgba(255, 255, 255, 0.1);
  border: none;
  border-radius: var(--radius-sm);
  color: var(--color-bg-primary);
  cursor: pointer;
  transition: all 0.2s ease;
}

.copy-button:hover {
  background: rgba(255, 255, 255, 0.2);
}
```

---

## 4. Layout System

### 4.1 Homepage Layout

```
┌─────────────────────────────────────────────────────┐
│  Header (Minimal Navigation)                         │
│  - Logo                                             │
│  - API Docs | Analytics                             │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  Hero Section (Above the Fold)                      │
│                                                      │
│  ┌────────────────────────────────────────────┐    │
│  │  Tagline: "Generate QR codes in             │    │
│  │            milliseconds"                    │    │
│  └────────────────────────────────────────────┘    │
│                                                      │
│  ┌────────────────────────────────────────────┐    │
│  │  [Large Input Field - Centered]            │    │
│  └────────────────────────────────────────────┘    │
│                                                      │
│  ┌──────────────┐  ┌──────────────┐               │
│  │  Format:     │  │  Size:       │               │
│  │  [PNG] [SVG] │  │  [Slider]    │               │
│  └──────────────┘  └──────────────┘               │
│                                                      │
│  ┌────────────────────────────────────────────┐    │
│  │  [QR Preview Container]                     │    │
│  │  (Appears as user types)                    │    │
│  └────────────────────────────────────────────┘    │
│                                                      │
│  ┌────────────────────────────────────────────┐    │
│  │  [Download Actions]                         │    │
│  │  (Appear after QR is generated)             │    │
│  └────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  Analytics Display                                   │
│  [Live Counter] [Status Indicator]                  │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  API Playground Section                              │
│  - Interactive API testing                           │
│  - Code examples                                     │
│  - Request/Response viewer                          │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  Documentation Section                                │
│  - Quick Start                                       │
│  - API Reference                                     │
│  - Code Examples                                     │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  Footer (Minimal)                                    │
│  - Links | Copyright                                 │
└─────────────────────────────────────────────────────┘
```

### 4.2 Grid System

```css
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--space-4);
}

@media (min-width: 640px) {
  .container { padding: 0 var(--space-6); }
}

@media (min-width: 1024px) {
  .container { padding: 0 var(--space-8); }
}

.grid {
  display: grid;
  gap: var(--space-6);
}

.grid-2 { grid-template-columns: repeat(2, 1fr); }
.grid-3 { grid-template-columns: repeat(3, 1fr); }

@media (max-width: 768px) {
  .grid-2, .grid-3 {
    grid-template-columns: 1fr;
  }
}
```

### 4.3 Spacing System

- **Section Spacing**: `var(--space-16)` to `var(--space-24)` between major sections
- **Component Spacing**: `var(--space-6)` to `var(--space-8)` between related components
- **Element Spacing**: `var(--space-2)` to `var(--space-4)` between related elements

---

## 5. Interaction Design

### 5.1 Animations

#### Fade In (QR Code Appearance)
```css
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
```

#### Slide Up (Download Actions)
```css
@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
```

#### Pulse (Status Indicator)
```css
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}
```

### 5.2 Transitions

**Standard Transition**: `transition: all 0.2s ease;`  
**Smooth Transition**: `transition: all 0.3s ease;`  
**Fast Transition**: `transition: all 0.15s ease;`

### 5.3 Hover States

- **Buttons**: Slight lift (`translateY(-2px)`), shadow increase
- **Links**: Color change, underline (if applicable)
- **Cards**: Shadow increase, slight scale
- **Icons**: Background color change, scale

### 5.4 Focus States

- **Visible Focus Ring**: `box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2);`
- **Keyboard Navigation**: All interactive elements focusable
- **Skip Links**: For accessibility

---

## 6. Responsive Design

### 6.1 Breakpoints

```css
/* Mobile First Approach */
/* Base: 0px - 639px (Mobile) */
@media (min-width: 640px) { /* Tablet */ }
@media (min-width: 768px) { /* Tablet Large */ }
@media (min-width: 1024px) { /* Desktop */ }
@media (min-width: 1280px) { /* Desktop Large */ }
```

### 6.2 Mobile Optimizations

- **Touch Targets**: Minimum 44x44px
- **Input Fields**: Large, easy to tap
- **Bottom Sheet**: Options appear in bottom sheet on mobile
- **Swipe Gestures**: Format switching via swipe
- **Stacked Layout**: Vertical stacking on mobile

### 6.3 Tablet Optimizations

- **Two-Column Layout**: Where appropriate
- **Larger Touch Targets**: Maintained
- **Optimized Spacing**: Adjusted for medium screens

### 6.4 Desktop Optimizations

- **Wide Layout**: Maximum content width 1200px
- **Hover Interactions**: Enhanced hover states
- **Keyboard Shortcuts**: Power user features
- **Multi-Column**: Where beneficial

---

## 7. Accessibility

### 7.1 WCAG 2.1 AA Compliance

- **Color Contrast**: Minimum 4.5:1 for text, 3:1 for UI components
- **Text Alternatives**: Alt text for images, ARIA labels
- **Keyboard Navigation**: All functionality accessible via keyboard
- **Focus Indicators**: Visible focus states
- **Screen Reader Support**: Proper semantic HTML, ARIA attributes

### 7.2 Semantic HTML

- Use proper heading hierarchy (h1, h2, h3)
- Use semantic elements (`<nav>`, `<main>`, `<section>`, `<article>`)
- Form labels properly associated
- Button vs link distinction

### 7.3 ARIA Attributes

```html
<!-- Example: QR Preview with ARIA -->
<div 
  role="img" 
  aria-label="QR Code Preview"
  aria-live="polite"
  aria-atomic="true"
>
  <!-- QR Code -->
</div>
```

### 7.4 Keyboard Navigation

- **Tab Order**: Logical, predictable
- **Skip Links**: Jump to main content
- **Keyboard Shortcuts**: 
  - `Enter`: Generate QR (if input focused)
  - `Escape`: Close modals/dropdowns
  - `Ctrl/Cmd + K`: Focus search (if applicable)

---

## 8. Performance Guidelines

### 8.1 Image Optimization

- **QR Codes**: Optimized PNG/SVG output
- **Lazy Loading**: Images below fold
- **Responsive Images**: Appropriate sizes for viewport
- **Format Selection**: SVG for scalability, PNG for compatibility

### 8.2 CSS Optimization

- **Critical CSS**: Inline above-fold styles
- **Unused CSS**: Remove unused styles
- **CSS Variables**: Use for theming
- **Minification**: Production builds minified

### 8.3 JavaScript Optimization

- **Code Splitting**: Route-based splitting
- **Tree Shaking**: Remove unused code
- **Lazy Loading**: Load components on demand
- **Debouncing**: Input debouncing (300ms for QR generation)

### 8.4 Font Optimization

- **System Fonts**: Use system font stack (no external fonts for MVP)
- **Font Display**: `font-display: swap` if using custom fonts
- **Subset Fonts**: Only include needed characters

---

## 9. Design Tokens

### 9.1 Complete Token List

```css
:root {
  /* Colors */
  --color-primary: #3B82F6;
  --color-primary-dark: #2563EB;
  --color-primary-light: #60A5FA;
  --color-success: #10B981;
  --color-warning: #F59E0B;
  --color-error: #EF4444;
  --color-text-primary: #111827;
  --color-text-secondary: #6B7280;
  --color-text-tertiary: #9CA3AF;
  --color-bg-primary: #FFFFFF;
  --color-bg-secondary: #FAFAFA;
  --color-bg-tertiary: #F3F4F6;
  --color-border: #E5E7EB;
  --color-border-dark: #D1D5DB;

  /* Typography */
  --font-primary: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  --font-mono: 'SF Mono', 'Monaco', monospace;
  --text-4xl: 2.25rem;
  --text-3xl: 1.875rem;
  --text-2xl: 1.5rem;
  --text-xl: 1.25rem;
  --text-lg: 1.125rem;
  --text-base: 1rem;
  --text-sm: 0.875rem;
  --text-xs: 0.75rem;
  --font-light: 300;
  --font-normal: 400;
  --font-medium: 500;
  --font-semibold: 600;
  --font-bold: 700;

  /* Spacing */
  --space-1: 0.25rem;
  --space-2: 0.5rem;
  --space-3: 0.75rem;
  --space-4: 1rem;
  --space-5: 1.25rem;
  --space-6: 1.5rem;
  --space-8: 2rem;
  --space-10: 2.5rem;
  --space-12: 3rem;
  --space-16: 4rem;
  --space-20: 5rem;
  --space-24: 6rem;

  /* Border Radius */
  --radius-sm: 0.25rem;
  --radius-md: 0.5rem;
  --radius-lg: 0.75rem;
  --radius-xl: 1rem;
  --radius-full: 9999px;

  /* Shadows */
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  --shadow-xl: 0 20px 25px -5px rgba(0, 0, 0, 0.1);

  /* Transitions */
  --transition-fast: 0.15s ease;
  --transition-base: 0.2s ease;
  --transition-slow: 0.3s ease;
}
```

---

## 10. Component Specifications

### 10.1 Homepage Hero Section

**Layout**: Centered, vertical stack  
**Max Width**: 800px  
**Spacing**: 
- Top padding: `var(--space-16)` (mobile) to `var(--space-24)` (desktop)
- Bottom padding: `var(--space-12)` (mobile) to `var(--space-16)` (desktop)
- Internal spacing: `var(--space-6)` between elements

**Elements**:
1. Tagline (h1): `var(--text-4xl)`, centered, `var(--color-text-primary)`
2. Input Field: Full width, max 600px, centered
3. Format/Size Controls: Horizontal flex, centered, subtle styling
4. QR Preview: Centered, max 400px, appears with animation
5. Download Actions: Horizontal flex, centered, appears after QR generation

### 10.2 API Playground Section

**Layout**: Full width container, card-style  
**Background**: `var(--color-bg-secondary)`  
**Padding**: `var(--space-6)`  
**Border Radius**: `var(--radius-lg)`  
**Margin Top**: `var(--space-16)`

**Elements**:
1. Section Title: `var(--text-2xl)`, `var(--font-semibold)`
2. Request Builder: Form with all API parameters
3. Response Viewer: Code block with syntax highlighting
4. Code Snippets: Tabs for different languages
5. Copy Button: Positioned absolutely in code blocks

---

## 11. Design Deliverables

### 11.1 Required Assets

1. **Logo**: SVG format, multiple sizes (16px, 32px, 64px, 128px)
2. **Favicon**: ICO and PNG formats (16x16, 32x32)
3. **Social Images**: Open Graph image (1200x630px)
4. **Component Mockups**: Key components in Figma/Sketch
5. **Homepage Mockup**: Full homepage design
6. **Mobile Mockups**: Key screens for mobile

### 11.2 Design Tools

- **Design Software**: Figma (recommended) or Sketch
- **Prototyping**: InVision or Figma Prototyping
- **Asset Export**: SVG for icons, PNG for images
- **Design Handoff**: Zeplin or Figma Dev Mode

---

## 12. Implementation Guidelines

### 12.1 CSS Architecture

- **Methodology**: BEM (Block Element Modifier) or utility-first (Tailwind)
- **Organization**: Component-based CSS files
- **Naming**: Descriptive, semantic class names
- **Specificity**: Keep specificity low, avoid !important

### 12.2 Component Structure

```
components/
├── QRPreview/
│   ├── QRPreview.tsx
│   ├── QRPreview.module.css
│   └── index.ts
├── InputField/
│   ├── InputField.tsx
│   ├── InputField.module.css
│   └── index.ts
└── ...
```

### 12.3 Styling Approach

**Option 1: CSS Modules**
- Scoped styles per component
- Type-safe with TypeScript
- Easy to maintain

**Option 2: Tailwind CSS**
- Utility-first approach
- Rapid development
- Consistent design system

**Recommendation**: Tailwind CSS for MVP (faster development, built-in design system)

---

## 13. Design Review Checklist

### 13.1 Visual Design
- [ ] Color contrast meets WCAG AA standards
- [ ] Typography is readable and hierarchical
- [ ] Spacing is consistent throughout
- [ ] Components are visually balanced
- [ ] Brand identity is clear and consistent

### 13.2 User Experience
- [ ] Primary actions are clear and prominent
- [ ] User flow is intuitive and logical
- [ ] Feedback is provided for all interactions
- [ ] Error states are clear and helpful
- [ ] Loading states are appropriate (minimal for instant generation)

### 13.3 Responsive Design
- [ ] Mobile layout is optimized
- [ ] Tablet layout is optimized
- [ ] Desktop layout is optimized
- [ ] Touch targets are appropriate size
- [ ] Text is readable at all sizes

### 13.4 Accessibility
- [ ] Semantic HTML is used
- [ ] ARIA attributes are correct
- [ ] Keyboard navigation works
- [ ] Screen reader testing completed
- [ ] Focus indicators are visible

### 13.5 Performance
- [ ] Images are optimized
- [ ] CSS is optimized
- [ ] JavaScript is optimized
- [ ] Font loading is optimized
- [ ] Bundle size is acceptable

---

**Document Status**: ✅ Ready for Review  
**Next Steps**: Design mockups, component implementation, style guide creation
