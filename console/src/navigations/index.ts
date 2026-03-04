import type { RouteKey as RouteKeyGeneric, RoutePath as RoutePathGeneric } from "nfx-ui/navigations";

import { createRouter, defineRouter } from "nfx-ui/navigations";

/**
 * 路由常量与工具（基于 nfx-ui defineRouter / createRouter）
 * 只定义实际使用的路由
 */

const routeMap = defineRouter({
  HOME: "/",
  LOGIN: "/login",
  DASHBOARD: "/dashboard",
  PROFILE: "/profile",
  IMAGES: "/images",
  EDIT_PROFILE: "/profile/edit",
  ACCOUNT_SECURITY: "/profile/security",
  USER_SECURITY: "/user-security",
  PERMISSION_MANAGEMENT: "/permission/management",
  PERMISSION_ROLES: "/permission/roles",
  PERMISSION_PERMISSIONS: "/permission/permissions",
  PERMISSION_ACTIONS: "/permission/actions",
  ADD_EDUCATION: "/profile/add-education",
  ADD_OCCUPATION: "/profile/add-occupation",
  EDIT_EDUCATION: "/profile/edit-education",
  EDIT_OCCUPATION: "/profile/edit-occupation",
  EDIT_PREFERENCE: "/profile/edit-preference",
});

const { ROUTES, matchRoute, isActiveRoute, getRouteByKey } = createRouter(routeMap);
type RouteKey = RouteKeyGeneric<typeof routeMap>;
type RoutePath = RoutePathGeneric<typeof routeMap>;

export { ROUTES, matchRoute, isActiveRoute, getRouteByKey, type RouteKey, type RoutePath };
