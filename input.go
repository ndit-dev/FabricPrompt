package main

import (
    "fmt"
    "io"
    "os"
    "strings"

    "github.com/manifoldco/promptui"
)

func getInput() (string, bool) {
    // Check if there is piped input
    stat, err := os.Stdin.Stat()
    if err != nil {
        fmt.Println("Error checking stdin:", err)
        return "", false
    }

    if (stat.Mode() & os.ModeCharDevice) == 0 {
        // There is piped input
        var pipedInput strings.Builder
        _, err := io.Copy(&pipedInput, os.Stdin)
        if err != nil {
            fmt.Println("Error reading piped input:", err)
            return "", false
        }
        return strings.TrimSpace(pipedInput.String()), false
    }

    // Use promptui for manual text input
    prompt := promptui.Prompt{
        Label: "Prompt",
        Validate: func(input string) error {
            return nil
        },
    }

    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return "", false
    }

    return result, false
}