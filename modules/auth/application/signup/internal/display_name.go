package internal

import "strings"

func GenerateForgerDisplayName(email string) string {
	localPart := strings.TrimSpace(email)
	if at := strings.IndexByte(localPart, '@'); at >= 0 {
		localPart = localPart[:at]
	}
	localPart = strings.TrimSpace(localPart)
	if localPart == "" {
		localPart = "user"
	}

	runes := []rune(localPart)
	if len(runes) > 10 {
		runes = runes[:10]
	}
	return "Forger-" + string(runes)
}
