import type { InputHTMLAttributes, ReactNode } from "react";

import { forwardRef } from "react";

import styles from "./styles.module.css";

export interface InputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "size"> {
  label?: string;
  error?: string;
  helperText?: string;
  leftIcon?: ReactNode;
  rightIcon?: ReactNode;
  /** 为 true 时 rightIcon 可点击（如密码切换、清除按钮） */
  rightIconInteractive?: boolean;
  size?: "small" | "medium" | "large";
  variant?: "default" | "filled";
  fullWidth?: boolean;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  (
    {
      label,
      error,
      helperText,
      leftIcon,
      rightIcon,
      rightIconInteractive = false,
      size = "medium",
      variant = "default",
      fullWidth = false,
      className = "",
      disabled,
      ...props
    },
    ref,
  ) => {
    const inputClasses = [
      styles.input,
      styles[size],
      styles[variant],
      error && styles.error,
      disabled && styles.disabled,
      className,
    ]
      .filter(Boolean)
      .join(" ");

    const wrapperClasses = [styles.wrapper, fullWidth && styles.fullWidth].filter(Boolean).join(" ");

    const containerClasses = [
      styles.inputContainer,
      leftIcon && styles.withLeftIcon,
      rightIcon && styles.withRightIcon,
      size && styles[`container${size.charAt(0).toUpperCase() + size.slice(1)}`],
    ]
      .filter(Boolean)
      .join(" ");

    return (
      <div className={wrapperClasses}>
        {label && (
          <label className={styles.label}>
            {label}
            {props.required && <span className={styles.required}>*</span>}
          </label>
        )}
        <div className={containerClasses}>
          {leftIcon && <div className={styles.leftIcon}>{leftIcon}</div>}
          <input ref={ref} className={inputClasses} disabled={disabled} {...props} />
          {rightIcon && (
            <div className={`${styles.rightIcon} ${rightIconInteractive ? styles.rightIconInteractive : ""}`}>
              {rightIcon}
            </div>
          )}
        </div>
        {error && <p className={styles.errorMessage}>{error}</p>}
        {helperText && !error && <p className={styles.helperText}>{helperText}</p>}
      </div>
    );
  },
);

Input.displayName = "Input";

export default Input;
