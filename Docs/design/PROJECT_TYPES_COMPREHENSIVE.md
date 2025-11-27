# Comprehensive Project Types List

## Overview

DoPlan supports a comprehensive range of project types, each with tailored questions and workflows. This document lists all supported project types and their characteristics.

## Project Type Categories

### 1. Web Projects

#### Website
- **Description**: Static or dynamic websites, portfolios, landing pages
- **Examples**: Company website, personal portfolio, agency site, blog
- **Technologies**: HTML, CSS, JavaScript, React, Next.js, Vue, Angular
- **Key Questions**: Pages needed, content structure, SEO requirements

#### Web App
- **Description**: Progressive Web Apps (PWA), Single Page Apps (SPA), web applications
- **Examples**: Gmail, Twitter, Trello, Notion
- **Technologies**: React, Vue, Angular, Svelte, Next.js, Nuxt
- **Key Questions**: Offline support, app-like features, performance requirements

#### SaaS
- **Description**: Software-as-a-Service, subscription-based applications
- **Examples**: Slack, Dropbox, Salesforce, HubSpot
- **Technologies**: Full-stack frameworks, cloud services, payment integration
- **Key Questions**: Pricing model, subscription tiers, multi-tenancy, scalability

### 2. Mobile Projects

#### Mobile App (Cross-Platform)
- **Description**: Apps that work on multiple platforms
- **Examples**: Instagram, WhatsApp, Uber
- **Technologies**: React Native, Flutter, Xamarin, Ionic
- **Key Questions**: Target platforms, native features needed, performance requirements

#### iOS App
- **Description**: iPhone/iPad applications (native iOS)
- **Examples**: iOS-only apps, App Store apps
- **Technologies**: Swift, Objective-C, SwiftUI, UIKit
- **Key Questions**: iOS version support, App Store requirements, device capabilities

#### Android App
- **Description**: Android phone/tablet applications (native Android)
- **Examples**: Android-only apps, Google Play apps
- **Technologies**: Java, Kotlin, Jetpack Compose, Android SDK
- **Key Questions**: Android version support, Google Play requirements, device fragmentation

### 3. Desktop Projects

#### Desktop App (Cross-Platform)
- **Description**: Apps that work on Windows, macOS, and Linux
- **Examples**: VS Code, Discord, Spotify
- **Technologies**: Electron, Tauri, Flutter Desktop, Qt
- **Key Questions**: Target platforms, native features, distribution method

#### Windows App
- **Description**: Native Windows desktop applications
- **Examples**: Windows-only software, .NET apps
- **Technologies**: .NET, C#, WPF, WinUI, UWP, WinForms
- **Key Questions**: Windows version support, installer type, Windows Store distribution

#### macOS App
- **Description**: Native macOS desktop applications
- **Examples**: macOS-only software, Mac App Store apps
- **Technologies**: Swift, Objective-C, SwiftUI, AppKit
- **Key Questions**: macOS version support, Mac App Store requirements, Apple Silicon support

#### Linux App
- **Description**: Native Linux desktop applications
- **Examples**: Linux-only software, open-source tools
- **Technologies**: GTK, Qt, C/C++, Python, Electron
- **Key Questions**: Distribution support (Ubuntu, Fedora, etc.), package format (deb, rpm, AppImage)

### 4. Command-Line & Tools

#### CLI Tool
- **Description**: Command-line interface tools and utilities
- **Examples**: Git, npm, Docker CLI, AWS CLI
- **Technologies**: Go, Rust, Python, Node.js, Shell scripts
- **Key Questions**: Command structure, help system, configuration, output format

#### Library/Package
- **Description**: Code libraries, npm packages, Python packages, SDKs
- **Examples**: React, Lodash, Express, TensorFlow
- **Technologies**: Language-specific (JavaScript, Python, Go, Rust, etc.)
- **Key Questions**: API design, versioning, dependencies, distribution, backward compatibility

#### Framework
- **Description**: Development frameworks, plugins, extensions
- **Examples**: Next.js, Django, Rails, Spring Boot
- **Technologies**: Framework-specific
- **Key Questions**: Plugin architecture, extension points, configuration system

### 5. Backend & Services

