"""Team model for Phylax platform"""

from typing import Optional, List, Dict, Any
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field
from .common import GeoLocation

# Owner: backend (team management)
# Consumers: dispatch-asset-service, mission-command-service


class TeamStatus(str, Enum):
    """Team operational status enumeration"""

    ACTIVE = "active"
    INACTIVE = "inactive"
    DEPLOYED = "deployed"


class Team(BaseModel):
    """Team represents a group of assets working together"""

    id: str = Field(alias="id")
    name: str
    description: Optional[str] = None
    status: TeamStatus
    color: Optional[str] = None  # Hex color for UI visualization

    # Team Composition
    asset_ids: List[str] = Field(default_factory=list, alias="assetIds")
    leader_id: Optional[str] = Field(default=None, alias="leaderId")

    # Team Capabilities
    capabilities: List[str] = Field(default_factory=list)

    # Location (optional - could be computed from assets)
    base_location: Optional[GeoLocation] = Field(default=None, alias="baseLocation")

    # Audit Trail
    created_by: str = Field(alias="createdBy")
    created_by_name: str = Field(alias="createdByName")
    created_at: datetime = Field(alias="createdAt")
    updated_at: datetime = Field(alias="updatedAt")

    # Metadata
    metadata: Optional[Dict[str, Any]] = None

    class Config:
        populate_by_name = True
        use_enum_values = True
