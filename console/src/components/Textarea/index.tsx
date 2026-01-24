import type { TextareaHTMLAttributes, ReactNode } from "react";

import { forwardRef } from "react";

import styles from "./styles.module.css";

export interface TextareaProps extends Omit<TextareaHTMLAttributes<HTMLTextAreaElement>, "size"> {
  label?: string;
  error?: string;
  helperText?: string;
  leftIcon?: ReactNode;
  rightIcon?: ReactNode;
  size?: "small" | "medium" | "large";
  variant?: "default" | "filled";
  fullWidth?: boolean;
}

const Textarea = forwardRef<HTMLTextAreaElement, TextareaProps>(
  (
    {
      label,
      error,
      helperText,
      leftIcon,
      rightIcon,
      size = "medium",
      variant = "default",
      fullWidth = false,
      className = "",
      disabled,
      ...props
    },
    ref,
  ) => {
    const textareaClasses = [
      styles.textarea,
      styles[size],
      styles[variant],
      error && styles.error,
      disabled && styles.disabled,
      className,
    ]
      .filter(Boolean)
      .join(" ");

    const wrapperClasses = [styles.wrapper, fullWidth && styles.fullWidth].filter(Boolean).join(" ");

    return (
      <div className={wrapperClasses}>
        {label && (
          <label className={styles.label}>
            {label}
            {props.required && <span className={styles.required}>*</span>}
          </label>
        )}
        <div className={styles.inputContainer}>
          {leftIcon && <div className={styles.leftIcon}>{leftIcon}</div>}
          <textarea ref={ref} className={textareaClasses} disabled={disabled} {...props} />
          {rightIcon && <div className={styles.rightIcon}>{rightIcon}</div>}
        </div>
        {error && <span className={styles.errorText}>{error}</span>}
        {helperText && !error && <span className={styles.helperText}>{helperText}</span>}
      </div>
    );
  },
);

Textarea.displayName = "Textarea";

export default Textarea;
