package deployment

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Platform represents a deployment platform
type Platform struct {
	Name        string
	Description string
	CLICommand  string
	DetectFunc  func(string) bool
	DeployFunc  func(string, map[string]string) error
}

// DetectPlatform detects which platform is best suited for the project
func DetectPlatform(projectRoot string) (string, error) {
	// Check for platform-specific config files
	if _, err := os.Stat(filepath.Join(projectRoot, "vercel.json")); err == nil {
		return "vercel", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "netlify.toml")); err == nil {
		return "netlify", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "railway.json")); err == nil {
		return "railway", nil
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "render.yaml")); err == nil {
		return "render", nil
	}

	// Detect by project type
	if _, err := os.Stat(filepath.Join(projectRoot, "package.json")); err == nil {
		// Check if it's Next.js
		data, err := os.ReadFile(filepath.Join(projectRoot, "package.json"))
		if err == nil {
			if strings.Contains(string(data), "next") {
				return "vercel", nil
			}
		}
		return "netlify", nil // Default for Node.js
	}

	if _, err := os.Stat(filepath.Join(projectRoot, "Dockerfile")); err == nil {
		return "railway", nil // Docker apps work well on Railway
	}

	return "vercel", nil // Default
}

// DeployToVercel deploys to Vercel
func DeployToVercel(projectRoot string, envVars map[string]string) error {
	// Check if Vercel CLI is installed
	if _, err := exec.LookPath("vercel"); err != nil {
		return fmt.Errorf("Vercel CLI not installed. Install with: npm i -g vercel")
	}

	fmt.Println("🚀 Deploying to Vercel...")

	// Set environment variables
	for key, value := range envVars {
		cmd := exec.Command("vercel", "env", "add", key, "production")
		cmd.Dir = projectRoot
		cmd.Stdin = strings.NewReader(value + "\n")
		if err := cmd.Run(); err != nil {
			fmt.Printf("⚠️  Warning: Failed to set env var %s: %v\n", key, err)
		}
	}

	// Deploy
	cmd := exec.Command("vercel", "--prod", "--yes")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("deployment failed: %w", err)
	}

	fmt.Println("✅ Successfully deployed to Vercel!")
	return nil
}

// DeployToNetlify deploys to Netlify
func DeployToNetlify(projectRoot string, envVars map[string]string) error {
	// Check if Netlify CLI is installed
	if _, err := exec.LookPath("netlify"); err != nil {
		return fmt.Errorf("Netlify CLI not installed. Install with: npm i -g netlify-cli")
	}

	fmt.Println("🚀 Deploying to Netlify...")

	// Set environment variables
	for key, value := range envVars {
		cmd := exec.Command("netlify", "env:set", key, value)
		cmd.Dir = projectRoot
		if err := cmd.Run(); err != nil {
			fmt.Printf("⚠️  Warning: Failed to set env var %s: %v\n", key, err)
		}
	}

	// Deploy
	cmd := exec.Command("netlify", "deploy", "--prod")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("deployment failed: %w", err)
	}

	fmt.Println("✅ Successfully deployed to Netlify!")
	return nil
}

// DeployToRailway deploys to Railway
func DeployToRailway(projectRoot string, envVars map[string]string) error {
	// Check if Railway CLI is installed
	if _, err := exec.LookPath("railway"); err != nil {
		return fmt.Errorf("Railway CLI not installed. Install with: npm i -g @railway/cli")
	}

	fmt.Println("🚀 Deploying to Railway...")

	// Set environment variables
	for key, value := range envVars {
		cmd := exec.Command("railway", "variables", "set", fmt.Sprintf("%s=%s", key, value))
		cmd.Dir = projectRoot
		if err := cmd.Run(); err != nil {
			fmt.Printf("⚠️  Warning: Failed to set env var %s: %v\n", key, err)
		}
	}

	// Deploy
	cmd := exec.Command("railway", "up")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("deployment failed: %w", err)
	}

	fmt.Println("✅ Successfully deployed to Railway!")
	return nil
}

// DeployToRender deploys to Render
func DeployToRender(projectRoot string, envVars map[string]string) error {
	fmt.Println("🚀 Deploying to Render...")
	fmt.Println("📝 Note: Render deployment requires manual setup via Render dashboard")
	fmt.Println("   1. Go to https://dashboard.render.com")
	fmt.Println("   2. Create a new service")
	fmt.Println("   3. Connect your GitHub repository")
	fmt.Println("   4. Configure environment variables")
	fmt.Println("   5. Deploy!")

	// Check for render.yaml
	renderYAML := filepath.Join(projectRoot, "render.yaml")
	if _, err := os.Stat(renderYAML); os.IsNotExist(err) {
		fmt.Println("💡 Tip: Create render.yaml for automated deployments")
	}

	return nil
}

// DeployToCoolify deploys to Coolify (self-hosted)
func DeployToCoolify(projectRoot string, envVars map[string]string) error {
	fmt.Println("🚀 Deploying to Coolify...")
	fmt.Println("📝 Note: Coolify deployment requires manual setup")
	fmt.Println("   1. Access your Coolify instance")
	fmt.Println("   2. Create a new application")
	fmt.Println("   3. Connect your Git repository")
	fmt.Println("   4. Configure environment variables")
	fmt.Println("   5. Deploy!")

	return nil
}

// DeployWithDocker deploys using Docker
func DeployWithDocker(projectRoot string, envVars map[string]string) error {
	fmt.Println("🐳 Building Docker image...")

	// Check for Dockerfile
	dockerfile := filepath.Join(projectRoot, "Dockerfile")
	if _, err := os.Stat(dockerfile); os.IsNotExist(err) {
		return fmt.Errorf("Dockerfile not found")
	}

	// Build image
	cmd := exec.Command("docker", "build", "-t", "doplan-app", ".")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}

	fmt.Println("✅ Docker image built successfully!")
	fmt.Println("💡 Next steps:")
	fmt.Println("   1. Tag the image: docker tag doplan-app <registry>/doplan-app:latest")
	fmt.Println("   2. Push to registry: docker push <registry>/doplan-app:latest")
	fmt.Println("   3. Deploy to your platform")

	return nil
}
