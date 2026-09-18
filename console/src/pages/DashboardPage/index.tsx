import { memo } from "react";
import { Button, Flex } from "@radix-ui/themes";
import { Edit, LayoutDashboard } from "lucide-react";
import { PageHeader, Suspense } from "@/components";
import { PageFrame } from "@/layouts";
import QuickStore, { useQuickStore } from "@/stores/quickStore";

import QuickNavigation from "./components/QuickNavigation";
import ResourceLinks from "./components/ResourceLinks";
import StatsCards from "./components/StatsCards";

const DashboardPage = memo(() => {
  const isEditMode = useQuickStore((state) => state.isEditMode);
  const toggleEditMode = QuickStore.getState().toggleEditMode;

  return (
    <PageFrame>
        <PageHeader
          icon={LayoutDashboard}
          title="仪表盘"
          actions={
            <Button variant="soft" onClick={toggleEditMode} aria-label={isEditMode ? "退出编辑模式" : "编辑快速导航"}>
              <Edit size={16} />
            </Button>
          }
        />
        <Flex direction="column" gap="5">
          <Suspense loadingText="加载统计数据...">
            <StatsCards />
          </Suspense>
          <QuickNavigation />
          <ResourceLinks />
        </Flex>
      </PageFrame>
  );
});

DashboardPage.displayName = "DashboardPage";
export default DashboardPage;
