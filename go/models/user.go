package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Owner: backend (auth service)
// Consumers: all services (authentication)

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleAdmin     UserRole = "admin"      // Full system access
	RoleOperator  UserRole = "operator"   // Main dashboard + upload video
	RoleFieldAgent UserRole = "field_agent" // Field agent dispatch view
)

// User represents a user in the system with role-based access
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password_hash" json:"-"` // Never expose in JSON
	Name         string             `bson:"name" json:"name"`
	Role         UserRole           `bson:"role" json:"role"`
	AssetID      string             `bson:"asset_id,omitempty" json:"assetId,omitempty"` // For field agents
	Active       bool               `bson:"active" json:"active"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updatedAt"`
	LastLoginAt  *time.Time         `bson:"last_login_at,omitempty" json:"lastLoginAt,omitempty"`
	Metadata     map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"`
}

// UserSession represents an active user session
type UserSession struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expiresAt"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	IPAddress string             `bson:"ip_address,omitempty" json:"ipAddress,omitempty"`
	UserAgent string             `bson:"user_agent,omitempty" json:"userAgent,omitempty"`
}

// HasRole checks if the user has the specified role
func (u *User) HasRole(role UserRole) bool {
	return u.Role == role
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// IsOperator checks if the user is an operator
func (u *User) IsOperator() bool {
	return u.Role == RoleOperator
}

// IsFieldAgent checks if the user is a field agent
func (u *User) IsFieldAgent() bool {
	return u.Role == RoleFieldAgent
}

// CanAccessDashboard checks if user can access main dashboard
func (u *User) CanAccessDashboard() bool {
	return u.Role == RoleAdmin || u.Role == RoleOperator
}

// CanAccessFieldAgentView checks if user can access field agent view
func (u *User) CanAccessFieldAgentView() bool {
	return u.Role == RoleAdmin || u.Role == RoleFieldAgent
}

// CanAccessAdminSettings checks if user can access admin settings
func (u *User) CanAccessAdminSettings() bool {
	return u.Role == RoleAdmin
}
