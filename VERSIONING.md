# Versioning Strategy for PHLX Contracts

This repository uses **semantic versioning** (MAJOR.MINOR.PATCH) for all contracts.

## Version Format: MAJOR.MINOR.PATCH

### MAJOR (Breaking Changes)
Increment when making **backward-incompatible** changes:
- Removing a field from a model
- Changing a field type (e.g., `string` → `number`)
- Renaming a field
- Making an optional field required
- Removing an endpoint from OpenAPI spec
- Changing request/response schema in a breaking way

**Example**: `1.2.3` → `2.0.0`

**Migration Required**: Yes - consumers must update their code

### MINOR (Non-Breaking Additions)
Increment when adding **backward-compatible** features:
- Adding a new optional field to a model
- Adding a new model
- Adding a new endpoint
- Adding a new event type
- Deprecating (but not removing) a field

**Example**: `1.2.3` → `1.3.0`

**Migration Required**: No - consumers can upgrade safely

### PATCH (Fixes & Documentation)
Increment for **backward-compatible** fixes:
- Fixing typos in documentation
- Improving code comments
- Fixing bugs that don't change API surface
- Updating README

**Example**: `1.2.3` → `1.2.4`

**Migration Required**: No

## Version Prefixes

### Go Modules
Go module versions are prefixed with `go/`:
```
go/v1.0.0
go/v1.1.0
go/v2.0.0
```

**Usage**:
```go
import "github.com/ai-project-787/phlx-contracts/go/models"
```

**Tagging**:
```bash
git tag go/v1.2.0
git push origin go/v1.2.0
```

### TypeScript Packages
TypeScript package versions are prefixed with `typescript/`:
```
typescript/v1.0.0
typescript/v1.1.0
typescript/v2.0.0
```

**Usage**:
```json
{
  "dependencies": {
    "@ai-project-787/contracts": "^1.2.0"
  }
}
```

**Tagging**:
```bash
# Update package.json version first
cd typescript
npm version minor  # or major/patch
cd ..
git tag typescript/v1.3.0
git push origin typescript/v1.3.0
```

## Examples

### Example 1: Adding a New Optional Field (MINOR)

**Before** (v1.2.0):
```go
type Asset struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

**After** (v1.3.0):
```go
type Asset struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description *string `json:"description,omitempty"` // NEW optional field
}
```

**Version**: `1.2.0` → `1.3.0` (minor bump)
**Breaking**: No

### Example 2: Changing Field Type (MAJOR)

**Before** (v1.3.0):
```go
type Asset struct {
    ID       string `json:"id"`
    Priority string `json:"priority"` // "low", "medium", "high"
}
```

**After** (v2.0.0):
```go
type Asset struct {
    ID       string `json:"id"`
    Priority int    `json:"priority"` // 1, 2, 3
}
```

**Version**: `1.3.0` → `2.0.0` (major bump)
**Breaking**: Yes - consumers must update code

### Example 3: Deprecating a Field (MINOR)

**Before** (v1.3.0):
```go
type Asset struct {
    ID       string `json:"id"`
    Location string `json:"location"` // Deprecated
}
```

**After** (v1.4.0):
```go
type Asset struct {
    ID        string     `json:"id"`
    Location  string     `json:"location,omitempty"` // DEPRECATED: Use Coordinates instead
    Latitude  *float64   `json:"latitude,omitempty"`
    Longitude *float64   `json:"longitude,omitempty"`
}
```

**Version**: `1.3.0` → `1.4.0` (minor bump)
**Breaking**: No - old field still exists

### Example 4: Removing a Deprecated Field (MAJOR)

**Before** (v1.4.0):
```go
type Asset struct {
    ID        string   `json:"id"`
    Location  string   `json:"location,omitempty"` // DEPRECATED
    Latitude  *float64 `json:"latitude,omitempty"`
    Longitude *float64 `json:"longitude,omitempty"`
}
```

**After** (v2.0.0):
```go
type Asset struct {
    ID        string   `json:"id"`
    Latitude  float64  `json:"latitude"`  // Now required
    Longitude float64  `json:"longitude"` // Now required
}
```

**Version**: `1.4.0` → `2.0.0` (major bump)
**Breaking**: Yes - removed field and made fields required

## Release Process

1. **Make changes** on a feature branch
2. **Determine version bump** (major/minor/patch)
3. **Update CHANGELOG** (if one exists)
4. **Create PR** with version info in title (e.g., "feat: add video model (minor)")
5. **Merge to main** after approval
6. **Create tags**:
   ```bash
   # For Go
   git tag go/v1.3.0
   git push origin go/v1.3.0
   
   # For TypeScript
   git tag typescript/v1.3.0
   git push origin typescript/v1.3.0
   ```
7. **GitHub Actions** automatically publishes the new version

## Synchronizing Versions

**Rule**: Go and TypeScript versions should stay in sync when possible.

If you make a breaking change to Go models, also update TypeScript types and bump both to the same major version:
```
go/v2.0.0
typescript/v2.0.0
```

If one language requires changes the other doesn't, versions can diverge temporarily, but should re-sync at the next breaking change.

## Handling Breaking Changes

1. **Announce in advance** - Give consumers notice before releasing
2. **Provide migration guide** - Document how to upgrade
3. **Support N-1 version** - Continue supporting the previous major version for at least 1 month
4. **Test thoroughly** - Breaking changes require extra testing

## Questions?

Refer to [CONTRIBUTING.md](./CONTRIBUTING.md) or open an issue.
