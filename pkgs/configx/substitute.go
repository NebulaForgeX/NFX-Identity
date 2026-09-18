package configx

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// readAndSubstituteToml 读取 TOML 文件并替换 ${VAR_NAME} 占位符
// 如果变量缺失，直接返回错误，不使用默认值
func readAndSubstituteToml(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read toml file: %w", err)
	}

	re := regexp.MustCompile(`\$\{([^}]+)\}`)
	missingVars := make(map[string]bool)

	substituted := re.ReplaceAllStringFunc(string(content), func(match string) string {
		varName := strings.TrimSpace(match[2 : len(match)-1])
		value := os.Getenv(varName)
		if value == "" {
			missingVars[varName] = true
			return match
		}
		return value
	})

	if len(missingVars) > 0 {
		varNames := make([]string, 0, len(missingVars))
		for name := range missingVars {
			varNames = append(varNames, name)
		}
		return "", fmt.Errorf("missing required environment variables: %s", strings.Join(varNames, ", "))
	}

	return substituted, nil
}
