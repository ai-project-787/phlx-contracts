/**
 * Kafka topic names for event routing
 * Converted from go/events/topics.go
 */

export const KafkaTopics = {
  ASSET_UPDATES: 'asset-updates',
  EMERGENCY_NOTIFICATIONS: 'emergency-notifications',
  CHAT_MESSAGES: 'chat-messages',
  TACTICAL_COMMANDS: 'tactical-commands',
  LOCATION_UPDATES: 'location-updates',
  VITALS_UPDATES: 'vitals-updates',
  SYSTEM_STATUS: 'system-status',
  VIDEO_UPLOADS: 'video-uploads',
  VIDEO_PROCESSING: 'video-processing',
  FRAME_EXTRACTION: 'frame-extraction',
  FRAME_UPLOAD_COMPLETE: 'frame-upload-complete',
  AI_ANALYSIS: 'ai-analysis',
  EVENT_ANALYSIS: 'event-analysis',
  CAMERA_EVENTS: 'camera-events',
  SUGGESTIONS: 'suggestions',
  MISSION_EVENTS: 'mission-events',
  AI_MISSION_SUGGESTIONS: 'ai-mission-suggestions',
  MISSION_CHAT: 'mission-chat',
} as const;

export type KafkaTopicName = typeof KafkaTopics[keyof typeof KafkaTopics];
