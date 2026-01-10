import { memo } from "react";

import { ExternalLink, FileText, History, Home, Info, Shield } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

interface ResourceLink {
  title: string;
  description: string;
  icon: React.ReactNode;
  url: string;
  color: string;
}

const ResourceLinks = memo(() => {
  // TODO: 更新为实际的项目资源链接
  const resourceLinks: ResourceLink[] = [
    {
      title: "项目主页",
      description: "访问项目主页",
      icon: <Home size={20} />,
      url: "#",
      color: "var(--color-fg-highlight)",
    },
    {
      title: "更新日志",
      description: "最新更新和功能",
      icon: <History size={20} />,
      url: "#",
      color: "var(--color-info)",
    },
    {
      title: "关于我们",
      description: "关于项目的信息",
      icon: <Info size={20} />,
      url: "#",
      color: "var(--color-success)",
    },
    {
      title: "服务条款",
      description: "用户条款和条件",
      icon: <FileText size={20} />,
      url: "#",
      color: "var(--color-fg-muted)",
    },
    {
      title: "隐私政策",
      description: "我们如何保护您的数据",
      icon: <Shield size={20} />,
      url: "#",
      color: "var(--color-danger)",
    },
  ];

  const handleLinkClick = (url: string) => {
    window.open(url, "_blank", "noopener,noreferrer");
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h2 className={styles.title}>资源与政策</h2>
        <p className={styles.subtitle}>了解更多关于平台和政策</p>
      </div>

      <div className={styles.grid}>
        {resourceLinks.map((link, index) => (
          <button
            key={index}
            className={styles.linkCard}
            onClick={() => handleLinkClick(link.url)}
            style={{ "--card-color": link.color } as React.CSSProperties}
          >
            <div className={styles.iconWrapper}>{link.icon}</div>
            <div className={styles.content}>
              <h3 className={styles.cardTitle}>{link.title}</h3>
              <p className={styles.cardDescription}>{link.description}</p>
            </div>
            <ExternalLink size={16} className={styles.externalIcon} />
          </button>
        ))}
      </div>
    </div>
  );
});

ResourceLinks.displayName = "ResourceLinks";

export default ResourceLinks;
