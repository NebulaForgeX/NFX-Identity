import { makeStore } from "@/utils";

type ModalType = "success" | "error" | "info" | "dateTimePicker" | "loading";

interface BaseModalProps {
  isOpen: boolean;
  message?: string;
  title?: string;
  confirmText?: string;
  onClick?: () => void;
  variant?: ModalType;
}

export interface DateTimePickerModalProps {
  isOpen: boolean;
  value: string;
  title?: string;
  minDate: Date | null;
  maxDate: Date | null;
  allowClear: boolean;
  onConfirm?: (value: string) => void;
  onCancel?: () => void;
}

interface LoadingModalProps {
  isOpen: boolean;
  message?: string;
  title?: string;
}

interface ModalState {
  modalType: ModalType | undefined;
  baseModal: BaseModalProps;
  dateTimePickerModal: DateTimePickerModalProps;
  loadingModal: LoadingModalProps;
}

interface ModalActions {
  showModal: (modalType: ModalType, props: BaseModalProps | DateTimePickerModalProps) => void;
  hideModal: (modalType?: ModalType) => void;
}

const defaultBaseModalProps: BaseModalProps = {
  isOpen: false,
  message: "No message",
  title: "No title",
  confirmText: "Confirm",
  onClick: undefined,
};

const defaultLoadingModalProps: LoadingModalProps = { isOpen: false };

const defaultDateTimePickerModalProps: DateTimePickerModalProps = {
  isOpen: false,
  value: "",
  title: undefined,
  minDate: null,
  maxDate: null,
  allowClear: true,
  onConfirm: undefined,
  onCancel: undefined,
};

function omitIsOpen<T extends { isOpen?: boolean }>(props: T): Omit<T, "isOpen"> {
  const rest = { ...props };
  delete rest.isOpen;
  return rest as Omit<T, "isOpen">;
}

const { store: ModalStore, useStore: useModalStore } = makeStore<ModalState, ModalActions>(
  {
    modalType: undefined,
    baseModal: defaultBaseModalProps,
    dateTimePickerModal: defaultDateTimePickerModalProps,
    loadingModal: defaultLoadingModalProps,
  },
  (set) => ({
    showModal: (modalType, props) => {
      if (modalType === "dateTimePicker") {
        const restProps = omitIsOpen(props as DateTimePickerModalProps);
        set({
          modalType,
          dateTimePickerModal: {
            ...defaultDateTimePickerModalProps,
            ...restProps,
            isOpen: true,
          },
        });
        return;
      }
      const restProps = omitIsOpen(props as BaseModalProps);
      set({
        modalType,
        baseModal: {
          isOpen: true,
          variant: modalType,
          ...restProps,
        },
      });
    },

    hideModal: (modalType) => {
      if (modalType === "dateTimePicker") {
        set({ dateTimePickerModal: defaultDateTimePickerModalProps });
        return;
      }
      if (modalType === undefined) {
        set({
          modalType: undefined,
          baseModal: defaultBaseModalProps,
          dateTimePickerModal: defaultDateTimePickerModalProps,
          loadingModal: defaultLoadingModalProps,
        });
        return;
      }
      set({
        modalType: undefined,
        baseModal: defaultBaseModalProps,
      });
    },
  }),
);

const showModal = ModalStore.getState().showModal;
const hideModal = ModalStore.getState().hideModal;

export { showModal, hideModal };
export { ModalStore, useModalStore };
export default ModalStore;

export const showInfo = (message: string, title?: string) => {
  showModal("info", {
    isOpen: true,
    message,
    title,
  });
};

export type ShowSuccessProps = {
  message: string;
  title?: string;
  onClick?: () => void;
};

export const showSuccess = (props: ShowSuccessProps | string) => {
  if (typeof props === "string") {
    showModal("success", {
      isOpen: true,
      message: props,
    });
    return;
  }
  showModal("success", {
    isOpen: true,
    message: props.message,
    title: props.title,
    onClick: props.onClick,
  });
};

export const showError = (message: string, title?: string) => {
  showModal("error", {
    isOpen: true,
    message,
    title,
  });
};

export interface ShowConfirmProps {
  message: string;
  onConfirm: () => void;
  onCancel?: () => void;
  title?: string;
  confirmText?: string;
  cancelText?: string;
}

export interface ShowDateTimePickerModalProps {
  value?: string;
  title?: string;
  minDate?: Date | null;
  maxDate?: Date | null;
  allowClear?: boolean;
  onConfirm: (value: string) => void;
  onCancel?: () => void;
}

export const showDateTimePickerModal = (props: ShowDateTimePickerModalProps) => {
  showModal("dateTimePicker", {
    isOpen: true,
    value: props.value ?? "",
    title: props.title,
    minDate: props.minDate ?? null,
    maxDate: props.maxDate ?? null,
    allowClear: props.allowClear ?? true,
    onConfirm: props.onConfirm,
    onCancel: props.onCancel,
  });
};

export const showLoading = (props?: { message?: string; title?: string }) => {
  ModalStore.setState({
    loadingModal: { isOpen: true, message: props?.message, title: props?.title },
  });
};

export const hideLoading = () => {
  ModalStore.setState({ loadingModal: defaultLoadingModalProps });
};

export const showConfirm = (props: ShowConfirmProps) => {
  showModal("info", {
    isOpen: true,
    message: props.message,
    title: props.title,
    confirmText: props.confirmText,
    onClick: props.onConfirm,
  });
};
