import { forwardRef, memo, useEffect, useRef, useState } from "react";
import { animate, motion, useMotionValue, useMotionValueEvent, useTransform } from "motion/react";
import { ChevronLeft, ChevronRight } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

const MAX_OVERFLOW = 50;

export interface SliderProps extends Omit<React.HTMLAttributes<HTMLDivElement>, "onChange"> {
  value: number;
  onChange: (value: number) => void;
  min?: number;
  max?: number;
  step?: number;
  showValue?: boolean;
  fullWidth?: boolean;
  label?: string;
  error?: string;
  helperText?: string;
  required?: boolean;
}

const Slider = memo(
  forwardRef<HTMLDivElement, SliderProps>(
    (
      {
        value: controlledValue,
        onChange,
        min = 0,
        max = 10,
        step = 1,
        showValue = true,
        fullWidth = false,
        label,
        error,
        helperText,
        required = false,
        className = "",
        ...props
      },
      ref,
    ) => {
      const [value, setValue] = useState<number>(controlledValue);
      const sliderRef = useRef<HTMLDivElement>(null);
      const [region, setRegion] = useState<"left" | "middle" | "right">("middle");
      const clientX = useMotionValue(0);
      const overflow = useMotionValue(0);

      useEffect(() => {
        setValue(controlledValue);
      }, [controlledValue]);

      useMotionValueEvent(clientX, "change", (latest: number) => {
        if (sliderRef.current) {
          const { left, right } = sliderRef.current.getBoundingClientRect();
          let newValue: number;
          if (latest < left) {
            setRegion("left");
            newValue = left - latest;
          } else if (latest > right) {
            setRegion("right");
            newValue = latest - right;
          } else {
            setRegion("middle");
            newValue = 0;
          }
          overflow.jump(decay(newValue, MAX_OVERFLOW));
        }
      });

      const handlePointerMove = (e: React.PointerEvent<HTMLDivElement>) => {
        if (e.buttons > 0 && sliderRef.current) {
          const { left, width } = sliderRef.current.getBoundingClientRect();
          let newValue = min + ((e.clientX - left) / width) * (max - min);
          if (step > 0) {
            newValue = Math.round(newValue / step) * step;
          }
          newValue = Math.min(Math.max(newValue, min), max);
          setValue(newValue);
          onChange(newValue);
          clientX.jump(e.clientX);
        }
      };

      const handlePointerDown = (e: React.PointerEvent<HTMLDivElement>) => {
        handlePointerMove(e);
        e.currentTarget.setPointerCapture(e.pointerId);
      };

      const handlePointerUp = () => {
        animate(overflow, 0, { type: "spring", bounce: 0.5 });
      };

      const getRangePercentage = (): number => {
        const totalRange = max - min;
        if (totalRange === 0) return 0;
        return ((value - min) / totalRange) * 100;
      };

      const wrapperClasses = [styles.sliderContainer, fullWidth && styles.fullWidth, className]
        .filter(Boolean)
        .join(" ");

      return (
        <div ref={ref} className={wrapperClasses} {...props}>
          {label && (
            <div className={styles.labelRow}>
              <label className={styles.label}>
                {label}
                {required && <span className={styles.required}>*</span>}
              </label>
            </div>
          )}
          <motion.div
            className={styles.sliderWrapper}
          >
            <motion.div
              animate={{
                scale: region === "left" ? [1, 1.4, 1] : 1,
                transition: { duration: 0.25 },
              }}
              style={{
                x: useTransform(() => (region === "left" ? -overflow.get() : 0)),
              }}
              className={styles.iconWrapper}
            >
              <ChevronLeft size={20} className={styles.icon} />
            </motion.div>

            <div
              ref={sliderRef}
              className={styles.sliderRoot}
              onPointerMove={handlePointerMove}
              onPointerDown={handlePointerDown}
              onPointerUp={handlePointerUp}
            >
              <motion.div
                style={{
                  scaleX: useTransform(() => {
                    if (sliderRef.current) {
                      const { width } = sliderRef.current.getBoundingClientRect();
                      return 1 + overflow.get() / width;
                    }
                    return 1;
                  }),
                  scaleY: useTransform(overflow, [0, MAX_OVERFLOW], [1, 0.8]),
                  transformOrigin: useTransform(() => {
                    if (sliderRef.current) {
                      const { left, width } = sliderRef.current.getBoundingClientRect();
                      return clientX.get() < left + width / 2 ? "right" : "left";
                    }
                    return "center";
                  }),
                }}
                className={styles.sliderTrackWrapper}
              >
                <div className={styles.sliderTrack}>
                  <motion.div
                    className={styles.sliderRange}
                    initial={false}
                    animate={{ width: `${getRangePercentage()}%` }}
                    transition={{ type: "spring", stiffness: 300, damping: 30 }}
                  />
                </div>
              </motion.div>
            </div>

            <motion.div
              animate={{
                scale: region === "right" ? [1, 1.4, 1] : 1,
                transition: { duration: 0.25 },
              }}
              style={{
                x: useTransform(() => (region === "right" ? overflow.get() : 0)),
              }}
              className={styles.iconWrapper}
            >
              <ChevronRight size={20} className={styles.icon} />
            </motion.div>
            {showValue && !label && (
              <span className={styles.valueIndicator}>{Math.round(value)}</span>
            )}
          </motion.div>
          {error && <p className={styles.errorMessage}>{error}</p>}
          {helperText && !error && <p className={styles.helperText}>{helperText}</p>}
        </div>
      );
    },
  ),
);

function decay(value: number, max: number): number {
  if (max === 0) {
    return 0;
  }
  const entry = value / max;
  const sigmoid = 2 * (1 / (1 + Math.exp(-entry)) - 0.5);
  return sigmoid * max;
}

Slider.displayName = "Slider";

export default Slider;
