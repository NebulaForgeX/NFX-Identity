import type { ReactNode } from "react";

import { Box, Container, Flex, Section } from "@radix-ui/themes";

import { safeStringable } from "@/utils";

import styles from "./s.module.css";

/** Fill the column beside the sidebar. */
const PAGE_FRAME_DEFAULT_MAX_WIDTH = "100%";

type PageFrameProps = {
  children: ReactNode;
  className?: string;
  maxWidth?: number | string;
  fullHeight?: boolean;
};

function PageFrame({ children, className, maxWidth = PAGE_FRAME_DEFAULT_MAX_WIDTH, fullHeight }: PageFrameProps) {
  const resolvedMaxWidth = typeof maxWidth === "number" ? `${maxWidth}px` : maxWidth;
  const callerClass = safeStringable(className);
  const shellClass = [fullHeight ? styles.fullHeightFrame : "", callerClass].filter(Boolean).join(" ");

  const frame = (
    <Container size="4" width="100%" maxWidth={resolvedMaxWidth} px="6">
      {fullHeight ? (
        <Flex direction="column" className={styles.fullHeightBody} width="100%" height="100%" minHeight="0">
          {children}
        </Flex>
      ) : (
        <Section size="1" py="6">
          <Flex direction="column" gap="6" width="100%" minWidth="0">
            {children}
          </Flex>
        </Section>
      )}
    </Container>
  );

  if (!shellClass) return frame;

  if (fullHeight) return <Flex direction="column" width="100%" className={shellClass}>{frame}</Flex>;

  return <Box className={shellClass}>{frame}</Box>;
}

export default PageFrame;
