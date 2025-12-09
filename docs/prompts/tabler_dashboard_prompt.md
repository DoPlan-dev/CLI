# Tabler.io Dashboard Recreation Prompt

Use this prompt with an AI assistant (Claude, ChatGPT, etc.) to recreate the DoPlan Dashboard using Tabler.io components with enhanced features.

---

## Prompt for AI Assistant

```
I need you to recreate a project management dashboard using Tabler.io (https://tabler.io) components. Use ONLY HTML and vanilla JavaScript - no frameworks. The dashboard should have a modern, dark theme with enhanced components and better UX.

### Current Dashboard Structure

The dashboard has 5 main pages:
1. **index.html** - Main dashboard with project overview, progress, tasks, and phase
2. **plan.html** - Task viewer with markdown rendering and search
3. **meetings.html** - Meeting history and upcoming meetings
4. **achievements.html** - XP, levels, streaks, and achievement cards
5. **settings.html** - Project settings, AI agents list, and memory insights

### Design Requirements

**Color Scheme:**
- Background: `#0A0E27` (dark navy)
- Primary/Accent: `#00D9FF` (cyan)
- Card Background: `#1a1f2e` (slightly lighter navy)
- Text: White with muted gray for secondary text
- Glow effect: `text-shadow: 0 0 20px rgba(0,217,255,0.6)`

**Layout:**
- Fixed sidebar navigation on the left (Tabler sidebar component)
- Main content area on the right
- Responsive design (mobile-friendly)
- Dark theme throughout

### Required Components & Enhancements

#### 1. Main Dashboard (index.html)

**Current Components:**
- Project Overview Card (avatar, name, type, created date)
- Progress Card (progress bar, completion stats)
- Tasks Stat Card (total tasks, completed count)
- Current Phase Card (phase name, status)

**ADD These Enhanced Components:**
- 📊 **Activity Timeline** - Use Tabler timeline component for recent project activities
- 📈 **Progress Chart** - Use Tabler chart component or Chart.js for visual progress
- 🎯 **Quick Actions** - Tabler button group for common actions (New Task, Start Meeting, etc.)
- 📋 **Recent Tasks** - Tabler table or list component showing last 5 tasks with status badges
- 🔔 **Notifications/Alerts** - Tabler alert component for system notifications
- ⚡ **Performance Metrics** - Tabler stat cards showing velocity, completion rate, etc.
- 🏆 **Recent Achievements** - Tabler card grid with mini achievement cards
- 📅 **Upcoming Deadlines** - Tabler calendar component or date list

#### 2. Plan Page (plan.html)

**Current Features:**
- Markdown task list viewer
- Search functionality

**ADD These Enhanced Components:**
- 🔍 **Advanced Filters** - Tabler form components (select dropdowns, checkboxes) for status, phase, priority
- 📊 **Task Board View** - Tabler card layout in Kanban style (To Do, In Progress, Done columns)
- 📈 **Task Statistics** - Tabler chart components showing task distribution
- 🏷️ **Tag System** - Tabler badges for categories/tags
- ⏱️ **Time Tracking** - Tabler progress bars for estimated vs actual time
- 📝 **Task Details Modal** - Tabler modal component for task details
- ✅ **Bulk Actions** - Tabler form with checkboxes for selecting multiple tasks
- 📑 **View Toggle** - Tabler tabs to switch between list and board view

#### 3. Meetings Page (meetings.html)

**Current Features:**
- Meeting history timeline
- Upcoming meetings

**ADD These Enhanced Components:**
- 📅 **Calendar View** - Tabler calendar component with meeting markers
- 🎥 **Meeting Notes** - Tabler form with textarea for meeting notes
- 📊 **Meeting Analytics** - Tabler stat cards and charts for meeting stats
- 👥 **Participants** - Tabler avatar group component for attendees
- 📎 **Attachments** - Tabler file input and file list components
- 🔔 **Meeting Reminders** - Tabler alert component for reminders
- 📝 **Action Items** - Tabler table or list for tracking action items
- 🎯 **Meeting Templates** - Tabler card grid for meeting templates

