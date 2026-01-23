import { memo } from "react";

import { BounceLoading } from "@/animations";
import ModalStore, { useModalStore } from "@/stores/modalStore";

import styles from "./styles.module.css";

const LoadingModal = memo(() => {
  const isOpen = useModalStore((state) => state.loadingModal.isOpen);
  const message = useModalStore((state) => state.loadingModal.message);
  const canClose = useModalStore((state) => state.loadingModal.canClose);

  if (!isOpen) return null;

  const handleClose = () => {
    if (canClose) {
      ModalStore.getState().hideModal("loading");
    }
  };

  return (
    <div
      className={styles.overlay}
      role="alert"
      aria-busy="true"
      aria-live="polite"
      onClick={canClose ? handleClose : undefined}
      style={{ cursor: canClose ? "pointer" : "default" }}
    >
      <div className={styles.content} onClick={(e) => e.stopPropagation()}>
        <BounceLoading size="medium" shape="circle" />
        {message && <p className={styles.message}>{message}</p>}
        {canClose && (
          <button className={styles.closeButton} onClick={handleClose} type="button">
            ✕
          </button>
        )}
      </div>
    </div>
  );
});

LoadingModal.displayName = "LoadingModal";

export default LoadingModal;
