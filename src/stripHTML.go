package main

import (
	"regexp"
	"strings"
)

func stripHTML(html string) string {
	re := regexp.MustCompile(`<.+?>`)
	return strings.TrimSpace(re.ReplaceAllString(html, ""))
}
