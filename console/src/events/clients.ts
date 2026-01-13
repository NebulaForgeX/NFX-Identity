export const clientsEvents = {
  // App 相关
  INVALIDATE_APP: "CLIENTS:INVALIDATE_APP",
  INVALIDATE_APPS: "CLIENTS:INVALIDATE_APPS",

  // APIKey 相关
  INVALIDATE_API_KEY: "CLIENTS:INVALIDATE_API_KEY",
  INVALIDATE_API_KEYS: "CLIENTS:INVALIDATE_API_KEYS",

  // ClientCredential 相关
  INVALIDATE_CLIENT_CREDENTIAL: "CLIENTS:INVALIDATE_CLIENT_CREDENTIAL",
  INVALIDATE_CLIENT_CREDENTIALS: "CLIENTS:INVALIDATE_CLIENT_CREDENTIALS",

  // ClientScope 相关
  INVALIDATE_CLIENT_SCOPE: "CLIENTS:INVALIDATE_CLIENT_SCOPE",
  INVALIDATE_CLIENT_SCOPES: "CLIENTS:INVALIDATE_CLIENT_SCOPES",

  // IPAllowlist 相关
  INVALIDATE_IP_ALLOWLIST: "CLIENTS:INVALIDATE_IP_ALLOWLIST",
  INVALIDATE_IP_ALLOWLISTS: "CLIENTS:INVALIDATE_IP_ALLOWLISTS",

  // RateLimit 相关
  INVALIDATE_RATE_LIMIT: "CLIENTS:INVALIDATE_RATE_LIMIT",
  INVALIDATE_RATE_LIMITS: "CLIENTS:INVALIDATE_RATE_LIMITS",
} as const;

type ClientsEvent = (typeof clientsEvents)[keyof typeof clientsEvents];

class ClientsEventEmitter {
  private listeners: Record<ClientsEvent, Set<Function>> = {
    [clientsEvents.INVALIDATE_APP]: new Set<Function>(),
    [clientsEvents.INVALIDATE_APPS]: new Set<Function>(),
    [clientsEvents.INVALIDATE_API_KEY]: new Set<Function>(),
    [clientsEvents.INVALIDATE_API_KEYS]: new Set<Function>(),
    [clientsEvents.INVALIDATE_CLIENT_CREDENTIAL]: new Set<Function>(),
    [clientsEvents.INVALIDATE_CLIENT_CREDENTIALS]: new Set<Function>(),
    [clientsEvents.INVALIDATE_CLIENT_SCOPE]: new Set<Function>(),
    [clientsEvents.INVALIDATE_CLIENT_SCOPES]: new Set<Function>(),
    [clientsEvents.INVALIDATE_IP_ALLOWLIST]: new Set<Function>(),
    [clientsEvents.INVALIDATE_IP_ALLOWLISTS]: new Set<Function>(),
    [clientsEvents.INVALIDATE_RATE_LIMIT]: new Set<Function>(),
    [clientsEvents.INVALIDATE_RATE_LIMITS]: new Set<Function>(),
  };

  on(event: ClientsEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: ClientsEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: ClientsEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const clientsEventEmitter = new ClientsEventEmitter();
