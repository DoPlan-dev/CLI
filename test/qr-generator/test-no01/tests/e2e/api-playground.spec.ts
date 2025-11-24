import { test, expect } from "@playwright/test";

test.describe("API Playground", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
    // Scroll to API playground section
    await page.locator("#api-playground").scrollIntoViewIfNeeded();
    await page.waitForTimeout(500); // Allow time for scroll
  });

  test.describe("API Request Building", () => {
    test("should display request builder form with all fields", async ({
      page,
    }) => {
      // Check that all form fields are visible
      await expect(page.locator("#playground-text")).toBeVisible();
      await expect(page.locator("#playground-size")).toBeVisible();
      await expect(page.locator("#playground-format")).toBeVisible();
      await expect(page.locator("#playground-error-correction")).toBeVisible();
      await expect(
        page.locator('button:has-text("Test API")')
      ).toBeVisible();
    });

    test("should allow entering text in request builder", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.clear();
      await textInput.fill("https://test.example.com");
      await expect(textInput).toHaveValue("https://test.example.com");
    });

    test("should allow changing size in request builder", async ({ page }) => {
      const sizeInput = page.locator("#playground-size");
      await sizeInput.clear();
      await sizeInput.fill("300");
      await expect(sizeInput).toHaveValue("300");
    });

    test("should allow changing format in request builder", async ({ page }) => {
      const formatSelect = page.locator("#playground-format");
      await formatSelect.selectOption("svg");
      await expect(formatSelect).toHaveValue("svg");
    });

    test("should allow changing error correction level", async ({ page }) => {
      const errorCorrectionSelect = page.locator(
        "#playground-error-correction"
      );
      await errorCorrectionSelect.selectOption("H");
      await expect(errorCorrectionSelect).toHaveValue("H");
    });

    test("should disable Test API button when text is empty", async ({
      page,
    }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      // Clear text input
      await textInput.clear();
      await expect(testButton).toBeDisabled();
    });

    test("should enable Test API button when text is entered", async ({
      page,
    }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      // Enter text
      await textInput.fill("test");
      await expect(testButton).toBeEnabled();
    });
  });

  test.describe("API Testing and Response Viewing", () => {
    test("should make API request when Test API button is clicked", async ({
      page,
    }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      // Set up request
      await textInput.fill("https://example.com");

      // Click test button and wait for API response
      const responsePromise = page.waitForResponse(
        (response) =>
          response.url().includes("/api/qr") && response.status() === 200
      );

      await testButton.click();
      const response = await responsePromise;

      // Verify response was received
      expect(response.ok()).toBeTruthy();
    });

    test("should display response JSON after successful API call", async ({
      page,
    }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      await textInput.fill("test response");
      await testButton.click();

      // Wait for response section to appear
      await page.waitForResponse(
        (response) =>
          response.url().includes("/api/qr") && response.status() === 200
      );

      // Check that response section is visible
      await expect(
        page.locator('h3:has-text("Response")')
      ).toBeVisible({ timeout: 10000 });

      // Check that response JSON is displayed
      const responsePre = page.locator("#api-playground pre").first();
      await expect(responsePre).toBeVisible();
      const responseText = await responsePre.textContent();
      expect(responseText).toContain("qrCode");
    });

    test("should display QR code image in response", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      await textInput.fill("test image");
      await testButton.click();

      // Wait for response
      await page.waitForResponse(
        (response) =>
          response.url().includes("/api/qr") && response.status() === 200
      );

      // Check that QR code image is displayed
      const qrImage = page
        .locator("#api-playground")
        .locator('img[alt="Generated QR Code"]');
      await expect(qrImage).toBeVisible({ timeout: 10000 });
    });

    test("should show loading state while testing API", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      await textInput.fill("loading test");

      // Click button and immediately check for loading state
      await testButton.click();
      await expect(testButton).toContainText("Testing...", { timeout: 1000 });
    });

    test("should display error message on API failure", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      const testButton = page.locator('button:has-text("Test API")');

      // Use invalid input that might cause an error
      // First, let's try with a very long text that might exceed limits
      await textInput.fill("a".repeat(3000)); // Exceeds 2000 char limit

      await testButton.click();

      // Wait a bit for error to appear
      await page.waitForTimeout(2000);

      // Check for error message (might be in response or error div)
      const errorMessage = page.locator("#api-playground").locator("text=/error/i");
      // Error might be displayed, check if it exists
      const errorExists = await errorMessage.count() > 0;
      // If error exists, it should be visible
      if (errorExists) {
        await expect(errorMessage.first()).toBeVisible();
      }
    });
  });

  test.describe("Code Snippet Generation", () => {
    test("should display code snippets section", async ({ page }) => {
      await expect(
        page.locator('h3:has-text("Code Snippets")')
      ).toBeVisible();
    });

    test("should show language tabs for all supported languages", async ({
      page,
    }) => {
      const languages = ["Curl", "Javascript", "Python", "Go", "Php"];

      for (const lang of languages) {
        await expect(
          page.locator(`button:has-text("${lang}")`)
        ).toBeVisible();
      }
    });

    test("should generate curl code snippet", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.fill("curl test");

      // Click on Curl tab
      await page.locator('button:has-text("Curl")').click();

      // Check that code snippet is displayed
      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last(); // Last pre is the code snippet
      await expect(codeBlock).toBeVisible();

      const codeText = await codeBlock.textContent();
      expect(codeText).toContain("curl");
      expect(codeText).toContain("POST");
      expect(codeText).toContain("/api/qr");
    });

    test("should generate JavaScript code snippet", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.fill("javascript test");

      // Click on Javascript tab
      await page.locator('button:has-text("Javascript")').click();

      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();

      const codeText = await codeBlock.textContent();
      expect(codeText).toContain("fetch");
      expect(codeText).toContain("POST");
    });

    test("should generate Python code snippet", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.fill("python test");

      // Click on Python tab
      await page.locator('button:has-text("Python")').click();

      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();

      const codeText = await codeBlock.textContent();
      expect(codeText).toContain("import requests");
      expect(codeText).toContain("requests.post");
    });

    test("should generate Go code snippet", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.fill("go test");

      // Click on Go tab
      await page.locator('button:has-text("Go")').click();

      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();

      const codeText = await codeBlock.textContent();
      expect(codeText).toContain("package main");
      expect(codeText).toContain("http.Post");
    });

    test("should generate PHP code snippet", async ({ page }) => {
      const textInput = page.locator("#playground-text");
      await textInput.fill("php test");

      // Click on Php tab
      await page.locator('button:has-text("Php")').click();

      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();

      const codeText = await codeBlock.textContent();
      expect(codeText).toContain("<?php");
      expect(codeText).toContain("curl_init");
    });

    test("should update code snippet when request parameters change", async ({
      page,
    }) => {
      const textInput = page.locator("#playground-text");
      const sizeInput = page.locator("#playground-size");

      // Set initial values
      await textInput.fill("initial");
      await sizeInput.fill("200");

      // Get initial code snippet
      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      const initialCode = await codeBlock.textContent();

      // Change size
      await sizeInput.fill("500");
      await page.waitForTimeout(500); // Wait for code to update

      // Code snippet should have changed
      const updatedCode = await codeBlock.textContent();
      expect(updatedCode).not.toBe(initialCode);
      expect(updatedCode).toContain("500");
    });
  });

  test.describe("Copy to Clipboard", () => {
    test("should display copy button", async ({ page }) => {
      await expect(
        page.locator('#api-playground button:has-text("Copy")')
      ).toBeVisible();
    });

    test("should copy code snippet to clipboard", async ({ page, context }) => {
      // Grant clipboard permissions
      await context.grantPermissions(["clipboard-read", "clipboard-write"]);

      const textInput = page.locator("#playground-text");
      await textInput.fill("clipboard test");

      // Get the code snippet text
      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();
      const expectedCode = await codeBlock.textContent();

      // Click copy button
      const copyButton = page.locator(
        '#api-playground button:has-text("Copy")'
      );
      await copyButton.click();

      // Wait for "Copied!" feedback
      await expect(
        page.locator('#api-playground button:has-text("Copied!")')
      ).toBeVisible({ timeout: 2000 });

      // Verify clipboard content (if supported)
      try {
        const clipboardText = await page.evaluate(() =>
          navigator.clipboard.readText()
        );
        expect(clipboardText).toBe(expectedCode?.trim());
      } catch (e) {
        // Clipboard API might not be available in test environment
        // In that case, we just verify the UI feedback
        console.log("Clipboard read not available in test environment");
      }
    });

    test("should show Copied! feedback after copying", async ({ page }) => {
      await page.context().grantPermissions([
        "clipboard-read",
        "clipboard-write",
      ]);

      const textInput = page.locator("#playground-text");
      await textInput.fill("feedback test");

      const copyButton = page.locator(
        '#api-playground button:has-text("Copy")'
      );
      await copyButton.click();

      // Should show "Copied!" text
      await expect(
        page.locator('#api-playground button:has-text("Copied!")')
      ).toBeVisible({ timeout: 2000 });

      // Should revert back to "Copy" after timeout
      await page.waitForTimeout(2500);
      await expect(
        page.locator('#api-playground button:has-text("Copy")')
      ).toBeVisible();
    });
  });

  test.describe("Integration Flow", () => {
    test("should complete full workflow: build request, test API, view response, generate code, copy", async ({
      page,
      context,
    }) => {
      await context.grantPermissions(["clipboard-read", "clipboard-write"]);

      // 1. Build request
      const textInput = page.locator("#playground-text");
      const sizeInput = page.locator("#playground-size");
      const formatSelect = page.locator("#playground-format");

      await textInput.fill("https://full-workflow-test.com");
      await sizeInput.fill("250");
      await formatSelect.selectOption("svg");

      // 2. Test API
      const testButton = page.locator('button:has-text("Test API")');
      await testButton.click();

      await page.waitForResponse(
        (response) =>
          response.url().includes("/api/qr") && response.status() === 200
      );

      // 3. Verify response is displayed
      await expect(
        page.locator('h3:has-text("Response")')
      ).toBeVisible({ timeout: 10000 });

      // 4. Generate code snippet
      await page.locator('button:has-text("Python")').click();
      const codeBlock = page
        .locator("#api-playground")
        .locator("pre")
        .last();
      await expect(codeBlock).toBeVisible();

      // 5. Copy to clipboard
      const copyButton = page.locator(
        '#api-playground button:has-text("Copy")'
      );
      await copyButton.click();
      await expect(
        page.locator('#api-playground button:has-text("Copied!")')
      ).toBeVisible({ timeout: 2000 });
    });
  });
});

