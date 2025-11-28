# Installation Guide

This guide will help you install DoPlan CLI on your platform. DoPlan CLI is a zero-install tool, meaning you can use it without any installation via `npx`, or install it permanently on your system.

---

## Prerequisites

Before installing DoPlan CLI, ensure you have:

- **Node.js** >= 14.0.0 (required for `npx` wrapper)
- **Go** >= 1.23.0 (only required if building from source)

### Checking Your Prerequisites

**Check Node.js version:**
```bash
node --version
```

**Check Go version (if building from source):**
```bash
go version
```

---

## Quick Install (Recommended)

The easiest way to use DoPlan CLI is via `npx` - **no installation required**!

```bash
npx @doplan-dev/cli
```

This will automatically:
1. Download the correct binary for your platform
2. Run DoPlan CLI immediately
3. No files are installed on your system

**Note**: The first run may take a few seconds to download the binary. Subsequent runs are instant.

---

## Platform-Specific Installation

### 🍎 macOS

#### Option 1: Using Homebrew (Recommended for Permanent Installation)

```bash
# Add tap (if needed)
brew tap doplan-dev/cli

# Install
brew install doplan
```

**Verify installation:**
```bash
doplan --version
```

#### Option 2: Using npx (No Installation)

```bash
npx @doplan-dev/cli
```

#### Option 3: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download the appropriate binary:
   - `doplan-darwin-amd64` for Intel Macs
   - `doplan-darwin-arm64` for Apple Silicon (M1/M2/M3)
3. Make it executable and move to PATH:

```bash
# For Intel Macs
chmod +x doplan-darwin-amd64
mv doplan-darwin-amd64 /usr/local/bin/doplan

# For Apple Silicon
chmod +x doplan-darwin-arm64
mv doplan-darwin-arm64 /usr/local/bin/doplan
```

#### Option 4: Build from Source

```bash
# Clone the repository
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI

# Build the binary
go build -o doplan ./cmd/doplan

# Install to system
sudo mv doplan /usr/local/bin/
```

---

### 🪟 Windows

#### Option 1: Using Scoop (Recommended for Permanent Installation)

```bash
# Add bucket
scoop bucket add doplan https://github.com/DoPlan-dev/scoop-bucket.git

# Install
scoop install doplan
```

**Verify installation:**
```powershell
doplan --version
```

#### Option 2: Using npx (No Installation)

```bash
npx @doplan-dev/cli
```

#### Option 3: Direct Binary Download

