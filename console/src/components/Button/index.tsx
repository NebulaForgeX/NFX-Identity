import type { ButtonHTMLAttributes, ReactNode } from "react";

import { forwardRef } from "react";

import styles from "./styles.module.css";

export interface ButtonProps extends Omit<ButtonHTMLAttributes<HTMLButtonElement>, "size"> {
  variant?: "primary" | "secondary" | "outline" | "ghost" | "danger";
  size?: "small" | "medium" | "large";
  fullWidth?: boolean;
  leftIcon?: ReactNode;
  rightIcon?: ReactNode;
  loading?: boolean;
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      variant = "primary",
      size = "medium",
      fullWidth = false,
      leftIcon,
      rightIcon,
      loading = false,
      disabled,
      children,
      className = "",
      ...props
    },
    ref,
  ) => {
    const buttonClasses = [
      styles.button,
      styles[variant],
      styles[size],
      fullWidth && styles.fullWidth,
      loading && styles.loading,
      className,
    ]
      .filter(Boolean)
      .join(" ");

    return (
      <button ref={ref} className={buttonClasses} disabled={disabled || loading} {...props}>
        {loading && <span className={styles.spinner} />}
        {leftIcon && !loading && <span className={styles.leftIcon}>{leftIcon}</span>}
        {children && <span className={styles.content}>{children}</span>}
        {rightIcon && !loading && <span className={styles.rightIcon}>{rightIcon}</span>}
      </button>
    );
  },
);

Button.displayName = "Button";

export default Button;
