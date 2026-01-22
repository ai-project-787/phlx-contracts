package models

import (
	"time"
)

// Owner: dispatch-asset-service
// Consumers: backend, mission-command-service, location-navigation-service, field-agent-app

// Asset status constants - simplified to 4 logical statuses
const (
	AssetStatusAvailable  = "available"  // Ready for dispatch
	AssetStatusDispatched = "dispatched" // On a mission
	AssetStatusReturning  = "returning"  // Coming back to base
	AssetStatusOffline    = "offline"    // Not operational
)

// Asset represents an asset in the system
type Asset struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Status           string                 `json:"status"`
	UseCase          string                 `json:"useCase"`
	TeamID           string                 `json:"teamId,omitempty"`      // Team this asset belongs to
	AssignedAreaIds  []string               `json:"assignedAreaIds,omitempty"` // Areas this asset patrols/monitors
	Latitude         float64                `json:"latitude"`
	Longitude        float64                `json:"longitude"`
	Altitude         float64                `json:"altitude,omitempty"`
	BatteryLevel     int                    `json:"batteryLevel,omitempty"`
	Members          int                    `json:"members,omitempty"`
	Vehicle          string                 `json:"vehicle,omitempty"`
	PulseRate        int                    `json:"pulseRate,omitempty"`
	OxygenLevel      int                    `json:"oxygenLevel,omitempty"`
	Location         string                 `json:"location,omitempty"`
	DispatchTime     *time.Time             `json:"dispatchTime,omitempty"`
	EstimatedArrival *time.Time             `json:"estimatedArrival,omitempty"`
	LastUpdated      time.Time              `json:"lastUpdated"`
	LastVitalUpdate  *time.Time             `json:"lastVitalUpdate,omitempty"`
	VideoSrc            string                 `json:"videoSrc,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	AutoPositionEnabled bool                   `json:"autoPositionEnabled"` // When false, position simulator skips this asset
}
