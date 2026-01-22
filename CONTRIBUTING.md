# Contributing to PHLX Contracts

Thank you for contributing to the PHLX Contracts repository! This document provides guidelines for adding or modifying shared data contracts.

## Development Workflow

1. **Always work on a feature branch**
   ```bash
   git checkout -b feat/add-new-model
   ```

2. **Make your changes** in the appropriate directory:
   - Go models: `go/models/` or `go/events/`
   - TypeScript types: `typescript/src/models/` or `typescript/src/events/`
   - OpenAPI specs: `openapi/`

3. **Test your changes locally**:
   ```bash
   # Test Go
   cd go && go build ./... && go test ./...
   
   # Test TypeScript
   cd typescript && npm run build && npm run lint
   
   # Validate OpenAPI (requires swagger-cli)
   npx @apidevtools/swagger-cli validate openapi/*.yaml
   ```

4. **Create a pull request** to the `main` branch
   - PRs require approval before merging
   - All validation checks must pass

## Adding New Models

### Go Models

1. Create a new file in `go/models/` (e.g., `video.go`)
2. Define your model with proper struct tags:
   ```go
   type Video struct {
       ID        string    `json:"id" bson:"_id,omitempty"`
       Name      string    `json:"name" bson:"name"`
       CreatedAt time.Time `json:"createdAt" bson:"created_at"`
   }
   ```
3. Add ownership comment at the top:
   ```go
   // Owner: video-processing-service
   // Consumers: backend, frontend
   ```

### TypeScript Types

1. Create a new file in `typescript/src/models/` (e.g., `Video.ts`)
2. Convert Go types to TypeScript:
   - `time.Time` → `string` (ISO 8601)
   - `*T` (pointer) → `T | null`
   - `omitempty` → optional field (`field?: type`)
   - `int`, `float64` → `number`
   - `map[string]interface{}` → `Record<string, any>`
3. Add to `typescript/src/models/index.ts`:
   ```typescript
   export * from './Video';
   ```

### OpenAPI Specs

1. Update the relevant service spec in `openapi/`
2. Add paths for new endpoints
3. Add schemas for new models in `components/schemas`
4. Validate with `swagger-cli validate openapi/<service>.yaml`

## Making Changes

### Breaking vs Non-Breaking Changes

**Breaking changes** (require major version bump):
- Removing a field
- Changing a field type
- Renaming a field
- Making an optional field required

**Non-breaking changes** (minor version bump):
- Adding a new optional field
- Adding a new model
- Adding a new endpoint
- Deprecating (but not removing) a field

**Patch changes**:
- Documentation updates
- Comment improvements
- Fixing typos

### Versioning

See [VERSIONING.md](./VERSIONING.md) for full details.

When making breaking changes:
1. Discuss in PR review
2. Update version in `package.json` (TypeScript) and create a new tag
3. Document migration path in PR description

## Code Style

### Go
- Follow standard Go conventions
- Run `go fmt` before committing
- Use descriptive struct field names

### TypeScript
- Use PascalCase for interfaces
- Use camelCase for properties
- Add JSDoc comments for complex types

### OpenAPI
- Use clear, descriptive summaries
- Provide examples where helpful
- Group related endpoints with tags

## Testing

Before submitting a PR:
- ✅ Go builds successfully (`go build ./...`)
- ✅ TypeScript builds successfully (`npm run build`)
- ✅ OpenAPI specs validate (`swagger-cli validate`)
- ✅ All tests pass (when tests exist)

## Getting Help

- Open an issue for questions
- Tag relevant team members in PRs
- Refer to existing models for examples
