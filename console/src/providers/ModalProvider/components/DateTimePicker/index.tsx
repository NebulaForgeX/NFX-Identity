import type { ReactDatePickerCustomHeaderProps } from "react-datepicker";

import { useEffect, useMemo, useState } from "react";
import { CheckIcon, Cross2Icon } from "@radix-ui/react-icons";
import { Box, Button, Dialog, Flex, Grid, IconButton, Select, Text } from "@radix-ui/themes";
import { format } from "date-fns";
import { enUS, fr, zhCN } from "date-fns/locale";
import { CalendarDays } from "lucide-react";
import { LanguageEnum } from "nfx-ui/enums";
import DatePicker, { registerLocale } from "react-datepicker";
import { useTranslation } from "react-i18next";

import LucideIcon from "@/components/LucideIcon";
import ModalStore, { hideModal, useModalStore } from "@/stores/modal";

import "react-datepicker/dist/react-datepicker.css";

import styles from "./s.module.css";

registerLocale(LanguageEnum.EN, enUS);
registerLocale(LanguageEnum.ZH, zhCN);
registerLocale(LanguageEnum.FR, fr);

const MONTH_INDEXES = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] as const;

function parseDateInput(value: string): Date | null {
  const match = /^(\d{4})-(\d{2})-(\d{2})$/.exec(value);
  if (!match) return null;
  const year = Number(match[1]);
  const month = Number(match[2]);
  const day = Number(match[3]);
  const date = new Date(year, month - 1, day);
  if (date.getFullYear() !== year || date.getMonth() !== month - 1 || date.getDate() !== day) return null;
  return date;
}

function formatDateInput(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  return `${year}-${month}-${day}`;
}

function startOfLocalDay(date: Date): Date {
  return new Date(date.getFullYear(), date.getMonth(), date.getDate());
}

function buildYearOptions(minDate: Date, maxDate: Date): number[] {
  const years: number[] = [];
  for (let year = minDate.getFullYear(); year <= maxDate.getFullYear(); year += 1) years.push(year);
  return years;
}

function clampDate(date: Date, minDate: Date, maxDate: Date): Date {
  const day = startOfLocalDay(date);
  if (day < minDate) return minDate;
  if (day > maxDate) return maxDate;
  return day;
}

