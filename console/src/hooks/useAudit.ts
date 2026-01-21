import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

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
  

  return useMutation({
    mutationFn: async (params: CreateEventRequest) => {
      return await CreateEvent(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENTS);
      showSuccess("事件创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建事件失败，请稍后重试。" + error.message);
    },
  });
};

// 删除事件
export const useDeleteEvent = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEvent(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT, id);
      showSuccess("事件删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除事件失败，请稍后重试。" + error.message);
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
  

  return useMutation({
    mutationFn: async (params: CreateActorSnapshotRequest) => {
      return await CreateActorSnapshot(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS);
      showSuccess("Actor Snapshot 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Actor Snapshot 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Actor Snapshot
export const useDeleteActorSnapshot = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteActorSnapshot(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, id);
      showSuccess("Actor Snapshot 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Actor Snapshot 失败，请稍后重试。" + error.message);
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
  

  return useMutation({
    mutationFn: async (params: CreateEventRetentionPolicyRequest) => {
      return await CreateEventRetentionPolicy(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      showSuccess("Event Retention Policy 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Event Retention Policy 失败，请稍后重试。" + error.message);
    },
  });
};

// 更新 Event Retention Policy
export const useUpdateEventRetentionPolicy = () => {
  

  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateEventRetentionPolicyRequest }) => {
      return await UpdateEventRetentionPolicy(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, variables.id);
      showSuccess("Event Retention Policy 更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新 Event Retention Policy 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Event Retention Policy
export const useDeleteEventRetentionPolicy = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEventRetentionPolicy(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, id);
      showSuccess("Event Retention Policy 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Event Retention Policy 失败，请稍后重试。" + error.message);
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
  

  return useMutation({
    mutationFn: async (params: CreateEventSearchIndexRequest) => {
      return await CreateEventSearchIndex(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES);
      showSuccess("Event Search Index 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Event Search Index 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Event Search Index
export const useDeleteEventSearchIndex = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteEventSearchIndex(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES);
      auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, id);
      showSuccess("Event Search Index 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Event Search Index 失败，请稍后重试。" + error.message);
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
  

  return useMutation({
    mutationFn: async (params: CreateHashChainCheckpointRequest) => {
      return await CreateHashChainCheckpoint(params);
    },
    onSuccess: () => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS);
      showSuccess("Hash Chain Checkpoint 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Hash Chain Checkpoint 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Hash Chain Checkpoint
export const useDeleteHashChainCheckpoint = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteHashChainCheckpoint(id);
    },
    onSuccess: (_, id) => {
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS);
      auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, id);
      showSuccess("Hash Chain Checkpoint 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Hash Chain Checkpoint 失败，请稍后重试。" + error.message);
    },
  });
};
