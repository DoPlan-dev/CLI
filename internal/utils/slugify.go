package utils

import (
	"strings"
	"unicode"
)

// Slugify converts a string to a URL-friendly slug (kebab-case)
// Examples:
//   - "My Feature Name" -> "my-feature-name"
//   - "User Authentication" -> "user-authentication"
//   - "API_Integration" -> "api-integration"
func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace spaces and underscores with hyphens
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	// Remove special characters, keep only alphanumeric and hyphens
	var result strings.Builder
	prevWasHyphen := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
			prevWasHyphen = false
		} else if r == '-' && !prevWasHyphen {
			result.WriteRune(r)
			prevWasHyphen = true
		}
	}

	slug := result.String()

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	// Collapse multiple consecutive hyphens into one
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	return slug
}
