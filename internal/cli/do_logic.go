package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DoPlan-dev/CLI/internal/utils"
)

// doLogic handles the complete /do command workflow
type doLogic struct {
	projectPath string
	systemDir   string
	coreDir     string
	historyDir  string
}

// newDoLogic creates a new doLogic instance
func newDoLogic(projectPath string) (*doLogic, error) {
	systemDir := filepath.Join(projectPath, ".do", "system")
	coreDir := filepath.Join(projectPath, ".do", "core")
	historyDir := filepath.Join(projectPath, ".do", "system", "history")

	// Ensure directories exist
	if err := utils.CreateDirectory(systemDir); err != nil {
		return nil, fmt.Errorf("failed to create system directory: %w", err)
	}
	if err := utils.CreateDirectory(coreDir); err != nil {
		return nil, fmt.Errorf("failed to create core directory: %w", err)
	}
	if err := utils.CreateDirectory(historyDir); err != nil {
		return nil, fmt.Errorf("failed to create history directory: %w", err)
	}

	return &doLogic{
		projectPath: projectPath,
		systemDir:   systemDir,
		coreDir:     coreDir,
		historyDir:  historyDir,
	}, nil
}

// runIdeationPhase captures and saves the project idea with iterative conversation
func (d *doLogic) runIdeationPhase(idea string, out io.Writer, memoryCard *MemoryCard) (string, error) {
	var allIdeas []string
	reader := bufio.NewReader(os.Stdin)

	// Load memory card for personalized greeting
	greeting := "💡 Let's capture your project idea!"
	if memoryCard != nil {
		greeting = memoryCard.GetGreeting() + "\n\n💡 " + memoryCard.GetEncouragement() + "\n\nLet's start by capturing your project idea!"
	}

	fmt.Fprintln(out, greeting)
	fmt.Fprintln(out, "")

	// Initial idea capture
	if idea == "" {
		fmt.Fprint(out, "What's your project idea? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read input: %w", err)
		}
		idea = strings.TrimSpace(input)
		if idea == "" {
			return "", fmt.Errorf("idea is required to proceed")
		}
	}

	allIdeas = append(allIdeas, idea)

	// Iterative conversation - ask for more details
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "✨ Great start! You can tell me more about your idea.")
	fmt.Fprintln(out, "   Often, users remember more details as they talk about it, and new improvements come to mind.")
	fmt.Fprintln(out, "")
	fmt.Fprint(out, "Tell me more about your idea (or type 'done' when you're finished): ")

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read input: %w", err)
		}

		additionalInput := strings.TrimSpace(input)
		lowerInput := strings.ToLower(additionalInput)

		if lowerInput == "done" || lowerInput == "that's all" || lowerInput == "that's it" || lowerInput == "finished" {
			break
		}

		if additionalInput != "" {
			allIdeas = append(allIdeas, additionalInput)

			// Thank user and encourage more
			fmt.Fprintln(out, "")
			fmt.Fprintln(out, "🙏 Thank you for those details! It's great you didn't forget about that - it will definitely improve the results.")
			fmt.Fprintln(out, "")
			fmt.Fprint(out, "You can tell me more, or if you have notes, additional features, or any other thoughts, share them now (or type 'done' to finish): ")
		} else {
			fmt.Fprint(out, "Anything else to add? (or type 'done' to finish): ")
		}
	}

	// Combine all ideas into final content
	finalIdea := strings.Join(allIdeas, "\n\n")

	// Save idea to IDEA.md
	ideaPath := filepath.Join(d.systemDir, "IDEA.md")
	content := fmt.Sprintf(`# Project Idea

## Overview
%s

## Additional Details
*Captured through iterative conversation*

## Goals
- [To be defined during meeting]

## Target Users
- [To be defined during meeting]

## Key Features
- [To be defined during meeting]

## Success Metrics
- [To be defined during meeting]

---
*Generated: %s*
*Command: /do ideation*
*Conversation rounds: %d*
`, finalIdea, time.Now().Format(time.RFC3339), len(allIdeas))

	if err := utils.WriteFile(ideaPath, []byte(content)); err != nil {
		return "", fmt.Errorf("failed to save IDEA.md: %w", err)
	}

	// Update memory card
	if memoryCard != nil {
		memoryCard.UpdateFromConversation(map[string]string{
			"ideation_rounds": fmt.Sprintf("%d", len(allIdeas)),
			"idea_length":     fmt.Sprintf("%d", len(finalIdea)),
		})
		SaveMemoryCard(memoryCard)
	}

	// Update active state
	if err := d.updateActiveState("idea", false); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Could not update active state: %v\n", err)
	}

	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "✅ Idea captured and saved to %s\n", ideaPath)
	fmt.Fprintf(out, "   Collected %d rounds of details - this will help create a better plan!\n", len(allIdeas))

	return finalIdea, nil
}

