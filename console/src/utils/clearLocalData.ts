import { AuthStore } from "nfx-ui/stores";
import { authEventEmitter } from "nfx-ui/events";
import { routerEventEmitter, routerEvents } from "@/events/router";

export function clearLocalData(options?: { navigateToLogin?: boolean }) {
  AuthStore.getState().clearAuth();
  authEventEmitter.logout();
  if (options?.navigateToLogin !== false) {
    routerEventEmitter.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }
}
