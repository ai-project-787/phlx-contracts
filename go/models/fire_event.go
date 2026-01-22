// Owner: backend (fire event processing)
// Consumers: frontend, external-data-service, event-correlation-service

package models

import (
	"time"
)

type FireEvent struct {
	ID           string       `bson:"_id" json:"id"`
	LocationID   string       `bson:"location_id" json:"location_id"`
	LocationName string       `bson:"location_name" json:"location_name"`
	LocationType string       `bson:"location_type" json:"location_type"` // "asset", "facility"

	EventType string  `bson:"event_type" json:"event_type"` // "detected", "updated", "cleared"
	RiskLevel string  `bson:"risk_level" json:"risk_level"` // "none", "low", "medium", "high", "critical"
	RiskScore float64 `bson:"risk_score" json:"risk_score"` // 0-100 (weighted from 4 factors)

	Fires []FireDetail `bson:"fires" json:"fires"` // Nearby active fires (FIRMS)
	FWI   *FWIInfo     `bson:"fwi,omitempty" json:"fwi,omitempty"` // Fire Weather Index (EFFIS)

	ScoreFactors *ScoreFactors `bson:"score_factors,omitempty" json:"score_factors,omitempty"` // Breakdown of 4 factors

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Status string `bson:"status" json:"status"` // "active", "cleared", "acknowledged"
}

type FWIInfo struct {
	Value    float64 `bson:"value" json:"value"`       // FWI numeric value (0-100)
	Category string  `bson:"category" json:"category"` // "very_low", "low", "moderate", "high", "very_high", "extreme"
	Rating   int     `bson:"rating" json:"rating"`     // 1-10 numeric rating
}

type FireDetail struct {
	FireID          string  `bson:"fire_id" json:"fire_id"`
	Source          string  `bson:"source" json:"source"`                     // "copernicus_effis", "nasa_firms"
	SatelliteSource string  `bson:"satellite_source,omitempty" json:"satellite_source,omitempty"` // "VIIRS_NOAA20_NRT", "VIIRS_SNPP_NRT", "MODIS_NRT"
	Distance        float64 `bson:"distance" json:"distance"`                 // km
	InFire          bool    `bson:"in_fire" json:"in_fire"`
	Intensity       float64 `bson:"intensity,omitempty" json:"intensity,omitempty"`
	Confidence      string  `bson:"confidence,omitempty" json:"confidence,omitempty"`
}

type ScoreFactors struct {
	DistanceScore   float64 `bson:"distance_score" json:"distance_score"`     // 0-100 (proximity to fires)
	IntensityScore  float64 `bson:"intensity_score" json:"intensity_score"`   // 0-100 (fire intensity)
	ConfidenceScore float64 `bson:"confidence_score" json:"confidence_score"` // 0-100 (detection confidence)
	FWIScore        float64 `bson:"fwi_score" json:"fwi_score"`               // 0-100 (Fire Weather Index)
}
