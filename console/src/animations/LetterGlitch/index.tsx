import { useRef, useEffect, useMemo, memo, useCallback } from 'react';
import styles from './style.module.css';
import { useTheme } from '@/providers/ThemeProvider/useTheme';

interface LetterGlitchProps {
  glitchColors?: string[];
  glitchSpeed?: number;
  centerVignette?: boolean;
  outerVignette?: boolean;
  smooth?: boolean;
  characters?: string;
  className?: string;
  style?: React.CSSProperties;
}

const LetterGlitch = memo(({
  glitchColors = ['#2b4539', '#61dca3', '#61b3dc'],
  glitchSpeed = 500, // 增加默认速度，减少更新频率
  centerVignette = false,
  outerVignette = true,
  smooth = true,
  characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$&*()-_+=/[]{};:<>.,0123456789',
  className = '',
  style = {}
}: LetterGlitchProps) => {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);
  const animationRef = useRef<number | null>(null);
  const letters = useRef<
    {
      char: string;
      color: string;
      targetColor: string;
      colorProgress: number;
    }[]
  >([]);
  const grid = useRef({ columns: 0, rows: 0 });
  const context = useRef<CanvasRenderingContext2D | null>(null);
  const lastGlitchTime = useRef(Date.now());

  const lettersAndSymbols = useMemo(() => Array.from(characters), [characters]);
  
  // 缓存 hexToRgb 结果
  const rgbCache = useRef<Map<string, { r: number; g: number; b: number } | null>>(new Map());

  const fontSize = 16;
  const charWidth = 10;
  const charHeight = 20;

  const getRandomChar = useCallback(() => {
    return lettersAndSymbols[Math.floor(Math.random() * lettersAndSymbols.length)];
  }, [lettersAndSymbols]);

  const getRandomColor = useCallback(() => {
    return glitchColors[Math.floor(Math.random() * glitchColors.length)];
  }, [glitchColors]);

  const hexToRgb = useCallback((hex: string) => {
    // 检查缓存
    if (rgbCache.current.has(hex)) {
      return rgbCache.current.get(hex)!;
    }
    
    const shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;
    const normalizedHex = hex.replace(shorthandRegex, (_m, r, g, b) => {
      return r + r + g + g + b + b;
    });

    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(normalizedHex);
    const rgb = result
      ? {
          r: parseInt(result[1], 16),
          g: parseInt(result[2], 16),
          b: parseInt(result[3], 16)
        }
      : null;
    
    // 缓存结果
    rgbCache.current.set(hex, rgb);
    return rgb;
  }, []);

  const interpolateColor = (
    start: { r: number; g: number; b: number },
    end: { r: number; g: number; b: number },
    factor: number
  ) => {
    const result = {
      r: Math.round(start.r + (end.r - start.r) * factor),
      g: Math.round(start.g + (end.g - start.g) * factor),
      b: Math.round(start.b + (end.b - start.b) * factor)
    };
    return `rgb(${result.r}, ${result.g}, ${result.b})`;
  };

  const calculateGrid = (width: number, height: number) => {
    const columns = Math.ceil(width / charWidth);
    const rows = Math.ceil(height / charHeight);
    return { columns, rows };
  };

  const initializeLetters = (columns: number, rows: number) => {
    grid.current = { columns, rows };
    const totalLetters = columns * rows;
    letters.current = Array.from({ length: totalLetters }, () => ({
      char: getRandomChar(),
      color: getRandomColor(),
      targetColor: getRandomColor(),
      colorProgress: 1
    }));
  };

  const resizeCanvas = () => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const parent = canvas.parentElement;
    if (!parent) return;

    const dpr = window.devicePixelRatio || 1;
    const rect = parent.getBoundingClientRect();

    canvas.width = rect.width * dpr;
    canvas.height = rect.height * dpr;

    canvas.style.width = `${rect.width}px`;
    canvas.style.height = `${rect.height}px`;

    if (context.current) {
      context.current.setTransform(dpr, 0, 0, dpr, 0, 0);
    }

    const { columns, rows } = calculateGrid(rect.width, rect.height);
    initializeLetters(columns, rows);
    drawLetters();
  };

  const drawLetters = useCallback(() => {
    if (!context.current || letters.current.length === 0) return;
    const ctx = context.current;
    const canvas = canvasRef.current;
    if (!canvas) return;
    
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.font = `${fontSize}px monospace`;
    ctx.textBaseline = 'top';
    ctx.textAlign = 'left';

    // 批量绘制，减少状态切换
    letters.current.forEach((letter, index) => {
      const x = (index % grid.current.columns) * charWidth;
      const y = Math.floor(index / grid.current.columns) * charHeight;
      ctx.fillStyle = letter.color;
      ctx.fillText(letter.char, x, y);
    });
  }, []);

  const updateLetters = useCallback(() => {
    if (!letters.current || letters.current.length === 0) return;

    // 减少更新比例，从 5% 降到 2%，降低性能消耗
    const updateCount = Math.max(1, Math.floor(letters.current.length * 0.02));

    for (let i = 0; i < updateCount; i++) {
      const index = Math.floor(Math.random() * letters.current.length);
      if (!letters.current[index]) continue;

      letters.current[index].char = getRandomChar();
      letters.current[index].targetColor = getRandomColor();

      if (!smooth) {
        letters.current[index].color = letters.current[index].targetColor;
        letters.current[index].colorProgress = 1;
      } else {
        letters.current[index].colorProgress = 0;
      }
    }
  }, [getRandomChar, getRandomColor, smooth]);

  const handleSmoothTransitions = useCallback(() => {
    let needsRedraw = false;
    let activeCount = 0;
    
    // 只处理正在过渡的字母，减少不必要的计算
    letters.current.forEach(letter => {
      if (letter.colorProgress < 1) {
        activeCount++;
        letter.colorProgress += 0.03; // 稍微降低过渡速度，减少重绘频率
        if (letter.colorProgress > 1) letter.colorProgress = 1;

        const startRgb = hexToRgb(letter.color);
        const endRgb = hexToRgb(letter.targetColor);
        if (startRgb && endRgb) {
          letter.color = interpolateColor(startRgb, endRgb, letter.colorProgress);
          needsRedraw = true;
        }
      }
    });

    // 只有在有活动过渡时才重绘
    if (needsRedraw && activeCount > 0) {
      drawLetters();
    }
  }, [hexToRgb, drawLetters]);

  const animate = useCallback(() => {
    const now = Date.now();
    const timeSinceLastGlitch = now - lastGlitchTime.current;
    
    // 只在需要更新时才执行
    if (timeSinceLastGlitch >= glitchSpeed) {
      updateLetters();
      drawLetters();
      lastGlitchTime.current = now;
    } else if (smooth) {
      // 只在平滑模式下才检查过渡
      handleSmoothTransitions();
    }

    animationRef.current = requestAnimationFrame(animate);
  }, [glitchSpeed, smooth, updateLetters, handleSmoothTransitions]);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    context.current = canvas.getContext('2d');
    resizeCanvas();
    animate();

    let resizeTimeout: ReturnType<typeof setTimeout>;

    const handleResize = () => {
      clearTimeout(resizeTimeout);
      resizeTimeout = setTimeout(() => {
        if (animationRef.current) {
          cancelAnimationFrame(animationRef.current);
        }
        resizeCanvas();
        animate();
      }, 100);
    };

    window.addEventListener('resize', handleResize);

    return () => {
      if (animationRef.current) {
        cancelAnimationFrame(animationRef.current);
      }
      window.removeEventListener('resize', handleResize);
      clearTimeout(resizeTimeout);
      // 清理缓存
      rgbCache.current.clear();
    };
  }, [glitchSpeed, smooth, glitchColors, animate]);

  return (
    <div className={`${styles.container} ${className}`} style={style}>
      <canvas ref={canvasRef} className={styles.canvas} />
      {outerVignette && <div className={styles.outerVignette}></div>}
      {centerVignette && <div className={styles.centerVignette}></div>}
    </div>
  );
});

