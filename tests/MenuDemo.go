package tests

import (
	"fmt"
	"window-go/colors"
	. "window-go/ui/gui"
)

func TestMenuApp() {
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
	win := NewWindow("🎯", "Window-Go Menu Demo", winX, winY, winWidth, winHeight,
		"rounded", colors.BoldCyan, colors.Cyan, colors.BgBlack, colors.White)

	// Create menu bar
	menuBar := NewMenuBar(0, 0, winWidth-2, colors.BgGray2, colors.Gray2, colors.BgBlack)

	// File Menu
	fileMenu := menuBar.AddSubMenu("File", colors.White, colors.BgWhite+colors.Black)
	fileMenu.AddItem(NewMenuItem("New", colors.White, colors.BgWhite+colors.Black, func() bool {
		return false
	}))
	fileMenu.AddItem(NewMenuItem("Open", colors.White, colors.BgWhite+colors.Black, func() bool {
		return false
	}))
	fileMenu.AddItem(NewMenuItem("Save", colors.White, colors.BgWhite+colors.Black, func() bool {
		return false
	}))
	fileMenu.AddItem(NewMenuItem("Exit", colors.White, colors.BgWhite+colors.Black, func() bool {
		return true // Quit
	}))

	// Edit Menu with Submenus
	editMenu := menuBar.AddSubMenu("Edit", colors.White, colors.BgWhite+colors.Black)
	formatSubmenu := editMenu.AddSubMenu("Format", colors.White, colors.BgWhite+colors.Black)
	formatSubmenu.AddItem(NewMenuItem("Bold", colors.White, colors.BgWhite+colors.Black, nil))
	formatSubmenu.AddItem(NewMenuItem("Italic", colors.White, colors.BgWhite+colors.Black, nil))
	formatSubmenu.AddItem(NewMenuItem("Underline", colors.White, colors.BgWhite+colors.Black, nil))

	// Add a deeply nested submenu for demonstration
	advancedSubmenu := formatSubmenu.AddSubMenu("Advanced", colors.White, colors.BgWhite+colors.Black)
	advancedSubmenu.AddItem(NewMenuItem("Option 1", colors.White, colors.BgWhite+colors.Black, nil))
	advancedSubmenu.AddItem(NewMenuItem("Option 2", colors.White, colors.BgWhite+colors.Black, nil))

	// View Menu
	viewMenu := menuBar.AddSubMenu("View", colors.White, colors.BgWhite+colors.Black)
	viewMenu.AddItem(NewMenuItem("Zoom In", colors.White, colors.BgWhite+colors.Black, nil))
	viewMenu.AddItem(NewMenuItem("Zoom Out", colors.White, colors.BgWhite+colors.Black, nil))
	viewMenu.AddItem(NewMenuItem("Reset Zoom", colors.White, colors.BgWhite+colors.Black, nil))

	// Help Menu
	helpMenu := menuBar.AddSubMenu("Help", colors.White, colors.BgWhite+colors.Black)
	helpMenu.AddItem(NewMenuItem("Documentation", colors.White, colors.BgWhite+colors.Black, nil))
	helpMenu.AddItem(NewMenuItem("About", colors.White, colors.BgWhite+colors.Black, nil))

	// Add the menu bar to the window
	win.AddElement(menuBar)

	// Add some content below the menu
	instructions := NewLabel("Menu Navigation:", 2, 3, colors.BoldWhite)
	win.AddElement(instructions)

	tips := []string{
		"• Use Tab/Shift+Tab to focus the menu bar",
		"• Use Left/Right arrows to navigate top-level menus",
		"• Use Up/Down arrows to navigate menu items",
		"• Use Enter to activate a menu item",
		"• Use Left arrow to close submenu",
		"• Use Escape to close all menus",
		"• Press 'q' or Ctrl+C to quit",
	}

	for i, tip := range tips {
		tipLabel := NewLabel(tip, 4, 5+i, colors.Gray)
		win.AddElement(tipLabel)
	}

	// Start the window interaction loop
	win.WindowActions()
}
