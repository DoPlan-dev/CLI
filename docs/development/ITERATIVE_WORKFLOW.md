# Iterative Development Workflow

## 🎯 Overview

This document outlines our collaborative development workflow for building new projects. We'll work iteratively, with feedback loops at each step.

## 📋 Workflow Process

### Phase 1: Idea Selection & Planning
1. **Choose Idea** → Share your app idea
2. **Refine Together** → Brainstorm and refine the concept
3. **Create Project** → Use DoPlan CLI to generate project structure
4. **Initial Scan** → I scan the generated project files
5. **Feedback Round 1** → You provide feedback, I analyze and suggest improvements

### Phase 2: Iterative Development

**For each development step:**

```
┌─────────────────────────────────────────┐
│ 1. You complete a step/feature          │
│    (code, implement, test, etc.)        │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│ 2. You request: "scan project"          │
│    I analyze all project files          │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│ 3. I provide feedback:                  │
│    - Code quality analysis              │
│    - Architecture review                │
│    - Suggestions for improvements       │
│    - Potential issues/risks             │
│    - Next steps recommendations         │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│ 4. You review feedback and provide:     │
│    - Your own feedback                  │
│    - Questions or concerns              │
│    - Decisions on suggested changes     │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│ 5. You proceed to next step              │
│    (implement, fix, add features, etc.) │
└──────────────┬──────────────────────────┘
               │
               ▼
         [Repeat cycle]
```

### Phase 3: Completion & Planning

1. **Project Complete** → All core features implemented
2. **Final Scan** → Comprehensive project review
3. **Feedback Collection** → Gather all feedback from development process
4. **Create Update Plan** → Based on feedback, create roadmap for:
   - Bug fixes
   - Feature enhancements
   - Performance improvements
   - Next release planning

## 🔍 What I'll Scan & Analyze

When you request a project scan, I'll review:

### Code Quality
- Code structure and organization
- Best practices adherence
- Potential bugs or issues
- Code duplication
- Performance concerns

### Architecture
- Project structure
- Design patterns usage
- Scalability considerations
- Separation of concerns
- Dependencies management

### Documentation
- README completeness
- Code comments
- API documentation
- Setup instructions

### Testing
- Test coverage
- Test quality
- Missing test cases

### Configuration
- Build setup
- CI/CD workflows
- Environment configuration
- Dependencies

### Security
- Security best practices
- Potential vulnerabilities
- Authentication/authorization
- Data handling

## 📝 Feedback Format

When I provide feedback, I'll structure it as:

### ✅ What's Working Well
- Positive observations
- Good practices identified
- Strong implementations

### ⚠️ Areas for Improvement
- Issues found
- Suggestions for refactoring
- Best practice recommendations

### 🚀 Next Steps
- Recommended actions
- Priority items
- Suggested features

### ❓ Questions
- Clarifications needed
- Decisions required
- Alternative approaches

## 🎯 Project Tracking

### Progress Checklist
- [ ] Idea selected and refined
- [ ] Project structure created
- [ ] Initial scan completed
- [ ] Core features implemented
- [ ] Testing completed
- [ ] Documentation updated
- [ ] Final review completed
- [ ] Update plan created

### Feedback Log
Keep track of feedback at each step:
- **Step 1**: [Date] - [Feedback summary]
- **Step 2**: [Date] - [Feedback summary]
- ...

## 💡 Tips for Effective Collaboration

1. **Be Specific**: When giving feedback, be specific about what you like/dislike
2. **Ask Questions**: Don't hesitate to ask for clarification
3. **Share Context**: If something doesn't work, share error messages or logs
4. **Prioritize**: Let me know what's most important to you
5. **Iterate Quickly**: Small, frequent iterations work better than large changes

## 🚀 Getting Started

1. **Share your idea** → Tell me what you want to build
2. **Create project** → Use `npx @doplan-dev/cli` to generate structure
3. **First scan** → Say "scan project" and I'll analyze everything
4. **Start building** → Implement features step by step
5. **Iterate** → Repeat the scan → feedback → implement cycle

---

**Ready to start?** Share your idea and let's begin! 🎉

