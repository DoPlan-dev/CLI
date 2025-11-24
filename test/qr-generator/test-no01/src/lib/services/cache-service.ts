import { createHash } from "crypto";
import { QRRequest } from "@/types/qr";
import { QRResponse } from "@/types/qr";

/**
 * Cached QR Code Entry
 */
interface CachedQR {
  response: QRResponse;
  expiresAt: number;
}

/**
 * Cache Service
 * Implements in-memory caching for QR code generation requests
 */
export class CacheService {
  private cache: Map<string, CachedQR> = new Map();
  private readonly defaultTTL: number = 3600000; // 1 hour in milliseconds

  /**
   * Generate cache key from request parameters
   */
  private generateCacheKey(request: QRRequest): string {
    const keyData = JSON.stringify({
      text: request.text,
      size: request.size || 200,
      format: request.format || "png",
      errorCorrection: request.errorCorrection || "M",
    });

    return createHash("sha256").update(keyData).digest("hex");
  }

  /**
   * Get cached QR code if available and not expired
   */
  get(request: QRRequest): QRResponse | null {
    const key = this.generateCacheKey(request);
    const cached = this.cache.get(key);

    if (!cached) {
      return null;
    }

    // Check if expired
    if (Date.now() > cached.expiresAt) {
      this.cache.delete(key);
      return null;
    }

    return cached.response;
  }

  /**
   * Store QR code in cache with TTL
   */
  set(request: QRRequest, response: QRResponse, ttl?: number): void {
    const key = this.generateCacheKey(request);
    const expiresAt = Date.now() + (ttl || this.defaultTTL);

    this.cache.set(key, {
      response,
      expiresAt,
    });
  }

  /**
   * Clear expired entries from cache
   */
  clearExpired(): void {
    const now = Date.now();
    for (const [key, value] of this.cache.entries()) {
      if (now > value.expiresAt) {
        this.cache.delete(key);
      }
    }
  }

  /**
   * Clear all cache entries
   */
  clear(): void {
    this.cache.clear();
  }

  /**
   * Get cache statistics
   */
  getStats(): { size: number; entries: number } {
    return {
      size: this.cache.size,
      entries: this.cache.size,
    };
  }
}

// Export singleton instance
export const cacheService = new CacheService();

// Periodically clean up expired entries (every 5 minutes)
if (typeof setInterval !== "undefined") {
  setInterval(() => {
    cacheService.clearExpired();
  }, 5 * 60 * 1000);
}

