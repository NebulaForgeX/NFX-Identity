import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";

import "./index.css";
import "@/assets/themes/global.css";
import "@/assets/languages/i18n";

import { BootstrapProvider, ModalProvider, QueryProvider, ThemeProvider, LenisProvider } from "@/providers";  

import App from "./App.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryProvider>
      <ThemeProvider>
        <BrowserRouter>
          <LenisProvider>
            <ModalProvider>
              <BootstrapProvider>
                <App />
              </BootstrapProvider>
            </ModalProvider>
          </LenisProvider>
        </BrowserRouter>
      </ThemeProvider>
    </QueryProvider>
  </StrictMode>,
);
