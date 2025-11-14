#!/usr/bin/env node

/**
 * Post-install script for npm package
 * Downloads the platform-specific binary after npm install
 */

const { downloadBinary } = require('../bin/doplan.js');

console.log('DoPlan CLI: Setting up platform-specific binary...');

try {
  downloadBinary();
  console.log('DoPlan CLI: Installation complete!');
} catch (error) {
  console.error('DoPlan CLI: Installation warning:', error.message);
  console.error('You may need to download the binary manually from GitHub releases.');
  // Don't exit with error - allow package to be installed even if download fails
  // User can manually download or use the wrapper script which will retry
}

