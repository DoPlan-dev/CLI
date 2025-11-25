"use client";

import React, { useState, useEffect } from "react";
import { QRRequest, QRResponse } from "@/types/qr";
import { getFriendlyErrorMessage } from "@/lib/utils/client-errors";
import { ErrorCode } from "@/lib/utils/errors";

interface QRPreviewProps {
  text: string;
  size?: number;
  format?: "png" | "svg";
  errorCorrection?: "L" | "M" | "Q" | "H";
  className?: string;
  onQRGenerated?: (qrCode: string | null) => void;
}

export function QRPreview({
  text,
  size = 200,
  format = "png",
  errorCorrection = "M",
  className = "",
  onQRGenerated,
}: QRPreviewProps) {
  const [qrCode, setQrCode] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    // Debounce API calls (300ms)
    const timeoutId = setTimeout(() => {
      if (!text || text.trim().length === 0) {
        setQrCode(null);
        setError(null);
        return;
      }

      const generateQR = async () => {
        setLoading(true);
        setError(null);

        try {
          const request: QRRequest = {
            text: text.trim(),
            size,
            format,
            errorCorrection,
          };

          const response = await fetch("/api/qr", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(request),
          });

          if (!response.ok) {
            let errorCode: ErrorCode | undefined;
            let errorMessage: string | undefined;
            try {
              const errorData = await response.json();
              errorCode = errorData.error?.code;
              errorMessage = errorData.error?.message;
            } catch {
              // ignore parse failures
            }
            throw new Error(
              getFriendlyErrorMessage(
                { code: errorCode, message: errorMessage },
                "We couldn't generate a QR code right now. Please try again."
              )
            );
          }

          const data: QRResponse = await response.json();
          setQrCode(data.qrCode);
          onQRGenerated?.(data.qrCode);
        } catch (err) {
          const message =
            err instanceof Error
              ? err.message
              : "We couldn't generate a QR code right now. Please try again.";
          setError(message);
          setQrCode(null);
          onQRGenerated?.(null);
        } finally {
          setLoading(false);
        }
      };

      generateQR();
    }, 300);

    return () => clearTimeout(timeoutId);
  }, [text, size, format, errorCorrection]);

  return (
    <div
      className={`w-full max-w-md aspect-square flex items-center justify-center bg-bg-primary border-2 rounded-lg p-6 transition-all duration-300 ${
        qrCode
          ? "border-primary shadow-md"
          : "border-dashed border-border"
      } ${className}`}
      role="img"
      aria-label="QR Code Preview"
      aria-live="polite"
      aria-atomic="true"
    >
      {loading && (
        <div className="text-text-secondary text-sm">Generating...</div>
      )}
      {error && (
        <div className="text-error text-sm text-center">{error}</div>
      )}
      {!loading && !error && !qrCode && (
        <div className="text-text-tertiary text-sm text-center">
          Enter text above to generate QR code
        </div>
      )}
      {!loading && !error && qrCode && (
        <img
          src={qrCode}
          alt="Generated QR Code"
          className="w-full h-full object-contain fade-in"
        />
      )}
    </div>
  );
}