#### API
- **Description**: REST APIs, GraphQL APIs, backend services
- **Examples**: Stripe API, GitHub API, custom APIs
- **Technologies**: Node.js, Python (FastAPI, Flask), Go, Java (Spring), Ruby (Rails)
- **Key Questions**: API design, authentication, rate limiting, documentation

#### Microservice
- **Description**: Microservices architecture, distributed systems
- **Examples**: Netflix, Amazon, Uber backend
- **Technologies**: Docker, Kubernetes, service mesh, message queues
- **Key Questions**: Service boundaries, communication patterns, deployment strategy

### 6. Specialized Projects

#### Game
- **Description**: Video games, game engines, game tools
- **Examples**: Unity games, Unreal Engine games, indie games
- **Technologies**: Unity (C#), Unreal (C++), Godot, Phaser, Three.js
- **Key Questions**: Platform targets, graphics requirements, multiplayer, monetization

#### Embedded/IoT
- **Description**: Embedded systems, IoT devices, firmware
- **Examples**: Smart home devices, wearables, industrial automation
- **Technologies**: C, C++, Rust, Arduino, Raspberry Pi, ESP32
- **Key Questions**: Hardware constraints, real-time requirements, power consumption

#### Data Science/ML
- **Description**: Machine learning, data analysis, AI projects
- **Examples**: ML models, data pipelines, analytics tools
- **Technologies**: Python (TensorFlow, PyTorch), R, Jupyter, pandas
- **Key Questions**: Data sources, model requirements, training infrastructure, deployment

#### Cloud
- **Description**: Cloud-native applications, serverless, cloud services
- **Examples**: AWS Lambda functions, Azure Functions, Google Cloud Functions
- **Technologies**: Serverless frameworks, cloud SDKs, container orchestration
- **Key Questions**: Cloud provider, scaling strategy, cost optimization

#### DevOps
- **Description**: DevOps tools, CI/CD pipelines, infrastructure as code
- **Examples**: GitHub Actions, Jenkins plugins, Terraform modules
- **Technologies**: YAML, Bash, Python, Go, Terraform, Ansible
- **Key Questions**: Automation scope, integration points, deployment targets

### 7. Patches & Updates

#### Patch (Windows)
- **Description**: Windows application patches, updates, hotfixes
- **Examples**: Windows app updates, .NET patches
- **Key Questions**: Change scope, backward compatibility, Windows version support

#### Patch (macOS)
- **Description**: macOS application patches, updates, hotfixes
- **Examples**: macOS app updates, Mac App Store updates
- **Key Questions**: Change scope, backward compatibility, macOS version support

#### Patch (Linux)
- **Description**: Linux application patches, updates, hotfixes
- **Examples**: Linux package updates, distribution patches
- **Key Questions**: Change scope, backward compatibility, distribution support

#### Patch (Web)
- **Description**: Web application patches, updates, hotfixes
- **Examples**: Website updates, web app bug fixes
- **Key Questions**: Change scope, browser compatibility, deployment strategy

## Project Type Detection

The `/meeting` command automatically detects project type from `.do/system/IDEA.md` by analyzing:
- Keywords in the project description
- Technology mentions
- Platform references
- Use case descriptions

## Question Library Structure

Each project type has tailored questions organized by experience level:

```
.do/core/brainstorm/
├── beginner/{project-type}/phase-*.md
├── intermediate/{project-type}/phase-*.md
└── advanced/{project-type}/phase-*.md
```

## Examples of Type-Specific Questions

### iOS App (Beginner)
- "Which iPhones should your app work on?"
- "Do you want it in the App Store?"
- "What's the main thing your app does?"

### CLI Tool (Advanced)
- "What's the command structure and subcommand hierarchy?"
- "How will you handle configuration management?"
- "What's the output format and parsing strategy?"

### Game (Intermediate)
- "What platforms will the game run on?"
- "What's the target frame rate and graphics quality?"
- "Do you need multiplayer support?"

### Microservice (Advanced)
- "What are the service boundaries and domain models?"
- "What communication patterns will you use? (REST, gRPC, message queues)"
- "What's the deployment and scaling strategy?"

## Fallback

If project type cannot be determined, the system uses `general/` templates with universal questions that work for any project type.

---

**Generated by: DoPlan CLI v1.2.0**

**Sub-Agent: Documentation Lead**

**Date: 2025-01-15**

