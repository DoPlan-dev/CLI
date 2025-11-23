# TypeScript Language Rules

## Type Safety
- Enable strict mode in `tsconfig.json`
- Avoid `any` type, use `unknown` if needed
- Use type inference where possible
- Define interfaces for object shapes

## Code Style
- Use ESLint with TypeScript plugin
- Use Prettier for formatting
- Maximum line length: 100 characters
- Use semicolons consistently

## Module System
- Use ES modules (`import`/`export`)
- Avoid default exports for better refactoring
- Use barrel exports (`index.ts`) for public APIs
- Organize imports: external, internal, relative

## Best Practices
- Use `const` by default, `let` when needed
- Prefer arrow functions for callbacks
- Use async/await over promises
- Handle errors explicitly

