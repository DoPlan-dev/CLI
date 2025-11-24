"use client";

import React from "react";

interface SizeSliderProps {
  value: number;
  onChange: (value: number) => void;
  min?: number;
  max?: number;
  className?: string;
}

export function SizeSlider({
  value,
  onChange,
  min = 50,
  max = 2000,
  className = "",
}: SizeSliderProps) {
  return (
    <div className={`w-full max-w-xs ${className}`}>
      <div className="flex items-center justify-between mb-2">
        <label
          htmlFor="size-slider"
          className="text-sm font-medium text-text-secondary"
        >
          Size
        </label>
        <span className="text-sm font-semibold text-text-primary">
          {value}px
        </span>
      </div>
      <input
        id="size-slider"
        type="range"
        min={min}
        max={max}
        value={value}
        onChange={(e) => onChange(Number(e.target.value))}
        className="w-full h-1.5 bg-bg-tertiary rounded-full appearance-none cursor-pointer accent-primary"
        style={{
          background: `linear-gradient(to right, var(--color-primary) 0%, var(--color-primary) ${
            ((value - min) / (max - min)) * 100
          }%, var(--color-bg-tertiary) ${
            ((value - min) / (max - min)) * 100
          }%, var(--color-bg-tertiary) 100%)`,
        }}
        aria-label="QR code size"
        aria-valuemin={min}
        aria-valuemax={max}
        aria-valuenow={value}
      />
    </div>
  );
}

