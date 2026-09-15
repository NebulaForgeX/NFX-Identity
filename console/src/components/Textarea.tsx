import type { ComponentPropsWithoutRef } from "react";

import { forwardRef } from "react";
import { Flex, Text, TextArea } from "@radix-ui/themes";

export type TextareaProps = Omit<ComponentPropsWithoutRef<"textarea">, "color"> & {
  label?: string;
  error?: string;
  helperText?: string;
  fullWidth?: boolean;
};

const Textarea = forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ label, error, helperText, fullWidth = true, className, rows, value, defaultValue, onChange, onBlur, name, placeholder, disabled, required, id }, ref) => {
    return (
      <Flex direction="column" gap="1" width={fullWidth ? "100%" : undefined} className={className}>
        {label ? (
          <Text as="label" size="2" weight="medium">
            {label}
          </Text>
        ) : null}
        <TextArea
          ref={ref}
          rows={rows}
          value={typeof value === "string" ? value : undefined}
          defaultValue={typeof defaultValue === "string" ? defaultValue : undefined}
          onChange={onChange}
          onBlur={onBlur}
          name={name}
          placeholder={placeholder}
          disabled={disabled}
          required={required}
          id={id}
          color={error ? "red" : undefined}
        />
        {error ? (
          <Text size="1" color="red">
            {error}
          </Text>
        ) : helperText ? (
          <Text size="1" color="gray">
            {helperText}
          </Text>
        ) : null}
      </Flex>
    );
  },
);

Textarea.displayName = "Textarea";
export default Textarea;
