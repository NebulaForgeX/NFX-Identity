import type { Theme } from "../types";

/**
 * Coffee 主题 - 基于 8 色盘重新设计
 * 色盘：奶油 #F8F4E1 | 浅灰米 #E4E0E1 | 浅棕 #D6C0B3 | 中棕 #AB886D / #AF8F6F
 *       深棕 #74512D | 暗棕 #493628 | 最深 #543310
 * 深棕背景 + 奶油/米色文字，语义色保持棕/绿棕/红棕调
 */
export const coffeeTheme: Theme = {
  name: "coffee",
  displayName: "Coffee",
  variables: {
    // 主色 - 奶油/拿铁系
    primary: "#AF8F6F",
    primaryLight: "#F8F4E1",
    primaryFg: "#2a1f18",

    // 语义色 - 棕调配色（微调以协调）
    success: "#6d8f72",
    successLight: "#8aaa8e",
    info: "#8a7a6a",
    infoLight: "#a89888",
    warning: "#b8926a",
    warningLight: "#d4b896",
    danger: "#a65d5d",
    dangerLight: "#c07d7d",

    // 背景 - 由深到浅：最深 → #543310 → #493628 → #74512D
    bg: "#2e1a08",
    bg2: "#543310",
    bg3: "#493628",
    bg4: "#74512D",

    // 边框 - 深棕到浅米
    border: "#493628",
    border2: "#543310",
    border3: "#74512D",
    border4: "#AF8F6F",
    border5: "#D6C0B3",

    // 文字
    fg: "#AF8F6F",
    fgText: "#F8F4E1",
    fgHeading: "#F8F4E1",
    fgHighlight: "#E4E0E1",

    separator: "#493628",

    temperature: {
      arcFill: ["#AF8F6F", "#AB886D", "#D6C0B3", "#E4E0E1", "#F8F4E1"],
      arcEmpty: "#543310",
      thumbBg: "#493628",
      thumbBorder: "#AF8F6F",
    },

    solar: {
      gradientLeft: "#AF8F6F",
      gradientRight: "#F8F4E1",
      shadowColor: "rgba(0, 0, 0, 0.35)",
      secondSeriesFill: "#493628",
      radius: ["80%", "90%"],
    },

    traffic: {
      tooltipBg: "#74512D",
      tooltipBorderColor: "#AF8F6F",
      tooltipExtraCss: "border-radius: 10px; padding: 4px 16px;",
      tooltipTextColor: "#F8F4E1",
      tooltipFontWeight: "normal",
      yAxisSplitLine: "#543310",
      lineBg: "#493628",
      lineShadowBlur: "1",
      itemColor: "#543310",
      itemBorderColor: "#AF8F6F",
      itemEmphasisBorderColor: "#F8F4E1",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
      shadowLineShadow: "rgba(0, 0, 0, 0)",
      gradFrom: "#493628",
      gradTo: "#2e1a08",
    },

    electricity: {
      tooltipBg: "#543310",
      tooltipLineColor: "#F8F4E1",
      tooltipLineWidth: "0",
      tooltipBorderColor: "#AF8F6F",
      tooltipExtraCss: "border-radius: 10px; padding: 8px 24px;",
      tooltipTextColor: "#F8F4E1",
      tooltipFontWeight: "normal",
      axisLineColor: "#AF8F6F",
      xAxisTextColor: "#AF8F6F",
      yAxisSplitLine: "#543310",
      itemBorderColor: "#AF8F6F",
      lineStyle: "solid",
      lineWidth: "4",
      lineGradFrom: "#AF8F6F",
      lineGradTo: "#F8F4E1",
      lineShadow: "rgba(175, 143, 111, 0.3)",
      areaGradFrom: "#493628",
      areaGradTo: "#2e1a08",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
    },

    echarts: {
      bg: "#2e1a08",
      textColor: "#F8F4E1",
      axisLineColor: "#AF8F6F",
      splitLineColor: "#493628",
      itemHoverShadowColor: "rgba(0, 0, 0, 0.4)",
      tooltipBackgroundColor: "#543310",
      areaOpacity: "0.6",
    },

    chartjs: {
      axisLineColor: "#AF8F6F",
      textColor: "#F8F4E1",
    },
  },
};
