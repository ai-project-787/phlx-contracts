package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Owner: mission-command-service
// Consumers: backend, dispatch-asset-service, field-agent-app

// MissionStatus represents the current state of a mission
type MissionStatus string

const (
	MissionStatusActive    MissionStatus = "active"
	MissionStatusCompleted MissionStatus = "completed"
	MissionStatusArchived  MissionStatus = "archived"
)

// Mission represents an operator-managed incident with correlated events
type Mission struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      MissionStatus      `json:"status" bson:"status"`
	Priority    string             `json:"priority" bson:"priority"` // low, medium, high, critical

	// Operator management
	ClaimedByOperatorID   *string    `json:"claimedByOperatorId,omitempty" bson:"claimed_by_operator_id,omitempty"`
	ClaimedByOperatorName *string    `json:"claimedByOperatorName,omitempty" bson:"claimed_by_operator_name,omitempty"`
	ClaimedAt             *time.Time `json:"claimedAt,omitempty" bson:"claimed_at,omitempty"`
	CompletedAt           *time.Time `json:"completedAt,omitempty" bson:"completed_at,omitempty"`
	CompletedByOperatorID *string    `json:"completedByOperatorId,omitempty" bson:"completed_by_operator_id,omitempty"`

	// Mission scope
	DispatchIDs []string `json:"dispatchIds" bson:"dispatch_ids"` // All dispatches in this mission
	AssetIDs    []string `json:"assetIds" bson:"asset_ids"`       // All dispatched assets
	EventIDs    []string `json:"eventIds" bson:"event_ids"`       // All correlated events (read from correlation service)

	// Location (centroid of all events/assets)
	Location *GeoLocation `json:"location,omitempty" bson:"location,omitempty"`

	// Metadata
	Tags      []string  `json:"tags" bson:"tags"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

// CreateMissionRequest represents the request to create a new mission
type CreateMissionRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Priority    string    `json:"priority" binding:"required"`
	DispatchID  string    `json:"dispatchId" binding:"required"`
	Location    *GeoLocation `json:"location,omitempty"`
}

// UpdateMissionRequest represents the request to update a mission
type UpdateMissionRequest struct {
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	Priority    *string   `json:"priority,omitempty"`
	Tags        *[]string `json:"tags,omitempty"`
}

// ClaimMissionRequest represents the request to claim a mission
type ClaimMissionRequest struct {
	OperatorID   string `json:"operatorId" binding:"required"`
	OperatorName string `json:"operatorName" binding:"required"`
}

// CompleteMissionRequest represents the request to complete a mission
type CompleteMissionRequest struct {
	OperatorID string `json:"operatorId" binding:"required"`
}

// DispatchResponseSummary represents a field agent's response to a dispatch
type DispatchResponseSummary struct {
	AssetID      string    `json:"assetId"`
	AssetName    string    `json:"assetName"`
	Accepted     bool      `json:"accepted"`
	ResponseTime time.Time `json:"responseTime"`
	Notes        string    `json:"notes,omitempty"`
}

// EnrichedDispatch represents a dispatch with its responses
type EnrichedDispatch struct {
	ID          string                    `json:"id"`
	EventID     string                    `json:"eventId"`
	Description string                    `json:"description"`
	Status      string                    `json:"status"`
	Priority    string                    `json:"priority"`
	Responses   []DispatchResponseSummary `json:"responses"`
	CreatedAt   time.Time                 `json:"createdAt"`
}

// EnrichedMission represents a mission with full dispatch details
type EnrichedMission struct {
	Mission    `json:",inline"`
	Dispatches []EnrichedDispatch `json:"dispatches"`
}

// AddDispatchToMissionRequest represents adding a dispatch to an existing mission
type AddDispatchToMissionRequest struct {
	DispatchID string `json:"dispatchId" binding:"required"`
}

// DirectCreateMissionRequest represents the request to create a mission directly (without dispatch)
type DirectCreateMissionRequest struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    string       `json:"priority"`
	EventID     string       `json:"eventId,omitempty"`
	Location    *GeoLocation `json:"location,omitempty"`
}
