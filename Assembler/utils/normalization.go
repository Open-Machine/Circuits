package utils

import (
	"regexp"
	"strings"
)

func LineNormalization(line string) string {
	lowerCaseStr := strings.ToLower(line)
	return removeUnecessarySpaces(lowerCaseStr)
}

func removeUnecessarySpaces(line string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(line, " ")
	return strings.TrimSpace(str)
}
