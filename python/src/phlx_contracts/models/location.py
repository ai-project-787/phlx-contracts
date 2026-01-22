"""Location model for Phylax platform"""

from typing import Optional, List
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field
from .common import Coordinate

# Owner: location-navigation-service
# Consumers: backend, dispatch-asset-service, mission-command-service


class LocType(str, Enum):
    """Location type enumeration"""

    PERIMETER = "perimeter"
    PATROL_ZONE = "patrol_zone"
    CHECKPOINT = "checkpoint"


class Area(BaseModel):
    """Area represents a named zone within a location with polygon boundary"""

    id: str
    name: str
    description: Optional[str] = None
    boundary: List[Coordinate]  # Polygon points
    fill_color: Optional[str] = Field(default=None, alias="fillColor")
    border_color: Optional[str] = Field(default=None, alias="borderColor")
    opacity: Optional[float] = None  # 0.0 to 1.0
    type: Optional[str] = None  # perimeter, patrol_zone, checkpoint, etc.
    priority: Optional[str] = None  # high, medium, low
    active: bool
    created_at: datetime = Field(alias="createdAt")
    updated_at: datetime = Field(alias="updatedAt")

    class Config:
        populate_by_name = True


class Location(BaseModel):
    """Location represents a geographic location with center coordinates and multiple areas"""

    id: str = Field(alias="id")
    name: str
    description: Optional[str] = None

    # Center coordinates
    latitude: float
    longitude: float

    # Areas within this location
    areas: List[Area] = Field(default_factory=list)

    # Visual settings
    color: Optional[str] = None
    icon: Optional[str] = None

    # Metadata
    use_case: Optional[str] = Field(default=None, alias="useCase")
    tags: List[str] = Field(default_factory=list)

    # Status
    active: bool

    # Audit
    created_by: Optional[str] = Field(default=None, alias="createdBy")
    created_at: datetime = Field(alias="createdAt")
    updated_by: Optional[str] = Field(default=None, alias="updatedBy")
    updated_at: datetime = Field(alias="updatedAt")

    class Config:
        populate_by_name = True


class CreateLocationRequest(BaseModel):
    """Request to create a new location"""

    name: str
    description: Optional[str] = None
    latitude: float
    longitude: float
    color: Optional[str] = None
    icon: Optional[str] = None
    use_case: Optional[str] = Field(default=None, alias="useCase")
    tags: Optional[List[str]] = None
    active: Optional[bool] = None

    class Config:
        populate_by_name = True


class GetLocationsRequest(BaseModel):
    """Request to get locations"""

    use_case: Optional[str] = Field(default=None, alias="useCase")
    active: Optional[bool] = None

    class Config:
        populate_by_name = True
