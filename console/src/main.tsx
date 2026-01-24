import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import "./index.css";
import "@/assets/themes/global.css";
import "@/assets/languages/i18n";

import { BootstrapProvider, BrowserRouterProvider, LayoutProvider, ModalProvider, QueryProvider, ThemeProvider, LenisProvider } from "@/providers";  

import App from "./App.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryProvider>
      <ThemeProvider>
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
    </QueryProvider>
  </StrictMode>,
);
