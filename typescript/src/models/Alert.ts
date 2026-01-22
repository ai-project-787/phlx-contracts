/**
 * Alert model
 * Converted from go/models/alert.go
 * Owner: backend (alert service)
 * Consumers: frontend, event-correlation-service
 */

export interface Alert {
  id: string;
  type: string; // "fire_risk", "asset_danger", "weather_warning"
  severity: string; // "low", "medium", "high", "critical"
  location_id: string;
  location_name: string;
  message: string;
  fire_event_id?: string; // Reference to source fire event

  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601

  status: string; // "active", "acknowledged", "resolved"
  acknowledged_at?: string | null; // ISO 8601
  acknowledged_by?: string;
}
