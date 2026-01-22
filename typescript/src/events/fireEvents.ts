/**
 * Fire event schema for Kafka messages
 * Converted from go/events/fire_event_schema.go
 */

import { FireEvent } from '../models/FireEvent';

/** FireEventSchema defines the Kafka message schema for fire events */
export interface FireEventSchema {
  schema_version: string; // "1.0"
  event_type: string; // "fire.risk.detected", "fire.risk.updated", "fire.risk.cleared"
  timestamp: string; // ISO 8601
  payload: FireEvent;
}
