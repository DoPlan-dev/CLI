# Brainstorm Interview Templates

This directory contains phase-based interview templates for the `/improve` command.

## How It Works

When a user types `/improve`, the AI agents conduct a structured 6-phase discovery interview:

1. **Phase 01: Vision & Outcomes** - Understanding the problem and success metrics
2. **Phase 02: Audience & Differentiation** - Defining users and competitive positioning
3. **Phase 03: Experience, UI/UX & Tech** - Design and technical expectations
4. **Phase 04: Content & SEO** - Content strategy and search optimization
5. **Phase 05: Marketing & Growth** - Go-to-market and growth strategy
6. **Phase 06: Delivery, Ops & Risks** - Launch planning and risk management

Each phase template (`phase-*.md`) contains the questions that will be asked during that phase.

## Customizing the Interview

### Edit Existing Phases

To modify questions in a phase, edit the corresponding `phase-*.md` file:

```markdown
# Phase 01 · Vision & Outcomes

- What problem are we solving and why now?
- [Your custom question here]
- What does success look like in 3, 6, and 12 months?
```

### Add New Phases

1. Create a new file: `phase-07-your-phase.md`
2. Follow the naming convention: `phase-XX-name.md`
3. Add your questions in bullet format
4. The `/improve` command will automatically include it (phases are loaded in order)

### Remove Phases

Simply delete the `phase-*.md` file you don't need. The interview will skip that phase.

### Reorder Phases

Rename files to change the order (e.g., `phase-01-` comes before `phase-02-`).

## Template Format

Each phase template should follow this format:

```markdown
# Phase XX · Phase Name

- Question 1?
- Question 2?
- Question 3?
```

**Tips:**
- Keep questions concise and clear
- One question per bullet point
- Questions should be open-ended (not yes/no)
- Order questions logically (general → specific)

## Output Format

After the interview, answers are compiled into `BRAINSTORM.md` using the structure in `TEMPLATE_BRAINSTORM.md`. You can customize that template to change how the final document is structured.

## Examples

### Domain-Specific Customization

**For E-commerce Projects:**
- Add questions about payment processing, inventory management, shipping
- Create `phase-07-ecommerce.md` with domain-specific questions

**For SaaS Projects:**
- Add questions about subscription models, user onboarding, feature gating
- Modify Phase 05 to focus on SaaS growth metrics

**For Content Sites:**
- Expand Phase 04 with detailed content strategy questions
- Add questions about CMS requirements, editorial workflows

## Best Practices

1. **Keep It Conversational**: Questions should feel natural, not like a form
2. **Probe Deeper**: Agents will ask follow-ups, but your base questions should be clear
3. **Domain Relevance**: Customize phases to match your project type
4. **Test Your Changes**: Run `/improve` after modifying templates to ensure they work as expected

## Files in This Directory

- `phase-01-vision.md` - Vision and outcomes questions
- `phase-02-audience.md` - Audience and differentiation questions
- `phase-03-experience.md` - Experience, UI/UX, and tech questions
- `phase-04-content.md` - Content and SEO questions
- `phase-05-marketing.md` - Marketing and growth questions
- `phase-06-delivery.md` - Delivery, operations, and risks questions
- `TEMPLATE_BRAINSTORM.md` - Output format template
- `README.md` - This file

---

**Note**: Changes to templates take effect immediately. No code changes needed!
