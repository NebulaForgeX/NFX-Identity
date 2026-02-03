import type { Theme } from "../types";

/**
 * Dark 主题 - 重新设计
 * - 背景层级分明（base → surface → elevated）
 * - 边框可见且不刺眼
 * - 主色为蓝色强调，语义色适配深色背景
 * - 文字对比度符合可读性
 */
export const darkTheme: Theme = {
  name: "dark",
  displayName: "Dark",
  variables: {
    // 主色 - 金黄色/琥珀色
    primary: "#d4a017",
    primaryLight: "#eab308",
    primaryFg: "#1a1a1a",

    // 语义色 - 深色模式、无蓝色，info 用青绿区分
    success: "#22c55e",
    successLight: "#4ade80",
    info: "#14b8a6",
    infoLight: "#2dd4bf",
    warning: "#eab308",
    warningLight: "#facc15",
    danger: "#ef4444",
    dangerLight: "#f87171",

    // 背景色 - 明确层级（由深到浅）
    bg: "#0f0f11",
    bg2: "#18181b",
    bg3: "#27272a",
    bg4: "#3f3f46",

    // 边框色 - 与背景区分，便于识别边界
    border: "#18181b",
    border2: "#27272a",
    border3: "#3f3f46",
    border4: "#52525b",
    border5: "#71717a",

    // 文字色 - 主文白、标题更亮、辅助文灰
    fg: "#a1a1aa",
    fgText: "#fafafa",
    fgHeading: "#ffffff",
    fgHighlight: "#fde047",

    // 分隔线
    separator: "#27272a",

    // 扩展配置 - Temperature
    temperature: {
      arcFill: ["#d4a017", "#eab308", "#fcd34d", "#fde047", "#fde047"],
      arcEmpty: "#27272a",
      thumbBg: "#27272a",
      thumbBorder: "#52525b",
    },

    // Solar
    solar: {
      gradientLeft: "#d4a017",
      gradientRight: "#eab308",
      shadowColor: "rgba(0, 0, 0, 0.3)",
      secondSeriesFill: "#27272a",
      radius: ["80%", "90%"],
    },

    // Traffic
    traffic: {
      tooltipBg: "#27272a",
      tooltipBorderColor: "#3f3f46",
      tooltipExtraCss: "border-radius: 10px; padding: 4px 16px;",
      tooltipTextColor: "#fafafa",
      tooltipFontWeight: "normal",
      yAxisSplitLine: "#27272a",
      lineBg: "#3f3f46",
      lineShadowBlur: "1",
      itemColor: "#3f3f46",
      itemBorderColor: "#52525b",
      itemEmphasisBorderColor: "#eab308",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
      shadowLineShadow: "rgba(0, 0, 0, 0)",
      gradFrom: "#27272a",
      gradTo: "#18181b",
    },

    // Electricity
    electricity: {
      tooltipBg: "#27272a",
      tooltipLineColor: "#fafafa",
      tooltipLineWidth: "0",
      tooltipBorderColor: "#3f3f46",
      tooltipExtraCss: "border-radius: 10px; padding: 8px 24px;",
      tooltipTextColor: "#fafafa",
      tooltipFontWeight: "normal",
      axisLineColor: "#3f3f46",
      xAxisTextColor: "#a1a1aa",
      yAxisSplitLine: "#27272a",
      itemBorderColor: "#eab308",
      lineStyle: "solid",
      lineWidth: "4",
      lineGradFrom: "#d4a017",
      lineGradTo: "#eab308",
      lineShadow: "rgba(212, 160, 23, 0.35)",
      areaGradFrom: "#27272a",
      areaGradTo: "#18181b",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
    },

    // ECharts
    echarts: {
      bg: "#18181b",
      textColor: "#fafafa",
      axisLineColor: "#52525b",
      splitLineColor: "#27272a",
      itemHoverShadowColor: "rgba(0, 0, 0, 0.4)",
      tooltipBackgroundColor: "#27272a",
      areaOpacity: "0.6",
    },

    // ChartJS
    chartjs: {
      axisLineColor: "#3f3f46",
      textColor: "#fafafa",
    },
  },
};
