package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

// Version of the program
const version = "1.0.1"

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
	fmt.Println("  You can input your query manually after starting the program, or pipe text directly in to it.")
	fmt.Println("  After entering your query, you can select additional options for the Fabric command.")
}

// the help text for the tview application
func tviewHelp() (*tview.TextView, *tview.TextView, *tview.TextView) {
	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Navigation

[yellow]Arrowkeys[white]: Move cursor.
[yellow]Ctrl-A, Home[white]: Move to the beginning of the current line.
[yellow]Ctrl-E, End[white]: Move to the end of the current line.
[yellow]Ctrl-F, page down[white]: Move down by one page.
[yellow]Ctrl-B, page up[white]: Move up by one page.
[yellow]Alt-Up arrow[white]: Scroll the page up.
[yellow]Alt-Down arrow[white]: Scroll the page down.
[yellow]Alt-Left arrow[white]: Scroll the page to the left.
[yellow]Alt-Right arrow[white]:  Scroll the page to the right.
[yellow]Alt-B, Ctrl-Left arrow[white]: Move back by one word.
[yellow]Alt-F, Ctrl-Right arrow[white]: Move forward by one word.

[blue]Press Enter for more help, press Escape to return.`)
	help2 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`)
	help3 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.
		
[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`)

	return help1, help2, help3
}

func init() {
    os.Setenv("LANG", "en_US.UTF-8")
}