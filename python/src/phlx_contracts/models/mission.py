"""Mission model for Phylax platform"""

from typing import Optional, List
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field
from .common import GeoLocation

# Owner: mission-command-service
# Consumers: backend, dispatch-asset-service, field-agent-app


class MissionStatus(str, Enum):
    """Mission status enumeration"""

    ACTIVE = "active"
    COMPLETED = "completed"
    ARCHIVED = "archived"


class Mission(BaseModel):
    """Mission represents an operator-managed incident with correlated events"""

    id: str = Field(alias="id")
    title: str
    description: str
    status: MissionStatus
    priority: str  # low, medium, high, critical

    # Operator management
    claimed_by_operator_id: Optional[str] = Field(default=None, alias="claimedByOperatorId")
    claimed_by_operator_name: Optional[str] = Field(default=None, alias="claimedByOperatorName")
    claimed_at: Optional[datetime] = Field(default=None, alias="claimedAt")
    completed_at: Optional[datetime] = Field(default=None, alias="completedAt")
    completed_by_operator_id: Optional[str] = Field(default=None, alias="completedByOperatorId")

    # Mission scope
    dispatch_ids: List[str] = Field(default_factory=list, alias="dispatchIds")
    asset_ids: List[str] = Field(default_factory=list, alias="assetIds")
    event_ids: List[str] = Field(default_factory=list, alias="eventIds")

    # Location (centroid of all events/assets)
    location: Optional[GeoLocation] = None

    # Metadata
    tags: List[str] = Field(default_factory=list)
    created_at: datetime = Field(alias="createdAt")
    updated_at: datetime = Field(alias="updatedAt")

    class Config:
        populate_by_name = True
        use_enum_values = True


class CreateMissionRequest(BaseModel):
    """Request to create a new mission"""

    title: str
    description: str
    priority: str
    dispatch_id: str = Field(alias="dispatchId")
    location: Optional[GeoLocation] = None

    class Config:
        populate_by_name = True


class UpdateMissionRequest(BaseModel):
    """Request to update a mission"""

    title: Optional[str] = None
    description: Optional[str] = None
    priority: Optional[str] = None
    tags: Optional[List[str]] = None

    class Config:
        populate_by_name = True


class ClaimMissionRequest(BaseModel):
    """Request to claim a mission"""

    operator_id: str = Field(alias="operatorId")
    operator_name: str = Field(alias="operatorName")

    class Config:
        populate_by_name = True


class CompleteMissionRequest(BaseModel):
    """Request to complete a mission"""

    operator_id: str = Field(alias="operatorId")

    class Config:
        populate_by_name = True


class DispatchResponseSummary(BaseModel):
    """Field agent's response to a dispatch"""

    asset_id: str = Field(alias="assetId")
    asset_name: str = Field(alias="assetName")
    accepted: bool
    response_time: datetime = Field(alias="responseTime")
    notes: Optional[str] = None

    class Config:
        populate_by_name = True


class EnrichedDispatch(BaseModel):
    """Dispatch with its responses"""

    id: str
    event_id: str = Field(alias="eventId")
    description: str
    status: str
    priority: str
    responses: List[DispatchResponseSummary]
    created_at: datetime = Field(alias="createdAt")

    class Config:
        populate_by_name = True


class EnrichedMission(BaseModel):
    """Mission with full dispatch details"""

    # Inline Mission fields
    id: str = Field(alias="id")
    title: str
    description: str
    status: MissionStatus
    priority: str
    claimed_by_operator_id: Optional[str] = Field(default=None, alias="claimedByOperatorId")
    claimed_by_operator_name: Optional[str] = Field(default=None, alias="claimedByOperatorName")
    claimed_at: Optional[datetime] = Field(default=None, alias="claimedAt")
    completed_at: Optional[datetime] = Field(default=None, alias="completedAt")
    completed_by_operator_id: Optional[str] = Field(default=None, alias="completedByOperatorId")
    dispatch_ids: List[str] = Field(default_factory=list, alias="dispatchIds")
    asset_ids: List[str] = Field(default_factory=list, alias="assetIds")
    event_ids: List[str] = Field(default_factory=list, alias="eventIds")
    location: Optional[GeoLocation] = None
    tags: List[str] = Field(default_factory=list)
    created_at: datetime = Field(alias="createdAt")
    updated_at: datetime = Field(alias="updatedAt")

    # Additional enrichment
    dispatches: List[EnrichedDispatch]

    class Config:
        populate_by_name = True
        use_enum_values = True


class AddDispatchToMissionRequest(BaseModel):
    """Request to add a dispatch to an existing mission"""

    dispatch_id: str = Field(alias="dispatchId")

    class Config:
        populate_by_name = True


class DirectCreateMissionRequest(BaseModel):
    """Request to create a mission directly without dispatch"""

    title: str
    description: str
    priority: str
    event_id: Optional[str] = Field(default=None, alias="eventId")
    location: Optional[GeoLocation] = None

    class Config:
        populate_by_name = True