1. Visit [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Download `doplan-windows-amd64.exe`
3. Rename to `doplan.exe`
4. Add to your PATH:
   - Option A: Move to a directory already in PATH (e.g., `C:\Windows\System32`)
   - Option B: Add the directory containing `doplan.exe` to your PATH environment variable

**Adding to PATH:**
1. Open System Properties → Environment Variables
2. Edit the `Path` variable under User variables
3. Add the directory containing `doplan.exe`
4. Click OK and restart your terminal

#### Option 4: Build from Source

```powershell
# Clone the repository
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI

# Build the binary
go build -o doplan.exe ./cmd/doplan

# Use from current directory or add to PATH
```

---

### 🐧 Linux

#### Option 1: Using npx (No Installation - Recommended)

```bash
npx @doplan-dev/cli
```

#### Option 2: Direct Binary Download

```bash
# Download latest release
curl -L https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan-linux-amd64 -o doplan

# Make executable
chmod +x doplan

# Move to PATH
sudo mv doplan /usr/local/bin/
```

**For ARM64 Linux:**
```bash
curl -L https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan-linux-arm64 -o doplan
chmod +x doplan
sudo mv doplan /usr/local/bin/
```

#### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/DoPlan-dev/CLI.git
cd CLI

# Build the binary
go build -o doplan ./cmd/doplan

# Install to system
sudo mv doplan /usr/local/bin/
```

#### Option 4: Using Package Managers

**Debian/Ubuntu:**

When `.deb` packages are available:
```bash
# Download .deb package
wget https://github.com/DoPlan-dev/CLI/releases/latest/download/doplan_amd64.deb

# Install
sudo dpkg -i doplan_amd64.deb
```

**Arch Linux:**

When AUR package is available:
```bash
# Using yay
yay -S doplan-cli

# Or using paru
paru -S doplan-cli
```

---

### 🐳 Docker

DoPlan CLI is available as a Docker image for containerized environments.

#### Basic Usage

```bash
docker run --rm -it -v $(pwd):/workspace doplan/cli
```

#### With Custom Working Directory

```bash
docker run --rm -it \
  -v $(pwd):/workspace \
  -w /workspace \
  doplan/cli
```

#### Using Specific Version

```bash
docker run --rm -it \
  -v $(pwd):/workspace \
  doplan/cli:1.0.0
```

**Note**: Replace `$(pwd)` with `%cd%` on Windows PowerShell.

---

## Verification

After installation, verify that DoPlan CLI is working correctly:

```bash
doplan --version
```

You should see output like:
```
doplan version 1.0.4
```

### Test Run

Try running DoPlan CLI to ensure everything works:

```bash
# Using npx (no installation)
npx @doplan-dev/cli

# Or if installed
doplan
```

You should see the interactive wizard start.

---

## Troubleshooting Installation Issues

### Issue: "Command not found" or "doplan: command not found"

**Solution:**
1. Verify the binary is in your PATH:
   ```bash
   # macOS/Linux
   which doplan
   
   # Windows
   where doplan
   ```

2. If not found, add the binary location to your PATH:
   - **macOS/Linux**: Add to `~/.bashrc`, `~/.zshrc`, or `~/.profile`
   - **Windows**: Add to System Environment Variables

3. Restart your terminal after updating PATH

### Issue: "Permission denied" (macOS/Linux)

**Solution:**
```bash
chmod +x doplan
```

Or if installed to system location:
```bash
sudo chmod +x /usr/local/bin/doplan
```

### Issue: "Binary not found" with npx

**Solution:**
1. Clear npm cache:
   ```bash
   npm cache clean --force
   ```

2. Try again:
   ```bash
   npx @doplan-dev/cli
   ```

3. Check your internet connection (first run requires download)

### Issue: Wrong architecture binary downloaded

**Solution:**
- **macOS**: Ensure you download the correct binary for your chip (Intel vs Apple Silicon)
- **Linux**: Check if you need `amd64` or `arm64` version
- **Windows**: Ensure you're using 64-bit Windows

### Issue: Homebrew installation fails

**Solution:**
1. Update Homebrew:
   ```bash
   brew update
   ```

2. Check if tap exists:
   ```bash
   brew tap doplan-dev/cli
   ```

3. Try installing again:
   ```bash
   brew install doplan
   ```

### Issue: Build from source fails

**Solution:**
1. Ensure Go >= 1.23.0 is installed:
   ```bash
   go version
   ```

2. Ensure you have internet access (for downloading dependencies)

3. Try building with verbose output:
   ```bash
   go build -v -o doplan ./cmd/doplan
   ```

### Issue: Docker image not found

**Solution:**
1. Pull the image explicitly:
   ```bash
   docker pull doplan/cli:latest
   ```

2. Check available tags:
   ```bash
   docker search doplan/cli
   ```

---

## Updating DoPlan CLI

### Using npx

No update needed! `npx` always uses the latest version.

### Using Homebrew

```bash
brew upgrade doplan
```

### Using Scoop

```bash
scoop update doplan
```

### Using Direct Binary

1. Download the latest release from [GitHub Releases](https://github.com/DoPlan-dev/CLI/releases/latest)
2. Replace the existing binary
3. Verify:
   ```bash
   doplan --version
   ```

### Using Docker

```bash
docker pull doplan/cli:latest
```

---

## Uninstallation

### Homebrew

```bash
brew uninstall doplan
```

### Scoop

```bash
scoop uninstall doplan
```

### Manual Installation

Remove the binary from your PATH:
```bash
# macOS/Linux
sudo rm /usr/local/bin/doplan

# Windows
# Delete doplan.exe from the directory in your PATH
```

---

## Next Steps

After successful installation:

1. **[Quick Start Guide](Quick-Start)** - Create your first project
2. **[Commands Reference](Commands)** - Learn all available commands
3. **[Workflow Guide](Workflow)** - Understand the development workflow

---

## Related Pages

- [Quick Start](Quick-Start) - Get started with your first project
- [Troubleshooting](Troubleshooting) - Common issues and solutions
- [FAQ](FAQ) - Frequently asked questions
- [Home](Home) - Wiki home page

---

**Last Updated**: 2025  
**Maintained By**: Documentation Team

