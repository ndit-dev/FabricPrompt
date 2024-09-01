package main

import (
	"regexp"
	"strings"
)

// Option represents a command option with a description
type Option struct {
	Value       string
	Description string
}

// Options available for the Fabric command
var options = []Option{
	{"-s", "-s:     Use this option if you want to see the results in realtime"},
	{"-c", "-c:     Use Context file (context.md) to add context"},
	{"save", "save:   Save the output to a file"},
}

func sanitizeFilename(filename string) string {
    // Replace any non-alphanumeric characters (except periods, hyphens, and underscores) with underscores
    reg := regexp.MustCompile(`[^a-zA-Z0-9.-]`)
    sanitized := reg.ReplaceAllString(filename, "_")
    
    // Remove leading and trailing periods
    return strings.Trim(sanitized, ".")
}

// truncateQuestion shortens the question for display purposes
func truncateQuestion(question string) string {
	lines := strings.Split(question, "\n")
	var truncated []string
	var totalChars int

	for i, line := range lines {
		if i >= 10 || totalChars+len(line) > 1000 {
			truncated = append(truncated, "...")
			break
		}
		truncated = append(truncated, line)
		totalChars += len(line)
	}

	return strings.Join(truncated, "\n")
}