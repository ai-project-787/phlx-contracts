/**
 * Event types for event-driven architecture
 * Converted from go/events/events.go
 */

import { CommandTarget } from '../models/TacticalCommand';

/** ImageBoundingBox represents a rectangular area in an image (pixel coordinates) */
export interface ImageBoundingBox {
  x: number;
  y: number;
  width: number;
  height: number;
}




export const EventType = {
  ASSET_UPDATE: 'asset_update' as const,
  ASSET_RECALL: 'asset_recall' as const,
  EMERGENCY_NOTIFICATION: 'emergency_notification' as const,
  CHAT_MESSAGE: 'chat_message' as const,
  SYSTEM_STATUS: 'system_status' as const,
  LOCATION_UPDATE: 'location_update' as const,
  VITALS_UPDATE: 'vitals_update' as const,
  VIDEO_UPLOAD: 'video_upload' as const,
  VIDEO_PROCESSING: 'video_processing' as const,
  FRAME_EXTRACTION: 'frame_extraction' as const,
  FRAME_UPLOAD_COMPLETE: 'frame_upload_complete' as const,
  AI_ANALYSIS: 'ai_analysis' as const,
  EVENT_ANALYSIS: 'event_analysis' as const,
  SUGGESTION_CREATED: 'suggestion_created' as const,
  MISSION_CREATED: 'mission_created' as const,
  AI_MISSION_SUGGESTION: 'ai_mission_suggestion' as const,
  TACTICAL_COMMAND_CREATED: 'tactical_command_created' as const,
  TACTICAL_COMMAND_RESPONSE: 'tactical_command_response' as const,
  TACTICAL_COMMAND_STATUS_CHANGED: 'tactical_command_status_changed' as const,
  TACTICAL_SUGGESTION_CREATED: 'tactical_suggestion_created' as const,
  FIRE_ALERT_CREATED: 'fire.alert.created' as const,
  MISSION_CHAT_MESSAGE: 'mission_chat_message' as const,
  MISSION_TYPING_INDICATOR: 'mission_typing_indicator' as const,
};

export type EventTypeValue = typeof EventType[keyof typeof EventType];

export interface BaseEvent {
  id: string;
  type: EventTypeValue;
  timestamp: string;
  source: string;
}



export interface LocationData {
  latitude: number;
  longitude: number;
  altitude?: number;
  accuracy?: number;
}

export interface AssetUpdateEventData extends BaseEvent {
  assetId: string;
  assetName: string;
  assetType: string;
  oldStatus: string;
  newStatus: string;
  location?: LocationData;
  metadata?: Record<string, any>;
}

export interface AssetRecallEventData extends BaseEvent {
  assetId: string;
  assetName: string;
  recalledBy: string;
  recalledByName: string;
  reason: string;
  location?: LocationData;
}

export interface EmergencyNotificationEventData extends BaseEvent {
  notificationId: string;
  title: string;
  message: string;
  severity: string;
  area: string;
  recipientCount: number;
  coordinates?: LocationData;
  acknowledged: boolean;
  acknowledgedBy?: string;
  acknowledgedAt?: string | null;
}

export interface ChatMessageEventData extends BaseEvent {
  messageId: string;
  text: string;
  sender: string;
  sessionId?: string;
  command?: string;
  response?: string;
}

export interface LocationUpdateEventData extends BaseEvent {
  assetId: string;
  assetName: string;
  location: LocationData;
  speed?: number;
  heading?: number;
  altitude?: number;
}

export interface VitalsUpdateEventData extends BaseEvent {
  personnelId: string;
  personnelName: string;
  pulseRate: number;
  oxygenLevel: number;
  temperature?: number;
  isAlert: boolean;
  alertReason?: string;
}

export interface SystemStatusEventData extends BaseEvent {
  status: string;
  previousStatus: string;
  changedBy: string;
  reason?: string;
  activeAssets: number;
  dispatchedAssets: number;
  metadata?: Record<string, any>;
}

