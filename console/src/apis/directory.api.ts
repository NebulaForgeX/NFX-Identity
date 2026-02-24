// Directory API - 基于 NFX-ID Backend

import type {
  Badge,
  BaseResponse,
  BatchUpdateUserImagesDisplayOrderRequest,
  CreateBadgeRequest,
  CreateOrUpdateUserAvatarRequest,
  CreateUserBadgeRequest,
  CreateUserEducationRequest,
  CreateUserEmailRequest,
  CreateUserImageRequest,
  CreateUserOccupationRequest,
  CreateUserPhoneRequest,
  CreateUserPreferenceRequest,
  CreateUserProfileRequest,
  CreateUserRequest,
  DataResponse,
  UpdateBadgeRequest,
  UpdateUserAvatarImageIDRequest,
  UpdateUserEducationRequest,
  UpdateUserEmailRequest,
  UpdateUserImageDisplayOrderRequest,
  UpdateUserImageImageIDRequest,
  UpdateUserOccupationRequest,
  UpdateUserPhoneRequest,
  UpdateUserPreferenceRequest,
  UpdateUserProfileRequest,
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

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 用户相关 ==========

// 创建用户
export const CreateUser = async (params: CreateUserRequest): Promise<User> => {
  const { data } = await protectedClient.post<DataResponse<User>>(URL_PATHS.DIRECTORY.users, params);
  return data.data;
};

// 根据 ID 获取用户
export const GetUser = async (id: string): Promise<User> => {
  const { data } = await protectedClient.get<DataResponse<User>>(URL_PATHS.DIRECTORY.users.byId(id));
  return data.data;
};

// 根据用户名获取用户
export const GetUserByUsername = async (username: string): Promise<User> => {
  const { data } = await protectedClient.get<DataResponse<User>>(URL_PATHS.DIRECTORY.users.byUsername(username));
  return data.data;
};

// 更新用户状态
export const UpdateUserStatus = async (id: string, params: UpdateUserStatusRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.users.byId(id).status, params);
  return data;
};

// 更新用户名
export const UpdateUserUsername = async (id: string, params: UpdateUserUsernameRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.users.byId(id).username, params);
  return data;
};

// 验证用户
export const VerifyUser = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.users.byId(id).verify);
  return data;
};

// 删除用户
export const DeleteUser = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.users.byId(id));
  return data;
};

// ========== 徽章相关 ==========

// 创建徽章
export const CreateBadge = async (params: CreateBadgeRequest): Promise<Badge> => {
  const { data } = await protectedClient.post<DataResponse<Badge>>(URL_PATHS.DIRECTORY.badges, params);
  return data.data;
};

// 根据 ID 获取徽章
export const GetBadge = async (id: string): Promise<Badge> => {
  const { data } = await protectedClient.get<DataResponse<Badge>>(URL_PATHS.DIRECTORY.badges.byId(id));
  return data.data;
};

// 根据名称获取徽章
export const GetBadgeByName = async (name: string): Promise<Badge> => {
  const { data } = await protectedClient.get<DataResponse<Badge>>(URL_PATHS.DIRECTORY.badges.byName(name));
  return data.data;
};

// 更新徽章
export const UpdateBadge = async (id: string, params: UpdateBadgeRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.badges.byId(id), params);
  return data;
};

// 删除徽章
export const DeleteBadge = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.badges.byId(id));
  return data;
};

// ========== 用户徽章相关 ==========

// 创建用户徽章
export const CreateUserBadge = async (params: CreateUserBadgeRequest): Promise<UserBadge> => {
  const { data } = await protectedClient.post<DataResponse<UserBadge>>(URL_PATHS.DIRECTORY.userBadges, params);
  return data.data;
};

// 根据 ID 获取用户徽章
export const GetUserBadge = async (id: string): Promise<UserBadge> => {
  const { data } = await protectedClient.get<DataResponse<UserBadge>>(URL_PATHS.DIRECTORY.userBadges.byId(id));
  return data.data;
};

// 删除用户徽章
export const DeleteUserBadge = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userBadges.byId(id));
  return data;
};

// ========== 用户教育相关 ==========

// 创建用户教育
export const CreateUserEducation = async (params: CreateUserEducationRequest): Promise<UserEducation> => {
  const { data } = await protectedClient.post<DataResponse<UserEducation>>(
    URL_PATHS.DIRECTORY.userEducations,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户教育
export const GetUserEducation = async (id: string): Promise<UserEducation> => {
  const { data } = await protectedClient.get<DataResponse<UserEducation>>(URL_PATHS.DIRECTORY.userEducations.byId(id));
  return data.data;
};

// 根据用户ID获取用户教育列表
export const GetUserEducationsByUserID = async (userId: string): Promise<UserEducation[]> => {
  const { data } = await protectedClient.get<DataResponse<UserEducation[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).userEducations,
  );
  return data.data;
};

// 更新用户教育
export const UpdateUserEducation = async (id: string, params: UpdateUserEducationRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userEducations.byId(id), params);
  return data;
};

