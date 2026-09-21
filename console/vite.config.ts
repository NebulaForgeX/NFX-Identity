import path from "node:path";
import { fileURLToPath } from "node:url";
import react from "@vitejs/plugin-react";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig } from "vite";

import {
  loadNfxConsoleEnv,
  nfxKillListenPortPlugin,
  nfxUiAtAliasPlugin,
  nfxUiDedupe,
  nfxUiOptimizeDepsExclude,
  nfxUiViteAliases,
  nfxViteDefine,
  nfxConsoleBase,
  nfxViteDevServer,
  resolveNfxUiRoot,
} from "./vite.nfx-ui.ts";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname);
const nfxUiRoot = resolveNfxUiRoot(root);

const PAGE_CHUNKS: Record<string, string> = {
  "/src/pages/Auth/Login": "page-login",
  "/src/pages/Dashboard": "page-dashboard",
  "/src/pages/NotFound": "page-404",
  "/src/pages/User/Profile": "page-profile",
  "/src/pages/Images": "page-images",
  "/src/pages/OwnerDirectory": "page-owner",
  "/src/pages/User/Settings": "page-settings",
};

const ELEMENT_CHUNKS: Record<string, string> = {
  "/src/elements/profile": "elements-profile",
};

export default defineConfig(({ mode, command }) => {
  const env = loadNfxConsoleEnv(root, mode);
  const port = Number(env.VITE_PORT) || 5173;
  const hasApiUrl = Boolean(env.VITE_API_URL);
  const proxyTarget = env.VITE_DEV_API_PROXY_TARGET || env.VITE_API_URL || "http://192.168.1.64/nfx-identity";

  return {
    base: nfxConsoleBase(env),
    define: nfxViteDefine(env),
    plugins: [
      nfxKillListenPortPlugin(port),
      nfxUiAtAliasPlugin(root, nfxUiRoot),
      react(),
      visualizer({
        filename: "./dist/stats.html",
        open: process.env.DOCKER !== "1" && process.env.DOCKER_BUILD !== "1",
        gzipSize: true,
        brotliSize: true,
      }),
    ],
    resolve: {
      alias: nfxUiViteAliases(root, nfxUiRoot),
      dedupe: nfxUiDedupe,
    },
    css: {
      modules: {
        localsConvention: "camelCase",
        generateScopedName: "[name]__[local]___[hash:base64:5]",
      },
    },
    optimizeDeps: {
      exclude: nfxUiOptimizeDepsExclude,
      holdUntilCrawlEnd: false,
    },
    server: {
      ...nfxViteDevServer(env, port),
      fs: { allow: [root, nfxUiRoot] },
      watch: {
        ignored: ["**/templates/**"],
      },
      ...(command === "serve" && !hasApiUrl
        ? {
            proxy: {
              "/auth": { target: proxyTarget, changeOrigin: true },
              "/asset": { target: proxyTarget, changeOrigin: true },
            },
          }
        : {}),
    },
    build: {
      outDir: "dist",
      sourcemap: true,
      chunkSizeWarningLimit: 300,
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes("node_modules")) {
              if (id.includes("react") || id.includes("react-dom") || id.includes("scheduler")) {
                return "react-vendor";
              }
              if (id.includes("react-router")) {
                return "router-vendor";
              }
              if (id.includes("i18next") || id.includes("react-i18next")) {
                return "i18n-vendor";
              }
              if (id.includes("react-hook-form") || id.includes("@hookform") || id.includes("zod")) {
                return "form-vendor";
              }
              if (id.includes("@tanstack/react-query") || id.includes("@tanstack/query-core")) {
                return "query-vendor";
              }
              if (id.includes("lucide-react")) {
                return "icons-vendor";
              }
              if (id.includes("axios") || id.includes("buffer") || id.includes("zustand")) {
                return "utils-vendor";
              }
            }

            for (const [pagePath, chunkName] of Object.entries(PAGE_CHUNKS)) {
              if (id.includes(pagePath)) {
                return chunkName;
              }
            }

            if (id.includes("/src/components/VirtualList")) {
              return "shared-virtual-list";
            }

            if (
              id.includes("/src/components/Header") ||
              id.includes("/src/components/Footer") ||
              id.includes("/src/components/Sidebar") ||
              id.includes("/src/layouts/")
            ) {
              return "shared-layout";
            }

            for (const [elementPath, chunkName] of Object.entries(ELEMENT_CHUNKS)) {
              if (id.includes(elementPath)) {
                return chunkName;
              }
            }

            if (id.includes("/src/hooks/")) {
              return "shared-hooks";
            }

            if (id.includes("/src/apis/")) {
              return "shared-apis";
            }

            if (id.includes("/src/stores/")) {
              return "shared-stores";
            }
          },
        },
      },
    },
    preview: {
      port,
      strictPort: true,
      host: "0.0.0.0",
    },
  };
});
