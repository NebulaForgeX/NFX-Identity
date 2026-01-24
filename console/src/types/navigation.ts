// 路由常量定义 - 只定义实际使用的路由
export const ROUTES = {
  HOME: "/",
  LOGIN: "/login",
  DASHBOARD: "/dashboard",
  PROFILE: "/profile",
  EDIT_PROFILE: "/profile/edit",
  ACCOUNT_SECURITY: "/profile/security",
  ADD_EDUCATION: "/profile/add-education",
  ADD_OCCUPATION: "/profile/add-occupation",
  EDIT_EDUCATION: "/profile/edit-education",
  EDIT_OCCUPATION: "/profile/edit-occupation",
  EDIT_PREFERENCE: "/profile/edit-preference",
} as const;

// 路由类型
export type RouteKey = keyof typeof ROUTES;
export type RoutePath = (typeof ROUTES)[RouteKey];

// 工具函数
export const isActiveRoute = (currentPath: string, targetPath: RoutePath): boolean => {
  return currentPath === targetPath;
};

export const getRouteByKey = (key: RouteKey): RoutePath => {
  return ROUTES[key];
};
