import type { ReactNode } from "react";

import { Container, Section } from "@radix-ui/themes";
import clsx from "clsx";

import styles from "./s.module.css";

type Space = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";

export type SurfaceProps = {
  children: ReactNode;
  /** `panel` follows the theme panel preference; `hero` adds the accent wash; `inset` sits inside another surface. */
  tone?: "panel" | "hero" | "inset";
  interactive?: boolean;
  selected?: boolean;
  sticky?: boolean;
  py?: Space;
  px?: Space;
  className?: string;
};

/** One themed panel: Section owns vertical space and look, Container owns horizontal space. */
export default function Surface({ children, tone = "panel", interactive, selected, sticky, py = "5", px = "5", className }: SurfaceProps) {
  return (
    <Section
      size="1"
      py={py}
      data-reveal={tone === "inset" ? undefined : ""}
      data-selected={selected ? "true" : undefined}
      className={clsx(styles.surface, styles[tone], interactive && styles.interactive, sticky && styles.sticky, className)}
    >
      <Container size="4" width="100%" maxWidth="100%" px={px}>
        {children}
      </Container>
    </Section>
  );
}
