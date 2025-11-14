#!/usr/bin/env node

/**
 * DoPlan CLI wrapper script
 * Downloads and executes the platform-specific binary from GitHub releases
 */

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const BINARY_NAME = 'doplan';
const REPO_OWNER = 'DoPlan-dev';
const REPO_NAME = 'CLI';
const VERSION = require('../package.json').version;

// Get platform information
function getPlatformInfo() {
  const platform = os.platform();
  const arch = os.arch();
  
  let osName, archName;
  
  switch (platform) {
    case 'darwin':
      osName = 'darwin';
      break;
    case 'linux':
      osName = 'linux';
      break;
    case 'win32':
      osName = 'windows';
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }
  
  switch (arch) {
    case 'x64':
      archName = 'amd64';
      break;
    case 'arm64':
      archName = 'arm64';
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }
  
  return { osName, archName, platform, arch };
}

// Get binary path
function getBinaryPath() {
  const { osName, archName, platform } = getPlatformInfo();
  const binaryDir = path.join(__dirname, '..', 'bin', `${osName}-${archName}`);
  const ext = platform === 'win32' ? '.exe' : '';
  return path.join(binaryDir, `${BINARY_NAME}${ext}`);
}

// Download binary from GitHub releases
function downloadBinary() {
  const { osName, archName, platform } = getPlatformInfo();
  const binaryDir = path.join(__dirname, '..', 'bin', `${osName}-${archName}`);
  const ext = platform === 'win32' ? '.exe' : '';
  const binaryPath = path.join(binaryDir, `${BINARY_NAME}${ext}`);
  
  // Create binary directory if it doesn't exist
  if (!fs.existsSync(binaryDir)) {
    fs.mkdirSync(binaryDir, { recursive: true });
  }
  
  // Check if binary already exists
  if (fs.existsSync(binaryPath)) {
    try {
      // Verify binary is executable
      fs.chmodSync(binaryPath, '755');
      return binaryPath;
    } catch (err) {
      // If chmod fails, try to download again
      console.warn('Binary exists but may be corrupted, re-downloading...');
    }
  }
  
  // Download from GitHub releases
  const archiveName = `${BINARY_NAME}_${VERSION}_${osName}_${archName}.tar.gz`;
  const releaseUrl = `https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/v${VERSION}/${archiveName}`;
  
  console.log(`Downloading DoPlan CLI v${VERSION} for ${osName}-${archName}...`);
  console.log(`URL: ${releaseUrl}`);
  
  try {
    // Download and extract
    const tempDir = path.join(os.tmpdir(), `doplan-${Date.now()}`);
    fs.mkdirSync(tempDir, { recursive: true });
    
    // Use curl or wget to download
    let downloadCmd;
    if (platform === 'win32') {
      // Windows: use PowerShell or curl if available
      downloadCmd = `powershell -Command "Invoke-WebRequest -Uri '${releaseUrl}' -OutFile '${path.join(tempDir, archiveName)}'"`;
    } else {
      // Unix-like: use curl or wget
      downloadCmd = `curl -L -o '${path.join(tempDir, archiveName)}' '${releaseUrl}'`;
    }
    
    execSync(downloadCmd, { stdio: 'inherit' });
    
    // Extract archive
    if (platform === 'win32') {
      // Windows: use tar if available, or 7zip
      try {
        execSync(`tar -xzf '${path.join(tempDir, archiveName)}' -C '${tempDir}'`, { stdio: 'inherit' });
      } catch (err) {
        // Try PowerShell extraction
        execSync(`powershell -Command "Expand-Archive -Path '${path.join(tempDir, archiveName)}' -DestinationPath '${tempDir}'"`, { stdio: 'inherit' });
      }
    } else {
      execSync(`tar -xzf '${path.join(tempDir, archiveName)}' -C '${tempDir}'`, { stdio: 'inherit' });
    }
    
    // Find the binary in extracted files (could be in root or subdirectory)
    let extractedBinary = path.join(tempDir, BINARY_NAME + ext);
    
    // If not in root, search for it recursively
    if (!fs.existsSync(extractedBinary)) {
      function findBinary(dir) {
        const files = fs.readdirSync(dir);
        for (const file of files) {
          const filePath = path.join(dir, file);
          const stat = fs.statSync(filePath);
          if (stat.isDirectory()) {
            const found = findBinary(filePath);
            if (found) return found;
          } else if (file === BINARY_NAME + ext) {
            return filePath;
          }
        }
        return null;
      }
      const found = findBinary(tempDir);
      if (found) {
        extractedBinary = found;
      }
    }
    
    if (fs.existsSync(extractedBinary)) {
      fs.copyFileSync(extractedBinary, binaryPath);
      fs.chmodSync(binaryPath, '755');
      
      // Cleanup
      fs.rmSync(tempDir, { recursive: true, force: true });
      
      console.log(`Successfully downloaded DoPlan CLI v${VERSION}`);
      return binaryPath;
    } else {
      throw new Error(`Binary not found in downloaded archive. Searched in: ${tempDir}`);
    }
  } catch (error) {
    console.error(`Failed to download binary: ${error.message}`);
    console.error(`\nPlease install manually from: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases`);
    process.exit(1);
  }
}

// Main execution
function main() {
  let binaryPath = getBinaryPath();
  
  // If binary doesn't exist, download it
  if (!fs.existsSync(binaryPath)) {
    binaryPath = downloadBinary();
  }
  
  // Execute the binary with all arguments
  const args = process.argv.slice(2);
  const spawn = require('child_process').spawn;
  
  const child = spawn(binaryPath, args, {
    stdio: 'inherit',
    shell: false
  });
  
  child.on('error', (error) => {
    if (error.code === 'ENOENT') {
      console.error(`Binary not found: ${binaryPath}`);
      console.error('Attempting to download...');
      binaryPath = downloadBinary();
      // Retry execution
      const retry = spawn(binaryPath, args, {
        stdio: 'inherit',
        shell: false
      });
      retry.on('error', (retryError) => {
        console.error(`Failed to execute binary: ${retryError.message}`);
        process.exit(1);
      });
      retry.on('exit', (code) => {
        process.exit(code || 0);
      });
    } else {
      console.error(`Error executing binary: ${error.message}`);
      process.exit(1);
    }
  });
  
  child.on('exit', (code) => {
    process.exit(code || 0);
  });
}

// Run if executed directly
if (require.main === module) {
  main();
}

module.exports = { getBinaryPath, downloadBinary, getPlatformInfo };

