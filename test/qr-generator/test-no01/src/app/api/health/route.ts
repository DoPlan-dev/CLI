import { NextResponse } from "next/server";
import { getDatabase } from "@/lib/db/database";

/**
 * GET /api/health
 * Health check endpoint
 */
export async function GET() {
  try {
    // Check database connection
    const db = getDatabase();
    const healthCheck = db.prepare("SELECT 1 as health").get() as {
      health: number;
    };

    if (!healthCheck || healthCheck.health !== 1) {
      return NextResponse.json(
        {
          status: "down",
          database: "unhealthy",
          timestamp: new Date().toISOString(),
        },
        { status: 503 }
      );
    }

    // All checks passed
    return NextResponse.json({
      status: "operational",
      database: "healthy",
      timestamp: new Date().toISOString(),
    });
  } catch (error) {
    console.error("Health check error:", error);
    return NextResponse.json(
      {
        status: "down",
        database: "unhealthy",
        error: error instanceof Error ? error.message : "Unknown error",
        timestamp: new Date().toISOString(),
      },
      { status: 503 }
    );
  }
}

