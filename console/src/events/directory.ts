export const directoryEvents = {
  // User 相关
  INVALIDATE_USER: "DIRECTORY:INVALIDATE_USER",
  INVALIDATE_USERS: "DIRECTORY:INVALIDATE_USERS",

  // Badge 相关
  INVALIDATE_BADGE: "DIRECTORY:INVALIDATE_BADGE",
  INVALIDATE_BADGES: "DIRECTORY:INVALIDATE_BADGES",

  // UserBadge 相关
  INVALIDATE_USER_BADGE: "DIRECTORY:INVALIDATE_USER_BADGE",
  INVALIDATE_USER_BADGES: "DIRECTORY:INVALIDATE_USER_BADGES",

  // UserEducation 相关
  INVALIDATE_USER_EDUCATION: "DIRECTORY:INVALIDATE_USER_EDUCATION",
  INVALIDATE_USER_EDUCATIONS: "DIRECTORY:INVALIDATE_USER_EDUCATIONS",

  // UserEmail 相关
  INVALIDATE_USER_EMAIL: "DIRECTORY:INVALIDATE_USER_EMAIL",
  INVALIDATE_USER_EMAILS: "DIRECTORY:INVALIDATE_USER_EMAILS",

  // UserOccupation 相关
  INVALIDATE_USER_OCCUPATION: "DIRECTORY:INVALIDATE_USER_OCCUPATION",
  INVALIDATE_USER_OCCUPATIONS: "DIRECTORY:INVALIDATE_USER_OCCUPATIONS",

  // UserPhone 相关
  INVALIDATE_USER_PHONE: "DIRECTORY:INVALIDATE_USER_PHONE",
  INVALIDATE_USER_PHONES: "DIRECTORY:INVALIDATE_USER_PHONES",

  // UserPreference 相关
  INVALIDATE_USER_PREFERENCE: "DIRECTORY:INVALIDATE_USER_PREFERENCE",
  INVALIDATE_USER_PREFERENCES: "DIRECTORY:INVALIDATE_USER_PREFERENCES",

  // UserProfile 相关
  INVALIDATE_USER_PROFILE: "DIRECTORY:INVALIDATE_USER_PROFILE",
  INVALIDATE_USER_PROFILES: "DIRECTORY:INVALIDATE_USER_PROFILES",

  // UserAvatar 相关
  INVALIDATE_USER_AVATAR: "DIRECTORY:INVALIDATE_USER_AVATAR",
  INVALIDATE_USER_AVATARS: "DIRECTORY:INVALIDATE_USER_AVATARS",

  // UserImage 相关
  INVALIDATE_USER_IMAGE: "DIRECTORY:INVALIDATE_USER_IMAGE",
  INVALIDATE_USER_IMAGES: "DIRECTORY:INVALIDATE_USER_IMAGES",
} as const;

type DirectoryEvent = (typeof directoryEvents)[keyof typeof directoryEvents];

class DirectoryEventEmitter {
  private listeners: Record<DirectoryEvent, Set<Function>> = {
    [directoryEvents.INVALIDATE_USER]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USERS]: new Set<Function>(),
    [directoryEvents.INVALIDATE_BADGE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_BADGES]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_BADGE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_BADGES]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_EDUCATION]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_EDUCATIONS]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_EMAIL]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_EMAILS]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_OCCUPATION]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_OCCUPATIONS]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PHONE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PHONES]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PREFERENCE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PREFERENCES]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PROFILE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_PROFILES]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_AVATAR]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_AVATARS]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_IMAGE]: new Set<Function>(),
    [directoryEvents.INVALIDATE_USER_IMAGES]: new Set<Function>(),
  };

  on(event: DirectoryEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: DirectoryEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: DirectoryEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const directoryEventEmitter = new DirectoryEventEmitter();
