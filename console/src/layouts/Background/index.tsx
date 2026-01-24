import type { ReactNode } from "react";

import { memo, useMemo } from "react";

import { WaveBackground, SquareBackground, LetterGlitchBackground, PixelBlastBackground } from "@/animations";
import { useUserPreferenceNormal } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";

import styles from "./styles.module.css";

interface BackgroundProps {
  children: ReactNode;
}

const Background = memo(({ children }: BackgroundProps) => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);

  // Get user preference for dashboard background
  const shouldFetch = !!currentUserId && isAuthValid;
  const { data: preference } = useUserPreferenceNormal({
    id: currentUserId || "00000000-0000-0000-0000-000000000000",
    options: {
      enabled: shouldFetch && !!currentUserId,
    },
  });

  // Determine which background to show (从 other 字段中读取)
  const dashboardBackground = useMemo(() => {
    if (!preference?.other) return "none";
    const other = preference.other as Record<string, unknown>;
    return (other.dashboardBackground as string) || "none";
  }, [preference]);

  // Render background component based on preference
  const renderBackground = () => {
    switch (dashboardBackground) {
      case "waves":
        return (
          <div className={styles.wavesWrapper}>
            <WaveBackground />
          </div>
        );
      case "squares":
        return (
          <div className={styles.squaresWrapper}>
            <SquareBackground />
          </div>
        );
      case "letterGlitch":
        return (
          <div className={styles.letterGlitchWrapper}>
            <LetterGlitchBackground />
          </div>
        );
      case "pixelBlast":
        return (
          <div className={styles.pixelBlastWrapper}>
            <PixelBlastBackground />
          </div>
        );
      case "none":
      default:
        return null;
    }
  };

  return (
    <>
      {children}
      {renderBackground()}
    </>
  );
});

Background.displayName = "Background";

export default Background;
