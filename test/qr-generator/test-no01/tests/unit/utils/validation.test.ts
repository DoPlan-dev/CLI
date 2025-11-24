import { describe, it, expect } from "vitest";
import {
  validateText,
  validateSize,
  validateFormat,
  validateErrorCorrection,
  validateQRRequest,
} from "@/lib/utils/validation";

describe("Validation Utilities", () => {
  describe("validateText", () => {
    it("should validate valid text", () => {
      const result = validateText("Hello World");
      expect(result.valid).toBe(true);
    });

    it("should reject empty text", () => {
      const result = validateText("");
      expect(result.valid).toBe(false);
      expect(result.error).toBeDefined();
      // Empty string triggers the "required" check first
      expect(result.error).toMatch(/required|at least 1 character/);
    });

    it("should reject text exceeding 2000 characters", () => {
      const longText = "a".repeat(2001);
      const result = validateText(longText);
      expect(result.valid).toBe(false);
      expect(result.error).toContain("2000 characters");
    });

    it("should reject non-string input", () => {
      const result = validateText(null as unknown as string);
      expect(result.valid).toBe(false);
    });

    it("should accept text at boundaries", () => {
      expect(validateText("a").valid).toBe(true);
      expect(validateText("a".repeat(2000)).valid).toBe(true);
    });
  });

  describe("validateSize", () => {
    it("should validate valid size", () => {
      expect(validateSize(200).valid).toBe(true);
      expect(validateSize(50).valid).toBe(true);
      expect(validateSize(2000).valid).toBe(true);
    });

    it("should accept undefined (optional)", () => {
      expect(validateSize(undefined).valid).toBe(true);
    });

    it("should reject size below minimum", () => {
      const result = validateSize(49);
      expect(result.valid).toBe(false);
      expect(result.error).toContain("50 and 2000");
    });

    it("should reject size above maximum", () => {
      const result = validateSize(2001);
      expect(result.valid).toBe(false);
      expect(result.error).toContain("50 and 2000");
    });

    it("should reject invalid number", () => {
      const result = validateSize(NaN);
      expect(result.valid).toBe(false);
    });
  });

  describe("validateFormat", () => {
    it("should validate PNG format", () => {
      expect(validateFormat("png").valid).toBe(true);
    });

    it("should validate SVG format", () => {
      expect(validateFormat("svg").valid).toBe(true);
    });

    it("should accept undefined (optional)", () => {
      expect(validateFormat(undefined).valid).toBe(true);
    });

    it("should reject invalid format", () => {
      const result = validateFormat("jpg" as "png" | "svg");
      expect(result.valid).toBe(false);
      expect(result.error).toBeDefined();
      // Error message contains format information
      expect(result.error).toMatch(/png|svg/);
    });
  });

  describe("validateErrorCorrection", () => {
    it("should validate all error correction levels", () => {
      expect(validateErrorCorrection("L").valid).toBe(true);
      expect(validateErrorCorrection("M").valid).toBe(true);
      expect(validateErrorCorrection("Q").valid).toBe(true);
      expect(validateErrorCorrection("H").valid).toBe(true);
    });

    it("should accept undefined (optional)", () => {
      expect(validateErrorCorrection(undefined).valid).toBe(true);
    });

    it("should reject invalid error correction level", () => {
      const result = validateErrorCorrection("X" as "L" | "M" | "Q" | "H");
      expect(result.valid).toBe(false);
      expect(result.error).toContain("L, M, Q, H");
    });
  });

  describe("validateQRRequest", () => {
    it("should validate complete valid request", () => {
      const request = {
        text: "Hello",
        size: 200,
        format: "png",
        errorCorrection: "M",
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(true);
    });

    it("should validate request with only required field", () => {
      const request = {
        text: "Hello",
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(true);
    });

    it("should reject request with invalid text", () => {
      const request = {
        text: "",
        size: 200,
        format: "png",
        errorCorrection: "M",
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(false);
      expect(result.field).toBe("text");
    });

    it("should reject request with invalid size", () => {
      const request = {
        text: "Hello",
        size: 30,
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(false);
      expect(result.field).toBe("size");
    });

    it("should reject request with invalid format", () => {
      const request = {
        text: "Hello",
        format: "jpg",
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(false);
      expect(result.field).toBe("format");
    });

    it("should reject request with invalid error correction", () => {
      const request = {
        text: "Hello",
        errorCorrection: "X",
      };

      const result = validateQRRequest(request);
      expect(result.valid).toBe(false);
      expect(result.field).toBe("errorCorrection");
    });
  });
});

