// 路由常量定义 - 只定义实际使用的路由
export const ROUTES = {
  HOME: "/",
  LOGIN: "/login",
  DASHBOARD: "/dashboard",
  // Permission routes
  PERMISSION_LIST: "/permissions",
  PERMISSION_DETAIL: "/permissions/:id",
  PERMISSION_EDIT: "/permissions/:id/edit",
  // User Permission routes
  USER_PERMISSIONS: "/users/:userId/permissions",
  // Authorization Code routes
  AUTHORIZATION_CODE_LIST: "/authorization-codes",
  AUTHORIZATION_CODE_DETAIL: "/authorization-codes/:id",
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
