import type { ButtonProps } from "@/components/Button";
import type { ReactNode } from "react";

import { cloneElement, forwardRef, isValidElement } from "react";

import Button from "@/components/Button";

import styles from "./styles.module.css";

export interface IconButtonProps extends Omit<ButtonProps, "leftIcon" | "rightIcon"> {
  topIcon?: ReactNode;
  rightIcon?: ReactNode;
  bottomIcon?: ReactNode;
  leftIcon?: ReactNode;
  iconOnly?: boolean;
  iconSize?: number;
}

const IconButton = forwardRef<HTMLButtonElement, IconButtonProps>(
  (
    {
      topIcon,
      rightIcon,
      bottomIcon,
      leftIcon,
      iconOnly = false,
      iconSize,
      children,
      className = "",
      ...buttonProps
    },
    ref,
  ) => {
    // 为图标添加 size 属性
    const renderIcon = (icon: ReactNode) => {
      if (!icon) return null;
      if (iconSize && isValidElement(icon) && typeof icon.type !== "string") {
        const existingProps = icon.props && typeof icon.props === "object" ? icon.props : {};
        return cloneElement(icon, { ...existingProps, size: iconSize });
      }
      return icon;
    };

    const hasIcons = topIcon || rightIcon || bottomIcon || leftIcon;
    const hasContent = !iconOnly && children;

    // 如果只有图标且 iconOnly，使用简化的布局
    if (iconOnly && hasIcons && !hasContent) {
      return (
        <Button
          ref={ref}
          className={`${styles.iconOnly} ${className}`}
          {...buttonProps}
        >
          <div className={styles.iconContainer}>
            {topIcon && <span className={styles.topIcon}>{renderIcon(topIcon)}</span>}
            {leftIcon && <span className={styles.leftIcon}>{renderIcon(leftIcon)}</span>}
            {rightIcon && <span className={styles.rightIcon}>{renderIcon(rightIcon)}</span>}
            {bottomIcon && <span className={styles.bottomIcon}>{renderIcon(bottomIcon)}</span>}
          </div>
        </Button>
      );
    }

    // 完整布局：支持四个方向的图标和内容
    return (
      <Button
        ref={ref}
        className={`${styles.iconButton} ${className}`}
        {...buttonProps}
      >
        <div className={styles.layout}>
          {topIcon && <span className={styles.topIcon}>{renderIcon(topIcon)}</span>}
          <div className={styles.horizontal}>
            {leftIcon && <span className={styles.leftIcon}>{renderIcon(leftIcon)}</span>}
            {hasContent && <span className={styles.content}>{children}</span>}
            {rightIcon && <span className={styles.rightIcon}>{renderIcon(rightIcon)}</span>}
          </div>
          {bottomIcon && <span className={styles.bottomIcon}>{renderIcon(bottomIcon)}</span>}
        </div>
      </Button>
    );
  },
);

IconButton.displayName = "IconButton";

export default IconButton;
