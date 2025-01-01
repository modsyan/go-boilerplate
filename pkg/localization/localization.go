package localization

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	messages map[string]map[string]string
	lang     string
	mu       sync.RWMutex
)

func SetLang(l string) {
	mu.Lock()
	defer mu.Unlock()
	lang = l
}

func LoadMessages(path string) error {
	// Ensure file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("localization file not found at path: %s", path)
	}

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open locales file: %w", err)
	}
	defer file.Close()

	// Read file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read locales file: %w", err)
	}

	// Parse JSON
	err = json.Unmarshal(bytes, &messages)
	if err != nil {
		return fmt.Errorf("failed to unmarshal locales file: %w", err)
	}

	return nil
}

func L(key string, placeholders ...string) string {
	mu.RLock()
	defer mu.RUnlock()

	if langMessages, ok := messages[lang]; ok {
		if msg, ok := langMessages[key]; ok {
			return replacePlaceholders(msg, langMessages, placeholders...)
		}
	}
	return key
}

func replacePlaceholders(message string, langMessages map[string]string, placeholders ...string) string {
	for i, placeholder := range placeholders {
		placeholderToken := fmt.Sprintf("{%d}", i)
		if localizedPlaceholder, ok := langMessages[placeholder]; ok {
			placeholder = localizedPlaceholder
		}
		message = strings.ReplaceAll(message, placeholderToken, placeholder)
	}
	return message
}
