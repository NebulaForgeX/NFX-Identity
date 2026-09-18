import { createStore, useStore } from "zustand";
import { persist, subscribeWithSelector } from "zustand/middleware";

import { ROUTES } from "@/navigations";

export interface QuickNavItem {
  id: string;
  title: string;
  description: string;
  icon: string;
  route: string;
  color: string;
}

export const ALL_AVAILABLE_ITEMS: QuickNavItem[] = [
  {
    id: "profile",
    title: "Profile",
    description: "Emails, phones, avatar, background",
    icon: "User",
    route: ROUTES.PROFILE,
    color: "var(--color-primary)",
  },
  {
    id: "assets",
    title: "Assets",
    description: "Images, files, videos, audios",
    icon: "Image",
    route: ROUTES.IMAGES,
    color: "var(--color-success)",
  },
  {
    id: "settings",
    title: "Settings",
    description: "Theme and console preferences",
    icon: "Settings",
    route: ROUTES.USER_SETTINGS,
    color: "var(--color-info)",
  },
];

interface QuickState {
  isEditMode: boolean;
  items: QuickNavItem[];
}

interface QuickActions {
  setEditMode: (editMode: boolean) => void;
  toggleEditMode: () => void;
  addItem: (item: QuickNavItem) => void;
  removeItem: (id: string) => void;
  updateItem: (id: string, item: Partial<QuickNavItem>) => void;
  reorderItems: (items: QuickNavItem[]) => void;
  resetItems: () => void;
}

const defaultItems: QuickNavItem[] = ALL_AVAILABLE_ITEMS.slice();

const defaultState: QuickState = {
  isEditMode: false,
  items: defaultItems,
};

export const QuickStore = createStore<QuickState & QuickActions>()(
  subscribeWithSelector(
    persist(
      (set) => ({
        ...defaultState,
        setEditMode: (editMode) => set({ isEditMode: editMode }),
        toggleEditMode: () => set((state) => ({ isEditMode: !state.isEditMode })),
        addItem: (item) => set((state) => ({ items: [...state.items, item] })),
        removeItem: (id) => set((state) => ({ items: state.items.filter((item) => item.id !== id) })),
        updateItem: (id, updatedItem) =>
          set((state) => ({
            items: state.items.map((item) => (item.id === id ? { ...item, ...updatedItem } : item)),
          })),
        reorderItems: (items) => set({ items }),
        resetItems: () => set({ items: defaultItems }),
      }),
      {
        name: "quick-nav-storage",
        version: 3,
        partialize: (state) => ({ items: state.items }),
        migrate: (_persistedState, version) => {
          if (version < 3) {
            return { items: defaultItems };
          }
          return _persistedState as QuickState;
        },
      },
    ),
  ),
);

export default QuickStore;
export const useQuickStore = <T>(selector: (state: QuickState & QuickActions) => T) => useStore(QuickStore, selector);
