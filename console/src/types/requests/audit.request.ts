// Audit Request Types - 基于 NFX-ID Backend

// ========== 事件相关 ==========

export interface CreateEventRequest {
  eventId: string;
  occurredAt: string;
  tenantId?: string;
  appId?: string;
  actorType: string;
  actorId: string;
  actorTenantMemberId?: string;
  action: string;
  targetType?: string;
  targetId?: string;
  result: string;
  failureReasonCode?: string;
  httpMethod?: string;
  httpPath?: string;
  httpStatus?: number;
  requestId?: string;
  traceId?: string;
  ip?: string;
  userAgent?: string;
  geoCountry?: string;
  riskLevel?: string;
  dataClassification?: string;
  prevHash?: string;
  eventHash?: string;
  metadata?: Record<string, unknown>;
}

// ========== Actor Snapshot 相关 ==========

export interface CreateActorSnapshotRequest {
  actorType: string;
  actorId: string;
  displayName?: string;
  email?: string;
  clientName?: string;
  tenantId?: string;
  snapshotAt: string;
  snapshotData?: Record<string, unknown>;
}

// ========== Event Retention Policy 相关 ==========

export interface CreateEventRetentionPolicyRequest {
  policyName: string;
  tenantId?: string;
  actionPattern?: string;
  dataClassification?: string;
  riskLevel?: string;
  retentionDays: number;
  retentionAction: string;
  archiveLocation?: string;
  status?: string;
}

export interface UpdateEventRetentionPolicyRequest {
  policyName?: string;
  actionPattern?: string;
  dataClassification?: string;
  riskLevel?: string;
  retentionDays?: number;
  retentionAction?: string;
  archiveLocation?: string;
  status?: string;
}

// ========== Event Search Index 相关 ==========

export interface CreateEventSearchIndexRequest {
  eventId: string;
  tenantId?: string;
  appId?: string;
  actorType: string;
  actorId: string;
  action: string;
  targetType?: string;
  targetId?: string;
  result: string;
  occurredAt: string;
  ip?: string;
}

// ========== Hash Chain Checkpoint 相关 ==========

export interface CreateHashChainCheckpointRequest {
  checkpointId: string;
  tenantId?: string;
  partitionDate: string;
  checkpointHash: string;
  prevCheckpointHash?: string;
  eventCount: number;
  firstEventId?: string;
  lastEventId?: string;
  createdBy?: string;
}
