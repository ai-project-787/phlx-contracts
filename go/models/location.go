// Owner: location-navigation-service
// Consumers: backend, dispatch-asset-service, mission-command-service

package models

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


// Area represents a named zone within a location with polygon boundary
type Area struct {
	ID          string       `json:"id" bson:"id"` // UUID within location
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description,omitempty" bson:"description,omitempty"`
	Boundary    []Coordinate `json:"boundary" bson:"boundary"` // Polygon points
	FillColor   string       `json:"fillColor,omitempty" bson:"fill_color,omitempty"`
	BorderColor string       `json:"borderColor,omitempty" bson:"border_color,omitempty"`
	Opacity     float64      `json:"opacity,omitempty" bson:"opacity,omitempty"` // 0.0 to 1.0
	Type        string       `json:"type,omitempty" bson:"type,omitempty"`       // perimeter, patrol_zone, checkpoint, etc.
	Priority    string       `json:"priority,omitempty" bson:"priority,omitempty"` // high, medium, low
	Active      bool         `json:"active" bson:"active"`
	CreatedAt   time.Time    `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time    `json:"updatedAt" bson:"updated_at"`
}

// Location represents a geographic location with center coordinates and multiple areas
type Location struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`

	// Center coordinates
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`

	// Areas within this location
	Areas []Area `json:"areas" bson:"areas"`

	// Visual settings
	Color string `json:"color,omitempty" bson:"color,omitempty"` // Hex color for map display
	Icon  string `json:"icon,omitempty" bson:"icon,omitempty"`   // Icon identifier

	// Metadata
	UseCase string   `json:"useCase,omitempty" bson:"use_case,omitempty"` // military, police, defense, etc.
	Tags    []string `json:"tags,omitempty" bson:"tags,omitempty"`

	// Status
	Active bool `json:"active" bson:"active"`

	// Audit
	CreatedBy string    `json:"createdBy,omitempty" bson:"created_by,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedBy string    `json:"updatedBy,omitempty" bson:"updated_by,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

// CreateLocationRequest represents the request to create a new location
type CreateLocationRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description,omitempty"`
	Latitude    float64  `json:"latitude" binding:"required"`
	Longitude   float64  `json:"longitude" binding:"required"`
	Color       string   `json:"color,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	UseCase     string   `json:"useCase,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Active      *bool    `json:"active,omitempty"` // Pointer to distinguish false from not set
}

// UpdateLocationRequest represents the request to update a location
type UpdateLocationRequest struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Latitude    *float64  `json:"latitude,omitempty"`
	Longitude   *float64  `json:"longitude,omitempty"`
	Color       *string   `json:"color,omitempty"`
	Icon        *string   `json:"icon,omitempty"`
	UseCase     *string   `json:"useCase,omitempty"`
	Tags        *[]string `json:"tags,omitempty"`
	Active      *bool     `json:"active,omitempty"`
}

// CreateAreaRequest represents the request to create a new area
type CreateAreaRequest struct {
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description,omitempty"`
	Boundary    []Coordinate `json:"boundary" binding:"required,min=3"`
	FillColor   string       `json:"fillColor,omitempty"`
	BorderColor string       `json:"borderColor,omitempty"`
	Opacity     *float64     `json:"opacity,omitempty"`
	Type        string       `json:"type,omitempty"`
	Priority    string       `json:"priority,omitempty"`
	Active      *bool        `json:"active,omitempty"`
}

// UpdateAreaRequest represents the request to update an area
type UpdateAreaRequest struct {
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Boundary    *[]Coordinate `json:"boundary,omitempty"`
	FillColor   *string       `json:"fillColor,omitempty"`
	BorderColor *string       `json:"borderColor,omitempty"`
	Opacity     *float64      `json:"opacity,omitempty"`
	Type        *string       `json:"type,omitempty"`
	Priority    *string       `json:"priority,omitempty"`
	Active      *bool         `json:"active,omitempty"`
}

