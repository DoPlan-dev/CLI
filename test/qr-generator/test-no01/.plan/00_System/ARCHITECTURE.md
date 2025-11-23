# Technical Architecture Document
## QR Code Generator API - Micro SaaS

**Version**: 1.0  
**Date**: 2024-12-19  
**Authors**: Engineering Lead, System Architect  
**Status**: Draft

---

## 1. Architecture Overview

### 1.1 System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Client Layer                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │   Web App    │  │  API Clients │  │  Mobile Web  │     │
│  │  (Next.js)   │  │  (cURL, etc) │  │  (Responsive)│     │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘     │
└─────────┼──────────────────┼──────────────────┼──────────────┘
          │                  │                  │
          └──────────────────┼──────────────────┘
                             │
          ┌──────────────────▼──────────────────┐
          │      Edge Network (CDN)              │
          │  ┌──────────────────────────────┐   │
          │  │   Static Assets (Edge)       │   │
          │  └──────────────────────────────┘   │
          └──────────────────┬───────────────────┘
                             │
          ┌──────────────────▼──────────────────┐
          │    Application Layer (Edge Functions) │
          │  ┌──────────────────────────────┐   │
          │  │   Next.js API Routes         │   │
          │  │   /api/qr                    │   │
          │  │   /api/analytics             │   │
          │  └──────────────────────────────┘   │
          └──────────────────┬───────────────────┘
                             │
          ┌──────────────────▼──────────────────┐
          │         Service Layer                │
          │  ┌──────────────┐  ┌──────────────┐ │
          │  │ QR Service   │  │ Analytics    │ │
          │  │ (qrcode lib) │  │ Service      │ │
          │  └──────────────┘  └──────────────┘ │
          └──────────────────┬───────────────────┘
                             │
          ┌──────────────────▼──────────────────┐
          │         Data Layer                   │
          │  ┌──────────────┐  ┌──────────────┐ │
          │  │   SQLite     │  │   Cache      │ │
          │  │  (Analytics) │  │  (Redis/     │ │
          │  │              │  │   In-Memory) │ │
          │  └──────────────┘  └──────────────┘ │
          └──────────────────────────────────────┘
```

### 1.2 Architecture Principles

1. **Edge-First**: Deploy compute at the edge for lowest latency
2. **Stateless Services**: API routes are stateless for horizontal scaling
3. **Caching Strategy**: Aggressive caching for identical requests
4. **Fail Fast**: Quick error responses, no long-running operations
5. **Type Safety**: TypeScript throughout for reliability
6. **Minimal Dependencies**: Keep bundle size small, dependencies minimal

---

## 2. Technology Stack

### 2.1 Frontend
- **Framework**: Next.js 14+ (App Router)
- **Language**: TypeScript (strict mode)
- **UI Library**: React 18+
- **Styling**: Tailwind CSS
- **State Management**: React hooks (useState, useEffect)
- **Build Tool**: Next.js built-in (Turbopack)

### 2.2 Backend
- **Runtime**: Node.js 20+ (LTS)
- **Framework**: Next.js API Routes (serverless functions)
- **Language**: TypeScript
- **QR Generation**: `qrcode` npm package
- **Image Processing**: `sharp` (optional, for advanced manipulation)

### 2.3 Database
- **MVP**: SQLite (via `better-sqlite3` or `sql.js`)
- **Future**: PostgreSQL (via Prisma ORM)
- **Schema**: Simple analytics table (generations, timestamps)

### 2.4 Caching
- **In-Memory**: Node.js Map for request deduplication
- **Edge Cache**: Vercel/Cloudflare edge caching
- **Future**: Redis for distributed caching

### 2.5 Deployment
- **Platform**: Vercel (primary) or Cloudflare Pages
- **Edge Functions**: Vercel Edge Functions / Cloudflare Workers
- **CDN**: Automatic via platform
- **Monitoring**: Vercel Analytics + Sentry

### 2.6 Development Tools
- **Package Manager**: npm or pnpm
- **Linting**: ESLint + Prettier
- **Testing**: Vitest (unit), Playwright (e2e)
- **Type Checking**: TypeScript compiler
- **CI/CD**: GitHub Actions

---

## 3. System Components

### 3.1 API Layer (`/api`)

#### 3.1.1 POST /api/qr
**Purpose**: Generate QR code

**Implementation**:
```typescript
// app/api/qr/route.ts
export async function POST(request: Request) {
  const body = await request.json();
  const { text, size = 200, format = 'png', errorCorrection = 'M' } = body;
  
  // Validate input
  // Generate QR code using qrcode library
  // Return base64 or file based on Accept header
}
```

**Flow**:
1. Validate request body (text, size, format, errorCorrection)
2. Check cache for identical request
3. Generate QR code using `qrcode` library
4. Convert to requested format (PNG/SVG)
5. Return response (JSON with base64 or binary file)
6. Log analytics event

**Performance Optimizations**:
- Request deduplication (cache identical requests)
- Edge function deployment (low latency)
- Lazy image processing (only if needed)

#### 3.1.2 GET /api/analytics
**Purpose**: Get usage statistics

**Implementation**:
```typescript
// app/api/analytics/route.ts
export async function GET() {
  // Query SQLite for total count
  // Query recent activity (last 24 hours)
  // Return aggregated statistics
}
```

**Flow**:
1. Query database for total generations
2. Query recent activity (aggregated by time intervals)
3. Calculate API status (based on recent errors)
4. Return JSON response

**Caching**: Cache response for 1 minute (stale-while-revalidate)

### 3.2 QR Service

**Location**: `lib/services/qr-service.ts`

**Responsibilities**:
- QR code generation logic
- Format conversion (PNG/SVG)
- Error correction level handling
- Size validation and optimization

**Implementation**:
```typescript
class QRService {
  async generateQR(text: string, options: QROptions): Promise<QRResult> {
    // Generate QR code using qrcode library
    // Convert to requested format
    // Return base64 or buffer
  }
  
