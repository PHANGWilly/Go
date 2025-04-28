package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[-=~*]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+([A-z,0-9]+)\s+`)
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			taggedLines[i] = "[USR] " + match[1] + " " + line
		} else {
			taggedLines[i] = line
		}
	}
	return taggedLines
}
