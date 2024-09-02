package main

import (
    "os"
    "fmt"

    "github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getInput() (string, bool) {
	app := tview.NewApplication()

	textArea := tview.NewTextArea().
		SetPlaceholder("Enter your promt here...")
	textArea.SetTitle(" FabricPrompt ").SetBorder(true)
	helpInfo := tview.NewTextView().
		SetText(" Press [yellow]F1[white] for help, press [yellow]Ctrl-V[white] to paste, press [yellow]Ctrl-C[white] to exit").
        SetDynamicColors(true)
	continueInfo := tview.NewTextView().
        SetText("Press [yellow]Ctrl-D[white] or [yellow]Ctrl-Space[white] to continue ").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(continueInfo, 1, 1, 1, 1, 0, 0, false)

	help1, help2, help3 := tviewHelp()
		
	help := tview.NewFrame(help1).
		SetBorders(1, 1, 0, 0, 2, 2)
	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	pages.AddAndSwitchToPage("main", mainView, true).
		AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help") 
			return nil
		}
        // Exit the textArea and continue with the prompt
        if event.Key() == tcell.KeyCtrlD || event.Key() == tcell.KeyCtrlSpace {
            app.Stop()
        }
        // Exit the program
        if event.Key() == tcell.KeyCtrlC {
            app.Stop()
            fmt.Println("User interupted, exiting...")
            os.Exit(0)
        }
        // Paste clipboard content at the cursor position
        if event.Key() == tcell.KeyCtrlV {
            clipboardContent, err := clipboard.ReadAll()
            if err == nil {
                currentText := textArea.GetText()
                newText := currentText + clipboardContent
                textArea.SetText(newText, true)
            }
        }
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	
	// store the prompt in a variable
	prompt := textArea.GetText()

    return prompt, false
}