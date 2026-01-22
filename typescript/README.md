# @ai-project-787/contracts (TypeScript)

TypeScript type definitions for PHLX microservices data models and event types.

## Installation

```bash
npm install @ai-project-787/contracts
# or
yarn add @ai-project-787/contracts
```

## Usage

```typescript
import { Asset, Mission, User, AssetUpdateEventData, KafkaTopics } from '@ai-project-787/contracts';

// Use data models
const asset: Asset = {
  id: '123',
  name: 'Drone-1',
  type: 'UAV',
  status: 'available',
  useCase: 'surveillance',
  latitude: 35.2401,
  longitude: 24.8093,
  lastUpdated: new Date().toISOString(),
  autoPositionEnabled: true,
};

// Use event types
const event: AssetUpdateEventData = {
  id: 'evt-123',
  type: 'asset_update',
  timestamp: new Date().toISOString(),
  source: 'dispatch-asset-service',
  assetId: asset.id,
  assetName: asset.name,
  assetType: asset.type,
  oldStatus: 'available',
  newStatus: 'dispatched',
};

// Use Kafka topics
console.log(KafkaTopics.ASSET_UPDATES); // 'asset-updates'
```

## Models

All data models are TypeScript interfaces converted from Go models:

- `Asset` - Assets (vehicles, personnel, equipment)
- `User` - User accounts and authentication
- `Mission` - Operator-managed incidents
- `Location` - Geographic locations and areas
- `Team` - Groups of assets
- `TacticalCommand` - Tactical commands and dispatches
- `MissionChat` - Mission chat messages
- `Alert` - Fire and event alerts
- `FireRisk` - Fire risk assessments
- `FireEvent` - Fire event data
- `AssetEventGroup` - Grouped events by asset
- `AuditLog` - Audit trail entries

## Events

All event types for Kafka-based event-driven architecture:

- `AssetUpdateEventData` - Asset status changes
- `MissionCreatedEventData` - New missions
- `TacticalCommandCreatedEventData` - New tactical commands
- `VideoUploadEventData` - Video uploads
- `AIAnalysisEventData` - AI analysis results
- And 20+ more event types...

## Topics

Kafka topic names:

```typescript
import { KafkaTopics } from '@ai-project-787/contracts';

KafkaTopics.ASSET_UPDATES           // 'asset-updates'
KafkaTopics.MISSION_EVENTS          // 'mission-events'
KafkaTopics.TACTICAL_COMMANDS       // 'tactical-commands'
// ... and more
```

## Type Safety

All types match the Go models exactly:

- `time.Time` → `string` (ISO 8601)
- `*T` (pointer) → `T | null`
- `omitempty` → optional field (`field?: type`)
- `int`, `float64` → `number`
- `map[string]interface{}` → `Record<string, any>`

## Versioning

This package follows semantic versioning:

- **MAJOR**: Breaking changes (removing fields, changing types)
- **MINOR**: Non-breaking additions (new optional fields)
- **PATCH**: Bugfixes (documentation, typos)

## License

MIT
