"""Asset event group model for Phylax platform"""

from typing import Optional, List, Dict, Any, Tuple
from pydantic import BaseModel, Field

# Owner: backend (event grouping)
# Consumers: frontend, event-correlation-service


class GeoJSONPoint(BaseModel):
    """GeoJSON point representation"""

    type: str = "Point"
    coordinates: Tuple[float, float]  # [longitude, latitude]

    class Config:
        populate_by_name = True


class Event(BaseModel):
    """Single event within a group"""

    id: str
    type: str
    timestamp: str
    location: Optional[GeoJSONPoint] = None
    severity: str
    description: str
    metadata: Optional[Dict[str, Any]] = None

    class Config:
        populate_by_name = True


class AssetEventGroup(BaseModel):
    """Grouped events by asset"""

    asset_id: str = Field(alias="assetId")
    asset_name: str = Field(alias="assetName")
    asset_type: str = Field(alias="assetType")  # "camera" or "fire_detector"
    event_count: int = Field(alias="eventCount")
    latest_event: Event = Field(alias="latestEvent")
    event_ids: List[str] = Field(alias="eventIds")
    events: Optional[List[Event]] = None  # Full event data for display

    class Config:
        populate_by_name = True


class GroupedEventsResponse(BaseModel):
    """Response for grouped events API"""

    groups: List[AssetEventGroup]
    count: int

    class Config:
        populate_by_name = True
