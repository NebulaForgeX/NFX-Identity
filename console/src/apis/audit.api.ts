// Audit API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "@/types";
import type {
  ActorSnapshot,
  Event,
  EventRetentionPolicy,
  EventSearchIndex,
  HashChainCheckpoint,
} from "@/types";
import type {
  CreateActorSnapshotRequest,
  CreateEventRequest,
  CreateEventRetentionPolicyRequest,
  CreateEventSearchIndexRequest,
  CreateHashChainCheckpointRequest,
  UpdateEventRetentionPolicyRequest,
} from "@/types";

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 事件相关 ==========

// 创建事件
export const CreateEvent = async (params: CreateEventRequest): Promise<Event> => {
  const { data } = await protectedClient.post<DataResponse<Event>>(URL_PATHS.AUDIT.CREATE_EVENT, params);
  return data.data;
};

// 根据 ID 获取事件
export const GetEvent = async (id: string): Promise<Event> => {
  const url = URL_PATHS.AUDIT.GET_EVENT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Event>>(url);
  return data.data;
};

// 删除事件
export const DeleteEvent = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.DELETE_EVENT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Actor Snapshot 相关 ==========

// 创建 Actor Snapshot
export const CreateActorSnapshot = async (params: CreateActorSnapshotRequest): Promise<ActorSnapshot> => {
  const { data } = await protectedClient.post<DataResponse<ActorSnapshot>>(URL_PATHS.AUDIT.CREATE_ACTOR_SNAPSHOT, params);
  return data.data;
};

// 根据 ID 获取 Actor Snapshot
export const GetActorSnapshot = async (id: string): Promise<ActorSnapshot> => {
  const url = URL_PATHS.AUDIT.GET_ACTOR_SNAPSHOT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ActorSnapshot>>(url);
  return data.data;
};

// 删除 Actor Snapshot
export const DeleteActorSnapshot = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.DELETE_ACTOR_SNAPSHOT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Event Retention Policy 相关 ==========

// 创建 Event Retention Policy
export const CreateEventRetentionPolicy = async (params: CreateEventRetentionPolicyRequest): Promise<EventRetentionPolicy> => {
  const { data } = await protectedClient.post<DataResponse<EventRetentionPolicy>>(URL_PATHS.AUDIT.CREATE_EVENT_RETENTION_POLICY, params);
  return data.data;
};

// 根据 ID 获取 Event Retention Policy
export const GetEventRetentionPolicy = async (id: string): Promise<EventRetentionPolicy> => {
  const url = URL_PATHS.AUDIT.GET_EVENT_RETENTION_POLICY.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<EventRetentionPolicy>>(url);
  return data.data;
};

// 更新 Event Retention Policy
export const UpdateEventRetentionPolicy = async (id: string, params: UpdateEventRetentionPolicyRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.UPDATE_EVENT_RETENTION_POLICY.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除 Event Retention Policy
export const DeleteEventRetentionPolicy = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.DELETE_EVENT_RETENTION_POLICY.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Event Search Index 相关 ==========

// 创建 Event Search Index
export const CreateEventSearchIndex = async (params: CreateEventSearchIndexRequest): Promise<EventSearchIndex> => {
  const { data } = await protectedClient.post<DataResponse<EventSearchIndex>>(URL_PATHS.AUDIT.CREATE_EVENT_SEARCH_INDEX, params);
  return data.data;
};

// 根据 ID 获取 Event Search Index
export const GetEventSearchIndex = async (id: string): Promise<EventSearchIndex> => {
  const url = URL_PATHS.AUDIT.GET_EVENT_SEARCH_INDEX.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<EventSearchIndex>>(url);
  return data.data;
};

// 删除 Event Search Index
export const DeleteEventSearchIndex = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.DELETE_EVENT_SEARCH_INDEX.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Hash Chain Checkpoint 相关 ==========

// 创建 Hash Chain Checkpoint
export const CreateHashChainCheckpoint = async (params: CreateHashChainCheckpointRequest): Promise<HashChainCheckpoint> => {
  const { data } = await protectedClient.post<DataResponse<HashChainCheckpoint>>(URL_PATHS.AUDIT.CREATE_HASH_CHAIN_CHECKPOINT, params);
  return data.data;
};

// 根据 ID 获取 Hash Chain Checkpoint
export const GetHashChainCheckpoint = async (id: string): Promise<HashChainCheckpoint> => {
  const url = URL_PATHS.AUDIT.GET_HASH_CHAIN_CHECKPOINT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<HashChainCheckpoint>>(url);
  return data.data;
};

// 删除 Hash Chain Checkpoint
export const DeleteHashChainCheckpoint = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUDIT.DELETE_HASH_CHAIN_CHECKPOINT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
