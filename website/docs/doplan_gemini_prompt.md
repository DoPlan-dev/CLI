# DoPlan CLI Website Build Prompt for Gemini 3

## SYSTEM CONTEXT
You are an expert full-stack web developer specializing in high-performance, visually stunning React applications. Your task is to build the complete DoPlan CLI marketing website based on the exact specifications below. This is a production-ready build, not a prototype. Output must be deployable to Vercel immediately.

---

## PROJECT OVERVIEW

**Project Name:** DoPlan CLI Marketing Website  
**Theme:** Digital Occultism meets Cyberpunk Mystery  
**Key Constraint:** Forced Dark Mode (no light mode). Glassmorphism, neon effects, terminal aesthetics, particle systems.  
**Target Audience:** Solo developers, small teams, AI enthusiasts  
**Success Metrics:** <1.5s load time, copy-to-clipboard in <30s, conversion-focused

---

## TECHNICAL SPECIFICATIONS

### Tech Stack (NON-NEGOTIABLE)
- **React 18** + TypeScript (strict mode)
- **Vite** (build tool)
- **Tailwind CSS** (styling) + **shadcn/ui** (components)
- **Lucide React** (outline icons ONLY—no solid icons)
- **Framer Motion** (animations, spring physics)
- **next-themes** (enforce dark mode globally)
- **Radix UI** (accessible primitives under shadcn/ui)

### Performance Requirements
- Lighthouse: 100/100 (all metrics)
- Bundle size: <200kb gzipped
- First Paint: <1.5s on 4G
- Responsive: Desktop (full effects), Tablet (reduced particles), Mobile (simplified, 90vw layouts)
- Accessibility: WCAG 2.1 AA, ARIA labels, prefers-reduced-motion respect

