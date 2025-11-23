# QR Code Generator API - Micro SaaS

## Project Overview
Build a QR Code Generator API micro SaaS that provides REST API endpoints for generating QR codes in multiple formats with customizable options.

## Core Features

- REST API endpoint that accepts text or URLs via POST request
- Generate QR codes in PNG and SVG formats
- Support customizable size (default 200x200 pixels)
- Support error correction levels (L, M, Q, H)
- Return base64 encoded image or direct file download
- Simple analytics tracking (generation count, timestamp)

## Technical Stack

- **Backend**: Node.js with Express framework
- **Language**: TypeScript for type safety
- **Database**: SQLite for MVP (easy to upgrade to PostgreSQL later)
- **QR Library**: qrcode npm package
- **Image Processing**: sharp for image manipulation if needed

## API Endpoints

### 1. POST /api/qr - Generate QR code
- **Request body**: 
  ```json
  {
    "text": "string",
    "size": "number (optional)",
    "format": "png|svg (optional)",
    "errorCorrection": "L|M|Q|H (optional)"
  }
  ```
- **Response**: 
  ```json
  {
    "qrCode": "string (base64)",
    "format": "string",
    "size": "number"
  }
  ```
  Or file download

### 2. GET /api/analytics - Get usage statistics
- **Response**: 
  ```json
  {
    "totalGenerations": "number",
    "recentActivity": "array"
  }
  ```

## Frontend

- Simple HTML/CSS/JavaScript interface
- Form to input text/URL
- Preview generated QR code
- Download button for PNG/SVG
- Display analytics

## MVP Scope

- Focus on core QR generation first
- Basic analytics (count only)
- Simple web interface
- No authentication required for MVP

## Future Enhancements (v2.0+)

- Custom colors and styling
- Logo embedding in center
- Batch generation
- API keys for rate limiting
- Advanced analytics dashboard
- Custom domains
