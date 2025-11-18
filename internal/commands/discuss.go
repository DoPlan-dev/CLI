package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// DiscussIdea starts the idea discussion workflow
func DiscussIdea() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	fmt.Println("💬 Starting idea discussion...")
	fmt.Println()

	// Check if idea notes already exist
	ideaNotesPath := filepath.Join(projectRoot, "doplan", "idea-notes.md")
	if _, err := os.Stat(ideaNotesPath); err == nil {
		fmt.Println("📝 Found existing idea notes")
		fmt.Println("   This will enhance and refine your existing idea")
		fmt.Println()
	}

	// Create or update idea notes
	ideaNotes := `# Idea Notes

## Project Idea
Describe your project idea here.

## Goals
- Goal 1
- Goal 2
- Goal 3

## Target Audience
Who is this project for?

## Key Features
- Feature 1
- Feature 2
- Feature 3

## Technical Considerations
- Consideration 1
- Consideration 2

## Next Steps
1. Refine the idea
2. Generate PRD
3. Create plan
`

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(ideaNotesPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write idea notes (append if exists, create if not)
	if _, err := os.Stat(ideaNotesPath); os.IsNotExist(err) {
		if err := os.WriteFile(ideaNotesPath, []byte(ideaNotes), 0644); err != nil {
			return fmt.Errorf("failed to write idea notes: %w", err)
		}
	}

	fmt.Println("✅ Idea discussion completed!")
	fmt.Println("📄 Idea notes saved to: doplan/idea-notes.md")
	fmt.Println()
	fmt.Println("💡 Next steps:")
	fmt.Println("   • Review and refine: doplan/idea-notes.md")
	fmt.Println("   • Generate PRD: Run 'doplan generate'")
	fmt.Println("   • Create plan: Run 'doplan plan'")

	return nil
}

