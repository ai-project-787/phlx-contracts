/**
 * Location model
 * Converted from go/models/location.go
 * Owner: location-navigation-service
 * Consumers: backend, dispatch-asset-service, mission-command-service
 */

import { Coordinate } from './common';

/** Area represents a named zone within a location with polygon boundary */
export interface Area {
  id: string; // UUID within location
  name: string;
  description?: string;
  boundary: Coordinate[]; // Polygon points
  fillColor?: string;
  borderColor?: string;
  opacity?: number; // 0.0 to 1.0
  type?: string; // perimeter, patrol_zone, checkpoint, etc.
  priority?: string; // high, medium, low
  active: boolean;
  createdAt: string; // ISO 8601
  updatedAt: string; // ISO 8601
}

/** Location represents a geographic location with center coordinates and multiple areas */
export interface Location {
  id: string;
  name: string;
  description?: string;

  // Center coordinates
  latitude: number;
  longitude: number;

  // Areas within this location
  areas: Area[];

  // Visual settings
  color?: string; // Hex color for map display
  icon?: string; // Icon identifier

  // Metadata
  useCase?: string; // military, police, defense, etc.
  tags?: string[];

  // Status
  active: boolean;

  // Audit
  createdBy?: string;
  createdAt: string; // ISO 8601
  updatedBy?: string;
  updatedAt: string; // ISO 8601
}

/** CreateLocationRequest represents the request to create a new location */
export interface CreateLocationRequest {
  name: string;
  description?: string;
  latitude: number;
  longitude: number;
  color?: string;
  icon?: string;
  useCase?: string;
  tags?: string[];
  active?: boolean;
}

/** UpdateLocationRequest represents the request to update a location */
export interface UpdateLocationRequest {
  name?: string;
  description?: string;
  latitude?: number;
  longitude?: number;
  color?: string;
  icon?: string;
  useCase?: string;
  tags?: string[];
  active?: boolean;
}

/** CreateAreaRequest represents the request to create a new area */
export interface CreateAreaRequest {
  name: string;
  description?: string;
  boundary: Coordinate[]; // min 3 points
  fillColor?: string;
  borderColor?: string;
  opacity?: number;
  type?: string;
  priority?: string;
  active?: boolean;
}

/** UpdateAreaRequest represents the request to update an area */
export interface UpdateAreaRequest {
  name?: string;
  description?: string;
  boundary?: Coordinate[];
  fillColor?: string;
  borderColor?: string;
  opacity?: number;
  type?: string;
  priority?: string;
  active?: boolean;
}
