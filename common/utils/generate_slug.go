package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with -
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non-word characters
	slug = regexp.MustCompile(`[^\w-]+`).ReplaceAllString(slug, "")

	// Replace multiple - with single -
	slug = regexp.MustCompile(`--+`).ReplaceAllString(slug, "-")

	// Trim - from start and end of text
	slug = strings.Trim(slug, "-")

	// Remove quotes and parentheses
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, `"`, "")
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")

	// Get current date in YYYYMMDD format
	dateStr := time.Now().Format("20060102")

	return fmt.Sprintf("%s_%s", slug, dateStr)
}
