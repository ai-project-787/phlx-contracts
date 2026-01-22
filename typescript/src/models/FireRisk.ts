/**
 * FireRisk model
 * Converted from go/models/fire_risk.go
 * Owner: backend (fire risk assessment)
 * Consumers: frontend, external-data-service
 */

import { GeoPoint } from './common';

export interface MonitoredLocation {
  id: string;
  name: string;
  type: string; // "asset", "facility", etc.
  location: GeoPoint;
  status: string;
}

export interface FireData {
  id: string;
  source: string;
  source_type: string;
  timestamp: string; // ISO 8601
  location: GeoPoint;
  data: Record<string, any>;
  tags: string[];
  source_metadata?: Record<string, any>;
}

