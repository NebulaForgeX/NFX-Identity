import { TruckLoading } from "@/animations";

import styles from "./styles.module.css";

/**
 * Bootstrap Page - 系统初始化页面
 * 当系统未初始化时显示此页面
 */
export default function BootstrapPage() {
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
