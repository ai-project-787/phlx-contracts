# Phylax Contracts - Python Package

Shared type contracts for Phylax microservices. This package provides Pydantic models for all data types used across the Phylax platform.

## Installation

### From GitHub (recommended for development)

```bash
pip install git+https://github.com/ai-project-787/phlx-contracts.git@main#subdirectory=python
```

### From local source

```bash
cd phlx-contracts/python
pip install -e .
```

## Usage

```python
from phlx_contracts.models import Asset, Mission, User, TacticalCommand

# Parse API response into typed model
asset_data = {"id": "123", "name": "UAV-1", "type": "uav", ...}
asset = Asset(**asset_data)

# Type-safe access
print(asset.name)  # IDE autocomplete works!
print(asset.battery_level)  # Optional[int], type-safe

# Runtime validation
try:
    invalid_asset = Asset(**{"invalid": "data"})
except ValidationError as e:
    print(f"Validation failed: {e}")
```

## Features

- **Type Safety**: Full type hints for IDE autocomplete and static analysis
- **Runtime Validation**: Pydantic validates all data at runtime
- **camelCase & snake_case Support**: Both naming conventions work (`populate_by_name=True`)
- **Enum Support**: Type-safe enums for status, roles, priorities, etc.
- **Optional Fields**: Proper Optional typing for nullable fields

## Models

### Data Models (12 total)

- `Asset` - Assets (UAVs, vehicles, personnel)
- `Mission` - Operator-managed incidents
- `User` - User accounts and authentication
- `Location` - Geographic locations and areas
- `Team` - Asset teams and compositions
- `TacticalCommand` - Operational commands
- `MissionChatMessage` - Mission chat messages
- `Alert` - System alerts
- `FireRisk` - Fire risk assessment
- `FireEvent` - Fire events and monitoring
- `AssetEventGroup` - Grouped events by asset
- `AuditLog` - Audit trail entries

### Common Types

- `GeoPoint` - GeoJSON point (MongoDB geospatial)
- `GeoLocation` - Geographic coordinates
- `Coordinate` - GPS coordinate (lat/lon)

## Versioning

This package follows semantic versioning and is synchronized with the Go and TypeScript packages:

- **Go**: `github.com/ai-project-787/phlx-contracts/go@v1.0.0`
- **TypeScript**: `@ai-project-787/contracts@1.0.0`
- **Python**: `phlx-contracts==1.0.0`

See `VERSIONING.md` in the root repository for details.

## Development

### Install development dependencies

```bash
pip install -e .[dev]
```

### Run tests

```bash
pytest
```

### Type checking

```bash
mypy src/
```

### Code formatting

```bash
black src/
isort src/
```

## License

MIT License - See LICENSE file in the root repository.
