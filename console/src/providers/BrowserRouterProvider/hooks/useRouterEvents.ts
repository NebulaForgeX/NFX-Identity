import { useCallback, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import { authEventEmitter, authEvents } from "@/events/auth";
import { routerEventEmitter, routerEvents } from "@/events/router";
import { ROUTES } from "@/types/navigation";

/**
 * useRouterEvents - 处理路由导航事件
 * 监听路由事件并执行实际的导航操作
 */
export function useRouterEvents() {
  const navigate = useNavigate();
  // 处理通用导航事件
  const handleNavigate = useCallback(
    (payload: { to: string; replace?: boolean; state?: unknown }) => {
      if (payload.replace) {
        navigate(payload.to, { replace: true, state: payload.state });
      } else {
        navigate(payload.to, { state: payload.state });
      }
    },
    [navigate],
  );

  // 处理替换导航事件
  const handleNavigateReplace = useCallback(
    (payload: { to: string; state?: unknown }) => {
      navigate(payload.to, { replace: true, state: payload.state });
    },
    [navigate],
  );

  // 处理返回事件
  const handleNavigateBack = useCallback(() => {
    navigate(-1);
  }, [navigate]);

  // 处理前进事件
  const handleNavigateForward = useCallback(() => {
    navigate(1);
  }, [navigate]);

  // 处理跳转到登录页
  const handleNavigateToLogin = useCallback(() => {
    navigate(ROUTES.LOGIN, { replace: true });
  }, [navigate]);

  // 处理跳转到仪表板
  const handleNavigateToDashboard = useCallback(() => {
    navigate(ROUTES.DASHBOARD, { replace: true });
  }, [navigate]);

  // 处理跳转到首页
  const handleNavigateToHome = useCallback(() => {
    navigate(ROUTES.HOME, { replace: true });
  }, [navigate]);

  // 处理登录成功事件
  const handleLoginSuccess = useCallback(() => {
    navigate(ROUTES.DASHBOARD, { replace: true });
  }, [navigate]);

  // 注册所有路由事件监听器
  useEffect(() => {
    routerEventEmitter.on(routerEvents.NAVIGATE, handleNavigate);
    routerEventEmitter.on(routerEvents.NAVIGATE_REPLACE, handleNavigateReplace);
    routerEventEmitter.on(routerEvents.NAVIGATE_BACK, handleNavigateBack);
    routerEventEmitter.on(routerEvents.NAVIGATE_FORWARD, handleNavigateForward);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_LOGIN, handleNavigateToLogin);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_DASHBOARD, handleNavigateToDashboard);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_HOME, handleNavigateToHome);
    
    // 监听登录成功事件
    authEventEmitter.on(authEvents.LOGIN_SUCCESS, handleLoginSuccess);

    return () => {
      routerEventEmitter.off(routerEvents.NAVIGATE, handleNavigate);
      routerEventEmitter.off(routerEvents.NAVIGATE_REPLACE, handleNavigateReplace);
      routerEventEmitter.off(routerEvents.NAVIGATE_BACK, handleNavigateBack);
      routerEventEmitter.off(routerEvents.NAVIGATE_FORWARD, handleNavigateForward);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_LOGIN, handleNavigateToLogin);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_DASHBOARD, handleNavigateToDashboard);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_HOME, handleNavigateToHome);
      
      // 移除登录成功事件监听
      authEventEmitter.off(authEvents.LOGIN_SUCCESS, handleLoginSuccess);
    };
  }, [
    handleNavigate,
    handleNavigateReplace,
    handleNavigateBack,
    handleNavigateForward,
    handleNavigateToLogin,
    handleNavigateToDashboard,
    handleNavigateToHome,
    handleLoginSuccess,
  ]);
}
