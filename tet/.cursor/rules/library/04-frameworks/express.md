# Express.js Framework Rules

## Project Structure
- Organize routes in separate files
- Use middleware for cross-cutting concerns
- Keep controllers thin
- Separate business logic from routes

## Error Handling
- Use centralized error handler
- Return consistent error responses
- Log errors appropriately
- Use HTTP status codes correctly

## Security
- Use `helmet` for security headers
- Validate and sanitize input
- Use `express-rate-limit` for rate limiting
- Implement proper authentication

## Best Practices
- Use async/await for async operations
- Use environment variables for configuration
- Implement proper logging
- Write integration tests

