package events

// FireEventSchema defines the Kafka message schema for fire events
type FireEventSchema struct {
	SchemaVersion string      `json:"schema_version"` // "1.0"
	EventType     string      `json:"event_type"`     // "fire.risk.detected", "fire.risk.updated", "fire.risk.cleared"
	Timestamp     string      `json:"timestamp"`      // ISO 8601
	Payload       interface{} `json:"payload"`        // models.FireEvent
}
