import { NextResponse } from "next/server";
import { analyticsService } from "@/lib/services/analytics-service";
import { withRateLimit } from "@/lib/middleware/rate-limit";

// Cache analytics response for 1 minute
let cachedAnalytics: { data: ReturnType<typeof analyticsService.getStatistics>; expiresAt: number } | null = null;
const CACHE_TTL = 60000; // 1 minute

/**
 * GET /api/analytics
 * Get usage statistics
 */
async function handleGET() {
  try {
    // Check cache
    if (cachedAnalytics && Date.now() < cachedAnalytics.expiresAt) {
      return NextResponse.json(cachedAnalytics.data);
    }

    // Get fresh analytics
    const analytics = analyticsService.getStatistics();

    // Cache the result
    cachedAnalytics = {
      data: analytics,
      expiresAt: Date.now() + CACHE_TTL,
    };

    return NextResponse.json(analytics);
  } catch (error) {
    console.error("Analytics error:", error);
    return NextResponse.json(
      {
        success: false,
        error: {
          code: "INTERNAL_ERROR",
          message: "Failed to retrieve analytics",
        },
      },
      { status: 500 }
    );
  }
}

// Export with rate limiting
export const GET = withRateLimit(handleGET);