  validateInput(text: string, size: number): boolean {
    // Validate text length, size constraints
  }
}
```

### 3.3 Analytics Service

**Location**: `lib/services/analytics-service.ts`

**Responsibilities**:
- Track QR code generations
- Aggregate statistics
- Query recent activity

**Implementation**:
```typescript
class AnalyticsService {
  async trackGeneration(params: GenerationParams): Promise<void> {
    // Insert into SQLite database
  }
  
  async getStatistics(): Promise<Analytics> {
    // Query aggregated statistics
  }
}
```

### 3.4 Cache Service

**Location**: `lib/services/cache-service.ts`

**Responsibilities**:
- Request deduplication
- QR code result caching
- Cache invalidation

**Implementation**:
```typescript
class CacheService {
  private cache: Map<string, CachedQR>;
  
  get(key: string): CachedQR | null {
    // Return cached result if exists and not expired
  }
  
  set(key: string, value: CachedQR, ttl: number): void {
    // Store in cache with TTL
  }
}
```

### 3.5 Frontend Components

#### 3.5.1 Homepage (`app/page.tsx`)
- Hero section with input field
- Live QR preview component
- Format and size controls
- Download actions
- API playground section
- Analytics display

#### 3.5.2 QR Preview Component
```typescript
// components/QRPreview.tsx
export function QRPreview({ text, size, format }: Props) {
  // Debounced input handling
  // Real-time QR generation
  // Smooth animations
}
```

#### 3.5.3 API Playground Component
```typescript
// components/APIPlayground.tsx
export function APIPlayground() {
  // Interactive API testing
  // Request builder
  // Response viewer
  // Code snippet generator
}
```

---

## 4. Data Models

### 4.1 Database Schema

#### Analytics Table
```sql
CREATE TABLE generations (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  text_hash TEXT NOT NULL,  -- Hash of input text for privacy
  size INTEGER NOT NULL,
  format TEXT NOT NULL,
  error_correction TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  response_time_ms INTEGER
);

CREATE INDEX idx_created_at ON generations(created_at);
CREATE INDEX idx_text_hash ON generations(text_hash);
```

### 4.2 TypeScript Interfaces

```typescript
// types/qr.ts
interface QRRequest {
  text: string;
  size?: number;  // 50-2000, default 200
  format?: 'png' | 'svg';  // default 'png'
  errorCorrection?: 'L' | 'M' | 'Q' | 'H';  // default 'M'
}

interface QRResponse {
  success: boolean;
  qrCode: string;  // base64 encoded
  format: string;
  size: number;
  errorCorrection: string;
  generatedAt: string;
}

