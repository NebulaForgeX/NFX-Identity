import type { QueryErrorResetBoundaryProps } from "@tanstack/react-query";
import type { ReactNode, SuspenseProps as ReactSuspenseProps } from "react";

import { memo, Suspense as ReactSuspense } from "react";
import { useTranslation } from "react-i18next";
import { QueryErrorResetBoundary } from "@tanstack/react-query";

import { AlertCircle } from "@/assets/icons/lucide";
import { BounceLoading, ECGLoading, TruckLoading } from "@/animations";

import styles from "./styles.module.css";
import SuspenseErrorBoundary from "./SuspenseErrorBoundary";

interface LoadingFallbackProps {
  fallback?: ReactNode;
  loadingType?: "ecg" | "truck" | "bounce";
  loadingShape?: "square" | "circle";
  loadingText?: string;
  loadingSize?: "small" | "medium" | "large";
  loadingContainerClassName?: string;
  loadingTextClassName?: string;
  loadingClassName?: string;
}

interface ErrorFallbackProps {
  error: Error | null;
  retry: () => void;
  errorFallback?: (args: { error: Error | null; retry: () => void }) => ReactNode;
  errorTitle?: string;
  errorDescription?: string;
  retryText?: string;
  errorContainerClassName?: string;
  errorDetailsClassName?: string;
  showErrorDetails?: boolean;
}

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

const LoadingFallback = memo((props: LoadingFallbackProps) => {
  const {
    fallback,
    loadingType = "ecg",
    loadingShape = "square",
    loadingText,
    loadingSize = "medium",
    loadingContainerClassName,
    loadingTextClassName,
    loadingClassName,
  } = props;

  const { t } = useTranslation("components");
  const defaultLoadingText = loadingText ?? t("suspense.loading");

  if (fallback) {
    return <>{fallback}</>;
  }

  const renderLoading = () => {
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
  };

  return (
    <div className={`${styles.loadingContainer} ${loadingContainerClassName || ""}`}>
      {renderLoading()}
      <p className={`${styles.loadingText} ${loadingTextClassName || ""}`}>{defaultLoadingText}</p>
    </div>
  );
});

LoadingFallback.displayName = "LoadingFallback";

const ErrorFallback = memo((props: ErrorFallbackProps) => {
  const {
    error,
    retry,
    errorFallback,
    errorTitle,
    errorDescription,
    retryText,
    errorContainerClassName,
    errorDetailsClassName,
    showErrorDetails = import.meta.env.DEV,
  } = props;

  const { t } = useTranslation("components");

  const defaultErrorTitle = errorTitle ?? t("suspense.errorTitle");
  const defaultErrorDescription = errorDescription ?? t("suspense.errorDescription");
  const defaultRetryText = retryText ?? t("suspense.retry");
  const defaultErrorDetailsText = t("suspense.errorDetails");

  if (errorFallback) {
    return <>{errorFallback({ error, retry })}</>;
  }

  return (
    <div className={`${styles.errorContainer} ${errorContainerClassName || ""}`}>
      <div className={styles.errorIconWrapper}>
        <AlertCircle className={styles.errorIcon} size={64} />
      </div>
      <div className={styles.errorContent}>
        <h3 className={styles.errorTitle}>{defaultErrorTitle}</h3>
        <p className={styles.errorDescription}>{defaultErrorDescription}</p>
        {showErrorDetails && error && (
          <div className={styles.errorDetailsWrapper}>
            <details className={styles.errorDetailsContainer}>
              <summary className={styles.errorDetailsSummary}>{defaultErrorDetailsText}</summary>
              <pre className={`${styles.errorDetails} ${errorDetailsClassName || ""}`}>{error.message}</pre>
            </details>
          </div>
        )}
        <button type="button" className={styles.retryButton} onClick={retry}>
          <span className={styles.retryButtonIcon}>↻</span>
          <span>{defaultRetryText}</span>
        </button>
      </div>
    </div>
  );
});

ErrorFallback.displayName = "ErrorFallback";

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

  return (
    <QueryErrorResetBoundary {...restProps}>
      {({ reset }) => (
        <SuspenseErrorBoundary
          onReset={reset}
          fallbackRender={({ error, retry }) => (
            <ErrorFallback
              error={error}
              retry={retry}
              errorFallback={errorFallback}
              errorTitle={defaultErrorTitle}
              errorDescription={defaultErrorDescription}
              retryText={defaultRetryText}
              errorContainerClassName={errorContainerClassName}
              errorDetailsClassName={errorDetailsClassName}
              showErrorDetails={showErrorDetails}
            />
          )}
        >
          <ReactSuspense
            fallback={
              <LoadingFallback
                fallback={fallback}
                loadingType={loadingType}
                loadingShape={loadingShape}
                loadingText={loadingText}
                loadingSize={loadingSize}
                loadingContainerClassName={loadingContainerClassName}
                loadingTextClassName={loadingTextClassName}
                loadingClassName={loadingClassName}
              />
            }
            {...restProps}
          >
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
