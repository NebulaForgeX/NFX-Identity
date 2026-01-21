import type { ReactNode } from "react";

import Suspense from "@/components/Suspense";
import { TruckLoading } from "@/components";

import { useSystemInit } from "@/hooks/useSystem";

import styles from "./style.module.css";

interface BootstrapProviderProps {
  children: ReactNode;
}

/**
 * BootstrapProvider - 系统初始化 Provider
 * 检查系统是否已初始化，如果未初始化则显示引导页面
 */
function BootstrapContent({ children }: { children: ReactNode }) {
  const systemState = useSystemInit();

  // 如果系统未初始化，显示引导页面
  if (!systemState.data.initialized) {
    return (
      <div className={styles.container}>
        <div className={styles.content}>
          <TruckLoading size="large" />
          <h1 className={styles.title}>欢迎使用 NFX-Identity</h1>
          <p className={styles.description}>
            系统正在初始化，请稍候...
          </p>
          <p className={styles.hint}>
            首次使用需要创建系统管理员账户
          </p>
        </div>
      </div>
    );
  }

  // 系统已初始化，渲染子组件
  return <>{children}</>;
}

export function BootstrapProvider({ children }: BootstrapProviderProps) {
  return (
    <Suspense
      loadingType="ecg"
      loadingText="检查系统状态..."
      loadingSize="medium"
    >
      <BootstrapContent>{children}</BootstrapContent>
    </Suspense>
  );
}
