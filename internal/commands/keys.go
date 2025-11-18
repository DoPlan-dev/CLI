package commands

import (
	"fmt"
	"os"

	"github.com/DoPlan-dev/CLI/internal/rakd"
	"github.com/DoPlan-dev/CLI/internal/sops"
	"github.com/DoPlan-dev/CLI/internal/wizard"
)

// ManageAPIKeys launches the API keys management TUI
func ManageAPIKeys() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Generate/update RAKD.md
	_, err = rakd.GenerateRAKD(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to generate RAKD: %w", err)
	}

	// Generate SOPS if they don't exist
	sopsGen := sops.NewGenerator(projectRoot)
	if err := sopsGen.GenerateAll(); err != nil {
		return fmt.Errorf("failed to generate SOPS: %w", err)
	}

	// Launch keys management wizard
	return wizard.RunKeysWizard()
}

// ValidateAPIKeys validates all API keys
func ValidateAPIKeys() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	data, err := rakd.GenerateRAKD(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to validate keys: %w", err)
	}

	fmt.Printf("✅ Validated %d services\n", len(data.Services))
	fmt.Printf("  Configured: %d\n", data.ConfiguredCount)
	fmt.Printf("  Required (Missing): %d\n", data.RequiredCount)
	fmt.Printf("  Pending: %d\n", data.PendingCount)
	fmt.Printf("  Optional: %d\n", data.OptionalCount)

	return nil
}

// SyncEnvExample syncs .env.example with detected keys
func SyncEnvExample() error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	detector := rakd.NewDetector(projectRoot)
	services, err := detector.DetectServices()
	if err != nil {
		return fmt.Errorf("failed to detect services: %w", err)
	}

	// Generate .env.example content
	envExamplePath := fmt.Sprintf("%s/.env.example", projectRoot)
	content := "# Environment Variables\n"
	content += "# Copy this file to .env and fill in your values\n\n"

	for _, service := range services {
		content += fmt.Sprintf("# %s\n", service.Name)
		for _, key := range service.Keys {
			if key.Required {
				content += fmt.Sprintf("%s=\n", key.EnvVar)
			} else {
				content += fmt.Sprintf("# %s=\n", key.EnvVar)
			}
		}
		content += "\n"
	}

	if err := os.WriteFile(envExamplePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write .env.example: %w", err)
	}

	fmt.Printf("✅ Synced .env.example with %d services\n", len(services))
	return nil
}
