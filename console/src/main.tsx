import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import "@radix-ui/themes/styles.css";
import "nfx-ui/themes/fonts";
import "nfx-ui/themes/styles.css";

import { LanguageEnum } from "nfx-ui/enums";
import { LanguageProvider, ThemeProvider, ModalProvider, DataProvider } from "nfx-ui/providers";
import { LayoutProvider } from "nfx-ui/layouts";

import "./index.css";
import "@/assets/themes/global.css";

import { getBuiltinI18nBundles } from "@/assets/languages/i18nResources";
import { DialogHost, QueryProvider, RouterProvider } from "@/providers";

import App from "./App.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryProvider>
      <LanguageProvider getBuiltinBundles={getBuiltinI18nBundles} fallbackLng={LanguageEnum.ZH}>
        <ThemeProvider>
          <LayoutProvider>
            <DataProvider>
              <RouterProvider>
                <ModalProvider>
                  <DialogHost>
                    <App />
                  </DialogHost>
                </ModalProvider>
              </RouterProvider>
            </DataProvider>
          </LayoutProvider>
        </ThemeProvider>
      </LanguageProvider>
    </QueryProvider>
  </StrictMode>,
);
