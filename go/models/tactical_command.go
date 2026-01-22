// Owner: mission-command-service
// Consumers: backend, field-agent-app, ai-analysis-service

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TacticalCommandStatus represents the status of a tactical command
type TacticalCommandStatus string

const (
	TacticalCommandStatusPendingApproval TacticalCommandStatus = "pending_approval" // AI suggestion awaiting operator approval
	TacticalCommandStatusPending         TacticalCommandStatus = "pending"          // Approved, waiting for target response
	TacticalCommandStatusAccepted        TacticalCommandStatus = "accepted"         // Target accepted
	TacticalCommandStatusRejected        TacticalCommandStatus = "rejected"         // Target rejected
	TacticalCommandStatusInProgress      TacticalCommandStatus = "in_progress"      // Execution in progress
	TacticalCommandStatusCompleted       TacticalCommandStatus = "completed"        // Command completed
	TacticalCommandStatusCancelled       TacticalCommandStatus = "cancelled"        // Cancelled by operator
)

// TacticalCommandCategory represents predefined categories for UI icon/color mapping
type TacticalCommandCategory string

const (
	CategoryMovement      TacticalCommandCategory = "movement"      // Move to location, patrol routes
	CategorySecurity      TacticalCommandCategory = "security"      // Secure area, lockdown, perimeter
	CategorySurveillance  TacticalCommandCategory = "surveillance"  // Observe, reconnaissance, drone deploy
	CategoryDispatch      TacticalCommandCategory = "dispatch"      // Asset dispatch to location
	CategoryCommunication TacticalCommandCategory = "communication" // Notifications, alerts, coordination
	CategoryMedical       TacticalCommandCategory = "medical"       // Medical triage, evacuation
	CategoryEvacuation    TacticalCommandCategory = "evacuation"    // Evacuate personnel/civilians
	CategorySupport       TacticalCommandCategory = "support"       // Provide assistance, backup
	CategoryInvestigation TacticalCommandCategory = "investigation" // Investigate incident
	CategoryOther         TacticalCommandCategory = "other"         // Fallback for unmatched
)

// TacticalCommandPriority represents command priority levels
type TacticalCommandPriority string

const (
	PriorityRoutine   TacticalCommandPriority = "routine"   // Normal priority
	PriorityPriority  TacticalCommandPriority = "priority"  // Elevated priority
	PriorityImmediate TacticalCommandPriority = "immediate" // High priority
	PriorityFlash     TacticalCommandPriority = "flash"     // Critical/Emergency
)

// CommandTarget represents a target (asset or team) for a tactical command
type CommandTarget struct {
	TargetType string `json:"target_type" bson:"target_type"` // "asset" or "team"
	TargetID   string `json:"target_id" bson:"target_id"`
	TargetName string `json:"target_name" bson:"target_name"`
}

// CommandResponse represents a target's response to a tactical command
type CommandResponse struct {
	TargetID        string    `json:"target_id" bson:"target_id"`
	TargetType      string    `json:"target_type" bson:"target_type"`
	TargetName      string    `json:"target_name" bson:"target_name"`
	Decision        string    `json:"decision" bson:"decision"` // "accepted" or "rejected"
	Notes           string    `json:"notes,omitempty" bson:"notes,omitempty"`
	RespondedBy     string    `json:"responded_by" bson:"responded_by"`           // User ID who responded
	RespondedByName string    `json:"responded_by_name" bson:"responded_by_name"` // User name
	RespondedAt     time.Time `json:"responded_at" bson:"responded_at"`
}

// CommandStatusUpdate represents a status change in the command lifecycle
type CommandStatusUpdate struct {
	Status    TacticalCommandStatus `json:"status" bson:"status"`
	ChangedBy string                `json:"changed_by" bson:"changed_by"`           // User ID
	ChangedByName string            `json:"changed_by_name" bson:"changed_by_name"` // User name
	Timestamp time.Time             `json:"timestamp" bson:"timestamp"`
	Notes     string                `json:"notes,omitempty" bson:"notes,omitempty"`
}

// TacticalGeoLocation represents geographic coordinates for command destinations
type TacticalGeoLocation struct {
	Latitude    float64 `json:"lat" bson:"lat"`
	Longitude   float64 `json:"lng" bson:"lng"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty"`               // "Training Area West"
	Description string  `json:"description,omitempty" bson:"description,omitempty"` // "2.4 km from perimeter"
}