// 删除用户教育
export const DeleteUserEducation = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userEducations.byId(id));
  return data;
};

// ========== 用户邮箱相关 ==========

// 创建用户邮箱
export const CreateUserEmail = async (params: CreateUserEmailRequest): Promise<UserEmail> => {
  const { data } = await protectedClient.post<DataResponse<UserEmail>>(URL_PATHS.DIRECTORY.userEmails, params);
  return data.data;
};

// 根据 ID 获取用户邮箱
export const GetUserEmail = async (id: string): Promise<UserEmail> => {
  const { data } = await protectedClient.get<DataResponse<UserEmail>>(URL_PATHS.DIRECTORY.userEmails.byId(id));
  return data.data;
};

// 根据用户ID获取用户邮箱列表
export const GetUserEmailsByUserID = async (userId: string): Promise<UserEmail[]> => {
  const { data } = await protectedClient.get<DataResponse<UserEmail[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).userEmails,
  );
  return data.data;
};

// 更新用户邮箱
export const UpdateUserEmail = async (id: string, params: UpdateUserEmailRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userEmails.byId(id), params);
  return data;
};

// 设置主邮箱
export const SetPrimaryUserEmail = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userEmails.setPrimary(id));
  return data;
};

// 验证用户邮箱
export const VerifyUserEmail = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userEmails.verify(id));
  return data;
};

// 删除用户邮箱
export const DeleteUserEmail = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userEmails.byId(id));
  return data;
};

// ========== 用户职业相关 ==========

// 创建用户职业
export const CreateUserOccupation = async (params: CreateUserOccupationRequest): Promise<UserOccupation> => {
  const { data } = await protectedClient.post<DataResponse<UserOccupation>>(
    URL_PATHS.DIRECTORY.userOccupations,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户职业
export const GetUserOccupation = async (id: string): Promise<UserOccupation> => {
  const { data } = await protectedClient.get<DataResponse<UserOccupation>>(
    URL_PATHS.DIRECTORY.userOccupations.byId(id),
  );
  return data.data;
};

// 根据用户ID获取用户职业列表
export const GetUserOccupationsByUserID = async (userId: string): Promise<UserOccupation[]> => {
  const { data } = await protectedClient.get<DataResponse<UserOccupation[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).userOccupations,
  );
  return data.data;
};

// 更新用户职业
export const UpdateUserOccupation = async (id: string, params: UpdateUserOccupationRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userOccupations.byId(id), params);
  return data;
};

// 删除用户职业
export const DeleteUserOccupation = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userOccupations.byId(id));
  return data;
};

// ========== 用户电话相关 ==========

// 创建用户电话
export const CreateUserPhone = async (params: CreateUserPhoneRequest): Promise<UserPhone> => {
  const { data } = await protectedClient.post<DataResponse<UserPhone>>(URL_PATHS.DIRECTORY.userPhones, params);
  return data.data;
};

// 根据 ID 获取用户电话
export const GetUserPhone = async (id: string): Promise<UserPhone> => {
  const { data } = await protectedClient.get<DataResponse<UserPhone>>(URL_PATHS.DIRECTORY.userPhones.byId(id));
  return data.data;
};

// 根据用户ID获取用户电话列表
export const GetUserPhonesByUserID = async (userId: string): Promise<UserPhone[]> => {
  const { data } = await protectedClient.get<DataResponse<UserPhone[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).userPhones,
  );
  return data.data;
};

// 更新用户电话
export const UpdateUserPhone = async (id: string, params: UpdateUserPhoneRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userPhones.byId(id), params);
  return data;
};

// 设置主电话
export const SetPrimaryUserPhone = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userPhones.setPrimary(id));
  return data;
};

// 验证用户电话
export const VerifyUserPhone = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userPhones.verify(id));
  return data;
};

// 删除用户电话
export const DeleteUserPhone = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userPhones.byId(id));
  return data;
};

// ========== 用户偏好相关 ==========

// 创建用户偏好
export const CreateUserPreference = async (params: CreateUserPreferenceRequest): Promise<UserPreference> => {
  const { data } = await protectedClient.post<DataResponse<UserPreference>>(
    URL_PATHS.DIRECTORY.userPreferences,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户偏好
export const GetUserPreference = async (id: string): Promise<UserPreference> => {
  const { data } = await protectedClient.get<DataResponse<UserPreference>>(
    URL_PATHS.DIRECTORY.userPreferences.byId(id),
  );
  return data.data;
};

// 更新用户偏好
export const UpdateUserPreference = async (id: string, params: UpdateUserPreferenceRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userPreferences.byId(id), params);
  return data;
};

// 删除用户偏好
export const DeleteUserPreference = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userPreferences.byId(id));
  return data;
};

