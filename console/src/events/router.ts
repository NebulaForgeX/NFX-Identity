import type { EventNamesOf } from "nfx-ui/events";
import { defineEvents, EventEmitter } from "nfx-ui/events";
import { singleton } from "nfx-ui/utils";
import { ROUTES } from "@/navigations";

export const routerEvents = defineEvents({
  NAVIGATE: "ROUTER:NAVIGATE",
  NAVIGATE_REPLACE: "ROUTER:NAVIGATE_REPLACE",
  NAVIGATE_BACK: "ROUTER:NAVIGATE_BACK",
  NAVIGATE_FORWARD: "ROUTER:NAVIGATE_FORWARD",
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
});

type RouterEvent = EventNamesOf<typeof routerEvents>;

interface NavigatePayload {
  to: string;
  replace?: boolean;
  state?: unknown;
}

class RouterEventEmitter extends EventEmitter<RouterEvent> {
  constructor() {
    super(routerEvents);
  }

  navigate(payload: NavigatePayload) {
    this.emit(routerEvents.NAVIGATE, payload);
  }

  navigateReplace(to: string, state?: unknown) {
    this.emit(routerEvents.NAVIGATE_REPLACE, { to, state });
  }

  navigateBack() {
    this.emit(routerEvents.NAVIGATE_BACK);
  }

  navigateForward() {
    this.emit(routerEvents.NAVIGATE_FORWARD);
  }

  navigateToLogin() {
    this.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }

  navigateToDashboard() {
    this.emit(routerEvents.NAVIGATE_TO_DASHBOARD);
  }

  navigateToHome() {
    this.emit(routerEvents.NAVIGATE_TO_HOME);
  }

  navigateToProfile() {
    this.emit(routerEvents.NAVIGATE_TO_PROFILE);
  }

  navigateToEditProfile() {
    this.emit(routerEvents.NAVIGATE_TO_EDIT_PROFILE);
  }

  navigateToAccountSecurity() {
    this.emit(routerEvents.NAVIGATE_TO_ACCOUNT_SECURITY);
  }

  navigateToUserSecurity() {
    this.emit(routerEvents.NAVIGATE_TO_USER_SECURITY);
  }

  navigateToAddEducation() {
    this.emit(routerEvents.NAVIGATE_TO_ADD_EDUCATION);
  }

  navigateToAddOccupation() {
    this.emit(routerEvents.NAVIGATE_TO_ADD_OCCUPATION);
  }

  navigateToEditEducation(id?: string) {
    const path = id ? `${ROUTES.EDIT_EDUCATION}?id=${id}` : ROUTES.EDIT_EDUCATION;
    this.emit(routerEvents.NAVIGATE_TO_EDIT_EDUCATION, { id, path });
  }

  navigateToEditOccupation(id?: string) {
    const path = id ? `${ROUTES.EDIT_OCCUPATION}?id=${id}` : ROUTES.EDIT_OCCUPATION;
    this.emit(routerEvents.NAVIGATE_TO_EDIT_OCCUPATION, { id, path });
  }

  navigateToEditPreference() {
    this.emit(routerEvents.NAVIGATE_TO_EDIT_PREFERENCE);
  }
}

export const routerEventEmitter = new (singleton(RouterEventEmitter))();
