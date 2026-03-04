/**
 * 公共组件：从 nfx-ui 再导出，与 NFX-Vault 一致；仅保留 IconButton、ThemeSwitcher、LanguageSwitcher、LayoutSwitcher 本地。
 */
export type {
  ButtonProps,
  DropdownOption,
  DropdownProps,
  IconName,
  IconProps,
  InputProps,
  KeyValueEditorProps,
  KeyValuePair,
  SearchInputProps,
  ShowFilterProps,
  ShowFilterValue,
  SliderProps,
  SuspenseProps,
  TextareaProps,
  VirtualListProps,
  VirtualWindowListProps,
} from "nfx-ui/components";

export {
  Button,
  Dropdown,
  Icon,
  Input,
  KeyValueEditor,
  SearchInput,
  ShowFilter,
  Slider,
  Suspense,
  Textarea,
  VirtualList,
  VirtualWindowList,
} from "nfx-ui/components";

export {
  BounceLoading,
  ECGLoading,
  LetterGlitchBackground,
  PixelBlastBackground,
  SquareBackground,
  TruckLoading,
  WaveBackground,
} from "nfx-ui/animations";

export { default as ThemeSwitcher } from "./ThemeSwitcher";
export { default as LanguageSwitcher } from "./LanguageSwitcher";
export { default as LayoutSwitcher } from "./LayoutSwitcher";
export { default as IconButton } from "./IconButton";
export type { IconButtonProps } from "./IconButton";
