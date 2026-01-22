// Owner: backend (alert service)
// Consumers: frontend, event-correlation-service

package models

import "time"

type Alert struct {
	ID        string    `bson:"_id" json:"id"`
	Type      string    `bson:"type" json:"type"`               // "fire_risk", "asset_danger", "weather_warning"
	Severity  string    `bson:"severity" json:"severity"`       // "low", "medium", "high", "critical"
	LocationID   string `bson:"location_id" json:"location_id"`
	LocationName string `bson:"location_name" json:"location_name"`
	Message   string    `bson:"message" json:"message"`
	FireEventID string `bson:"fire_event_id,omitempty" json:"fire_event_id,omitempty"` // Reference to source fire event
	
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	
	Status          string     `bson:"status" json:"status"` // "active", "acknowledged", "resolved"
	AcknowledgedAt  *time.Time `bson:"acknowledged_at,omitempty" json:"acknowledged_at,omitempty"`
	AcknowledgedBy  string     `bson:"acknowledged_by,omitempty" json:"acknowledged_by,omitempty"`
}
