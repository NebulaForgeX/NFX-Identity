import type { ReactNode } from "react";

import { Base, DateTimePicker } from "./components";
import { useSystemFeedbackInv } from "./hooks/useSystemFeedbackInv";

const ModalProvider = ({ children }: { children: ReactNode }) => {
  useSystemFeedbackInv();
  return (
    <>
      {children}
      <Base />
      <DateTimePicker />
    </>
  );
};

export default ModalProvider;
