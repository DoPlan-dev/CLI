# Go Language Rules

## Code Style
- Use `gofmt` for formatting
- Follow Go naming conventions (exported: PascalCase, unexported: camelCase)
- Use `golint` and `golangci-lint` for linting
- Maximum line length: 120 characters

## Error Handling
- Always check errors explicitly
- Use `fmt.Errorf` with `%w` for error wrapping
- Return errors, don't ignore them
- Use custom error types for domain errors

## Testing
- Write tests in `*_test.go` files
- Use table-driven tests for multiple cases
- Aim for >80% test coverage
- Use `testify` for assertions

## Package Structure
- Keep packages focused and cohesive
- Use `internal/` for private packages
- Use `pkg/` for public packages
- Keep main package minimal

## Performance
- Use `sync.Pool` for frequently allocated objects
- Prefer `strings.Builder` for string concatenation
- Use `context.Context` for cancellation
- Profile before optimizing

