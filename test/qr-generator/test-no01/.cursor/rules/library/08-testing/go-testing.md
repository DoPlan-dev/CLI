# Go Testing Rules

## Test Structure
- Write tests in `*_test.go` files
- Use table-driven tests for multiple cases
- Use subtests with `t.Run()`
- Keep tests in same package (white-box testing)

## Best Practices
- Use `testify` for assertions
- Test both success and error cases
- Use test fixtures for complex data
- Aim for >80% coverage

## Test Utilities
- Use `httptest` for HTTP testing
- Use `testify/mock` for mocking
- Use `testify/assert` for assertions
- Use `testify/require` for fatal assertions

## Performance Testing
- Use `testing.B` for benchmarks
- Run benchmarks with `go test -bench`
- Compare benchmark results
- Profile with `go test -cpuprofile`