#### 4. Achievements Page (achievements.html)

**Current Features:**
- XP, Level, Streak display
- Achievement cards grid

**ADD These Enhanced Components:**
- 📊 **Progress Visualization** - Tabler progress component (circular or linear) for XP to next level
- 🏅 **Achievement Categories** - Tabler tabs or pills for filtering by category
- 📈 **Stats Dashboard** - Tabler chart components showing achievement trends
- 🎖️ **Badge Collection** - Tabler card grid with badge/icon components
- 🏆 **Leaderboard** - Tabler table component for rankings (if multi-user)
- 📅 **Achievement Timeline** - Tabler timeline component showing when achievements were unlocked
- 🎯 **Goals/Targets** - Tabler form components for setting and tracking goals
- 🎉 **Unlock Animation** - JavaScript animation when viewing achievements

#### 5. Settings Page (settings.html)

**Current Features:**
- Project settings table
- AI agents list
- Memory insights

**ADD These Enhanced Components:**
- ⚙️ **Settings Categories** - Tabler tabs component (General, Agents, Integrations, Notifications)
- 🔧 **Agent Configuration** - Tabler form components (switches, inputs) for agent settings
- 🔌 **Integrations** - Tabler card grid for integration options with toggle switches
- 📊 **Analytics Dashboard** - Tabler chart and stat card components
- 🔐 **Access Control** - Tabler table with form controls for permissions
- 📝 **Project Documentation** - Tabler list group with links to docs
- 🎨 **Theme Customization** - Tabler color picker or select for theme
- 🔔 **Notification Settings** - Tabler form with switches and checkboxes

### Technical Requirements

1. **Use Tabler.io Components ONLY:**
   - Include Tabler CSS: `https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler.min.css`
   - Include Tabler JS: `https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/js/tabler.min.js`
   - Include Tabler Icons: `https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/dist/tabler-icons.min.css`
   - Use Tabler's sidebar, cards, tables, modals, forms, charts, badges, alerts, buttons, progress bars
   - Reference: https://tabler.io/docs

2. **Additional Libraries (CDN only):**
   - Chart.js (for advanced charts): `https://cdn.jsdelivr.net/npm/chart.js@latest`
   - Marked.js (for markdown): `https://cdn.jsdelivr.net/npm/marked/marked.min.js`
   - Canvas Confetti (for achievements): `https://cdn.jsdelivr.net/npm/canvas-confetti@latest/dist/confetti.browser.min.js`

3. **JavaScript Requirements:**
   - Use ONLY vanilla JavaScript (no frameworks like React, Vue, etc.)
   - Keep all existing data loading functions (fetch API)
   - Add smooth transitions using CSS and JavaScript
   - Implement loading states with Tabler's loading components
   - Error handling with Tabler alert components
   - Use Tabler's built-in JavaScript utilities where possible

4. **HTML Structure:**
   - Use Tabler's page layout structure
   - Use Tabler's sidebar navigation component
   - Use Tabler's card components for content
   - Use Tabler's grid system
   - Semantic HTML5

