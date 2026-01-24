// Audit Domain Types - 基于 NFX-ID Backend

import {
  ActorType,
  ResultType,
  RiskLevel,
  DataClassification,
  RetentionAction,
} from "./enums";

export interface Event {
  id: string;
  eventId: string;
  occurredAt: string;
  receivedAt: string;
  tenantId?: string;
  appId?: string;
  actorType: ActorType;
  actorId: string;
  actorTenantMemberId?: string;
  action: string;
  targetType?: string;
  targetId?: string;
  result: ResultType;
  failureReasonCode?: string;
  httpMethod?: string;
  httpPath?: string;
  httpStatus?: number;
  requestId?: string;
  traceId?: string;
  ip?: string;
  userAgent?: string;
  geoCountry?: string;
  riskLevel: RiskLevel;
  dataClassification: DataClassification;
  prevHash?: string;
  eventHash?: string;
  metadata?: Record<string, unknown>;
  createdAt: string;
}

export interface ActorSnapshot {
  id: string;
  actorType: string;
  actorId: string;
  displayName?: string;
  email?: string;
  clientName?: string;
  tenantId?: string;
  snapshotAt: string;
  snapshotData?: Record<string, unknown>;
  createdAt: string;
}

export interface EventRetentionPolicy {
  id: string;
  policyName: string;
  tenantId?: string;
  actionPattern?: string;
  dataClassification?: DataClassification;
  riskLevel?: RiskLevel;
  retentionDays: number;
  retentionAction: RetentionAction;
  archiveLocation?: string;
  status: string;
  createdAt: string;
  updatedAt: string;
  createdBy?: string;
}

export interface EventSearchIndex {
  id: string;
  eventId: string;
  tenantId?: string;
  appId?: string;
  actorType: ActorType;
  actorId: string;
  action: string;
  targetType?: string;
  targetId?: string;
  result: ResultType;
  occurredAt: string;
  ip?: string;
  createdAt: string;
}

export interface HashChainCheckpoint {
  id: string;
  checkpointId: string;
  tenantId?: string;
  partitionDate: string;
  checkpointHash: string;
  prevCheckpointHash?: string;
  eventCount: number;
  firstEventId?: string;
  lastEventId?: string;
  createdAt: string;
  createdBy?: string;
}
