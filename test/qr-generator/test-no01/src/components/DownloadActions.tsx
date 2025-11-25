"use client";

import React, { useState } from "react";
import { QRRequest } from "@/types/qr";

interface DownloadActionsProps {
  qrRequest: QRRequest;
  qrCodeBase64: string | null;
  className?: string;
}

export function DownloadActions({
  qrRequest,
  qrCodeBase64,
  className = "",
}: DownloadActionsProps) {
  const [copied, setCopied] = useState(false);
  const [downloading, setDownloading] = useState<string | null>(null);

  if (!qrCodeBase64) {
    return null;
  }

  const handleDownload = async (format: "png" | "svg") => {
    setDownloading(format);
    try {
      const response = await fetch("/api/qr", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/octet-stream",
        },
        body: JSON.stringify({
          ...qrRequest,
          format,
        }),
      });

      if (!response.ok) {
        throw new Error("Download failed");
      }

      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = `qr-code-${Date.now()}.${format}`;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Download error:", error);
      alert("Failed to download QR code");
    } finally {
      setDownloading(null);
    }
  };

  const handleCopyBase64 = async () => {
    try {
      await navigator.clipboard.writeText(qrCodeBase64);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (error) {
      console.error("Copy error:", error);
      alert("Failed to copy to clipboard");
    }
  };

  return (
    <div className={`flex gap-3 flex-wrap justify-center ${className}`}>
      <button
        onClick={() => handleDownload("png")}
        disabled={downloading === "png"}
        className="flex items-center gap-2 px-5 py-3 text-sm font-medium text-text-primary bg-bg-primary border border-border rounded-md transition-all duration-200 hover:border-primary hover:text-primary hover:-translate-y-0.5 hover:shadow-sm disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {downloading === "png" ? "Downloading..." : "Download PNG"}
      </button>
      <button
        onClick={() => handleDownload("svg")}
        disabled={downloading === "svg"}
        className="flex items-center gap-2 px-5 py-3 text-sm font-medium text-text-primary bg-bg-primary border border-border rounded-md transition-all duration-200 hover:border-primary hover:text-primary hover:-translate-y-0.5 hover:shadow-sm disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {downloading === "svg" ? "Downloading..." : "Download SVG"}
      </button>
      <button
        onClick={handleCopyBase64}
        className="flex items-center gap-2 px-5 py-3 text-sm font-medium text-text-primary bg-bg-primary border border-border rounded-md transition-all duration-200 hover:border-primary hover:text-primary hover:-translate-y-0.5 hover:shadow-sm"
      >
        {copied ? "Copied!" : "Copy Base64"}
      </button>
    </div>
  );
}

