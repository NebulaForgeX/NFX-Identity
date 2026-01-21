import type { LenisRef } from "lenis/react";
import { useEffect, useRef } from "react";
import gsap from "gsap";
import { cancelFrame, frame } from "framer-motion";

interface UseLenisRafOptions {
  /**
   * Whether the RAF is enabled
   * @default true
   */
  enabled?: boolean;
}

/**
 * Hook for managing Lenis RAF (requestAnimationFrame) integration
 * Supports three modes simultaneously:
 * - Custom requestAnimationFrame loop
 * - GSAP ticker integration
 * - Framer Motion frame integration
 * 
 * @returns The Lenis ref that should be passed to ReactLenis component
 */
export function useLenisRaf(options: UseLenisRafOptions = {}) {
  const { enabled = true } = options;
  const lenisRef = useRef<LenisRef>(null);
  const rafIdRef = useRef<number | null>(null);

  useEffect(() => {
    if (!enabled || !lenisRef.current) {
      return;
    }

    // Custom requestAnimationFrame loop
    function rafUpdate(time: number) {
      lenisRef.current?.lenis?.raf(time);
      rafIdRef.current = requestAnimationFrame(rafUpdate);
    }

    // GSAP integration
    function gsapUpdate(time: number) {
      lenisRef.current?.lenis?.raf(time * 1000);
    }

    // Framer Motion integration
    function framerUpdate(data: { timestamp: number }) {
      const time = data.timestamp;
      lenisRef.current?.lenis?.raf(time);
    }

    // Start all three
    rafIdRef.current = requestAnimationFrame(rafUpdate);
    gsap.ticker.add(gsapUpdate);
    frame.update(framerUpdate, true);

    // Cleanup
    return () => {
      if (rafIdRef.current !== null) {
        cancelAnimationFrame(rafIdRef.current);
        rafIdRef.current = null;
      }
      gsap.ticker.remove(gsapUpdate);
      cancelFrame(framerUpdate);
    };
  }, [enabled]);

  return lenisRef;
}
