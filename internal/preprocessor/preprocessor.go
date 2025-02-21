package preprocessor

import (
	"regexp"
	"strings"
)

const syntax = "---"

func PreProcess(source string) string {
	// remove entire line after `---`
	comment := regexp.MustCompile(syntax + ".*")
	source = comment.ReplaceAllString(source, "")

	// trim whitespaces and empty lines
	lines := strings.Split(source, "\n")
	var processedLines []string

	for _, line := range lines {
		// trim from both ends
		trimmed := strings.TrimSpace(line)

		// reduce multiple spaces within the line to a single space
		trimmed = regexp.MustCompile(`\s+`).ReplaceAllString(trimmed, " ")

		if len(trimmed) > 0 {
			processedLines = append(processedLines, trimmed)
		}
	}

	return strings.Join(processedLines, "\n")
}
