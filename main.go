package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ktr0731/go-fuzzyfinder"
)

// Version of the program
const version = "1.0.0"

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

// main is the entry point of the program
func main() {
	// Check for help or version flags
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			printHelp()
			return
		case "--version":
			fmt.Printf("FabricPrompt version %s\n", version)
			return
		}
	}

	fmt.Println("Enter your query (type or paste your text, then press Enter and Ctrl-D to finish):")
	fmt.Println("Or press Ctrl-B and then Enter to use clipboard content directly.")
	fmt.Println("To cancel, press Ctrl-C\n")

	reader := bufio.NewReader(os.Stdin)
	var query strings.Builder
	var usedClipboard bool

	// Read input from user
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading input:", err)
			return
		}

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "\x02" { // Ctrl-B
			usedClipboard = true
			break
		}

		query.WriteString(line)
	}

	// Process the final query
	var finalQuery string
	if usedClipboard {
		finalQuery = getClipboardContent()
	} else {
		finalQuery = strings.TrimSpace(query.String())
	}

	if finalQuery == "" {
		fmt.Println("No input provided. Exiting.")
		return
	}

	generateFabricCommand(finalQuery, usedClipboard)
}

// printHelp displays the help message
func printHelp() {
	fmt.Println("FabricPrompt - A CLI tool to interact with Fabric patterns")
	fmt.Println("\nUsage:")
	fmt.Println("  fabricp [OPTIONS]")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  --version      Show version information")
	fmt.Println("\nDescription:")
	fmt.Println("  FabricPrompt allows you to input a query and select a Fabric pattern to process it.")
	fmt.Println("  You can input your query manually or use the clipboard content.")
	fmt.Println("  After entering your query, you can select additional options for the Fabric command.")
}

// getClipboardContent retrieves the content from the system clipboard
func getClipboardContent() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbpaste")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard", "-o")
	default:
		fmt.Println("Unsupported operating system")
		return ""
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting clipboard content: %v\n", err)
		return ""
	}
	return string(output)
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

func sanitizeFilename(filename string) string {
    // Replace any non-alphanumeric characters (except periods, hyphens, and underscores) with underscores
    reg := regexp.MustCompile(`[^a-zA-Z0-9.-]`)
    sanitized := reg.ReplaceAllString(filename, "_")
    
    // Remove leading and trailing periods
    return strings.Trim(sanitized, ".")
}

// generateFabricCommand creates and executes the Fabric command
func generateFabricCommand(question string, usedClipboard bool) {
    pattern, err := selectPattern()
    if err != nil {
        fmt.Printf("Error selecting pattern: %v\n", err)
        return
    }

    // Define all available options
    allOptions := []string{"-s", "-c", "save"}
    
    // Ask for options, without any pre-selected
    selected := []string{}
    prompt := &survey.MultiSelect{
        Message: "Select options (use space to toggle, enter to confirm):",
        Options: allOptions,
        Default: []string{}, // Empty slice means no default selection
    }
    survey.AskOne(prompt, &selected)

    var options []string
	var saveFilename string
	for _, option := range selected {
		if option == "save" {
			for {
				filenamePrompt := &survey.Input{
					Message: "Enter filename to save:",
				}
				err := survey.AskOne(filenamePrompt, &saveFilename)
				if err != nil {
					fmt.Println("Error reading input:", err)
					return
				}
				
				saveFilename = sanitizeFilename(strings.TrimSpace(saveFilename))
				
				if saveFilename != "" {
					break
				}
				fmt.Println("Filename cannot be empty. Please try again.")
			}
		} else {
			options = append(options, option)
		}
	}

    // Display query information
    if usedClipboard {
        truncatedQuestion := truncateQuestion(question)
        fmt.Println("\nRunning query:")
        fmt.Println(truncatedQuestion)
        if truncatedQuestion != question {
            fmt.Println("(Query truncated for display. Full query will be used in the command.)")
        }
    } else {
        fmt.Println("\nRunning your query")
    }

    fmt.Printf("with pattern: %s\n", pattern)
    if len(options) > 0 {
        fmt.Printf("and options: %s\n", strings.Join(options, " "))
    } else {
        fmt.Println("with no additional options")
    }
    if saveFilename != "" {
        fmt.Printf("Saving output to: %s\n", saveFilename)
    }
    fmt.Println()

    // Escape backticks and single quotes in the question to prevent command injection
    escapedQuestion := strings.ReplaceAll(strings.ReplaceAll(question, "`", "\\`"), "'", "'\\''")

    // Build the command
    fabricCmd := fmt.Sprintf("fabric -p %s %s", pattern, strings.Join(options, " "))
    command := fmt.Sprintf("echo '%s' | %s", escapedQuestion, fabricCmd)

    if saveFilename != "" {
        command += fmt.Sprintf(" | save '%s'", saveFilename)
    }

    cmd := exec.Command("bash", "-c", command)

    if saveFilename == "" {
        // If not saving, we stream the output to console
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
    }

    err = cmd.Run()
    if err != nil {
        fmt.Printf("Command finished with error: %v\n", err)
    }

    if saveFilename != "" {
        fmt.Printf("Output saved to %s\n", saveFilename)
    }
}

// selectPattern allows the user to choose a Fabric pattern
func selectPattern() (string, error) {
	patterns, err := getPatterns()
	if err != nil {
		return "", fmt.Errorf("error fetching patterns: %v", err)
	}

	idx, err := fuzzyfinder.Find(
		patterns,
		func(i int) string {
			return patterns[i]
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			content, err := ioutil.ReadFile(filepath.Join(getPatternDir(), patterns[i], "system.md"))
			if err != nil {
				return fmt.Sprintf("Could not read file: %v", err)
			}
			return string(content)
		}),
	)

	if err != nil {
		return "", err
	}

	return patterns[idx], nil
}

// getPatterns retrieves all available Fabric patterns
func getPatterns() ([]string, error) {
	patternDir := getPatternDir()
	entries, err := ioutil.ReadDir(patternDir)
	if err != nil {
		return nil, err
	}

	var patterns []string
	for _, entry := range entries {
		if entry.IsDir() {
			patterns = append(patterns, entry.Name())
		}
	}
	return patterns, nil
}

// getPatternDir returns the directory path for Fabric patterns
func getPatternDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".config", "fabric", "patterns")
}

// init sets up the initial environment
func init() {
	os.Setenv("LANG", "en_US.UTF-8")
}