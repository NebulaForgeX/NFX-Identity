import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { LanguageEnum, LanguageProvider } from "nfx-ui/languages";
import { ThemeEnum, ThemeProvider } from "nfx-ui/themes";
import { LayoutProvider } from "nfx-ui/layouts";

import "./index.css";
import "@/assets/themes/global.css";

import { NAME_SPACES, NAME_SPACES_MAP, RESOURCES } from "@/assets/languages/i18nResources";
import { getErrorTranslations } from "@/apis/system.api";
import { BootstrapProvider, BrowserRouterProvider, ModalProvider, QueryProvider, LenisProvider } from "@/providers";

import App from "./App.tsx";

async function loadErrorTranslationsBundle(
  lng: string,
): Promise<{ namespace: string; bundle: Record<string, unknown> } | null> {
  try {
    const bundle = await getErrorTranslations(lng);
    return { namespace: "errors", bundle: bundle as Record<string, unknown> };
  } catch {
    return null;
  }
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryProvider>
      <LanguageProvider
        bundles={{ RESOURCES, NAME_SPACES_MAP, NAME_SPACES }}
        fallbackLng={LanguageEnum.ZH}
        onLoadExtraBundles={async (lng) => loadErrorTranslationsBundle(lng)}
      >
        <ThemeProvider defaultTheme={ThemeEnum.DEFAULT}>
          <LayoutProvider>
            <BrowserRouterProvider>
              <LenisProvider>
                <ModalProvider>
                  <BootstrapProvider>
                    <App />
                  </BootstrapProvider>
                </ModalProvider>
              </LenisProvider>
            </BrowserRouterProvider>
          </LayoutProvider>
        </ThemeProvider>
      </LanguageProvider>
    </QueryProvider>
  </StrictMode>,
);
