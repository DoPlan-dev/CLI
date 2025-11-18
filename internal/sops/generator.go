package sops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ServiceTemplate represents a service setup template
type ServiceTemplate struct {
	Name         string
	Category     string
	Description  string
	Keys         []KeyTemplate
	SetupSteps   []string
	CodeExample  string
	CommonIssues []string
	Resources    []string
}

// KeyTemplate represents an API key template
type KeyTemplate struct {
	Name        string
	EnvVar      string
	Description string
	Required    bool
	Format      string
}

// Generator generates SOPS (Service Operating Procedures) documents
type Generator struct {
	projectRoot string
}

// NewGenerator creates a new SOPS generator
func NewGenerator(projectRoot string) *Generator {
	return &Generator{
		projectRoot: projectRoot,
	}
}

// GenerateAll generates all SOPS documents
func (g *Generator) GenerateAll() error {
	templates := g.getServiceTemplates()

	for _, template := range templates {
		if err := g.generateServiceSOP(template); err != nil {
			return fmt.Errorf("failed to generate SOP for %s: %w", template.Name, err)
		}
	}

	return nil
}

// GenerateForService generates SOPS for a specific service
func (g *Generator) GenerateForService(serviceName string) error {
	templates := g.getServiceTemplates()

	for _, template := range templates {
		if strings.EqualFold(template.Name, serviceName) {
			return g.generateServiceSOP(template)
		}
	}

	return fmt.Errorf("service template not found: %s", serviceName)
}

func (g *Generator) generateServiceSOP(template ServiceTemplate) error {
	sopDir := filepath.Join(g.projectRoot, "doplan", "SOPS", template.Category)
	if err := os.MkdirAll(sopDir, 0755); err != nil {
		return fmt.Errorf("failed to create SOP directory: %w", err)
	}

	filename := strings.ToLower(strings.ReplaceAll(template.Name, " ", "-")) + ".md"
	sopPath := filepath.Join(sopDir, filename)

	content := g.generateSOPContent(template)

	return os.WriteFile(sopPath, []byte(content), 0644)
}

func (g *Generator) generateSOPContent(template ServiceTemplate) string {
	var content strings.Builder

	// Header
	content.WriteString(fmt.Sprintf("# %s Setup Guide\n\n", template.Name))
	content.WriteString(fmt.Sprintf("**Category:** %s  \n", strings.Title(template.Category)))
	content.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", time.Now().Format("January 2, 2006")))
	content.WriteString("---\n\n")

	// Service Overview
	content.WriteString("## Service Overview\n\n")
	content.WriteString(fmt.Sprintf("%s\n\n", template.Description))
	content.WriteString("**When to Use:**\n")
	content.WriteString("- Use this service when you need to integrate ")
	content.WriteString(strings.ToLower(template.Description))
	content.WriteString("\n\n")

	// Setup Steps
	content.WriteString("## Setup Steps\n\n")
	for i, step := range template.SetupSteps {
		content.WriteString(fmt.Sprintf("%d. %s\n", i+1, step))
	}
	content.WriteString("\n")

	// API Key Creation
	content.WriteString("## API Key Creation\n\n")
	content.WriteString("### Required Keys\n\n")
	content.WriteString("| Key Name | Environment Variable | Required | Format |\n")
	content.WriteString("|----------|---------------------|----------|--------|\n")
	for _, key := range template.Keys {
		required := "No"
		if key.Required {
			required = "Yes"
		}
		format := key.Format
		if format == "" {
			format = "-"
		}
		content.WriteString(fmt.Sprintf("| %s | `%s` | %s | %s |\n",
			key.Name, key.EnvVar, required, format))
	}
	content.WriteString("\n")

	// Environment Variables
	content.WriteString("## Environment Variables\n\n")
	content.WriteString("Add the following to your `.env` file:\n\n")
	content.WriteString("```env\n")
	for _, key := range template.Keys {
		if key.Required {
			content.WriteString(fmt.Sprintf("%s=your_%s_here\n", key.EnvVar, strings.ToLower(key.Name)))
		} else {
			content.WriteString(fmt.Sprintf("# %s=optional_%s_here\n", key.EnvVar, strings.ToLower(key.Name)))
		}
	}
	content.WriteString("```\n\n")

	// Code Examples
	if template.CodeExample != "" {
		content.WriteString("## Code Examples\n\n")
		content.WriteString(template.CodeExample)
		content.WriteString("\n\n")
	}

	// Testing
	content.WriteString("## Testing\n\n")
	content.WriteString("### Verify Configuration\n\n")
	content.WriteString("1. Ensure all required environment variables are set\n")
	content.WriteString("2. Run `doplan keys validate` to verify keys\n")
	content.WriteString("3. Test API connection:\n\n")
	content.WriteString("```bash\n")
	content.WriteString(fmt.Sprintf("doplan keys test --service %s\n", template.Name))
	content.WriteString("```\n\n")

	// Common Issues
	if len(template.CommonIssues) > 0 {
		content.WriteString("## Common Issues\n\n")
		for i, issue := range template.CommonIssues {
			content.WriteString(fmt.Sprintf("%d. **%s**\n", i+1, issue))
		}
		content.WriteString("\n")
	}

	// Resources
	if len(template.Resources) > 0 {
		content.WriteString("## Resources\n\n")
		for _, resource := range template.Resources {
			content.WriteString(fmt.Sprintf("- %s\n", resource))
		}
		content.WriteString("\n")
	}

	return content.String()
}

