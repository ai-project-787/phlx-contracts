/**
 * AssetEventGroup model
 * Converted from go/models/asset_event_group.go
 * Owner: backend (event grouping)
 * Consumers: frontend, event-correlation-service
 */

/** GeoJSONPoint represents a geographical point in GeoJSON format */
export interface GeoJSONPoint {
  type: string; // Always "Point"
  coordinates: [number, number]; // [longitude, latitude]
}

/** Event represents a single event within a group */
export interface Event {
  id: string;
  type: string;
  timestamp: string; // ISO 8601
  location?: GeoJSONPoint;
  severity: string;
  description: string;
  metadata?: Record<string, any>;
}

/** AssetEventGroup represents grouped events by asset (camera or fire_detector) */
export interface AssetEventGroup {
  assetId: string;
  assetName: string;
  assetType: string; // "camera" or "fire_detector"
  eventCount: number;
  latestEvent: Event;
  eventIds: string[];
  events?: Event[]; // Full event data for display
}

/** GroupedEventsResponse is the response structure for grouped events API */
export interface GroupedEventsResponse {
  groups: AssetEventGroup[];
  count: number;
}