export interface VideoUploadEventData extends BaseEvent {
  videoId: string;
  videoName: string;
  format: string;
  duration: number;
  fileSize: number;
  uploadedBy: string;
  gcsPath: string;
  status: string;
  cameraId: string;
  location?: LocationData;
}

export interface VideoProcessingEventData extends BaseEvent {
  videoId: string;
  jobType: string;
  status: string;
  progress: number;
  errorMsg?: string;
  startedAt?: string | null;
  completedAt?: string | null;
}

export interface FrameExtractionEventData extends BaseEvent {
  videoId: string;
  frameId: string;
  frameNumber: number;
  frameTimestamp: number;
  gcsPath: string;
  url: string;
  fileSize: number;
  cameraId: string;
  location?: LocationData;
}

export interface FrameUploadCompleteEventData extends BaseEvent {
  videoId: string;
  frameId: string;
  frameNumber: number;
  frameTimestamp: number;
  gcsPath: string;
  url: string;
  fileSize: number;
  verifiedAt: string;
  retryCount: number;
  cameraId: string;
  location?: LocationData;
}

export interface DetectedObject {
  type: string;
  confidence: number;
  boundingBox: ImageBoundingBox;
  attributes: Record<string, any>;
}

export interface DetectedEvent {
  type: string;
  confidence: number;
  description: string;
  severity: string;
  location?: LocationData;
  metadata: Record<string, any>;
}

export interface AIAnalysisEventData extends BaseEvent {
  videoId: string;
  frameId: string;
  confidence: number;
  objects: DetectedObject[];
  events: DetectedEvent[];
  metadata: Record<string, any>;
}

export interface EventAnalysisEventData extends BaseEvent {
  videoId: string;
  frameId: string;
  frameNumber: number;
  analysisType: string;
  description: string;
  summary: string;
  detectedItems: string[];
  confidence: number;
  severity: string;
  category: string;
  cameraId: string;
  location?: LocationData;
  metadata: Record<string, any>;
  rawResponse?: string;
}

export interface SuggestionCreatedEventData extends BaseEvent {
  suggestionId: string;
  eventId: string;
  missionId: string;
  missionTitle: string;
  confidence: number;
  reasoning: string;
}

export interface MissionCreatedEventData extends BaseEvent {
  missionId: string;
  title: string;
  description: string;
  priority: string;
  location?: LocationData;
  createdBy: string;
  createdByName: string;
}

export interface AIMissionSuggestionEventData extends BaseEvent {
  suggestionId: string;
  title: string;
  description: string;
  priority: string;
  reasoning: string;
  confidence: number;
  correlatedEventIds: string[];
  location?: LocationData;
}

export interface TacticalCommandCreatedEventData extends BaseEvent {
  commandId: string;
  missionId: string;
  title: string;
  description: string;
  category: string;
  priority: string;
  targets: CommandTarget[];
  createdBy: string;
  createdByName: string;
  source: string;
}

export interface TacticalCommandResponseEventData extends BaseEvent {
  commandId: string;
  targetId: string;
  targetType: string;
  targetName: string;
  decision: string;
  notes?: string;
  respondedBy: string;
  respondedByName: string;
}

export interface TacticalCommandStatusChangedEventData extends BaseEvent {
  commandId: string;
  oldStatus: string;
  newStatus: string;
  changedBy: string;
  changedByName: string;
  notes?: string;
}

export interface TacticalSuggestionCreatedEventData extends BaseEvent {
  suggestionId: string;
  missionId: string;
  title: string;
  description: string;
  category: string;
  priority: string;
  suggestedTargets: string[];
  reasoning: string;
  confidence: number;
}

export interface FireAlertCreatedEventData extends BaseEvent {
  alertId: string;
  fireEventId: string;
  locationId: string;
  locationName: string;
  severity: string;
  message: string;
  riskScore: number;
  location?: LocationData;
}

export interface MissionChatMessageEventData extends BaseEvent {
  messageId: string;
  missionId: string;
  senderId: string;
  senderName: string;
  senderRole: string;
  content: string;
}

export interface MissionTypingIndicatorEventData extends BaseEvent {
  missionId: string;
  userId: string;
  userName: string;
  isTyping: boolean;
}
