package main

import (
	"fmt"
    "os"
    "path/filepath"
    "io/ioutil"
	"os/exec"
	"strings"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ktr0731/go-fuzzyfinder"
)

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

func getPatternDir() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    return filepath.Join(homeDir, ".config", "fabric", "patterns")
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

// generateFabricCommand creates and executes the Fabric command
func generateFabricCommand(question string) {
    pattern, err := selectPattern()
    if err != nil {
        fmt.Printf("Error selecting pattern: %v\n", err)
        return
    }

    // Define all available options with descriptions
    optionsMap := map[string]string{
        "-s":   "stream output",
        "-c":   "use context file",
        "save": "save to file",
    }

    // Create a slice of options with descriptions for the prompt
    var allOptionsWithDesc []string
    for option, desc := range optionsMap {
        if desc != "" {
            allOptionsWithDesc = append(allOptionsWithDesc, fmt.Sprintf("%s (%s)", option, desc))
        } else {
            allOptionsWithDesc = append(allOptionsWithDesc, option)
        }
    }

    // Ask for options, without any pre-selected
    selectedWithDesc := []string{}
    prompt := &survey.MultiSelect{
        Message: "Select options (use space to toggle, enter to confirm):",
        Options: allOptionsWithDesc,
        Default: []string{}, // Empty slice means no default selection
    }
    survey.AskOne(prompt, &selectedWithDesc)

    // Parse the selected options to remove descriptions
    var selected []string
    for _, optionWithDesc := range selectedWithDesc {
        for option := range optionsMap {
            if optionWithDesc == option || optionWithDesc == fmt.Sprintf("%s (%s)", option, optionsMap[option]) {
                selected = append(selected, option)
                break
            }
        }
    }

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
	fmt.Println("Answer:")

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