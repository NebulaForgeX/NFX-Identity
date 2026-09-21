import type { ReactNode } from "react";

import { Box, Flex, Heading, Text } from "@radix-ui/themes";

import styles from "./s.module.css";

export function LedgerSection({
  title,
  description,
  actions,
  children,
}: {
  title: string;
  description?: string;
  actions?: ReactNode;
  children: ReactNode;
}) {
  return (
    <Flex direction="column" gap="4" width="100%">
      <Flex align="start" justify="between" gap="4" wrap="wrap">
        <Flex direction="column" gap="1" minWidth="0">
          <Heading as="h2" size="3" className={styles.sectionTitle}>
            {title}
          </Heading>
          {description ? (
            <Text as="p" size="1" color="gray">
              {description}
            </Text>
          ) : null}
        </Flex>
        {actions ? (
          <Flex gap="2" wrap="wrap">
            {actions}
          </Flex>
        ) : null}
      </Flex>
      {children}
    </Flex>
  );
}

export function FieldRow({ label, value, children }: { label: string; value?: ReactNode; children?: ReactNode }) {
  return (
    <Box className={styles.hairline}>
      <Box py="4">
        <div className={styles.row}>
          <Text as="p" size="1" color="gray" className={styles.label}>
            {label}
          </Text>
          <div className={styles.value}>
            {children ?? (
              <Text as="p" size="2">
                {value}
              </Text>
            )}
          </div>
        </div>
      </Box>
    </Box>
  );
}

export function FieldList({ children }: { children: ReactNode }) {
  return <Flex direction="column">{children}</Flex>;
}
