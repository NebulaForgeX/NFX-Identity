import { createStore, useStore } from "zustand";
import { subscribeWithSelector } from "zustand/middleware";

type ModalType = "success" | "error" | "info" | "confirm" | "search" | "yearSelect";

interface BaseModalProps {
  isOpen: boolean;
  message?: string;
  title?: string;
  confirmText?: string;
  onClick?: () => void;
}

interface ConfirmModalProps {
  isOpen: boolean;
  message?: string;
  title?: string;
  onConfirm?: () => void;
  onCancel?: () => void;
  confirmText?: string;
  cancelText?: string;
}

interface SearchModalProps {
  isOpen: boolean;
  title?: string;
}

interface YearSelectModalProps {
  isOpen: boolean;
  initialYear?: number;
  minOffset?: number;
  maxOffset?: number;
  onSelect?: (year: number) => void;
}

interface ModalState {
  modalType: ModalType;
  baseModal: BaseModalProps;
  confirmModal: ConfirmModalProps;
  searchModal: SearchModalProps;
  yearSelectModal: YearSelectModalProps;
}

interface ModalActions {
  showModal: (
    modalType: ModalType,
    props: BaseModalProps | ConfirmModalProps | SearchModalProps | YearSelectModalProps,
  ) => void;
  hideModal: (modalType?: ModalType) => void; // undefined 表示关闭所有模态框
}

const defaultBaseModalProps: BaseModalProps = {
  isOpen: false,
  message: "No message",
  title: "No title",
  confirmText: "Confirm",
  onClick: undefined,
};
const defaultConfirmModalProps: ConfirmModalProps = {
  isOpen: false,
  message: "",
  title: undefined,
  onConfirm: undefined,
  onCancel: undefined,
  confirmText: "Confirm",
  cancelText: "Cancel",
};
const defaultSearchModalProps: SearchModalProps = {
  isOpen: false,
  title: "Search",
};
const defaultYearSelectModalProps: YearSelectModalProps = {
  isOpen: false,
  initialYear: undefined,
  minOffset: -500,
  maxOffset: 500,
  onSelect: undefined,
};

export const ModalStore = createStore<ModalState & ModalActions>()(
  subscribeWithSelector((set) => ({
    modalType: "info",
    baseModal: defaultBaseModalProps,
    confirmModal: defaultConfirmModalProps,
    searchModal: defaultSearchModalProps,
    yearSelectModal: defaultYearSelectModalProps,

    showModal: (modalType, props) => {
      // 根据 modalType 设置对应的模态框状态
      if (modalType === "success" || modalType === "error" || modalType === "info") {
        const { isOpen, ...restProps } = props as BaseModalProps;
        set({
          modalType,
          baseModal: {
            isOpen: true,
            ...restProps,
          },
        });
      } else if (modalType === "confirm") {
        const { isOpen, ...restProps } = props as ConfirmModalProps;
        set({
          modalType,
          confirmModal: {
            isOpen: true,
            ...restProps,
          },
        });
      } else if (modalType === "search") {
        const { isOpen, ...restProps } = props as SearchModalProps;
        set({
          modalType,
          searchModal: {
            isOpen: true,
            ...restProps,
          },
        });
      } else if (modalType === "yearSelect") {
        const { isOpen, ...restProps } = props as YearSelectModalProps;
        set({
          modalType,
          yearSelectModal: {
            isOpen: true,
            ...restProps,
          },
        });
      }
    },

    hideModal: (modalType) => {
      // 如果 modalType 为 undefined，关闭所有模态框
      if (modalType === undefined) {
        set({
          modalType: undefined,
          baseModal: defaultBaseModalProps,
          confirmModal: defaultConfirmModalProps,
          searchModal: defaultSearchModalProps,
          yearSelectModal: defaultYearSelectModalProps,
        });
        return;
      }
      // 根据 modalType 只关闭对应的模态框
      if (modalType === "success" || modalType === "error" || modalType === "info") {
        set({
          modalType: undefined,
          baseModal: defaultBaseModalProps,
        });
      } else if (modalType === "confirm") {
        set({
          modalType: undefined,
          confirmModal: defaultConfirmModalProps,
        });
      } else if (modalType === "search") {
        set({
          modalType: undefined,
          searchModal: defaultSearchModalProps,
        });
      } else if (modalType === "yearSelect") {
        set({
          modalType: undefined,
          yearSelectModal: defaultYearSelectModalProps,
        });
      }
    },
  })),
);

export default ModalStore;
export const useModalStore = <T>(selector: (state: ModalState) => T) => useStore(ModalStore, selector);

export const showInfo = (message: string, title?: string) => {
  ModalStore.getState().showModal("info", {
    isOpen: true,
    message,
    title,
  });
};

export interface ShowSuccessProps {
  message: string;
  title?: string;
  onClick?: () => void;
}

export const showSuccess = (props: ShowSuccessProps | string) => {
  // 兼容旧的字符串参数形式
  if (typeof props === "string") {
    ModalStore.getState().showModal("success", {
      isOpen: true,
      message: props,
    });
    return;
  }

  ModalStore.getState().showModal("success", {
    isOpen: true,
    message: props.message,
    title: props.title,
    onClick: props.onClick,
  });
};

export const showError = (message: string, title?: string) => {
  ModalStore.getState().showModal("error", {
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

export const showConfirm = (props: ShowConfirmProps) => {
  ModalStore.getState().showModal("confirm", {
    isOpen: true,
    message: props.message,
    title: props.title,
    onConfirm: props.onConfirm,
    onCancel: props.onCancel,
    confirmText: props.confirmText,
    cancelText: props.cancelText,
  });
};

export const showSearch = () => {
  ModalStore.getState().showModal("search", {
    isOpen: true,
  });
};

export interface ShowYearSelectProps {
  initialYear?: number;
  minOffset?: number;
  maxOffset?: number;
  onSelect: (year: number) => void;
}

export const showYearSelect = (props: ShowYearSelectProps) => {
  ModalStore.getState().showModal("yearSelect", {
    isOpen: true,
    initialYear: props.initialYear,
    minOffset: props.minOffset,
    maxOffset: props.maxOffset,
    onSelect: props.onSelect,
  });
};
