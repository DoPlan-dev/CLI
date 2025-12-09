# Achievement Graphics

This directory contains PNG graphics for all DoPlan achievements.

## Generating Graphics

To generate achievement graphics, run:

```bash
python3 scripts/generate_achievement_graphics.py
```

### Requirements

- Python 3.6+
- Pillow library: `pip install Pillow`

### Output

Graphics are generated as PNG files named `{achievement_id}.png` (e.g., `first_project.png`, `score_1000.png`).

## Graphics Design

Each achievement graphic follows a consistent design system:

- **Common** 🟢 - Simple badge with gray background
- **Uncommon** 🔵 - Enhanced border with blue accent
- **Rare** 🟣 - Glow effect with purple accent
- **Epic** 🟠 - Animated-style frame with orange/gold accent
- **Legendary** 🔴 - Premium gold/platinum style with special effects

## Manual Generation

If you prefer to create graphics manually or use a different tool:

1. Use the achievement list in `../ACHIEVEMENTS.md`
2. Create 400x300px PNG images
3. Include the achievement icon, title, points, and rarity
4. Name files using the achievement ID

## Usage

Graphics are referenced in:
- Achievement documentation
- IDE notifications (when implemented)
- Engagement dashboard (future enhancement)
