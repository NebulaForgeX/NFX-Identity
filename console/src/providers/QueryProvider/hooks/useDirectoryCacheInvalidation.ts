import type { QueryClient } from "@tanstack/react-query";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import {
  DIRECTORY_BADGE,
  DIRECTORY_BADGE_LIST,
  DIRECTORY_USER,
  DIRECTORY_USER_BADGE,
  DIRECTORY_USER_BADGE_LIST,
  DIRECTORY_USER_EDUCATION,
  DIRECTORY_USER_EDUCATION_LIST,
  DIRECTORY_USER_EMAIL,
  DIRECTORY_USER_EMAIL_LIST,
  DIRECTORY_USER_IMAGE,
  DIRECTORY_USER_IMAGE_LIST,
  DIRECTORY_USER_LIST,
  DIRECTORY_USER_OCCUPATION,
  DIRECTORY_USER_OCCUPATION_LIST,
  DIRECTORY_USER_PHONE,
  DIRECTORY_USER_PHONE_LIST,
  DIRECTORY_USER_PREFERENCE,
  DIRECTORY_USER_PREFERENCE_LIST,
  DIRECTORY_USER_PROFILE,
  DIRECTORY_USER_PROFILE_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Directory 相关的缓存失效事件处理
 */
export const useDirectoryCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateUsers = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_LIST });
  const handleInvalidateUser = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER(item) });
  const handleInvalidateBadges = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_BADGE_LIST });
  const handleInvalidateBadge = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_BADGE(item) });
  const handleInvalidateUserBadges = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_BADGE_LIST });
  const handleInvalidateUserBadge = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_BADGE(item) });
  const handleInvalidateUserEducations = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_EDUCATION_LIST });
  const handleInvalidateUserEducation = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_EDUCATION(item) });
  const handleInvalidateUserEmails = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_EMAIL_LIST });
  const handleInvalidateUserEmail = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_EMAIL(item) });
  const handleInvalidateUserOccupations = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_OCCUPATION_LIST });
  const handleInvalidateUserOccupation = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_OCCUPATION(item) });
  const handleInvalidateUserPhones = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PHONE_LIST });
  const handleInvalidateUserPhone = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PHONE(item) });
  const handleInvalidateUserPreferences = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PREFERENCE_LIST });
  const handleInvalidateUserPreference = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PREFERENCE(item) });
  const handleInvalidateUserProfiles = () => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PROFILE_LIST });
  const handleInvalidateUserProfile = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_PROFILE(item) });
  const handleInvalidateUserImages = () => {
    queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_IMAGE_LIST });
    queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_IMAGE.getPrefix });
  };
  const handleInvalidateUserImage = (item: string) => queryClient.invalidateQueries({ queryKey: DIRECTORY_USER_IMAGE(item) });

  // 注册监听器
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER, handleInvalidateUser as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupations as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_IMAGES, handleInvalidateUserImages as EventCb);
  directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_IMAGE, handleInvalidateUserImage as EventCb);

  return () => {
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER, handleInvalidateUser as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupation as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_IMAGES, handleInvalidateUserImages as EventCb);
    directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_IMAGE, handleInvalidateUserImage as EventCb);
  };
};
