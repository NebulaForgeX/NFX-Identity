import React, { memo } from "react";

import { useLayout } from "@/hooks/useLayout";

import SideHideLayout from "./SideHideLayout";
import SideShowLayout from "./SideShowLayout";
import MainWrapper from "./MainWrapper";

interface LayoutSwitcherProps {
  children: React.ReactNode;
}

export const LayoutSwitcher = memo(({ children }: LayoutSwitcherProps) => {
  const { layoutMode } = useLayout();
  return (
    <MainWrapper>
      {(headerHeight, footerHeight) => {
        if (layoutMode === "hide") {
          return (
            <SideHideLayout headerHeight={headerHeight} footerHeight={footerHeight}>
              {children}
            </SideHideLayout>
          );
        } else {
          return (
            <SideShowLayout headerHeight={headerHeight} footerHeight={footerHeight}>
              {children}
            </SideShowLayout>
          );
        }
      }}
    </MainWrapper>
  );
});

LayoutSwitcher.displayName = "LayoutSwitcher";

export default LayoutSwitcher;
