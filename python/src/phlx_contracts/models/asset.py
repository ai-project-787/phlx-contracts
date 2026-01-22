"""Asset model for Phylax platform"""

from typing import Optional, List, Dict, Any
from datetime import datetime
from pydantic import BaseModel, Field

# Owner: dispatch-asset-service
# Consumers: backend, mission-command-service, location-navigation-service, field-agent-app

# Asset status constants
ASSET_STATUS_AVAILABLE = "available"
ASSET_STATUS_DISPATCHED = "dispatched"
ASSET_STATUS_RETURNING = "returning"
ASSET_STATUS_OFFLINE = "offline"


class Asset(BaseModel):
    """Asset represents an asset in the system"""

    id: str = Field(alias="id")
    name: str
    type: str
    status: str
    use_case: str = Field(alias="useCase")
    team_id: Optional[str] = Field(default=None, alias="teamId")
    assigned_area_ids: Optional[List[str]] = Field(default=None, alias="assignedAreaIds")
    latitude: float
    longitude: float
    altitude: Optional[float] = None
    battery_level: Optional[int] = Field(default=None, alias="batteryLevel")
    members: Optional[int] = None
    vehicle: Optional[str] = None
    pulse_rate: Optional[int] = Field(default=None, alias="pulseRate")
    oxygen_level: Optional[int] = Field(default=None, alias="oxygenLevel")
    location: Optional[str] = None
    dispatch_time: Optional[datetime] = Field(default=None, alias="dispatchTime")
    estimated_arrival: Optional[datetime] = Field(default=None, alias="estimatedArrival")
    last_updated: datetime = Field(alias="lastUpdated")
    last_vital_update: Optional[datetime] = Field(default=None, alias="lastVitalUpdate")
    video_src: Optional[str] = Field(default=None, alias="videoSrc")
    metadata: Optional[Dict[str, Any]] = None
    auto_position_enabled: bool = Field(alias="autoPositionEnabled")

    class Config:
        populate_by_name = True  # Allow both snake_case and camelCase
