"""Data models for Phylax platform"""

from .common import GeoPoint, GeoLocation, Coordinate
from .asset import Asset
from .mission import (
    Mission,
    MissionStatus,
    CreateMissionRequest,
    UpdateMissionRequest,
    ClaimMissionRequest,
    CompleteMissionRequest,
    DispatchResponseSummary,
    EnrichedDispatch,
    EnrichedMission,
    AddDispatchToMissionRequest,
    DirectCreateMissionRequest,
)
from .user import User
from .location import Location, LocType, GetLocationsRequest
from .team import Team
from .tactical_command import TacticalCommand, TacticalCommandType
from .mission_chat import MissionChatMessage
from .alert import Alert
from .fire_risk import FireRisk
from .fire_event import FireEvent
from .asset_event_group import AssetEventGroup
from .audit_log import AuditLog

__all__ = [
    # Common types
    "GeoPoint",
    "GeoLocation",
    "Coordinate",
    # Models
    "Asset",
    "Mission",
    "MissionStatus",
    "CreateMissionRequest",
    "UpdateMissionRequest",
    "ClaimMissionRequest",
    "CompleteMissionRequest",
    "DispatchResponseSummary",
    "EnrichedDispatch",
    "EnrichedMission",
    "AddDispatchToMissionRequest",
    "DirectCreateMissionRequest",
    "User",
    "Location",
    "LocType",
    "GetLocationsRequest",
    "Team",
    "TacticalCommand",
    "TacticalCommandType",
    "MissionChatMessage",
    "Alert",
    "FireRisk",
    "FireEvent",
    "AssetEventGroup",
    "AuditLog",
]
