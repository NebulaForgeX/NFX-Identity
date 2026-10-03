import type { RefObject } from "react";

import { useGSAP } from "@gsap/react";
import gsap from "gsap";

gsap.registerPlugin(useGSAP);

/** Staggers every `[data-reveal]` node inside `scope` into view once per mount. */
export default function usePanelReveal(scope: RefObject<Nullable<HTMLElement>>) {
  useGSAP(
    () => {
      const mm = gsap.matchMedia();
      mm.add("(prefers-reduced-motion: no-preference)", () => {
        const nodes = gsap.utils.toArray<HTMLElement>("[data-reveal]");
        if (!nodes.length) return;
        gsap.fromTo(
          nodes,
          { autoAlpha: 0, y: 14 },
          { autoAlpha: 1, y: 0, duration: 0.55, ease: "power3.out", stagger: 0.06, clearProps: "transform,opacity,visibility" },
        );
      });
      return () => mm.revert();
    },
    { scope },
  );
}
