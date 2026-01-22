/**
 * AuditLog model
 * Converted from go/models/audit_log.go
 * Owner: backend (audit logging)
 * Consumers: all services (audit trail)
 */

/** AuditActionType represents the type of action in the audit log */
export const AuditActionType = {
  // Mission actions
  MISSION_CREATED: 'mission_created' as const,
  MISSION_CLAIMED: 'mission_claimed' as const,
  MISSION_COMPLETED: 'mission_completed' as const,
  MISSION_ARCHIVED: 'mission_archived' as const,
  MISSION_DELETED: 'mission_deleted' as const,

  // Asset actions
  ASSET_STATUS_CHANGED: 'asset_status_changed' as const,

  // Event actions (created by correlation service, read by backend)
  EVENT_CORRELATED: 'event_correlated' as const,
  EVENT_SUGGESTED: 'event_suggested' as const,
  EVENT_APPROVED: 'event_approved' as const,
  EVENT_REJECTED: 'event_rejected' as const,

  // Operator actions
  OPERATOR_ORDER: 'operator_order' as const,
  OPERATOR_NOTE: 'operator_note' as const,

  // Operational command actions (from AI Command Center)
  COMMAND_RECEIVED: 'command_received' as const,
  COMMAND_ACCEPTED: 'command_accepted' as const,
  COMMAND_DECLINED: 'command_declined' as const,
  COMMAND_STARTED: 'command_started' as const,
  COMMAND_COMPLETED: 'command_completed' as const,
};

export type AuditActionTypeValue = typeof AuditActionType[keyof typeof AuditActionType];

/** AuditLog represents a complete audit trail entry for mission actions */
export interface AuditLog {
  id: string;
  missionId: string;
  timestamp: string; // ISO 8601

  // Action details
  actionType: AuditActionTypeValue;
  actorType: string; // "operator", "asset", "ai_agent", "system"
  actorId: string;
  actorName?: string;

  // Target of action
  targetType?: string; // "asset", "dispatch", "event"
  targetId?: string;
  targetName?: string;

  // Details
  action: string; // Human-readable action
  details: Record<string, any>;
  createdAt: string; // ISO 8601
}
