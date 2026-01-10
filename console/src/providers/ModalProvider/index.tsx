import type { ReactNode } from "react";

import { memo } from "react";

import { BaseModal, ConfirmModal, SearchModal, YearSelectModal } from "./components";

interface ModalProviderProps {
  children: ReactNode;
}

const ModalProvider = memo(({ children }: ModalProviderProps) => {
  return (
    <>
      {children}
      <BaseModal />
      <ConfirmModal />
      <SearchModal />
      <YearSelectModal />
    </>
  );
});

ModalProvider.displayName = "ModalProvider";
export default ModalProvider;
