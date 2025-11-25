import QRCode from "qrcode";
import { QRRequest, QRResponse } from "@/types/qr";
import { validateQRRequest } from "@/lib/utils/validation";

/**
 * QR Service
 * Handles QR code generation with various options
 */
export class QRService {
  /**
   * Generate QR code from text with options
   */
  async generateQR(request: QRRequest): Promise<QRResponse> {
    // Validate input
    const validation = validateQRRequest(request);
    if (!validation.valid) {
      throw new Error(validation.error || "Invalid input");
    }

    // Set defaults
    const size = request.size || 200;
    const format = request.format || "png";
    const errorCorrection = request.errorCorrection || "M";

    // Map error correction level to qrcode library format
    const errorCorrectionLevel = errorCorrection as "L" | "M" | "Q" | "H";

    try {
      let qrCode: string;

      if (format === "svg") {
        // Generate SVG
        qrCode = await QRCode.toString(request.text, {
          type: "svg",
          width: size,
          errorCorrectionLevel,
          margin: 1,
        });
        // Convert SVG to base64
        qrCode = `data:image/svg+xml;base64,${Buffer.from(qrCode).toString(
          "base64"
        )}`;
      } else {
        // Generate PNG (default)
        const buffer = await QRCode.toBuffer(request.text, {
          type: "png",
          width: size,
          errorCorrectionLevel,
          margin: 1,
        });
        // Convert to base64
        qrCode = `data:image/png;base64,${buffer.toString("base64")}`;
      }

      return {
        success: true,
        qrCode,
        format,
        size,
        errorCorrection,
        generatedAt: new Date().toISOString(),
      };
    } catch (error) {
      throw new Error(
        `Failed to generate QR code: ${error instanceof Error ? error.message : "Unknown error"}`
      );
    }
  }

  /**
   * Generate QR code buffer (for file downloads)
   */
  async generateQRBuffer(
    request: QRRequest
  ): Promise<{ buffer: Buffer; mimeType: string; filename: string }> {
    // Validate input
    const validation = validateQRRequest(request);
    if (!validation.valid) {
      throw new Error(validation.error || "Invalid input");
    }

    // Set defaults
    const size = request.size || 200;
    const format = request.format || "png";
    const errorCorrection = request.errorCorrection || "M";
    const errorCorrectionLevel = errorCorrection as "L" | "M" | "Q" | "H";

    try {
      let buffer: Buffer;
      let mimeType: string;
      let filename: string;

      if (format === "svg") {
        const svgString = await QRCode.toString(request.text, {
          type: "svg",
          width: size,
          errorCorrectionLevel,
          margin: 1,
        });
        buffer = Buffer.from(svgString);
        mimeType = "image/svg+xml";
        filename = `qr-code-${Date.now()}.svg`;
      } else {
        buffer = await QRCode.toBuffer(request.text, {
          type: "png",
          width: size,
          errorCorrectionLevel,
          margin: 1,
        });
        mimeType = "image/png";
        filename = `qr-code-${Date.now()}.png`;
      }

      return { buffer, mimeType, filename };
    } catch (error) {
      throw new Error(
        `Failed to generate QR code: ${error instanceof Error ? error.message : "Unknown error"}`
      );
    }
  }
}

// Export singleton instance
export const qrService = new QRService();