// runMeetingPhase conducts the discovery meeting
func (d *doLogic) runMeetingPhase(idea string, out io.Writer) error {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🤝 Starting discovery meeting...")
	fmt.Fprintln(out, "")

	// Determine meeting speed (simplified - can be enhanced)
	speed, err := d.selectMeetingSpeed(out)
	if err != nil {
		return err
	}

	// Load phase templates
	phases, err := d.loadPhaseTemplates(speed)
	if err != nil {
		return fmt.Errorf("failed to load phase templates: %w", err)
	}

	// Conduct meeting
	answers := make(map[string]map[string]string)
	for phaseNum, phase := range phases {
		fmt.Fprintf(out, "\n📋 Phase %s: %s\n", phaseNum, phase.Title)
		fmt.Fprintln(out, strings.Repeat("─", 60))

		phaseAnswers := make(map[string]string)
		for _, question := range phase.Questions {
			answer, err := d.askQuestion(question, out)
			if err != nil {
				return err
			}
			phaseAnswers[question] = answer
		}
		answers[phaseNum] = phaseAnswers
		fmt.Fprintln(out, "")
	}

	// Generate BRAINSTORM.md
	if err := d.generateBrainstorm(idea, speed, answers, out); err != nil {
		return fmt.Errorf("failed to generate BRAINSTORM.md: %w", err)
	}

	// Update active state
	if err := d.updateActiveState("brainstorm", false); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Could not update active state: %v\n", err)
	}

	fmt.Fprintln(out, "✅ Discovery meeting completed!")
	return nil
}

// runRefiningPhase provides enhancements and suggestions
func (d *doLogic) runRefiningPhase(out io.Writer) error {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "✨ Refining, enhancing, and generating suggestions...")
	fmt.Fprintln(out, "")

	// Read IDEA.md and BRAINSTORM.md
	ideaPath := filepath.Join(d.systemDir, "IDEA.md")
	brainstormPath := filepath.Join(d.systemDir, "BRAINSTORM.md")

	ideaContent, err := os.ReadFile(ideaPath)
	if err != nil {
		return fmt.Errorf("failed to read IDEA.md: %w", err)
	}

	brainstormContent, err := os.ReadFile(brainstormPath)
	if err != nil {
		return fmt.Errorf("failed to read BRAINSTORM.md: %w", err)
	}

	// Generate refinements
	refinements := d.generateRefinements(string(ideaContent), string(brainstormContent))

	// Save refinements
	refinementsPath := filepath.Join(d.systemDir, "REFINEMENTS.md")
	content := fmt.Sprintf(`# Project Refinements and Suggestions

*Generated: %s*
*Command: /do refining*

## Summary
Based on your idea and discovery meeting, here are our recommendations for enhancement:

%s

## Next Steps
1. Review the suggestions above
2. Update IDEA.md or BRAINSTORM.md if needed
3. Type /plan to generate your execution plan
4. Type /dev to start development

---
*Generated by DoPlan CLI*
`, time.Now().Format(time.RFC3339), refinements)

	if err := utils.WriteFile(refinementsPath, []byte(content)); err != nil {
		return fmt.Errorf("failed to save REFINEMENTS.md: %w", err)
	}

	fmt.Fprintf(out, "✅ Refinements and suggestions saved to %s\n", refinementsPath)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "💡 Key Suggestions:")
	fmt.Fprintln(out, refinements)

	return nil
}

// Helper functions

func (d *doLogic) selectMeetingSpeed(out io.Writer) (string, error) {
	fmt.Fprintln(out, "Select meeting speed:")
	fmt.Fprintln(out, "  1. Quick Start (5-10 min) - Essential questions only")
	fmt.Fprintln(out, "  2. Standard (15-20 min) - Balanced depth [Recommended]")
	fmt.Fprintln(out, "  3. Comprehensive (30-45 min) - Detailed planning")
	fmt.Fprint(out, "Your choice (1-3, default: 2): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		input = "2"
	}

	switch input {
	case "1":
		return "quick", nil
	case "2":
		return "standard", nil
	case "3":
		return "comprehensive", nil
	default:
		return "standard", nil
	}
}