const DateTimePicker = () => {
  const { t, i18n } = useTranslation("pages.User.Profile.Edit");
  const isOpen = useModalStore((state) => state.dateTimePickerModal.isOpen);
  const value = useModalStore((state) => state.dateTimePickerModal.value);
  const title = useModalStore((state) => state.dateTimePickerModal.title);
  const minDateProp = useModalStore((state) => state.dateTimePickerModal.minDate);
  const maxDateProp = useModalStore((state) => state.dateTimePickerModal.maxDate);
  const allowClear = useModalStore((state) => state.dateTimePickerModal.allowClear);

  const [selected, setSelected] = useState<Date | null>(null);
  const [calendarMonth, setCalendarMonth] = useState<Date>(() => new Date());
  const [pickerEpoch, setPickerEpoch] = useState(0);

  const locale = (() => {
    const lng = i18n.language?.toLowerCase();
    if (lng?.startsWith(LanguageEnum.ZH)) return LanguageEnum.ZH;
    if (lng?.startsWith(LanguageEnum.FR)) return LanguageEnum.FR;
    return LanguageEnum.EN;
  })();
  const dateFnsLocale = locale === LanguageEnum.ZH ? zhCN : locale === LanguageEnum.FR ? fr : enUS;

  const fallbackMax = useMemo(() => startOfLocalDay(new Date()), [isOpen]);
  const maxDate = maxDateProp ?? fallbackMax;
  const minDate = useMemo(() => minDateProp ?? new Date(maxDate.getFullYear() - 120, 0, 1), [minDateProp, maxDate]);
  const yearOptions = useMemo(() => buildYearOptions(minDate, maxDate), [minDate, maxDate]);

  useEffect(() => {
    if (!isOpen) return;
    const parsed = parseDateInput(value);
    const next = parsed ? clampDate(parsed, minDate, maxDate) : null;
    setSelected(next);
    setCalendarMonth(next ?? maxDate);
  }, [isOpen, value, minDate, maxDate]);

  const dismiss = () => {
    hideModal("dateTimePicker");
  };

  const handleOpenChange = (open: boolean) => {
    if (!open) dismiss();
  };

  const handleCancel = () => {
    ModalStore.getState().dateTimePickerModal.onCancel?.();
    dismiss();
  };

  const handleConfirm = () => {
    if (!selected) return;
    ModalStore.getState().dateTimePickerModal.onConfirm?.(formatDateInput(selected));
    dismiss();
  };

  const handleClear = () => {
    ModalStore.getState().dateTimePickerModal.onConfirm?.("");
    dismiss();
  };

  const handleToday = () => {
    const next = clampDate(new Date(), minDate, maxDate);
    setSelected(next);
    setCalendarMonth(next);
    setPickerEpoch((epoch) => epoch + 1);
  };

  const renderCustomHeader = ({ date, changeYear, changeMonth }: ReactDatePickerCustomHeaderProps) => (
    <Box className={styles.calendarHeader}>
      <Flex align="center" justify="center" gap="1" wrap="wrap">
        <Select.Root
          size="2"
          value={String(date.getMonth())}
          onValueChange={(next) => {
            const month = Number(next);
            changeMonth(month);
            setCalendarMonth(new Date(date.getFullYear(), month, 1));
          }}
        >
          <Select.Trigger variant="surface" className={styles.headerSelectTrigger} />
          <Select.Content position="popper">
            {MONTH_INDEXES.map((month) => (
              <Select.Item key={month} value={String(month)}>
                {format(new Date(2000, month, 1), "LLLL", { locale: dateFnsLocale })}
              </Select.Item>
            ))}
          </Select.Content>
        </Select.Root>
        <Select.Root
          size="2"
          value={String(date.getFullYear())}
          onValueChange={(next) => {
            const year = Number(next);
            changeYear(year);
            setCalendarMonth(new Date(year, date.getMonth(), 1));
          }}
        >
          <Select.Trigger variant="surface" className={styles.headerSelectTrigger} />
          <Select.Content position="popper">
            {yearOptions.map((year) => (
              <Select.Item key={year} value={String(year)}>
                {year}
              </Select.Item>
            ))}
          </Select.Content>
        </Select.Root>
      </Flex>
    </Box>
  );

  const summary = selected ? format(selected, "PPP", { locale: dateFnsLocale }) : t("datePicker.empty");

  return (
    <Dialog.Root open={isOpen} onOpenChange={handleOpenChange}>
      <Dialog.Content aria-describedby={undefined} maxWidth="min(720px, calc(100vw - var(--space-4) * 2))" style={{ padding: 0 }}>
        <Box py="4">
          <Box px="4">
            <Flex direction="column" gap="4">
              <Flex justify="between" align="start" gap="3">
                <Flex align="center" gap="3" minWidth="0">
                  <Box className={styles.headerSize}>
                    <Box className={`${styles.headerRadius} ${styles.iconInk}`}>
                      <Flex align="center" justify="center" className={`${styles.headerFill} ${styles.iconGlyph}`}>
                        <LucideIcon icon={CalendarDays} size={22} />
                      </Flex>
                    </Box>
                  </Box>
                  <Box>
                    <Dialog.Title size="2" weight="medium">
                      {title ?? t("datePicker.title")}
                    </Dialog.Title>
                    <Text as="p" size="2" color="gray" mt="1">
                      {t("datePicker.hint")}
                    </Text>
                  </Box>
                </Flex>
                <IconButton type="button" variant="ghost" size="2" aria-label={t("datePicker.cancel")} onClick={handleCancel}>
                  <Cross2Icon width="15" height="15" />
                </IconButton>
              </Flex>

              <Box className={styles.panelEdge}>
                <Box className={styles.panelFill}>
                <Box py="2">
                  <Box px="3">
                    <Flex align="center" gap="2">
                        <Box className={styles.summarySize}>
                          <Box className={`${styles.summaryRadius} ${styles.iconInk}`}>
                            <Flex align="center" justify="center" className={`${styles.summaryFill} ${styles.iconGlyph}`}>
                              <LucideIcon icon={CalendarDays} size={18} />
                            </Flex>
                          </Box>
                        </Box>
                      <Box>
                        <Text as="span" size="1" weight="medium" color="gray">
                          {t("datePicker.selected")}
                        </Text>
                        <Text as="p" size="2" weight="medium">
                          {summary}
                        </Text>
                      </Box>
                    </Flex>
                  </Box>
                </Box>
                </Box>
              </Box>

              <Grid columns={{ initial: "1", sm: "minmax(0, 1fr) 220px" }} gap="3" align="stretch">
                <Box className={`${styles.panelEdge} ${styles.calendarPanel}`}>
                  <Box className={styles.panelFill}>
                  <Box py="2">
                    <Box px="2">
                      <Flex align="center" justify="center" overflow="hidden">
                        <DatePicker
                          key={pickerEpoch}
                          selected={selected}
                          onChange={(date: Nullable<Date>) => {
                            if (!date) {
                              setSelected(null);
                              return;
                            }
                            const next = clampDate(date, minDate, maxDate);
                            setSelected(next);
                            setCalendarMonth(next);
                          }}
                          inline
                          locale={locale}
                          minDate={minDate}
                          maxDate={maxDate}
                          openToDate={calendarMonth}
                          onMonthChange={setCalendarMonth}
                          onYearChange={setCalendarMonth}
                          renderCustomHeader={renderCustomHeader}
                          dateFormat="yyyy-MM-dd"
                          calendarStartDay={1}
                        />
                      </Flex>
                    </Box>
                  </Box>
                  </Box>
                </Box>

                <Flex direction="column" gap="3">
                  <Button type="button" variant="outline" size="2" onClick={handleToday}>
                    <LucideIcon icon={CalendarDays} size={16} />
                    {t("datePicker.today")}
                  </Button>
                  <Box className={styles.panelEdge}>
                    <Box className={styles.noteFill}>
                    <Box py="3">
                      <Box px="3">
                        <Flex align="start" gap="2">
                          <Text color="gray">
                            <LucideIcon icon={CalendarDays} size={18} />
                          </Text>
                          <Text as="p" size="2" color="gray">
                            {t("datePicker.note")}
                          </Text>
                        </Flex>
                      </Box>
                    </Box>
                    </Box>
                  </Box>
                </Flex>
              </Grid>

              <Flex gap="2" justify="end" align="center" wrap="wrap">
                {allowClear ? (
                  <Button type="button" variant="outline" color="gray" onClick={handleClear} mr="auto">
                    {t("datePicker.clear")}
                  </Button>
                ) : null}
                <Button type="button" variant="outline" onClick={handleCancel}>
                  {t("datePicker.cancel")}
                </Button>
                <Button type="button" onClick={handleConfirm} disabled={!selected}>
                  <CheckIcon />
                  {t("datePicker.confirm")}
                </Button>
              </Flex>
            </Flex>
          </Box>
        </Box>
      </Dialog.Content>
    </Dialog.Root>
  );
};

DateTimePicker.displayName = "DateTimePicker";

export default DateTimePicker;
