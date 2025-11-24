"use client";

import React from "react";

interface FormatToggleProps {
  format: "png" | "svg";
  onChange: (format: "png" | "svg") => void;
  className?: string;
}

export function FormatToggle({
  format,
  onChange,
  className = "",
}: FormatToggleProps) {
  return (
    <div className={`flex gap-2 p-1 bg-bg-tertiary rounded-md ${className}`}>
      <button
        type="button"
        onClick={() => onChange("png")}
        className={`px-4 py-2 text-sm font-medium rounded-sm transition-all duration-200 ${
          format === "png"
            ? "text-text-primary bg-bg-primary shadow-sm"
            : "text-text-secondary bg-transparent hover:text-text-primary"
        }`}
        aria-pressed={format === "png"}
        aria-label="PNG format"
      >
        PNG
      </button>
      <button
        type="button"
        onClick={() => onChange("svg")}
        className={`px-4 py-2 text-sm font-medium rounded-sm transition-all duration-200 ${
          format === "svg"
            ? "text-text-primary bg-bg-primary shadow-sm"
            : "text-text-secondary bg-transparent hover:text-text-primary"
        }`}
        aria-pressed={format === "svg"}
        aria-label="SVG format"
      >
        SVG
      </button>
    </div>
  );
}

