/**
 * FireEvent model
 * Converted from go/models/fire_event.go
 * Owner: backend (fire event processing)
 * Consumers: frontend, external-data-service, event-correlation-service
 */

export interface FWIInfo {
  value: number; // FWI numeric value (0-100)
  category: string; // "very_low", "low", "moderate", "high", "very_high", "extreme"
  rating: number; // 1-10 numeric rating
}

export interface FireDetail {
  fire_id: string;
  source: string; // "copernicus_effis", "nasa_firms"
  satellite_source?: string; // "VIIRS_NOAA20_NRT", "VIIRS_SNPP_NRT", "MODIS_NRT"
  distance: number; // km
  in_fire: boolean;
  intensity?: number;
  confidence?: string;
}

export interface ScoreFactors {
  distance_score: number; // 0-100 (proximity to fires)
  intensity_score: number; // 0-100 (fire intensity)
  confidence_score: number; // 0-100 (detection confidence)
  fwi_score: number; // 0-100 (Fire Weather Index)
}

export interface FireEvent {
  id: string;
  location_id: string;
  location_name: string;
  location_type: string; // "asset", "facility"

  event_type: string; // "detected", "updated", "cleared"
  risk_level: string; // "none", "low", "medium", "high", "critical"
  risk_score: number; // 0-100 (weighted from 4 factors)

  fires: FireDetail[]; // Nearby active fires (FIRMS)
  fwi?: FWIInfo; // Fire Weather Index (EFFIS)

  score_factors?: ScoreFactors; // Breakdown of 4 factors

  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601

  status: string; // "active", "cleared", "acknowledged"
}
