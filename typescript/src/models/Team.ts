/**
 * Team model
 * Converted from go/models/team.go
 * Owner: backend (team management)
 * Consumers: dispatch-asset-service, mission-command-service
 */

import { GeoLocation } from './common';
import { Asset } from './Asset';

/** TeamStatus represents the operational status of a team */
export const TeamStatus = {
  ACTIVE: 'active' as const,   // Team is operational
  INACTIVE: 'inactive' as const, // Team is not operational
  DEPLOYED: 'deployed' as const, // Team is currently on a mission
};

export type TeamStatusType = typeof TeamStatus[keyof typeof TeamStatus];

/** Team represents a group of assets working together */
export interface Team {
  id: string;
  name: string;
  description?: string;
  status: TeamStatusType;
  color?: string; // Hex color for UI visualization (e.g., "#FF5733")

  // Team Composition
  assetIds: string[]; // Assets assigned to this team
  leaderId?: string; // Optional team leader asset ID

  // Team Capabilities
  capabilities?: string[]; // e.g., ["firefighting", "medical", "rescue"]

  // Location (optional - could be computed from assets)
  baseLocation?: GeoLocation;

  // Audit Trail
  createdBy: string;
  createdByName: string;
  createdAt: string; // ISO 8601
  updatedAt: string; // ISO 8601

  // Metadata
  metadata?: Record<string, any>;
}

/** TeamWithAssets extends Team with full asset details for display */
export interface TeamWithAssets extends Team {
  assets: Asset[];
}
