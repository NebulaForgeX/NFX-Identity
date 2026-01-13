export const auditEvents = {
  // Event 相关
  INVALIDATE_EVENT: "AUDIT:INVALIDATE_EVENT",
  INVALIDATE_EVENTS: "AUDIT:INVALIDATE_EVENTS",

  // ActorSnapshot 相关
  INVALIDATE_ACTOR_SNAPSHOT: "AUDIT:INVALIDATE_ACTOR_SNAPSHOT",
  INVALIDATE_ACTOR_SNAPSHOTS: "AUDIT:INVALIDATE_ACTOR_SNAPSHOTS",

  // EventRetentionPolicy 相关
  INVALIDATE_EVENT_RETENTION_POLICY: "AUDIT:INVALIDATE_EVENT_RETENTION_POLICY",
  INVALIDATE_EVENT_RETENTION_POLICIES: "AUDIT:INVALIDATE_EVENT_RETENTION_POLICIES",

  // EventSearchIndex 相关
  INVALIDATE_EVENT_SEARCH_INDEX: "AUDIT:INVALIDATE_EVENT_SEARCH_INDEX",
  INVALIDATE_EVENT_SEARCH_INDICES: "AUDIT:INVALIDATE_EVENT_SEARCH_INDICES",

  // HashChainCheckpoint 相关
  INVALIDATE_HASH_CHAIN_CHECKPOINT: "AUDIT:INVALIDATE_HASH_CHAIN_CHECKPOINT",
  INVALIDATE_HASH_CHAIN_CHECKPOINTS: "AUDIT:INVALIDATE_HASH_CHAIN_CHECKPOINTS",
} as const;

type AuditEvent = (typeof auditEvents)[keyof typeof auditEvents];

class AuditEventEmitter {
  private listeners: Record<AuditEvent, Set<Function>> = {
    [auditEvents.INVALIDATE_EVENT]: new Set<Function>(),
    [auditEvents.INVALIDATE_EVENTS]: new Set<Function>(),
    [auditEvents.INVALIDATE_ACTOR_SNAPSHOT]: new Set<Function>(),
    [auditEvents.INVALIDATE_ACTOR_SNAPSHOTS]: new Set<Function>(),
    [auditEvents.INVALIDATE_EVENT_RETENTION_POLICY]: new Set<Function>(),
    [auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES]: new Set<Function>(),
    [auditEvents.INVALIDATE_EVENT_SEARCH_INDEX]: new Set<Function>(),
    [auditEvents.INVALIDATE_EVENT_SEARCH_INDICES]: new Set<Function>(),
    [auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT]: new Set<Function>(),
    [auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS]: new Set<Function>(),
  };

  on(event: AuditEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: AuditEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: AuditEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const auditEventEmitter = new AuditEventEmitter();
