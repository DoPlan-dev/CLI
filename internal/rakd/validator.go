package rakd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Validator validates API keys from .env file
type Validator struct {
	projectRoot string
	envVars     map[string]string
}

// NewValidator creates a new validator
func NewValidator(projectRoot string) *Validator {
	return &Validator{
		projectRoot: projectRoot,
		envVars:     make(map[string]string),
	}
}

// LoadEnvFile loads environment variables from .env file
func (v *Validator) LoadEnvFile() error {
	envPath := filepath.Join(v.projectRoot, ".env")
	data, err := os.ReadFile(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // .env file doesn't exist, that's okay
		}
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// Remove quotes if present
			value = strings.Trim(value, `"'`)
			v.envVars[key] = value
		}
	}

	return scanner.Err()
}

// ValidateService validates all keys for a service
func (v *Validator) ValidateService(service *Service) {
	for i := range service.Keys {
		key := &service.Keys[i]
		value, exists := v.envVars[key.EnvVar]

		if !exists || value == "" {
			if key.Required {
				key.Status = StatusRequired
			} else {
				key.Status = StatusOptional
			}
			continue
		}

		key.Value = value
		key.Validated = true

		// Validate format if format hint is provided
		if key.Format != "" {
			if v.validateFormat(value, key.Format) {
				key.Status = StatusConfigured
			} else {
				key.Status = StatusInvalid
				key.Error = fmt.Sprintf("Key format doesn't match expected pattern: %s", key.Format)
			}
		} else {
			// Basic validation: key should not be empty and have reasonable length
			if len(value) >= 8 {
				key.Status = StatusConfigured
			} else {
				key.Status = StatusInvalid
				key.Error = "Key appears to be too short"
			}
		}
	}

	// Update service status based on keys
	v.updateServiceStatus(service)
}

// validateFormat validates key format against expected pattern
func (v *Validator) validateFormat(value, format string) bool {
	// Convert format hint to regex pattern
	pattern := format
	pattern = strings.ReplaceAll(pattern, "...", ".*")
	pattern = strings.ReplaceAll(pattern, ".", "\\.")
	pattern = "^" + pattern + "$"

	matched, err := regexp.MatchString(pattern, value)
	return err == nil && matched
}

// updateServiceStatus updates the service status based on key statuses
func (v *Validator) updateServiceStatus(service *Service) {
	hasRequired := false
	hasMissingRequired := false
	allConfigured := true

	for _, key := range service.Keys {
		if key.Required {
			hasRequired = true
			if key.Status != StatusConfigured {
				hasMissingRequired = true
				allConfigured = false
			}
		}
		if key.Status != StatusConfigured {
			allConfigured = false
		}
	}

	if allConfigured {
		service.Status = StatusConfigured
	} else if hasMissingRequired {
		service.Status = StatusRequired
	} else if hasRequired {
		service.Status = StatusPending
	} else {
		service.Status = StatusOptional
	}
}

// ValidateAll validates all services
func (v *Validator) ValidateAll(services []Service) error {
	if err := v.LoadEnvFile(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	for i := range services {
		v.ValidateService(&services[i])
	}

	return nil
}
