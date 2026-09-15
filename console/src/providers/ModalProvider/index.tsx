import type { ReactNode } from "react";

import { memo } from "react";

import { BaseModal, ConfirmModal, LoadingModal, SearchModal, YearSelectModal } from "./components";

interface DialogHostProps {
  children: ReactNode;
}

const DialogHost = memo(({ children }: DialogHostProps) => {
  return (
    <>
      {children}
      <BaseModal />
      <ConfirmModal />
      <LoadingModal />
      <SearchModal />
      <YearSelectModal />
    </>
  );
});

DialogHost.displayName = "DialogHost";
export default DialogHost;
