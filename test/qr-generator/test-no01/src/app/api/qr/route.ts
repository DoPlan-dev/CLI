import { NextRequest, NextResponse } from "next/server";
import { QRRequest } from "@/types/qr";
import { qrService } from "@/lib/services/qr-service";
import { cacheService } from "@/lib/services/cache-service";
import { analyticsService } from "@/lib/services/analytics-service";
import {
  createErrorResponseWithStatus,
  ErrorCode,
} from "@/lib/utils/errors";
import { validateQRRequest } from "@/lib/utils/validation";
import { withRateLimit } from "@/lib/middleware/rate-limit";

/**
 * POST /api/qr
 * Generate QR code from text
 */
async function handlePOST(request: NextRequest) {
  const startTime = Date.now();

  try {
    // Parse request body
    const body = await request.json().catch(() => ({}));
    const qrRequest: QRRequest = {
      text: body.text,
      size: body.size,
      format: body.format,
      errorCorrection: body.errorCorrection,
    };

    // Validate request
    const validation = validateQRRequest(qrRequest);
    if (!validation.valid) {
      const errorResponse = createErrorResponseWithStatus(
        ErrorCode.INVALID_INPUT,
        validation.error || "Invalid request",
        validation.field
      );
      return NextResponse.json(errorResponse.body, {
        status: errorResponse.status,
      });
    }

    // Check cache first
    const cached = cacheService.get(qrRequest);
    if (cached) {
      return NextResponse.json(cached);
    }

    // Generate QR code
    const response = await qrService.generateQR(qrRequest);
    const responseTime = Date.now() - startTime;

    // Store in cache
    cacheService.set(qrRequest, response);

    // Track analytics
    const textHash = analyticsService.hashText(qrRequest.text);
    analyticsService.trackGeneration({
      textHash,
      size: response.size,
      format: response.format,
      errorCorrection: response.errorCorrection,
      responseTimeMs: responseTime,
    });

    // Check Accept header for file download
    const acceptHeader = request.headers.get("accept") || "";
    if (acceptHeader.includes("application/octet-stream") || acceptHeader.includes("image/")) {
      // Return file download
      const { buffer, mimeType, filename } = await qrService.generateQRBuffer(
        qrRequest
      );

      return new NextResponse(buffer as unknown as BodyInit, {
        headers: {
          "Content-Type": mimeType,
          "Content-Disposition": `attachment; filename="${filename}"`,
        },
      });
    }

    // Return JSON response
    return NextResponse.json(response);
  } catch (error) {
    console.error("QR generation error:", error);
    const errorResponse = createErrorResponseWithStatus(
      ErrorCode.INTERNAL_ERROR,
      error instanceof Error ? error.message : "Internal server error"
    );
    return NextResponse.json(errorResponse.body, {
      status: errorResponse.status,
    });
  }
}

// Export with rate limiting
export const POST = withRateLimit(handlePOST);

