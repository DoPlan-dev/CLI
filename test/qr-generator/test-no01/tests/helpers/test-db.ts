import Database from "better-sqlite3";
import path from "path";
import fs from "fs";

/**
 * Create a test database instance
 */
export function createTestDatabase(): Database.Database {
  const testDbPath = path.join(process.cwd(), "tests", "test.db");
  const testDbDir = path.dirname(testDbPath);

  // Ensure test directory exists
  if (!fs.existsSync(testDbDir)) {
    fs.mkdirSync(testDbDir, { recursive: true });
  }

  // Remove existing test database
  if (fs.existsSync(testDbPath)) {
    fs.unlinkSync(testDbPath);
  }

  const db = new Database(testDbPath);

  // Initialize schema
  db.exec(`
    CREATE TABLE generations (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      text_hash TEXT NOT NULL,
      size INTEGER NOT NULL,
      format TEXT NOT NULL,
      error_correction TEXT NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      response_time_ms INTEGER
    );

    CREATE INDEX idx_created_at ON generations(created_at);
    CREATE INDEX idx_text_hash ON generations(text_hash);
  `);

  return db;
}

/**
 * Clean up test database
 */
export function cleanupTestDatabase(db: Database.Database): void {
  db.close();
  const testDbPath = path.join(process.cwd(), "tests", "test.db");
  if (fs.existsSync(testDbPath)) {
    fs.unlinkSync(testDbPath);
  }
}

