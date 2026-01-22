"""Mission chat model for Phylax platform"""

from typing import List
from datetime import datetime
from pydantic import BaseModel, Field

# Owner: backend (mission chat)
# Consumers: frontend, field-agent-app


class MissionChatMessage(BaseModel):
    """Mission chat message model"""

    id: str = Field(alias="id")
    mission_id: str = Field(alias="missionId")
    sender_id: str = Field(alias="senderId")
    sender_name: str = Field(alias="senderName")
    sender_role: str = Field(alias="senderRole")  # "operator" | "field_agent"
    content: str
    timestamp: datetime
    created_at: datetime = Field(alias="createdAt")

    class Config:
        populate_by_name = True


class TypingUser(BaseModel):
    """User currently typing"""

    user_id: str = Field(alias="userId")
    user_name: str = Field(alias="userName")

    class Config:
        populate_by_name = True


class TypingStatus(BaseModel):
    """Typing status in mission chat"""

    mission_id: str = Field(alias="missionId")
    typing_users: List[TypingUser] = Field(default_factory=list, alias="typingUsers")

    class Config:
        populate_by_name = True


class SendMissionChatMessageRequest(BaseModel):
    """Request to send a chat message"""

    content: str

    class Config:
        populate_by_name = True


class UpdateTypingStatusRequest(BaseModel):
    """Request to update typing status"""

    is_typing: bool = Field(alias="isTyping")

    class Config:
        populate_by_name = True


class MissionChatResponse(BaseModel):
    """Paginated chat history response"""

    messages: List[MissionChatMessage]
    total_count: int = Field(alias="totalCount")
    has_more: bool = Field(alias="hasMore")

    class Config:
        populate_by_name = True
