import { Flex, Slider as RadixSlider, Text } from "@radix-ui/themes";

export type SliderProps = {
  value?: number;
  onChange?: (value: number) => void;
  min?: number;
  max?: number;
  step?: number;
  showValue?: boolean;
  className?: string;
};

export default function Slider({ value = 0, onChange, min = 0, max = 10, step = 1, showValue, className }: SliderProps) {
  return (
    <Flex align="center" gap="2" width="100%" className={className}>
      <RadixSlider
        value={[value]}
        min={min}
        max={max}
        step={step}
        onValueChange={(next) => onChange?.(next[0] ?? min)}
      />
      {showValue ? (
        <Text size="1" color="gray">
          {value}
        </Text>
      ) : null}
    </Flex>
  );
}
