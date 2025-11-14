#!/usr/bin/env node

/**
 * Build script for npm package
 * Prepares the package for publishing by verifying required files
 */

const fs = require('fs');
const path = require('path');

console.log('DoPlan CLI: Building npm package...');

// Verify required files exist
const requiredFiles = [
  'bin/doplan.js',
  'scripts/postinstall.js',
  'scripts/prepublish.js',
  'README.md',
  'LICENSE'
];

let hasErrors = false;

for (const file of requiredFiles) {
  const filePath = path.join(__dirname, '..', file);
  if (!fs.existsSync(filePath)) {
    console.error(`Error: Required file missing: ${file}`);
    hasErrors = true;
  } else {
    console.log(`✓ Found: ${file}`);
  }
}

// Verify bin/doplan.js is executable (on Unix systems)
const binPath = path.join(__dirname, '..', 'bin', 'doplan.js');
if (fs.existsSync(binPath)) {
  try {
    // Check if file has shebang
    const content = fs.readFileSync(binPath, 'utf8');
    if (!content.startsWith('#!/usr/bin/env node')) {
      console.warn('Warning: bin/doplan.js missing shebang line');
    }
  } catch (error) {
    console.warn(`Warning: Could not read bin/doplan.js: ${error.message}`);
  }
}

if (hasErrors) {
  console.error('\nBuild failed. Please fix the errors above.');
  process.exit(1);
}

console.log('\n✅ Build completed successfully!');
console.log('Package is ready for publishing.');

