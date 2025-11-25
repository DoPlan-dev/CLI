/**
 * Error handling utilities
 */

/**
 * Error codes for API responses
 */
export enum ErrorCode {
  INVALID_INPUT = "INVALID_INPUT",
  TEXT_TOO_LONG = "TEXT_TOO_LONG",
  INVALID_SIZE = "INVALID_SIZE",
  RATE_LIMIT_EXCEEDED = "RATE_LIMIT_EXCEEDED",
  INTERNAL_ERROR = "INTERNAL_ERROR",
}

/**
 * API Error Response Interface
 */
export interface APIError {
  success: false;
  error: {
    code: ErrorCode;
    message: string;
    field?: string;
  };
}

/**
 * Create error response object
 */
export function createErrorResponse(
  code: ErrorCode,
  message: string,
  field?: string
): APIError {
  return {
    success: false,
    error: {
      code,
      message,
      ...(field && { field }),
    },
  };
}

/**
 * Map error code to HTTP status code
 */
export function getHttpStatus(errorCode: ErrorCode): number {
  const statusMap: Record<ErrorCode, number> = {
    [ErrorCode.INVALID_INPUT]: 400,
    [ErrorCode.TEXT_TOO_LONG]: 400,
    [ErrorCode.INVALID_SIZE]: 400,
    [ErrorCode.RATE_LIMIT_EXCEEDED]: 429,
    [ErrorCode.INTERNAL_ERROR]: 500,
  };

  return statusMap[errorCode] || 500;
}

/**
 * Create standardized error response for API routes
 */
export function createErrorResponseWithStatus(
  code: ErrorCode,
  message: string,
  field?: string
): { status: number; body: APIError } {
  return {
    status: getHttpStatus(code),
    body: createErrorResponse(code, message, field),
  };
}

