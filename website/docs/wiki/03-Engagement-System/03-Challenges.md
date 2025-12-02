# Challenges Guide

Challenges are **high-scoring rewards** (300-2000 points) for completing important first-time tasks. They're designed to make you **extremely excited** to finish them!

---

## 🎯 What Are Challenges?

Challenges reward you with **high scores** for completing significant first-time tasks:
- **300-2000 points** each (much higher than regular achievements)
- **First-time focus** - Especially rewarding for first attempts
- **Automatic detection** - No manual claiming needed
- **Celebrated achievements** - Special recognition for completion

---

## 🏆 Challenge Categories

### 1. Integration Challenges (4+ Challenges)

Rewards for API and integration work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| API Integration Master | Generate and integrate first API | 500 | Epic | 🔌 |
| Integration Tested | Add API and pass all tests | 750 | Epic | ✅ |
| Third-Party Connector | Integrate with third-party service | 600 | Epic | 🔗 |
| Webhook Wizard | Set up first webhook | 400 | Rare | 🪝 |

**How to Earn**: Create APIs, integrate services, set up webhooks.

---

### 2. Database Challenges (4+ Challenges)

Rewards for database work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| Database Connected | Link project to database | 500 | Epic | 🗄️ |
| Data Migration Master | Perform first migration | 400 | Uncommon | ➡️ |
| Database Merger | Merge data from two sources | 600 | Rare | ➕ |
| Backup Strategist | Implement backup solution | 700 | Epic | 💾 |

**How to Earn**: Connect databases, run migrations, merge data, set up backups.

---

### 3. Deployment Challenges (5+ Challenges)

Rewards for deployment work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| First Launch! | Deploy to staging first time | 1,000 | Legendary | 🚀 |
| Production Ready | Deploy to production | 1,500 | Legendary | 🌐 |
| CI/CD Pioneer | Set up CI/CD pipeline | 800 | Epic | ⚙️ |
| Docker Deployer | Containerize and deploy | 700 | Rare | 🐳 |
| Kubernetes Captain | Deploy to Kubernetes | 1,000 | Epic | ☸️ |

**How to Earn**: Deploy to staging/production, set up CI/CD, use containers.

---

### 4. Testing Challenges (5+ Challenges)

Rewards for testing work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| Test Coverage Champion | Achieve 80%+ coverage | 700 | Epic | 📊 |
| E2E Master | Implement E2E test suite | 600 | Rare | 🌐 |
| Performance Tester | Conduct performance tests | 500 | Rare | ⚡ |
| Bug Squasher | Fix 10 bugs | 400 | Uncommon | 🐛 |
| Flawless Code | All tests passing | 750 | Epic | ✅ |

**How to Earn**: Write tests, achieve coverage, fix bugs, pass all tests.

---

### 5. Workflow Challenges (4+ Challenges)

Rewards for workflow and Git work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| GitHub Workflow Master | Implement robust workflow | 800 | Epic | 🌳 |
| Conventional Committer | 20 conventional commits | 300 | Rare | 📝 |
| Code Reviewer | First PR approved | 400 | Uncommon | 👀 |
| Branch Strategist | Manage 5 branches concurrently | 500 | Rare | 🌿 |

**How to Earn**: Set up workflows, use conventional commits, manage branches.

---

### 6. Release Challenges (4+ Challenges)

Rewards for release work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| Public Launch! | First public release | 2,000 | Legendary | 🚀 |
| Version 1.0 | Release v1.0.0 | 1,500 | Legendary | 🥇 |
| Release Notes Artisan | Generate release notes | 700 | Epic | 📝 |
| Hotfix Hero | Deploy critical hotfix | 1,200 | Epic | 🚑 |

**How to Earn**: Make public releases, release v1.0, create release notes.

---

### 7. Performance Challenges (3+ Challenges)

Rewards for performance work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| Performance Optimizer | 20%+ performance improvement | 800 | Epic | ⚡ |
| Load Balancer | Implement load balancing | 700 | Rare | ⚖️ |
| Caching Expert | Implement caching strategy | 600 | Rare | 📦 |

**How to Earn**: Optimize performance, implement load balancing, add caching.

---

### 8. Security Challenges (3+ Challenges)

Rewards for security work:

| Challenge | Requirement | Points | Rarity | Icon |
|-----------|-------------|--------|--------|------|
| Security Auditor | Complete security audit | 1,000 | Legendary | 🔒 |
| Vulnerability Fixer | Fix 5 critical vulnerabilities | 800 | Epic | 🩹 |
| RBAC Implementer | Implement RBAC | 700 | Rare | 🔑 |

