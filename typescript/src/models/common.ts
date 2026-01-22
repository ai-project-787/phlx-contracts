/**
 * Common types used across all models
 * Converted from go/models/common.go
 */

/**
 * GeoPoint represents a geographical point in GeoJSON format
 * Used for MongoDB geospatial queries
 */
export interface GeoPoint {
  type: string; // Always "Point"
  coordinates: [number, number]; // [longitude, latitude]
}

/**
 * GeoLocation represents geographic coordinates in GeoJSON format
 * Used for mission and team locations
 */
export interface GeoLocation {
  type: string; // "Point"
  coordinates: [number, number]; // [longitude, latitude]
}

/**
 * Coordinate represents a GPS coordinate
 * Used for location boundaries and areas
 */
export interface Coordinate {
  latitude: number;
  longitude: number;
}


/**
 * BoundingBox represents a rectangular bounding box
 * Used for object detection and spatial queries
 */
export interface GeoBoundingBox {
  west: number;
  south: number;
  east: number;
  north: number;
}
