import type { ReactNode } from "react";

import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { Box } from "@radix-ui/themes";
import gsap from "gsap";
import { PreferencesPopover } from "nfx-ui/components";

import styles from "./AuthShell.module.css";

gsap.registerPlugin(useGSAP);

export type AuthShellProps = {
  brandTitle: string;
  brandEyebrow?: string;
  heroFooter: string;
  children: ReactNode;
};

export default function AuthShell({ brandTitle, brandEyebrow, heroFooter, children }: AuthShellProps) {
  const pageRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) return;
      gsap.set(".js-booklet", { autoAlpha: 0, y: 36, rotateX: 8, transformOrigin: "50% 100%" });
      gsap.set(".js-mark", { autoAlpha: 0, letterSpacing: "0.8em" });
      gsap.set(".js-mrz", { autoAlpha: 0, x: 24 });
      const tl = gsap.timeline({ defaults: { ease: "power3.out" } });
      tl.to(".js-mark", { autoAlpha: 1, letterSpacing: "0.42em", duration: 0.7 })
        .to(".js-booklet", { autoAlpha: 1, y: 0, rotateX: 0, duration: 0.85 }, "-=0.25")
        .to(".js-mrz", { autoAlpha: 1, x: 0, duration: 0.55 }, "-=0.3");
    },
    { scope: pageRef },
  );

  return (
    <Box ref={pageRef} className={styles.page} asChild>
      <main>
        <div className={styles.grain} aria-hidden />
        <div className={styles.desk}>
          <div className={styles.toolbar}>
            <PreferencesPopover />
          </div>
          <p className={`${styles.mark} js-mark`}>{brandEyebrow ?? "NFX Identity"}</p>
          <section className={`${styles.booklet} js-booklet`}>
            <div className={styles.bookletHead}>
              <span>PASSPORT DESK</span>
              <span className={styles.serial}>NFX · 00 19 72</span>
            </div>
            <h1 className={styles.title}>{brandTitle}</h1>
            <p className={styles.lede}>{heroFooter}</p>
            <div className={styles.formWell}>{children}</div>
            <p className={`${styles.mrz} js-mrz`} aria-hidden>
              P&lt;NFXIDENTITY&lt;&lt;DIRECTORY&lt;&lt;&lt;ACCESS&lt;&lt;&lt;CONSOLE
            </p>
          </section>
        </div>
      </main>
    </Box>
  );
}
