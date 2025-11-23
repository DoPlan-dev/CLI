# Authentication Security Rules

## Password Security
- Never store plain text passwords
- Use bcrypt or Argon2 for hashing
- Enforce strong password policies
- Implement password reset securely

## Session Management
- Use secure, HTTP-only cookies
- Implement session expiration
- Regenerate session IDs on login
- Use CSRF tokens for state-changing operations

## OAuth/OpenID Connect
- Validate ID tokens properly
- Use PKCE for public clients
- Store tokens securely
- Implement token refresh

## Best Practices
- Use HTTPS for all authentication
- Implement rate limiting on auth endpoints
- Log authentication events
- Use multi-factor authentication when possible

