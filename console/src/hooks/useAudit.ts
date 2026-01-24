import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import { useTranslation } from "react-i18next";

import {
  CreateActorSnapshot,
  CreateEvent,
  CreateEventRetentionPolicy,
  CreateEventSearchIndex,
  CreateHashChainCheckpoint,
  DeleteActorSnapshot,
  DeleteEvent,
  DeleteEventRetentionPolicy,
  DeleteEventSearchIndex,
  DeleteHashChainCheckpoint,
  GetActorSnapshot,
  GetEvent,
  GetEventRetentionPolicy,
  GetEventSearchIndex,
  GetHashChainCheckpoint,
  UpdateEventRetentionPolicy,
} from "@/apis/audit.api";
import type {
  ActorSnapshot,
  CreateActorSnapshotRequest,
  CreateEventRequest,
  CreateEventRetentionPolicyRequest,
  CreateEventSearchIndexRequest,
  CreateHashChainCheckpointRequest,
  Event,
  EventRetentionPolicy,
  EventSearchIndex,
  HashChainCheckpoint,
  UpdateEventRetentionPolicyRequest,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { auditEventEmitter, auditEvents } from "@/events/audit";
import { showError, showSuccess } from "@/stores/modalStore";
import {
  AUDIT_EVENT,
  AUDIT_ACTOR_SNAPSHOT,
  AUDIT_EVENT_RETENTION_POLICY,
  AUDIT_EVENT_SEARCH_INDEX,
  AUDIT_HASH_CHAIN_CHECKPOINT,
} from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== Event 相关 ==========

// 根据 ID 获取事件
export const useEvent = (params: UnifiedQueryParams<Event> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetEvent(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUDIT_EVENT(id), { id }, options);
};

// 创建事件
export const useCreateEvent = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: CreateEventRequest) => {
      return await CreateEvent(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENTS);
      showSuccess(t("event.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("event.createError") + error.message);
    },
  });
};

// 删除事件
export const useDeleteEvent = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEvent(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT, id);
      showSuccess(t("event.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("event.deleteError") + error.message);
    },
  });
};

// ========== ActorSnapshot 相关 ==========

// 根据 ID 获取 Actor Snapshot
export const useActorSnapshot = (params: UnifiedQueryParams<ActorSnapshot> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetActorSnapshot(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUDIT_ACTOR_SNAPSHOT(id), { id }, options);
};

// 创建 Actor Snapshot
export const useCreateActorSnapshot = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: CreateActorSnapshotRequest) => {
      return await CreateActorSnapshot(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS);
      showSuccess(t("actorSnapshot.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("actorSnapshot.createError") + error.message);
    },
  });
};

// 删除 Actor Snapshot
export const useDeleteActorSnapshot = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteActorSnapshot(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, id);
      showSuccess(t("actorSnapshot.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("actorSnapshot.deleteError") + error.message);
    },
  });
};

// ========== EventRetentionPolicy 相关 ==========

// 根据 ID 获取 Event Retention Policy
export const useEventRetentionPolicy = (params: UnifiedQueryParams<EventRetentionPolicy> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetEventRetentionPolicy(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUDIT_EVENT_RETENTION_POLICY(id), { id }, options);
};

// 创建 Event Retention Policy
export const useCreateEventRetentionPolicy = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: CreateEventRetentionPolicyRequest) => {
      return await CreateEventRetentionPolicy(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      showSuccess(t("eventRetentionPolicy.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("eventRetentionPolicy.createError") + error.message);
    },
  });
};

// 更新 Event Retention Policy
export const useUpdateEventRetentionPolicy = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateEventRetentionPolicyRequest }) => {
      return await UpdateEventRetentionPolicy(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, variables.id);
      showSuccess(t("eventRetentionPolicy.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("eventRetentionPolicy.updateError") + error.message);
    },
  });
};

// 删除 Event Retention Policy
export const useDeleteEventRetentionPolicy = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEventRetentionPolicy(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, id);
      showSuccess(t("eventRetentionPolicy.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("eventRetentionPolicy.deleteError") + error.message);
    },
  });
};

// ========== EventSearchIndex 相关 ==========

// 根据 ID 获取 Event Search Index
export const useEventSearchIndex = (params: UnifiedQueryParams<EventSearchIndex> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetEventSearchIndex(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUDIT_EVENT_SEARCH_INDEX(id), { id }, options);
};

// 创建 Event Search Index
export const useCreateEventSearchIndex = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: CreateEventSearchIndexRequest) => {
      return await CreateEventSearchIndex(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES);
      showSuccess(t("eventSearchIndex.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("eventSearchIndex.createError") + error.message);
    },
  });
};

// 删除 Event Search Index
export const useDeleteEventSearchIndex = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEventSearchIndex(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, id);
      showSuccess(t("eventSearchIndex.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("eventSearchIndex.deleteError") + error.message);
    },
  });
};

// ========== HashChainCheckpoint 相关 ==========

// 根据 ID 获取 Hash Chain Checkpoint
export const useHashChainCheckpoint = (params: UnifiedQueryParams<HashChainCheckpoint> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetHashChainCheckpoint(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUDIT_HASH_CHAIN_CHECKPOINT(id), { id }, options);
};

// 创建 Hash Chain Checkpoint
export const useCreateHashChainCheckpoint = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (params: CreateHashChainCheckpointRequest) => {
      return await CreateHashChainCheckpoint(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS);
      showSuccess(t("hashChainCheckpoint.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("hashChainCheckpoint.createError") + error.message);
    },
  });
};

// 删除 Hash Chain Checkpoint
export const useDeleteHashChainCheckpoint = () => {
  const { t } = useTranslation("hooks.audit");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteHashChainCheckpoint(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, id);
      showSuccess(t("hashChainCheckpoint.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("hashChainCheckpoint.deleteError") + error.message);
    },
  });
};
