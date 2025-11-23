# Design System
## DoPlan CLI v1.0

**Version**: 1.0  
**Status**: Draft  
**Last Updated**: November 2025  
**Owners**: Design & UX Manager, UI/UX Designer

---

## 🎨 Design Philosophy

### Core Principles
1. **Clarity First**: Every element should be clear and purposeful
2. **Delightful Experience**: Make developers smile, not frustrated
3. **Accessibility**: Works for everyone, regardless of ability
4. **Consistency**: Predictable patterns and interactions
5. **Performance**: Fast, responsive, never laggy

### Design Goals
- **Magical, not tedious**: The wizard should feel like a delightful experience
- **Professional yet friendly**: Serious tool, approachable interface
- **Keyboard-first**: Full functionality without mouse
- **Error-tolerant**: Help users recover from mistakes gracefully

---

## 🎨 Color Palette

### Primary Colors
- **Purple** (`#A855F7`): Primary actions, headers, important information
- **Pink** (`#EC4899`): Accents, highlights, secondary actions

### Status Colors
- **Green** (`#10B981`): Success states, completed actions, positive feedback
- **Blue** (`#3B82F6`): Information, hints, helpful messages
- **Yellow** (`#F59E0B`): Warnings, non-critical issues, attention needed
- **Red** (`#EF4444`): Errors, critical issues, destructive actions

### Neutral Colors
- **White** (`#FFFFFF`): Text on dark backgrounds, primary text
- **Gray-100** (`#F3F4F6`): Light backgrounds, subtle borders
- **Gray-800** (`#1F2937`): Dark backgrounds, high contrast text
- **Gray-900** (`#111827`): Terminal background, deepest contrast

### Color Usage Guidelines
- **Primary Actions**: Purple/Pink
- **Success Messages**: Green with ✅ icon
- **Info Messages**: Blue with ℹ️ icon
- **Warnings**: Yellow with ⚠️ icon
- **Errors**: Red with ❌ icon
- **Text**: White on dark, Gray-900 on light

---

## 📝 Typography

### Font Family
- **Primary**: System monospace font (Courier New, Monaco, Consolas)
- **Fallback**: Any monospace font available on system

### Font Sizes
- **Header**: 24px (2 lines) - Welcome screen, major sections
- **Title**: 18px (1.5 lines) - Step titles, section headers
- **Body**: 14px (1.2 lines) - Main content, descriptions
- **Small**: 12px (1.1 lines) - Hints, help text, footers

### Font Weights
- **Bold**: Headers, important text, selected items
- **Regular**: Body text, descriptions
- **Light**: Hints, secondary information

### Line Height
- **Headers**: 1.2 (tight, for impact)
- **Body**: 1.5 (comfortable reading)
- **Code**: 1.2 (compact, for code blocks)

---

## 🧩 Component Library

### 1. Welcome Screen
```
┌─────────────────────────────────────────┐
│                                         │
│         🚀 DoPlan CLI                   │
│   Create professional projects          │
│         in seconds                      │
│                                         │
│  Press Enter to continue                │
│  Press 'q' to quit                      │
│                                         │
└─────────────────────────────────────────┘
```

**Design Specs**:
- Centered ASCII art with emojis
- Purple/Pink gradient for title
- Clear instructions at bottom
- Border around content area

### 2. Text Input Field
```
Project Name: [my-awesome-app        ]
              ^cursor
```

**Design Specs**:
- Label in purple
- Input field with border
- Cursor indicator
- Real-time validation feedback
- Character count (optional)

### 3. Selection Menu
```
Select your IDE:
  ○ Cursor (Recommended) ⭐
  ○ Claude Code
  ○ Antigravity
  ○ Windsurf
  ○ Cline
  ○ OpenCode

Use ↑/↓ to navigate, Enter to select
```

**Design Specs**:
- Radio button indicators (○/●)
- Highlight selected item in purple
- Show recommended options with ⭐
- Keyboard navigation hints
- Clear visual hierarchy

### 4. Progress Screen
```
Generating your project...

  ✅ Creating directory structure
  ✅ Generating AI agents
  ✅ Extracting rules library
  ⏳ Creating GitHub workflows...
  ⏸  Setting up boilerplate

[████████░░] 60% complete
```

**Design Specs**:
- List of steps with status icons
- Progress bar at bottom
- Percentage indicator
- Smooth animations
- Color-coded status (green ✅, yellow ⏳, gray ⏸)

### 5. Success Screen
```
✨ Project created successfully!

✅ my-awesome-app/
  ├── .cursor/
  ├── .plan/
  ├── .github/
  └── src/

Next steps:
  1. Open with: code ./my-awesome-app
  2. Then type /tell to begin

Press Enter to exit
```

**Design Specs**:
- Celebration emoji (✨)
- Green checkmarks for success
- Tree view of created structure
- Clear next steps
- Actionable instructions

### 6. Error Message
```
❌ Error: Directory 'my-project' already exists

Suggestion: Choose a different name or delete the existing directory.

Press Enter to try again
```

**Design Specs**:
- Red ❌ icon
- Clear error message
- Actionable suggestion
- Recovery option
- Helpful, not technical

---

## 🎭 UI Patterns

### 1. Wizard Flow Pattern
**Structure**:
1. Welcome → 2. Input → 3. Selection → 4. Progress → 5. Success

**Transitions**:
- Smooth fade between steps
- Clear progress indicator (Step 2 of 4)
- Back navigation (optional, with Backspace)

### 2. Input Validation Pattern
**Real-time Feedback**:
- Valid input: Green border, ✅ icon
- Invalid input: Red border, ❌ icon, error message
- Neutral: Gray border, no icon

