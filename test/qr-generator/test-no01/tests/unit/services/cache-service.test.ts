import { describe, it, expect, beforeEach } from "vitest";
import { cacheService } from "@/lib/services/cache-service";
import { QRRequest, QRResponse } from "@/types/qr";

describe("CacheService", () => {
  beforeEach(() => {
    cacheService.clear();
  });

  describe("get and set", () => {
    it("should store and retrieve cached QR code", () => {
      const request: QRRequest = {
        text: "https://example.com",
        size: 200,
        format: "png",
        errorCorrection: "M",
      };

      const response: QRResponse = {
        success: true,
        qrCode: "data:image/png;base64,test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      cacheService.set(request, response);
      const cached = cacheService.get(request);

      expect(cached).not.toBeNull();
      expect(cached?.qrCode).toBe(response.qrCode);
    });

    it("should return null for non-existent cache entry", () => {
      const request: QRRequest = {
        text: "not cached",
      };

      const cached = cacheService.get(request);
      expect(cached).toBeNull();
    });

    it("should generate same cache key for identical requests", () => {
      const request1: QRRequest = {
        text: "test",
        size: 200,
        format: "png",
        errorCorrection: "M",
      };

      const request2: QRRequest = {
        text: "test",
        size: 200,
        format: "png",
        errorCorrection: "M",
      };

      const response: QRResponse = {
        success: true,
        qrCode: "test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      cacheService.set(request1, response);
      const cached = cacheService.get(request2);

      expect(cached).not.toBeNull();
    });

    it("should generate different cache keys for different requests", () => {
      const request1: QRRequest = {
        text: "test1",
        size: 200,
      };

      const request2: QRRequest = {
        text: "test2",
        size: 200,
      };

      const response: QRResponse = {
        success: true,
        qrCode: "test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      cacheService.set(request1, response);
      const cached = cacheService.get(request2);

      expect(cached).toBeNull();
    });
  });

  describe("TTL expiration", () => {
    it("should expire cache entries after TTL", async () => {
      const request: QRRequest = {
        text: "test",
      };

      const response: QRResponse = {
        success: true,
        qrCode: "test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      // Set with short TTL (100ms)
      cacheService.set(request, response, 100);

      // Should be available immediately
      expect(cacheService.get(request)).not.toBeNull();

      // Wait for expiration
      await new Promise((resolve) => setTimeout(resolve, 150));

      // Should be expired
      expect(cacheService.get(request)).toBeNull();
    });
  });

  describe("clearExpired", () => {
    it("should remove expired entries", async () => {
      const request1: QRRequest = { text: "test1" };
      const request2: QRRequest = { text: "test2" };

      const response: QRResponse = {
        success: true,
        qrCode: "test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      // Set one with short TTL, one with long TTL
      cacheService.set(request1, response, 50);
      cacheService.set(request2, response, 10000);

      await new Promise((resolve) => setTimeout(resolve, 100));

      cacheService.clearExpired();

      expect(cacheService.get(request1)).toBeNull();
      expect(cacheService.get(request2)).not.toBeNull();
    });
  });

  describe("clear", () => {
    it("should clear all cache entries", () => {
      const request: QRRequest = { text: "test" };
      const response: QRResponse = {
        success: true,
        qrCode: "test",
        format: "png",
        size: 200,
        errorCorrection: "M",
        generatedAt: new Date().toISOString(),
      };

      cacheService.set(request, response);
      expect(cacheService.get(request)).not.toBeNull();

      cacheService.clear();
      expect(cacheService.get(request)).toBeNull();
    });
  });
});

