import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  CreateBadge,
  CreateUser,
  CreateUserBadge,
  CreateUserEducation,
  CreateUserEmail,
  CreateUserOccupation,
  CreateUserPhone,
  CreateOrUpdateUserAvatar,
  CreateUserImage,
  CreateUserPreference,
  CreateUserProfile,
  DeleteBadge,
  DeleteUser,
  DeleteUserAvatar,
  DeleteUserBadge,
  DeleteUserEducation,
  DeleteUserEmail,
  DeleteUserImage,
  DeleteUserOccupation,
  DeleteUserPhone,
  DeleteUserPreference,
  DeleteUserProfile,
  GetBadge,
  GetBadgeByName,
  GetCurrentUserImageByUserID,
  GetUser,
  GetUserAvatar,
  GetUserBadge,
  GetUserByUsername,
  GetUserEducation,
  GetUserEducationsByUserID,
  GetUserEmail,
  GetUserEmailsByUserID,
  GetUserImage,
  GetUserImagesByUserID,
  GetUserOccupation,
  GetUserOccupationsByUserID,
  GetUserPhone,
  GetUserPhonesByUserID,
  GetUserPreference,
  GetUserProfile,
  SetPrimaryUserEmail,
  SetPrimaryUserPhone,
  UpdateBadge,
  UpdateUserAvatar,
  UpdateUserEducation,
  UpdateUserEmail,
  UpdateUserImage,
  UpdateUserImageDisplayOrder,
  UpdateUserOccupation,
  UpdateUserPhone,
  UpdateUserPreference,
  UpdateUserProfile,
  UpdateUserStatus,
  UpdateUserUsername,
  VerifyUser,
  VerifyUserEmail,
  VerifyUserPhone,
} from "@/apis/directory.api";
import type {
  Badge,
  CreateBadgeRequest,
  CreateUserBadgeRequest,
  CreateUserEducationRequest,
  CreateUserEmailRequest,
  CreateUserOccupationRequest,
  CreateUserPhoneRequest,
  CreateOrUpdateUserAvatarRequest,
  CreateUserImageRequest,
  CreateUserPreferenceRequest,
  CreateUserProfileRequest,
  CreateUserRequest,
  UpdateBadgeRequest,
  UpdateUserEducationRequest,
  UpdateUserEmailRequest,
  UpdateUserImageDisplayOrderRequest,
  UpdateUserImageImageIDRequest,
  UpdateUserOccupationRequest,
  UpdateUserPhoneRequest,
  UpdateUserPreferenceRequest,
  UpdateUserProfileRequest,
  UpdateUserAvatarImageIDRequest,
  UpdateUserStatusRequest,
  UpdateUserUsernameRequest,
  User,
  UserAvatar,
  UserBadge,
  UserEducation,
  UserEmail,
  UserImage,
  UserOccupation,
  UserPhone,
  UserPreference,
  UserProfile,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import { showError, showSuccess } from "@/stores/modalStore";
import {
  DIRECTORY_USER,
  DIRECTORY_BADGE,
  DIRECTORY_USER_BADGE,
  DIRECTORY_USER_EDUCATION,
  DIRECTORY_USER_EDUCATION_LIST,
  DIRECTORY_USER_EMAIL,
  DIRECTORY_USER_EMAIL_LIST,
  DIRECTORY_USER_OCCUPATION,
  DIRECTORY_USER_OCCUPATION_LIST,
  DIRECTORY_USER_PHONE,
  DIRECTORY_USER_PHONE_LIST,
  DIRECTORY_USER_PREFERENCE,
  DIRECTORY_USER_PROFILE,
  DIRECTORY_USER_AVATAR,
  DIRECTORY_USER_IMAGE,
  DIRECTORY_USER_IMAGE_LIST,
} from "@/constants";
import type { UnifiedQueryParams, suspenseUnifiedQueryOptions, SuspenseUnifiedQueryOptions } from "./core/type";

// ========== User 相关 ==========

// 根据 ID 获取用户
export const useUser = (params: UnifiedQueryParams<User> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUser(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER(id), { id }, options);
};

// 根据用户名获取用户
export const useUserByUsername = (params: UnifiedQueryParams<User> & { username: string }) => {
  const { username, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { username: string }) => {
      return await GetUserByUsername(params.username);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER(username), { username }, options);
};

// 创建用户
export const useCreateUser = () => {
  return useMutation({
    mutationFn: async (params: CreateUserRequest) => {
      return await CreateUser(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
      showSuccess("用户创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户状态
export const useUpdateUserStatus = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserStatusRequest }) => {
      return await UpdateUserStatus(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER, variables.id);
      showSuccess("用户状态更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户状态失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户名
export const useUpdateUserUsername = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserUsernameRequest }) => {
      return await UpdateUserUsername(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER, variables.id);
      showSuccess("用户名更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户名失败，请稍后重试。" + error.message);
    },
  });
};

// 验证用户
export const useVerifyUser = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await VerifyUser(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER, id);
      showSuccess("用户验证成功！");
    },
    onError: (error: AxiosError) => {
      showError("验证用户失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户
export const useDeleteUser = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUser(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER, id);
      showSuccess("用户删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Badge 相关 ==========

// 根据 ID 获取徽章
export const useBadge = (params: UnifiedQueryParams<Badge> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetBadge(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_BADGE(id), { id }, options);
};

// 根据名称获取徽章
export const useBadgeByName = (params: UnifiedQueryParams<Badge> & { name: string }) => {
  const { name, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { name: string }) => {
      return await GetBadgeByName(params.name);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_BADGE(name), { name }, options);
};

// 创建徽章
export const useCreateBadge = () => {
  return useMutation({
    mutationFn: async (params: CreateBadgeRequest) => {
      return await CreateBadge(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
      showSuccess("徽章创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建徽章失败，请稍后重试。" + error.message);
    },
  });
};

// 更新徽章
export const useUpdateBadge = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateBadgeRequest }) => {
      return await UpdateBadge(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGE, variables.id);
      showSuccess("徽章更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新徽章失败，请稍后重试。" + error.message);
    },
  });
};