**Examples**:
- ✅ "my-awesome-app" (valid)
- ❌ "my awesome app" (spaces not allowed)
- ❌ "MyAwesomeApp" (uppercase not recommended)

### 3. Selection Pattern
**Visual Indicators**:
- Selected: Purple background, white text, ● indicator
- Unselected: Transparent background, gray text, ○ indicator
- Recommended: ⭐ icon next to option

**Navigation**:
- ↑/↓ to move selection
- Enter to confirm
- Esc to cancel

### 4. Progress Pattern
**Status Icons**:
- ✅ Completed (green)
- ⏳ In progress (yellow, animated spinner)
- ⏸ Pending (gray)
- ❌ Failed (red)

**Progress Bar**:
- Filled portion: Purple gradient
- Unfilled portion: Gray
- Percentage text: White

### 5. Message Pattern
**Message Types**:
- **Success**: Green ✅ + message + optional action
- **Info**: Blue ℹ️ + message + optional link
- **Warning**: Yellow ⚠️ + message + optional suggestion
- **Error**: Red ❌ + message + recovery suggestion

**Layout**:
```
[Icon] Message text
       Optional action/suggestion
```

---

## ⌨️ Interaction Patterns

### Keyboard Navigation
- **Enter**: Confirm, proceed to next step
- **↑/↓**: Navigate menu options
- **Tab**: Move between input fields (if multiple)
- **Esc**: Cancel, go back
- **q**: Quit application
- **Ctrl+C**: Force quit

### Input Handling
- **Real-time validation**: Validate as user types
- **Auto-focus**: Focus first input automatically
- **Clear feedback**: Show validation state immediately
- **Error prevention**: Prevent invalid input when possible

### Animation Guidelines
- **Transitions**: 200-300ms for smooth feel
- **Loading spinners**: Rotate continuously, 1s per rotation
- **Progress bars**: Animate smoothly, not jumpy
- **Fade effects**: 150ms fade in/out

---

## ♿ Accessibility

### Keyboard-Only Navigation
- All functionality accessible via keyboard
- Clear keyboard shortcuts
- Focus indicators visible
- Logical tab order

### Screen Reader Support
- Text labels for all icons
- Descriptive alt text
- Status announcements
- Progress announcements

### Color Blind Support
- Don't rely solely on color
- Use icons + text for status
- High contrast ratios (4.5:1 minimum)
- Patterns/textures for differentiation

### Low Vision Support
- High contrast mode option
- Adjustable font sizes (if possible)
- Clear visual hierarchy
- Sufficient spacing

---

## 📱 Responsive Design

### Terminal Size Considerations
- **Minimum**: 80x24 characters
- **Optimal**: 120x40 characters
- **Adaptive**: Layout adjusts to terminal size
- **Graceful degradation**: Works on small terminals

### Layout Strategies
- **Centered content**: For welcome/success screens
- **Left-aligned**: For lists and menus
- **Full-width**: For progress bars
- **Flexible**: Adapts to terminal width

---

## 🎬 Animation & Transitions

### Loading States
- **Spinner**: Rotating character (|, /, -, \)
- **Progress bar**: Animated fill
- **Pulse**: Subtle pulse for active elements

### Transitions
- **Fade**: Smooth fade between screens
- **Slide**: Optional slide effect (if supported)
- **Duration**: 200-300ms for responsiveness

### Micro-interactions
- **Hover effects**: Highlight on selection (if mouse supported)
- **Focus indicators**: Clear focus ring
- **Button press**: Visual feedback on action

---

## 📐 Spacing System

### Vertical Spacing
- **Tight**: 1 line (related items)
- **Normal**: 2 lines (sections)
- **Loose**: 4 lines (major sections)

### Horizontal Spacing
- **Tight**: 1 character (related items)
- **Normal**: 2 characters (standard spacing)
- **Loose**: 4 characters (separated items)

### Padding
- **Content padding**: 2 characters from edges
- **Section padding**: 4 characters between sections
- **Component padding**: 1 character inside components

---

## 🎯 Design Tokens

### Borders
- **Thin**: Single character (─, │, ┌, ┐, └, ┘)
- **Thick**: Double character (═, ║, ╔, ╗, ╚, ╝)
- **Rounded**: Use Unicode box-drawing characters

### Icons
- ✅ Success/Complete
- ❌ Error/Failed
- ⚠️ Warning
- ℹ️ Information
- ⏳ In Progress
- ⏸ Pending
- ⭐ Recommended
- 🚀 Launch/Start
- ✨ Success/Celebration

### Patterns
- **Stripe**: For progress bars
- **Dots**: For loading states
- **Borders**: For content containers

---

## 📋 Design Checklist

### Before Implementation
- [ ] All colors meet contrast requirements
- [ ] All interactions have keyboard alternatives
- [ ] All icons have text labels
- [ ] Error messages are actionable
- [ ] Success messages are clear
- [ ] Loading states are visible
- [ ] Transitions are smooth
- [ ] Layout works on small terminals

### User Testing
- [ ] Test with keyboard-only navigation
- [ ] Test with screen reader
- [ ] Test with color blindness simulators
- [ ] Test on different terminal sizes
- [ ] Test error scenarios
- [ ] Test on different operating systems

---

## 🔄 Design Iteration Process

1. **Design**: Create initial design
2. **Review**: Team review and feedback
3. **Refine**: Update based on feedback
4. **Implement**: Build in code
5. **Test**: User testing and feedback
6. **Iterate**: Continuous improvement

---

**Document Status**: ✅ Complete  
**Next Step**: Review and approve, then type `/good` to lock plan and generate tasks.
