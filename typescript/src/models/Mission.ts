/**
 * Mission model
 * Converted from go/models/mission.go
 * Owner: mission-command-service
 * Consumers: backend, dispatch-asset-service, field-agent-app
 */

import { GeoLocation } from './common';

/** MissionStatus represents the current state of a mission */
export const MissionStatus = {
  ACTIVE: 'active' as const,
  COMPLETED: 'completed' as const,
  ARCHIVED: 'archived' as const,
};

export type MissionStatusType = typeof MissionStatus[keyof typeof MissionStatus];

/** Mission represents an operator-managed incident with correlated events */
export interface Mission {
  id: string;
  title: string;
  description: string;
  status: MissionStatusType;
  priority: string; // low, medium, high, critical

  // Operator management
  claimedByOperatorId?: string | null;
  claimedByOperatorName?: string | null;
  claimedAt?: string | null; // ISO 8601
  completedAt?: string | null; // ISO 8601
  completedByOperatorId?: string | null;

  // Mission scope
  dispatchIds: string[]; // All dispatches in this mission
  assetIds: string[]; // All dispatched assets
  eventIds: string[]; // All correlated events (read from correlation service)

  // Location (centroid of all events/assets)
  location?: GeoLocation;

  // Metadata
  tags: string[];
  createdAt: string; // ISO 8601
  updatedAt: string; // ISO 8601
}

/** CreateMissionRequest represents the request to create a new mission */
export interface CreateMissionRequest {
  title: string;
  description: string;
  priority: string;
  dispatchId: string;
  location?: GeoLocation;
}

/** UpdateMissionRequest represents the request to update a mission */
export interface UpdateMissionRequest {
  title?: string;
  description?: string;
  priority?: string;
  tags?: string[];
}

/** ClaimMissionRequest represents the request to claim a mission */
export interface ClaimMissionRequest {
  operatorId: string;
  operatorName: string;
}

/** CompleteMissionRequest represents the request to complete a mission */
export interface CompleteMissionRequest {
  operatorId: string;
}

/** DispatchResponseSummary represents a field agent's response to a dispatch */
export interface DispatchResponseSummary {
  assetId: string;
  assetName: string;
  accepted: boolean;
  responseTime: string; // ISO 8601
  notes?: string;
}

/** EnrichedDispatch represents a dispatch with its responses */
export interface EnrichedDispatch {
  id: string;
  eventId: string;
  description: string;
  status: string;
  priority: string;
  responses: DispatchResponseSummary[];
  createdAt: string; // ISO 8601
}

/** EnrichedMission represents a mission with full dispatch details */
export interface EnrichedMission extends Mission {
  dispatches: EnrichedDispatch[];
}

/** AddDispatchToMissionRequest represents adding a dispatch to an existing mission */
export interface AddDispatchToMissionRequest {
  dispatchId: string;
}

/** DirectCreateMissionRequest represents the request to create a mission directly (without dispatch) */
export interface DirectCreateMissionRequest {
  title: string;
  description: string;
  priority: string;
  eventId?: string;
  location?: GeoLocation;
}
