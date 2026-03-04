import type { QueryClient } from "@tanstack/react-query";
import { auditEventEmitter, auditEvents } from "@/events/audit";
import {
  AUDIT_ACTOR_SNAPSHOT,
  AUDIT_ACTOR_SNAPSHOT_LIST,
  AUDIT_EVENT,
  AUDIT_EVENT_LIST,
  AUDIT_EVENT_RETENTION_POLICY,
  AUDIT_EVENT_RETENTION_POLICY_LIST,
  AUDIT_EVENT_SEARCH_INDEX,
  AUDIT_EVENT_SEARCH_INDEX_LIST,
  AUDIT_HASH_CHAIN_CHECKPOINT,
  AUDIT_HASH_CHAIN_CHECKPOINT_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Audit 相关的缓存失效事件处理
 */
export const useAuditCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateEvents = () => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT_LIST });
  const handleInvalidateEvent = (item: string) => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT(item) });
  const handleInvalidateActorSnapshots = () => queryClient.invalidateQueries({ queryKey: AUDIT_ACTOR_SNAPSHOT_LIST });
  const handleInvalidateActorSnapshot = (item: string) => queryClient.invalidateQueries({ queryKey: AUDIT_ACTOR_SNAPSHOT(item) });
  const handleInvalidateEventRetentionPolicies = () => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT_RETENTION_POLICY_LIST });
  const handleInvalidateEventRetentionPolicy = (item: string) => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT_RETENTION_POLICY(item) });
  const handleInvalidateEventSearchIndices = () => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT_SEARCH_INDEX_LIST });
  const handleInvalidateEventSearchIndex = (item: string) => queryClient.invalidateQueries({ queryKey: AUDIT_EVENT_SEARCH_INDEX(item) });
  const handleInvalidateHashChainCheckpoints = () => queryClient.invalidateQueries({ queryKey: AUDIT_HASH_CHAIN_CHECKPOINT_LIST });
  const handleInvalidateHashChainCheckpoint = (item: string) => queryClient.invalidateQueries({ queryKey: AUDIT_HASH_CHAIN_CHECKPOINT(item) });

  auditEventEmitter.on(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints as EventCb);
  auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint as EventCb);

  return () => {
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints as EventCb);
    auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint as EventCb);
  };
};
