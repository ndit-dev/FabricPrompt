package main

import (
    "fmt"
    "os"
)

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

	finalQuery, usedClipboard := getInput()
	
	// Prcoess the query with Fabric
	generateFabricCommand(finalQuery, usedClipboard)
}