import type { ReactNode } from "react";

import { useRef } from "react";
import { Box, Container, Flex, Section } from "@radix-ui/themes";

import { usePanelReveal } from "@/hooks";

const PAGE_FRAME_MAX_WIDTH = "1440px";

/** Panel page column. Children marked `data-reveal` stagger in on mount. */
function PageFrame({ children }: { children: ReactNode }) {
  const rootRef = useRef<HTMLDivElement>(null);
  usePanelReveal(rootRef);

  return (
    <Box ref={rootRef} width="100%">
      <Container size="4" width="100%" maxWidth={PAGE_FRAME_MAX_WIDTH} px={{ initial: "4", md: "6" }}>
        <Section size="1" py="6">
          <Flex direction="column" gap="5" width="100%" minWidth="0">
            {children}
          </Flex>
        </Section>
      </Container>
    </Box>
  );
}

export default PageFrame;