type phaseTemplate struct {
	Title     string
	Questions []string
}

func (d *doLogic) loadPhaseTemplates(speed string) (map[string]phaseTemplate, error) {
	phases := make(map[string]phaseTemplate)

	// Define phases based on speed
	var phaseNumbers []string
	switch speed {
	case "quick":
		phaseNumbers = []string{"01", "03"} // Vision and Experience only
	case "standard":
		phaseNumbers = []string{"01", "02", "03", "06"} // Vision, Audience, Experience, Delivery
	case "comprehensive":
		phaseNumbers = []string{"01", "02", "03", "04", "05", "06"} // All phases
	default:
		phaseNumbers = []string{"01", "02", "03", "06"}
	}

	// Load phase templates from .do/core/brainstorm/
	brainstormDir := filepath.Join(d.coreDir, "brainstorm")
	for _, num := range phaseNumbers {
		// Try different filename patterns
		possibleNames := []string{
			fmt.Sprintf("phase-%s-vision.md", num),
			fmt.Sprintf("phase-%s-audience.md", num),
			fmt.Sprintf("phase-%s-experience.md", num),
			fmt.Sprintf("phase-%s-content.md", num),
			fmt.Sprintf("phase-%s-marketing.md", num),
			fmt.Sprintf("phase-%s-delivery.md", num),
			fmt.Sprintf("phase-%s.md", num),
		}

		var phasePath string
		var found bool
		for _, name := range possibleNames {
			candidate := filepath.Join(brainstormDir, name)
			if utils.PathExists(candidate) {
				phasePath = candidate
				found = true
				break
			}
		}

		if !found {
			// Use default template
			phases[num] = d.getDefaultPhaseTemplate(num)
			continue
		}

		// Read and parse phase template
		content, err := os.ReadFile(phasePath)
		if err != nil {
			phases[num] = d.getDefaultPhaseTemplate(num)
			continue
		}

		phase := d.parsePhaseTemplate(string(content), num)
		phases[num] = phase
	}

	return phases, nil
}

func (d *doLogic) getDefaultPhaseTemplate(num string) phaseTemplate {
	templates := map[string]phaseTemplate{
		"01": {
			Title: "Vision & Outcomes",
			Questions: []string{
				"What is the primary problem we're solving?",
				"What is the vision for this product?",
				"What are the key success metrics?",
			},
		},
		"02": {
			Title: "Audience & Differentiation",
			Questions: []string{
				"Who is the primary target audience?",
				"How do we differentiate from competitors?",
			},
		},
		"03": {
			Title: "Experience, UI/UX & Tech",
			Questions: []string{
				"What user experience do we want to create?",
				"What are the technical requirements?",
			},
		},
		"04": {
			Title: "Content & SEO",
			Questions: []string{
				"What content needs do we have?",
				"What SEO requirements are important?",
			},
		},
		"05": {
			Title: "Marketing & Growth",
			Questions: []string{
				"What marketing channels will we use?",
				"How will we grow the user base?",
			},
		},
		"06": {
			Title: "Delivery, Ops & Risks",
			Questions: []string{
				"What are the delivery timelines?",
				"What are the main risks and how will we mitigate them?",
			},
		},
	}

	if phase, ok := templates[num]; ok {
		return phase
	}
	return phaseTemplate{Title: "General", Questions: []string{"Tell us more about your project."}}
}

func (d *doLogic) parsePhaseTemplate(content, num string) phaseTemplate {
	lines := strings.Split(content, "\n")
	phase := phaseTemplate{}
	var questions []string
	inQuestions := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			// Extract title
			title := strings.TrimPrefix(line, "#")
			title = strings.TrimPrefix(title, "Phase")
			title = strings.TrimSpace(title)
			title = strings.TrimPrefix(title, num+":")
			title = strings.TrimSpace(title)
			if title != "" {
				phase.Title = title
			}
		}
		if strings.Contains(strings.ToLower(line), "question") {
			inQuestions = true
		}
		if inQuestions && strings.HasPrefix(line, "-") {
			// Extract question
			q := strings.TrimPrefix(line, "-")
			q = strings.TrimSpace(q)
			if q != "" && !strings.HasPrefix(q, "**") {
				questions = append(questions, q)
			}
		}
	}

	if len(questions) == 0 {
		// Fallback to default
		return d.getDefaultPhaseTemplate(num)
	}

	phase.Questions = questions
	if phase.Title == "" {
		phase.Title = d.getDefaultPhaseTemplate(num).Title
	}

	return phase
}