// 删除徽章
export const useDeleteBadge = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteBadge(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGE, id);
      showSuccess("徽章删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除徽章失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserBadge 相关 ==========

// 根据 ID 获取用户徽章
export const useUserBadge = (params: UnifiedQueryParams<UserBadge> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserBadge(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_BADGE(id), { id }, options);
};

// 创建用户徽章
export const useCreateUserBadge = () => {
  return useMutation({
    mutationFn: async (params: CreateUserBadgeRequest) => {
      return await CreateUserBadge(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_BADGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
      showSuccess("用户徽章创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户徽章失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户徽章
export const useDeleteUserBadge = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserBadge(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_BADGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_BADGE, id);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
      showSuccess("用户徽章删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户徽章失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserEducation 相关 ==========

// 根据 ID 获取用户教育
export const useUserEducation = (params: UnifiedQueryParams<UserEducation> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserEducation(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_EDUCATION(id), { id }, options);
};

// 根据用户ID获取用户教育列表
export const useUserEducationsByUserID = (params: UnifiedQueryParams<UserEducation[]> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      return await GetUserEducationsByUserID(params.userId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery([...DIRECTORY_USER_EDUCATION_LIST, userId], { userId }, options);
};

// 创建用户教育
export const useCreateUserEducation = () => {
  return useMutation({
    mutationFn: async (params: CreateUserEducationRequest) => {
      return await CreateUserEducation(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATIONS);
      showSuccess("用户教育创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户教育失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户教育
export const useUpdateUserEducation = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserEducationRequest }) => {
      return await UpdateUserEducation(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATIONS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATION, variables.id);
      showSuccess("用户教育更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户教育失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户教育
export const useDeleteUserEducation = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserEducation(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATIONS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATION, id);
      showSuccess("用户教育删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户教育失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserEmail 相关 ==========

// 根据 ID 获取用户邮箱
export const useUserEmail = (params: UnifiedQueryParams<UserEmail> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserEmail(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_EMAIL(id), { id }, options);
};

// 根据用户ID获取用户邮箱列表
export const useUserEmailsByUserID = (params: UnifiedQueryParams<UserEmail[]> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      return await GetUserEmailsByUserID(params.userId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery([...DIRECTORY_USER_EMAIL_LIST, userId], { userId }, options);
};

// 创建用户邮箱
export const useCreateUserEmail = () => {
  return useMutation({
    mutationFn: async (params: CreateUserEmailRequest) => {
      return await CreateUserEmail(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
      showSuccess("用户邮箱创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户邮箱失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户邮箱
export const useUpdateUserEmail = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserEmailRequest }) => {
      return await UpdateUserEmail(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAIL, variables.id);
      showSuccess("用户邮箱更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户邮箱失败，请稍后重试。" + error.message);
    },
  });
};

// 设置主邮箱
export const useSetPrimaryUserEmail = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await SetPrimaryUserEmail(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAIL, id);
      showSuccess("主邮箱设置成功！");
    },
    onError: (error: AxiosError) => {
      showError("设置主邮箱失败，请稍后重试。" + error.message);
    },
  });
};

// 验证用户邮箱
export const useVerifyUserEmail = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await VerifyUserEmail(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAIL, id);
      showSuccess("用户邮箱验证成功！");
    },
    onError: (error: AxiosError) => {
      showError("验证用户邮箱失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户邮箱
export const useDeleteUserEmail = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserEmail(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAIL, id);
      showSuccess("用户邮箱删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户邮箱失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserOccupation 相关 ==========

// 根据 ID 获取用户职业
export const useUserOccupation = (params: UnifiedQueryParams<UserOccupation> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserOccupation(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_OCCUPATION(id), { id }, options);
};

// 根据用户ID获取用户职业列表
export const useUserOccupationsByUserID = (params: UnifiedQueryParams<UserOccupation[]> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      return await GetUserOccupationsByUserID(params.userId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery([...DIRECTORY_USER_OCCUPATION_LIST, userId], { userId }, options);
};

// 创建用户职业
export const useCreateUserOccupation = () => {
  return useMutation({
    mutationFn: async (params: CreateUserOccupationRequest) => {
      return await CreateUserOccupation(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATIONS);
      showSuccess("用户职业创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户职业失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户职业
export const useUpdateUserOccupation = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserOccupationRequest }) => {
      return await UpdateUserOccupation(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATIONS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATION, variables.id);
      showSuccess("用户职业更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户职业失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户职业
export const useDeleteUserOccupation = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserOccupation(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATIONS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATION, id);
      showSuccess("用户职业删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户职业失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserPhone 相关 ==========

// 根据 ID 获取用户电话
export const useUserPhone = (params: UnifiedQueryParams<UserPhone> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserPhone(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_PHONE(id), { id }, options);
};

// 根据用户ID获取用户电话列表
export const useUserPhonesByUserID = (params: UnifiedQueryParams<UserPhone[]> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      return await GetUserPhonesByUserID(params.userId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery([...DIRECTORY_USER_PHONE_LIST, userId], { userId }, options);
};

// 创建用户电话
export const useCreateUserPhone = () => {
  return useMutation({
    mutationFn: async (params: CreateUserPhoneRequest) => {
      return await CreateUserPhone(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
      showSuccess("用户电话创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户电话失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户电话
export const useUpdateUserPhone = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserPhoneRequest }) => {
      return await UpdateUserPhone(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONE, variables.id);
      showSuccess("用户电话更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户电话失败，请稍后重试。" + error.message);
    },
  });
};

// 设置主电话
export const useSetPrimaryUserPhone = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await SetPrimaryUserPhone(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONE, id);
      showSuccess("主电话设置成功！");
    },
    onError: (error: AxiosError) => {
      showError("设置主电话失败，请稍后重试。" + error.message);
    },
  });
};

// 验证用户电话
export const useVerifyUserPhone = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await VerifyUserPhone(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONE, id);
      showSuccess("用户电话验证成功！");
    },
    onError: (error: AxiosError) => {
      showError("验证用户电话失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户电话
export const useDeleteUserPhone = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserPhone(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONE, id);
      showSuccess("用户电话删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户电话失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserPreference 相关 ==========

// 根据 ID 获取用户偏好（Suspense 模式）
export const useUserPreference = (params: UnifiedQueryParams<UserPreference> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserPreference(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_PREFERENCE(id), { id }, options);
};

// 根据 ID 获取用户偏好（普通模式，支持 enabled 选项）
export const useUserPreferenceNormal = (params: {
  id: string;
  options?: suspenseUnifiedQueryOptions<UserPreference>;
  postProcess?: (data: UserPreference) => void;
}) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserPreference(params.id);
    },
    "normal",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_PREFERENCE(id), { id }, options);
};

// 创建用户偏好
export const useCreateUserPreference = () => {
  return useMutation({
    mutationFn: async (params: CreateUserPreferenceRequest) => {
      return await CreateUserPreference(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCES);
      showSuccess("用户偏好创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户偏好失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户偏好
export const useUpdateUserPreference = (options?: { silent?: boolean }) => {
  const silent = options?.silent ?? false;
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserPreferenceRequest }) => {
      return await UpdateUserPreference(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCE, variables.id);
      if (!silent) {
        showSuccess("用户偏好更新成功！");
      }
    },
    onError: (error: AxiosError) => {
      if (!silent) {
        showError("更新用户偏好失败，请稍后重试。" + error.message);
      }
    },
  });
};

// 删除用户偏好
export const useDeleteUserPreference = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserPreference(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCE, id);
      showSuccess("用户偏好删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户偏好失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserProfile 相关 ==========

// 根据 ID 获取用户资料
export const useUserProfile = (params: UnifiedQueryParams<UserProfile> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserProfile(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_PROFILE(id), { id }, options);
};

// 创建用户资料
export const useCreateUserProfile = () => {
  return useMutation({
    mutationFn: async (params: CreateUserProfileRequest) => {
      return await CreateUserProfile(params);
    },
    onSuccess: () => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILES);
      showSuccess("用户资料创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户资料失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户资料
export const useUpdateUserProfile = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserProfileRequest }) => {
      return await UpdateUserProfile(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILE, variables.id);
      showSuccess("用户资料更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户资料失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户资料
export const useDeleteUserProfile = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserProfile(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILE, id);
      showSuccess("用户资料删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户资料失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserAvatar 相关 ==========

// 根据用户ID获取用户头像
export const useUserAvatar = (params: UnifiedQueryParams<UserAvatar> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      const avatar = await GetUserAvatar(params.userId);
      // 如果返回 null，返回一个默认的空对象，避免 suspense 错误
      return avatar || { userId: params.userId, imageId: "", createdAt: "", updatedAt: "" };
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_AVATAR(userId), { userId }, options);
};

// 创建或更新用户头像
export const useCreateOrUpdateUserAvatar = () => {
  return useMutation({
    mutationFn: async (params: CreateOrUpdateUserAvatarRequest) => {
      return await CreateOrUpdateUserAvatar(params);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATARS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATAR, variables.userId);
      showSuccess("用户头像设置成功！");
    },
    onError: (error: AxiosError) => {
      showError("设置用户头像失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户头像
export const useUpdateUserAvatar = () => {
  return useMutation({
    mutationFn: async (params: { userId: string; data: UpdateUserAvatarImageIDRequest }) => {
      return await UpdateUserAvatar(params.userId, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATARS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATAR, variables.userId);
      showSuccess("用户头像更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户头像失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户头像
export const useDeleteUserAvatar = () => {
  return useMutation({
    mutationFn: async (userId: string) => {
      return await DeleteUserAvatar(userId);
    },
    onSuccess: (_, userId) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATARS);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATAR, userId);
      showSuccess("用户头像删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户头像失败，请稍后重试。" + error.message);
    },
  });
};

// ========== UserImage 相关 ==========

// 根据ID获取用户图片
export const useUserImage = (params: UnifiedQueryParams<UserImage> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserImage(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_IMAGE(id), { id }, options);
};

// 根据用户ID获取用户图片列表
export const useUserImagesByUserID = (params: UnifiedQueryParams<UserImage[]> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      return await GetUserImagesByUserID(params.userId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_IMAGE_LIST, { userId }, options);
};

// 获取用户当前图片（display_order = 0）
export const useCurrentUserImageByUserID = (params: UnifiedQueryParams<UserImage> & { userId: string }) => {
  const { userId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { userId: string }) => {
      const image = await GetCurrentUserImageByUserID(params.userId);
      // 如果返回 null，返回一个默认的空对象，避免 suspense 错误
      return image || { id: "", userId: params.userId, imageId: "", displayOrder: 0, createdAt: "", updatedAt: "" };
    },
    "suspense",
    postProcess,
  );
  return makeQuery(DIRECTORY_USER_IMAGE(`current-${userId}`), { userId }, options as SuspenseUnifiedQueryOptions<UserImage> | undefined);
};

// 创建用户图片
export const useCreateUserImage = () => {
  return useMutation({
    mutationFn: async (params: CreateUserImageRequest) => {
      return await CreateUserImage(params);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGE, variables.userId);
      showSuccess("用户图片创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户图片失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户图片
export const useUpdateUserImage = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserImageImageIDRequest }) => {
      return await UpdateUserImage(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGE, variables.id);
      showSuccess("用户图片更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户图片失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户图片显示顺序
export const useUpdateUserImageDisplayOrder = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserImageDisplayOrderRequest }) => {
      return await UpdateUserImageDisplayOrder(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGE, variables.id);
      showSuccess("用户图片顺序更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户图片顺序失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户图片
export const useDeleteUserImage = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserImage(id);
    },
    onSuccess: (_, id) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGE, id);
      showSuccess("用户图片删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户图片失败，请稍后重试。" + error.message);
    },
  });
};
