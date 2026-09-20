import type { ReactNode } from "react";

import { Box, Flex, Heading, Text } from "@radix-ui/themes";

import styles from "./ledger.module.css";

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
    <section className={styles.section}>
      <Flex align="start" justify="between" gap="3" wrap="wrap" className={styles.sectionHead}>
        <Box minWidth="0">
          <Heading as="h2" size="3" className={styles.sectionTitle}>
            {title}
          </Heading>
          {description ? (
            <Text as="p" size="1" color="gray">
              {description}
            </Text>
          ) : null}
        </Box>
        {actions ? <Flex gap="2" wrap="wrap">{actions}</Flex> : null}
      </Flex>
      {children}
    </section>
  );
}

export function FieldRow({ label, value, children }: { label: string; value?: ReactNode; children?: ReactNode }) {
  return (
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
  );
}

export function FieldList({ children }: { children: ReactNode }) {
  return <div className={styles.list}>{children}</div>;
}