interface Analytics {
  totalGenerations: number;
  todayGenerations: number;
  recentActivity: ActivityPoint[];
  apiStatus: 'operational' | 'degraded' | 'down';
  averageResponseTime: number;
}

interface ActivityPoint {
  timestamp: string;
  count: number;
}
```

---

## 5. API Design

### 5.1 RESTful Endpoints

| Method | Endpoint | Purpose | Auth Required |
|--------|----------|---------|---------------|
| POST | `/api/qr` | Generate QR code | No (MVP) |
| GET | `/api/analytics` | Get statistics | No (MVP) |
| GET | `/api/health` | Health check | No |

### 5.2 Request/Response Formats

#### POST /api/qr
**Request**:
```json
{
  "text": "https://example.com",
  "size": 200,
  "format": "png",
  "errorCorrection": "M"
}
```

**Response (JSON)**:
```json
{
  "success": true,
  "qrCode": "data:image/png;base64,iVBORw0KG...",
  "format": "png",
  "size": 200,
  "errorCorrection": "M",
  "generatedAt": "2024-12-19T10:30:00Z"
}
```

**Response (File)**:
- Content-Type: `image/png` or `image/svg+xml`
- Content-Disposition: `attachment; filename="qr-code-{timestamp}.{format}"`
- Binary image data

### 5.3 Error Handling

**Error Response Format**:
```json
{
  "success": false,
  "error": {
    "code": "INVALID_INPUT",
    "message": "Text must be between 1 and 2000 characters",
    "field": "text"
  }
}
```

**Error Codes**:
- `INVALID_INPUT`: Invalid request parameters
- `TEXT_TOO_LONG`: Text exceeds maximum length
- `INVALID_SIZE`: Size out of valid range
- `RATE_LIMIT_EXCEEDED`: Too many requests
- `INTERNAL_ERROR`: Server error

### 5.4 Rate Limiting

**MVP Strategy**: IP-based rate limiting
- 100 requests per minute per IP
- 1000 requests per hour per IP
- Headers: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`

---

## 6. Performance Optimization

### 6.1 Caching Strategy

1. **Request Deduplication**: Cache identical requests (same text + params) for 1 hour
2. **Edge Caching**: Cache static assets and API responses at edge
3. **Browser Caching**: Set appropriate cache headers for static resources
4. **Database Query Caching**: Cache analytics queries for 1 minute

### 6.2 Code Optimization

1. **Lazy Loading**: Load QR library only when needed
2. **Tree Shaking**: Remove unused code from bundle
3. **Image Optimization**: Use appropriate formats, compression
4. **Bundle Splitting**: Separate vendor code from app code

### 6.3 Network Optimization

1. **Edge Deployment**: Deploy API routes to edge functions
2. **CDN**: Serve static assets from CDN
3. **Compression**: Gzip/Brotli compression for responses
4. **HTTP/2**: Use HTTP/2 for multiplexing

---

## 7. Security Considerations

### 7.1 Input Validation
- Validate text length (1-2000 characters)
- Validate size range (50-2000 pixels)
- Sanitize input to prevent injection attacks
- Rate limiting to prevent abuse

