#!/usr/bin/env node
/**
 * DoPlan CLI wrapper
 * Executes the platform-specific binary
 */

const path = require('path');
const { spawn } = require('child_process');
const os = require('os');
const fs = require('fs');

const binDir = path.join(__dirname);
const platform = os.platform();
const arch = os.arch();

let binaryName = 'doplan';
if (platform === 'win32') {
  binaryName = 'doplan.exe';
}

const binaryPath = path.join(binDir, binaryName);

if (!fs.existsSync(binaryPath)) {
  console.log('Binary not found. Attempting to download...');
  const downloadScript = path.join(__dirname, '..', 'scripts', 'download.js');
  if (fs.existsSync(downloadScript)) {
    try {
      require(downloadScript).main();
      // Check again after download attempt
      if (!fs.existsSync(binaryPath)) {
        console.error('Error: DoPlan CLI binary not found.');
        console.error(`Expected: ${binaryPath}`);
        console.error('Please ensure the GitHub release exists with binaries.');
        console.error('Or build from source: https://github.com/DoPlan-dev/CLI');
        process.exit(1);
      }
    } catch (error) {
      console.error('Error downloading binary:', error.message);
      console.error('Please ensure the GitHub release exists with binaries.');
      console.error('Or build from source: https://github.com/DoPlan-dev/CLI');
      process.exit(1);
    }
  } else {
    console.error('Error: DoPlan CLI binary not found.');
    console.error(`Expected: ${binaryPath}`);
    console.error('Please run: npm install');
    process.exit(1);
  }
}

// Make executable on Unix systems
if (platform !== 'win32') {
  try {
    fs.chmodSync(binaryPath, '755');
  } catch (e) {
    // Ignore errors
  }
}

// Spawn the binary with all arguments
const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: 'inherit',
  cwd: process.cwd(),
});

child.on('error', (error) => {
  console.error('Error executing DoPlan CLI:', error.message);
  process.exit(1);
});

child.on('exit', (code) => {
  process.exit(code || 0);
});

