/**
 * User model
 * Converted from go/models/user.go
 * Owner: backend (auth service)
 * Consumers: all services (authentication)
 */

/** UserRole represents the role of a user in the system */
export const UserRole = {
  ADMIN: 'admin' as const,      // Full system access
  OPERATOR: 'operator' as const,   // Main dashboard + upload video
  FIELD_AGENT: 'field_agent' as const, // Field agent dispatch view
};

export type UserRoleType = typeof UserRole[keyof typeof UserRole];

/** User represents a user in the system with role-based access */
export interface User {
  id: string;
  email: string;
  // passwordHash is never exposed in JSON
  name: string;
  role: UserRoleType;
  assetId?: string; // For field agents
  active: boolean;
  createdAt: string; // ISO 8601
  updatedAt: string; // ISO 8601
  lastLoginAt?: string; // ISO 8601
  metadata?: Record<string, any>;
}

/** UserSession represents an active user session */
export interface UserSession {
  id: string;
  userId: string;
  token: string;
  expiresAt: string; // ISO 8601
  createdAt: string; // ISO 8601
  ipAddress?: string;
  userAgent?: string;
}
