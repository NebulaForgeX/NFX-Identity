import type { ReactNode } from "react";

import { Box, Flex, Heading, Text } from "@radix-ui/themes";
import { AnimatedIcon, type AnimatedIconComponent } from "nfx-ui/icons";

import styles from "./s.module.css";

export type PageHeaderProps = {
  icon: AnimatedIconComponent;
  title: string;
  description?: string;
  actions?: ReactNode;
};

export default function PageHeader({ icon, title, description, actions }: PageHeaderProps) {
  return (
    <Flex direction={{ initial: "column", md: "row" }} align={{ initial: "start", md: "end" }} justify="between" gap="4" width="100%" data-reveal="">
      <Flex align="center" gap="4" minWidth="0">
        <Box className={styles.stamp}>
          <Flex align="center" justify="center" width="100%" height="100%">
            <AnimatedIcon icon={icon} size={22} />
          </Flex>
        </Box>
        <Flex direction="column" gap="1" minWidth="0">
          <Heading as="h1" size="7" weight="bold" className={styles.title}>
            {title}
          </Heading>
          {description ? (
            <Text as="p" size="2" color="gray" className={styles.lede}>
              {description}
            </Text>
          ) : null}
        </Flex>
      </Flex>
      {actions ? (
        <Flex gap="2" wrap="wrap" align="center" flexShrink="0">
          {actions}
        </Flex>
      ) : null}
    </Flex>
  );
}
