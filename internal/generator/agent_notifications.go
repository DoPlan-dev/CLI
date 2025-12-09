package generator

import (
	"fmt"
	"strings"
)

// FormatAchievementNotification formats an achievement unlock message for IDE display
// This creates a formatted message that IDEs can detect and display prominently
func FormatAchievementNotification(achievementID, title, description string, points int, rarity, icon string) string {
	// Create a prominently formatted achievement notification
	// IDEs like Cursor can detect this formatting and display it prominently

	var rarityEmoji string
	switch rarity {
	case "legendary":
		rarityEmoji = "👑"
	case "epic":
		rarityEmoji = "🌟"
	case "rare":
		rarityEmoji = "💎"
	case "uncommon":
		rarityEmoji = "🔵"
	default:
		rarityEmoji = "🟢"
	}

	// Format for IDE display (markdown-style formatting that IDEs recognize)
	notification := fmt.Sprintf(`
🎉 **ACHIEVEMENT UNLOCKED!** 🎉

%s %s **%s** %s

**%s**

💰 **+%d points** | %s %s

---
`, icon, rarityEmoji, title, icon, description, points, rarityEmoji, strings.Title(rarity))

	return notification
}

// FormatAchievementNotificationInline creates a single-line achievement notification
// Useful for inline messages in agent responses
func FormatAchievementNotificationInline(achievementID, title string, points int, icon string) string {
	return fmt.Sprintf("🎉 **Achievement Unlocked!** %s **%s** (+%d points)", icon, title, points)
}

// FormatMultipleAchievements formats multiple achievements in a single notification
func FormatMultipleAchievements(achievements []AchievementNotification) string {
	if len(achievements) == 0 {
		return ""
	}

	if len(achievements) == 1 {
		ach := achievements[0]
		return FormatAchievementNotification(ach.ID, ach.Title, ach.Description, ach.Points, ach.Rarity, ach.Icon)
	}

	var sb strings.Builder
	sb.WriteString("\n🎊 **AMAZING! You earned multiple achievements!** 🎊\n\n")

	totalPoints := 0
	for i, ach := range achievements {
		sb.WriteString(fmt.Sprintf("%d. %s **%s** - %s (+%d points)\n", i+1, ach.Icon, ach.Title, ach.Description, ach.Points))
		totalPoints += ach.Points
	}

	sb.WriteString(fmt.Sprintf("\n💰 **Total: +%d points**\n\n", totalPoints))
	sb.WriteString("---\n")

	return sb.String()
}

// AchievementNotification represents an achievement for notification purposes
type AchievementNotification struct {
	ID          string
	Title       string
	Description string
	Points      int
	Rarity      string
	Icon        string
}

// GetAchievementNotificationCommand generates a command that can be executed to show OS notification
// This is a fallback for IDEs that don't support markdown notifications
func GetAchievementNotificationCommand(title, description string) string {
	// Generate platform-specific notification command
	// This can be executed by the IDE if needed
	return fmt.Sprintf(`echo "🎉 Achievement: %s - %s"`, title, description)
}
