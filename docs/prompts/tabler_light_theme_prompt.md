# Tabler.io Light Theme Dashboard Prompt

Recreate the DoPlan Dashboard using **Tabler.io inspired design** with a **light theme** (not dark). Use ONLY HTML and vanilla JavaScript - no frameworks.

---

## Prompt for AI Assistant

```
I need you to recreate a project management dashboard using Tabler.io inspired design with a light, modern theme. Use ONLY HTML and vanilla JavaScript - no frameworks like React, Vue, or Angular.

### Design Style Reference

This dashboard should match the Tabler.io aesthetic with:
- **Light theme** (NOT dark)
- **Background**: `#f4f6fa` (light gray)
- **Cards**: White background (`bg-white`)
- **Primary Color**: `#206bc4` (Tabler blue)
- **Hover Color**: `#1a5c98` (darker blue)
- **Text Colors**: Dark gray (`#1e293b`, `#64748b` for secondary)
- **Borders**: Light gray (`#e2e8f0`, `border-gray-200`)
- **Rounded Corners**: `rounded-xl` (12px), `rounded-lg` (8px)
- **Shadows**: Subtle (`shadow-sm`)
- **Max Width**: `1600px` container

### Layout Structure

- **Top Navigation Bar** (NOT sidebar) - sticky header with logo, nav links, search, notifications, profile
- **Main Content Area** - centered with max-width container
- **Footer** - simple footer with copyright
- **Responsive** - mobile-friendly with collapsible mobile menu

### Current Dashboard Pages

The dashboard has 5 main pages:
1. **index.html** - Main dashboard with stats, charts, activity feed
2. **plan.html** - Kanban board for task management
3. **meetings.html** - Meeting history and calendar
4. **achievements.html** - XP, levels, achievement badges
5. **settings.html** - Project settings and configuration

### Required Components & Style

#### 1. Top Navigation Bar (All Pages)

**Style:**
- White background with bottom border
- Sticky positioning (`sticky top-0 z-50`)
- Height: `64px` (h-16)
- Shadow: `shadow-sm`

**Components:**
- Logo on left: "DoPlan.dev" in Tabler blue (`#206bc4`)
- Navigation links: Dashboard, Plan, Meetings, Achievements, Settings
- Active link: Blue text with light blue background (`bg-[#206bc4]/10`)
- Search bar: Rounded input with search icon (desktop only)
- Notifications bell: With red dot indicator
- User profile: Avatar with name/role (desktop), mobile menu button (mobile)
- Mobile menu: Expandable dropdown below header

#### 2. Main Dashboard (index.html)

**Welcome Hero Section:**
- White card with border (`bg-white rounded-xl border border-gray-200`)
- Padding: `p-8`
- Flex layout: Text on left, illustration on right
- Heading: "Welcome back, [Name]!" (text-3xl font-bold)
- Description with highlighted stats
- Action buttons: Primary (blue) and secondary (white with border)
- Illustration image on right side

**Stats Grid (4 columns):**
- White cards (`bg-white rounded-xl border border-gray-200`)
- Hover effect: Border changes to blue (`hover:border-[#206bc4]/50`)
- Each card shows:
  - Label (text-gray-500 text-sm)
  - Large number (text-3xl font-bold text-gray-800)
  - Progress bar or indicator
  - Trend badge (green for positive, yellow for negative)

**Content Grid:**
- **Chart Section** (2/3 width):
  - White card with Chart.js line chart
  - Gradient fill (blue with transparency)
  - Custom tooltip styling
  - Title with info icon button
  
- **Activity Feed** (1/3 width):
  - White card with scrollable list
  - Each item: Colored dot, text, timestamp
  - Hover effects
  - "View all activity" link at bottom

#### 3. Plan Page (plan.html)

**Header:**
- Title: "Project Plan" (text-3xl font-bold)
- Subtitle: Description text
- Toolbar: Search input, filter button, "New Task" button

**Kanban Board:**
- Three columns: "To Do", "In Progress", "Done"
- Column style: Light gray background (`bg-gray-100/80 rounded-xl border`)
- Column headers: Status dot, title, count badge
- Task cards:
  - White background (`bg-white rounded-lg`)
  - Border (blue left border for in-progress)
  - Tag badge at top
  - Title
  - Progress bar (for in-progress)
  - Assignee avatar
  - Action buttons (move to next column)
  - Hover effects
- "Add Card" button at bottom of each column (dashed border)

#### 4. Achievements Page (achievements.html)

**Hero Section:**
- White card with gradient background decoration
- Circular progress indicator (SVG circle with stroke-dasharray)
- Level display in center
- Description text
- Badge pills (streak, contributor status)
- Illustration on right

