package models

import "time"

// GridSlot represents a camera feed in a specific grid position
// Owner: composition-service
type GridSlot struct {
	CameraURL string `json:"camera_url" bson:"camera_url"`
	Position  int    `json:"position" bson:"position"` // 0=TL, 1=TR, 2=BL, 3=BR
}

// GridConfig defines a 2x2 grid composition configuration
// Owner: composition-service
type GridConfig struct {
	SessionID string     `json:"session_id" bson:"session_id"`
	MissionID string     `json:"mission_id,omitempty" bson:"mission_id,omitempty"` // Optional, for priority lookup
	Slots     []GridSlot `json:"slots" bson:"slots"`
	OutputURL string     `json:"output_url" bson:"output_url"`
}

// CompositionStatus reports health status of a composition process
// Owner: composition-service
type CompositionStatus struct {
	SessionID   string    `json:"session_id"`
	IsRunning   bool      `json:"is_running"`
	StartTime   time.Time `json:"start_time"`
	Restarts    int       `json:"restarts"`
	Encoder     string    `json:"encoder"`
	OutputURL   string    `json:"output_url"`
	LastError   string    `json:"last_error,omitempty"`
	Profile     string    `json:"profile,omitempty"`     // Current bandwidth profile
	BitrateKbps int       `json:"bitrate_kbps,omitempty"` // Current allocated bitrate
}

// StreamProfile represents quality levels for composite streams
// Owner: composition-service
type StreamProfile int

const (
	// ProfileBackground is for inactive monitoring (lowest quality)
	ProfileBackground StreamProfile = iota // 480p @ 15fps, 1 Mbps
	// ProfileMonitoring is for general use (medium quality)
	ProfileMonitoring // 720p @ 30fps, 2.5 Mbps
	// ProfileMissionCritical is for active missions (highest quality)
	ProfileMissionCritical // 1080p @ 30fps, 4 Mbps
)

// String returns the string representation of a StreamProfile
func (p StreamProfile) String() string {
	switch p {
	case ProfileBackground:
		return "Background"
	case ProfileMonitoring:
		return "Monitoring"
	case ProfileMissionCritical:
		return "MissionCritical"
	default:
		return "Unknown"
	}
}

// ParseStreamProfile converts a string to a StreamProfile
func ParseStreamProfile(s string) (StreamProfile, bool) {
	switch s {
	case "Background":
		return ProfileBackground, true
	case "Monitoring":
		return ProfileMonitoring, true
	case "MissionCritical":
		return ProfileMissionCritical, true
	default:
		return ProfileMonitoring, false // Default to Monitoring
	}
}

// ProfileConfig defines the encoding parameters for a stream profile
// Owner: composition-service
type ProfileConfig struct {
	Resolution string // "854x480", "1280x720", "1920x1080"
	Bitrate    int    // Mbps (1, 2.5, 4)
	FPS        int    // 15, 30, 30
	Preset     string // "ultrafast", "fast", "fast"
}

// GetProfileConfig returns the configuration for a given profile
func GetProfileConfig(profile StreamProfile) ProfileConfig {
	switch profile {
	case ProfileBackground:
		return ProfileConfig{
			Resolution: "854x480",
			Bitrate:    1,
			FPS:        15,
			Preset:     "ultrafast",
		}
	case ProfileMonitoring:
		return ProfileConfig{
			Resolution: "1280x720",
			Bitrate:    2,
			FPS:        30,
			Preset:     "fast",
		}
	case ProfileMissionCritical:
		return ProfileConfig{
			Resolution: "1920x1080",
			Bitrate:    4,
			FPS:        30,
			Preset:     "fast",
		}
	default:
		return GetProfileConfig(ProfileMonitoring)
	}
}
