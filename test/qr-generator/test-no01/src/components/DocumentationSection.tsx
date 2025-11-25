"use client";

import React, { useState } from "react";

interface DocumentationSectionProps {
  className?: string;
}

export function DocumentationSection({
  className = "",
}: DocumentationSectionProps) {
  const [expandedSection, setExpandedSection] = useState<string | null>("quick-start");

  const toggleSection = (section: string) => {
    setExpandedSection(expandedSection === section ? null : section);
  };

  return (
    <section id="api-docs" className={`${className}`}>
      <h2 className="text-3xl font-semibold text-center mb-12 text-text-primary">
        Documentation
      </h2>

      <div className="space-y-4 max-w-4xl mx-auto">
        {/* Quick Start */}
        <div className="bg-bg-secondary rounded-lg overflow-hidden">
          <button
            type="button"
            onClick={() => toggleSection("quick-start")}
            className="w-full px-6 py-4 text-left flex items-center justify-between hover:bg-bg-tertiary transition-colors"
          >
            <h3 className="text-xl font-semibold text-text-primary">
              Quick Start
            </h3>
            <svg
              className={`w-5 h-5 text-text-secondary transition-transform ${
                expandedSection === "quick-start" ? "rotate-180" : ""
              }`}
              fill="none"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          {expandedSection === "quick-start" && (
            <div className="px-6 py-4 border-t border-border">
              <ol className="space-y-4 text-text-secondary">
                <li className="flex gap-3">
                  <span className="flex-shrink-0 w-6 h-6 bg-primary text-white rounded-full flex items-center justify-center text-sm font-semibold">
                    1
                  </span>
                  <div>
                    <p className="font-medium text-text-primary mb-1">
                      Enter your text or URL
                    </p>
                    <p>
                      Type any text or URL in the input field above. The QR code
                      will generate automatically as you type.
                    </p>
                  </div>
                </li>
                <li className="flex gap-3">
                  <span className="flex-shrink-0 w-6 h-6 bg-primary text-white rounded-full flex items-center justify-center text-sm font-semibold">
                    2
                  </span>
                  <div>
                    <p className="font-medium text-text-primary mb-1">
                      Customize (optional)
                    </p>
                    <p>
                      Adjust the size, format (PNG/SVG), or error correction
                      level to suit your needs.
                    </p>
                  </div>
                </li>
                <li className="flex gap-3">
                  <span className="flex-shrink-0 w-6 h-6 bg-primary text-white rounded-full flex items-center justify-center text-sm font-semibold">
                    3
                  </span>
                  <div>
                    <p className="font-medium text-text-primary mb-1">
                      Download or use the API
                    </p>
                    <p>
                      Download your QR code or use our REST API to integrate QR
                      generation into your application.
                    </p>
                  </div>
                </li>
              </ol>
            </div>
          )}
        </div>

        {/* API Reference */}
        <div className="bg-bg-secondary rounded-lg overflow-hidden">
          <button
            type="button"
            onClick={() => toggleSection("api-reference")}
            className="w-full px-6 py-4 text-left flex items-center justify-between hover:bg-bg-tertiary transition-colors"
          >
            <h3 className="text-xl font-semibold text-text-primary">
              API Reference
            </h3>
            <svg
              className={`w-5 h-5 text-text-secondary transition-transform ${
                expandedSection === "api-reference" ? "rotate-180" : ""
              }`}
              fill="none"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          {expandedSection === "api-reference" && (
            <div className="px-6 py-4 border-t border-border space-y-6">
              {/* POST /api/qr */}
              <div>
                <h4 className="text-lg font-semibold text-text-primary mb-2">
                  POST /api/qr
                </h4>
                <p className="text-text-secondary mb-4">
                  Generate a QR code from text or URL.
                </p>

                <div className="mb-4">
                  <p className="text-sm font-medium text-text-primary mb-2">
                    Request Body:
                  </p>
                  <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
                    <pre className="text-bg-primary text-sm font-mono">
                      {JSON.stringify(
                        {
                          text: "string (required)",
                          size: "number (optional, 50-2000, default: 200)",
                          format: "string (optional, 'png' | 'svg', default: 'png')",
                          errorCorrection:
                            "string (optional, 'L' | 'M' | 'Q' | 'H', default: 'M')",
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                </div>

                <div className="mb-4">
                  <p className="text-sm font-medium text-text-primary mb-2">
                    Response (200 OK):
                  </p>
                  <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
                    <pre className="text-bg-primary text-sm font-mono">
                      {JSON.stringify(
                        {
                          success: true,
                          qrCode: "data:image/png;base64,...",
                          format: "png",
                          size: 200,
                          errorCorrection: "M",
                          generatedAt: "2024-12-19T10:30:00Z",
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                </div>
              </div>

              {/* GET /api/analytics */}
              <div>
                <h4 className="text-lg font-semibold text-text-primary mb-2">
                  GET /api/analytics
                </h4>
                <p className="text-text-secondary mb-4">
                  Get usage statistics and analytics.
                </p>

                <div className="mb-4">
                  <p className="text-sm font-medium text-text-primary mb-2">
                    Response (200 OK):
                  </p>
                  <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
                    <pre className="text-bg-primary text-sm font-mono">
                      {JSON.stringify(
                        {
                          totalGenerations: 1234,
                          todayGenerations: 56,
                          recentActivity: [
                            { timestamp: "2024-12-19 10:00", count: 5 },
                          ],
                          apiStatus: "operational",
                          averageResponseTime: 45,
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                </div>
              </div>

              {/* GET /api/health */}
              <div>
                <h4 className="text-lg font-semibold text-text-primary mb-2">
                  GET /api/health
                </h4>
                <p className="text-text-secondary mb-4">
                  Health check endpoint.
                </p>

                <div className="mb-4">
                  <p className="text-sm font-medium text-text-primary mb-2">
                    Response (200 OK):
                  </p>
                  <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
                    <pre className="text-bg-primary text-sm font-mono">
                      {JSON.stringify(
                        {
                          status: "operational",
                          database: "healthy",
                          timestamp: "2024-12-19T10:30:00Z",
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* Code Examples */}
        <div className="bg-bg-secondary rounded-lg overflow-hidden">
          <button
            type="button"
            onClick={() => toggleSection("code-examples")}
            className="w-full px-6 py-4 text-left flex items-center justify-between hover:bg-bg-tertiary transition-colors"
          >
            <h3 className="text-xl font-semibold text-text-primary">
              Code Examples
            </h3>
            <svg
              className={`w-5 h-5 text-text-secondary transition-transform ${
                expandedSection === "code-examples" ? "rotate-180" : ""
              }`}
              fill="none"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          {expandedSection === "code-examples" && (
            <div className="px-6 py-4 border-t border-border">
              <p className="text-text-secondary mb-4">
                Use the API Playground above to generate code snippets in your
                preferred language. The playground supports:
              </p>
              <ul className="list-disc list-inside space-y-2 text-text-secondary">
                <li>cURL</li>
                <li>JavaScript (Fetch API)</li>
                <li>Python (requests library)</li>
                <li>Go (net/http)</li>
                <li>PHP (cURL)</li>
              </ul>
            </div>
          )}
        </div>

        {/* FAQ */}
        <div className="bg-bg-secondary rounded-lg overflow-hidden">
          <button
            type="button"
            onClick={() => toggleSection("faq")}
            className="w-full px-6 py-4 text-left flex items-center justify-between hover:bg-bg-tertiary transition-colors"
          >
            <h3 className="text-xl font-semibold text-text-primary">FAQ</h3>
            <svg
              className={`w-5 h-5 text-text-secondary transition-transform ${
                expandedSection === "faq" ? "rotate-180" : ""
              }`}
              fill="none"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          {expandedSection === "faq" && (
            <div className="px-6 py-4 border-t border-border space-y-4">
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  What is the maximum text length?
                </h4>
                <p className="text-text-secondary">
                  The maximum text length is 2000 characters. This ensures QR
                  codes remain scannable and maintain good error correction
                  capabilities.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  What size should I use?
                </h4>
                <p className="text-text-secondary">
                  QR codes can be generated from 50px to 2000px. For most use
                  cases, 200px is a good default. Larger sizes are useful for
                  print materials, while smaller sizes work well for digital
                  displays.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  What&apos;s the difference between PNG and SVG?
                </h4>
                <p className="text-text-secondary">
                  PNG is a raster format (pixel-based) and works well for most
                  use cases. SVG is a vector format that scales without quality
                  loss, making it ideal for print or when you need to resize
                  the QR code.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  What error correction level should I use?
                </h4>
                <p className="text-text-secondary">
                  <strong>L (Low):</strong> ~7% error correction - Use for
                  clean environments. <br />
                  <strong>M (Medium):</strong> ~15% error correction - Default,
                  good for most cases. <br />
                  <strong>Q (Quartile):</strong> ~25% error correction - Use
                  when QR codes may be damaged. <br />
                  <strong>H (High):</strong> ~30% error correction - Maximum
                  error correction, use for challenging environments.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  Is there a rate limit?
                </h4>
                <p className="text-text-secondary">
                  Yes, the API has rate limits: 100 requests per minute and
                  1000 requests per hour per IP address. Rate limit information
                  is included in response headers.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-text-primary mb-1">
                  Is my data stored?
                </h4>
                <p className="text-text-secondary">
                  For privacy, we hash the input text before storing analytics
                  data. The actual text content is never stored in our database.
                </p>
              </div>
            </div>
          )}
        </div>
      </div>
    </section>
  );
}

