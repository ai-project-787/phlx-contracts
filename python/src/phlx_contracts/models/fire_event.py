"""Fire event model for Phylax platform"""

from typing import Optional, List
from datetime import datetime
from pydantic import BaseModel, Field

# Owner: backend (fire event processing)
# Consumers: frontend, external-data-service, event-correlation-service


class FWIInfo(BaseModel):
    """Fire Weather Index information"""

    value: float  # FWI numeric value (0-100)
    category: str  # "very_low", "low", "moderate", "high", "very_high", "extreme"
    rating: int  # 1-10 numeric rating

    class Config:
        populate_by_name = True


class FireDetail(BaseModel):
    """Fire detail information"""

    fire_id: str = Field(alias="fire_id")
    source: str  # "copernicus_effis", "nasa_firms"
    satellite_source: Optional[str] = Field(default=None, alias="satellite_source")
    distance: float  # km
    in_fire: bool = Field(alias="in_fire")
    intensity: Optional[float] = None
    confidence: Optional[str] = None

    class Config:
        populate_by_name = True


class ScoreFactors(BaseModel):
    """Risk score breakdown"""

    distance_score: float = Field(alias="distance_score")  # 0-100
    intensity_score: float = Field(alias="intensity_score")  # 0-100
    confidence_score: float = Field(alias="confidence_score")  # 0-100
    fwi_score: float = Field(alias="fwi_score")  # 0-100

    class Config:
        populate_by_name = True


class FireEvent(BaseModel):
    """Fire event model"""

    id: str = Field(alias="id")
    location_id: str = Field(alias="location_id")
    location_name: str = Field(alias="location_name")
    location_type: str = Field(alias="location_type")  # "asset", "facility"

    event_type: str = Field(alias="event_type")  # "detected", "updated", "cleared"
    risk_level: str = Field(alias="risk_level")  # "none", "low", "medium", "high", "critical"
    risk_score: float = Field(alias="risk_score")  # 0-100

    fires: List[FireDetail] = Field(default_factory=list)
    fwi: Optional[FWIInfo] = None

    score_factors: Optional[ScoreFactors] = Field(default=None, alias="score_factors")

    created_at: datetime = Field(alias="created_at")
    updated_at: datetime = Field(alias="updated_at")

    status: str  # "active", "cleared", "acknowledged"

    class Config:
        populate_by_name = True