### 7.2 Data Privacy
- Hash input text before storing (don't store actual text)
- No PII collection in MVP
- Analytics are aggregated (no individual tracking)

### 7.3 API Security
- HTTPS only (TLS 1.2+)
- CORS configuration (allow specific origins)
- Rate limiting to prevent DDoS
- Input size limits

### 7.4 Infrastructure Security
- Environment variables for secrets
- No secrets in code or logs
- Regular dependency updates
- Security headers (CSP, HSTS, etc.)

---

## 8. Monitoring & Observability

### 8.1 Logging
- Structured logging (JSON format)
- Log levels: ERROR, WARN, INFO, DEBUG
- Request/response logging (sanitized)
- Error stack traces

### 8.2 Metrics
- API response times (P50, P95, P99)
- Request count by endpoint
- Error rate by endpoint
- Cache hit rate
- Database query performance

### 8.3 Alerting
- Error rate threshold (>1%)
- Response time threshold (P95 >200ms)
- Uptime monitoring
- Database connection issues

### 8.4 Tools
- **Error Tracking**: Sentry
- **Performance**: Vercel Analytics, Web Vitals
- **Logging**: Vercel Logs or CloudWatch
- **Uptime**: UptimeRobot or similar

---

## 9. Testing Strategy

### 9.1 Unit Tests
- QR service functions
- Analytics service functions
- Input validation
- Error handling

**Tools**: Vitest

### 9.2 Integration Tests
- API endpoint testing
- Database operations
- Cache operations

**Tools**: Vitest + Supertest

### 9.3 End-to-End Tests
- Homepage user flow
- QR generation flow
- Download functionality
- API playground

**Tools**: Playwright

### 9.4 Performance Tests
- Load testing (10,000 req/min)
- Stress testing
- Response time benchmarks

**Tools**: k6 or Artillery

---

## 10. Deployment Architecture

### 10.1 Development Environment
- Local development with Next.js dev server
- SQLite database file
- Hot reload enabled

### 10.2 Staging Environment
- Vercel preview deployments
- Separate database (SQLite or PostgreSQL)
- Full monitoring enabled

### 10.3 Production Environment
- Vercel production deployment
- Edge functions enabled
- CDN for static assets
- Production database
- Full monitoring and alerting

### 10.4 CI/CD Pipeline

```
Git Push → GitHub Actions
  ├─ Lint & Type Check
  ├─ Unit Tests
  ├─ Build
  └─ Deploy to Vercel
      ├─ Preview (PR)
      └─ Production (main branch)
```

---

## 11. Scalability Considerations

### 11.1 Horizontal Scaling
- Stateless API routes (no session state)
- Edge functions scale automatically
- Database connection pooling (future)

### 11.2 Vertical Scaling
- Optimize QR generation algorithm
- Database query optimization
- Cache frequently accessed data

### 11.3 Future Enhancements
- Redis for distributed caching
- PostgreSQL for better concurrency
- Read replicas for analytics queries
- Queue system for batch operations

---

## 12. Migration Path

### 12.1 MVP to v2.0
1. Add user authentication
2. Migrate SQLite to PostgreSQL
3. Add API keys and per-user rate limiting
4. Implement advanced features (colors, logos)

### 12.2 Database Migration
- Use Prisma for database migrations
- Zero-downtime migration strategy
- Data validation and rollback plan

---

## 13. File Structure

```
qr-generator-api/
├── app/
│   ├── api/
│   │   ├── qr/
│   │   │   └── route.ts
│   │   ├── analytics/
│   │   │   └── route.ts
│   │   └── health/
│   │       └── route.ts
│   ├── page.tsx              # Homepage
│   ├── layout.tsx
│   └── globals.css
├── components/
│   ├── QRPreview.tsx
│   ├── APIPlayground.tsx
│   ├── AnalyticsDisplay.tsx
│   └── DownloadActions.tsx
├── lib/
│   ├── services/
│   │   ├── qr-service.ts
│   │   ├── analytics-service.ts
│   │   └── cache-service.ts
│   ├── db/
│   │   └── database.ts
│   └── utils/
│       ├── validation.ts
│       └── errors.ts
├── types/
│   ├── qr.ts
│   └── analytics.ts
├── public/
│   └── (static assets)
├── tests/
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── package.json
├── tsconfig.json
├── next.config.js
└── README.md
```

---

## 14. Dependencies

### 14.1 Production Dependencies
```json
{
  "next": "^14.0.0",
  "react": "^18.0.0",
  "react-dom": "^18.0.0",
  "qrcode": "^1.5.3",
  "better-sqlite3": "^9.0.0",
  "sharp": "^0.32.0"
}
```

### 14.2 Development Dependencies
```json
{
  "typescript": "^5.0.0",
  "@types/node": "^20.0.0",
  "@types/react": "^18.0.0",
  "@types/qrcode": "^1.5.0",
  "vitest": "^1.0.0",
  "playwright": "^1.40.0",
  "eslint": "^8.0.0",
  "prettier": "^3.0.0"
}
```

---

## 15. Risk Mitigation

### 15.1 Technical Risks
- **QR Library Issues**: Have fallback library ready
- **Performance Degradation**: Monitor and optimize continuously
- **Database Issues**: Regular backups, migration strategy

### 15.2 Operational Risks
- **High Traffic**: Auto-scaling, rate limiting
- **Service Outage**: Health checks, alerting, runbooks
- **Data Loss**: Regular backups, replication (future)

---

**Document Status**: ✅ Ready for Review  
**Next Steps**: Implementation planning, development kickoff