**Achievements Grid:**
- Filter buttons: "All", "Locked", "Unlocked"
- Grid: 4 columns (responsive)
- Achievement cards:
  - White background
  - Icon in colored circle
  - Title and description
  - Progress bar at bottom
  - Locked state: Grayscale, reduced opacity
  - Unlocked state: Blue border, full color
  - Hover: Lift effect (`hover:-translate-y-1 hover:shadow-lg`)

#### 5. Meetings Page (meetings.html)

**Style:**
- Similar card-based layout
- Timeline or calendar view
- Meeting cards with details
- Action items list

#### 6. Settings Page (settings.html)

**Style:**
- Tabbed interface (General, Agents, Integrations)
- Form components with labels
- Toggle switches
- Tables for data display

### Technical Requirements

1. **CSS Framework:**
   - Use Tailwind CSS via CDN: `https://cdn.tailwindcss.com`
   - OR use Tabler.io CSS: `https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler.min.css`
   - Custom CSS for specific styling

2. **JavaScript Libraries:**
   - Chart.js: `https://cdn.jsdelivr.net/npm/chart.js@latest`
   - Marked.js: `https://cdn.jsdelivr.net/npm/marked/marked.min.js`
   - Canvas Confetti: `https://cdn.jsdelivr.net/npm/canvas-confetti@latest/dist/confetti.browser.min.js`

3. **Icons:**
   - Use Tabler icons via webfont: `https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css`
   - OR use inline SVG icons (like in the reference)

4. **JavaScript:**
   - Vanilla JavaScript only
   - Keep all existing data loading functions (fetch API)
   - Add smooth transitions and hover effects
   - Mobile menu toggle functionality
   - Chart initialization
   - Task management functions

### Color Palette

```css
Primary Blue: #206bc4
Hover Blue: #1a5c98
Background: #f4f6fa
Card Background: #ffffff
Text Primary: #1e293b
Text Secondary: #64748b
Border: #e2e8f0
Success: #10b981 (green)
Warning: #f59e0b (yellow)
Error: #ef4444 (red)
```

### Typography

- Font: System sans-serif (Inter, -apple-system, etc.)
- Headings: Bold, dark gray
- Body: Medium weight, gray
- Small text: Light gray, smaller size

### Spacing & Layout

- Container: `max-w-[1600px] mx-auto px-4 sm:px-6 lg:px-8`
- Card padding: `p-6` or `p-8`
- Gap between cards: `gap-6`
- Section spacing: `space-y-8` or `mb-6`

### Interactive Elements

- Buttons: Rounded, with hover states
- Cards: Hover border color change, shadow increase
- Links: Blue color, underline on hover
- Inputs: Focus ring in blue
- Transitions: `transition-all` for smooth effects

### Responsive Design

- Mobile: Stack columns, hide search, show mobile menu
- Tablet: 2-column grids
- Desktop: Full layout with all features
- Breakpoints: sm (640px), md (768px), lg (1024px), xl (1280px)

### Code Structure

Generate complete HTML files with:
- Semantic HTML5
- Tailwind CSS classes (or Tabler classes)
- Inline SVG icons
- Vanilla JavaScript in `<script>` tags
- All data fetching functions preserved
- Error handling
- Loading states

### File Structure

```
dashboard/
├── index.html          (Main dashboard)
├── plan.html           (Kanban board)
├── meetings.html       (Meetings)
├── achievements.html   (Achievements)
├── settings.html       (Settings)
└── data/
    └── project.json
```

### Key Design Principles

1. **Clean & Modern**: Lots of white space, clean lines
2. **Light & Airy**: Light background, not heavy
3. **Professional**: Business-appropriate styling
4. **Consistent**: Same card style, spacing, colors throughout
5. **Accessible**: Good contrast, readable text
6. **Interactive**: Smooth hover effects, transitions

### Deliverables

Please provide:
1. Complete HTML code for all 5 pages
2. Matching the exact light theme style described
3. Tabler.io inspired design
4. Vanilla JavaScript for all interactions
5. Responsive design
6. All existing functionality preserved

Make sure the design is clean, modern, professional, and matches the Tabler.io light theme aesthetic exactly.
```

---

## Key Style Points

### Colors
- **Background**: `#f4f6fa` (light gray, not dark)
- **Cards**: White (`#ffffff`)
- **Primary**: `#206bc4` (Tabler blue)
- **Text**: Dark gray (`#1e293b`)

### Layout
- **Top navigation** (not sidebar)
- **Max width**: `1600px`
- **Padding**: Responsive (`px-4 sm:px-6 lg:px-8`)

### Components
- **Cards**: White with borders, rounded corners
- **Hover effects**: Border color change, shadow increase
- **Buttons**: Blue primary, white secondary
- **Badges**: Colored backgrounds with borders

### Typography
- **Headings**: Bold, large
- **Body**: Medium weight
- **Small text**: Light gray

## Usage

1. Copy the prompt above
2. Paste into your AI assistant
3. Review generated code
4. Test locally
5. Integrate into `dashboard.go`
