"use client";

import React, { useState } from "react";
import { InputField } from "@/components/InputField";
import { FormatToggle } from "@/components/FormatToggle";
import { SizeSlider } from "@/components/SizeSlider";
import { ErrorCorrectionSelector } from "@/components/ErrorCorrectionSelector";
import { QRPreview } from "@/components/QRPreview";
import { DownloadActions } from "@/components/DownloadActions";
import { AnalyticsDisplay } from "@/components/AnalyticsDisplay";
import { APIPlayground } from "@/components/APIPlayground";
import { DocumentationSection } from "@/components/DocumentationSection";
import { QRRequest } from "@/types/qr";

export default function Home() {
  const [text, setText] = useState("");
  const [format, setFormat] = useState<"png" | "svg">("png");
  const [size, setSize] = useState(200);
  const [errorCorrection, setErrorCorrection] = useState<"L" | "M" | "Q" | "H">(
    "M"
  );
  const [qrCodeBase64, setQrCodeBase64] = useState<string | null>(null);

  const qrRequest: QRRequest = {
    text,
    size,
    format,
    errorCorrection,
  };

  return (
    <main className="min-h-screen bg-bg-primary">
      {/* Hero Section */}
      <section className="container mx-auto px-4 sm:px-6 lg:px-8 py-12 md:py-16 lg:py-24">
        <div className="max-w-4xl mx-auto">
          {/* Tagline */}
          <h1 className="text-4xl md:text-5xl font-bold text-center mb-4 text-text-primary">
            Generate QR codes in milliseconds
          </h1>
          <p className="text-lg text-center text-text-secondary mb-12">
            Fast, simple, and developer-friendly QR code generation API
          </p>

          {/* Input Field */}
          <div className="mb-8 flex justify-center">
            <InputField
              value={text}
              onChange={setText}
              placeholder="Enter text or URL to generate QR code..."
            />
          </div>

          {/* Options */}
          <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mb-8">
            <FormatToggle format={format} onChange={setFormat} />
            <SizeSlider value={size} onChange={setSize} />
            <div className="w-full sm:w-auto sm:min-w-[200px]">
              <ErrorCorrectionSelector
                value={errorCorrection}
                onChange={setErrorCorrection}
              />
            </div>
          </div>

          {/* QR Preview */}
          <div className="flex justify-center mb-8">
            <QRPreview
              text={text}
              size={size}
              format={format}
              errorCorrection={errorCorrection}
              onQRGenerated={setQrCodeBase64}
            />
          </div>

          {/* Download Actions */}
          {qrCodeBase64 && (
            <div className="mb-12 slide-up">
              <DownloadActions
                qrRequest={qrRequest}
                qrCodeBase64={qrCodeBase64}
              />
            </div>
          )}
        </div>
      </section>

      {/* Analytics Section */}
      <section
        id="analytics"
        className="container mx-auto px-4 sm:px-6 lg:px-8 py-12"
      >
        <div className="max-w-4xl mx-auto">
          <h2 className="text-2xl font-semibold text-center mb-6 text-text-primary">
            Live Statistics
          </h2>
          <AnalyticsDisplay />
        </div>
      </section>

      {/* API Playground Section */}
      <section className="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="max-w-4xl mx-auto">
          <APIPlayground />
        </div>
      </section>

      {/* Documentation Section */}
      <section className="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="max-w-4xl mx-auto">
          <DocumentationSection />
        </div>
      </section>
    </main>
  );
}
