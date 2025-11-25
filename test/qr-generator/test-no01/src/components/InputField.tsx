"use client";

import React from "react";

interface InputFieldProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
  label?: string;
  id?: string;
  className?: string;
}

export function InputField({
  value,
  onChange,
  placeholder = "Enter text or URL...",
  label,
  id = "qr-input",
  className = "",
}: InputFieldProps) {
  return (
    <div className={`w-full max-w-2xl ${className}`}>
      {label && (
        <label
          htmlFor={id}
          className="block text-sm font-medium text-text-secondary mb-2"
        >
          {label}
        </label>
      )}
      <input
        id={id}
        type="text"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
        className="w-full px-6 py-4 text-lg font-normal text-text-primary bg-bg-primary border-2 border-border rounded-md transition-all duration-200 focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
        aria-label={label || "QR code input"}
        aria-required="true"
      />
    </div>
  );
}

