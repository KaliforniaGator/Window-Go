package ui

import (
	"fmt"
	"window-go/colors"
)

func ClearLine() {
	// Clear the entire current line and return carriage
	fmt.Print("\033[2K\r")

}

func ClearLineAndPrintBottom() {
	// Clear the entire current line and return carriage
	fmt.Print("\033[2K\r")
	// Print only the bottom prompt exactly as defined
	fmt.Print(colors.BoldGreen + "╰─" + colors.Reset + "$ ")
}

func NewLine() {
	// Print a new line
	fmt.Print("\n")
}

func ClearScreen() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")
}

func ClearScreenAndBuffer() {
	// Clear the screen and buffer
	fmt.Print("\033[H\033[2J\033[3J")
	// Clear the scrollback buffer
}
