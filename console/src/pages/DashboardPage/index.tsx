import { memo, useMemo } from "react";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserPreferenceSync } from "@/hooks/useUserPreferenceSync";
import { useUserPreferenceNormal } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import QuickStore, { useQuickStore } from "@/stores/quickStore";
import { WaveBackground, SquareBackground, LetterGlitchBackground, PixelBlastBackground } from "@/animations";

import { QuickNavigation, ResourceLinks, StatsCards } from "./components";
import styles from "./styles.module.css";

const DashboardPage = memo(() => {
  // Sync and apply user preferences on dashboard load (after login)
  useUserPreferenceSync();
  const isEditMode = useQuickStore((state) => state.isEditMode);
  const toggleEditMode = QuickStore.getState().toggleEditMode;
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
      <div className={styles.container}>
      <div className={styles.header}>
        <h1 className={styles.title}>仪表盘</h1>
        <button
          className={styles.editButton}
          onClick={toggleEditMode}
          aria-label={isEditMode ? "退出编辑模式" : "编辑快速导航"}
          title={isEditMode ? "退出编辑模式" : "编辑快速导航"}
        >
          <Edit size={20} />
        </button>
      </div>

      {/* Stats Cards */}
      <Suspense
        loadingType="ecg"
        loadingText="加载统计数据..."
        loadingSize="small"
        loadingContainerClassName={styles.loading}
      >
        <StatsCards />
      </Suspense>

      {/* Quick Navigation */}
      <QuickNavigation />

      {/* Resources & Policies Section */}
      <ResourceLinks />
      </div>

      {renderBackground()}
    </>
  );
});

DashboardPage.displayName = "DashboardPage";

export default DashboardPage;
