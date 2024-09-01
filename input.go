package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "os/exec"
    "strings"
)

func getClipboardContent() string {
    cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Error getting clipboard content: %v\n", err)
        return ""
    }
    return string(output)
}

func getInput() (string, bool) {
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
            return "", false
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
        return "", false
    }

    return finalQuery, usedClipboard
}