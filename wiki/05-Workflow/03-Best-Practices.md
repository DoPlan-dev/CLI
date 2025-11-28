# Best Practices

Tips and best practices for getting the most out of DoPlan. Optimize your workflow, maximize efficiency, and build a strong relationship with DoPlan.

---

## 🎯 General Best Practices

### 1. Complete Onboarding First Time

**Why**: Sets up your preferences and builds foundation
**How**: 
- Run `/hey` when you first use DoPlan
- Complete the interactive tutorial
- Set your experience level accurately
- Review reference materials

**Benefit**: Better personalization from the start

---

### 2. Use Iterative Conversation

**Why**: Captures complete ideas and improves plan quality
**How**:
- Use `/do` for full ideation workflow
- Provide multiple rounds of input
- Don't rush - add details as you think of them
- Type "done" when truly finished

**Benefit**: Better plans, fewer revisions

---

### 3. Review Plan Before Development

**Why**: Ensures plan is correct and complete
**How**:
- Read TASKS.md after `/plan`
- Check phase structure
- Verify task dependencies
- Review feature documentation

**Benefit**: Fewer surprises during development

---

### 4. Always Use `/done` When Complete

**Why**: Tracks completion, commits changes, checks achievements
**How**:
- Run `/done` after every task
- Don't skip - even for small tasks
- Let DoPlan handle Git operations
- Check achievements earned

**Benefit**: Automatic tracking, achievements, Git automation

---

### 5. Check Progress Regularly

**Why**: Stay aware of progress and plan next steps
**How**:
- Run `/status` after completing tasks
- Check before starting new work
- Monitor completion percentage
- Track state changes

**Benefit**: Better planning, awareness of progress

---

## ⚡ Efficiency Tips

### 1. Use Fast-Track When Appropriate

**When**: You have a detailed prompt or PRD
**How**:
```bash
/do now --prompt "Build a todo app with React and Node.js"
```

**Benefit**: Skips meeting/refinement, saves time

---

### 2. Use Feature Command for Single Features

**When**: Adding a feature to existing project
**How**:
```bash
/do feature
```

**Benefit**: Quick feature addition, no full workflow

---

### 3. Use Lucky Mode for Inspiration

**When**: You want AI-suggested ideas
**How**:
```bash
/do i'm lucky
```

**Benefit**: Creative ideas, learning from rejections

---

### 4. Check Engagement for Motivation

**When**: Need motivation or want to see progress
**How**:
```bash
/sys engagement
```

**Benefit**: See achievements, stay motivated

---

## 🎯 Workflow Optimization

### 1. Batch Similar Tasks

**Why**: More efficient development
**How**:
- Group related tasks together
- Complete similar features in sequence
- Reduce context switching

**Benefit**: Faster development, better flow

---

### 2. Review Documentation Before Coding

**Why**: Understand requirements before starting
**How**:
- Read feature documentation in `.do/plan/`
- Review design and plan files
- Check task dependencies
- Understand context

**Benefit**: Fewer mistakes, better code

---

### 3. Use State History for Reference

**Why**: See what changed and when
**How**:
- Check `.do/system/history/state-*.json`
- Review state deltas in `/status`
- Understand project evolution

**Benefit**: Better project understanding

---

### 4. Track Time for Analysis

**Why**: Understand where time goes
**How**:
- Review `.do/system/time-tracker.jsonl`
- Analyze task durations
- Identify time sinks
- Optimize workflow

**Benefit**: Better time management

---

## 🎓 Learning Best Practices

### 1. Set Learning Goals

**Why**: Tracks what you want to learn
**How**:
- Set goals in `/hey` or through conversation
- Update goals as you learn
- Track progress in Memory Card

**Benefit**: Focused learning, achievement tracking

---

### 2. Explore New Technologies

**Why**: Expands your skills
**How**:
- Try new technologies in projects
- DoPlan tracks your tech stack
- Earn "Tech Explorer" achievements
- Build diverse experience

**Benefit**: Skill growth, achievement earning

---

### 3. Learn from Pain Points

**Why**: Improves your skills
**How**:
- DoPlan remembers where you struggle
- Provides extra help in those areas
- Tracks pain points in Memory Card
- Offers targeted assistance

**Benefit**: Targeted learning, skill improvement

---

### 4. Use Educational Explanations

**Why**: Understand why, not just how
**How**:
- Set experience level accurately
- Ask for explanations
- Review detailed responses
- Learn best practices

**Benefit**: Deeper understanding, better skills

---

## 🎮 Engagement Best Practices

### 1. Complete Challenges for High Scores

**Why**: Challenges award 300-2000 points
**How**:
- Complete first-time tasks
- Deploy to staging/production
- Integrate APIs
- Achieve test coverage

