package tests

import (
	"fmt"
	"window-go/colors"
	. "window-go/ui/gui"
)

func TestDialogApp() {
	// Clear screen and get terminal dimensions
	fmt.Print(ClearScreenAndBuffer())
	termWidth := GetTerminalWidth()
	termHeight := GetTerminalHeight()

	// Window setup
	winWidth := termWidth * 8 / 10
	winHeight := termHeight * 8 / 10
	if winWidth < 80 {
		winWidth = 80
	}
	if winHeight < 24 {
		winHeight = 24
	}
	winX := (termWidth - winWidth) / 2
	winY := (termHeight - winHeight) / 2

	// Create main window
	win := NewWindow("💬", "Window-Go Dialog Demo", winX, winY, winWidth, winHeight,
		"rounded", colors.BoldMagenta, colors.Magenta, colors.BgBlack, colors.White)

	// Track current status
	var statusLabel *Label
	var currentPrompt *Prompt
	var currentDialog *Prompt

	// Helper function to update status
	updateStatus := func(text string, color string) {
		if statusLabel != nil {
			statusLabel.Text = text
			statusLabel.Color = color
		}
	}

	// Add status label
	statusLabel = NewLabel("Select a dialog type to display", 2, 2, colors.Gray)
	win.AddElement(statusLabel)

	// Create different dialog type buttons
	buttonY := 5
	buttonSpacing := 2

	// Single Line Prompt
	singleLineBtn := NewButton("Single Line Prompt", 2, buttonY, 20, colors.BoldCyan, colors.BgWhite+colors.Cyan, func() bool {
		if currentPrompt != nil {
			currentPrompt.SetActive(false)
			win.RemoveElement(currentPrompt)
			currentPrompt = nil
		}
		buttons := []*PromptButton{
			NewPromptButton("Yes", colors.Green, colors.BgWhite+colors.BoldGreen, func() bool {
				updateStatus("User selected Yes!", colors.Green)
				win.RemoveElement(currentPrompt)
				currentPrompt = nil
				return false
			}),
			NewPromptButton("No", colors.Red, colors.BgWhite+colors.BoldRed, func() bool {
				updateStatus("User selected No!", colors.Red)
				win.RemoveElement(currentPrompt)
				currentPrompt = nil
				return false
			}),
		}
		currentPrompt = NewSingleLinePrompt(
			"Confirm",
			"Do you want to proceed?",
			2, winHeight-4, winWidth-4,
			colors.BoldWhite, colors.White,
			buttons,
		)
		win.AddElement(currentPrompt)
		currentPrompt.SetActive(true)
		return false
	})
	win.AddElement(singleLineBtn)

	// Info Dialog
	infoBtn := NewButton("Info Dialog", 2, buttonY+buttonSpacing, 20, colors.BoldBlue, colors.BgWhite+colors.Blue, func() bool {
		if currentDialog != nil {
			currentDialog.SetActive(false)
			win.RemoveElement(currentDialog)
			currentDialog = nil
		}
		buttons := []*PromptButton{
			NewPromptButton("OK", colors.BoldWhite, colors.BgWhite+colors.Blue, func() bool {
				updateStatus("Info dialog closed", colors.Blue)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
		}
		currentDialog = NewDialogPrompt(
			"Information",
			"This is an information dialog box.\nIt can contain multiple lines of text and will automatically adjust its size based on content.",
			winWidth/4, winHeight/4, winWidth/2,
			colors.BgBlue, colors.Blue, colors.BoldWhite, colors.White,
			buttons,
		)
		win.AddElement(currentDialog)
		currentDialog.SetActive(true)
		return false
	})
	win.AddElement(infoBtn)

	// Warning Dialog
	warningBtn := NewButton("Warning Dialog", 2, buttonY+buttonSpacing*2, 20, colors.BoldYellow, colors.BgWhite+colors.Yellow, func() bool {
		if currentDialog != nil {
			currentDialog.SetActive(false)
			win.RemoveElement(currentDialog)
			currentDialog = nil
		}
		buttons := []*PromptButton{
			NewPromptButton("Continue", colors.BoldYellow, colors.BgWhite+colors.Yellow, func() bool {
				updateStatus("Warning acknowledged", colors.Yellow)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
			NewPromptButton("Cancel", colors.BoldRed, colors.BgWhite+colors.Red, func() bool {
				updateStatus("Warning dialog cancelled", colors.Red)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
		}
		currentDialog = NewDialogPrompt(
			"Warning",
			"This operation might have unexpected consequences.\nAre you sure you want to continue?",
			winWidth/4, winHeight/4, winWidth/2,
			colors.BgYellow, colors.Yellow, colors.BoldBlack, colors.Black,
			buttons,
		)
		win.AddElement(currentDialog)
		currentDialog.SetActive(true)
		return false
	})
	win.AddElement(warningBtn)

	// Error Dialog
	errorBtn := NewButton("Error Dialog", 2, buttonY+buttonSpacing*3, 20, colors.BoldRed, colors.BgWhite+colors.Red, func() bool {
		if currentDialog != nil {
			currentDialog.SetActive(false)
			win.RemoveElement(currentDialog)
			currentDialog = nil
		}
		buttons := []*PromptButton{
			NewPromptButton("Retry", colors.BoldYellow, colors.BgWhite+colors.Yellow, func() bool {
				updateStatus("Retrying operation...", colors.Yellow)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
			NewPromptButton("Abort", colors.BoldRed, colors.BgWhite+colors.Red, func() bool {
				updateStatus("Operation aborted", colors.Red)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
		}
		currentDialog = NewDialogPrompt(
			"Error",
			"An error occurred while processing your request.\nError code: 0x8007000E\nWould you like to retry the operation?",
			winWidth/4, winHeight/4, winWidth/2,
			colors.BgRed, colors.Red, colors.BoldWhite, colors.White,
			buttons,
		)
		win.AddElement(currentDialog)
		currentDialog.SetActive(true)
		return false
	})
	win.AddElement(errorBtn)

	// Custom Dialog
	customBtn := NewButton("Custom Dialog", 2, buttonY+buttonSpacing*4, 20, colors.BoldMagenta, colors.BgWhite+colors.Magenta, func() bool {
		if currentDialog != nil {
			currentDialog.SetActive(false)
			win.RemoveElement(currentDialog)
			currentDialog = nil
		}
		buttons := []*PromptButton{
			NewPromptButton("Option 1", colors.BoldCyan, colors.BgWhite+colors.Cyan, func() bool {
				updateStatus("Selected Option 1", colors.Cyan)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
			NewPromptButton("Option 2", colors.BoldGreen, colors.BgWhite+colors.Green, func() bool {
				updateStatus("Selected Option 2", colors.Green)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
			NewPromptButton("Option 3", colors.BoldYellow, colors.BgWhite+colors.Yellow, func() bool {
				updateStatus("Selected Option 3", colors.Yellow)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
			NewPromptButton("Cancel", colors.BoldRed, colors.BgWhite+colors.Red, func() bool {
				updateStatus("Custom dialog cancelled", colors.Red)
				win.RemoveElement(currentDialog)
				currentDialog = nil
				return false
			}),
		}
		currentDialog = NewDialogPrompt(
			"Custom Dialog",
			"This is a custom dialog with multiple options.\nUse arrow keys to navigate between buttons.\nPress Enter to select an option.",
			winWidth/4, winHeight/4, winWidth/2,
			colors.BgMagenta, colors.Magenta, colors.BoldWhite, colors.White,
			buttons,
		)
		win.AddElement(currentDialog)
		currentDialog.SetActive(true)
		return false
	})
	win.AddElement(customBtn)

	// Add quit button
	quitBtn := NewButton("Quit", 2, buttonY+buttonSpacing*6, 20, colors.BoldRed, colors.BgWhite+colors.Red, func() bool {
		return true
	})
	win.AddElement(quitBtn)

	// Add instructions
	instructions := []string{
		"Instructions:",
		"• Use Tab/Shift+Tab to navigate between buttons",
		"• Press Enter to show the selected dialog type",
		"• In dialogs, use arrow keys to select buttons",
		"• Press Enter to activate the selected button",
		"• Press Escape to close non-modal dialogs",
		"• Press 'q' or Ctrl+C to quit",
	}

	for i, text := range instructions {
		label := NewLabel(text, winWidth/2-20, buttonY+i, colors.Gray)
		win.AddElement(label)
	}

	// Start the window interaction loop
	win.WindowActions()
}
