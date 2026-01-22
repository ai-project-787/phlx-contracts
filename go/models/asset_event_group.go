// Owner: backend (event grouping)
// Consumers: frontend, event-correlation-service

package models

import "time"

// AssetEventGroup represents grouped events by asset (camera or fire_detector)
type AssetEventGroup struct {
	AssetID     string   `json:"assetId"`
	AssetName   string   `json:"assetName"`
	AssetType   string   `json:"assetType"` // "camera" or "fire_detector"
	EventCount  int      `json:"eventCount"`
	LatestEvent Event    `json:"latestEvent"`
	EventIDs    []string `json:"eventIds"`
	Events      []Event  `json:"events,omitempty"` // Full event data for display
}

// Event represents a single event within a group
type Event struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Timestamp   string                 `json:"timestamp"`
	Location    *GeoJSONPoint          `json:"location,omitempty"`
	Severity    string                 `json:"severity"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// GeoJSONPoint represents a geographical point in GeoJSON format
type GeoJSONPoint struct {
	Type        string    `json:"type"` // Always "Point"
	Coordinates [2]float64 `json:"coordinates"` // [longitude, latitude]
}

// GroupedEventsResponse is the response structure for grouped events API
type GroupedEventsResponse struct {
	Groups []AssetEventGroup `json:"groups"`
	Count  int               `json:"count"`
}

// ConvertAlertToEventGroup converts an Alert to AssetEventGroup format
func ConvertAlertToEventGroup(alert *Alert, fireEvent *FireEvent) *AssetEventGroup {
	// Convert fire event location to GeoJSON if available
	var location *GeoJSONPoint
	if fireEvent != nil && fireEvent.Fires != nil && len(fireEvent.Fires) > 0 {
		// Use first fire location as representative location
		// In production, you might calculate centroid or use facility location
		// For now, we'll need to fetch location from a location service
		// This is a placeholder - actual implementation should fetch from location service
		location = nil // Will be populated by handler logic
	}

	// Extract metadata from fire event
	metadata := make(map[string]interface{})
	if fireEvent != nil {
		metadata["fireEventId"] = fireEvent.ID
		metadata["riskScore"] = fireEvent.RiskScore
		metadata["fireCount"] = len(fireEvent.Fires)

		// Add satellite source from first fire if available
		if len(fireEvent.Fires) > 0 && fireEvent.Fires[0].SatelliteSource != "" {
			metadata["satelliteSource"] = fireEvent.Fires[0].SatelliteSource
		}

		if fireEvent.FWI != nil {
			metadata["fwiValue"] = fireEvent.FWI.Value
			metadata["fwiCategory"] = fireEvent.FWI.Category
		}

		if fireEvent.ScoreFactors != nil {
			metadata["distanceScore"] = fireEvent.ScoreFactors.DistanceScore
			metadata["intensityScore"] = fireEvent.ScoreFactors.IntensityScore
		}
	}

	return &AssetEventGroup{
		AssetID:    alert.LocationID,
		AssetName:  alert.LocationName,
		AssetType:  "fire_detector",
		EventCount: 1,
		LatestEvent: Event{
			ID:          alert.ID,
			Type:        "fire_risk",
			Timestamp:   alert.CreatedAt.Format(time.RFC3339),
			Location:    location,
			Severity:    alert.Severity,
			Description: alert.Message,
			Metadata:    metadata,
		},
		EventIDs: []string{alert.ID},
	}
}
