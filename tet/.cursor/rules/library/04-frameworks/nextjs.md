# Next.js Framework Rules

## Project Structure
- Use App Router (app directory) for new projects
- Keep pages/components organized by feature
- Use `lib/` for utilities
- Use `components/` for reusable UI components

## Data Fetching
- Use Server Components by default
- Use `fetch` with caching for data fetching
- Implement proper error boundaries
- Handle loading states

## Performance
- Use `next/image` for images
- Implement code splitting
- Use dynamic imports for heavy components
- Optimize fonts with `next/font`

## API Routes
- Keep API routes in `app/api/` directory
- Validate request data
- Handle errors gracefully
- Return proper HTTP status codes

## Best Practices
- Use TypeScript for type safety
- Implement proper SEO with metadata
- Use middleware for authentication
- Follow Next.js conventions

