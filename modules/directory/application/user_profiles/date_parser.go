package user_profiles

import (
	"time"
)

// parseDateString 尝试解析多种日期格式
// 支持：
// - RFC3339 (2006-01-02T15:04:05Z07:00)
// - Date only (2006-01-02)
func parseDateString(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	// 尝试 RFC3339 格式
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return &t, nil
	}

	// 尝试日期格式 (YYYY-MM-DD)
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		// 设置为当天的开始时间 (UTC)
		utcTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		return &utcTime, nil
	}

	// 如果都失败了，返回错误
	return nil, &time.ParseError{
		Layout:  "RFC3339 or 2006-01-02",
		Value:   dateStr,
		Message: "unable to parse date string",
	}
}