func (d *doLogic) askQuestion(question string, out io.Writer) (string, error) {
	fmt.Fprintf(out, "❓ %s\n", question)
	fmt.Fprint(out, "   Your answer: ")

	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(answer), nil
}

func (d *doLogic) generateBrainstorm(idea, speed string, answers map[string]map[string]string, out io.Writer) error {
	var sb strings.Builder

	sb.WriteString("# Brainstorm Session\n\n")
	sb.WriteString(fmt.Sprintf("**Date**: %s\n", time.Now().Format("January 2, 2006")))
	sb.WriteString(fmt.Sprintf("**Meeting Speed**: %s\n", strings.Title(speed)))
	sb.WriteString(fmt.Sprintf("**Project Idea**: %s\n\n", idea))
	sb.WriteString("---\n\n")

	// Organize answers by phase (in order)
	phaseOrder := []string{"01", "02", "03", "04", "05", "06"}
	for _, num := range phaseOrder {
		if phaseAnswers, ok := answers[num]; ok {
			phase := d.getDefaultPhaseTemplate(num)
			sb.WriteString(fmt.Sprintf("## Phase %s: %s\n\n", num, phase.Title))

			for question, answer := range phaseAnswers {
				sb.WriteString(fmt.Sprintf("### %s\n", question))
				sb.WriteString(fmt.Sprintf("%s\n\n", answer))
			}
		}
	}

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("*Generated: %s*\n", time.Now().Format(time.RFC3339)))
	sb.WriteString("*Command: /do meeting*\n")

	brainstormPath := filepath.Join(d.systemDir, "BRAINSTORM.md")
	return utils.WriteFile(brainstormPath, []byte(sb.String()))
}

func (d *doLogic) generateRefinements(ideaContent, brainstormContent string) string {
	var sb strings.Builder

	sb.WriteString("### 1. Idea Clarity\n")
	sb.WriteString("- Ensure your idea is specific and actionable\n")
	sb.WriteString("- Consider adding more detail about target users\n")
	sb.WriteString("- Define clear success metrics\n\n")

	sb.WriteString("### 2. Technical Considerations\n")
	sb.WriteString("- Review technical requirements from the meeting\n")
	sb.WriteString("- Consider scalability and performance early\n")
	sb.WriteString("- Plan for security and data privacy\n\n")

	sb.WriteString("### 3. User Experience\n")
	sb.WriteString("- Focus on core user journeys first\n")
	sb.WriteString("- Plan for accessibility and mobile responsiveness\n")
	sb.WriteString("- Consider onboarding and user education\n\n")

	sb.WriteString("### 4. Next Steps\n")
	sb.WriteString("- Review and refine IDEA.md with more details\n")
	sb.WriteString("- Update BRAINSTORM.md with any additional insights\n")
	sb.WriteString("- Use /plan to generate your execution plan\n")

	return sb.String()
}

func (d *doLogic) updateActiveState(phase string, locked bool) error {
	statePath := filepath.Join(d.historyDir, "active_state.json")

	state := map[string]interface{}{
		"phase":  phase,
		"locked": locked,
	}

	// Try to read existing state
	if data, err := os.ReadFile(statePath); err == nil {
		var existing map[string]interface{}
		if err := json.Unmarshal(data, &existing); err == nil {
			// Preserve existing fields
			if activeTask, ok := existing["active_task"]; ok {
				state["active_task"] = activeTask
			}
			if completed, ok := existing["completed"]; ok {
				state["completed"] = completed
			}
		}
	}

	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return utils.WriteFile(statePath, jsonData)
}

