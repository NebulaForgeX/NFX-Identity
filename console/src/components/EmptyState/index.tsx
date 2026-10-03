import type { ReactNode } from "react";

import { Box, Container, Flex, Heading, Section, Text } from "@radix-ui/themes";
import { AnimatedIcon, type AnimatedIconComponent } from "nfx-ui/icons";

import styles from "./s.module.css";

export type EmptyStateProps = {
  icon: AnimatedIconComponent;
  title: string;
  description?: string;
  action?: ReactNode;
  titleAs?: "h1" | "h2" | "h3";
};

export default function EmptyState({ icon, title, description, action, titleAs = "h3" }: EmptyStateProps) {
  return (
    <Section size="1" py="7" className={styles.frame}>
      <Container size="1" px="4">
        <Flex direction="column" align="center" justify="center" gap="3">
          <Box className={styles.mark}>
            <Flex align="center" justify="center" width="100%" height="100%">
              <AnimatedIcon icon={icon} size={22} />
            </Flex>
          </Box>
          <Heading as={titleAs} size="3" align="center">
            {title}
          </Heading>
          {description ? (
            <Text as="p" size="2" color="gray" align="center" className={styles.copy}>
              {description}
            </Text>
          ) : null}
          {action}
        </Flex>
      </Container>
    </Section>
  );
}
