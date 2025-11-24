import { test, expect } from "@playwright/test";

test.describe("Homepage", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
  });

  test("should display homepage with hero section", async ({ page }) => {
    await expect(page.locator("h1")).toContainText(
      "Generate QR codes in milliseconds"
    );
    await expect(page.locator('input#qr-input')).toBeVisible();
  });

  test("should generate QR code when text is entered", async ({ page }) => {
    const input = page.locator('input#qr-input');
    await input.fill("https://example.com");

    // Wait for API call to complete and QR code to appear
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );

    // Check that QR code image is visible (with longer timeout)
    const qrImage = page.locator('img[alt="Generated QR Code"]');
    await expect(qrImage).toBeVisible({ timeout: 10000 });
  });

  test("should toggle format between PNG and SVG", async ({ page }) => {
    const input = page.locator('input#qr-input');
    await input.fill("test");

    // Wait for initial QR code
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );
    await expect(page.locator('img[alt="Generated QR Code"]')).toBeVisible({
      timeout: 10000,
    });

    // Click SVG button
    await page.locator('button:has-text("SVG")').click();

    // Wait for new API call with SVG format
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );

    // QR code should still be visible
    const qrImage = page.locator('img[alt="Generated QR Code"]');
    await expect(qrImage).toBeVisible({ timeout: 10000 });
  });

  test("should adjust size with slider", async ({ page }) => {
    const input = page.locator('input#qr-input');
    await input.fill("test");

    // Wait for initial QR code
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );
    await expect(page.locator('img[alt="Generated QR Code"]')).toBeVisible({
      timeout: 10000,
    });

    const slider = page.locator('input[type="range"]');
    await slider.fill("500");

    // Wait for new API call with updated size
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );

    // QR code should still be visible
    const qrImage = page.locator('img[alt="Generated QR Code"]');
    await expect(qrImage).toBeVisible({ timeout: 10000 });
  });

  test("should show download buttons after QR generation", async ({ page }) => {
    const input = page.locator('input#qr-input');
    await input.fill("test");

    // Wait for QR code to be generated
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );
    await expect(page.locator('img[alt="Generated QR Code"]')).toBeVisible({
      timeout: 10000,
    });

    // Download buttons should appear
    await expect(
      page.locator('button:has-text("Download PNG")')
    ).toBeVisible({ timeout: 5000 });
    await expect(
      page.locator('button:has-text("Download SVG")')
    ).toBeVisible({ timeout: 5000 });
    await expect(
      page.locator('button:has-text("Copy Base64")')
    ).toBeVisible({ timeout: 5000 });
  });

  test("should display analytics section", async ({ page }) => {
    // Wait for analytics API call
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/analytics") && response.status() === 200
    );

    // Scroll to analytics section
    await page.locator("#analytics").scrollIntoViewIfNeeded();
    await page.waitForTimeout(500); // Allow time for scroll

    // Check for analytics display
    await expect(
      page.locator("text=Total Generations")
    ).toBeVisible({ timeout: 10000 });
    await expect(page.locator("text=Today")).toBeVisible({ timeout: 5000 });
  });

  test("should display API playground section", async ({ page }) => {
    await page.locator("#api-playground").scrollIntoViewIfNeeded();
    await page.waitForTimeout(500); // Allow time for scroll
    await expect(page.locator("text=API Playground")).toBeVisible({
      timeout: 10000,
    });
  });

  test("should display documentation section", async ({ page }) => {
    await page.locator("#api-docs").scrollIntoViewIfNeeded();
    await expect(page.locator("text=Documentation")).toBeVisible();
  });

  test("should handle empty input gracefully", async ({ page }) => {
    const input = page.locator('input#qr-input');
    await input.fill("");
    await page.waitForTimeout(500);

    // Should show placeholder message
    await expect(
      page.locator("text=Enter text above to generate QR code")
    ).toBeVisible();
  });

  test("should be responsive on mobile", async ({ page }) => {
    await page.setViewportSize({ width: 375, height: 667 });
    await page.goto("/");
    await page.waitForLoadState("networkidle");

    // Header should be visible
    await expect(page.locator("header")).toBeVisible({ timeout: 10000 });

    // Input should be visible and usable
    const input = page.locator('input#qr-input');
    await expect(input).toBeVisible({ timeout: 10000 });
    await input.fill("test");

    // Wait for QR code to appear
    await page.waitForResponse(
      (response) =>
        response.url().includes("/api/qr") && response.status() === 200
    );
  });
});

test.describe("Navigation", () => {
  test("should navigate to API docs section", async ({ page }) => {
    await page.goto("/");
    await page.waitForLoadState("networkidle");

    // Use getByRole for better reliability
    const apiDocsLink = page.getByRole("navigation").getByRole("link", {
      name: "API Docs",
    });
    await apiDocsLink.scrollIntoViewIfNeeded();
    await apiDocsLink.click();

    // Wait for scroll/animation
    await page.waitForTimeout(500);
    await expect(page.locator("#api-docs")).toBeVisible({ timeout: 10000 });
  });

  test("should navigate to analytics section", async ({ page }) => {
    await page.goto("/");
    await page.waitForLoadState("networkidle");

    // Use getByRole for better reliability
    const analyticsLink = page.getByRole("navigation").getByRole("link", {
      name: "Analytics",
    });
    await analyticsLink.scrollIntoViewIfNeeded();
    await analyticsLink.click();

    // Wait for scroll/animation
    await page.waitForTimeout(500);
    await expect(page.locator("#analytics")).toBeVisible({ timeout: 10000 });
  });
});

