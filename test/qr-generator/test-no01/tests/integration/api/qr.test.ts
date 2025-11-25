import { describe, it, expect, beforeAll, afterAll } from "vitest";

// Note: Integration tests for Next.js API routes require a running server
// These tests are skipped by default - run with TEST_BASE_URL set and server running
// Example: TEST_BASE_URL=http://localhost:3000 npm test

const baseUrl = process.env.TEST_BASE_URL || "http://localhost:3000";
const skipIntegration = !process.env.TEST_BASE_URL;

describe.skipIf(skipIntegration)("POST /api/qr", () => {

  it("should generate QR code with valid request", async () => {
    const response = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: "https://example.com",
        size: 200,
        format: "png",
        errorCorrection: "M",
      }),
    });

    expect(response.status).toBe(200);
    const data = await response.json();
    expect(data.success).toBe(true);
    expect(data.qrCode).toContain("data:image/png;base64,");
    expect(data.format).toBe("png");
    expect(data.size).toBe(200);
  });

  it("should generate SVG QR code", async () => {
    const response = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: "Hello World",
        format: "svg",
      }),
    });

    expect(response.status).toBe(200);
    const data = await response.json();
    expect(data.success).toBe(true);
    expect(data.format).toBe("svg");
    expect(data.qrCode).toContain("data:image/svg+xml;base64,");
  });

  it("should return error for empty text", async () => {
    const response = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: "",
      }),
    });

    expect(response.status).toBe(400);
    const data = await response.json();
    expect(data.success).toBe(false);
    expect(data.error.code).toBe("INVALID_INPUT");
  });

  it("should return error for text exceeding 2000 characters", async () => {
    const response = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: "a".repeat(2001),
      }),
    });

    expect(response.status).toBe(400);
    const data = await response.json();
    expect(data.success).toBe(false);
  });

  it("should return error for invalid size", async () => {
    const response = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: "test",
        size: 30,
      }),
    });

    expect(response.status).toBe(400);
    const data = await response.json();
    expect(data.success).toBe(false);
  });

  it("should cache identical requests", async () => {
    const request = {
      text: "cache-test",
      size: 200,
      format: "png" as const,
      errorCorrection: "M" as const,
    };

    const response1 = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    });

    const response2 = await fetch(`${baseUrl}/api/qr`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    });

    const data1 = await response1.json();
    const data2 = await response2.json();

    expect(data1.qrCode).toBe(data2.qrCode);
  });
});

describe.skipIf(skipIntegration)("GET /api/analytics", () => {

  it("should return analytics data", async () => {
    const response = await fetch(`${baseUrl}/api/analytics`);

    expect(response.status).toBe(200);
    const data = await response.json();
    expect(data).toHaveProperty("totalGenerations");
    expect(data).toHaveProperty("todayGenerations");
    expect(data).toHaveProperty("recentActivity");
    expect(data).toHaveProperty("apiStatus");
    expect(data).toHaveProperty("averageResponseTime");
  });
});

describe.skipIf(skipIntegration)("GET /api/health", () => {

  it("should return health status", async () => {
    const response = await fetch(`${baseUrl}/api/health`);

    expect(response.status).toBe(200);
    const data = await response.json();
    expect(data).toHaveProperty("status");
    expect(data).toHaveProperty("database");
    expect(data).toHaveProperty("timestamp");
  });
});

