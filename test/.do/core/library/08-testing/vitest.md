# Vitest Testing Framework Rules

## Configuration
- Use `vitest.config.ts` for configuration
- Enable coverage reporting
- Configure test environment (node, jsdom)
- Set up test globals

## Test Structure
- Use `describe` and `it` blocks
- Use descriptive test names
- Follow AAA pattern
- Keep tests focused and isolated

## Best Practices
- Write fast, parallel tests
- Use `vi.mock()` for mocking
- Test both unit and integration scenarios
- Maintain high test coverage

## Performance
- Run tests in watch mode during development
- Use `test.concurrent` for independent tests
- Optimize test setup and teardown
- Profile slow tests

