import type { QueryClient } from "@tanstack/react-query";
import { auditEventEmitter, auditEvents } from "@/events/audit";
import { AUDIT_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Audit 相关的缓存失效事件处理
 */
export const useAuditCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateEvents = () => queryClient.invalidateQueries({ queryKey: AUDIT_QUERY_KEY_PREFIXES.EVENTS });
  const handleInvalidateEvent = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUDIT_QUERY_KEY_PREFIXES.EVENT, item] });
  const handleInvalidateActorSnapshots = () => queryClient.invalidateQueries({ queryKey: AUDIT_QUERY_KEY_PREFIXES.ACTOR_SNAPSHOTS });
  const handleInvalidateActorSnapshot = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUDIT_QUERY_KEY_PREFIXES.ACTOR_SNAPSHOT, item] });
  const handleInvalidateEventRetentionPolicies = () => queryClient.invalidateQueries({ queryKey: AUDIT_QUERY_KEY_PREFIXES.EVENT_RETENTION_POLICIES });
  const handleInvalidateEventRetentionPolicy = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUDIT_QUERY_KEY_PREFIXES.EVENT_RETENTION_POLICY, item] });
  const handleInvalidateEventSearchIndices = () => queryClient.invalidateQueries({ queryKey: AUDIT_QUERY_KEY_PREFIXES.EVENT_SEARCH_INDICES });
  const handleInvalidateEventSearchIndex = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUDIT_QUERY_KEY_PREFIXES.EVENT_SEARCH_INDEX, item] });
  const handleInvalidateHashChainCheckpoints = () => queryClient.invalidateQueries({ queryKey: AUDIT_QUERY_KEY_PREFIXES.HASH_CHAIN_CHECKPOINTS });
  const handleInvalidateHashChainCheckpoint = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUDIT_QUERY_KEY_PREFIXES.HASH_CHAIN_CHECKPOINT, item] });

  // 注册监听器
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent);
  auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots);
  auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices);
  auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex);
  auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints);
  auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint);

  // 清理监听器
  return () => {
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent);
    auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots);
    auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices);
    auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex);
    auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints);
    auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint);
  };
};
