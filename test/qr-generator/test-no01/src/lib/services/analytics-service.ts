import { createHash } from "crypto";
import { getDatabase } from "@/lib/db/database";
import { Analytics, ActivityPoint, GenerationParams } from "@/types/analytics";

/**
 * Analytics Service
 * Handles tracking and querying QR code generation statistics
 */
export class AnalyticsService {
  /**
   * Track a QR code generation event
   */
  trackGeneration(params: GenerationParams): void {
    const db = getDatabase();
    const stmt = db.prepare(`
      INSERT INTO generations (text_hash, size, format, error_correction, response_time_ms)
      VALUES (?, ?, ?, ?, ?)
    `);

    stmt.run(
      params.textHash,
      params.size,
      params.format,
      params.errorCorrection,
      params.responseTimeMs
    );
  }

  /**
   * Hash text for privacy (SHA-256)
   */
  hashText(text: string): string {
    return createHash("sha256").update(text).digest("hex");
  }

  /**
   * Get aggregated statistics
   */
  getStatistics(): Analytics {
    const db = getDatabase();

    // Get total generations
    const totalResult = db
      .prepare("SELECT COUNT(*) as count FROM generations")
      .get() as { count: number };
    const totalGenerations = totalResult.count;

    // Get today's generations
    const todayResult = db
      .prepare(
        "SELECT COUNT(*) as count FROM generations WHERE DATE(created_at) = DATE('now')"
      )
      .get() as { count: number };
    const todayGenerations = todayResult.count;

    // Get recent activity (last 24 hours, aggregated by 5-minute intervals)
    const recentActivity = this.getRecentActivity();

    // Calculate average response time
    const avgResponseTimeResult = db
      .prepare(
        "SELECT AVG(response_time_ms) as avg FROM generations WHERE response_time_ms IS NOT NULL"
      )
      .get() as { avg: number | null };
    const averageResponseTime = avgResponseTimeResult.avg
      ? Math.round(avgResponseTimeResult.avg)
      : 0;

    // Determine API status based on recent errors and response times
    const apiStatus = this.calculateAPIStatus(averageResponseTime);

    return {
      totalGenerations,
      todayGenerations,
      recentActivity,
      apiStatus,
      averageResponseTime,
    };
  }

  /**
   * Get recent activity aggregated by 5-minute intervals (last 24 hours)
   */
  private getRecentActivity(): ActivityPoint[] {
    const db = getDatabase();
    const stmt = db.prepare(`
      SELECT 
        strftime('%Y-%m-%d %H:%M', datetime(created_at, '-' || (strftime('%M', created_at) % 5) || ' minutes')) as timestamp,
        COUNT(*) as count
      FROM generations
      WHERE created_at >= datetime('now', '-24 hours')
      GROUP BY timestamp
      ORDER BY timestamp ASC
    `);

    const results = stmt.all() as Array<{ timestamp: string; count: number }>;
    return results.map((row) => ({
      timestamp: row.timestamp,
      count: row.count,
    }));
  }

  /**
   * Calculate API status based on metrics
   */
  private calculateAPIStatus(
    averageResponseTime: number
  ): "operational" | "degraded" | "down" {
    // Simple heuristic: if average response time is reasonable, operational
    // If high but not too high, degraded
    // If very high or no data, down
    if (averageResponseTime === 0) {
      return "down"; // No data
    }
    if (averageResponseTime < 200) {
      return "operational";
    }
    if (averageResponseTime < 1000) {
      return "degraded";
    }
    return "down";
  }
}

// Export singleton instance
export const analyticsService = new AnalyticsService();

