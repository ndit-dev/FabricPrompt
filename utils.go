package main

import (
	"fmt"
	"os"
)

// Version of the program
const version = "1.0.0"

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

func init() {
    os.Setenv("LANG", "en_US.UTF-8")
}