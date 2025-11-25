import { beforeAll, afterAll } from "vitest";
import { closeDatabase } from "@/lib/db/database";

// Clean up after all tests
afterAll(async () => {
  closeDatabase();
});

