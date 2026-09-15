import type { ReactNode } from "react";

import { memo } from "react";

import { LetterGlitchBackground, PixelBlastBackground, SquareBackground, WaveBackground } from "@/animations";
import { DashboardBackgroundEnum, DEFAULT_DASHBOARD_BACKGROUND } from "nfx-ui/enums";
import { usePreferenceStore } from "nfx-ui/stores";

import styles from "./styles.module.css";

interface BackgroundProps {
  children: ReactNode;
}

const Background = memo(({ children }: BackgroundProps) => {
  const dashboardBackground = usePreferenceStore((s) => s.dashboardBackground) || DEFAULT_DASHBOARD_BACKGROUND;

  const renderBackground = () => {
    switch (dashboardBackground) {
      case DashboardBackgroundEnum.WAVES:
        return (
          <div className={styles.wavesWrapper}>
            <WaveBackground />
          </div>
        );
      case DashboardBackgroundEnum.SQUARES:
        return (
          <div className={styles.squaresWrapper}>
            <SquareBackground />
          </div>
        );
      case DashboardBackgroundEnum.LETTER_GLITCH:
        return (
          <div className={styles.letterGlitchWrapper}>
            <LetterGlitchBackground />
          </div>
        );
      case DashboardBackgroundEnum.PIXEL_BLAST:
        return (
          <div className={styles.pixelBlastWrapper}>
            <PixelBlastBackground />
          </div>
        );
      case DashboardBackgroundEnum.NONE:
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
