import { memo } from "react";

import styles from "./styles.module.css";

interface BounceLoadingProps {
  size?: "small" | "medium" | "large";
  shape?: "square" | "circle";
  className?: string;
}

const BounceLoading = memo(({ size = "medium", shape = "square", className }: BounceLoadingProps) => {
  const sizeMap = {
    small: "32px",
    medium: "48px",
    large: "64px",
  };

  const loaderSize = sizeMap[size];

  return (
    <div
      className={`${styles.loader} ${styles[shape]} ${className || ""}`}
      style={{ width: loaderSize, height: loaderSize }}
    />
  );
});

BounceLoading.displayName = "BounceLoading";

export default BounceLoading;
