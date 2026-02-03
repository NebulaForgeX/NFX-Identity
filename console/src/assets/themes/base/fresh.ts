import type { Theme } from "../types";

/**
 * Fresh 主题 - 围绕 5D866C / C2A68C / E6D8C3 / F5F5F0 设计，非直接照搬
 * - 主色：青绿 #5D866C（按钮、链接、强调）
 * - 页面底：偏白 #FAFAF8，不用最浅色做大底再叠米色，避免「咖啡色块压在象牙上」的违和感
 * - 层级背景：从 F5F5F0 微暖白 → 极浅灰，卡片/输入框用浅中性
 * - 米/咖 E6D8C3、C2A68C 仅作点缀：边框、分隔线、图表辅助，不占大块面积
 */
export const freshTheme: Theme = {
  name: "fresh",
  displayName: "Fresh",
  variables: {
    // 主色 - 青绿（来自 5D866C）
    primary: "#5D866C",
    primaryLight: "#7a9d7e",
    primaryFg: "#ffffff",

    // 语义色 - 与主色协调，不抢戏
    success: "#5D866C",
    successLight: "#8ab88f",
    info: "#6b8f7a",
    infoLight: "#8fa895",
    warning: "#c9a227",
    warningLight: "#ddb84d",
    danger: "#b85c4a",
    dangerLight: "#d47868",

    // 背景 - 干净浅底 + 微暖层级，不大块用米/咖
    bg: "#FAFAF8",
    bg2: "#F5F5F0",
    bg3: "#EFEDE9",
    bg4: "#E8E6E1",

    // 边框 - 浅层用中性，深层用一点暖色点缀（E6D8C3 / C2A68C）
    border: "#FAFAF8",
    border2: "#F5F5F0",
    border3: "#E6D8C3",
    border4: "#ddceb8",
    border5: "#C2A68C",

    // 文字 - 深绿灰，在浅底上可读
    fg: "#5a7260",
    fgText: "#2d3b2d",
    fgHeading: "#1e2a1e",
    fgHighlight: "#4a6b52",

    // 分隔线 - 用米色做细线点缀
    separator: "#E6D8C3",

    // 扩展配置 - 图表主色用绿，辅助用浅色
    temperature: {
      arcFill: ["#5D866C", "#7a9d7e", "#8ab88f", "#9bc49e", "#9bc49e"],
      arcEmpty: "#EFEDE9",
      thumbBg: "#F5F5F0",
      thumbBorder: "#5D866C",
    },

    solar: {
      gradientLeft: "#5D866C",
      gradientRight: "#7a9d7e",
      shadowColor: "rgba(93, 134, 108, 0.15)",
      secondSeriesFill: "#F5F5F0",
      radius: ["80%", "90%"],
    },

    traffic: {
      tooltipBg: "#FAFAF8",
      tooltipBorderColor: "#E6D8C3",
      tooltipExtraCss: "border-radius: 10px; padding: 4px 16px;",
      tooltipTextColor: "#2d3b2d",
      tooltipFontWeight: "normal",
      yAxisSplitLine: "#E6D8C3",
      lineBg: "#EFEDE9",
      lineShadowBlur: "1",
      itemColor: "#EFEDE9",
      itemBorderColor: "#ddceb8",
      itemEmphasisBorderColor: "#5D866C",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
      shadowLineShadow: "rgba(0, 0, 0, 0)",
      gradFrom: "#F5F5F0",
      gradTo: "#FAFAF8",
    },

    electricity: {
      tooltipBg: "#FAFAF8",
      tooltipLineColor: "#2d3b2d",
      tooltipLineWidth: "0",
      tooltipBorderColor: "#E6D8C3",
      tooltipExtraCss: "border-radius: 10px; padding: 8px 24px;",
      tooltipTextColor: "#2d3b2d",
      tooltipFontWeight: "normal",
      axisLineColor: "#E6D8C3",
      xAxisTextColor: "#5a7260",
      yAxisSplitLine: "#E6D8C3",
      itemBorderColor: "#5D866C",
      lineStyle: "solid",
      lineWidth: "4",
      lineGradFrom: "#5D866C",
      lineGradTo: "#7a9d7e",
      lineShadow: "rgba(93, 134, 108, 0.2)",
      areaGradFrom: "#F5F5F0",
      areaGradTo: "#FAFAF8",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
    },

    echarts: {
      bg: "#FAFAF8",
      textColor: "#2d3b2d",
      axisLineColor: "#5D866C",
      splitLineColor: "#E6D8C3",
      itemHoverShadowColor: "rgba(93, 134, 108, 0.25)",
      tooltipBackgroundColor: "#F5F5F0",
      areaOpacity: "0.7",
    },

    chartjs: {
      axisLineColor: "#E6D8C3",
      textColor: "#2d3b2d",
    },
  },
};
