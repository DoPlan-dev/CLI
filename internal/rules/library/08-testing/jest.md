# Jest Testing Framework Rules

## Test Structure
- Use `describe` blocks to group related tests
- Use descriptive test names
- Follow AAA pattern: Arrange, Act, Assert
- Keep tests independent and isolated

## Best Practices
- Write tests before or alongside code (TDD)
- Test edge cases and error conditions
- Use mocks for external dependencies
- Aim for >80% code coverage

## Mocking
- Mock external API calls
- Mock file system operations
- Use `jest.fn()` for function mocks
- Reset mocks between tests

## Assertions
- Use descriptive assertion messages
- Test one thing per test
- Use appropriate matchers
- Verify both success and failure cases

