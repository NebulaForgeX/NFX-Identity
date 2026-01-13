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
  CreateUserPreference,
  CreateUserProfile,
  DeleteBadge,
  DeleteUser,
  DeleteUserBadge,
  DeleteUserEducation,
  DeleteUserEmail,
  DeleteUserOccupation,
  DeleteUserPhone,
  DeleteUserPreference,
  DeleteUserProfile,
  GetBadge,
  GetBadgeByName,
  GetUser,
  GetUserBadge,
  GetUserByUsername,
  GetUserEducation,
  GetUserEmail,
  GetUserOccupation,
  GetUserPhone,
  GetUserPreference,
  GetUserProfile,
  SetPrimaryUserEmail,
  SetPrimaryUserPhone,
  UpdateBadge,
  UpdateUserEducation,
  UpdateUserEmail,
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
  CreateBadgeRequest,
  CreateUserBadgeRequest,
  CreateUserEducationRequest,
  CreateUserEmailRequest,
  CreateUserOccupationRequest,
  CreateUserPhoneRequest,
  CreateUserPreferenceRequest,
  CreateUserProfileRequest,
  CreateUserRequest,
  UpdateBadgeRequest,
  UpdateUserEducationRequest,
  UpdateUserEmailRequest,
  UpdateUserOccupationRequest,
  UpdateUserPhoneRequest,
  UpdateUserPreferenceRequest,
  UpdateUserProfileRequest,
  UpdateUserStatusRequest,
  UpdateUserUsernameRequest,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import { showError, showSuccess } from "@/stores/modalStore";

// ========== User 相关 ==========

// 根据 ID 获取用户
export const useUser = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUser(params.id);
  },
  "normal",
);

// 根据用户名获取用户
export const useUserByUsername = makeUnifiedQuery(
  async (params: { username: string }) => {
    return await GetUserByUsername(params.username);
  },
  "normal",
);

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
export const useBadge = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetBadge(params.id);
  },
  "normal",
);

// 根据名称获取徽章
export const useBadgeByName = makeUnifiedQuery(
  async (params: { name: string }) => {
    return await GetBadgeByName(params.name);
  },
  "normal",
);

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
export const useUserBadge = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserBadge(params.id);
  },
  "normal",
);

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
export const useUserEducation = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserEducation(params.id);
  },
  "normal",
);

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
export const useUserEmail = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserEmail(params.id);
  },
  "normal",
);

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
export const useUserOccupation = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserOccupation(params.id);
  },
  "normal",
);

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
export const useUserPhone = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserPhone(params.id);
  },
  "normal",
);

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

// 根据 ID 获取用户偏好
export const useUserPreference = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserPreference(params.id);
  },
  "normal",
);

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
export const useUpdateUserPreference = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserPreferenceRequest }) => {
      return await UpdateUserPreference(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCES);
      directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCE, variables.id);
      showSuccess("用户偏好更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户偏好失败，请稍后重试。" + error.message);
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
export const useUserProfile = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetUserProfile(params.id);
  },
  "normal",
);

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
