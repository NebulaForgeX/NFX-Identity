import type { QueryClient } from "@tanstack/react-query";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import { DIRECTORY_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Directory 相关的缓存失效事件处理
 */
export const useDirectoryCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateUsers = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USERS });
  const handleInvalidateUser = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER, item] });
  const handleInvalidateBadges = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.BADGES });
  const handleInvalidateBadge = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.BADGE, item] });
  const handleInvalidateUserBadges = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_BADGES });
  const handleInvalidateUserBadge = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_BADGE, item] });
  const handleInvalidateUserEducations = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_EDUCATIONS });
  const handleInvalidateUserEducation = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_EDUCATION, item] });
  const handleInvalidateUserEmails = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_EMAILS });
  const handleInvalidateUserEmail = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_EMAIL, item] });
  const handleInvalidateUserOccupations = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_OCCUPATIONS });
  const handleInvalidateUserOccupation = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_OCCUPATION, item] });
  const handleInvalidateUserPhones = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_PHONES });
  const handleInvalidateUserPhone = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_PHONE, item] });
  const handleInvalidateUserPreferences = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_PREFERENCES });
  const handleInvalidateUserPreference = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_PREFERENCE, item] });
  const handleInvalidateUserProfiles = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_QUERY_KEY_PREFIXES.USER_PROFILES });
  const handleInvalidateUserProfile = (item: string) => queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_PROFILE, item] });

  // 注册监听器
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER, handleInvalidateUser);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupations);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile);

  // 清理监听器
  return () => {
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER, handleInvalidateUser);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupation);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile);
  };
};
