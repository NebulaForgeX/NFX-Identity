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
