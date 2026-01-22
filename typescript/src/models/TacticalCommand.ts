/**
 * TacticalCommand model
 * Converted from go/models/tactical_command.go
 * Owner: mission-command-service
 * Consumers: backend, field-agent-app, ai-analysis-service
 */

/** TacticalCommandStatus represents the status of a tactical command */
export const TacticalCommandStatus = {
  PENDING_APPROVAL: 'pending_approval' as const, // AI suggestion awaiting operator approval
  PENDING: 'pending' as const,          // Approved, waiting for target response
  ACCEPTED: 'accepted' as const,        // Target accepted
  REJECTED: 'rejected' as const,        // Target rejected
  IN_PROGRESS: 'in_progress' as const,  // Execution in progress
  COMPLETED: 'completed' as const,      // Command completed
  CANCELLED: 'cancelled' as const,      // Cancelled by operator
};

export type TacticalCommandStatusType = typeof TacticalCommandStatus[keyof typeof TacticalCommandStatus];

/** TacticalCommandCategory represents predefined categories for UI icon/color mapping */
export const TacticalCommandCategory = {
  MOVEMENT: 'movement' as const,
  SECURITY: 'security' as const,
  SURVEILLANCE: 'surveillance' as const,
  DISPATCH: 'dispatch' as const,
  COMMUNICATION: 'communication' as const,
  MEDICAL: 'medical' as const,
  EVACUATION: 'evacuation' as const,
  SUPPORT: 'support' as const,
  INVESTIGATION: 'investigation' as const,
  OTHER: 'other' as const,
};

export type TacticalCommandCategoryType = typeof TacticalCommandCategory[keyof typeof TacticalCommandCategory];

/** TacticalCommandPriority represents command priority levels */
export const TacticalCommandPriority = {
  ROUTINE: 'routine' as const,
  PRIORITY: 'priority' as const,
  IMMEDIATE: 'immediate' as const,
  FLASH: 'flash' as const,
};

export type TacticalCommandPriorityType = typeof TacticalCommandPriority[keyof typeof TacticalCommandPriority];

/** CommandTarget represents a target (asset or team) for a tactical command */
export interface CommandTarget {
  target_type: string; // "asset" or "team"
  target_id: string;
  target_name: string;
}

/** CommandResponse represents a target's response to a tactical command */
export interface CommandResponse {
  target_id: string;
  target_type: string;
  target_name: string;
  decision: string; // "accepted" or "rejected"
  notes?: string;
  responded_by: string; // User ID who responded
  responded_by_name: string; // User name
  responded_at: string; // ISO 8601
}

/** CommandStatusUpdate represents a status change in the command lifecycle */
export interface CommandStatusUpdate {
  status: TacticalCommandStatusType;
  changed_by: string; // User ID
  changed_by_name: string; // User name
  timestamp: string; // ISO 8601
  notes?: string;
}

/** TacticalGeoLocation represents geographic coordinates for command destinations */
export interface TacticalGeoLocation {
  lat: number;
  lng: number;
  name?: string; // "Training Area West"
  description?: string; // "2.4 km from perimeter"
}

/** TacticalGeoArea represents an area of operation (circle, polygon, or route) */
export interface TacticalGeoArea {
  type: string; // "circle", "polygon", "route"
  center?: TacticalGeoLocation;
  radius?: number; // meters (for circle)
  coordinates?: TacticalGeoLocation[];
  name?: string;
}

/** TacticalCommand represents a unified command model replacing OperationalCommand and DispatchRequest */
export interface TacticalCommand {
  id: string;

  // Mission Context (required - commands must link to existing missions)
  mission_id: string;
  mission_title: string;
  situation_summary?: string;

  // Command Definition (AI-generated flexible titles)
  title: string; // AI-generated: "Deploy ISR drones to map fire front"
  description: string; // Detailed instructions
  category: TacticalCommandCategoryType; // From predefined list for UI icons

  // Multi-Target Assignment (can target multiple assets/teams)
  targets: CommandTarget[];

  // Location & Navigation
  destination?: TacticalGeoLocation;
  waypoints?: TacticalGeoLocation[];
  area_of_operation?: TacticalGeoArea;

  // Mission Parameters
  objective?: string; // What to achieve
  priority: TacticalCommandPriorityType;

  // Status & Responses (each target responds independently)
  status: TacticalCommandStatusType;
  responses?: CommandResponse[];
  status_history: CommandStatusUpdate[];

  // Metadata
  source: string; // "ai" or "operator"
  created_by: string;
  created_by_name: string;
  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601

  // Additional metadata for extensibility
  metadata?: Record<string, any>;
}

/** CreateTacticalCommandRequest represents the request to create a new tactical command */
export interface CreateTacticalCommandRequest {
  mission_id: string;
  title: string;
  description: string;
  category: TacticalCommandCategoryType;
  targets?: CommandTarget[]; // Either targets or target_name must be provided
  target_name?: string; // Comma-separated asset names from Langflow
  destination?: TacticalGeoLocation;
  waypoints?: TacticalGeoLocation[];
  area_of_operation?: TacticalGeoArea;
  objective?: string;
  priority: TacticalCommandPriorityType;
  source?: string;
  situation_summary?: string;
}
