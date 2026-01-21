import type { ReactNode } from "react";
import type { LenisOptions } from "lenis";

import { ReactLenis } from "lenis/react";

import { useLenisRaf } from "./hooks/useLenisRaf";

interface LenisProviderProps {
  children: ReactNode;
  options?: LenisOptions;
  root?: boolean | "asChild";
}

export function LenisProvider({
  children,
  options,
  root = true,
}: LenisProviderProps) {
  const lenisRef = useLenisRaf({ enabled: true });
  const finalOptions: LenisOptions = { ...options, autoRaf: false };

  return (
    <ReactLenis root={root} options={finalOptions} ref={lenisRef}>
      {children}
    </ReactLenis>
  );
}
