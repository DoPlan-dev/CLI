import { NextRequest, NextResponse } from "next/server";

/**
 * Rate limit configuration
 */
const RATE_LIMITS = {
  perMinute: 100,
  perHour: 1000,
};

/**
 * Rate limit store (in-memory)
 * In production, consider using Redis or similar
 */
interface RateLimitEntry {
  count: number;
  resetAt: number;
}

const rateLimitStore = new Map<string, {
  minute: RateLimitEntry;
  hour: RateLimitEntry;
}>();

/**
 * Get client IP address from request
 */
function getClientIP(request: NextRequest): string {
  // Check various headers for IP (for proxy/load balancer scenarios)
  const forwarded = request.headers.get("x-forwarded-for");
  const realIP = request.headers.get("x-real-ip");
  const cfConnectingIP = request.headers.get("cf-connecting-ip");

  if (cfConnectingIP) {
    return cfConnectingIP;
  }
  if (realIP) {
    return realIP;
  }
  if (forwarded) {
    return forwarded.split(",")[0].trim();
  }

  // Fallback - NextRequest doesn't have ip property directly
  // In production, this should be handled by the platform (Vercel, etc.)
  return "unknown";
}

/**
 * Check rate limit for a client
 */
export function checkRateLimit(
  request: NextRequest
): { allowed: boolean; headers: Record<string, string> } {
  const ip = getClientIP(request);
  const now = Date.now();

  // Get or create rate limit entry for this IP
  let entry = rateLimitStore.get(ip);
  if (!entry) {
    entry = {
      minute: { count: 0, resetAt: now + 60000 },
      hour: { count: 0, resetAt: now + 3600000 },
    };
    rateLimitStore.set(ip, entry);
  }

  // Reset counters if time window has passed
  if (now >= entry.minute.resetAt) {
    entry.minute = { count: 0, resetAt: now + 60000 };
  }
  if (now >= entry.hour.resetAt) {
    entry.hour = { count: 0, resetAt: now + 3600000 };
  }

  // Increment counters
  entry.minute.count++;
  entry.hour.count++;

  // Check limits
  const minuteLimitExceeded = entry.minute.count > RATE_LIMITS.perMinute;
  const hourLimitExceeded = entry.hour.count > RATE_LIMITS.perHour;

  // Prepare headers
  const headers: Record<string, string> = {
    "X-RateLimit-Limit-Minute": String(RATE_LIMITS.perMinute),
    "X-RateLimit-Remaining-Minute": String(
      Math.max(0, RATE_LIMITS.perMinute - entry.minute.count)
    ),
    "X-RateLimit-Reset-Minute": String(Math.ceil(entry.minute.resetAt / 1000)),
    "X-RateLimit-Limit-Hour": String(RATE_LIMITS.perHour),
    "X-RateLimit-Remaining-Hour": String(
      Math.max(0, RATE_LIMITS.perHour - entry.hour.count)
    ),
    "X-RateLimit-Reset-Hour": String(Math.ceil(entry.hour.resetAt / 1000)),
  };

  if (minuteLimitExceeded || hourLimitExceeded) {
    return {
      allowed: false,
      headers,
    };
  }

  return {
    allowed: true,
    headers,
  };
}

/**
 * Rate limit middleware wrapper
 */
export function withRateLimit(
  handler: (request: NextRequest) => Promise<NextResponse>
) {
  return async (request: NextRequest): Promise<NextResponse> => {
    const rateLimit = checkRateLimit(request);

    // Add rate limit headers to response
    const response = rateLimit.allowed
      ? await handler(request)
      : NextResponse.json(
          {
            success: false,
            error: {
              code: "RATE_LIMIT_EXCEEDED",
              message: "Too many requests. Please try again later.",
            },
          },
          { status: 429 }
        );

    // Add rate limit headers
    Object.entries(rateLimit.headers).forEach(([key, value]) => {
      response.headers.set(key, value);
    });

    return response;
  };
}

// Clean up old entries periodically (every 5 minutes)
if (typeof setInterval !== "undefined") {
  setInterval(() => {
    const now = Date.now();
    for (const [ip, entry] of rateLimitStore.entries()) {
      // Remove entries that are older than 1 hour
      if (now > entry.hour.resetAt + 3600000) {
        rateLimitStore.delete(ip);
      }
    }
  }, 5 * 60 * 1000);
}