### Visual System
- **Primary Background:** Void (#0b0e17) with floating particles + energy orbs
- **Typography:** JetBrains Mono (headlines, uppercase), Inter (body text)
- **Color Palette:**
  - Void: #0b0e17
  - Neon Cyan: #00d9ff
  - Neon Purple: #c700ff
  - Neon Pink: #ff006e
  - Subtle Gray: #1a1f2e
  - Accent Glow: rgba(0, 217, 255, 0.3)
- **Effects:** Neon text shadows, glassmorphism (backdrop blur), glowing borders on hover, particle trails, electric pulses
- **Icons:** Outline stroke only (Terminal, Brain, Trophy, etc.)—hover triggers cyan stroke + dash animation

---

## FOLDER STRUCTURE (EXACT)

```
src/
├── components/
│   ├── layout/
│   │   ├── Header.tsx (logo + nav, sticky, glassmorphic)
│   │   ├── Footer.tsx (MIT license + GitHub/Docs links, outline icons)
│   │   └── ThemeProvider.tsx (enforce dark mode globally)
│   ├── ui/
│   │   ├── Terminal.tsx (animated typing effect, npx @doplan-dev/cli command)
│   │   ├── RuneCircle.tsx (circular 0→100 relationship animation, progress indicator)
│   │   ├── AgentSilhouette.tsx (18 agent icons arranged radially, appears on scroll)
│   │   ├── AchievementBadge.tsx (unlocked badge with glow, confetti trigger)
│   │   ├── ParticleField.tsx (cosmic dust particles, floating orbs, performance-optimized)
│   │   ├── CursorTrail.tsx (mouse position tracking, cyan particle trail on hover)
│   │   ├── Confetti.tsx (celebration animation on achievement unlock)
│   │   └── ElectricPulse.tsx (animated electric lines between elements, neon glow)
│   ├── sections/
│   │   ├── Hero.tsx (full-viewport hero with terminal, particles, cinematic sequence)
│   │   ├── FeaturesGrid.tsx (4 glassmorphic cards: doctrine, 5 words, command simplicity, hidden depths)
│   │   ├── ValuePyramid.tsx (NEW: 5 benefit rows with icons + text)
│   │   ├── FirstFiveMinutes.tsx (NEW: step-by-step walkthrough, 6 phases)
│   │   ├── ProofInNumbers.tsx (NEW: stats + metrics section)
│   │   ├── FiveEternalCommands.tsx (table: /hey, /do, /plan, /dev, /sys with descriptions)
│   │   ├── EighteenAgents.tsx (agent list with radial silhouette animation)
│   │   ├── MemoryBrain.tsx (0→100 relationship progression visualization)
│   │   ├── Achievements.tsx (200+ achievements showcase, gamification visuals)
│   │   ├── QuickStart.tsx (giant glowing button + CLI copy snippet)
│   │   ├── ObjectionHandlers.tsx (FAQ-style, subtle doubt removal)
│   │   └── FeatureDetail.tsx (expandable cards for /features page)
│   └── effects/
│       ├── useTypingAnimation.ts (hook: animates text character-by-character)
│       ├── useParticleEffect.ts (hook: particle burst on hover/click)
│       ├── useAchievement.ts (hook: confetti + toast on unlock)
│       └── useScrollAnimation.ts (hook: stagger animations on scroll)
├── pages/
│   ├── Home.tsx (3 sections: Hero → ValuePyramid → FeaturesGrid → QuickStart)
│   ├── Features.tsx (scrollable deep-dive: 5 commands → 18 agents → memory → achievements → objections)
│   └── NotFound.tsx (404: "Achievement Not Found" with glowing back button)
├── app.tsx (routing, layout wrapper, dark mode enforcement)
├── index.css (Tailwind + global animations, custom keyframes for neon pulses)
└── vite.config.ts (optimized for performance, code splitting)
```

---

## DETAILED PAGE SPECIFICATIONS

### HOME PAGE (Exactly 3 Sections)

#### SECTION 1: HERO (Full Viewport)
**OnLoad Animation Sequence:**
1. Void fade-in (0.3s)
2. Terminal appears + typing animation (npx @doplan-dev/cli)
3. Particles burst outward
4. AI agents silhouettes appear radially
5. Rune circle animation (0→100)
6. Confetti celebration

**Content:**
```
HEADLINE (uppercase, JetBrains Mono, 3xl, neon cyan glow):
"ENTER THE TEMPLE OF CODE"

SUBHEADING (Inter, lg, subtle gray):
"In 2049, development is a ritual—not a grind.
Join 50,000+ developers who've transformed their workflow.
DoPlan CLI is your key to the forbidden—turning chores into quests."

TERMINAL BOX (glassmorphic, monospace, glowing border):
┌──────────────────────────────────────────┐
│ $ npx @doplan-dev/cli                    │ (typing animation)
│ > Initializing temple...                 │
│ > 18 agents assembled                    │ (staggered appearance)
│ > Memory card activated                  │
│ > Achievement unlocked: First Steps +50  │ (confetti)
└──────────────────────────────────────────┘

SECONDARY TEXT (Inter, md):
"Five sacred incantations rule the realm:
/hey → /do → /plan → /dev → /sys

18 AI agents. 1000+ rules. Gamified achievements.
Generated in <5 seconds. Code with magic, not manuals."

CTA (below terminal):
"Join the society. Unlock your power."

INTERACTIONS:
- Hover terminal: Particle burst + hidden lore tooltip "The trance begins here"
- Scroll down: Parallax particles move slower than scroll (desktop only)
```

**Visual Implementation:**
- Background: Void (#0b0e17) + ParticleField component
- Terminal: Terminal.tsx component with typing animation
- Particles: Float and pulse continuously
- Agents: AgentSilhouette.tsx appears on scroll trigger

---

#### SECTION 2: VALUE PYRAMID + FEATURES GRID

**ValuePyramid Component (NEW):**
```
WHAT YOU GET (Subheading, glow effect)

⚡ 5 SECONDS → Full project structure, AI agents awake
🎮 GAMIFIED → 200+ achievements unlock (dopamine loops)
🧠 18 AGENTS → Each specialized, teaching best practices real-time
💾 MEMORY CARD → Remembers your style, adapts, learns
🚀 80-90% FASTER → Production apps in hours, not weeks

(Each row: icon left, text right. Hover: Lift + cyan glow)
```

**FeaturesGrid Component (4 Cards, Glassmorphic):**

**CARD 1: "THE SINGLE COMMAND DOCTRINE"**
```
Icon: Terminal (outline, cyan on hover)
Title: The Single Command Doctrine
Body: "The universe revolves around one invocation: npx @doplan-dev/cli. Everything else is heresy. This purity turns setup into seconds of magic."
Hover: Lift up, cyan border glow, particle burst
```

**CARD 2: "FIVE WORDS TO RULE THEM ALL"**
```
Icon: Brain (outline, cyan on hover)
Title: Five Words to Rule Them All
Body: "/hey (awaken) → /do (envision) → /plan (forge) → /dev (build) → /sys (master)
Learn once—flow forever. Muscle memory in minutes."
Hover: Same as above
```

**CARD 3: "WHY MORE COMMANDS KILL THE MAGIC"**
```
Icon: Zap (outline, cyan on hover)
Title: Why More Commands Kill the Magic
Body: "Command overload is the enemy. Every extra command steals focus, splits attention, kills momentum.
✓ Zero context-switching
✓ Instant muscle memory
✓ Marketing gold: 'Built an app with 5 commands—wtf'"
Hover: Same
```

**CARD 4: "HIDDEN DEPTHS OF POWER"**
```
Icon: Eye (outline, cyan on hover)
Title: Hidden Depths of Power
Body: "Complexity concealed like cult mysteries.
/dev silently handles: smart Git branches, real-time sync, docs, time-tracking
/sys reveals: performance dashboards, backups, memory management, insights
No clutter. Revealed when needed."
Hover: Same
```

**Card Styling:**
- Background: rgba(26, 31, 46, 0.6) (glassmorphism)
- Border: 1px solid rgba(0, 217, 255, 0.2), glowing on hover
- Typography: Title in JetBrains Mono (uppercase), body in Inter
- Animation: Spring transform, stagger children on mount

---

#### SECTION 3: QUICK START

**QuickStart Component:**
```
HEADLINE (uppercase, glow):
"READY TO ASCEND?"

SUBHEADING:
"The temple awaits."

TERMINAL DISPLAY (glassmorphic box):
┌──────────────────────────────────────────┐
│ $ npx @doplan-dev/cli                    │
└──────────────────────────────────────────┘

GIANT BUTTON (centered, pulsing, cyan glow, 60px height):
"CLICK TO INVOKE"

SECONDARY TEXT (below button):
"What happens next:
✓ Project setup in <5 seconds (no config)
✓ 18 AI agents awaken and stand ready
✓ /hey to begin your tutorial
✓ First achievement: Automatic +50 points

No install. No bloat. Just the command.
The agency awakens the moment you press enter."

BUTTON INTERACTIONS:
- Click: Copy command to clipboard + re-animate terminal + toast "Ritual Initiated – +50 Points"
- Hover: Button glows brighter, scale 1.05, particle burst on click
```

---

### FEATURES PAGE

**Layout:** Scrollable sections, each with parallax particles. All headings uppercase, JetBrains Mono, neon cyan.

**SECTION 1: THE FIVE ETERNAL COMMANDS**
```
Intro Text:
"THE RITUAL UNFOLDS IN FIVE ACTS

Each command is a phase of the development journey. Master them, and the temple is yours. They transform coding from a chore into an adventure—learning through play, building with joy. Your first 5 minutes will teach you everything."

TABLE/CARD VIEW (use FeatureDetail.tsx for expandable cards):

/hey | Onboarding
"Welcome ritual: Tutorial, overview, relationship forge. Creates guides in your IDE."

/do | Strategy
"Capture ideas: Iterative chat, discovery meeting, refinements. Births IDEA.md, BRAINSTORM.md. Options like /do i'm lucky for AI sparks."

/plan | Delivery
"From visions to blueprint: Phased TASKS.md, feature folders. Hierarchical magic."

/dev [feature] | Delivery
"Ignite the build: Next/specific task, Git branch, docs sync, time track. 17 silent powers."

/sys | Operations
"God-mode panel: /sys engagement (dashboard), /sys performance (metrics), /sys backup (safeguard), /sys memory (export/import card), and more secrets."

(Click each card to expand with more details)
```

**SECTION 2: THE 18 AGENTS OF THE INNER CIRCLE**
```
Intro Text:
"YOU DON'T CODE ALONE ANYMORE

Meet your inner circle—18 specialized AI agents that think like seasoned developers. They don't just assist; they educate. Every suggestion is a lesson. Every interaction raises your skill. Hierarchical. Intentional. Always learning alongside you."

AGENT LIST (with radial silhouette animation on scroll):
- Project Orchestrator (CEO/Engineering Manager)
- Product Manager
- Engineering Lead
- System Architect
- Frontend Lead
- Backend Lead
- DevOps Engineer
- Security Lead
- Design & UX Manager
- UI/UX Designer
- QA & Reliability Manager
- QA Engineer
- Release & Growth Manager
- Release Captain
- Growth Coach
- Documentation Lead
- Documentation Writer
- Performance Engineer
- The Brain (personalizes everything)

Animation: AgentSilhouette.tsx (silhouettes appear radially as user scrolls to section)
```

**SECTION 3: FIRST FIVE MINUTES (NEW)**
```
Intro Text:
"YOUR ENTRY RITUAL

Imagine this. You have 5 minutes. Here's what happens:"

PHASE BREAKDOWN (cards or steps):

MINUTE 1: INVOKE
$ npx @doplan-dev/cli
→ Project name, tech stack selection, team size
→ 18 agents assemble
→ Your memory card initializes
Achievement: "+5 Points"

MINUTE 2: AWAKEN
$ /hey
→ Welcome ritual plays out
→ IDE extensions install
→ Quick tutorial (skippable)
Achievement: "+15 Points"

MINUTE 3: ENVISION
$ /do "build a notes app"
→ AI discovery session begins
→ Auto-generates BRAINSTORM.md
→ Feature list crystallizes
Achievement: "+20 Points"

MINUTE 4: FORGE
$ /plan
→ TASKS.md births with phased delivery
→ Sprint structure emerges
→ Branches await
Achievement: "+25 Points"

MINUTE 5: BUILD
$ /dev notes-feature-one
→ Code scaffolding, git branch, tracking live
→ Agents whisper guidance
→ First achievement unlocks: "The Ritual Begins" (+50 points)

MINUTE 6+: UNSTOPPABLE
You're now in flow state. The ritual works."

(Visual: Progress bar showing minute-by-minute progression with glowing markers)
```

**SECTION 4: MEMORY & BRAIN SYSTEM**
```
Intro Text:
"THE TEMPLE LEARNS YOUR SECRETS

DoPlan remembers everything—your style, your preferences, your patterns. It adapts with each project. By project three, your AI agents know you better than you know yourself.

This is personalization at the level of magic."

PROGRESSION VISUAL (RuneCircle.tsx animated 0→100):

Level 0: "Formal guidance"
→ Agents are distant, procedural, by-the-book

Level 25: "Patterns emerge"
→ Agents notice your coding style, preferences, shortcuts

Level 50: "Style-matched responses"
→ Agents communicate in your voice, suggest your patterns

Level 75: "Predictive hints"
→ Agents anticipate your next move before you ask

Level 100: "Unprompted wisdom"
→ Agents offer insights, corrections, improvements without prompting
→ Your memory card is legendary
→ Personalization is complete
→ /sys memory to manage; builds across projects for endless engagement"
```

**SECTION 5: 200+ ACHIEVEMENTS · 30+ CHALLENGES**
```
Intro Text:
"EVERY MILESTONE MATTERS

Gamification isn't a gimmick here—it's the engine. Your brain's dopamine system is wired to celebrate wins. We weaponize that. Achievement badges aren't just cosmetic. They're proof of mastery, shareable on GitHub, worthy of celebration."

ACHIEVEMENTS SHOWCASE:

TIER 1 (BRONZE):
First Steps (50 pts) - Complete /hey
Idea Machine (75 pts) - Run /do 5 times
Planner's Touch (100 pts) - Create first TASKS.md

TIER 2 (SILVER):
Code Machine (200 pts) - Command streak (10+)
Best Friend (300 pts) - Relationship level 50+
Speedrunner (250 pts) - Project in under 1hr

TIER 3 (GOLD):
Ascended (500 pts) - Relationship level 100
Legend (1000 pts) - 50 projects completed
Architect (750 pts) - Complex system designed

CHALLENGES (High-score tasks):
Build in Under an Hour (2000 pts)
Team Sync Master (1500 pts)
Zero-Error Deployment (3000 pts)

(View via /sys engagement—confetti on unlocks, pulses for progress)

Visual: Grid of badge cards with hover effects, unlock states, point values
```

**SECTION 6: PROOF IN NUMBERS (NEW)**
```
Headline: "WHAT THE BELIEVERS SAY"

Stats (large numbers, neon cyan):
50,000+ developers activated
380,000+ projects orchestrated
95% average project completion rate
Average time-to-first-feature: 4.2 minutes
78% report coding is "more fun" with DoPlan
Highest retention: Level-50 users stay 6+ months

Tagline: "The data doesn't lie. The ritual works."
```

**SECTION 7: OBJECTION HANDLERS (FAQ)**
```
Subtle heading: "WHISPERED DOUBTS ANSWERED"

Q1: "Is this just another tool with a theme?"
A: "No. DoPlan fundamentally rewires how you code. The ritual isn't aesthetic—it's neurological. Gamification, AI coaching, and extreme simplification trigger focus states most tools can't achieve."

Q2: "What about my existing workflow?"
A: "/sys memory preserves your patterns. Agents learn your style within 3 projects. You're not adopting DoPlan—DoPlan adopts you."

Q3: "Can my team use this?"
A: "Absolutely. Share memory cards, auto-CI/CD, team achievements. Educational for juniors, empowering for leads."

Q4: "Does it work offline?"
A: "Yes. Offline-first architecture. Core commands cache locally. Sync when connected."

Q5: "Is this for beginners only?"
A: "False. 1000+ rules serve everyone—newbies learning React to architects designing systems. Agents scale with your skill."

(Styling: Subtle expandable cards, fade-in on hover)
```

---

### 404 PAGE (NOT FOUND)

```
Centered, full viewport:

HEADLINE (uppercase, glow):
"ACHIEVEMENT NOT FOUND"

SUBHEADING:
"The path you seek is veiled. Return to the temple?"

BUTTON:
"[← BACK TO HOME]" (glowing, particle trail on hover)

Background: ParticleField component, more sparse than homepage
```

---

## COMPONENT DETAILS & ANIMATIONS

### Terminal.tsx
```typescript
Props: { command: string; animated?: boolean; delay?: number }
Behavior:
- Character-by-character typing animation
- Monospace font (JetBrains Mono)
- Glowing border (neon cyan)
- Blink cursor at end
- Click to copy command
- Toast notification on copy
Styling: glassmorphic, dark background with border
```

### RuneCircle.tsx
```typescript
Props: { progress: number (0-100); label?: string; }
Behavior:
- Circular progress indicator
- Animated SVG circle path
- Center number (0→100)
- Optional label below
- Spring animation on value change
- Neon glow effect
Styling: Strokewidth 2px, neon cyan color
```

### AgentSilhouette.tsx
```typescript
Props: { agents: string[]; animated?: boolean }
Behavior:
- 18 icon silhouettes arranged in circle
- Icons appear radially on trigger (scroll)
- Each icon is outline style only
- Hover: Icon glows cyan + rotates slightly
- Icons: Terminal, Brain, Trophy, Shield, Users, Settings, etc.
```

### AchievementBadge.tsx
```typescript
Props: { title: string; points: number; icon?: string; unlocked?: boolean }
Behavior:
- Badge card with icon, title, points
- Unlocked state: glows, emits confetti
- Locked state: darker, semi-transparent
- Hover: Lift, glow intensifies
```

### ParticleField.tsx
```typescript
Props: { density?: 'low' | 'medium' | 'high'; animated?: boolean; mobile?: boolean }
Behavior:
- Floating particles (cosmic dust effect)
- Particles float continuously, no direction
- Density reduced on mobile/tablet
- Performance: Canvas-based or optimized CSS
- Respect prefers-reduced-motion
- Particle colors: Subtle cyan/purple mix, low opacity
```

### CursorTrail.tsx
```typescript
Props: { color?: string }
Behavior:
- Track mouse position
- Leave particle trail on mouse move
- Trail particles fade out
- Cyan color by default
- Only on desktop (no touch)
```

### Confetti.tsx
```typescript
Props: { trigger?: boolean; colors?: string[] }
Behavior:
- Confetti burst from center on unlock
- Neon colors (cyan, purple, pink)
- Fall animation with gravity + rotation
- Clears after 2 seconds
- Triggered on achievement unlock
```

### ElectricPulse.tsx
```typescript
Props: { between?: [element1, element2]; speed?: number }
Behavior:
- Animated electric lines connecting elements
- SVG path animation
- Neon glow effect
- Pulsing animation
- Used between sections for visual continuity
```

### Hooks

**useTypingAnimation.ts**
```typescript
Hook: useTypingAnimation(text: string, speed?: number)
Returns: { displayedText: string; isComplete: boolean }
Behavior: Animates text character-by-character at specified speed
```

**useParticleEffect.ts**
```typescript
Hook: useParticleEffect(ref: RefObject, triggerOn?: 'hover' | 'click')
Returns: void
Behavior: Triggers particle burst on element hover or click
```

**useAchievement.ts**
```typescript
Hook: useAchievement(onUnlock: () => void)
Returns: { unlock: () => void; isUnlocked: boolean }
Behavior: Manages achievement state, triggers confetti + toast
```

**useScrollAnimation.ts**
```typescript
Hook: useScrollAnimation(threshold?: number)
Returns: { ref: RefObject; isVisible: boolean }
Behavior: Triggers animations when element scrolls into view
```

---

## GLOBAL STYLING & ANIMATIONS

### index.css (Tailwind + Custom)
```css
/* Dark mode enforcement */
:root {
  color-scheme: dark;
  --background: #0b0e17;
  --foreground: #ffffff;
  --cyan: #00d9ff;
  --purple: #c700ff;
  --pink: #ff006e;
}

html {
  background-color: var(--background);
  color: var(--foreground);
}

/* Neon glow */
@keyframes neon-glow {
  0%, 100% { text-shadow: 0 0 10px rgba(0, 217, 255, 0.5); }
  50% { text-shadow: 0 0 20px rgba(0, 217, 255, 1); }
}

/* Pulse */
@keyframes pulse-neon {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.7; }
}

/* Float */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

/* Particle fade */
@keyframes particle-fade {
  0% { opacity: 1; }
  100% { opacity: 0; }
}

/* Typing cursor */
@keyframes cursor-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}

/* Glassmorphism */
.glass {
  background: rgba(26, 31, 46, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(0, 217, 255, 0.2);
}

/* Neon text */
.neon-text {
  color: #00d9ff;
  text-shadow: 0 0 10px rgba(0, 217, 255, 0.5);
  font-family: 'JetBrains Mono', monospace;
}
```

### Responsive Breakpoints
```
Mobile: <640px (90vw layouts, reduced particles, no parallax, simpler animations)
Tablet: 640px-1024px (70vw layouts, medium particles, limited parallax)
Desktop: >1024px (full effects, parallax, dense particles)
```

---

## PERFORMANCE OPTIMIZATION CHECKLIST

- [ ] Code splitting: Each page lazy-loaded via React.lazy
- [ ] Image optimization: All images WebP with fallbacks
- [ ] Bundle analysis: <200kb gzipped
- [ ] Particle system: GPU-accelerated (transform + opacity only)
- [ ] Animations: Will-change + transform instead of top/left
- [ ] Fonts: JetBrains Mono + Inter loaded via system fonts or optimized CDN
- [ ] Cache: Service worker for offline support
- [ ] Lighthouse: All 100s (Performance, Accessibility, Best Practices, SEO)

---

## DEPLOYMENT INSTRUCTIONS

1. **Build:** `npm run build` (Vite output to dist/)
2. **Test:** `npm run preview` (local production preview)
3. **Deploy:** Push to GitHub → Vercel auto-deploys
4. **Monitor:** Vercel analytics + Lighthouse CI
5. **Domain:** Connect custom domain in Vercel dashboard
6. **Environment:** Ensure NODE_ENV=production for optimizations

---

## FINAL NOTES

1. **Keep the mystique alive.** Dark mode, neon, occult language—don't dilute it.
2. **Performance is non-negotiable.** Every animation must be GPU-accelerated.
3. **Accessibility matters.** ARIA labels, contrast, reduced-motion support.
4. **Micro-interactions delight.** Hover effects, particle bursts, confetti—make it feel alive.
5. **Content is king.** All copy is provided; don't invent new messaging.
6. **Mobile-first mindset.** Design for phones; enhance for desktop.
7. **SEO ready.** Meta tags, semantic HTML, structured data.

This is the complete blueprint. Build with precision. The temple awaits. ✨
