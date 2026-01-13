export const systemEvents = {
  // SystemState 相关
  INVALIDATE_SYSTEM_STATE: "SYSTEM:INVALIDATE_SYSTEM_STATE",
  INVALIDATE_SYSTEM_STATES: "SYSTEM:INVALIDATE_SYSTEM_STATES",
} as const;

type SystemEvent = (typeof systemEvents)[keyof typeof systemEvents];

class SystemEventEmitter {
  private listeners: Record<SystemEvent, Set<Function>> = {
    [systemEvents.INVALIDATE_SYSTEM_STATE]: new Set<Function>(),
    [systemEvents.INVALIDATE_SYSTEM_STATES]: new Set<Function>(),
  };

  on(event: SystemEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: SystemEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: SystemEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const systemEventEmitter = new SystemEventEmitter();