// Validate validates coordinate values
func (c *Coordinate) Validate() error {
	if c.Latitude < -90 || c.Latitude > 90 {
		return fmt.Errorf("invalid latitude: %f (must be between -90 and 90)", c.Latitude)
	}
	if c.Longitude < -180 || c.Longitude > 180 {
		return fmt.Errorf("invalid longitude: %f (must be between -180 and 180)", c.Longitude)
	}
	return nil
}

// Validate validates a boundary polygon
func ValidateBoundary(boundary []Coordinate) error {
	// Minimum 3 points for a polygon
	if len(boundary) < 3 {
		return errors.New("boundary must have at least 3 points to form a polygon")
	}

	// Check for valid coordinates
	for i, coord := range boundary {
		if err := coord.Validate(); err != nil {
			return fmt.Errorf("invalid coordinate at point %d: %w", i, err)
		}
	}

	return nil
}

// IsPointInPolygon determines if a point is inside a polygon boundary using ray casting algorithm
func IsPointInPolygon(point Coordinate, polygon []Coordinate) bool {
	if len(polygon) < 3 {
		return false
	}

	inside := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		xi, yi := polygon[i].Longitude, polygon[i].Latitude
		xj, yj := polygon[j].Longitude, polygon[j].Latitude

		intersect := ((yi > point.Latitude) != (yj > point.Latitude)) &&
			(point.Longitude < (xj-xi)*(point.Latitude-yi)/(yj-yi)+xi)

		if intersect {
			inside = !inside
		}

		j = i
	}

	return inside
}

// GetAssetsInArea returns which assets from a list are inside this area
func (a *Area) GetAssetsInArea(assets []Asset) []Asset {
	var assetsInArea []Asset

	for _, asset := range assets {
		// Skip assets without valid location (0,0 is considered invalid)
		if asset.Latitude == 0 && asset.Longitude == 0 {
			continue
		}

		point := Coordinate{
			Latitude:  asset.Latitude,
			Longitude: asset.Longitude,
		}

		if IsPointInPolygon(point, a.Boundary) {
			assetsInArea = append(assetsInArea, asset)
		}
	}

	return assetsInArea
}

// FindAreaByID finds an area by ID within a location
func (l *Location) FindAreaByID(areaID string) *Area {
	for i := range l.Areas {
		if l.Areas[i].ID == areaID {
			return &l.Areas[i]
		}
	}
	return nil
}

// AddArea adds a new area to the location
func (l *Location) AddArea(area Area) {
	l.Areas = append(l.Areas, area)
	l.UpdatedAt = time.Now()
}

// UpdateArea updates an existing area in the location
func (l *Location) UpdateArea(areaID string, updates UpdateAreaRequest) error {
	for i := range l.Areas {
		if l.Areas[i].ID == areaID {
			if updates.Name != nil {
				l.Areas[i].Name = *updates.Name
			}
			if updates.Description != nil {
				l.Areas[i].Description = *updates.Description
			}
			if updates.Boundary != nil {
				if err := ValidateBoundary(*updates.Boundary); err != nil {
					return err
				}
				l.Areas[i].Boundary = *updates.Boundary
			}
			if updates.FillColor != nil {
				l.Areas[i].FillColor = *updates.FillColor
			}
			if updates.BorderColor != nil {
				l.Areas[i].BorderColor = *updates.BorderColor
			}
			if updates.Opacity != nil {
				l.Areas[i].Opacity = *updates.Opacity
			}
			if updates.Type != nil {
				l.Areas[i].Type = *updates.Type
			}
			if updates.Priority != nil {
				l.Areas[i].Priority = *updates.Priority
			}
			if updates.Active != nil {
				l.Areas[i].Active = *updates.Active
			}
			l.Areas[i].UpdatedAt = time.Now()
			l.UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("area with ID %s not found", areaID)
}

// RemoveArea removes an area from the location
func (l *Location) RemoveArea(areaID string) error {
	for i := range l.Areas {
		if l.Areas[i].ID == areaID {
			l.Areas = append(l.Areas[:i], l.Areas[i+1:]...)
			l.UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("area with ID %s not found", areaID)
}
