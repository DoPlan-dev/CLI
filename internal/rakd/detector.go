package rakd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Detector detects required API keys from project dependencies and code
type Detector struct {
	projectRoot string
}

// NewDetector creates a new detector
func NewDetector(projectRoot string) *Detector {
	return &Detector{
		projectRoot: projectRoot,
	}
}

// DetectServices detects services from dependencies and code
func (d *Detector) DetectServices() ([]Service, error) {
	var services []Service

	// Detect from package.json (Node.js)
	if nodeServices := d.detectFromPackageJSON(); len(nodeServices) > 0 {
		services = append(services, nodeServices...)
	}

	// Detect from go.mod (Go)
	if goServices := d.detectFromGoMod(); len(goServices) > 0 {
		services = append(services, goServices...)
	}

	// Detect from requirements.txt (Python)
	if pythonServices := d.detectFromRequirements(); len(pythonServices) > 0 {
		services = append(services, pythonServices...)
	}

	// Detect from code imports
	if codeServices := d.detectFromCode(); len(codeServices) > 0 {
		services = append(services, codeServices...)
	}

	// Remove duplicates
	services = d.deduplicateServices(services)

	return services, nil
}

// detectFromPackageJSON detects services from package.json
func (d *Detector) detectFromPackageJSON() []Service {
	packageJSONPath := filepath.Join(d.projectRoot, "package.json")
	data, err := os.ReadFile(packageJSONPath)
	if err != nil {
		return nil
	}

	var pkg struct {
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil
	}

	var services []Service
	allDeps := make(map[string]string)
	for k, v := range pkg.Dependencies {
		allDeps[k] = v
	}
	for k, v := range pkg.DevDependencies {
		allDeps[k] = v
	}

	// Known service patterns
	servicePatterns := map[string]Service{
		"@stripe/stripe-js": {
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Publishable Key", EnvVar: "NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY", Required: true, Format: "pk_test_... or pk_live_..."},
				{Name: "Secret Key", EnvVar: "STRIPE_SECRET_KEY", Required: true, Format: "sk_test_... or sk_live_..."},
			},
		},
		"stripe": {
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Publishable Key", EnvVar: "STRIPE_PUBLISHABLE_KEY", Required: true, Format: "pk_test_... or pk_live_..."},
				{Name: "Secret Key", EnvVar: "STRIPE_SECRET_KEY", Required: true, Format: "sk_test_... or sk_live_..."},
			},
		},
		"@sendgrid/mail": {
			Name:        "SendGrid",
			Category:    "email",
			Description: "Email delivery service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "API Key", EnvVar: "SENDGRID_API_KEY", Required: true, Format: "SG...."},
			},
		},
		"nodemailer": {
			Name:        "Email Service",
			Category:    "email",
			Description: "Email sending service (SMTP)",
			Priority:    "optional",
			Detected:    true,
			Keys: []APIKey{
				{Name: "SMTP Host", EnvVar: "SMTP_HOST", Required: false},
				{Name: "SMTP Port", EnvVar: "SMTP_PORT", Required: false},
				{Name: "SMTP User", EnvVar: "SMTP_USER", Required: false},
				{Name: "SMTP Password", EnvVar: "SMTP_PASSWORD", Required: false},
			},
		},
		"@aws-sdk/client-s3": {
			Name:        "AWS S3",
			Category:    "storage",
			Description: "Object storage service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Access Key ID", EnvVar: "AWS_ACCESS_KEY_ID", Required: true},
				{Name: "Secret Access Key", EnvVar: "AWS_SECRET_ACCESS_KEY", Required: true},
				{Name: "Region", EnvVar: "AWS_REGION", Required: true},
			},
		},
		"@supabase/supabase-js": {
			Name:        "Supabase",
			Category:    "database",
			Description: "Backend-as-a-Service platform",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "URL", EnvVar: "NEXT_PUBLIC_SUPABASE_URL", Required: true},
				{Name: "Anon Key", EnvVar: "NEXT_PUBLIC_SUPABASE_ANON_KEY", Required: true},
				{Name: "Service Role Key", EnvVar: "SUPABASE_SERVICE_ROLE_KEY", Required: true},
			},
		},
		"@auth0/nextjs-auth0": {
			Name:        "Auth0",
			Category:    "authentication",
			Description: "Authentication service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Domain", EnvVar: "AUTH0_DOMAIN", Required: true},
				{Name: "Client ID", EnvVar: "AUTH0_CLIENT_ID", Required: true},
				{Name: "Client Secret", EnvVar: "AUTH0_CLIENT_SECRET", Required: true},
				{Name: "Base URL", EnvVar: "AUTH0_BASE_URL", Required: true},
			},
		},
		"@clerk/nextjs": {
			Name:        "Clerk",
			Category:    "authentication",
			Description: "Authentication service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Publishable Key", EnvVar: "NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY", Required: true},
				{Name: "Secret Key", EnvVar: "CLERK_SECRET_KEY", Required: true},
			},
		},
		"openai": {
			Name:        "OpenAI",
			Category:    "ai",
			Description: "AI/ML service",
			Priority:    "optional",
			Detected:    true,
			Keys: []APIKey{
				{Name: "API Key", EnvVar: "OPENAI_API_KEY", Required: false, Format: "sk-..."},
			},
		},
		"@google-cloud/storage": {
			Name:        "Google Cloud Storage",
			Category:    "storage",
			Description: "Object storage service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Project ID", EnvVar: "GOOGLE_CLOUD_PROJECT_ID", Required: true},
				{Name: "Credentials JSON", EnvVar: "GOOGLE_APPLICATION_CREDENTIALS", Required: true},
			},
		},
		"@sentry/nextjs": {
			Name:        "Sentry",
			Category:    "analytics",
			Description: "Error tracking and monitoring",
			Priority:    "optional",
			Detected:    true,
			Keys: []APIKey{
				{Name: "DSN", EnvVar: "NEXT_PUBLIC_SENTRY_DSN", Required: false},
			},
		},
	}

	for depName := range allDeps {
		if service, ok := servicePatterns[depName]; ok {
			services = append(services, service)
		}
	}

	return services
}

