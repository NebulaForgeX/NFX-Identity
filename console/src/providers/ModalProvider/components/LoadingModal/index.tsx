import { memo } from "react";

import { TruckLoading } from "@/animations";
import { useModalStore } from "@/stores/modalStore";

import styles from "./styles.module.css";

const LoadingModal = memo(() => {
  const isOpen = useModalStore((state) => state.loadingModal.isOpen);
  const message = useModalStore((state) => state.loadingModal.message);

  if (!isOpen) return null;

  return (
    <div className={styles.overlay} role="alert" aria-busy="true" aria-live="polite">
      <div className={styles.content}>
        <TruckLoading size="large" />
        {message && <p className={styles.message}>{message}</p>}
      </div>
    </div>
  );
});

LoadingModal.displayName = "LoadingModal";

export default LoadingModal;
