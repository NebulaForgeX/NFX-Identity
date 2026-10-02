import { Box, Container, Flex, Section } from "@radix-ui/themes";

import { PreferencesPopover } from "@/components";

import styles from "./s.module.css";

export default function UserTopBar() {
  return (
    <Box position="sticky" top="0" className={styles.bar}>
      <Container size="4" px="6" width="100%">
        <Section size="1" py="3">
          <Flex align="center" justify="end">
            <PreferencesPopover />
          </Flex>
        </Section>
      </Container>
    </Box>
  );
}