**How to Earn**: Run security audits, fix vulnerabilities, implement security.

---

## 🎯 How Challenges Work

### Automatic Detection

Challenges are **automatically detected** when you:
- Complete first-time tasks
- Reach specific milestones
- Achieve quality metrics
- Implement workflows properly

### Context-Based Detection

Challenges use context from commands to detect completion:

**Example - API Integration**:
```go
context := map[string]interface{}{
    "api_created": true,
    "api_tested": true,
    "tests_passed": true,
}
```

**Result**: 
- 🏆 "API Integration Master" (500 points)
- ✅ "Integration Tested" (750 points)
- **Total: 1,250 points!**

### First-Time Focus

Challenges are especially rewarding for **first-time** tasks:
- First API integration
- First deployment
- First database connection
- First public release

**This makes trying new things exciting!**

---

## 💰 Challenge Points

### Point Ranges

- **Rare Challenges**: 300-500 points
- **Epic Challenges**: 500-1000 points
- **Legendary Challenges**: 1000-2000 points

### High Scores

Challenges award **much higher** points than regular achievements:
- Regular achievement: 10-100 points
- Challenge: 300-2000 points

**Example**: Complete "Public Launch" challenge → 2,000 points!

---

## 🎉 Challenge Celebration

### Single Challenge

```
🏆  CHALLENGE COMPLETED!  🏆
  🔌  API Integration Master  🔌
  Generate and integrate your first API
  💰 Points Earned: +500
  ⭐ Rarity: epic
  ⚡ Completed on first try - Amazing!
```

### Multiple Challenges

```
🚀🚀🚀  INCREDIBLE! You earned multiple rewards!  🚀🚀🚀

  1. 🔌 API Integration Master (Challenge) (+500 points)
  2. ✅ Integration Tested (Challenge) (+750 points)
  3. 🎯 On the Rise (Achievement) (+25 points)

  💰 Total Points: +1,275
  📊 New Score: 1,275
```

---

## 🎯 Challenge Examples

### Example 1: First Deployment

**What you do**:
1. Deploy to staging environment
2. Run `/done` to complete task

**What happens**:
- 🏆 "First Launch!" challenge detected (1,000 points)
- Score increases significantly
- Special celebration displayed
- Achievement unlocked!

### Example 2: API Integration

**What you do**:
1. Create API endpoint
2. Write tests
3. All tests pass
4. Run `/done`

**What happens**:
- 🏆 "API Integration Master" (500 points)
- ✅ "Integration Tested" (750 points)
- **Total: 1,250 points!**
- Multiple rewards celebration

### Example 3: Public Release

**What you do**:
1. Release project publicly
2. Version 1.0.0

**What happens**:
- 🏆 "Public Launch!" (2,000 points)
- 🥇 "Version 1.0" (1,500 points)
- **Total: 3,500 points!**
- Legendary celebration

---

## 📊 Tracking Challenges

### View Completed Challenges

```bash
/sys engagement
```

**Shows**:
- Challenges completed count
- Total score from challenges
- Next challenge opportunities

### Challenge Progress

The engagement dashboard shows:
- Completed challenges
- Challenge categories explored
- High-scoring challenges earned
- Next challenge hints

---

## 🎓 Challenge Tips

### For Beginners
- **Start with integration** - API challenges are achievable
- **Focus on testing** - Testing challenges reward good practices
- **Try deployment** - First deployment is highly rewarding

### For Intermediate
- **Complete workflows** - Workflow challenges reward best practices
- **Achieve coverage** - Testing coverage challenges
- **Optimize performance** - Performance challenges

### For Advanced
- **Aim for legendary** - Public launch, production deployment
- **Security focus** - Security challenges are high-scoring
- **Complete all categories** - Explore all challenge types

---

## 🚀 Earning Your First Challenge

### Quick Wins

1. **Connect database** → "Database Connected" (500 points)
2. **Create API** → "API Integration Master" (500 points)
3. **Write tests** → "Integration Tested" (750 points)
4. **Deploy to staging** → "First Launch!" (1,000 points)

**Total**: 2,750+ points from first challenges!

---

## 💡 Pro Tips

- **Complete challenges** for highest scores (300-2000 points)
- **First-time tasks** are especially rewarding
- **Multiple challenges** can be earned together
- **Check `/sys engagement`** to see challenge progress
- **Aim for legendary** challenges for maximum points

---

**Next**: [Score System](./04-Score-System.md)

