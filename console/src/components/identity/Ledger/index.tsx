import type { ReactNode } from "react";

import { Box, Flex, Grid, Heading, Section, Text } from "@radix-ui/themes";

import { Surface } from "@/components";

import styles from "./s.module.css";

export function LedgerSection({
  title,
  description,
  actions,
  sticky,
  children,
}: {
  title: string;
  description?: string;
  actions?: ReactNode;
  sticky?: boolean;
  children: ReactNode;
}) {
  return (
    <Surface sticky={sticky}>
      <Flex direction="column" gap="5" width="100%">
        <Flex align="start" justify="between" gap="4" wrap="wrap">
          <Flex direction="column" gap="1" minWidth="0">
            <Heading as="h2" size="4" weight="bold">
              {title}
            </Heading>
            {description ? (
              <Text as="p" size="2" color="gray">
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
    </Surface>
  );
}

export function FieldRow({ label, value, children }: { label: string; value?: ReactNode; children?: ReactNode }) {
  return (
    <Section size="1" py="3" className={styles.hairline}>
      <Grid columns={{ initial: "1", sm: "minmax(10rem, 14rem) minmax(0, 1fr)" }} gap={{ initial: "1", sm: "5" }} align="center">
        <Text as="p" size="1" color="gray" weight="medium" className={styles.label}>
          {label}
        </Text>
        <Box className={styles.value}>
          {children ?? (
            <Text as="p" size="2">
              {value}
            </Text>
          )}
        </Box>
      </Grid>
    </Section>
  );
}

export function FieldList({ children }: { children: ReactNode }) {
  return <Flex direction="column">{children}</Flex>;
}
