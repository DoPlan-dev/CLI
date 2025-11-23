package generator

import (
	"fmt"
	"path/filepath"

	"github.com/doplan/cli/internal/utils"
	"github.com/doplan/cli/pkg/models"
)

// BoilerplateGenerator generates boilerplate code for the project
type BoilerplateGenerator struct{}

// Name returns the name of the generator
func (g *BoilerplateGenerator) Name() string {
	return "Boilerplate"
}

// Generate creates boilerplate code based on project type
func (g *BoilerplateGenerator) Generate(request *models.ProjectRequest, projectPath string) error {
	// For now, we'll generate Next.js boilerplate (default for Fullstack)
	if request.ProjectType == "Fullstack" || request.ProjectType == "" {
		return g.generateNextJSBoilerplate(projectPath, request)
	}

	// Future: Add support for other project types
	return nil
}

// generateNextJSBoilerplate generates Next.js boilerplate files
func (g *BoilerplateGenerator) generateNextJSBoilerplate(projectPath string, request *models.ProjectRequest) error {
	// Generate package.json
	if err := generatePackageJSON(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate package.json: %w", err)
	}

	// Generate tsconfig.json
	if err := generateTSConfig(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate tsconfig.json: %w", err)
	}

	// Generate tailwind.config.ts
	if err := generateTailwindConfig(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate tailwind.config.ts: %w", err)
	}

	// Generate ESLint config
	if err := generateESLintConfig(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate ESLint config: %w", err)
	}

	// Generate Next.js app structure
	if err := generateNextJSAppStructure(projectPath, request); err != nil {
		return fmt.Errorf("failed to generate Next.js app structure: %w", err)
	}

	return nil
}

// generatePackageJSON generates package.json
func generatePackageJSON(projectPath string, request *models.ProjectRequest) error {
	content := `{
  "name": "` + request.ProjectName + `",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint"
  },
  "dependencies": {
    "next": "15.2.1",
    "react": "19.0.0",
    "react-dom": "19.0.0"
  },
  "devDependencies": {
    "@types/node": "^20",
    "@types/react": "^19",
    "@types/react-dom": "^19",
    "typescript": "5.6.0",
    "tailwindcss": "3.4.10",
    "postcss": "^8",
    "autoprefixer": "^10",
    "eslint": "^8",
    "eslint-config-next": "15.2.1"
  }
}
`
	path := filepath.Join(projectPath, "package.json")
	return utils.WriteFile(path, []byte(content))
}

// generateTSConfig generates tsconfig.json
func generateTSConfig(projectPath string, request *models.ProjectRequest) error {
	content := `{
  "compilerOptions": {
    "target": "ES2017",
    "lib": ["dom", "dom.iterable", "esnext"],
    "allowJs": true,
    "skipLibCheck": true,
    "strict": true,
    "noEmit": true,
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "bundler",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "jsx": "preserve",
    "incremental": true,
    "plugins": [
      {
        "name": "next"
      }
    ],
    "paths": {
      "@/*": ["./src/*"]
    }
  },
  "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx", ".next/types/**/*.ts"],
  "exclude": ["node_modules"]
}
`
	path := filepath.Join(projectPath, "tsconfig.json")
	return utils.WriteFile(path, []byte(content))
}

// generateTailwindConfig generates tailwind.config.ts
func generateTailwindConfig(projectPath string, request *models.ProjectRequest) error {
	content := `import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
      },
    },
  },
  plugins: [],
};
export default config;
`
	path := filepath.Join(projectPath, "tailwind.config.ts")
	return utils.WriteFile(path, []byte(content))
}

// generateESLintConfig generates .eslintrc.json
func generateESLintConfig(projectPath string, request *models.ProjectRequest) error {
	content := `{
  "extends": ["next/core-web-vitals", "next/typescript"]
}
`
	path := filepath.Join(projectPath, ".eslintrc.json")
	return utils.WriteFile(path, []byte(content))
}

// generateNextJSAppStructure generates Next.js app directory structure
func generateNextJSAppStructure(projectPath string, request *models.ProjectRequest) error {
	// Create src/app directory
	appDir := filepath.Join(projectPath, "src", "app")
	if err := utils.CreateDirectory(appDir); err != nil {
		return fmt.Errorf("failed to create src/app directory: %w", err)
	}

	// Generate layout.tsx
	layoutContent := `import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "` + request.ProjectName + `",
  description: "Generated by DoPlan",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
`
	layoutPath := filepath.Join(appDir, "layout.tsx")
	if err := utils.WriteFile(layoutPath, []byte(layoutContent)); err != nil {
		return fmt.Errorf("failed to generate layout.tsx: %w", err)
	}

	// Generate page.tsx
	pageContent := `export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm">
        <h1 className="text-4xl font-bold text-center mb-8">
          Welcome to ` + request.ProjectName + `
        </h1>
        <p className="text-center text-gray-600">
          Generated by DoPlan - Zero-install AI Project Director
        </p>
      </div>
    </main>
  );
}
`
	pagePath := filepath.Join(appDir, "page.tsx")
	if err := utils.WriteFile(pagePath, []byte(pageContent)); err != nil {
		return fmt.Errorf("failed to generate page.tsx: %w", err)
	}

	// Generate globals.css
	globalsContent := `@tailwind base;
@tailwind components;
@tailwind utilities;

:root {
  --background: #ffffff;
  --foreground: #171717;
}

@media (prefers-color-scheme: dark) {
  :root {
    --background: #0a0a0a;
    --foreground: #ededed;
  }
}

body {
  color: var(--foreground);
  background: var(--background);
  font-family: Arial, Helvetica, sans-serif;
}
`
	globalsPath := filepath.Join(appDir, "globals.css")
	if err := utils.WriteFile(globalsPath, []byte(globalsContent)); err != nil {
		return fmt.Errorf("failed to generate globals.css: %w", err)
	}

	// Create postcss.config.js
	postcssContent := `module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
};
`
	postcssPath := filepath.Join(projectPath, "postcss.config.js")
	if err := utils.WriteFile(postcssPath, []byte(postcssContent)); err != nil {
		return fmt.Errorf("failed to generate postcss.config.js: %w", err)
	}

	return nil
}

// GenerateBoilerplate is a convenience function that creates a BoilerplateGenerator and generates boilerplate
func GenerateBoilerplate(request *models.ProjectRequest, projectPath string) error {
	generator := &BoilerplateGenerator{}
	return generator.Generate(request, projectPath)
}
