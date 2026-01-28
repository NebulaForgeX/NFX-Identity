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

  // 处理跳转到个人资料页
  const handleNavigateToProfile = useCallback(() => {
    navigate(ROUTES.PROFILE);
  }, [navigate]);

  // 处理跳转到编辑个人资料页
  const handleNavigateToEditProfile = useCallback(() => {
    navigate(ROUTES.EDIT_PROFILE);
  }, [navigate]);

  // 处理跳转到账户安全页
  const handleNavigateToAccountSecurity = useCallback(() => {
    navigate(ROUTES.ACCOUNT_SECURITY);
  }, [navigate]);

  // 处理跳转到用户安全页
  const handleNavigateToUserSecurity = useCallback(() => {
    navigate(ROUTES.USER_SECURITY);
  }, [navigate]);

  // 处理跳转到添加教育经历页
  const handleNavigateToAddEducation = useCallback(() => {
    navigate(ROUTES.ADD_EDUCATION);
  }, [navigate]);

  // 处理跳转到添加工作经历页
  const handleNavigateToAddOccupation = useCallback(() => {
    navigate(ROUTES.ADD_OCCUPATION);
  }, [navigate]);

  // 处理跳转到编辑教育经历页
  const handleNavigateToEditEducation = useCallback(
    (payload?: { id?: string; path?: string }) => {
      if (payload?.path) {
        navigate(payload.path);
      } else if (payload?.id) {
        navigate(`${ROUTES.EDIT_EDUCATION}?id=${payload.id}`);
      } else {
        navigate(ROUTES.EDIT_EDUCATION);
      }
    },
    [navigate],
  );

  // 处理跳转到编辑工作经历页
  const handleNavigateToEditOccupation = useCallback(
    (payload?: { id?: string; path?: string }) => {
      if (payload?.path) {
        navigate(payload.path);
      } else if (payload?.id) {
        navigate(`${ROUTES.EDIT_OCCUPATION}?id=${payload.id}`);
      } else {
        navigate(ROUTES.EDIT_OCCUPATION);
      }
    },
    [navigate],
  );

  // 处理跳转到编辑偏好设置页
  const handleNavigateToEditPreference = useCallback(() => {
    navigate(ROUTES.EDIT_PREFERENCE);
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
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_PROFILE, handleNavigateToProfile);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_EDIT_PROFILE, handleNavigateToEditProfile);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_ACCOUNT_SECURITY, handleNavigateToAccountSecurity);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_USER_SECURITY, handleNavigateToUserSecurity);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_ADD_EDUCATION, handleNavigateToAddEducation);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_ADD_OCCUPATION, handleNavigateToAddOccupation);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_EDIT_EDUCATION, handleNavigateToEditEducation);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_EDIT_OCCUPATION, handleNavigateToEditOccupation);
    routerEventEmitter.on(routerEvents.NAVIGATE_TO_EDIT_PREFERENCE, handleNavigateToEditPreference);

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
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_PROFILE, handleNavigateToProfile);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_EDIT_PROFILE, handleNavigateToEditProfile);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_ACCOUNT_SECURITY, handleNavigateToAccountSecurity);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_USER_SECURITY, handleNavigateToUserSecurity);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_ADD_EDUCATION, handleNavigateToAddEducation);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_ADD_OCCUPATION, handleNavigateToAddOccupation);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_EDIT_EDUCATION, handleNavigateToEditEducation);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_EDIT_OCCUPATION, handleNavigateToEditOccupation);
      routerEventEmitter.off(routerEvents.NAVIGATE_TO_EDIT_PREFERENCE, handleNavigateToEditPreference);

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
    handleNavigateToProfile,
    handleNavigateToEditProfile,
    handleNavigateToAccountSecurity,
    handleNavigateToUserSecurity,
    handleNavigateToAddEducation,
    handleNavigateToAddOccupation,
    handleNavigateToEditEducation,
    handleNavigateToEditOccupation,
    handleNavigateToEditPreference,
    handleLoginSuccess,
  ]);
}
