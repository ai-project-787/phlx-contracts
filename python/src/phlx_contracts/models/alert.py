"""Alert model for Phylax platform"""

from typing import Optional
from datetime import datetime
from pydantic import BaseModel, Field

# Owner: backend (alert service)
# Consumers: frontend, event-correlation-service


class Alert(BaseModel):
    """Alert represents a system alert"""

    id: str = Field(alias="id")
    type: str  # "fire_risk", "asset_danger", "weather_warning"
    severity: str  # "low", "medium", "high", "critical"
    location_id: str = Field(alias="location_id")
    location_name: str = Field(alias="location_name")
    message: str
    fire_event_id: Optional[str] = Field(default=None, alias="fire_event_id")

    created_at: datetime = Field(alias="created_at")
    updated_at: datetime = Field(alias="updated_at")

    status: str  # "active", "acknowledged", "resolved"
    acknowledged_at: Optional[datetime] = Field(default=None, alias="acknowledged_at")
    acknowledged_by: Optional[str] = Field(default=None, alias="acknowledged_by")

    class Config:
        populate_by_name = True
