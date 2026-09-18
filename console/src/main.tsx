import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import "@radix-ui/themes/styles.css";
import "nfx-ui/themes/fonts";
import "nfx-ui/themes/styles.css";

import { LanguageEnum } from "nfx-ui/enums";
import { LanguageProvider, ThemeProvider, ModalProvider, DataProvider } from "nfx-ui/providers";
import { LayoutProvider } from "nfx-ui/layouts";

import "./index.css";

import { ApiAssetRepository, ApiAuthRepository } from "nfx-ui/apis";
import { getBuiltinI18nBundles } from "@/assets/languages/i18nResources";
import { BootstrapProvider, QueryProvider, RouterProvider } from "@/providers";
import PreferenceSync from "@/providers/PreferenceSync";

import App from "./App.tsx";

const authErrors = new ApiAuthRepository();
const assetErrors = new ApiAssetRepository();

async function onLoadExtraBundles(lng: LanguageEnum) {
  try {
    const [authBundle, assetBundle] = await Promise.all([
      authErrors.GetErrorTranslations(lng).catch(() => ({})),
      assetErrors.GetErrorTranslations(lng).catch(() => ({})),
    ]);
    return { namespace: "errors", bundle: { ...assetBundle, ...authBundle } };
  } catch {
    return null;
  }
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryProvider>
      <LanguageProvider getBuiltinBundles={getBuiltinI18nBundles} fallbackLng={LanguageEnum.ZH} onLoadExtraBundles={onLoadExtraBundles}>
        <ThemeProvider>
          <LayoutProvider>
            <DataProvider>
              <PreferenceSync />
              <RouterProvider>
                <ModalProvider>
                  <BootstrapProvider>
                    <App />
                  </BootstrapProvider>
                </ModalProvider>
              </RouterProvider>
            </DataProvider>
          </LayoutProvider>
        </ThemeProvider>
      </LanguageProvider>
    </QueryProvider>
  </StrictMode>,
);
