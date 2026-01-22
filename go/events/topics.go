package events

// Kafka topic names for event routing
// Topics are managed by Kafka admin and must be pre-created in production

// KafkaTopics defines the Kafka topics used by the system
var KafkaTopics = struct {
	AssetUpdates           string
	EmergencyNotifications string
	ChatMessages           string
	TacticalCommands       string
	LocationUpdates        string
	VitalsUpdates          string
	SystemStatus           string
	VideoUploads           string
	VideoProcessing        string
	FrameExtraction        string
	FrameUploadComplete    string
	AIAnalysis             string
	EventAnalysis          string
	CameraEvents           string
	Suggestions            string
	MissionEvents          string
	AIMissionSuggestions   string
	MissionChat            string
}{
	AssetUpdates:           "asset-updates",
	EmergencyNotifications: "emergency-notifications",
	ChatMessages:           "chat-messages",
	TacticalCommands:       "tactical-commands",
	LocationUpdates:        "location-updates",
	VitalsUpdates:          "vitals-updates",
	SystemStatus:           "system-status",
	VideoUploads:           "video-uploads",
	VideoProcessing:        "video-processing",
	FrameExtraction:        "frame-extraction",
	FrameUploadComplete:    "frame-upload-complete",
	AIAnalysis:             "ai-analysis",
	EventAnalysis:          "event-analysis",
	CameraEvents:           "camera-events",
	Suggestions:            "suggestions",
	MissionEvents:          "mission-events",
	AIMissionSuggestions:   "ai-mission-suggestions",
	MissionChat:            "mission-chat",
}
