#!/usr/bin/env node
/**
 * DoPlan CLI wrapper - routes to doplan or goplan based on command
 * This allows npx @doplan-dev/cli to work with multiple commands
 */

const path = require('path');
const { spawn } = require('child_process');
const fs = require('fs');

// Get the command name from process.argv
// npx @doplan-dev/cli doplan -> argv[2] = 'doplan'
// npx @doplan-dev/cli goplan -> argv[2] = 'goplan'
// npx @doplan-dev/cli -> argv[2] = undefined (default to doplan)
const command = process.argv[2];

let scriptPath;
if (command === 'goplan') {
  scriptPath = path.join(__dirname, 'goplan.js');
} else {
  // Default to doplan for any other command or no command
  scriptPath = path.join(__dirname, 'doplan.js');
}

if (!fs.existsSync(scriptPath)) {
  console.error(`Error: Script not found: ${scriptPath}`);
  process.exit(1);
}

// Execute the script with remaining arguments
const child = spawn('node', [scriptPath, ...process.argv.slice(3)], {
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

