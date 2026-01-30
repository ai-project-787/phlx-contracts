/**
 * Composition models
 * Converted from go/models/composition.go
 * Owner: composition-service
 * Consumers: frontend, admin-settings
 */

/** GridSlot represents a camera feed in a specific grid position */
export interface GridSlot {
  cameraUrl: string;
  position: number; // 0=TL, 1=TR, 2=BL, 3=BR (for 2x2), extends for larger grids
}

/** GridSize represents the supported grid dimensions (1x1 to 4x4) */
export type GridSize = 1 | 2 | 3 | 4;

/** GridConfig defines an NxN grid composition configuration (1x1 to 4x4) */
export interface GridConfig {
  sessionId: string;
  missionId?: string; // Optional, for priority lookup
  gridSize?: GridSize; // 1-4, defaults to 2 for backward compatibility
  slots: GridSlot[];
  outputUrl: string;
}

/** CompositionStatus reports health status of a composition process */
export interface CompositionStatus {
  sessionId: string;
  isRunning: boolean;
  startTime: string; // ISO 8601
  restarts: number;
  encoder: string;
  outputUrl: string;
  lastError?: string;
  profile?: string; // Current bandwidth profile
  bitrateKbps?: number; // Current allocated bitrate
}

/** StreamProfile represents quality levels for composite streams */
export const StreamProfile = {
  BACKGROUND: 'Background' as const, // 480p @ 15fps, 1 Mbps
  MONITORING: 'Monitoring' as const, // 720p @ 30fps, 2.5 Mbps
  MISSION_CRITICAL: 'MissionCritical' as const, // 1080p @ 30fps, 4 Mbps
};

export type StreamProfileType = typeof StreamProfile[keyof typeof StreamProfile];

/** ProfileConfig defines the encoding parameters for a stream profile */
export interface ProfileConfig {
  resolution: string; // "854x480", "1280x720", "1920x1080"
  bitrate: number; // Mbps (1, 2.5, 4)
  fps: number; // 15, 30, 30
  preset: string; // "ultrafast", "fast", "fast"
}

/** Grid size configuration for UI */
export const GRID_SIZE_CONFIG: Record<GridSize, { label: string; slots: number }> = {
  1: { label: '1x1', slots: 1 },
  2: { label: '2x2', slots: 4 },
  3: { label: '3x3', slots: 9 },
  4: { label: '4x4', slots: 16 },
};

/** Helper to get grid size, defaulting to 2 for backward compatibility */
export function getGridSize(config: GridConfig): GridSize {
  if (!config.gridSize || config.gridSize < 1 || config.gridSize > 4) {
    return 2;
  }
  return config.gridSize as GridSize;
}

/** Helper to get max slots for a grid size */
export function getMaxSlots(gridSize: GridSize): number {
  return gridSize * gridSize;
}
