import { describe, it, expect, beforeEach } from "vitest";
import { qrService } from "@/lib/services/qr-service";
import { QRRequest } from "@/types/qr";

describe("QRService", () => {
  describe("generateQR", () => {
    it("should generate PNG QR code with default options", async () => {
      const request: QRRequest = {
        text: "https://example.com",
      };

      const result = await qrService.generateQR(request);

      expect(result.success).toBe(true);
      expect(result.format).toBe("png");
      expect(result.size).toBe(200);
      expect(result.errorCorrection).toBe("M");
      expect(result.qrCode).toContain("data:image/png;base64,");
      expect(result.generatedAt).toBeDefined();
    });

    it("should generate SVG QR code", async () => {
      const request: QRRequest = {
        text: "Hello World",
        format: "svg",
      };

      const result = await qrService.generateQR(request);

      expect(result.success).toBe(true);
      expect(result.format).toBe("svg");
      expect(result.qrCode).toContain("data:image/svg+xml;base64,");
    });

    it("should respect custom size", async () => {
      const request: QRRequest = {
        text: "Test",
        size: 500,
      };

      const result = await qrService.generateQR(request);

      expect(result.success).toBe(true);
      expect(result.size).toBe(500);
    });

    it("should support all error correction levels", async () => {
      const levels: Array<"L" | "M" | "Q" | "H"> = ["L", "M", "Q", "H"];

      for (const level of levels) {
        const request: QRRequest = {
          text: "Test",
          errorCorrection: level,
        };

        const result = await qrService.generateQR(request);

        expect(result.success).toBe(true);
        expect(result.errorCorrection).toBe(level);
      }
    });

    it("should throw error for empty text", async () => {
      const request: QRRequest = {
        text: "",
      };

      await expect(qrService.generateQR(request)).rejects.toThrow();
    });

    it("should throw error for text exceeding 2000 characters", async () => {
      const request: QRRequest = {
        text: "a".repeat(2001),
      };

      await expect(qrService.generateQR(request)).rejects.toThrow();
    });

    it("should throw error for invalid size", async () => {
      const request: QRRequest = {
        text: "Test",
        size: 30, // Below minimum
      };

      await expect(qrService.generateQR(request)).rejects.toThrow();
    });
  });

  describe("generateQRBuffer", () => {
    it("should generate PNG buffer", async () => {
      const request: QRRequest = {
        text: "Test",
        format: "png",
      };

      const result = await qrService.generateQRBuffer(request);

      expect(result.buffer).toBeInstanceOf(Buffer);
      expect(result.mimeType).toBe("image/png");
      expect(result.filename).toContain(".png");
    });

    it("should generate SVG buffer", async () => {
      const request: QRRequest = {
        text: "Test",
        format: "svg",
      };

      const result = await qrService.generateQRBuffer(request);

      expect(result.buffer).toBeInstanceOf(Buffer);
      expect(result.mimeType).toBe("image/svg+xml");
      expect(result.filename).toContain(".svg");
    });
  });
});

