package models

// GeoPoint represents a geographical point in GeoJSON format
// Used for MongoDB geospatial queries
type GeoPoint struct {
	Type        string    `json:"type" bson:"type"`               // Always "Point"
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}

// GeoLocation represents geographic coordinates in GeoJSON format
// Used for mission and team locations
type GeoLocation struct {
	Type        string    `json:"type" bson:"type"`               // "Point"
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}

// Coordinate represents a GPS coordinate
// Used for location boundaries and areas
type Coordinate struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
