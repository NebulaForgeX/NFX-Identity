import type { ReactNode } from "react";

import { BrowserRouter } from "react-router-dom";

import { useRouterEvents } from "./hooks/useRouterEvents";

interface BrowserRouterProviderProps {
  children: ReactNode;
}

/**
 * RouterEventsHandler - 在 BrowserRouter 内部处理路由导航事件
 * 必须在 BrowserRouter 的上下文中才能使用 useNavigate
 */
function RouterEventsHandler({ children }: { children: ReactNode }) {
  // 处理路由导航事件（这里可以使用 useNavigate，因为已经在 BrowserRouter 内部）
  useRouterEvents();
  return <>{children}</>;
}

/**
 * BrowserRouterProvider - 统一处理路由导航事件
 * 包含 BrowserRouter 和路由事件监听
 * 所有路由跳转都通过事件系统，而不是直接使用 navigate
 */
export function BrowserRouterProvider({ children }: BrowserRouterProviderProps) {
  return (
    <BrowserRouter>
      <RouterEventsHandler>{children}</RouterEventsHandler>
    </BrowserRouter>
  );
}
