# PostgreSQL Database Rules

## Schema Design
- Use appropriate data types
- Normalize database design (3NF)
- Use indexes for frequently queried columns
- Use foreign keys for referential integrity

## Queries
- Use parameterized queries to prevent SQL injection
- Use transactions for multi-step operations
- Avoid N+1 queries, use JOINs
- Use EXPLAIN ANALYZE for query optimization

## Migrations
- Use migration tools (e.g., `golang-migrate`, `knex`)
- Never modify existing migrations
- Test migrations on staging first
- Keep migrations small and focused

## Best Practices
- Use connection pooling
- Set appropriate connection limits
- Monitor query performance
- Regular backups and maintenance

