import React, { memo } from "react";

import { useLayoutStore } from "@/stores/layoutStore";

import SideHideLayout from "./SideHideLayout";
import SideShowLayout from "./SideShowLayout";
import MainWrapper from "./MainWrapper";

interface LayoutSwitcherProps {
  children: React.ReactNode;
}

export const LayoutSwitcher = memo(({ children }: LayoutSwitcherProps) => {
  const layoutMode = useLayoutStore((s) => s.layoutMode);
  
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
