# Installation

Get DoPlan CLI running on your system in seconds. DoPlan is designed to be **zero-install** - you can use it immediately with `npx` without any setup!

---

## 🚀 Quick Install (Recommended)

The easiest way to use DoPlan is via `npx` - **no installation required**:

```bash
npx @doplan-dev/cli
```

That's it! This will automatically:
- Download the correct binary for your platform
- Run DoPlan immediately
- Work offline after the first run

---

## 📦 Platform-Specific Installation

### macOS

#### Option 1: Using npx (Recommended - No Installation)
```bash
npx @doplan-dev/cli
```

#### Option 2: Using Homebrew
```bash
# Add tap (if needed)
brew tap doplan-dev/cli

# Install
brew install doplan
```

#### Option 3: Direct Binary Download
1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-darwin-amd64` (Intel) or `doplan-darwin-arm64` (Apple Silicon)
3. Make it executable:
   ```bash
   chmod +x doplan-darwin-amd64
   mv doplan-darwin-amd64 /usr/local/bin/doplan
   ```

#### Option 4: Build from Source
```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

---

### Windows

#### Option 1: Using npx (Recommended - No Installation)
```bash
npx @doplan-dev/cli
```

#### Option 2: Using Scoop
```bash
scoop bucket add doplan https://github.com/DoPlan-dev/scoop-bucket.git
scoop install doplan
```

#### Option 3: Direct Binary Download
1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-windows-amd64.exe`
3. Rename to `doplan.exe` and add to your PATH

#### Option 4: Build from Source
```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan.exe ./cmd/doplan
# Add to PATH or use from current directory
```

---

### Linux

#### Option 1: Using npx (Recommended - No Installation)
```bash
npx @doplan-dev/cli
```

#### Option 2: Direct Binary Download
```bash
# Download latest release
wget https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan-linux-amd64

# Make executable
chmod +x doplan-linux-amd64

# Move to PATH
sudo mv doplan-linux-amd64 /usr/local/bin/doplan
```

#### Option 3: Build from Source
```bash
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

---

## ✅ Verify Installation

After installation, verify it works:

```bash
doplan --version
```

You should see the version number, for example:
```
doplan version 1.0.0
```

Or if using npx:
```bash
npx @doplan-dev/cli --version
```

---

## 🔧 Prerequisites

### Required
- **Node.js** >= 14.0.0 (for npx wrapper)
  - Check: `node --version`
  - Install: [nodejs.org](https://nodejs.org/)

### Optional (for building from source)
- **Go** >= 1.23.0
  - Check: `go version`
  - Install: [golang.org](https://golang.org/)

### Recommended
- **Git** (for version control automation)
  - Check: `git --version`
  - Install: [git-scm.com](https://git-scm.com/)

---

## 🎯 First Run

Once installed, run DoPlan:

```bash
# If installed globally
doplan

# Or with npx
npx @doplan-dev/cli
```

You'll see the interactive wizard that guides you through creating your first project!

---

## 🆘 Troubleshooting

### "Command not found"
- **npx**: Make sure Node.js is installed (`node --version`)
- **Global install**: Check that the binary is in your PATH
- **macOS/Linux**: Try `sudo` for system-wide installation

### "Permission denied"
```bash
# Make executable
chmod +x doplan

# Or on macOS/Linux
sudo chmod +x /usr/local/bin/doplan
```

### "Binary not found for platform"
- Check [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases) for your platform
- Build from source if your platform isn't supported

### "npx download fails"
- Check your internet connection
- Try clearing npm cache: `npm cache clean --force`
- Check npm registry: `npm config get registry`

---

## 🔄 Updating DoPlan

### npx (Automatic)
npx always uses the latest version automatically - no update needed!

### Global Install
```bash
# Homebrew (macOS)
brew upgrade doplan

# Scoop (Windows)
scoop update doplan

# Manual
# Download latest from GitHub Releases and replace binary
```

### From Source
```bash
cd CLI
git pull
go build -o doplan ./cmd/doplan
sudo mv doplan /usr/local/bin/
```

---

## 📝 Next Steps

Now that DoPlan is installed:

1. **[Create Your First Project](./03-First-Project.md)** - See DoPlan in action
2. **[Take a Quick Tour](./04-Quick-Tour.md)** - Explore the features
3. **[Learn the Commands](../02-Commands/)** - Master the workflow

---

**Ready to start?** → [First Project](./03-First-Project.md)

