/**
 * Validation utilities for QR code generation requests
 */

export interface ValidationResult {
  valid: boolean;
  error?: string;
  field?: string;
}

/**
 * Validate text length (1-2000 characters)
 */
export function validateText(text: string): ValidationResult {
  if (!text || typeof text !== "string") {
    return {
      valid: false,
      error: "Text is required and must be a string",
      field: "text",
    };
  }

  if (text.length < 1) {
    return {
      valid: false,
      error: "Text must be at least 1 character long",
      field: "text",
    };
  }

  if (text.length > 2000) {
    return {
      valid: false,
      error: "Text must not exceed 2000 characters",
      field: "text",
    };
  }

  return { valid: true };
}

/**
 * Validate size range (50-2000px)
 */
export function validateSize(size: number | undefined): ValidationResult {
  if (size === undefined) {
    return { valid: true }; // Optional, will use default
  }

  if (typeof size !== "number" || isNaN(size)) {
    return {
      valid: false,
      error: "Size must be a valid number",
      field: "size",
    };
  }

  if (size < 50 || size > 2000) {
    return {
      valid: false,
      error: "Size must be between 50 and 2000 pixels",
      field: "size",
    };
  }

  return { valid: true };
}

/**
 * Validate format (png/svg)
 */
export function validateFormat(
  format: string | undefined
): ValidationResult {
  if (format === undefined) {
    return { valid: true }; // Optional, will use default
  }

  if (format !== "png" && format !== "svg") {
    return {
      valid: false,
      error: "Format must be either 'png' or 'svg'",
      field: "format",
    };
  }

  return { valid: true };
}

/**
 * Validate error correction level (L/M/Q/H)
 */
export function validateErrorCorrection(
  errorCorrection: string | undefined
): ValidationResult {
  if (errorCorrection === undefined) {
    return { valid: true }; // Optional, will use default
  }

  const validLevels = ["L", "M", "Q", "H"];
  if (!validLevels.includes(errorCorrection)) {
    return {
      valid: false,
      error: "Error correction must be one of: L, M, Q, H",
      field: "errorCorrection",
    };
  }

  return { valid: true };
}

/**
 * Validate complete QR request
 */
export function validateQRRequest(request: {
  text?: unknown;
  size?: unknown;
  format?: unknown;
  errorCorrection?: unknown;
}): ValidationResult {
  // Validate text
  const textValidation = validateText(request.text as string);
  if (!textValidation.valid) {
    return textValidation;
  }

  // Validate size
  const sizeValidation = validateSize(request.size as number | undefined);
  if (!sizeValidation.valid) {
    return sizeValidation;
  }

  // Validate format
  const formatValidation = validateFormat(request.format as string | undefined);
  if (!formatValidation.valid) {
    return formatValidation;
  }

  // Validate error correction
  const errorCorrectionValidation = validateErrorCorrection(
    request.errorCorrection as string | undefined
  );
  if (!errorCorrectionValidation.valid) {
    return errorCorrectionValidation;
  }

  return { valid: true };
}

