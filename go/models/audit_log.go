// Owner: backend (audit logging)
// Consumers: all services (audit trail)

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuditActionType represents the type of action in the audit log
type AuditActionType string

const (
	// Mission actions
	AuditActionMissionCreated   AuditActionType = "mission_created"
	AuditActionMissionClaimed   AuditActionType = "mission_claimed"
	AuditActionMissionCompleted AuditActionType = "mission_completed"
	AuditActionMissionArchived  AuditActionType = "mission_archived"
	AuditActionMissionDeleted   AuditActionType = "mission_deleted"

	// Asset actions
	AuditActionAssetStatusChanged AuditActionType = "asset_status_changed"

	// Event actions (created by correlation service, read by backend)
	AuditActionEventCorrelated AuditActionType = "event_correlated"
	AuditActionEventSuggested  AuditActionType = "event_suggested"
	AuditActionEventApproved   AuditActionType = "event_approved"
	AuditActionEventRejected   AuditActionType = "event_rejected"

	// Operator actions
	AuditActionOperatorOrder AuditActionType = "operator_order"
	AuditActionOperatorNote  AuditActionType = "operator_note"

	// Operational command actions (from AI Command Center)
	AuditActionCommandReceived AuditActionType = "command_received"
	AuditActionCommandAccepted AuditActionType = "command_accepted"
	AuditActionCommandDeclined AuditActionType = "command_declined"
	AuditActionCommandStarted  AuditActionType = "command_started"
	AuditActionCommandCompleted AuditActionType = "command_completed"
)

// AuditLog represents a complete audit trail entry for mission actions
type AuditLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MissionID string             `json:"missionId" bson:"mission_id"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`

	// Action details
	ActionType AuditActionType `json:"actionType" bson:"action_type"`
	ActorType  string          `json:"actorType" bson:"actor_type"` // "operator", "asset", "ai_agent", "system"
	ActorID    string          `json:"actorId" bson:"actor_id"`
	ActorName  string          `json:"actorName,omitempty" bson:"actor_name,omitempty"`

	// Target of action
	TargetType string `json:"targetType,omitempty" bson:"target_type,omitempty"` // "asset", "dispatch", "event"
	TargetID   string `json:"targetId,omitempty" bson:"target_id,omitempty"`
	TargetName string `json:"targetName,omitempty" bson:"target_name,omitempty"`

	// Details
	Action    string                 `json:"action" bson:"action"` // Human-readable action
	Details   map[string]interface{} `json:"details" bson:"details"`
	CreatedAt time.Time              `json:"createdAt" bson:"created_at"`
}
