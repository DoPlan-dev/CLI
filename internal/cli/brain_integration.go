package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
)

// EnhanceAgentFile enhances an agent markdown file with memory card context
func EnhanceAgentFile(agentFilePath string, agentRole string) error {
	// Load brain
	brain, err := NewBrain()
	if err != nil {
		// If brain can't load, return without error (graceful degradation)
		return nil
	}

	// Read agent file
	data, err := os.ReadFile(agentFilePath)
	if err != nil {
		return fmt.Errorf("failed to read agent file: %w", err)
	}

	agentContent := string(data)

	// Extract system prompt (between ## System Prompt and next ##)
	parts := strings.Split(agentContent, "## System Prompt")
	if len(parts) < 2 {
		// No system prompt section found, skip enhancement
		return nil
	}

	beforePrompt := parts[0] + "## System Prompt"
	afterPrompt := strings.SplitN(parts[1], "##", 2)

	if len(afterPrompt) < 1 {
		return nil
	}

	// Extract base prompt (everything before next ##)
	basePrompt := strings.TrimSpace(afterPrompt[0])
	restOfFile := ""
	if len(afterPrompt) > 1 {
		restOfFile = "##" + afterPrompt[1]
	}

	// Enhance prompt with brain
	enhancedPrompt := brain.EnhanceAgentPrompt(basePrompt, agentRole)

	// Reconstruct file
	enhancedContent := beforePrompt + "\n" + enhancedPrompt + "\n\n" + restOfFile

	// Write back
	return utils.WriteFile(agentFilePath, []byte(enhancedContent))
}

// EnhanceAllAgentFiles enhances all agent files in a project
func EnhanceAllAgentFiles(projectPath string) error {
	agentDir := filepath.Join(projectPath, ".cursor", "agents")
	if !utils.PathExists(agentDir) {
		// Also check .do/core/agents
		agentDir = filepath.Join(projectPath, ".do", "core", "agents")
		if !utils.PathExists(agentDir) {
			return nil // No agents directory, skip
		}
	}

	// Read all agent files
	files, err := os.ReadDir(agentDir)
	if err != nil {
		return fmt.Errorf("failed to read agents directory: %w", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}

		agentPath := filepath.Join(agentDir, file.Name())

		// Determine agent role from filename
		agentRole := strings.TrimSuffix(file.Name(), ".md")
		agentRole = strings.ReplaceAll(agentRole, "_", " ")
		agentRole = strings.Title(agentRole)

		// Enhance agent file
		if err := EnhanceAgentFile(agentPath, agentRole); err != nil {
			// Log error but continue with other files
			fmt.Printf("Warning: Failed to enhance agent file %s: %v\n", file.Name(), err)
		}
	}

	return nil
}

// ProcessCommandWithBrain processes a command with brain integration
func ProcessCommandWithBrain(
	command string,
	userInput string,
	baseAgentPrompt string,
	agentRole string,
	out io.Writer,
) (string, error) {
	// Initialize brain
	agentBrain, err := NewAgentBrain()
	if err != nil {
		// Graceful degradation - continue without brain
		return baseAgentPrompt, nil
	}

	startTime := time.Now()

	// Enhance agent prompt
	_ = agentBrain.ProcessAgentPrompt(baseAgentPrompt, agentRole)

	// In a real implementation, you would call the agent here with enhancedPrompt
	// For now, we'll simulate by returning the enhanced prompt
	// enhancedPrompt := agentBrain.ProcessAgentPrompt(baseAgentPrompt, agentRole)
	// agentResponse := callAgent(enhancedPrompt, userInput)

	// For demonstration, create a sample response
	agentResponse := fmt.Sprintf("Based on your input: %s\n\nI've processed your request using the enhanced context from your memory card.", userInput)

	// Process and personalize response
	personalizedResponse := agentBrain.ProcessAgentResponse(agentResponse, command)

	// Calculate duration
	duration := time.Since(startTime).Seconds()

	// Determine sentiment (in real implementation, this would be analyzed)
	sentiment := "positive" // Default

	// Track interaction
	agentBrain.TrackInteraction(command, userInput, personalizedResponse, duration, sentiment)

	// Format response for user
	finalResponse := agentBrain.FormatResponseForUser(personalizedResponse, "response")

	return finalResponse, nil
}

// GetBrainEnhancedGreeting returns a greeting enhanced by brain
func GetBrainEnhancedGreeting(context string) string {
	brain, err := NewBrain()
	if err != nil {
		return "Hello! 👋 I'm DoPlan, your AI development partner."
	}

	return brain.GetContextualGreeting(context)
}

// GetBrainPersonalizedSuggestion returns a personalized suggestion
func GetBrainPersonalizedSuggestion(topic string) string {
	brain, err := NewBrain()
	if err != nil {
		return ""
	}

	return brain.GetPersonalizedSuggestion(topic)
}

// ShouldUseBrain returns whether brain system should be used
func ShouldUseBrain() bool {
	_, err := LoadMemoryCard()
	return err == nil
}
