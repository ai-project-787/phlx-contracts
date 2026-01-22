// Owner: backend (team management)
// Consumers: dispatch-asset-service, mission-command-service

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TeamStatus represents the operational status of a team
type TeamStatus string

const (
	TeamStatusActive   TeamStatus = "active"   // Team is operational
	TeamStatusInactive TeamStatus = "inactive" // Team is not operational
	TeamStatusDeployed TeamStatus = "deployed" // Team is currently on a mission
)

// Team represents a group of assets working together
type Team struct {
	ID           primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Name         string                 `bson:"name" json:"name"`
	Description  string                 `bson:"description,omitempty" json:"description,omitempty"`
	Status       TeamStatus             `bson:"status" json:"status"`
	Color        string                 `bson:"color,omitempty" json:"color,omitempty"` // Hex color for UI visualization (e.g., "#FF5733")

	// Team Composition
	AssetIDs     []string               `bson:"asset_ids" json:"assetIds"`       // Assets assigned to this team
	LeaderID     string                 `bson:"leader_id,omitempty" json:"leaderId,omitempty"` // Optional team leader asset ID

	// Team Capabilities
	Capabilities []string               `bson:"capabilities,omitempty" json:"capabilities,omitempty"` // e.g., ["firefighting", "medical", "rescue"]

	// Location (optional - could be computed from assets)
	BaseLocation *GeoLocation           `bson:"base_location,omitempty" json:"baseLocation,omitempty"`

	// Audit Trail
	CreatedBy    primitive.ObjectID     `bson:"created_by" json:"createdBy"`
	CreatedByName string                `bson:"created_by_name" json:"createdByName"`
	CreatedAt    time.Time              `bson:"created_at" json:"createdAt"`
	UpdatedAt    time.Time              `bson:"updated_at" json:"updatedAt"`

	// Metadata
	Metadata     map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"`
}

// TeamWithAssets extends Team with full asset details for display
type TeamWithAssets struct {
	Team
	Assets []Asset `json:"assets"`
}
