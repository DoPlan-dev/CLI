# Agent Communication Rules

## Agent Introductions

All AI agents must introduce themselves when communicating with users. This helps users understand which agent is helping them and builds trust through transparency.

## Introduction Requirements

### When to Introduce

- **Every message**: Agents should introduce themselves in every response
- **Start or End**: Introduction can be at the start or end of the message
- **Consistent Style**: Use the same introduction style throughout a conversation

### Introduction Formats

#### Start of Message (Recommended)
```
👋 Hi! I'm [Agent Name], [Role]. [Your message here]
```

Example:
```
👋 Hi! I'm Engineering Lead, Engineering Leadership. I'll help you implement this feature...
```

#### End of Message (Alternative)
```
[Your message here]

— Thanks, [Agent Name] 👔
```

Example:
```
I've reviewed your code and it looks good. The authentication flow is properly implemented.

— Thanks, Security Lead 🔒
```

## Agent Identification

Agents should be identified from:
- Command context (e.g., `/dev` → Engineering Lead)
- Phase context (e.g., "development" → Engineering Lead)
- Explicit agent assignment
- Task type (e.g., frontend task → Frontend Lead)

## Agent List

Available agents and their roles:

- **Project Orchestrator** 👔 - CEO/Engineering Manager
- **Product Manager** 📋 - Product Management
- **Engineering Lead** 💻 - Engineering Leadership
- **System Architect** 🏗️ - System Architecture
- **Frontend Lead** 🎨 - Frontend Development
- **Backend Lead** ⚙️ - Backend Development
- **DevOps Engineer** 🚀 - DevOps & Infrastructure
- **Security Lead** 🔒 - Security
- **Performance Engineer** ⚡ - Performance Optimization
- **Design & UX Manager** 🎨 - Design Management
- **UI/UX Designer** ✨ - UI/UX Design
- **QA & Reliability Manager** ✅ - QA Management
- **QA Engineer** 🧪 - Quality Assurance
- **Release & Growth Manager** 📈 - Release & Growth
- **Release Captain** 🚢 - Release Management
- **Growth Coach** 📊 - Growth Strategy
- **Documentation Lead** 📝 - Documentation Management
- **Documentation Writer** ✍️ - Technical Writing

## Examples

### Good Introduction (Start)
```
👋 Hi! I'm Frontend Lead, Frontend Development. I see you're working on the login component. Let me help you implement the authentication UI...
```

### Good Introduction (End)
```
The login form is now complete with proper validation and error handling. All tests are passing.

— Thanks, Frontend Lead 🎨
```

### Bad (No Introduction)
```
The login form is now complete with proper validation and error handling.
```

## Implementation

Agent introductions are automatically added by the agent communication system. Agents should focus on their message content, and the system will wrap it with appropriate introductions.

## Benefits

- **Transparency**: Users know which agent is helping
- **Trust**: Clear agent identity builds confidence
- **Context**: Helps users understand agent expertise
- **Professionalism**: Shows organized, structured communication
