"""Fire risk model for Phylax platform"""

from typing import List, Dict, Any
from datetime import datetime
from pydantic import BaseModel, Field
from .common import GeoPoint

# Owner: backend (fire risk assessment)
# Consumers: frontend, external-data-service


class MonitoredLocation(BaseModel):
    """Monitored location model"""

    id: str = Field(alias="id")
    name: str
    type: str  # "asset", "facility", etc.
    location: GeoPoint
    status: str

    class Config:
        populate_by_name = True


class FireData(BaseModel):
    """Fire data model"""

    id: str = Field(alias="id")
    source: str
    source_type: str = Field(alias="source_type")
    timestamp: datetime
    location: GeoPoint
    data: Dict[str, Any]
    tags: List[str] = Field(default_factory=list)
    source_metadata: Dict[str, Any] = Field(default_factory=dict, alias="source_metadata")

    class Config:
        populate_by_name = True


class BoundingBox(BaseModel):
    """Geographic bounding box"""

    west: float
    south: float
    east: float
    north: float


# Alias for backwards compatibility
FireRisk = MonitoredLocation
