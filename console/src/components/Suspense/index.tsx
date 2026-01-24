import type { QueryErrorResetBoundaryProps } from "@tanstack/react-query";
import type { ReactNode, SuspenseProps as ReactSuspenseProps } from "react";

import { memo, Suspense as ReactSuspense, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { QueryErrorResetBoundary } from "@tanstack/react-query";

import { BounceLoading, ECGLoading, TruckLoading } from "@/animations";

import styles from "./styles.module.css";
import SuspenseErrorBoundary from "./SuspenseErrorBoundary";

interface SuspenseProps extends Omit<ReactSuspenseProps, "fallback">, Omit<QueryErrorResetBoundaryProps, "children"> {
  fallback?: ReactNode;
  test?: boolean;
  loadingType?: "ecg" | "truck" | "bounce";
  loadingShape?: "square" | "circle";
  loadingText?: string;
  loadingSize?: "small" | "medium" | "large";
  loadingContainerClassName?: string;
  loadingTextClassName?: string;
  loadingClassName?: string;
  errorFallback?: (args: { error: Error | null; retry: () => void }) => ReactNode;
  errorTitle?: string;
  errorDescription?: string;
  retryText?: string;
  errorContainerClassName?: string;
  errorDetailsClassName?: string;
  showErrorDetails?: boolean;
}

const Suspense = memo((props: SuspenseProps) => {
  const { t } = useTranslation("components");
  const {
    fallback,
    test,
    loadingType = "ecg",
    loadingShape = "square",
    loadingClassName,
    loadingText,
    loadingSize = "medium",
    loadingContainerClassName,
    loadingTextClassName,
    children,
    errorFallback,
    errorTitle,
    errorDescription,
    retryText,
    errorContainerClassName,
    errorDetailsClassName,
    showErrorDetails = import.meta.env.DEV,
    ...restProps
  } = props;

  // 使用国际化文本作为默认值
  const defaultErrorTitle = errorTitle ?? t("suspense.errorTitle");
  const defaultErrorDescription = errorDescription ?? t("suspense.errorDescription");
  const defaultRetryText = retryText ?? t("suspense.retry");
  const defaultLoadingText = loadingText ?? t("suspense.loading");
  const renderLoading = useCallback(() => {
    switch (loadingType) {
      case "ecg":
        return <ECGLoading size={loadingSize} className={loadingClassName} />;
      case "truck":
        return <TruckLoading size={loadingSize} className={loadingClassName} />;
      case "bounce":
        return <BounceLoading size={loadingSize} shape={loadingShape} className={loadingClassName} />;
      default:
        return <ECGLoading size={loadingSize} className={loadingClassName} />;
    }
  }, [loadingType, loadingShape, loadingSize, loadingClassName]);

  const renderFallback = useCallback(() => {
    if (fallback) {
      return fallback;
    }
    return (
      <div className={`${styles.loadingContainer} ${loadingContainerClassName || ""}`}>
        {renderLoading()}
        <p className={`${styles.loadingText} ${loadingTextClassName || ""}`}>{defaultLoadingText}</p>
      </div>
    );
  }, [fallback, loadingContainerClassName, defaultLoadingText, loadingTextClassName, renderLoading]);

  const renderErrorFallback = useCallback(
    (error: Error | null, retry: () => void) => {
      if (errorFallback) return errorFallback({ error, retry });

      return (
        <div className={`${styles.errorContainer} ${errorContainerClassName || ""}`}>
          <h3 className={styles.errorTitle}>{defaultErrorTitle}</h3>
          <p className={styles.errorDescription}>{defaultErrorDescription}</p>
          {showErrorDetails && error && (
            <pre className={`${styles.errorDetails} ${errorDetailsClassName || ""}`}>{error.message}</pre>
          )}
          <button type="button" className={styles.retryButton} onClick={retry}>
            {defaultRetryText}
          </button>
        </div>
      );
    },
    [
      errorFallback,
      errorContainerClassName,
      defaultErrorDescription,
      defaultErrorTitle,
      defaultRetryText,
      showErrorDetails,
      errorDetailsClassName,
    ],
  );

  return (
    <QueryErrorResetBoundary {...restProps}>
      {({ reset }) => (
        <SuspenseErrorBoundary onReset={reset} fallbackRender={({ error, retry }) => renderErrorFallback(error, retry)}>
          <ReactSuspense fallback={renderFallback()} {...restProps}>
            {test ? <AlwaysPending /> : children}
          </ReactSuspense>
        </SuspenseErrorBoundary>
      )}
    </QueryErrorResetBoundary>
  );
});

Suspense.displayName = "Suspense";
export default Suspense;

const AlwaysPending = () => {
  throw new Promise(() => {}); // 永不 resolve 的 Promise
};
