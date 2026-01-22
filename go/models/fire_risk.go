// Owner: backend (fire risk assessment)
// Consumers: frontend, external-data-service

package models

import (
	"time"
)

type MonitoredLocation struct {
	ID       string   `bson:"_id" json:"id"`
	Name     string   `bson:"name" json:"name"`
	Type     string   `bson:"type" json:"type"` // "asset", "facility", etc.
	Location GeoPoint `bson:"location" json:"location"`
	Status   string   `bson:"status" json:"status"`
}


type FireData struct {
	ID             string                 `bson:"_id" json:"id"`
	Source         string                 `bson:"source" json:"source"`
	SourceType     string                 `bson:"source_type" json:"source_type"`
	Timestamp      time.Time              `bson:"timestamp" json:"timestamp"`
	Location       GeoPoint               `bson:"location" json:"location"`
	Data           map[string]interface{} `bson:"data" json:"data"`
	Tags           []string               `bson:"tags" json:"tags"`
	SourceMetadata map[string]interface{} `bson:"source_metadata,omitempty" json:"source_metadata,omitempty"`
}

type BoundingBox struct {
	West  float64
	South float64
	East  float64
	North float64
}