// runFeatureIdeation handles adding a single feature idea
func (d *doLogic) runFeatureIdeation(out io.Writer, memoryCard *MemoryCard) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Fprintln(out, "✨ Let's add a feature idea to your project!")
	fmt.Fprintln(out, "")
	fmt.Fprint(out, "What feature would you like to add? ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	featureIdea := strings.TrimSpace(input)
	if featureIdea == "" {
		return "", fmt.Errorf("feature idea cannot be empty")
	}

	// Read existing IDEA.md or create new
	ideaPath := filepath.Join(d.systemDir, "IDEA.md")
	var existingContent string

	if utils.PathExists(ideaPath) {
		data, err := os.ReadFile(ideaPath)
		if err == nil {
			existingContent = string(data)
		}
	}

	// Append feature to existing content or create new
	var newContent string
	if existingContent != "" {
		// Add feature section if it doesn't exist
		if !strings.Contains(existingContent, "## Features") {
			newContent = existingContent + "\n\n## Features\n- " + featureIdea + "\n"
		} else {
			// Append to existing features
			newContent = strings.Replace(existingContent, "## Features", "## Features\n- "+featureIdea, 1)
		}
	} else {
		newContent = fmt.Sprintf(`# Project Idea

## Overview
[Main project idea]

## Features
- %s

---
*Generated: %s*
*Command: /do feature*
`, featureIdea, time.Now().Format(time.RFC3339))
	}

	if err := utils.WriteFile(ideaPath, []byte(newContent)); err != nil {
		return "", fmt.Errorf("failed to save feature: %w", err)
	}

	return featureIdea, nil
}

// runFastTrack handles fast track mode with detailed prompt/PRD
func (d *doLogic) runFastTrack(out io.Writer, prompt, prdFile string) error {
	reader := bufio.NewReader(os.Stdin)
	var finalContent string

	fmt.Fprintln(out, "⚡ Fast Track Mode - Skipping discovery meeting")
	fmt.Fprintln(out, "")

	// Collect prompt
	if prompt == "" {
		fmt.Fprint(out, "Enter your detailed project prompt (or press Enter to skip): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read prompt: %w", err)
		}
		prompt = strings.TrimSpace(input)
	}

	// Collect PRD file
	if prdFile == "" {
		fmt.Fprint(out, "Path to PRD file (or press Enter to skip): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read PRD path: %w", err)
		}
		prdFile = strings.TrimSpace(input)
	}

	// Load PRD content if provided
	var prdContent string
	if prdFile != "" && utils.PathExists(prdFile) {
		data, err := os.ReadFile(prdFile)
		if err == nil {
			prdContent = string(data)
		}
	}

	// Combine prompt and PRD
	if prompt != "" {
		finalContent += "# Project Prompt\n\n" + prompt + "\n\n"
	}
	if prdContent != "" {
		finalContent += "# PRD Document\n\n" + prdContent + "\n\n"
	}

	if finalContent == "" {
		return fmt.Errorf("either prompt or PRD file must be provided")
	}

	// Save to IDEA.md
	ideaPath := filepath.Join(d.systemDir, "IDEA.md")
	content := finalContent + fmt.Sprintf(`---
*Generated: %s*
*Command: /do now (Fast Track)*
*Skipped: Discovery Meeting*
`, time.Now().Format(time.RFC3339))

	if err := utils.WriteFile(ideaPath, []byte(content)); err != nil {
		return fmt.Errorf("failed to save IDEA.md: %w", err)
	}

	// Update state to ready for planning
	if err := d.updateActiveState("ready_for_planning", false); err != nil {
		fmt.Fprintf(out, "⚠️  Warning: Could not update active state: %v\n", err)
	}

	fmt.Fprintf(out, "✅ Fast track complete! Saved to %s\n", ideaPath)
	return nil
}

