import type { Theme } from "../types";

/**
 * Wine 主题（深色）- 8 色盘（Burgundy 体系）
 * #2A0B12 Background | #3A0F1A Surface | #4A1322 Secondary | #800020 Primary
 * #9A1F3D Primary Hover | #B23A5C Accent | #C97A8A Muted Highlight | #F4E6E9 Text
 */
export const wineTheme: Theme = {
  name: "wine",
  displayName: "Wine",
  variables: {
    primary: "#800020",
    primaryLight: "#F4E6E9",
    primaryFg: "#F4E6E9",

    success: "#6b8f6b",
    successLight: "#8aaa8e",
    info: "#8a6a6a",
    infoLight: "#a88a8a",
    warning: "#c49050",
    warningLight: "#d4b080",
    danger: "#b85050",
    dangerLight: "#d07070",

    bg: "#2A0B12",
    bg2: "#3A0F1A",
    bg3: "#4A1322",
    bg4: "#4A1322",

    border: "#2A0B12",
    border2: "#3A0F1A",
    border3: "#4A1322",
    border4: "#800020",
    border5: "#B23A5C",

    fg: "#C97A8A",
    fgText: "#F4E6E9",
    fgHeading: "#F4E6E9",
    fgHighlight: "#F4E6E9",

    separator: "#3A0F1A",

    temperature: {
      arcFill: ["#800020", "#9A1F3D", "#B23A5C", "#C97A8A", "#F4E6E9"],
      arcEmpty: "#2A0B12",
      thumbBg: "#3A0F1A",
      thumbBorder: "#800020",
    },

    solar: {
      gradientLeft: "#800020",
      gradientRight: "#F4E6E9",
      shadowColor: "rgba(0, 0, 0, 0.4)",
      secondSeriesFill: "#3A0F1A",
      radius: ["80%", "90%"],
    },

    traffic: {
      tooltipBg: "#4A1322",
      tooltipBorderColor: "#B23A5C",
      tooltipExtraCss: "border-radius: 10px; padding: 4px 16px;",
      tooltipTextColor: "#F4E6E9",
      tooltipFontWeight: "normal",
      yAxisSplitLine: "#3A0F1A",
      lineBg: "#3A0F1A",
      lineShadowBlur: "1",
      itemColor: "#2A0B12",
      itemBorderColor: "#800020",
      itemEmphasisBorderColor: "#F4E6E9",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
      shadowLineShadow: "rgba(0, 0, 0, 0)",
      gradFrom: "#3A0F1A",
      gradTo: "#2A0B12",
    },

    electricity: {
      tooltipBg: "#3A0F1A",
      tooltipLineColor: "#F4E6E9",
      tooltipLineWidth: "0",
      tooltipBorderColor: "#800020",
      tooltipExtraCss: "border-radius: 10px; padding: 8px 24px;",
      tooltipTextColor: "#F4E6E9",
      tooltipFontWeight: "normal",
      axisLineColor: "#B23A5C",
      xAxisTextColor: "#C97A8A",
      yAxisSplitLine: "#3A0F1A",
      itemBorderColor: "#800020",
      lineStyle: "solid",
      lineWidth: "4",
      lineGradFrom: "#800020",
      lineGradTo: "#F4E6E9",
      lineShadow: "rgba(128, 0, 32, 0.35)",
      areaGradFrom: "#3A0F1A",
      areaGradTo: "#2A0B12",
      shadowLineDarkBg: "rgba(0, 0, 0, 0)",
    },

    echarts: {
      bg: "#2A0B12",
      textColor: "#F4E6E9",
      axisLineColor: "#B23A5C",
      splitLineColor: "#3A0F1A",
      itemHoverShadowColor: "rgba(0, 0, 0, 0.4)",
      tooltipBackgroundColor: "#3A0F1A",
      areaOpacity: "0.6",
    },

    chartjs: {
      axisLineColor: "#B23A5C",
      textColor: "#F4E6E9",
    },
  },
};
