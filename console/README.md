# NFX Console

NFX Console 是一个现代化的后台管理系统控制台，基于 React 19 + TypeScript + Vite 构建。

## 📋 项目简介

NFX Console 是一个功能完善的后台管理系统，用于管理系统用户、认证、图片以及相关业务功能。系统采用最新的前端技术栈，提供流畅的用户体验和高效的开发体验。

## ✨ 主要功能

### 🔐 用户认证
- 用户登录/登出
- 身份验证与授权
- 会话管理

### 📊 仪表板
- 数据概览
- 系统统计信息

### 👤 个人中心
- 个人资料查看与编辑
- 账户安全设置
- 密码修改

### 📦 分类管理
- 分类列表查看
- 分类添加/编辑/删除
- 分类详情查看
- 分类面板管理

### 🏷️ 子分类管理
- 子分类列表查看
- 子分类添加/编辑/删除
- 子分类详情查看

### 🍵 茶叶管理
- 用户与认证管理
- 图片资源管理
- 系统配置管理

## 🛠️ 技术栈

### 核心框架
- **React 19.1.1** - 最新的 React 框架
- **TypeScript 5.9.3** - 类型安全的 JavaScript
- **Vite (Rolldown)** - 极速构建工具

### 路由与状态管理
- **React Router DOM 7.9.5** - 声明式路由
- **Zustand 5.0.8** - 轻量级状态管理
- **TanStack Query 5.90.6** - 强大的数据获取与缓存

### UI 组件库
- **Headless UI 2.2.9** - 无样式 UI 组件
- **Lucide React 0.552.0** - 现代化图标库
- **React Pro Sidebar 1.1.0** - 专业侧边栏组件

### 表单处理
- **React Hook Form 7.66.0** - 高性能表单库
- **Zod 4.1.12** - TypeScript 优先的 schema 验证
- **@hookform/resolvers 5.2.2** - 表单验证解析器

### 国际化
- **i18next 25.6.0** - 国际化框架
- **react-i18next 16.2.3** - React 国际化绑定
- **i18next-browser-languagedetector 8.2.0** - 语言检测

### HTTP 客户端
- **Axios 1.13.1** - HTTP 请求库
- **axios-case-converter 1.1.1** - 请求/响应数据格式转换

### 开发工具
- **ESLint 9.36.0** - 代码检查
- **Prettier 3.4.2** - 代码格式化
- **Husky 9.1.6** - Git hooks 管理
- **TypeScript ESLint 8.45.0** - TypeScript 代码检查

## 📁 项目结构

```
nfx-console/
├── src/
│   ├── apis/              # API 接口定义
│   ├── assets/            # 静态资源
│   ├── components/        # 共享组件
│   ├── elements/          # 业务元素组件
│   │   └── profile/        # 个人资料相关组件
│   ├── events/            # 事件系统
│   ├── hooks/             # 自定义 Hooks
│   ├── layouts/           # 布局组件
│   ├── pages/             # 页面组件
│   ├── providers/         # Context Providers
│   ├── services/          # 业务服务
│   ├── stores/            # 状态管理
│   ├── types/             # TypeScript 类型定义
│   └── utils/             # 工具函数
├── public/                # 公共静态文件
├── dist/                  # 构建输出目录
├── vite.config.ts         # Vite 配置文件
├── tsconfig.json          # TypeScript 配置
├── eslint.config.js       # ESLint 配置
└── package.json           # 项目依赖配置
```

## 🚀 快速开始

### 环境要求

- **Node.js**: >= 18.0.0
- **npm**: >= 9.0.0 或 **pnpm**: >= 8.0.0

### 安装依赖

```bash
npm install
# 或
pnpm install
```

### 开发模式

```bash
# 开发环境
npm run dev

# 生产环境预览
npm run dev:prod
```

开发服务器将在 `http://localhost:5173` 启动，支持热模块替换（HMR）。

### 构建项目

```bash
# 生产环境构建
npm run build

# 开发环境构建
npm run build:dev
```

构建产物将输出到 `dist/` 目录。

### 预览构建结果

```bash
npm run preview
```

## 📜 可用脚本

| 脚本 | 描述 |
|------|------|
| `npm run dev` | 启动开发服务器（开发环境模式） |
| `npm run dev:prod` | 启动开发服务器（生产环境模式） |
| `npm run build` | 构建生产版本 |
| `npm run build:dev` | 构建开发版本 |
| `npm run preview` | 预览构建结果 |
| `npm run lint` | 运行 ESLint 代码检查 |
| `npm run lint:code-format` | 检查代码格式（Prettier） |

## 🎯 核心特性

### 代码分割优化
项目采用精细化的代码分割策略，将代码拆分为多个 chunk：
- **Vendor chunks**: React、Router、i18n、Form、Query、Icons、Utils
- **Page chunks**: 按页面拆分，实现按需加载
- **Element chunks**: 按业务模块拆分组件
- **Shared chunks**: 共享的 hooks、APIs、stores

### 类型安全
- 完整的 TypeScript 类型定义
- Zod schema 验证确保运行时类型安全
- 严格的 ESLint TypeScript 规则

### 国际化支持
- 支持中英文切换
- 自动语言检测
- 完整的翻译体系

### 性能优化
- 虚拟滚动支持（VirtualList）
- 代码分割与懒加载
- React Query 智能缓存
- 构建产物分析（rollup-plugin-visualizer）

### 开发体验
- 热模块替换（HMR）
- 类型检查与代码提示
- 自动代码格式化
- Git hooks 自动检查

## 🔧 配置说明

### 环境变量

项目支持多环境配置，通过 `--mode` 参数指定：
- `dev`: 开发环境
- `prod`: 生产环境

### Vite 配置

主要配置项：
- **端口**: 5173
- **主机**: 0.0.0.0（支持局域网访问）
- **别名**: `@` 指向 `src/` 目录
- **CSS Modules**: 支持 camelCase 命名
- **Source Map**: 生产构建包含 source map

### TypeScript 配置

- `tsconfig.json`: 根配置
- `tsconfig.app.json`: 应用代码配置
- `tsconfig.node.json`: Node.js 工具配置

## 📝 代码规范

### ESLint
项目使用 ESLint 9.x 扁平配置，包含：
- React Hooks 规则
- TypeScript 推荐规则
- React 最佳实践

### Prettier
统一的代码格式化配置，支持：
- 自动导入排序
- 代码格式化
- 文件类型过滤

### Git Hooks
通过 Husky 配置 Git hooks，确保：
- 提交前代码检查
- 代码格式验证

## 🌐 浏览器支持

- Chrome (最新版)
- Firefox (最新版)
- Safari (最新版)
- Edge (最新版)

## 📄 许可证

本项目为私有项目，仅供内部使用。

## 👥 贡献

本项目为内部项目，如有问题或建议，请联系开发团队。

## 📞 联系方式

---

**NFX Console** - 专业的后台管理系统控制台
