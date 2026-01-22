// Event Schema Version: 1.0
// Breaking changes require major version bump
//
// Event ownership:
// - AssetUpdateEvent → dispatch-asset-service
// - MissionCreatedEvent → mission-command-service
// - VideoUploadEvent → video-processing-service
// - TacticalCommandCreated → mission-command-service
// - FireAlertCreatedEvent → backend (temporary, to be moved)
//
// Consumers: All services (event-driven architecture)

package events

import (
	"time"
)

// EventType represents the type of event
type EventType string

const (
	AssetUpdateEvent             EventType = "asset_update"
	AssetRecallEvent             EventType = "asset_recall"
	EmergencyNotification        EventType = "emergency_notification"
	ChatMessageEvent             EventType = "chat_message"
	SystemStatusEvent            EventType = "system_status"
	LocationUpdateEvent          EventType = "location_update"
	VitalsUpdateEvent            EventType = "vitals_update"
	VideoUploadEvent             EventType = "video_upload"
	VideoProcessingEvent         EventType = "video_processing"
	FrameExtractionEvent         EventType = "frame_extraction"
	FrameUploadCompleteEvent     EventType = "frame_upload_complete"
	AIAnalysisEvent              EventType = "ai_analysis"
	EventAnalysisEvent           EventType = "event_analysis"
	SuggestionCreated            EventType = "suggestion_created"
	MissionCreated               EventType = "mission_created"
	AIMissionSuggestion          EventType = "ai_mission_suggestion"
	TacticalCommandCreated       EventType = "tactical_command_created"
	TacticalCommandResponse      EventType = "tactical_command_response"
	TacticalCommandStatusChanged EventType = "tactical_command_status_changed"
	TacticalSuggestionCreated    EventType = "tactical_suggestion_created"
	FireAlertCreatedEvent        EventType = "fire.alert.created"
	MissionChatMessageEvent      EventType = "mission_chat_message"
	MissionTypingIndicatorEvent  EventType = "mission_typing_indicator"
)

// BaseEvent represents the common structure for all events
type BaseEvent struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
}