const LetterGlitchBackground: React.FC<Omit<LetterGlitchProps, 'glitchColors'>> = memo((props) => {
  const { currentTheme } = useTheme();

  // 从主题中获取故障颜色数组
  const glitchColors = useMemo(() => {
    const primary = currentTheme.variables.primary || '#FA1E16';
    const info = currentTheme.variables.info || '#0095ff';
    const success = currentTheme.variables.success || '#00d68f';
    
    const toHex = (color: string): string => {
      if (color.startsWith('#')) {
        return color;
      }
      const match = color.match(/rgb\((\d+),\s*(\d+),\s*(\d+)\)/);
      if (match) {
        const r = parseInt(match[1], 10).toString(16).padStart(2, '0');
        const g = parseInt(match[2], 10).toString(16).padStart(2, '0');
        const b = parseInt(match[3], 10).toString(16).padStart(2, '0');
        return `#${r}${g}${b}`;
      }
      return color;
    };
    
    return [toHex(primary), toHex(info), toHex(success)];
  }, [currentTheme]);

  return (
    <LetterGlitch
      glitchColors={glitchColors}
      glitchSpeed={300} // 增加默认速度，减少更新频率
      centerVignette={true}
      outerVignette={false}
      smooth={true}
      {...props}
    />
  );
});

LetterGlitchBackground.displayName = 'LetterGlitchBackground';

export default LetterGlitchBackground;