// TacticalGeoArea represents an area of operation (circle, polygon, or route)
type TacticalGeoArea struct {
	Type        string                `json:"type" bson:"type"` // "circle", "polygon", "route"
	Center      *TacticalGeoLocation  `json:"center,omitempty" bson:"center,omitempty"`
	Radius      float64               `json:"radius,omitempty" bson:"radius,omitempty"` // meters (for circle)
	Coordinates []TacticalGeoLocation `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	Name        string                `json:"name,omitempty" bson:"name,omitempty"`
}

// TacticalCommand represents a unified command model replacing OperationalCommand and DispatchRequest
type TacticalCommand struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// Mission Context (required - commands must link to existing missions)
	MissionID        primitive.ObjectID `json:"mission_id" bson:"mission_id"`
	MissionTitle     string             `json:"mission_title" bson:"mission_title"`
	SituationSummary string             `json:"situation_summary,omitempty" bson:"situation_summary,omitempty"`

	// Command Definition (AI-generated flexible titles)
	Title       string                  `json:"title" bson:"title"`             // AI-generated: "Deploy ISR drones to map fire front"
	Description string                  `json:"description" bson:"description"` // Detailed instructions
	Category    TacticalCommandCategory `json:"category" bson:"category"`       // From predefined list for UI icons

	// Multi-Target Assignment (can target multiple assets/teams)
	Targets []CommandTarget `json:"targets" bson:"targets"`

	// Location & Navigation
	Destination     *TacticalGeoLocation  `json:"destination,omitempty" bson:"destination,omitempty"`
	Waypoints       []TacticalGeoLocation `json:"waypoints,omitempty" bson:"waypoints,omitempty"`
	AreaOfOperation *TacticalGeoArea      `json:"area_of_operation,omitempty" bson:"area_of_operation,omitempty"`

	// Mission Parameters
	Objective string                  `json:"objective,omitempty" bson:"objective,omitempty"` // What to achieve
	Priority  TacticalCommandPriority `json:"priority" bson:"priority"`

	// Status & Responses (each target responds independently)
	Status        TacticalCommandStatus `json:"status" bson:"status"`
	Responses     []CommandResponse     `json:"responses,omitempty" bson:"responses,omitempty"`
	StatusHistory []CommandStatusUpdate `json:"status_history" bson:"status_history"`

	// Metadata
	Source    string             `json:"source" bson:"source"` // "ai" or "operator"
	CreatedBy primitive.ObjectID `json:"created_by" bson:"created_by"`
	CreatedByName string         `json:"created_by_name" bson:"created_by_name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`

	// Additional metadata for extensibility
	Metadata map[string]interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
}

// CreateTacticalCommandRequest represents the request to create a new tactical command
type CreateTacticalCommandRequest struct {
	MissionID        string                  `json:"mission_id" binding:"required"`
	Title            string                  `json:"title" binding:"required"`
	Description      string                  `json:"description" binding:"required"`
	Category         TacticalCommandCategory `json:"category" binding:"required"`
	Targets          []CommandTarget         `json:"targets"` // Either Targets or TargetNames must be provided
	TargetNames      string                  `json:"target_name,omitempty"` // Comma-separated asset names from Langflow
	Destination      *TacticalGeoLocation    `json:"destination,omitempty"`
	Waypoints        []TacticalGeoLocation   `json:"waypoints,omitempty"`
	AreaOfOperation  *TacticalGeoArea        `json:"area_of_operation,omitempty"`
	Objective        string                  `json:"objective,omitempty"`
	Priority         TacticalCommandPriority `json:"priority" binding:"required"`
	SituationSummary string                  `json:"situation_summary,omitempty"`
	Source           string                  `json:"source"` // "ai" or "operator", defaults to "operator"
	Metadata         map[string]interface{}  `json:"metadata,omitempty"`
}

// RespondToTacticalCommandRequest represents a target's response to a command
type RespondToTacticalCommandRequest struct {
	TargetID   string `json:"target_id" binding:"required"`
	TargetType string `json:"target_type" binding:"required"` // "asset" or "team"
	Decision   string `json:"decision" binding:"required"`    // "accepted" or "rejected"
	Notes      string `json:"notes,omitempty"`
}

// UpdateTacticalCommandStatusRequest represents a status update request
type UpdateTacticalCommandStatusRequest struct {
	Status TacticalCommandStatus `json:"status" binding:"required"`
	Notes  string                `json:"notes,omitempty"`
}

// TacticalCommandFilter represents query filters for listing commands
type TacticalCommandFilter struct {
	MissionID  string                `json:"mission_id,omitempty"`
	Status     TacticalCommandStatus `json:"status,omitempty"`
	TargetID   string                `json:"target_id,omitempty"`
	TargetType string                `json:"target_type,omitempty"`
	Category   TacticalCommandCategory `json:"category,omitempty"`
	Priority   TacticalCommandPriority `json:"priority,omitempty"`
	Source     string                `json:"source,omitempty"`
}

// ValidCategories returns all valid category values
func ValidCategories() []TacticalCommandCategory {
	return []TacticalCommandCategory{
		CategoryMovement,
		CategorySecurity,
		CategorySurveillance,
		CategoryDispatch,
		CategoryCommunication,
		CategoryMedical,
		CategoryEvacuation,
		CategorySupport,
		CategoryInvestigation,
		CategoryOther,
	}
}

// IsValidCategory checks if a category is valid
func IsValidCategory(category TacticalCommandCategory) bool {
	for _, c := range ValidCategories() {
		if c == category {
			return true
		}
	}
	return false
}

// ValidPriorities returns all valid priority values
func ValidPriorities() []TacticalCommandPriority {
	return []TacticalCommandPriority{
		PriorityRoutine,
		PriorityPriority,
		PriorityImmediate,
		PriorityFlash,
	}
}

// IsValidPriority checks if a priority is valid
func IsValidPriority(priority TacticalCommandPriority) bool {
	for _, p := range ValidPriorities() {
		if p == priority {
			return true
		}
	}
	return false
}

// ValidStatuses returns all valid status values
func ValidStatuses() []TacticalCommandStatus {
	return []TacticalCommandStatus{
		TacticalCommandStatusPendingApproval,
		TacticalCommandStatusPending,
		TacticalCommandStatusAccepted,
		TacticalCommandStatusRejected,
		TacticalCommandStatusInProgress,
		TacticalCommandStatusCompleted,
		TacticalCommandStatusCancelled,
	}
}

// IsValidStatus checks if a status is valid
func IsValidStatus(status TacticalCommandStatus) bool {
	for _, s := range ValidStatuses() {
		if s == status {
			return true
		}
	}
	return false
}
