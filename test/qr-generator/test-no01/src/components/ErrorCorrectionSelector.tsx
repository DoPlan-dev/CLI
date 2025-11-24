"use client";

import React from "react";

type ErrorCorrectionLevel = "L" | "M" | "Q" | "H";

interface ErrorCorrectionSelectorProps {
  value: ErrorCorrectionLevel;
  onChange: (value: ErrorCorrectionLevel) => void;
  className?: string;
}

const errorCorrectionInfo: Record<
  ErrorCorrectionLevel,
  { label: string; description: string }
> = {
  L: { label: "L (Low)", description: "~7% error correction" },
  M: { label: "M (Medium)", description: "~15% error correction" },
  Q: { label: "Q (Quartile)", description: "~25% error correction" },
  H: { label: "H (High)", description: "~30% error correction" },
};

export function ErrorCorrectionSelector({
  value,
  onChange,
  className = "",
}: ErrorCorrectionSelectorProps) {
  return (
    <div className={className}>
      <label
        htmlFor="error-correction"
        className="block text-sm font-medium text-text-secondary mb-2"
      >
        Error Correction
      </label>
      <select
        id="error-correction"
        value={value}
        onChange={(e) => onChange(e.target.value as ErrorCorrectionLevel)}
        className="w-full px-4 py-2 text-sm font-medium text-text-primary bg-bg-primary border border-border rounded-md transition-all duration-200 focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
        aria-label="Error correction level"
      >
        {Object.entries(errorCorrectionInfo).map(([level, info]) => (
          <option key={level} value={level}>
            {info.label} - {info.description}
          </option>
        ))}
      </select>
    </div>
  );
}

