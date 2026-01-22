# PHLX Contracts

Shared type contracts for Phylax microservices ecosystem. This repository serves as the single source of truth for data models and event schemas used across all services.

## Overview

PHLX Contracts provides:
- **Go Models** - Shared data structures for MongoDB entities (Asset, Mission, User, etc.)
- **Event Schemas** - Kafka event definitions with versioning
- **Type Safety** - Compile-time guarantees across service boundaries

## Repository Structure

```
phlx-contracts/
├── go/
│   ├── models/        # Data models (Asset, Mission, User, etc.)
│   ├── events/        # Kafka event schemas
│   ├── go.mod
│   └── go.sum
├── .gitignore
├── LICENSE
└── README.md
```

## Usage

### Go Services

Add the contracts module to your service:

```bash
go get github.com/ai-project-787/phlx-contracts/go
```

Import models and events:

```go
import (
    "github.com/ai-project-787/phlx-contracts/go/models"
    "github.com/ai-project-787/phlx-contracts/go/events"
)

// Use data models
asset := &models.Asset{
    ID:     "asset-123",
    Name:   "UAV-Alpha-1",
    Type:   models.AssetTypeUAV,
    Status: models.AssetStatusActive,
}

// Use event schemas
event := &events.AssetUpdateEvent{
    BaseEvent: events.BaseEvent{
        EventType: events.EventTypeAssetUpdate,
        Timestamp: time.Now(),
    },
    AssetID: asset.ID,
    Updates: map[string]interface{}{
        "position": asset.Position,
        "status":   asset.Status,
    },
}
```

## Versioning

- **Data Models**: Breaking changes require coordination across all services
- **Event Schemas**: Versioned via `schema_version` field (e.g., "1.0", "2.0")
- **Semantic Versioning**: This repository follows SemVer for releases

## Contributing

Changes to contracts must be:
1. **Coordinated**: Discuss with all affected service teams
2. **Backward Compatible**: Avoid breaking changes when possible
3. **Documented**: Update ownership comments and README
4. **Tested**: Verify no compilation errors in dependent services

## Ownership

Each model/event documents its owning service:

- `Asset` → dispatch-asset-service
- `Mission` → mission-command-service
- `User` → backend (auth service)
- `AssetUpdateEvent` → dispatch-asset-service
- `MissionCreatedEvent` → mission-command-service

## License

MIT License - See LICENSE file for details
