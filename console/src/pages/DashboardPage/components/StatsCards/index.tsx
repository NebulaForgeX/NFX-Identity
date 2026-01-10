import { memo } from "react";

// import { FolderTree, Layers } from "@/assets/icons/lucide"; // TODO: 根据实际需要启用

import styles from "./styles.module.css";

interface StatCard {
  title: string;
  icon: React.ReactNode;
  color: string;
  route?: string;
}

const StatsCardsContent = memo(() => {
  // TODO: 后续接入真实 API 后替换为动态数据
  // 暂时使用硬编码的示例数据，保持 UI 风格
  const stats: (StatCard & { value: number })[] = [
    // TODO: 根据实际需求添加统计卡片
    // {
    //   title: "示例统计",
    //   value: 0,
    //   icon: <FolderTree size={32} />,
    //   color: "var(--color-fg-highlight)",
    // },
  ];

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h2 className={styles.title}>数据概览</h2>
        <p className={styles.subtitle}>系统核心数据统计</p>
      </div>

      <div className={styles.grid}>
        {stats.map((stat, index) => (
          <div key={index} className={styles.statCard} style={{ "--stat-color": stat.color } as React.CSSProperties}>
            <div className={styles.iconWrapper}>{stat.icon}</div>
            <div className={styles.content}>
              <div className={styles.value}>{stat.value.toLocaleString()}</div>
              <div className={styles.label}>{stat.title}</div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
});

StatsCardsContent.displayName = "StatsCardsContent";

export default StatsCardsContent;
