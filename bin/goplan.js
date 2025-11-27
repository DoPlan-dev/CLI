#!/usr/bin/env node

/**
 * GoPlan helper CLI
 * Provides beginner-friendly patches (access, etc.) without Go tooling.
 */

const { spawnSync } = require('child_process');
const fs = require('fs');
const path = require('path');

const ACCESS_SCRIPT = path.join(__dirname, '..', 'scripts', 'doplan', 'access.sh');

const args = process.argv.slice(2);
const command = args[0];

function printHelp() {
  console.log(`GoPlan Helper
Usage:
  goplan access [all|.do/system|docs|.do/plan]

Examples:
  goplan access all
  goplan access docs`);
}

function ensureAccessScript() {
  if (!fs.existsSync(ACCESS_SCRIPT)) {
    console.error('Access patch script is missing. Expected at:', ACCESS_SCRIPT);
    process.exit(1);
  }
}

function runAccess(target = 'all') {
  ensureAccessScript();

  const result = spawnSync('bash', [ACCESS_SCRIPT, target], {
    stdio: 'inherit',
    cwd: process.cwd(),
  });

  if (result.error) {
    console.error('Failed to run access patch:', result.error.message);
    process.exit(result.status ?? 1);
  }

  process.exit(result.status ?? 0);
}

function main() {
  if (!command || command === '--help' || command === '-h') {
    printHelp();
    process.exit(0);
  }

  if (command === 'access') {
    runAccess(args[1] || 'all');
    return;
  }

  console.error(`Unknown command: ${command}\n`);
  printHelp();
  process.exit(1);
}

main();

