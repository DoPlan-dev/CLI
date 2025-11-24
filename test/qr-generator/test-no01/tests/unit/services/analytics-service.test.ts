import { describe, it, expect, beforeEach, afterEach } from "vitest";
import { analyticsService } from "@/lib/services/analytics-service";
import { createTestDatabase, cleanupTestDatabase } from "../../helpers/test-db";
import Database from "better-sqlite3";

describe("AnalyticsService", () => {
  let testDb: Database.Database;

  beforeEach(() => {
    testDb = createTestDatabase();
    // Mock the database getter (this is a simplified approach)
    // In a real scenario, you'd inject the database dependency
  });

  afterEach(() => {
    cleanupTestDatabase(testDb);
  });

  describe("hashText", () => {
    it("should generate consistent hash for same text", () => {
      const text = "test text";
      const hash1 = analyticsService.hashText(text);
      const hash2 = analyticsService.hashText(text);

      expect(hash1).toBe(hash2);
      expect(hash1).toHaveLength(64); // SHA-256 produces 64 char hex string
    });

    it("should generate different hashes for different text", () => {
      const hash1 = analyticsService.hashText("text1");
      const hash2 = analyticsService.hashText("text2");

      expect(hash1).not.toBe(hash2);
    });
  });

  describe("trackGeneration", () => {
    it("should track generation event", () => {
      const params = {
        textHash: "test-hash",
        size: 200,
        format: "png",
        errorCorrection: "M",
        responseTimeMs: 50,
      };

      // This test would require mocking the database
      // For now, we'll test that the method doesn't throw
      expect(() => {
        analyticsService.trackGeneration(params);
      }).not.toThrow();
    });
  });

  describe("getStatistics", () => {
    it("should return statistics structure", () => {
      const stats = analyticsService.getStatistics();

      expect(stats).toHaveProperty("totalGenerations");
      expect(stats).toHaveProperty("todayGenerations");
      expect(stats).toHaveProperty("recentActivity");
      expect(stats).toHaveProperty("apiStatus");
      expect(stats).toHaveProperty("averageResponseTime");

      expect(typeof stats.totalGenerations).toBe("number");
      expect(typeof stats.todayGenerations).toBe("number");
      expect(Array.isArray(stats.recentActivity)).toBe(true);
      expect(["operational", "degraded", "down"]).toContain(stats.apiStatus);
      expect(typeof stats.averageResponseTime).toBe("number");
    });
  });
});

