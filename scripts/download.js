#!/usr/bin/env node
/**
 * Download script for DoPlan CLI
 * Downloads the appropriate binary for the current platform
 */

const https = require('https');
const fs = require('fs');
const path = require('path');
const os = require('os');

const VERSION = process.env.DOPLAN_VERSION || 'latest';

async function getLatestVersion() {
  return new Promise((resolve, reject) => {
    https.get('https://api.github.com/repos/DoPlan-dev/CLI/releases/latest', {
      headers: {
        'User-Agent': 'DoPlan-CLI-Downloader'
      }
    }, (response) => {
      let data = '';
      
      response.on('data', (chunk) => {
        data += chunk;
      });
      
      response.on('end', () => {
        if (response.statusCode === 200) {
          try {
            const release = JSON.parse(data);
            // Extract version from tag (e.g., "v1.0.0" -> "1.0.0")
            const version = release.tag_name.replace(/^v/, '');
            resolve(version);
          } catch (e) {
            reject(new Error('Failed to parse release data'));
          }
        } else {
          reject(new Error(`Failed to fetch latest version: ${response.statusCode}`));
        }
      });
    }).on('error', reject);
  });
}

function getPlatformInfo() {
  const platform = os.platform();
  const arch = os.arch();
  
  let goos, goarch, ext;
  
  switch (platform) {
    case 'darwin':
      goos = 'darwin';
      goarch = arch === 'arm64' ? 'arm64' : 'amd64';
      ext = 'tar.gz';
      break;
    case 'linux':
      goos = 'linux';
      goarch = arch === 'arm64' ? 'arm64' : 'amd64';
      ext = 'tar.gz';
      break;
    case 'win32':
      goos = 'windows';
      goarch = 'amd64';
      ext = 'zip';
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }
  
  return { goos, goarch, ext };
}

function downloadFile(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    
    https.get(url, (response) => {
      if (response.statusCode === 302 || response.statusCode === 301) {
        // Follow redirect
        return downloadFile(response.headers.location, dest).then(resolve).catch(reject);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode}`));
        return;
      }
      
      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        resolve();
      });
    }).on('error', (err) => {
      fs.unlink(dest, () => {});
      reject(err);
    });
  });
}

function extractArchive(archivePath, ext) {
  const { execSync } = require('child_process');
  const binDir = path.join(__dirname, '..', 'bin');
  
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }
  
  if (ext === 'zip') {
    // Windows - use unzip or PowerShell
    try {
      execSync(`unzip -o "${archivePath}" -d "${binDir}"`, { stdio: 'inherit' });
    } catch (e) {
      // Fallback to PowerShell on Windows
      execSync(`powershell -Command "Expand-Archive -Path '${archivePath}' -DestinationPath '${binDir}' -Force"`, { stdio: 'inherit' });
    }
  } else {
    // Unix - use tar
    execSync(`tar -xzf "${archivePath}" -C "${binDir}"`, { stdio: 'inherit' });
  }
  
  // Make binary executable (Unix)
  if (ext !== 'zip') {
    const binaryPath = path.join(binDir, 'doplan');
    if (fs.existsSync(binaryPath)) {
      fs.chmodSync(binaryPath, '755');
    }
  }
}

async function main() {
  try {
    // Get the actual version to use
    let versionToUse = VERSION;
    if (VERSION === 'latest') {
      console.log('Fetching latest release version...');
      versionToUse = await getLatestVersion();
      console.log(`Latest version: ${versionToUse}`);
    }
    
    const { goos, goarch, ext } = getPlatformInfo();
    const archiveName = `doplan-${versionToUse}-${goos}-${goarch}.${ext}`;
    const baseUrl = `https://github.com/DoPlan-dev/CLI/releases/download/v${versionToUse}`;
    const url = `${baseUrl}/${archiveName}`;
    
    console.log(`Downloading DoPlan CLI for ${goos}/${goarch}...`);
    console.log(`URL: ${url}`);
    
    const archivePath = path.join(__dirname, '..', archiveName);
    
    await downloadFile(url, archivePath);
    console.log('Download complete!');
    
    console.log('Extracting...');
    extractArchive(archivePath, ext);
    
    // Clean up archive
    fs.unlinkSync(archivePath);
    
    console.log('Installation complete!');
  } catch (error) {
    // In CI environments or development, don't fail
    const isCI = process.env.CI === 'true' || process.env.GITHUB_ACTIONS === 'true' || process.env.CI === '1';
    const isDev = process.env.NODE_ENV === 'development' || !process.env.npm_config_user_config;
    
    if (isCI || isDev || error.message.includes('404')) {
      console.warn('Warning: Could not download binary:', error.message);
      if (error.message.includes('404')) {
        console.warn('The GitHub release may not exist yet or binaries are not available.');
        console.warn('The binary will be downloaded automatically when you run the CLI.');
        console.warn('Or check: https://github.com/DoPlan-dev/CLI/releases');
      } else {
        console.warn('Binary will be downloaded on first use or when release is available.');
      }
      process.exit(0);
    }
    // For production npm installs, fail so they know something is wrong
    console.error('Installation failed:', error.message);
    console.error('Please check: https://github.com/DoPlan-dev/CLI/releases');
    console.error('Or build from source: https://github.com/DoPlan-dev/CLI');
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}

module.exports = { main };

