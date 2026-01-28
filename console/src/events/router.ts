import { ROUTES } from "@/types/navigation";

export const routerEvents = {
  // 导航事件
  NAVIGATE: "ROUTER:NAVIGATE",
  NAVIGATE_REPLACE: "ROUTER:NAVIGATE_REPLACE",
  NAVIGATE_BACK: "ROUTER:NAVIGATE_BACK",
  NAVIGATE_FORWARD: "ROUTER:NAVIGATE_FORWARD",
  // 特定路由跳转事件
  NAVIGATE_TO_LOGIN: "ROUTER:NAVIGATE_TO_LOGIN",
  NAVIGATE_TO_DASHBOARD: "ROUTER:NAVIGATE_TO_DASHBOARD",
  NAVIGATE_TO_HOME: "ROUTER:NAVIGATE_TO_HOME",
  NAVIGATE_TO_PROFILE: "ROUTER:NAVIGATE_TO_PROFILE",
  NAVIGATE_TO_EDIT_PROFILE: "ROUTER:NAVIGATE_TO_EDIT_PROFILE",
  NAVIGATE_TO_ACCOUNT_SECURITY: "ROUTER:NAVIGATE_TO_ACCOUNT_SECURITY",
  NAVIGATE_TO_USER_SECURITY: "ROUTER:NAVIGATE_TO_USER_SECURITY",
  NAVIGATE_TO_ADD_EDUCATION: "ROUTER:NAVIGATE_TO_ADD_EDUCATION",
  NAVIGATE_TO_ADD_OCCUPATION: "ROUTER:NAVIGATE_TO_ADD_OCCUPATION",
  NAVIGATE_TO_EDIT_EDUCATION: "ROUTER:NAVIGATE_TO_EDIT_EDUCATION",
  NAVIGATE_TO_EDIT_OCCUPATION: "ROUTER:NAVIGATE_TO_EDIT_OCCUPATION",
  NAVIGATE_TO_EDIT_PREFERENCE: "ROUTER:NAVIGATE_TO_EDIT_PREFERENCE",
} as const;

type RouterEvent = (typeof routerEvents)[keyof typeof routerEvents];

interface NavigatePayload {
  to: string;
  replace?: boolean;
  state?: unknown;
}

class RouterEventEmitter {
  private listeners: Record<RouterEvent, Set<Function>> = {
    [routerEvents.NAVIGATE]: new Set<Function>(),
    [routerEvents.NAVIGATE_REPLACE]: new Set<Function>(),
    [routerEvents.NAVIGATE_BACK]: new Set<Function>(),
    [routerEvents.NAVIGATE_FORWARD]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_LOGIN]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_DASHBOARD]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_HOME]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_PROFILE]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_EDIT_PROFILE]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_ACCOUNT_SECURITY]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_USER_SECURITY]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_ADD_EDUCATION]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_ADD_OCCUPATION]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_EDIT_EDUCATION]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_EDIT_OCCUPATION]: new Set<Function>(),
    [routerEvents.NAVIGATE_TO_EDIT_PREFERENCE]: new Set<Function>(),
  };

  on(event: RouterEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: RouterEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: RouterEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }

  // 便捷方法：导航到指定路径
  navigate(payload: NavigatePayload) {
    this.emit(routerEvents.NAVIGATE, payload);
  }

  // 便捷方法：替换当前路径
  navigateReplace(to: string, state?: unknown) {
    this.emit(routerEvents.NAVIGATE_REPLACE, { to, state });
  }

  // 便捷方法：导航到登录页
  navigateToLogin() {
    this.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }

  // 便捷方法：导航到仪表板
  navigateToDashboard() {
    this.emit(routerEvents.NAVIGATE_TO_DASHBOARD);
  }

  // 便捷方法：导航到首页
  navigateToHome() {
    this.emit(routerEvents.NAVIGATE_TO_HOME);
  }

  // 便捷方法：导航到个人资料页
  navigateToProfile() {
    this.emit(routerEvents.NAVIGATE_TO_PROFILE);
  }

  // 便捷方法：导航到编辑个人资料页
  navigateToEditProfile() {
    this.emit(routerEvents.NAVIGATE_TO_EDIT_PROFILE);
  }

  // 便捷方法：导航到账户安全页
  navigateToAccountSecurity() {
    this.emit(routerEvents.NAVIGATE_TO_ACCOUNT_SECURITY);
  }

  // 便捷方法：导航到用户安全页
  navigateToUserSecurity() {
    this.emit(routerEvents.NAVIGATE_TO_USER_SECURITY);
  }

  // 便捷方法：导航到添加教育经历页
  navigateToAddEducation() {
    this.emit(routerEvents.NAVIGATE_TO_ADD_EDUCATION);
  }

  // 便捷方法：导航到添加工作经历页
  navigateToAddOccupation() {
    this.emit(routerEvents.NAVIGATE_TO_ADD_OCCUPATION);
  }

  // 便捷方法：导航到编辑教育经历页
  navigateToEditEducation(id?: string) {
    const path = id ? `${ROUTES.EDIT_EDUCATION}?id=${id}` : ROUTES.EDIT_EDUCATION;
    this.emit(routerEvents.NAVIGATE_TO_EDIT_EDUCATION, { id, path });
  }

  // 便捷方法：导航到编辑工作经历页
  navigateToEditOccupation(id?: string) {
    const path = id ? `${ROUTES.EDIT_OCCUPATION}?id=${id}` : ROUTES.EDIT_OCCUPATION;
    this.emit(routerEvents.NAVIGATE_TO_EDIT_OCCUPATION, { id, path });
  }

  // 便捷方法：导航到编辑偏好设置页
  navigateToEditPreference() {
    this.emit(routerEvents.NAVIGATE_TO_EDIT_PREFERENCE);
  }

  // 便捷方法：返回上一页
  navigateBack() {
    this.emit(routerEvents.NAVIGATE_BACK);
  }

  // 便捷方法：前进到下一页
  navigateForward() {
    this.emit(routerEvents.NAVIGATE_FORWARD);
  }
}

export const routerEventEmitter = new RouterEventEmitter();