**Benefit**: High scores, exciting rewards

---

### 2. Build Achievement Streaks

**Why**: Consistency achievements
**How**:
- Use DoPlan regularly
- Complete projects consistently
- Build daily streaks
- Maintain momentum

**Benefit**: Streak achievements, relationship growth

---

### 3. Track Your Progress

**Why**: See growth over time
**How**:
- Check `/sys engagement` regularly
- Monitor score milestones
- Track relationship level
- View achievement progress

**Benefit**: Motivation, progress awareness

---

### 4. Celebrate Achievements

**Why**: Stay motivated
**How**:
- View achievements when earned
- Check engagement dashboard
- Track next milestones
- Celebrate wins

**Benefit**: Motivation, engagement

---

## 🔒 Safety Best Practices

### 1. Verify Branch Before `/done`

**Why**: Prevents committing to wrong branch
**How**:
- DoPlan checks automatically
- Warns if on main/master
- Confirms before proceeding

**Benefit**: Safety, correct commits

---

### 2. Check Dependencies

**Why**: Ensures tasks are ready
**How**:
- DoPlan checks automatically
- Lists missing dependencies
- Blocks completion if needed

**Benefit**: Correct order, fewer issues

---

### 3. Use State History

**Why**: Rollback capability
**How**:
- State snapshots created automatically
- Review state history
- Use for rollback if needed

**Benefit**: Safety net, project recovery

---

### 4. Review Commits

**Why**: Ensure correct commit messages
**How**:
- DoPlan uses conventional format
- Review before pushing
- Verify changes

**Benefit**: Clean Git history

---

## 💡 Communication Best Practices

### 1. Provide Complete Information

**Why**: Better plans and results
**How**:
- Use iterative conversation in `/do`
- Add details as you think of them
- Don't rush the process
- Be thorough

**Benefit**: Better outcomes, fewer revisions

---

### 2. Set Preferences Accurately

**Why**: Better personalization
**How**:
- Set experience level correctly
- Choose work style accurately
- Set communication preferences
- Update as you learn

**Benefit**: Better personalization, appropriate responses

---

### 3. Use Appropriate Commands

**Why**: Efficient workflow
**How**:
- Use `/do` for full ideation
- Use `/do now` for fast-tracking
- Use `/do feature` for single features
- Use `/do i'm lucky` for inspiration

**Benefit**: Efficient workflow, right tool for job

---

## 🚀 Advanced Best Practices

### 1. Leverage Time Tracking

**Why**: Understand productivity
**How**:
- Review time-tracker.jsonl
- Analyze task durations
- Identify patterns
- Optimize workflow

**Benefit**: Better time management

---

### 2. Use State Management

**Why**: Project understanding
**How**:
- Review active_state.json
- Check state history
- Understand project state
- Use for rollback

**Benefit**: Better project control

---

### 3. Customize System Settings

**Why**: Tailor to your needs
**How**:
- Use `/sys role` for permissions
- Use `/sys control` for features
- Configure security settings
- Customize engagement

**Benefit**: Personalized experience

---

### 4. Build Long-Term Relationship

**Why**: Better personalization
**How**:
- Use DoPlan regularly
- Complete projects
- Earn achievements
- Build relationship level

**Benefit**: Maximum personalization, best experience

---

## 🎯 Common Mistakes to Avoid

### 1. Skipping Onboarding

**Mistake**: Not running `/hey` first time
**Fix**: Always complete onboarding
**Benefit**: Better setup, personalization

---

### 2. Rushing Ideation

**Mistake**: Not using iterative conversation
**Fix**: Take time, provide details
**Benefit**: Better plans, fewer revisions

---

### 3. Forgetting `/done`

**Mistake**: Not running `/done` after tasks
**Fix**: Always use `/done` when complete
**Benefit**: Tracking, achievements, Git automation

---

### 4. Not Checking Progress

**Mistake**: Not using `/status`
**Fix**: Check regularly
**Benefit**: Awareness, better planning

---

### 5. Ignoring Engagement

**Mistake**: Not viewing achievements
**Fix**: Check `/sys engagement` regularly
**Benefit**: Motivation, progress tracking

---

## 💡 Pro Tips

### For Beginners
- Complete tutorial thoroughly
- Set experience level accurately
- Use iterative conversation
- Ask for explanations
- Check progress regularly

### For Intermediate
- Use fast-track when appropriate
- Complete challenges for high scores
- Build achievement streaks
- Track time for analysis
- Explore new technologies

### For Advanced
- Leverage all features
- Customize system settings
- Build long-term relationship
- Use state management
- Optimize workflow

---

**Next**: [Common Patterns](./04-Common-Patterns.md)

