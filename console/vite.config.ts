import path from "path";
import react from "@vitejs/plugin-react";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig, loadEnv } from "vite";

const PAGE_CHUNKS: Record<string, string> = {
  "/src/pages/LoginPage": "page-login",
  "/src/pages/DashboardPage": "page-dashboard",
  "/src/pages/NotFoundPage": "page-404",
  "/src/pages/ProfilePage": "page-profile",
  "/src/pages/ProfileEditPage": "page-profile-edit",
  "/src/pages/AccountSecurityPage": "page-account-security",
  "/src/pages/ViewProfilePage": "page-profile-view",
  "/src/pages/CategoryListPage": "page-category-list",
  "/src/pages/CategoryAddPage": "page-category-add",
  "/src/pages/CategoryDetailPage": "page-category-detail",
  "/src/pages/CategoryEditPage": "page-category-edit",
  "/src/pages/CategoryPanelPage": "page-category-panel",
  "/src/pages/SubcategoryListPage": "page-subcategory-list",
  "/src/pages/SubcategoryAddPage": "page-subcategory-add",
  "/src/pages/SubcategoryDetailPage": "page-subcategory-detail",
  "/src/pages/SubcategoryEditPage": "page-subcategory-edit",
  "/src/pages/TeaListPage": "page-tea-list",
  "/src/pages/TeaAddPage": "page-tea-add",
  "/src/pages/TeaDetailPage": "page-tea-detail",
  "/src/pages/TeaEditPage": "page-tea-edit",
};

const TEA_COMPONENT_CHUNKS: Record<string, string> = {
  "/src/elements/tea/components/TeaListItem": "elements-tea-list",
  "/src/elements/tea/components/TeaImageUploader": "elements-tea-media",
  "/src/elements/tea/components/ImagesController": "elements-tea-media",
  "/src/elements/tea/components/CategorySelector": "elements-tea-category-selector",
  "/src/elements/tea/components/CategoryController": "elements-tea-form",
  "/src/elements/tea/components/NameController": "elements-tea-form",
  "/src/elements/tea/components/DescriptionController": "elements-tea-form",
  "/src/elements/tea/components/PriceController": "elements-tea-form",
  "/src/elements/tea/components/StockController": "elements-tea-form",
  "/src/elements/tea/components/ShowController": "elements-tea-form",
};

const ELEMENT_CHUNKS: Record<string, string> = {
  "/src/elements/profile": "elements-profile",
  "/src/elements/category/components": "elements-category-components",
  "/src/elements/category/controllers": "elements-category-controllers",
  "/src/elements/category/hooks": "elements-category-hooks",
  "/src/elements/subCategory/components": "elements-subcategory-components",
  "/src/elements/subCategory/controllers": "elements-subcategory-controllers",
  "/src/elements/subCategory/hooks": "elements-subcategory-hooks",
  "/src/elements/tea/controllers": "elements-tea-controllers",
  "/src/elements/tea/hooks": "elements-tea-hooks",
  ...TEA_COMPONENT_CHUNKS,
};

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "");
  const port = Number(env.VITE_PORT) || 5173;

  return {
    plugins: [
      react(),
      visualizer({
        filename: "./dist/stats.html", // 分析图生成的文件名
        open: true, // 构建完成后自动打开浏览器
        gzipSize: true, // 显示 gzip 后的大小
        brotliSize: true, // 显示 brotli 压缩后的大小
      }),
    ],
    base: "/", // 添加相对路径base，确保资源路径正确
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
        "lucide-react/icons": path.resolve(__dirname, "./node_modules/lucide-react/dist/esm/icons"),
      },
    },
    css: {
      modules: {
        localsConvention: "camelCase",
        generateScopedName: "[name]__[local]___[hash:base64:5]",
      },
    },
    server: {
      port: port,
      host: "0.0.0.0", // 允许局域网访问
      open: true,
      watch: {
        // 排除 templates 目录，避免监听其他项目的文件
        ignored: ["**/templates/**", "**/node_modules/**"],
      },
    },
    build: {
      outDir: "dist",
      sourcemap: true,
      chunkSizeWarningLimit: 300, // 调整警告阈值为 300 kB (gzip 后的大小更重要)
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes("node_modules")) {
              if (id.includes("react") || id.includes("react-dom") || id.includes("scheduler")) {
                return "react-vendor";
              }
              if (id.includes("react-router-dom") || id.includes("react-router")) {
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
    optimizeDeps: {
      // 排除 templates 目录的文件
      exclude: ["templates"],
    },
    preview: {
      port: port,
      host: "0.0.0.0",
    },
  };
});