// runLuckyMode handles the iterative idea suggestion flow
func (d *doLogic) runLuckyMode(out io.Writer, memoryCard *MemoryCard) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var selectedIdeas []string
	var rejectedIdeas []string
	round := 1

	fmt.Fprintln(out, "🍀 Lucky Mode - Let's discover something amazing!")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "I'll suggest 2 ideas. Choose one you like, or ask for 2 more.")
	fmt.Fprintln(out, "We'll keep going until you find something you love!")
	fmt.Fprintln(out, "")

	for {
		// Generate 2 unique ideas (avoiding rejected ones)
		ideas := d.generateIdeas(2, rejectedIdeas, memoryCard)

		fmt.Fprintf(out, "--- Round %d ---\n", round)
		fmt.Fprintln(out, "")
		fmt.Fprintf(out, "💡 Idea 1: %s\n", ideas[0])
		fmt.Fprintf(out, "💡 Idea 2: %s\n", ideas[1])
		fmt.Fprintln(out, "")
		fmt.Fprint(out, "Choose (1, 2, or 'more' for 2 new ideas, 'done' to finish): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read input: %w", err)
		}

		choice := strings.TrimSpace(strings.ToLower(input))

		switch choice {
		case "1":
			selectedIdeas = append(selectedIdeas, ideas[0])
			rejectedIdeas = append(rejectedIdeas, ideas[1])
			fmt.Fprintf(out, "\n✅ Selected: %s\n", ideas[0])
			return ideas[0], d.saveLuckyIdea(ideas[0], selectedIdeas)
		case "2":
			selectedIdeas = append(selectedIdeas, ideas[1])
			rejectedIdeas = append(rejectedIdeas, ideas[0])
			fmt.Fprintf(out, "\n✅ Selected: %s\n", ideas[1])
			return ideas[1], d.saveLuckyIdea(ideas[1], selectedIdeas)
		case "more", "m":
			rejectedIdeas = append(rejectedIdeas, ideas...)
			round++
			fmt.Fprintln(out, "\n🔄 Generating 2 new ideas based on your preferences...")
			fmt.Fprintln(out, "")
			continue
		case "done", "d":
			if len(selectedIdeas) > 0 {
				// Use the last selected idea
				lastIdea := selectedIdeas[len(selectedIdeas)-1]
				return lastIdea, d.saveLuckyIdea(lastIdea, selectedIdeas)
			}
			return "", fmt.Errorf("no idea selected")
		default:
			fmt.Fprintln(out, "Invalid choice. Please enter 1, 2, 'more', or 'done'")
			continue
		}
	}
}

// generateIdeas generates unique project ideas
func (d *doLogic) generateIdeas(count int, rejected []string, memoryCard *MemoryCard) []string {
	// Idea pool - in a real implementation, this could use AI or a larger database
	ideaPool := []string{
		"Build a personal finance tracker with AI-powered insights",
		"Create a social learning platform for coding bootcamps",
		"Develop a habit tracking app with gamification",
		"Build a local business discovery and review platform",
		"Create a collaborative note-taking tool for teams",
		"Develop a recipe sharing app with dietary restrictions",
		"Build a fitness challenge platform with community support",
		"Create a language learning app with conversation practice",
		"Develop a project management tool for freelancers",
		"Build a mental health journaling app with mood tracking",
		"Create a sustainable living tips and challenges app",
		"Develop a local event discovery and RSVP platform",
		"Build a skill-sharing marketplace for professionals",
		"Create a book recommendation engine based on reading history",
		"Develop a travel planning tool with budget tracking",
	}

	// Filter out rejected ideas
	available := make([]string, 0)
	for _, idea := range ideaPool {
		isRejected := false
		for _, r := range rejected {
			if strings.Contains(strings.ToLower(idea), strings.ToLower(r)) {
				isRejected = true
				break
			}
		}
		if !isRejected {
			available = append(available, idea)
		}
	}

	// If we don't have enough, add some generic ones
	if len(available) < count {
		available = append(available, ideaPool...)
	}

	// Return requested count
	if len(available) < count {
		count = len(available)
	}

	return available[:count]
}

// saveLuckyIdea saves the selected idea from lucky mode
func (d *doLogic) saveLuckyIdea(idea string, allSelected []string) error {
	ideaPath := filepath.Join(d.systemDir, "IDEA.md")
	content := fmt.Sprintf(`# Project Idea (Lucky Mode)

## Overview
%s

## Selection Journey
*You explored %d ideas before choosing this one*

## Goals
- [To be defined during planning]

## Target Users
- [To be defined during planning]

## Key Features
- [To be defined during planning]

## Success Metrics
- [To be defined during planning]

---
*Generated: %s*
*Command: /do i'm lucky*
*Skipped: Discovery Meeting (Fast Track)*
`, idea, len(allSelected), time.Now().Format(time.RFC3339))

	if err := utils.WriteFile(ideaPath, []byte(content)); err != nil {
		return fmt.Errorf("failed to save IDEA.md: %w", err)
	}

	// Update state
	return d.updateActiveState("ready_for_planning", false)
}
