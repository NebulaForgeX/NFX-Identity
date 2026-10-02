import type { ReactNode } from "react";

import { Container, Flex, Heading, Section, Text } from "@radix-ui/themes";
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
    <Container size="1" px="4">
      <Section size="1" py="9">
        <Flex direction="column" align="center" justify="center" gap="3">
          <Flex align="center" justify="center" className={styles.mark}>
            <AnimatedIcon icon={icon} size={24} />
          </Flex>
          <Heading as={titleAs} size="4" align="center">
            {title}
          </Heading>
          {description ? (
            <Text as="p" size="2" color="gray" align="center" className={styles.copy}>
              {description}
            </Text>
          ) : null}
          {action}
        </Flex>
      </Section>
    </Container>
  );
}