// ========== 用户资料相关 ==========

// 创建用户资料
export const CreateUserProfile = async (params: CreateUserProfileRequest): Promise<UserProfile> => {
  const { data } = await protectedClient.post<DataResponse<UserProfile>>(
    URL_PATHS.DIRECTORY.userProfiles,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户资料
export const GetUserProfile = async (id: string): Promise<UserProfile> => {
  const { data } = await protectedClient.get<DataResponse<UserProfile>>(URL_PATHS.DIRECTORY.userProfiles.byId(id));
  return data.data;
};

// 更新用户资料
export const UpdateUserProfile = async (id: string, params: UpdateUserProfileRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userProfiles.byId(id), params);
  return data;
};

// 删除用户资料
export const DeleteUserProfile = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userProfiles.byId(id));
  return data;
};

// ========== 用户头像相关 ==========

// 创建或更新用户头像
export const CreateOrUpdateUserAvatar = async (params: CreateOrUpdateUserAvatarRequest): Promise<UserAvatar> => {
  const { data } = await protectedClient.post<DataResponse<UserAvatar>>(
    URL_PATHS.DIRECTORY.userAvatars.byUserId(params.userId),
    params,
  );
  return data.data;
};

// 根据用户ID获取用户头像（后端 404 时返回 err_code + message，axios 抛错，不返回 null）
export const GetUserAvatar = async (userId: string): Promise<UserAvatar> => {
  const { data } = await protectedClient.get<DataResponse<UserAvatar>>(
    URL_PATHS.DIRECTORY.userAvatars.byUserId(userId),
  );
  return data.data;
};

// 更新用户头像
export const UpdateUserAvatar = async (
  userId: string,
  params: UpdateUserAvatarImageIDRequest,
): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.DIRECTORY.userAvatars.byUserId(userId),
    params,
  );
  return data;
};

// 删除用户头像
export const DeleteUserAvatar = async (userId: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.DIRECTORY.userAvatars.byUserId(userId),
  );
  return data;
};

// ========== 用户图片相关 ==========

// 创建用户图片
export const CreateUserImage = async (params: CreateUserImageRequest): Promise<UserImage> => {
  const { data } = await protectedClient.post<DataResponse<UserImage>>(
    URL_PATHS.DIRECTORY.users.byId(params.userId).userImages,
    params,
  );
  return data.data;
};

// 根据ID获取用户图片
export const GetUserImage = async (id: string): Promise<UserImage> => {
  const { data } = await protectedClient.get<DataResponse<UserImage>>(URL_PATHS.DIRECTORY.userImages.byId(id));
  return data.data;
};

// 根据用户ID获取用户图片列表
export const GetUserImagesByUserID = async (userId: string): Promise<UserImage[]> => {
  const { data } = await protectedClient.get<DataResponse<UserImage[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).userImages,
  );
  return data.data;
};

// 获取用户当前图片（display_order = 0）（后端 404 时返回 err_code + message，axios 抛错，不返回 null）
export const GetCurrentUserImageByUserID = async (userId: string): Promise<UserImage> => {
  const { data } = await protectedClient.get<DataResponse<UserImage>>(
    URL_PATHS.DIRECTORY.users.byId(userId).currentImage,
  );
  return data.data;
};

// 更新用户图片
export const UpdateUserImage = async (id: string, params: UpdateUserImageImageIDRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.DIRECTORY.userImages.byId(id), params);
  return data;
};

// 设置主图（背景图）
export const SetPrimaryUserImage = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userImages.setPrimary(id));
  return data;
};

// 更新用户图片显示顺序
export const UpdateUserImageDisplayOrder = async (
  id: string,
  params: UpdateUserImageDisplayOrderRequest,
): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(URL_PATHS.DIRECTORY.userImages.displayOrder(id), params);
  return data;
};

// 批量更新用户图片显示顺序
export const UpdateUserImagesDisplayOrderBatch = async (
  userId: string,
  params: BatchUpdateUserImagesDisplayOrderRequest,
): Promise<UserImage[]> => {
  const { data } = await protectedClient.patch<DataResponse<UserImage[]>>(
    URL_PATHS.DIRECTORY.users.byId(userId).imagesDisplayOrder,
    params,
  );
  return data.data;
};

// 删除用户图片
export const DeleteUserImage = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.DIRECTORY.userImages.byId(id));
  return data;
};
