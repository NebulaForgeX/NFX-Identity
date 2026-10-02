import type { CSSProperties } from "react";

import { Section } from "@radix-ui/themes";
import { useLayoutStore } from "nfx-ui/stores";
import { Outlet } from "react-router";

import Header from "@/layouts/Header";

import styles from "./s.module.css";

function Main() {
  const headerHeight = useLayoutStore((state) => state.headerHeight);

  return (
    <Section
      size="1"
      pt="0"
      pb="0"
      position="relative"
      minHeight="100dvh"
      width="100%"
      className={styles.page}
      style={{ "--app-header-height": `${headerHeight}px` } as CSSProperties}
    >
      <Header />
      <main>
        <Outlet />
      </main>
    </Section>
  );
}

export default Main;
