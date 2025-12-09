package cli

import (
	"fmt"
	"strings"
)

// AgentInfo represents information about an agent
type AgentInfo struct {
	Name        string
	Role        string
	Emoji       string
	DisplayName string
}

// GetAgentInfo returns agent information by name
func GetAgentInfo(agentName string) *AgentInfo {
	// Map of agent names to their info
	agentMap := map[string]AgentInfo{
		"Project Orchestrator": {
			Name:        "Project Orchestrator",
			Role:        "CEO/Engineering Manager",
			Emoji:       "👔",
			DisplayName: "Project Orchestrator",
		},
		"Product Manager": {
			Name:        "Product Manager",
			Role:        "Product Management",
			Emoji:       "📋",
			DisplayName: "Product Manager",
		},
		"Engineering Lead": {
			Name:        "Engineering Lead",
			Role:        "Engineering Leadership",
			Emoji:       "💻",
			DisplayName: "Engineering Lead",
		},
		"System Architect": {
			Name:        "System Architect",
			Role:        "System Architecture",
			Emoji:       "🏗️",
			DisplayName: "System Architect",
		},
		"Frontend Lead": {
			Name:        "Frontend Lead",
			Role:        "Frontend Development",
			Emoji:       "🎨",
			DisplayName: "Frontend Lead",
		},
		"Backend Lead": {
			Name:        "Backend Lead",
			Role:        "Backend Development",
			Emoji:       "⚙️",
			DisplayName: "Backend Lead",
		},
		"DevOps Engineer": {
			Name:        "DevOps Engineer",
			Role:        "DevOps & Infrastructure",
			Emoji:       "🚀",
			DisplayName: "DevOps Engineer",
		},
		"Security Lead": {
			Name:        "Security Lead",
			Role:        "Security",
			Emoji:       "🔒",
			DisplayName: "Security Lead",
		},
		"Performance Engineer": {
			Name:        "Performance Engineer",
			Role:        "Performance Optimization",
			Emoji:       "⚡",
			DisplayName: "Performance Engineer",
		},
		"Design & UX Manager": {
			Name:        "Design & UX Manager",
			Role:        "Design Management",
			Emoji:       "🎨",
			DisplayName: "Design & UX Manager",
		},
		"UI/UX Designer": {
			Name:        "UI/UX Designer",
			Role:        "UI/UX Design",
			Emoji:       "✨",
			DisplayName: "UI/UX Designer",
		},
		"QA & Reliability Manager": {
			Name:        "QA & Reliability Manager",
			Role:        "QA Management",
			Emoji:       "✅",
			DisplayName: "QA & Reliability Manager",
		},
		"QA Engineer": {
			Name:        "QA Engineer",
			Role:        "Quality Assurance",
			Emoji:       "🧪",
			DisplayName: "QA Engineer",
		},
		"Release & Growth Manager": {
			Name:        "Release & Growth Manager",
			Role:        "Release & Growth",
			Emoji:       "📈",
			DisplayName: "Release & Growth Manager",
		},
		"Release Captain": {
			Name:        "Release Captain",
			Role:        "Release Management",
			Emoji:       "🚢",
			DisplayName: "Release Captain",
		},
		"Growth Coach": {
			Name:        "Growth Coach",
			Role:        "Growth Strategy",
			Emoji:       "📊",
			DisplayName: "Growth Coach",
		},
		"Documentation Lead": {
			Name:        "Documentation Lead",
			Role:        "Documentation Management",
			Emoji:       "📝",
			DisplayName: "Documentation Lead",
		},
		"Documentation Writer": {
			Name:        "Documentation Writer",
			Role:        "Technical Writing",
			Emoji:       "✍️",
			DisplayName: "Documentation Writer",
		},
	}

	// Try exact match first
	if info, ok := agentMap[agentName]; ok {
		return &info
	}

	// Try case-insensitive match
	for name, info := range agentMap {
		if strings.EqualFold(name, agentName) {
			return &info
		}
	}

	// Default agent info if not found
	return &AgentInfo{
		Name:        agentName,
		Role:        "AI Assistant",
		Emoji:       "🤖",
		DisplayName: agentName,
	}
}

// FormatAgentIntroduction formats an agent introduction message
// style: "start" for beginning of message, "end" for end of message
func FormatAgentIntroduction(agentName, message, style string) string {
	info := GetAgentInfo(agentName)

	if style == "end" {
		// End of message: "Thanks, {Agent Name}"
		return fmt.Sprintf("%s\n\n— Thanks, %s %s", message, info.Emoji, info.DisplayName)
	}

	// Start of message: "Hi! I'm {Agent Name}, {Role}. {Message}"
	return fmt.Sprintf("%s Hi! I'm **%s**, %s. %s", info.Emoji, info.DisplayName, info.Role, message)
}

// WrapAgentMessage wraps a message with agent introduction
// If agentName is empty, returns message as-is
func WrapAgentMessage(agentName, message string, introAtStart bool) string {
	if agentName == "" {
		return message
	}

	if introAtStart {
		return FormatAgentIntroduction(agentName, message, "start")
	}

	return FormatAgentIntroduction(agentName, message, "end")
}

// IdentifyAgentFromContext attempts to identify the agent from context
// This can be enhanced to detect agent from command, phase, or other context
func IdentifyAgentFromContext(context map[string]interface{}) string {
	// Check for explicit agent name
	if agent, ok := context["agent"].(string); ok && agent != "" {
		return agent
	}

	// Infer from command
	if command, ok := context["command"].(string); ok {
		commandAgentMap := map[string]string{
			"/do":      "Product Manager",
			"/plan":    "Product Manager",
			"/dev":     "Engineering Lead",
			"/done":    "Engineering Lead",
			"/write":   "Documentation Writer",
			"/content": "Documentation Writer",
			"/ship":    "Release Captain",
			"/safe":    "Security Lead",
			"/cheap":   "Performance Engineer",
		}
		if agent, ok := commandAgentMap[command]; ok {
			return agent
		}
	}

	// Infer from phase
	if phase, ok := context["phase"].(string); ok {
		phaseAgentMap := map[string]string{
			"ideation":    "Product Manager",
			"meeting":     "Product Manager",
			"planning":    "Product Manager",
			"development": "Engineering Lead",
			"writing":     "Documentation Writer",
			"design":      "UI/UX Designer",
			"testing":     "QA Engineer",
			"release":     "Release Captain",
		}
		if agent, ok := phaseAgentMap[phase]; ok {
			return agent
		}
	}

	return "" // No agent identified
}
