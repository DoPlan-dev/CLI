import { ErrorCode } from "@/lib/utils/errors";

const DEFAULT_ERROR_MESSAGE =
  "Something went wrong while generating your QR code. Please try again.";

const errorMessages: Record<ErrorCode, string> = {
  [ErrorCode.INVALID_INPUT]:
    "We couldn't read your request. Double-check the fields and try again.",
  [ErrorCode.TEXT_TOO_LONG]:
    "That text is a bit too long. Keep it under 2,000 characters and try again.",
  [ErrorCode.INVALID_SIZE]:
    "The size you chose isn't supported. Pick a value between 50px and 2000px.",
  [ErrorCode.RATE_LIMIT_EXCEEDED]:
    "You're generating codes pretty quickly. Please wait a moment before trying again.",
  [ErrorCode.INTERNAL_ERROR]:
    "Our servers hit a snag. Please try again in a moment.",
};

interface APIErrorShape {
  code?: ErrorCode;
  message?: string;
}

export function getFriendlyErrorMessage(
  error?: APIErrorShape,
  fallback?: string
): string {
  if (!error?.code) {
    return fallback || error?.message || DEFAULT_ERROR_MESSAGE;
  }

  return errorMessages[error.code] || fallback || error.message || DEFAULT_ERROR_MESSAGE;
}

