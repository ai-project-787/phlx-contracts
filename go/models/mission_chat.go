// Owner: backend (mission chat)
// Consumers: frontend, field-agent-app

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MissionChatMessage represents a human chat message in a mission
type MissionChatMessage struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MissionID primitive.ObjectID `json:"missionId" bson:"mission_id"`
	SenderID  string             `json:"senderId" bson:"sender_id"`
	SenderName string            `json:"senderName" bson:"sender_name"`
	SenderRole string            `json:"senderRole" bson:"sender_role"` // "operator" | "field_agent"
	Content   string             `json:"content" bson:"content"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
}

// TypingUser represents a user who is currently typing
type TypingUser struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

// TypingStatus represents the typing status in a mission chat
type TypingStatus struct {
	MissionID   string       `json:"missionId"`
	TypingUsers []TypingUser `json:"typingUsers"`
}

// SendMissionChatMessageRequest represents the request to send a chat message
type SendMissionChatMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

// UpdateTypingStatusRequest represents the request to update typing status
type UpdateTypingStatusRequest struct {
	IsTyping bool `json:"isTyping"`
}

// MissionChatResponse represents a paginated chat history response
type MissionChatResponse struct {
	Messages   []MissionChatMessage `json:"messages"`
	TotalCount int64                `json:"totalCount"`
	HasMore    bool                 `json:"hasMore"`
}
