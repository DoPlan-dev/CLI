# Flowbite Dashboard Recreation Prompt

Use this prompt with an AI assistant (Claude, ChatGPT, etc.) to recreate the DoPlan Dashboard using Flowbite instead of Tabler/Bootstrap.

---

## Prompt for AI Assistant

```
I need you to recreate a project management dashboard using Flowbite (Tailwind CSS component library) instead of Tabler/Bootstrap. The dashboard should have a modern, dark theme with enhanced components and better UX.

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
- Fixed sidebar navigation on the left
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
- 📊 **Activity Timeline** - Recent project activities/commits
- 📈 **Progress Chart** - Visual chart showing progress over time (use Chart.js or similar)
- 🎯 **Quick Actions** - Buttons for common actions (New Task, Start Meeting, etc.)
- 📋 **Recent Tasks** - List of last 5 tasks with status
- 🔔 **Notifications/Alerts** - System notifications or reminders
- ⚡ **Performance Metrics** - Cards showing velocity, completion rate, etc.
- 🏆 **Recent Achievements** - Mini achievement cards
- 📅 **Upcoming Deadlines** - Calendar widget or list

#### 2. Plan Page (plan.html)

**Current Features:**
- Markdown task list viewer
- Search functionality

**ADD These Enhanced Components:**
- 🔍 **Advanced Filters** - Filter by status, phase, assignee, priority
- 📊 **Task Board View** - Kanban-style board (To Do, In Progress, Done)
- 📈 **Task Statistics** - Charts showing task distribution
- 🏷️ **Tag System** - Visual tags for categories
- ⏱️ **Time Tracking** - Estimated vs actual time
- 📝 **Task Details Modal** - Click task to see full details
- ✅ **Bulk Actions** - Select multiple tasks for batch operations

#### 3. Meetings Page (meetings.html)

**Current Features:**
- Meeting history timeline
- Upcoming meetings

**ADD These Enhanced Components:**
- 📅 **Calendar View** - Full calendar with meeting markers
- 🎥 **Meeting Notes** - Rich text editor for meeting notes
- 📊 **Meeting Analytics** - Stats on meeting frequency, duration
- 👥 **Participants** - List of meeting attendees
- 📎 **Attachments** - File attachments for meetings
- 🔔 **Meeting Reminders** - Notifications before meetings
- 📝 **Action Items** - Track action items from meetings

#### 4. Achievements Page (achievements.html)

**Current Features:**
- XP, Level, Streak display
- Achievement cards grid

**ADD These Enhanced Components:**
- 📊 **Progress Visualization** - Circular progress for XP to next level
- 🏅 **Achievement Categories** - Filter by category (tasks, meetings, milestones)
- 📈 **Stats Dashboard** - Charts showing achievement trends
- 🎖️ **Badge Collection** - Visual badge gallery
- 🏆 **Leaderboard** - If multi-user, show rankings
- 📅 **Achievement Timeline** - When achievements were unlocked
- 🎯 **Goals/Targets** - Set and track achievement goals

#### 5. Settings Page (settings.html)

**Current Features:**
- Project settings table
- AI agents list
- Memory insights

**ADD These Enhanced Components:**
- ⚙️ **Settings Categories** - Tabs for General, Agents, Integrations, etc.
- 🔧 **Agent Configuration** - Edit agent settings, enable/disable
- 🔌 **Integrations** - Connect external services (GitHub, Slack, etc.)
- 📊 **Analytics Dashboard** - Project analytics and insights
- 🔐 **Access Control** - User permissions and roles
- 📝 **Project Documentation** - Quick links to docs
- 🎨 **Theme Customization** - Color scheme picker
- 🔔 **Notification Settings** - Configure alerts and notifications

### Technical Requirements

1. **Use Flowbite Components:**
   - Sidebar navigation
   - Cards
   - Tables
   - Modals
   - Dropdowns
   - Forms
   - Charts/Graphs
   - Badges
   - Alerts
   - Buttons
   - Progress bars

2. **Include These Libraries:**
   - Tailwind CSS (via CDN)
   - Flowbite JS (via CDN)
   - Heroicons (for icons)
   - Chart.js or ApexCharts (for data visualization)
   - Marked.js (for markdown rendering)
   - Canvas Confetti (for achievements)

3. **JavaScript Functionality:**
   - Keep all existing data loading functions
   - Add smooth transitions and animations
   - Implement dark mode toggle (optional)
   - Add loading states and skeletons
   - Error handling with user-friendly messages

4. **Responsive Design:**
   - Mobile-first approach
   - Collapsible sidebar on mobile
   - Touch-friendly interactions
   - Optimized for tablets

5. **Performance:**
   - Lazy load heavy components
   - Optimize images and assets
   - Minimize JavaScript bundle
   - Use CDN for all libraries

### Code Structure

Generate complete HTML files with:
- Proper semantic HTML5
- Inline Tailwind CSS classes
- Flowbite component classes
- JavaScript in `<script>` tags
- All data fetching functions preserved
- Error handling
- Loading states

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

### Additional Features to Consider

- **Search Bar** - Global search across all pages
- **Notifications Bell** - Notification center
- **User Profile** - User avatar and dropdown
- **Keyboard Shortcuts** - Power user features
- **Export/Import** - Data export functionality
- **Print Views** - Printer-friendly layouts
- **Accessibility** - ARIA labels, keyboard navigation
- **Animations** - Smooth page transitions
- **Tooltips** - Helpful tooltips on hover

### Deliverables

Please provide:
1. Complete HTML code for all 5 pages
2. Explanation of new components added
3. Instructions for integration
4. Any additional CSS/JS files needed
5. Migration notes from old to new version

Make sure the design is modern, professional, and maintains the dark cyan theme while being more feature-rich and user-friendly.
```

---

## Usage Instructions

1. **Copy the prompt above** and paste it into your AI assistant (Claude, ChatGPT, etc.)

2. **Review the generated code** and test it locally

3. **Customize as needed** - Adjust colors, add/remove components

4. **Integrate into Go code** - Replace the HTML strings in `dashboard.go`

5. **Test thoroughly** - Ensure all JavaScript functions work correctly

## Additional Customization Ideas

- Add more chart types (pie charts, line graphs, etc.)
- Implement drag-and-drop for task management
- Add real-time updates with WebSockets
- Create a mobile app version
- Add data export (PDF, CSV, JSON)
- Implement user preferences persistence
- Add keyboard shortcuts
- Create a command palette (Cmd+K style)

## Migration Checklist

- [ ] Replace Bootstrap/Tabler with Flowbite
- [ ] Update all CSS classes to Tailwind
- [ ] Test all JavaScript functionality
- [ ] Verify responsive design
- [ ] Check browser compatibility
- [ ] Test data loading functions
- [ ] Verify all links work
- [ ] Test on mobile devices
- [ ] Update Go generator code
- [ ] Run automated tests

