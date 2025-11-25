/**
 * Activity Point Interface
 * Represents a single data point in the activity timeline
 */
export interface ActivityPoint {
  timestamp: string;
  count: number;
}

/**
 * Analytics Interface
 * Contains aggregated statistics about QR code generations
 */
export interface Analytics {
  totalGenerations: number;
  todayGenerations: number;
  recentActivity: ActivityPoint[];
  apiStatus: "operational" | "degraded" | "down";
  averageResponseTime: number;
}

/**
 * Generation Parameters
 * Used for tracking QR code generation events
 */
export interface GenerationParams {
  textHash: string; // Hashed input text for privacy
  size: number;
  format: string;
  errorCorrection: string;
  responseTimeMs: number;
}

