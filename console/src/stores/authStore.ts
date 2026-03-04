import type { Nullable } from "nfx-ui/types";
import type { Tokens } from "@/types";

import { makePersistStore } from "nfx-ui/stores";

interface AuthState {
  isAuthValid: boolean;
  accessToken: Nullable<string>;
  refreshToken: Nullable<string>;
  currentUserId: string;
}

interface AuthActions {
  setIsAuthValid: (isAuthValid: boolean) => void;
  setTokens: (tokens: Tokens) => void;
  setRefreshToken: (refreshToken: string) => void;
  setCurrentUserId: (userId: string) => void;
  getCurrentUserId: () => Nullable<string>;
  clearAuth: () => void;
}

const { store: AuthStore, useStore: useAuthStore } = makePersistStore<AuthState, AuthActions>({
  name: "auth-storage",
  initialState: {
    isAuthValid: false,
    accessToken: null,
    refreshToken: null,
    currentUserId: "00000000-0000-0000-0000-000000000000",
  },
  actions: (set, get) => ({
    setIsAuthValid: (isAuthValid) => set({ isAuthValid }),

    setTokens: (tokens) =>
      set({
        accessToken: tokens.accessToken,
        refreshToken: tokens.refreshToken ?? null,
      }),

    setRefreshToken: (refreshToken) => set({ refreshToken }),

    setCurrentUserId: (userId) => set({ currentUserId: userId }),

    getCurrentUserId: () => (get().isAuthValid ? get().currentUserId : null),

    clearAuth: () =>
      set({
        isAuthValid: false,
        accessToken: null,
        refreshToken: null,
        currentUserId: "00000000-0000-0000-0000-000000000000",
      }),
  }),
});

export { AuthStore, useAuthStore };
export default AuthStore;