// AssetUpdateEventData represents asset status changes
type AssetUpdateEventData struct {
	BaseEvent
	AssetID   string                 `json:"assetId"`
	AssetName string                 `json:"assetName"`
	AssetType string                 `json:"assetType"`
	OldStatus string                 `json:"oldStatus"`
	NewStatus string                 `json:"newStatus"`
	Location  *LocationData          `json:"location,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// EmergencyNotificationEventData represents emergency alerts
type EmergencyNotificationEventData struct {
	BaseEvent
	NotificationID string        `json:"notificationId"`
	Title          string        `json:"title"`
	Message        string        `json:"message"`
	Severity       string        `json:"severity"`
	Area           string        `json:"area"`
	RecipientCount int           `json:"recipientCount"`
	Coordinates    *LocationData `json:"coordinates,omitempty"`
	Acknowledged   bool          `json:"acknowledged"`
	AcknowledgedBy string        `json:"acknowledgedBy,omitempty"`
	AcknowledgedAt *time.Time    `json:"acknowledgedAt,omitempty"`
}

// ChatMessageEventData represents command panel messages
type ChatMessageEventData struct {
	BaseEvent
	MessageID string `json:"messageId"`
	Text      string `json:"text"`
	Sender    string `json:"sender"` // 'user', 'system', 'update', 'recommendation'
	SessionID string `json:"sessionId,omitempty"`
	Command   string `json:"command,omitempty"`
	Response  string `json:"response,omitempty"`
}

// LocationUpdateEventData represents real-time location updates
type LocationUpdateEventData struct {
	BaseEvent
	AssetID   string        `json:"assetId"`
	AssetName string        `json:"assetName"`
	Location  *LocationData `json:"location"`
	Speed     float64       `json:"speed,omitempty"`
	Heading   float64       `json:"heading,omitempty"`
	Altitude  float64       `json:"altitude,omitempty"`
}

// VitalsUpdateEventData represents personnel vital signs updates
type VitalsUpdateEventData struct {
	BaseEvent
	PersonnelID   string  `json:"personnelId"`
	PersonnelName string  `json:"personnelName"`
	PulseRate     int     `json:"pulseRate"`
	OxygenLevel   int     `json:"oxygenLevel"`
	Temperature   float64 `json:"temperature,omitempty"`
	IsAlert       bool    `json:"isAlert"`
	AlertReason   string  `json:"alertReason,omitempty"`
}

// SystemStatusEventData represents system-wide status changes
type SystemStatusEventData struct {
	BaseEvent
	Status           string                 `json:"status"` // 'Normal', 'Emergency', 'Maintenance'
	PreviousStatus   string                 `json:"previousStatus"`
	ChangedBy        string                 `json:"changedBy"`
	Reason           string                 `json:"reason,omitempty"`
	ActiveAssets     int                    `json:"activeAssets"`
	DispatchedAssets int                    `json:"dispatchedAssets"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

// VideoUploadEventData represents video upload events
type VideoUploadEventData struct {
	BaseEvent
	VideoID     string        `json:"videoId"`
	VideoName   string        `json:"videoName"`
	Format      string        `json:"format"`
	Duration    float64       `json:"duration"`
	FileSize    int64         `json:"fileSize"`
	UploadedBy  string        `json:"uploadedBy"`
	GCSPath     string        `json:"gcsPath"`
	Status      string        `json:"status"`
	CameraID    string        `json:"cameraId"`
	Location    *LocationData `json:"location,omitempty"`
}

// VideoProcessingEventData represents video processing status events
type VideoProcessingEventData struct {
	BaseEvent
	VideoID     string     `json:"videoId"`
	JobType     string     `json:"jobType"`  // 'frame_extraction', 'ai_analysis', 'transcoding'
	Status      string     `json:"status"`   // 'pending', 'running', 'completed', 'failed'
	Progress    float64    `json:"progress"` // 0-100
	ErrorMsg    string     `json:"errorMsg,omitempty"`
	StartedAt   *time.Time `json:"startedAt,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// FrameExtractionEventData represents frame extraction events
type FrameExtractionEventData struct {
	BaseEvent
	VideoID     string        `json:"videoId"`
	FrameID     string        `json:"frameId"`
	FrameNumber int           `json:"frameNumber"`
	Timestamp   float64       `json:"timestamp"` // seconds from video start
	GCSPath     string        `json:"gcsPath"`
	URL         string        `json:"url"`
	FileSize    int64         `json:"fileSize"`
	CameraID    string        `json:"cameraId"`
	Location    *LocationData `json:"location,omitempty"`
}

// FrameUploadCompleteEventData represents frame upload completion events
type FrameUploadCompleteEventData struct {
	BaseEvent
	VideoID     string        `json:"videoId"`
	FrameID     string        `json:"frameId"`
	FrameNumber int           `json:"frameNumber"`
	Timestamp   float64       `json:"timestamp"` // seconds from video start
	GCSPath     string        `json:"gcsPath"`
	URL         string        `json:"url"`
	FileSize    int64         `json:"fileSize"`
	VerifiedAt  time.Time     `json:"verifiedAt"`
	RetryCount  int           `json:"retryCount"`
	CameraID    string        `json:"cameraId"`
	Location    *LocationData `json:"location,omitempty"`
}

// AIAnalysisEventData represents AI analysis results
type AIAnalysisEventData struct {
	BaseEvent
	VideoID    string                 `json:"videoId"`
	FrameID    string                 `json:"frameId"`
	Confidence float64                `json:"confidence"`
	Objects    []DetectedObject       `json:"objects"`
	Events     []DetectedEvent        `json:"events"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// EventAnalysisEventData represents analyzed events from AI processing
type EventAnalysisEventData struct {
	BaseEvent
	VideoID       string                 `json:"videoId"`
	FrameID       string                 `json:"frameId"`
	FrameNumber   int                    `json:"frameNumber"`
	// Note: Timestamp is inherited from BaseEvent as time.Time
	AnalysisType  string                 `json:"analysisType"` // 'openai_vision', 'custom_model', etc.
	Description   string                 `json:"description"`
	Summary       string                 `json:"summary"`
	DetectedItems []string               `json:"detectedItems"`
	Confidence    float64                `json:"confidence"`
	Severity      string                 `json:"severity"` // 'low', 'medium', 'high', 'critical'
	Category      string                 `json:"category"` // 'security', 'safety', 'emergency', 'normal'
	CameraID      string                 `json:"cameraId"`
	Location      *LocationData          `json:"location,omitempty"`
	Metadata      map[string]interface{} `json:"metadata"`
	RawResponse   string                 `json:"rawResponse,omitempty"`
}

// DetectedObject represents an object detected in a frame
type DetectedObject struct {
	Type        string                 `json:"type"`
	Confidence  float64                `json:"confidence"`
	BoundingBox BoundingBox            `json:"boundingBox"`
	Attributes  map[string]interface{} `json:"attributes"`
}

// DetectedEvent represents an event detected in a frame
type DetectedEvent struct {
	Type        string                 `json:"type"`
	Confidence  float64                `json:"confidence"`
	Description string                 `json:"description"`
	Severity    string                 `json:"severity"`
	Location    *LocationData          `json:"location,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// SuggestionCreatedEventData represents a new AI correlation suggestion created
type SuggestionCreatedEventData struct {
	BaseEvent
	SuggestionID string  `json:"suggestionId"`
	EventID      string  `json:"eventId"`
	MissionID    string  `json:"missionId"`
	MissionTitle string  `json:"missionTitle"`
	Confidence   float64 `json:"confidence"`
	Reasoning    string  `json:"reasoning"`
}

// MissionCreatedEventData represents a new mission being created
type MissionCreatedEventData struct {
	BaseEvent
	MissionID   string        `json:"missionId"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Priority    string        `json:"priority"`
	Status      string        `json:"status"`
	Location    *LocationData `json:"location,omitempty"`
	AssetIDs    []string      `json:"assetIds"`
	Tags        []string      `json:"tags"`
	CreatedBy   string        `json:"createdBy"`
}

// AIMissionSuggestionEventData represents AI-generated suggestions for a mission
type AIMissionSuggestionEventData struct {
	BaseEvent
	MissionID        string                        `json:"missionId"`
	MissionTitle     string                        `json:"missionTitle"`
	TacticalCommands []TacticalCommandSuggestion   `json:"tacticalCommands,omitempty"`
	Analysis         string                        `json:"analysis"`
	Confidence       float64                       `json:"confidence"`
}

// TacticalCommandSuggestion represents a suggested tactical command
type TacticalCommandSuggestion struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	TargetType  string `json:"targetType"` // "team" or "asset"
	TargetID    string `json:"targetId,omitempty"`
	TargetName  string `json:"targetName,omitempty"`
	Priority    string `json:"priority"`
	Reasoning   string `json:"reasoning"`
}

// TacticalCommandTarget represents a target for tactical command events
type TacticalCommandTarget struct {
	TargetType string `json:"targetType"`
	TargetID   string `json:"targetId"`
	TargetName string `json:"targetName"`
}

// TacticalGeoLocation represents a location for tactical commands
type TacticalGeoLocation struct {
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}

// TacticalGeoArea represents an area of operation
type TacticalGeoArea struct {
	Type        string                `json:"type"`
	Center      *TacticalGeoLocation  `json:"center,omitempty"`
	Radius      float64               `json:"radius,omitempty"`
	Coordinates []TacticalGeoLocation `json:"coordinates,omitempty"`
	Name        string                `json:"name,omitempty"`
}

// TacticalCommandCreatedEventData represents a new tactical command created
type TacticalCommandCreatedEventData struct {
	BaseEvent
	CommandID        string                  `json:"commandId"`
	MissionID        string                  `json:"missionId"`
	MissionTitle     string                  `json:"missionTitle"`
	Title            string                  `json:"title"`
	Description      string                  `json:"description"`
	Category         string                  `json:"category"`
	Targets          []TacticalCommandTarget `json:"targets"`
	Priority         string                  `json:"priority"`
	CommandSource    string                  `json:"commandSource"` // "ai" or "operator"
	Destination      *TacticalGeoLocation    `json:"destination,omitempty"`
	AreaOfOperation  *TacticalGeoArea        `json:"areaOfOperation,omitempty"`
	Objective        string                  `json:"objective,omitempty"`
	SituationSummary string                  `json:"situationSummary,omitempty"`
}

// TacticalCommandResponseEventData represents a target's response to a tactical command
type TacticalCommandResponseEventData struct {
	BaseEvent
	CommandID       string `json:"commandId"`
	MissionID       string `json:"missionId"`
	TargetID        string `json:"targetId"`
	TargetType      string `json:"targetType"`
	TargetName      string `json:"targetName"`
	Decision        string `json:"decision"` // "accepted" or "rejected"
	Notes           string `json:"notes,omitempty"`
	RespondedBy     string `json:"respondedBy"`
	RespondedByName string `json:"respondedByName"`
	NewStatus       string `json:"newStatus"`
}

// TacticalCommandStatusEventData represents status changes for tactical commands
type TacticalCommandStatusEventData struct {
	BaseEvent
	CommandID     string `json:"commandId"`
	MissionID     string `json:"missionId"`
	CommandTitle  string `json:"commandTitle"`
	OldStatus     string `json:"oldStatus"`
	NewStatus     string `json:"newStatus"`
	UpdatedBy     string `json:"updatedBy"`
	UpdatedByName string `json:"updatedByName"`
	Notes         string `json:"notes,omitempty"`
}

// BoundingBox represents a rectangular area in an image
type BoundingBox struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// LocationData represents geographical coordinates
type LocationData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude,omitempty"`
	Address   string  `json:"address,omitempty"`
	Area      string  `json:"area,omitempty"`
}

// MissionChatMessageEventData represents a mission team chat message
type MissionChatMessageEventData struct {
	BaseEvent
	MissionID  string `json:"missionId"`
	MessageID  string `json:"messageId"`
	SenderID   string `json:"senderId"`
	SenderName string `json:"senderName"`
	SenderRole string `json:"senderRole"` // "operator" | "field_agent"
	Content    string `json:"content"`
}

// MissionTypingUser represents a user who is typing
type MissionTypingUser struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

// MissionTypingIndicatorEventData represents typing status in a mission chat
type MissionTypingIndicatorEventData struct {
	BaseEvent
	MissionID   string              `json:"missionId"`
	TypingUsers []MissionTypingUser `json:"typingUsers"`
}

// WebSocketMessage represents messages sent to frontend clients
type WebSocketMessage struct {
	Type      string      `json:"type"`
	Event     EventType   `json:"event"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
	ClientID  string      `json:"clientId,omitempty"`
}

