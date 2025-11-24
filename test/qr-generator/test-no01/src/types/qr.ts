/**
 * QR Code Request Interface
 */
export interface QRRequest {
  text: string;
  size?: number; // 50-2000, default 200
  format?: "png" | "svg"; // default 'png'
  errorCorrection?: "L" | "M" | "Q" | "H"; // default 'M'
}

/**
 * QR Code Response Interface
 */
export interface QRResponse {
  success: boolean;
  qrCode: string; // base64 encoded
  format: string;
  size: number;
  errorCorrection: string;
  generatedAt: string;
}

