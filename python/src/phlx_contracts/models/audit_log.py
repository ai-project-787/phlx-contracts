"""Audit log model for Phylax platform"""

from typing import Optional, Dict, Any
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field

# Owner: backend (audit logging)
# Consumers: all services (audit trail)


class AuditActionType(str, Enum):
    """Audit action type enumeration"""

    # Mission actions
    MISSION_CREATED = "mission_created"
    MISSION_CLAIMED = "mission_claimed"
    MISSION_COMPLETED = "mission_completed"
    MISSION_ARCHIVED = "mission_archived"
    MISSION_DELETED = "mission_deleted"

    # Asset actions
    ASSET_STATUS_CHANGED = "asset_status_changed"

    # Event actions
    EVENT_CORRELATED = "event_correlated"
    EVENT_SUGGESTED = "event_suggested"
    EVENT_APPROVED = "event_approved"
    EVENT_REJECTED = "event_rejected"

    # Operator actions
    OPERATOR_ORDER = "operator_order"
    OPERATOR_NOTE = "operator_note"

    # Operational command actions
    COMMAND_RECEIVED = "command_received"
    COMMAND_ACCEPTED = "command_accepted"
    COMMAND_DECLINED = "command_declined"
    COMMAND_STARTED = "command_started"
    COMMAND_COMPLETED = "command_completed"


class AuditLog(BaseModel):
    """Audit log entry"""

    id: str = Field(alias="id")
    mission_id: str = Field(alias="missionId")
    timestamp: datetime

    # Action details
    action_type: AuditActionType = Field(alias="actionType")
    actor_type: str = Field(alias="actorType")  # "operator", "asset", "ai_agent", "system"
    actor_id: str = Field(alias="actorId")
    actor_name: Optional[str] = Field(default=None, alias="actorName")

    # Target of action
    target_type: Optional[str] = Field(default=None, alias="targetType")  # "asset", "dispatch", "event"
    target_id: Optional[str] = Field(default=None, alias="targetId")
    target_name: Optional[str] = Field(default=None, alias="targetName")

    # Details
    action: str  # Human-readable action
    details: Dict[str, Any]
    created_at: datetime = Field(alias="createdAt")

    class Config:
        populate_by_name = True
        use_enum_values = True
