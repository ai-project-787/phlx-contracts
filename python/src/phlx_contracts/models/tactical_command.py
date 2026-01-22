"""Tactical command model for Phylax platform"""

from typing import Optional, List, Dict, Any
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field

# Owner: mission-command-service
# Consumers: backend, field-agent-app, ai-analysis-service


class TacticalCommandStatus(str, Enum):
    """Tactical command status enumeration"""

    PENDING_APPROVAL = "pending_approval"
    PENDING = "pending"
    ACCEPTED = "accepted"
    REJECTED = "rejected"
    IN_PROGRESS = "in_progress"
    COMPLETED = "completed"
    CANCELLED = "cancelled"


class TacticalCommandCategory(str, Enum):
    """Tactical command category enumeration"""

    MOVEMENT = "movement"
    SECURITY = "security"
    SURVEILLANCE = "surveillance"
    DISPATCH = "dispatch"
    COMMUNICATION = "communication"
    MEDICAL = "medical"
    EVACUATION = "evacuation"
    SUPPORT = "support"
    INVESTIGATION = "investigation"
    OTHER = "other"


class TacticalCommandPriority(str, Enum):
    """Tactical command priority enumeration"""

    ROUTINE = "routine"
    PRIORITY = "priority"
    IMMEDIATE = "immediate"
    FLASH = "flash"


class CommandTarget(BaseModel):
    """Command target (asset or team)"""

    target_type: str = Field(alias="target_type")  # "asset" or "team"
    target_id: str = Field(alias="target_id")
    target_name: str = Field(alias="target_name")

    class Config:
        populate_by_name = True


class CommandResponse(BaseModel):
    """Target's response to a tactical command"""

    target_id: str = Field(alias="target_id")
    target_type: str = Field(alias="target_type")
    target_name: str = Field(alias="target_name")
    decision: str  # "accepted" or "rejected"
    notes: Optional[str] = None
    responded_by: str = Field(alias="responded_by")
    responded_by_name: str = Field(alias="responded_by_name")
    responded_at: datetime = Field(alias="responded_at")

    class Config:
        populate_by_name = True


class CommandStatusUpdate(BaseModel):
    """Status change in command lifecycle"""

    status: TacticalCommandStatus
    changed_by: str = Field(alias="changed_by")
    changed_by_name: str = Field(alias="changed_by_name")
    timestamp: datetime
    notes: Optional[str] = None

    class Config:
        populate_by_name = True
        use_enum_values = True


class TacticalGeoLocation(BaseModel):
    """Geographic coordinates for command destinations"""

    lat: float
    lng: float
    name: Optional[str] = None
    description: Optional[str] = None

    class Config:
        populate_by_name = True


class TacticalGeoArea(BaseModel):
    """Area of operation (circle, polygon, or route)"""

    type: str  # "circle", "polygon", "route"
    center: Optional[TacticalGeoLocation] = None
    radius: Optional[float] = None  # meters (for circle)
    coordinates: Optional[List[TacticalGeoLocation]] = None
    name: Optional[str] = None

    class Config:
        populate_by_name = True


class TacticalCommand(BaseModel):
    """Tactical command model"""

    id: str = Field(alias="id")

    # Mission Context
    mission_id: str = Field(alias="mission_id")
    mission_title: str = Field(alias="mission_title")
    situation_summary: Optional[str] = Field(default=None, alias="situation_summary")

    # Command Definition
    title: str
    description: str
    category: TacticalCommandCategory

    # Multi-Target Assignment
    targets: List[CommandTarget]

    # Location & Navigation
    destination: Optional[TacticalGeoLocation] = None
    waypoints: Optional[List[TacticalGeoLocation]] = None
    area_of_operation: Optional[TacticalGeoArea] = Field(default=None, alias="area_of_operation")

    # Mission Parameters
    objective: Optional[str] = None
    priority: TacticalCommandPriority

    # Status & Responses
    status: TacticalCommandStatus
    responses: List[CommandResponse] = Field(default_factory=list)
    status_history: List[CommandStatusUpdate] = Field(default_factory=list, alias="status_history")

    # Metadata
    source: str  # "ai" or "operator"
    created_by: str = Field(alias="created_by")
    created_by_name: str = Field(alias="created_by_name")
    created_at: datetime = Field(alias="created_at")
    updated_at: datetime = Field(alias="updated_at")

    metadata: Optional[Dict[str, Any]] = None

    class Config:
        populate_by_name = True
        use_enum_values = True


class CreateTacticalCommandRequest(BaseModel):
    """Request to create a new tactical command"""

    mission_id: str = Field(alias="mission_id")
    title: str
    description: str
    category: TacticalCommandCategory
    targets: Optional[List[CommandTarget]] = None
    target_name: Optional[str] = Field(default=None, alias="target_name")  # Comma-separated from Langflow
    destination: Optional[TacticalGeoLocation] = None
    waypoints: Optional[List[TacticalGeoLocation]] = None
    area_of_operation: Optional[TacticalGeoArea] = Field(default=None, alias="area_of_operation")
    objective: Optional[str] = None
    priority: TacticalCommandPriority
    situation_summary: Optional[str] = Field(default=None, alias="situation_summary")
    source: str = "operator"  # "ai" or "operator"
    metadata: Optional[Dict[str, Any]] = None

    class Config:
        populate_by_name = True
        use_enum_values = True


class RespondToTacticalCommandRequest(BaseModel):
    """Target's response to a command"""

    target_id: str = Field(alias="target_id")
    target_type: str = Field(alias="target_type")
    decision: str  # "accepted" or "rejected"
    notes: Optional[str] = None

    class Config:
        populate_by_name = True


class UpdateTacticalCommandStatusRequest(BaseModel):
    """Status update request"""

    status: TacticalCommandStatus
    notes: Optional[str] = None

    class Config:
        populate_by_name = True
        use_enum_values = True


class TacticalCommandFilter(BaseModel):
    """Query filters for listing commands"""

    mission_id: Optional[str] = Field(default=None, alias="mission_id")
    status: Optional[TacticalCommandStatus] = None
    target_id: Optional[str] = Field(default=None, alias="target_id")
    target_type: Optional[str] = Field(default=None, alias="target_type")
    category: Optional[TacticalCommandCategory] = None
    priority: Optional[TacticalCommandPriority] = None
    source: Optional[str] = None

    class Config:
        populate_by_name = True
        use_enum_values = True


# Define type alias for simpler imports
TacticalCommandType = TacticalCommandCategory