func (g *Generator) getServiceTemplates() []ServiceTemplate {
	return []ServiceTemplate{
		{
			Name:        "Stripe",
			Category:    "payment",
			Description: "Payment processing and subscription management",
			Keys: []KeyTemplate{
				{Name: "Publishable Key", EnvVar: "STRIPE_PUBLISHABLE_KEY", Required: true, Format: "pk_test_... or pk_live_..."},
				{Name: "Secret Key", EnvVar: "STRIPE_SECRET_KEY", Required: true, Format: "sk_test_... or sk_live_..."},
			},
			SetupSteps: []string{
				"Sign up for a Stripe account at https://stripe.com",
				"Navigate to Developers > API keys",
				"Copy your Publishable key (starts with pk_)",
				"Copy your Secret key (starts with sk_)",
				"Add keys to your .env file",
			},
			CodeExample: "```javascript\nimport Stripe from 'stripe';\n\nconst stripe = new Stripe(process.env.STRIPE_SECRET_KEY);\n\n// Create a payment intent\nconst paymentIntent = await stripe.paymentIntents.create({\n  amount: 2000,\n  currency: 'usd',\n});\n```",
			CommonIssues: []string{
				"Using test keys in production - ensure you use live keys (pk_live_/sk_live_)",
				"Exposing secret keys in client-side code - never expose secret keys",
				"Invalid API key format - ensure keys start with pk_ or sk_",
			},
			Resources: []string{
				"[Stripe Documentation](https://stripe.com/docs)",
				"[Stripe API Reference](https://stripe.com/docs/api)",
				"[Stripe Testing](https://stripe.com/docs/testing)",
			},
		},
		{
			Name:        "SendGrid",
			Category:    "email",
			Description: "Email delivery service",
			Keys: []KeyTemplate{
				{Name: "API Key", EnvVar: "SENDGRID_API_KEY", Required: true, Format: "SG...."},
			},
			SetupSteps: []string{
				"Sign up for SendGrid at https://sendgrid.com",
				"Navigate to Settings > API Keys",
				"Create a new API key with 'Full Access' or 'Mail Send' permissions",
				"Copy the API key (starts with SG.)",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nconst sgMail = require('@sendgrid/mail');\nsgMail.setApiKey(process.env.SENDGRID_API_KEY);\n\nconst msg = {\n  to: 'recipient@example.com',\n  from: 'sender@example.com',\n  subject: 'Hello World',\n  text: 'Hello plain world!',\n  html: '<p>Hello HTML world!</p>',\n};\n\nawait sgMail.send(msg);\n```",
			CommonIssues: []string{
				"API key not working - ensure key has 'Mail Send' permissions",
				"Sender verification - verify your sender email in SendGrid dashboard",
				"Rate limits - check your SendGrid plan limits",
			},
			Resources: []string{
				"[SendGrid Documentation](https://docs.sendgrid.com)",
				"[SendGrid API Reference](https://docs.sendgrid.com/api-reference)",
			},
		},
		{
			Name:        "AWS S3",
			Category:    "storage",
			Description: "Object storage service",
			Keys: []KeyTemplate{
				{Name: "Access Key ID", EnvVar: "AWS_ACCESS_KEY_ID", Required: true},
				{Name: "Secret Access Key", EnvVar: "AWS_SECRET_ACCESS_KEY", Required: true},
				{Name: "Region", EnvVar: "AWS_REGION", Required: true},
			},
			SetupSteps: []string{
				"Sign up for AWS account at https://aws.amazon.com",
				"Navigate to IAM > Users",
				"Create a new user with S3 access permissions",
				"Create access keys for the user",
				"Copy Access Key ID and Secret Access Key",
				"Add to your .env file along with your preferred region",
			},
			CodeExample: "```javascript\nconst AWS = require('aws-sdk');\n\nconst s3 = new AWS.S3({\n  accessKeyId: process.env.AWS_ACCESS_KEY_ID,\n  secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY,\n  region: process.env.AWS_REGION,\n});\n\n// Upload a file\nconst params = {\n  Bucket: 'your-bucket-name',\n  Key: 'file.txt',\n  Body: 'File content',\n};\n\nawait s3.upload(params).promise();\n```",
			CommonIssues: []string{
				"Access denied - check IAM user permissions",
				"Region mismatch - ensure region matches bucket region",
				"Bucket doesn't exist - create bucket in AWS console first",
			},
			Resources: []string{
				"[AWS S3 Documentation](https://docs.aws.amazon.com/s3/)",
				"[AWS SDK Documentation](https://docs.aws.amazon.com/sdk-for-javascript/)",
			},
		},
		{
			Name:        "Supabase",
			Category:    "database",
			Description: "Backend-as-a-Service platform with PostgreSQL database",
			Keys: []KeyTemplate{
				{Name: "URL", EnvVar: "NEXT_PUBLIC_SUPABASE_URL", Required: true},
				{Name: "Anon Key", EnvVar: "NEXT_PUBLIC_SUPABASE_ANON_KEY", Required: true},
				{Name: "Service Role Key", EnvVar: "SUPABASE_SERVICE_ROLE_KEY", Required: true},
			},
			SetupSteps: []string{
				"Sign up for Supabase at https://supabase.com",
				"Create a new project",
				"Navigate to Settings > API",
				"Copy your Project URL",
				"Copy your anon/public key",
				"Copy your service_role key (keep this secret!)",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nimport { createClient } from '@supabase/supabase-js';\n\nconst supabase = createClient(\n  process.env.NEXT_PUBLIC_SUPABASE_URL,\n  process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY\n);\n\n// Query data\nconst { data, error } = await supabase\n  .from('users')\n  .select('*');\n```",
			CommonIssues: []string{
				"Row Level Security (RLS) - ensure RLS policies are configured",
				"Service role key exposure - never expose service_role key in client code",
				"Connection issues - check project URL and keys",
			},
			Resources: []string{
				"[Supabase Documentation](https://supabase.com/docs)",
				"[Supabase JavaScript Client](https://supabase.com/docs/reference/javascript/introduction)",
			},
		},
		{
			Name:        "Auth0",
			Category:    "authentication",
			Description: "Authentication and authorization service",
			Keys: []KeyTemplate{
				{Name: "Domain", EnvVar: "AUTH0_DOMAIN", Required: true},
				{Name: "Client ID", EnvVar: "AUTH0_CLIENT_ID", Required: true},
				{Name: "Client Secret", EnvVar: "AUTH0_CLIENT_SECRET", Required: true},
				{Name: "Base URL", EnvVar: "AUTH0_BASE_URL", Required: true},
			},
			SetupSteps: []string{
				"Sign up for Auth0 at https://auth0.com",
				"Create a new Application",
				"Copy your Domain (e.g., your-tenant.auth0.com)",
				"Copy your Client ID",
				"Copy your Client Secret",
				"Set your Application's callback URLs",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nimport { handleAuth, handleLogin } from '@auth0/nextjs-auth0';\n\nexport default handleAuth({\n  login: handleLogin({\n    authorizationParams: {\n      audience: 'your-api-identifier',\n    },\n  }),\n});\n```",
			CommonIssues: []string{
				"Callback URL mismatch - ensure callback URLs match in Auth0 dashboard",
				"Invalid audience - check API identifier configuration",
				"Session issues - check cookie settings",
			},
			Resources: []string{
				"[Auth0 Documentation](https://auth0.com/docs)",
				"[Auth0 Next.js SDK](https://auth0.com/docs/quickstart/webapp/nextjs)",
			},
		},
		{
			Name:        "Clerk",
			Category:    "authentication",
			Description: "Authentication and user management service",
			Keys: []KeyTemplate{
				{Name: "Publishable Key", EnvVar: "NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY", Required: true},
				{Name: "Secret Key", EnvVar: "CLERK_SECRET_KEY", Required: true},
			},
			SetupSteps: []string{
				"Sign up for Clerk at https://clerk.com",
				"Create a new application",
				"Copy your Publishable Key",
				"Copy your Secret Key",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nimport { ClerkProvider } from '@clerk/nextjs';\n\nfunction MyApp({ Component, pageProps }) {\n  return (\n    <ClerkProvider>\n      <Component {...pageProps} />\n    </ClerkProvider>\n  );\n}\n```",
			CommonIssues: []string{
				"API key not working - ensure keys are from the correct environment",
				"Redirect URLs - configure allowed redirect URLs in Clerk dashboard",
			},
			Resources: []string{
				"[Clerk Documentation](https://clerk.com/docs)",
				"[Clerk Next.js Guide](https://clerk.com/docs/quickstarts/nextjs)",
			},
		},
		{
			Name:        "OpenAI",
			Category:    "ai",
			Description: "AI/ML service for language models and embeddings",
			Keys: []KeyTemplate{
				{Name: "API Key", EnvVar: "OPENAI_API_KEY", Required: false, Format: "sk-..."},
			},
			SetupSteps: []string{
				"Sign up for OpenAI at https://platform.openai.com",
				"Navigate to API Keys",
				"Create a new secret key",
				"Copy the API key (starts with sk-)",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nimport OpenAI from 'openai';\n\nconst openai = new OpenAI({\n  apiKey: process.env.OPENAI_API_KEY,\n});\n\nconst completion = await openai.chat.completions.create({\n  messages: [{ role: 'user', content: 'Hello!' }],\n  model: 'gpt-3.5-turbo',\n});\n```",
			CommonIssues: []string{
				"Rate limits - check your OpenAI plan limits",
				"Invalid API key - ensure key starts with sk-",
				"Billing - ensure your account has credits",
			},
			Resources: []string{
				"[OpenAI Documentation](https://platform.openai.com/docs)",
				"[OpenAI API Reference](https://platform.openai.com/docs/api-reference)",
			},
		},
		{
			Name:        "Sentry",
			Category:    "analytics",
			Description: "Error tracking and performance monitoring",
			Keys: []KeyTemplate{
				{Name: "DSN", EnvVar: "NEXT_PUBLIC_SENTRY_DSN", Required: false},
			},
			SetupSteps: []string{
				"Sign up for Sentry at https://sentry.io",
				"Create a new project",
				"Copy your DSN (Data Source Name)",
				"Add to your .env file",
			},
			CodeExample: "```javascript\nimport * as Sentry from '@sentry/nextjs';\n\nSentry.init({\n  dsn: process.env.NEXT_PUBLIC_SENTRY_DSN,\n  tracesSampleRate: 1.0,\n});\n```",
			CommonIssues: []string{
				"Events not appearing - check DSN is correct",
				"Source maps - configure source maps for better error tracking",
			},
			Resources: []string{
				"[Sentry Documentation](https://docs.sentry.io)",
				"[Sentry Next.js Guide](https://docs.sentry.io/platforms/javascript/guides/nextjs/)",
			},
		},
	}
}
