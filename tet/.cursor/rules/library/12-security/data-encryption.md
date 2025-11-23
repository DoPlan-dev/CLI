# Data Encryption Rules

## Encryption at Rest
- Encrypt sensitive data in databases
- Use strong encryption algorithms (AES-256)
- Protect encryption keys securely
- Rotate keys periodically

## Encryption in Transit
- Use TLS 1.2+ for all connections
- Enforce HTTPS for web applications
- Use certificate pinning for mobile apps
- Validate SSL certificates

## Key Management
- Never hardcode encryption keys
- Use key management services (AWS KMS, HashiCorp Vault)
- Rotate keys regularly
- Implement key versioning

## Best Practices
- Encrypt PII and sensitive data
- Use environment variables for secrets
- Never commit secrets to version control
- Audit encryption usage

