import type { ReactNode } from "react";

import { memo, useCallback, useLayoutEffect, useRef, useState } from "react";

import Footer from "../Footer";
import Header from "../Header";
import styles from "./styles.module.css";

interface MainWrapperProps {
  children: (headerHeight: number, footerHeight: number) => ReactNode;
}

function useElementHeight<T extends HTMLElement>() {
  const ref = useRef<T | null>(null);
  const [height, setHeight] = useState(0);

  useLayoutEffect(() => {
    const node = ref.current;
    if (!node) return;

    const observer = new ResizeObserver(([entry]) => {
      const nextHeight = entry?.contentRect.height ?? 0;
      setHeight((prev) => (prev !== nextHeight ? nextHeight : prev));
    });

    observer.observe(node);
    return () => observer.disconnect();
  }, []);

  const callbackRef = useCallback((instance: T | null) => {
    if (instance) {
      ref.current = instance;
    }
  }, []);

  return [callbackRef, height] as const;
}

const MainWrapper = memo(({ children }: MainWrapperProps) => {
  const [headerRef, headerHeight] = useElementHeight<HTMLDivElement>();
  const [footerRef, footerHeight] = useElementHeight<HTMLDivElement>();

  return (
    <div className={styles.layout}>
      {/* Header */}
      <header ref={headerRef} className={styles.header}>
        <Header />
      </header>

        {children(headerHeight, footerHeight)}  

      {/* Footer */}
      <footer ref={footerRef} className={styles.footer}>
        <Footer />
      </footer>
    </div>
  );
});

MainWrapper.displayName = "MainWrapper";

export default MainWrapper;
