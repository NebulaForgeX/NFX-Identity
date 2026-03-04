// Audit API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "nfx-ui/types";
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

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 事件相关 ==========

// 创建事件
export const CreateEvent = async (params: CreateEventRequest): Promise<Event> => {
  const { data } = await protectedClient.post<DataResponse<Event>>(
    URL_PATHS.AUDIT.events,
    params,
  );
  return data.data;
};

// 根据 ID 获取事件
export const GetEvent = async (id: string): Promise<Event> => {
  const { data } = await protectedClient.get<DataResponse<Event>>(
    URL_PATHS.AUDIT.events.byId(id),
  );
  return data.data;
};

// 删除事件
export const DeleteEvent = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.AUDIT.events.byId(id),
  );
  return data;
};

// ========== Actor Snapshot 相关 ==========

// 创建 Actor Snapshot
export const CreateActorSnapshot = async (params: CreateActorSnapshotRequest): Promise<ActorSnapshot> => {
  const { data } = await protectedClient.post<DataResponse<ActorSnapshot>>(
    URL_PATHS.AUDIT.actorSnapshots,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Actor Snapshot
export const GetActorSnapshot = async (id: string): Promise<ActorSnapshot> => {
  const { data } = await protectedClient.get<DataResponse<ActorSnapshot>>(
    URL_PATHS.AUDIT.actorSnapshots.byId(id),
  );
  return data.data;
};

// 删除 Actor Snapshot
export const DeleteActorSnapshot = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.AUDIT.actorSnapshots.byId(id),
  );
  return data;
};

// ========== Event Retention Policy 相关 ==========

// 创建 Event Retention Policy
export const CreateEventRetentionPolicy = async (
  params: CreateEventRetentionPolicyRequest,
): Promise<EventRetentionPolicy> => {
  const { data } = await protectedClient.post<DataResponse<EventRetentionPolicy>>(
    URL_PATHS.AUDIT.eventRetentionPolicies,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Event Retention Policy
export const GetEventRetentionPolicy = async (id: string): Promise<EventRetentionPolicy> => {
  const { data } = await protectedClient.get<DataResponse<EventRetentionPolicy>>(
    URL_PATHS.AUDIT.eventRetentionPolicies.byId(id),
  );
  return data.data;
};

// 更新 Event Retention Policy
export const UpdateEventRetentionPolicy = async (
  id: string,
  params: UpdateEventRetentionPolicyRequest,
): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.AUDIT.eventRetentionPolicies.byId(id),
    params,
  );
  return data;
};

// 删除 Event Retention Policy
export const DeleteEventRetentionPolicy = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.AUDIT.eventRetentionPolicies.byId(id),
  );
  return data;
};

// ========== Event Search Index 相关 ==========

// 创建 Event Search Index
export const CreateEventSearchIndex = async (params: CreateEventSearchIndexRequest): Promise<EventSearchIndex> => {
  const { data } = await protectedClient.post<DataResponse<EventSearchIndex>>(
    URL_PATHS.AUDIT.eventSearchIndex,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Event Search Index
export const GetEventSearchIndex = async (id: string): Promise<EventSearchIndex> => {
  const { data } = await protectedClient.get<DataResponse<EventSearchIndex>>(
    URL_PATHS.AUDIT.eventSearchIndex.byId(id),
  );
  return data.data;
};

// 删除 Event Search Index
export const DeleteEventSearchIndex = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.AUDIT.eventSearchIndex.byId(id),
  );
  return data;
};

// ========== Hash Chain Checkpoint 相关 ==========

// 创建 Hash Chain Checkpoint
export const CreateHashChainCheckpoint = async (
  params: CreateHashChainCheckpointRequest,
): Promise<HashChainCheckpoint> => {
  const { data } = await protectedClient.post<DataResponse<HashChainCheckpoint>>(
    URL_PATHS.AUDIT.hashChainCheckpoints,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Hash Chain Checkpoint
export const GetHashChainCheckpoint = async (id: string): Promise<HashChainCheckpoint> => {
  const { data } = await protectedClient.get<DataResponse<HashChainCheckpoint>>(
    URL_PATHS.AUDIT.hashChainCheckpoints.byId(id),
  );
  return data.data;
};

// 删除 Hash Chain Checkpoint
export const DeleteHashChainCheckpoint = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.AUDIT.hashChainCheckpoints.byId(id),
  );
  return data;
};
