import { ProfileKindEnum } from "nfx-ui/enums";
import { createRouter, defineRouter } from "@/utils";

const routeMap = defineRouter({
  HOME: "/",
  LOGIN: "/auth/login",
  SIGNUP: "/auth/signup",
  SELECT_PROFILE: "/auth/select-profile",

  FORGER: "/forger",
  FORGER_DESK: "/forger/desk",
  FORGER_PROFILE: "/forger/profile",
  FORGER_PROFILE_OVERVIEW: "/forger/profile/overview",
  FORGER_PROFILE_EDIT: "/forger/profile/edit",
  FORGER_PROFILE_IDENTITY: "/forger/profile/identity",
  FORGER_PROFILE_SECURITY: "/forger/profile/security",
  FORGER_ASSETS: "/forger/assets",
  FORGER_SETTINGS: "/forger/settings",

  AUTHORITY: "/authority",
  AUTHORITY_DESK: "/authority/desk",
  AUTHORITY_PROFILE: "/authority/profile",
  AUTHORITY_PROFILE_OVERVIEW: "/authority/profile/overview",
  AUTHORITY_PROFILE_EDIT: "/authority/profile/edit",
  AUTHORITY_PROFILE_IDENTITY: "/authority/profile/identity",
  AUTHORITY_PROFILE_SECURITY: "/authority/profile/security",
  AUTHORITY_ASSETS: "/authority/assets",
  AUTHORITY_SETTINGS: "/authority/settings",
  AUTHORITY_DIRECTORY: "/authority/directory",
});

const { ROUTES, matchRoute, isActiveRoute, buildPath } = createRouter(routeMap);

export type RouteKey = keyof typeof ROUTES;
export { ROUTES, matchRoute, isActiveRoute, buildPath };

export type ScopePaths = {
  root: string;
  desk: string;
  profile: string;
  overview: string;
  edit: string;
  identity: string;
  security: string;
  assets: string;
  settings: string;
  directory?: string;
};

export function scopePaths(kind: ProfileKindEnum): ScopePaths {
  if (kind === ProfileKindEnum.AUTHORITY) {
    return {
      root: ROUTES.AUTHORITY,
      desk: ROUTES.AUTHORITY_DESK,
      profile: ROUTES.AUTHORITY_PROFILE,
      overview: ROUTES.AUTHORITY_PROFILE_OVERVIEW,
      edit: ROUTES.AUTHORITY_PROFILE_EDIT,
      identity: ROUTES.AUTHORITY_PROFILE_IDENTITY,
      security: ROUTES.AUTHORITY_PROFILE_SECURITY,
      assets: ROUTES.AUTHORITY_ASSETS,
      settings: ROUTES.AUTHORITY_SETTINGS,
      directory: ROUTES.AUTHORITY_DIRECTORY,
    };
  }
  return {
    root: ROUTES.FORGER,
    desk: ROUTES.FORGER_DESK,
    profile: ROUTES.FORGER_PROFILE,
    overview: ROUTES.FORGER_PROFILE_OVERVIEW,
    edit: ROUTES.FORGER_PROFILE_EDIT,
    identity: ROUTES.FORGER_PROFILE_IDENTITY,
    security: ROUTES.FORGER_PROFILE_SECURITY,
    assets: ROUTES.FORGER_ASSETS,
    settings: ROUTES.FORGER_SETTINGS,
  };
}

export function profileHome(kind: Nullable<ProfileKindEnum>): string {
  if (kind === ProfileKindEnum.AUTHORITY) return ROUTES.AUTHORITY_DESK;
  return ROUTES.FORGER_DESK;
}