5. **Responsive Design:**
   - Use Tabler's responsive utilities
   - Collapsible sidebar on mobile (Tabler's built-in functionality)
   - Touch-friendly interactions
   - Mobile-optimized tables and forms

6. **Styling:**
   - Override Tabler's default colors with custom CSS variables
   - Use Tabler's dark theme as base
   - Custom CSS for glow effects and accent colors
   - Maintain the dark cyan theme (#0A0E27, #00D9FF)

### Code Structure

Generate complete HTML files with:
- Tabler.io CSS and JS via CDN
- Tabler component classes (sidebar, cards, tables, etc.)
- Vanilla JavaScript in `<script>` tags
- All existing data fetching functions preserved
- Custom CSS for theme customization
- Error handling with Tabler alerts
- Loading states with Tabler components

### File Structure

```
dashboard/
├── index.html          (Main dashboard with enhanced components)
├── plan.html           (Enhanced task management)
├── meetings.html       (Enhanced meeting management)
├── achievements.html   (Enhanced achievements)
├── settings.html       (Enhanced settings)
└── data/
    └── project.json
```

### Tabler Components to Use

- **Layout**: `page`, `page-wrapper`, `navbar`, `navbar-vertical`, `sidebar`
- **Cards**: `card`, `card-header`, `card-body`, `card-title`
- **Tables**: `table`, `table-vcenter`, `table-striped`
- **Forms**: `form-control`, `form-select`, `form-check`, `form-switch`
- **Buttons**: `btn`, `btn-primary`, `btn-icon`
- **Badges**: `badge`, `bg-primary`, `bg-success`
- **Alerts**: `alert`, `alert-info`
- **Progress**: `progress`, `progress-bar`
- **Charts**: Tabler's chart components or Chart.js integration
- **Modals**: `modal`, `modal-dialog`, `modal-content`
- **Tabs**: `nav-tabs`, `tab-content`
- **Timeline**: Tabler timeline components
- **Avatars**: `avatar`, `avatar-lg`
- **Icons**: Tabler icons (`ti ti-dashboard`, etc.)

### Additional Features to Consider

- **Search Bar** - Tabler form input with search icon
- **Notifications Bell** - Tabler dropdown with notification list
- **User Profile** - Tabler avatar with dropdown menu
- **Keyboard Shortcuts** - JavaScript keyboard event handlers
- **Export/Import** - JavaScript functions for data export
- **Print Views** - CSS print media queries
- **Accessibility** - ARIA labels, keyboard navigation
- **Animations** - CSS transitions and JavaScript for smooth effects
- **Tooltips** - Tabler's tooltip component

### Deliverables

Please provide:
1. Complete HTML code for all 5 pages using Tabler.io components
2. Vanilla JavaScript for all interactions (no frameworks)
3. Custom CSS for theme customization (dark cyan theme)
4. Explanation of new components added
5. Instructions for integration
6. Notes on Tabler component usage

Make sure the design is modern, professional, and maintains the dark cyan theme while being more feature-rich and user-friendly. Use Tabler.io's component library extensively for a consistent, polished look.
```

---

## Usage Instructions

1. **Copy the prompt above** and paste it into your AI assistant (Claude, ChatGPT, etc.)

2. **Review the generated code** and test it locally

3. **Customize as needed** - Adjust colors, add/remove components

4. **Integrate into Go code** - Replace the HTML strings in `dashboard.go`

5. **Test thoroughly** - Ensure all JavaScript functions work correctly

## Tabler.io Resources

- **Documentation**: https://tabler.io/docs
- **Components**: https://tabler.io/docs/components
- **Icons**: https://tabler.io/icons
- **Examples**: https://tabler.io/preview

## Key Tabler Features to Leverage

- **Sidebar Navigation**: Built-in responsive sidebar
- **Dark Theme**: Native dark mode support
- **Charts**: Built-in chart components
- **Forms**: Comprehensive form components
- **Tables**: Advanced table features
- **Modals**: Easy modal dialogs
- **Icons**: 2000+ free icons included

## Migration Checklist

- [ ] Replace current components with Tabler equivalents
- [ ] Update CSS to use Tabler classes
- [ ] Convert JavaScript to vanilla JS (remove any framework code)
- [ ] Test all Tabler components
- [ ] Verify responsive design
- [ ] Check browser compatibility
- [ ] Test data loading functions
- [ ] Verify all links work
- [ ] Test on mobile devices
- [ ] Update Go generator code
- [ ] Run automated tests

## Notes

- Tabler.io is a free, open-source admin dashboard template
- It's built on Bootstrap but has its own component library
- All components are accessible and responsive
- Icons are included via webfont
- JavaScript is minimal and framework-free
