/**
 * Asset model
 * Converted from go/models/asset.go
 * Owner: dispatch-asset-service
 * Consumers: backend, mission-command-service, location-navigation-service, field-agent-app
 */

/** Asset status constants - simplified to 4 logical statuses */
export const AssetStatus = {
  AVAILABLE: 'available' as const,  // Ready for dispatch
  DISPATCHED: 'dispatched' as const, // On a mission
  RETURNING: 'returning' as const,   // Coming back to base
  OFFLINE: 'offline' as const,       // Not operational
};

export type AssetStatusType = typeof AssetStatus[keyof typeof AssetStatus];

/** Asset represents an asset in the system */
export interface Asset {
  id: string;
  name: string;
  type: string;
  status: string;
  useCase: string;
  teamId?: string; // Team this asset belongs to
  assignedAreaIds?: string[]; // Areas this asset patrols/monitors
  latitude: number;
  longitude: number;
  altitude?: number;
  batteryLevel?: number;
  members?: number;
  vehicle?: string;
  pulseRate?: number;
  oxygenLevel?: number;
  location?: string;
  dispatchTime?: string; // ISO 8601
  estimatedArrival?: string; // ISO 8601
  lastUpdated: string; // ISO 8601
  lastVitalUpdate?: string; // ISO 8601
  videoSrc?: string;
  metadata?: Record<string, any>;
  autoPositionEnabled: boolean; // When false, position simulator skips this asset
}
