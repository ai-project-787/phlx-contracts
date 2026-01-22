"""User model for Phylax platform"""

from typing import Optional, Dict, Any
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field

# Owner: backend (auth service)
# Consumers: all services (authentication)


class UserRole(str, Enum):
    """User role enumeration"""

    ADMIN = "admin"
    OPERATOR = "operator"
    FIELD_AGENT = "field_agent"


class User(BaseModel):
    """User represents a user in the system with role-based access"""

    id: str = Field(alias="id")
    email: str
    password_hash: Optional[str] = Field(default=None, exclude=True)  # Never expose in JSON
    name: str
    role: UserRole
    asset_id: Optional[str] = Field(default=None, alias="assetId")
    active: bool
    created_at: datetime = Field(alias="createdAt")
    updated_at: datetime = Field(alias="updatedAt")
    last_login_at: Optional[datetime] = Field(default=None, alias="lastLoginAt")
    metadata: Optional[Dict[str, Any]] = None

    class Config:
        populate_by_name = True
        use_enum_values = True

    def has_role(self, role: UserRole) -> bool:
        """Check if the user has the specified role"""
        return self.role == role

    def is_admin(self) -> bool:
        """Check if the user is an admin"""
        return self.role == UserRole.ADMIN

    def is_operator(self) -> bool:
        """Check if the user is an operator"""
        return self.role == UserRole.OPERATOR

    def is_field_agent(self) -> bool:
        """Check if the user is a field agent"""
        return self.role == UserRole.FIELD_AGENT

    def can_access_dashboard(self) -> bool:
        """Check if user can access main dashboard"""
        return self.role in [UserRole.ADMIN, UserRole.OPERATOR]

    def can_access_field_agent_view(self) -> bool:
        """Check if user can access field agent view"""
        return self.role in [UserRole.ADMIN, UserRole.FIELD_AGENT]

    def can_access_admin_settings(self) -> bool:
        """Check if user can access admin settings"""
        return self.role == UserRole.ADMIN


class UserSession(BaseModel):
    """User session model"""

    id: str = Field(alias="id")
    user_id: str = Field(alias="userId")
    token: str
    expires_at: datetime = Field(alias="expiresAt")
    created_at: datetime = Field(alias="createdAt")
    ip_address: Optional[str] = Field(default=None, alias="ipAddress")
    user_agent: Optional[str] = Field(default=None, alias="userAgent")

    class Config:
        populate_by_name = True