// detectFromGoMod detects services from go.mod
func (d *Detector) detectFromGoMod() []Service {
	goModPath := filepath.Join(d.projectRoot, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return nil
	}

	var services []Service
	content := string(data)

	// Detect common Go service patterns
	if strings.Contains(content, "github.com/stripe/stripe-go") {
		services = append(services, Service{
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "API Key", EnvVar: "STRIPE_API_KEY", Required: true, Format: "sk_test_... or sk_live_..."},
			},
		})
	}

	if strings.Contains(content, "github.com/aws/aws-sdk-go") {
		services = append(services, Service{
			Name:        "AWS",
			Category:    "storage",
			Description: "AWS services",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "Access Key ID", EnvVar: "AWS_ACCESS_KEY_ID", Required: true},
				{Name: "Secret Access Key", EnvVar: "AWS_SECRET_ACCESS_KEY", Required: true},
				{Name: "Region", EnvVar: "AWS_REGION", Required: true},
			},
		})
	}

	return services
}

// detectFromRequirements detects services from requirements.txt (Python)
func (d *Detector) detectFromRequirements() []Service {
	requirementsPath := filepath.Join(d.projectRoot, "requirements.txt")
	data, err := os.ReadFile(requirementsPath)
	if err != nil {
		return nil
	}

	var services []Service
	content := string(data)

	if strings.Contains(content, "stripe") {
		services = append(services, Service{
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing service",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "API Key", EnvVar: "STRIPE_API_KEY", Required: true, Format: "sk_test_... or sk_live_..."},
			},
		})
	}

	return services
}

// detectFromCode detects services from code imports/usage
func (d *Detector) detectFromCode() []Service {
	// This is a simplified version - could be enhanced with AST parsing
	var services []Service

	// Check for common patterns in code files
	patterns := map[string]Service{
		"STRIPE": {
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing (detected from code)",
			Priority:    "required",
			Detected:    true,
			Keys: []APIKey{
				{Name: "API Key", EnvVar: "STRIPE_API_KEY", Required: true},
			},
		},
	}

	// Simple regex search for now
	for pattern, service := range patterns {
		if d.searchInCode(pattern) {
			services = append(services, service)
		}
	}

	return services
}

// searchInCode searches for a pattern in code files
func (d *Detector) searchInCode(pattern string) bool {
	extensions := []string{".js", ".ts", ".jsx", ".tsx", ".go", ".py"}
	for _, ext := range extensions {
		patternRegex := regexp.MustCompile(fmt.Sprintf(`(?i)%s`, regexp.QuoteMeta(pattern)))
		err := filepath.Walk(d.projectRoot, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				// Skip common directories
				if info.Name() == "node_modules" || info.Name() == ".git" || info.Name() == ".doplan" {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.HasSuffix(path, ext) {
				data, err := os.ReadFile(path)
				if err == nil && patternRegex.Match(data) {
					return fmt.Errorf("found") // Signal found
				}
			}
			return nil
		})
		if err != nil && err.Error() == "found" {
			return true
		}
	}
	return false
}

// deduplicateServices removes duplicate services
func (d *Detector) deduplicateServices(services []Service) []Service {
	seen := make(map[string]bool)
	var unique []Service

	for _, service := range services {
		key := fmt.Sprintf("%s-%s", service.Name, service.Category)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, service)
		}
	}

	return unique
}
