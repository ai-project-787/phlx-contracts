/**
 * MissionChat model
 * Converted from go/models/mission_chat.go
 * Owner: backend (mission chat)
 * Consumers: frontend, field-agent-app
 */

/** MissionChatMessage represents a human chat message in a mission */
export interface MissionChatMessage {
  id: string;
  missionId: string;
  senderId: string;
  senderName: string;
  senderRole: string; // "operator" | "field_agent"
  content: string;
  timestamp: string; // ISO 8601
  createdAt: string; // ISO 8601
}

/** TypingUser represents a user who is currently typing */
export interface TypingUser {
  userId: string;
  userName: string;
}

/** TypingStatus represents the typing status in a mission chat */
export interface TypingStatus {
  missionId: string;
  typingUsers: TypingUser[];
}

/** SendMissionChatMessageRequest represents the request to send a chat message */
export interface SendMissionChatMessageRequest {
  content: string;
}

/** UpdateTypingStatusRequest represents the request to update typing status */
export interface UpdateTypingStatusRequest {
  isTyping: boolean;
}

/** MissionChatResponse represents a paginated chat history response */
export interface MissionChatResponse {
  messages: MissionChatMessage[];
  totalCount: number;
  hasMore: boolean;
}
