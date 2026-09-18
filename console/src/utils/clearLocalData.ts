import { AuthStore } from "nfx-ui/stores";
import { authEventEmitter } from "nfx-ui/events";
import { ApiAuthRepository } from "nfx-ui/apis";
import { routerEventEmitter, routerEvents } from "@/events/router";

const auth = new ApiAuthRepository();

export async function clearLocalData(options?: { navigateToLogin?: boolean }) {
  const refreshToken = AuthStore.getState().refreshToken;
  if (refreshToken) {
    try {
      await auth.Logout({ refreshToken });
    } catch {
      // Local session must still be cleared if revoke fails.
    }
  }
  AuthStore.getState().clearAuth();
  authEventEmitter.logout();
  if (options?.navigateToLogin !== false) {
    routerEventEmitter.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }
}
