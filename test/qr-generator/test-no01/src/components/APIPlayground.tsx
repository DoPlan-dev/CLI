"use client";

import React, { useState, useEffect } from "react";
import { QRRequest, QRResponse } from "@/types/qr";
import { getFriendlyErrorMessage } from "@/lib/utils/client-errors";
import { ErrorCode } from "@/lib/utils/errors";

type CodeLanguage = "curl" | "javascript" | "python" | "go" | "php";

interface APIPlaygroundProps {
  className?: string;
}

export function APIPlayground({ className = "" }: APIPlaygroundProps) {
  const [text, setText] = useState("https://example.com");
  const [size, setSize] = useState(200);
  const [format, setFormat] = useState<"png" | "svg">("png");
  const [errorCorrection, setErrorCorrection] = useState<"L" | "M" | "Q" | "H">(
    "M"
  );
  const [response, setResponse] = useState<QRResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [selectedLanguage, setSelectedLanguage] =
    useState<CodeLanguage>("javascript");
  const [copied, setCopied] = useState(false);
  const [baseUrl, setBaseUrl] = useState("");

  // Set base URL only on client to avoid hydration mismatch
  useEffect(() => {
    setBaseUrl(typeof window !== "undefined" ? window.location.origin : "");
  }, []);

  const request: QRRequest = {
    text,
    size,
    format,
    errorCorrection,
  };

  const handleTest = async () => {
    setLoading(true);
    setError(null);
    setResponse(null);

    try {
      const res = await fetch("/api/qr", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(request),
      });

      if (!res.ok) {
        let errorCode: ErrorCode | undefined;
        let errorMessage: string | undefined;
        try {
          const errorData = await res.json();
          errorCode = errorData.error?.code;
          errorMessage = errorData.error?.message;
        } catch {
          // ignore parse failure
        }
        throw new Error(
          getFriendlyErrorMessage(
            { code: errorCode, message: errorMessage },
            "We couldn't complete that request. Please adjust the inputs and try again."
          )
        );
      }

      const data: QRResponse = await res.json();
      setResponse(data);
    } catch (err) {
      const message =
        err instanceof Error
          ? err.message
          : "We couldn't complete that request. Please try again.";
      setError(message);
    } finally {
      setLoading(false);
    }
  };

  const generateCodeSnippet = (lang: CodeLanguage): string => {
    // Use relative URL for SSR compatibility, absolute URL only after hydration
    const url = baseUrl || "/api/qr";
    const requestBody = JSON.stringify(request, null, 2);

    switch (lang) {
      case "curl":
        return `curl -X POST ${url} \\
  -H "Content-Type: application/json" \\
  -d '${JSON.stringify(request)}'`;

      case "javascript":
        return `fetch("${url}", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify(${requestBody}),
})
  .then((response) => response.json())
  .then((data) => console.log(data))
  .catch((error) => console.error("Error:", error));`;

      case "python":
        return `import requests

response = requests.post(
    "${url}",
    json=${requestBody},
    headers={"Content-Type": "application/json"}
)

print(response.json())`;

      case "go":
        return `package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    requestBody := ${requestBody.replace(
      /"/g,
      '\\"'
    )}
    
    jsonData, _ := json.Marshal(requestBody)
    resp, err := http.Post(
        "${url}",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Println(result)
}`;

      case "php":
        return `<?php

$data = ${requestBody.replace(/"/g, '\\"')};

$ch = curl_init("${url}");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    "Content-Type: application/json"
]);

$response = curl_exec($ch);
curl_close($ch);

echo $response;
?>`;

      default:
        return "";
    }
  };

  const handleCopyCode = async () => {
    const code = generateCodeSnippet(selectedLanguage);
    try {
      await navigator.clipboard.writeText(code);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error("Failed to copy:", err);
    }
  };

  return (
    <section
      id="api-playground"
      className={`bg-bg-secondary rounded-lg p-6 md:p-8 ${className}`}
    >
      <h2 className="text-2xl font-semibold mb-6 text-text-primary">
        API Playground
      </h2>

      {/* Request Builder */}
      <div className="mb-8">
        <h3 className="text-lg font-medium mb-4 text-text-primary">
          Request Builder
        </h3>
        <div className="space-y-4">
          <div>
            <label
              htmlFor="playground-text"
              className="block text-sm font-medium text-text-secondary mb-2"
            >
              Text
            </label>
            <input
              id="playground-text"
              type="text"
              value={text}
              onChange={(e) => setText(e.target.value)}
              className="w-full px-4 py-2 text-base text-text-primary bg-bg-primary border border-border rounded-md focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
              placeholder="Enter text or URL"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label
                htmlFor="playground-size"
                className="block text-sm font-medium text-text-secondary mb-2"
              >
                Size (px)
              </label>
              <input
                id="playground-size"
                type="number"
                min="50"
                max="2000"
                value={size}
                onChange={(e) => setSize(Number(e.target.value))}
                className="w-full px-4 py-2 text-base text-text-primary bg-bg-primary border border-border rounded-md focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
              />
            </div>

            <div>
              <label
                htmlFor="playground-format"
                className="block text-sm font-medium text-text-secondary mb-2"
              >
                Format
              </label>
              <select
                id="playground-format"
                value={format}
                onChange={(e) => setFormat(e.target.value as "png" | "svg")}
                className="w-full px-4 py-2 text-base text-text-primary bg-bg-primary border border-border rounded-md focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
              >
                <option value="png">PNG</option>
                <option value="svg">SVG</option>
              </select>
            </div>

            <div>
              <label
                htmlFor="playground-error-correction"
                className="block text-sm font-medium text-text-secondary mb-2"
              >
                Error Correction
              </label>
              <select
                id="playground-error-correction"
                value={errorCorrection}
                onChange={(e) =>
                  setErrorCorrection(e.target.value as "L" | "M" | "Q" | "H")
                }
                className="w-full px-4 py-2 text-base text-text-primary bg-bg-primary border border-border rounded-md focus:outline-none focus:border-primary focus:ring-3 focus:ring-primary/10"
              >
                <option value="L">L (Low)</option>
                <option value="M">M (Medium)</option>
                <option value="Q">Q (Quartile)</option>
                <option value="H">H (High)</option>
              </select>
            </div>
          </div>

          <button
            onClick={handleTest}
            disabled={loading || !text.trim()}
            className="px-6 py-3 bg-primary text-white font-semibold rounded-md transition-all duration-200 hover:bg-primary-dark hover:-translate-y-0.5 hover:shadow-md disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? "Testing..." : "Test API"}
          </button>
        </div>
      </div>

      {/* Response Viewer */}
      {response && (
        <div className="mb-8">
          <h3 className="text-lg font-medium mb-4 text-text-primary">
            Response
          </h3>
          <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
            <pre className="text-bg-primary text-sm font-mono">
              {JSON.stringify(response, null, 2)}
            </pre>
          </div>
          {response.qrCode && (
            <div className="mt-4 flex justify-center">
              <img
                src={response.qrCode}
                alt="Generated QR Code"
                className="max-w-xs border border-border rounded-md"
              />
            </div>
          )}
        </div>
      )}

      {error && (
        <div className="mb-8 p-4 bg-error/10 border border-error rounded-md">
          <p className="text-error text-sm">{error}</p>
        </div>
      )}

      {/* Code Snippet Generator */}
      <div>
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-lg font-medium text-text-primary">
            Code Snippets
          </h3>
          <button
            onClick={handleCopyCode}
            className="px-4 py-2 text-sm font-medium text-text-secondary bg-bg-primary border border-border rounded-md hover:text-text-primary hover:border-primary transition-colors"
          >
            {copied ? "Copied!" : "Copy"}
          </button>
        </div>

        {/* Language Tabs */}
        <div className="flex gap-2 mb-4 overflow-x-auto">
          {(["curl", "javascript", "python", "go", "php"] as CodeLanguage[]).map(
            (lang) => (
              <button
                key={lang}
                type="button"
                onClick={() => setSelectedLanguage(lang)}
                className={`px-4 py-2 text-sm font-medium rounded-md transition-colors whitespace-nowrap ${
                  selectedLanguage === lang
                    ? "bg-primary text-white"
                    : "bg-bg-tertiary text-text-secondary hover:text-text-primary"
                }`}
              >
                {lang.charAt(0).toUpperCase() + lang.slice(1)}
              </button>
            )
          )}
        </div>

        {/* Code Block */}
        <div className="relative">
          <div className="bg-text-primary rounded-md p-4 overflow-x-auto">
            <pre className="text-bg-primary text-sm font-mono">
              {generateCodeSnippet(selectedLanguage)}
            </pre>
          </div>
        </div>
      </div>
    </section>
  );
}

