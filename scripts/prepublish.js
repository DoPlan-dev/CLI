#!/usr/bin/env node

/**
 * Pre-publish script for npm package
 * Validates package before publishing
 */

const fs = require('fs');
const path = require('path');

console.log('DoPlan CLI: Validating package before publish...');

// Check required files
const requiredFiles = [
  'package.json',
  'bin/doplan.js',
  'README.md',
  'LICENSE'
];

let hasErrors = false;

for (const file of requiredFiles) {
  const filePath = path.join(__dirname, '..', file);
  if (!fs.existsSync(filePath)) {
    console.error(`Error: Required file missing: ${file}`);
    hasErrors = true;
  }
}

// Validate package.json
try {
  const packageJson = require('../package.json');
  
  if (!packageJson.name) {
    console.error('Error: package.json missing "name" field');
    hasErrors = true;
  }
  
  if (!packageJson.version) {
    console.error('Error: package.json missing "version" field');
    hasErrors = true;
  }
  
  if (!packageJson.bin || !packageJson.bin.doplan) {
    console.error('Error: package.json missing "bin.doplan" field');
    hasErrors = true;
  }
  
  // Validate version format (semver)
  const semverRegex = /^\d+\.\d+\.\d+(-.*)?$/;
  if (!semverRegex.test(packageJson.version)) {
    console.error(`Error: Invalid version format: ${packageJson.version}`);
    hasErrors = true;
  }
} catch (error) {
  console.error(`Error reading package.json: ${error.message}`);
  hasErrors = true;
}

if (hasErrors) {
  console.error('\nPre-publish validation failed. Please fix the errors above.');
  process.exit(1);
}

console.log('DoPlan CLI: Package validation passed!');
console.log('Ready to publish to npm.');

