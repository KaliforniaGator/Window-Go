package gui

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
	"window-go/colors"

	"golang.org/x/term"
)

// ANSI Escape Codes
const (
	clearScreen          = "\x1b[2J"
	clearScreenAndBuffer = "\033[H\033[2J\033[3J"
	moveCursorFormat     = "\x1b[%d;%dH" // row, col (1-based) - Renamed format string
	hideCursor           = "\x1b[?25l"
	showCursor           = "\x1b[?25h"
)

// ClearScreen clears the entire terminal screen.
func ClearScreen() string { // Return string instead of printing directly
	return clearScreen
}

// ClearScreenAndBuffer clears the terminal screen and scrollback buffer.
func ClearScreenAndBuffer() string { // Return string instead of printing directly
	return clearScreenAndBuffer
}

// MoveCursor positions the cursor at the specified row and column.
// Note: row and col are 0-based for convenience, but converted to 1-based for ANSI.
func MoveCursor(row, col int) { // Keep this for direct printing if needed elsewhere
	fmt.Printf(moveCursorFormat, row+1, col+1)
}

// MoveCursorCmd returns the ANSI escape code string to move the cursor.
// Note: row and col are 0-based for convenience.
func MoveCursorCmd(row, col int) string {
	return fmt.Sprintf(moveCursorFormat, row+1, col+1)
}

// HideCursor makes the terminal cursor invisible.
func HideCursor() string { // Return string
	return hideCursor
}

// ShowCursor makes the terminal cursor visible.
func ShowCursor() string { // Return string
	return showCursor
}

// ClearLineSuffix returns ANSI sequence to clear from cursor to end of line
func ClearLineSuffix() string {
	return "\x1b[K"
}

// ResetStyle returns ANSI escape sequence to reset text formatting
func ResetStyle() string {
	return "\x1b[0m"
}

// ReverseVideo returns ANSI escape sequence for reverse video (inverted colors)
func ReverseVideo() string {
	return "\x1b[7m"
}

// ResetVideo returns ANSI escape sequence to reset text formatting
func ResetVideo() string {
	return "\x1b[27m"
}

// BoxType defines the structure for different box styles
type BoxType struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
}

// TextAlignment defines the structure for text alignment
type TextAlignment struct {
	Horizontal string
	Vertical   string
}

var (
	BoxTypes = map[string]BoxType{
		"single": {
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
			Horizontal:  "─",
			Vertical:    "│",
		},
		"double": {
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
			Horizontal:  "═",
			Vertical:    "║",
		},
		"round": {
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "╰",
			BottomRight: "╯",
			Horizontal:  "─",
			Vertical:    "│",
		},
		"bold": {
			TopLeft:     "┏",
			TopRight:    "┓",
			BottomLeft:  "┗",
			BottomRight: "┛",
			Horizontal:  "━",
			Vertical:    "┃",
		},
	}
)

func PrintColoredText(text string, color string) {
	// Print colored text
	fmt.Printf("%s%s%s", color, text, colors.Reset)
}
func PrintError(text string) {
	// Print error message
	fmt.Printf("%s%s%s", colors.BoldRed, text, colors.Reset)
}
func PrintSuccess(text string) {
	// Print success message
	fmt.Printf("%s%s%s", colors.BoldGreen, text, colors.Reset)
}
func PrintWarning(text string) {
	// Print warning message
	fmt.Printf("%s%s%s", colors.BoldYellow, text, colors.Reset)
}
func PrintInfo(text string) {
	// Print info message
	fmt.Printf("%s%s%s", colors.BoldCyan, text, colors.Reset)
}
func PrintDebug(text string) {
	// Print debug message
	fmt.Printf("%s%s%s", colors.BoldGray, text, colors.Reset)
}
func PrintAlert(text string) {
	// Print alert message
	fmt.Printf("%s%s%s", colors.BoldWhite, text, colors.Reset)
}

func GetTerminalWidth() int {
	// Get the terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // Default width if unable to get terminal size
	}
	return width
}

func GetTerminalHeight() int {
	// Get the terminal height
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 24 // Default height if unable to get terminal size
	}
	return height
}

func newLine() {
	// Print a new line
	fmt.Print("\n")
}

// Estimate the width of a string based on average character width
func EstimateStringWidth(s string) int {
	// Assume an average width of 8 pixels per character
	// You can adjust this value based on your needs
	const averageCharWidth = 8

	// Count the number of runes (characters) in the string
	charCount := utf8.RuneCountInString(s)

	// Calculate the estimated width
	return charCount * averageCharWidth
}

func NormalizeWidth(text string) int {
	maxWidth := GetTerminalWidth() / 2

	width := EstimateStringWidth(text)
	if width > maxWidth {
		width = maxWidth - 4
	}
	return width
}

func TitleBox(text string) {
	width := -1
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "double", colors.BoldWhite, "", colors.BoldWhite, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}

func ErrorBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldRed, "", colors.BoldRed, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}
func SuccessBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldGreen, "", colors.BoldGreen, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}
func WarningBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldYellow, "", colors.BoldYellow, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}
func InfoBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldCyan, "", colors.BoldCyan, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}
func DebugBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldGray, "", colors.BoldGray, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}
func AlertBox(text string) {
	width := NormalizeWidth(text)
	height := -1
	if width < 0 {
		width = len(text) + 4
	}
	if height < 0 {
		height = 3
	}
	newLine()
	PrintBanner(text, "single", colors.BoldYellow, "", colors.BoldYellow, width, height, TextAlignment{
		Horizontal: "center",
		Vertical:   "center",
	})
	fmt.Println()
}

func wrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	currentLine := words[0]
	for _, word := range words[1:] {
		if len(currentLine)+len(word)+1 <= width {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

func PrintBanner(text string, boxStyle string, textColor string, bgColor string, borderColor string, width int, height int, alignment TextAlignment) {
	fmt.Print(colors.Reset)
	box, exists := BoxTypes[boxStyle]
	if !exists {
		box = BoxTypes["single"]
	}

	padding := 2
	effectiveWidth := width - (padding * 2)
	wrappedText := wrapText(text, effectiveWidth)
	textHeight := len(wrappedText)

	if width < padding*2 {
		width = padding * 2
	}

	if height < textHeight+2 {
		height = textHeight + 2
	}

	// Top border with border color
	fmt.Print(bgColor + borderColor)
	fmt.Print(box.TopLeft)
	for i := 0; i < width; i++ {
		fmt.Print(box.Horizontal)
	}
	fmt.Print(box.TopRight + "\n")

	// Calculate vertical position
	var startRow int
	switch alignment.Vertical {
	case "top":
		startRow = padding
	case "bottom":
		startRow = height - padding - textHeight
	default: // center
		startRow = height/2 - textHeight/2 - 1
	}

	// Print empty lines before text
	for i := 1; i < startRow; i++ {
		fmt.Print(borderColor + box.Vertical + colors.Reset + bgColor)
		for j := 0; j < width; j++ {
			fmt.Print(" ")
		}
		fmt.Print(borderColor + box.Vertical + "\n")
	}

	// Print text lines
	for _, line := range wrappedText {
		fmt.Print(borderColor + box.Vertical + colors.Reset + textColor + bgColor)
		lineLength := len(line)
		leftPadding := padding

		switch alignment.Horizontal {
		case "left":
			leftPadding = padding
		case "right":
			leftPadding = width - lineLength - padding
		default: // center
			leftPadding = (width - lineLength) / 2
		}

		for i := 0; i < leftPadding; i++ {
			fmt.Print(" ")
		}
		fmt.Print(textColor + line + colors.Reset + bgColor)
		rightPadding := width - leftPadding - lineLength
		for i := 0; i < rightPadding; i++ {
			fmt.Print(" ")
		}
		fmt.Print(borderColor + box.Vertical + "\n")
	}

	// Print empty lines after text
	for i := startRow + textHeight + 1; i < height-1; i++ {
		fmt.Print(borderColor + box.Vertical + colors.Reset + bgColor)
		for j := 0; j < width; j++ {
			fmt.Print(" ")
		}
		fmt.Print(borderColor + box.Vertical + "\n")
	}

	// Bottom border
	fmt.Print(borderColor)
	fmt.Print(box.BottomLeft)
	for i := 0; i < width; i++ {
		fmt.Print(box.Horizontal)
	}
	fmt.Print(box.BottomRight)
	fmt.Print(colors.Reset)
}

func PrintBannerColors() {
	// Group colors by type
	colorGroups := map[string][]string{
		"Regular Colors":    {"red", "green", "yellow", "blue", "purple", "cyan", "gray", "white", "black"},
		"Bold Colors":       {"bold_red", "bold_green", "bold_yellow", "bold_blue", "bold_purple", "bold_cyan", "bold_gray", "bold_white", "bold_black"},
		"Background Colors": {"bg_red", "bg_green", "bg_yellow", "bg_blue", "bg_purple", "bg_cyan", "bg_gray", "bg_white", "bg_black"},
		"Gray Shades":       {"gray1", "gray2", "gray3", "gray4", "gray5"},
		"Background Grays":  {"bg_gray1", "bg_gray2", "bg_gray3", "bg_gray4", "bg_gray5"},
	}

	fmt.Println("\nAvailable Colors:")
	for group, colorNames := range colorGroups {
		fmt.Printf("\n%s%s%s\n", colors.BoldWhite, group, colors.Reset)
		for _, name := range colorNames {
			if color, exists := colors.ColorMap[name]; exists {
				fmt.Printf("  %s%-15s%s %s(%-12s)%s\n",
					color, "Sample Text", colors.Reset,
					colors.Gray, name, colors.Reset)
			}
		}
	}
	fmt.Println()
}

func PrintWindow(icon string, title string, content string, bgColor string, borderColor string,
	titleColor string, contentColor string, width int, height int) {

	// Determine position (e.g., centered)
	termWidth := GetTerminalWidth()
	termHeight := GetTerminalHeight()
	winX := (termWidth - width) / 2
	winY := (termHeight - height) / 2
	if winX < 0 {
		winX = 0
	}
	if winY < 0 {
		winY = 0
	}

	// Create a new Window instance
	// Using "single" style as the original PrintWindow implicitly did.
	// Content color is set as the default for the window.
	win := NewWindow(icon, title, winX, winY, width, height, "single", titleColor, borderColor, bgColor, contentColor)

	// Add the main content as a Label element spanning the width
	// Wrap the text first to fit the content area width
	contentWidth := width - 2 // Account for borders
	wrappedContent := wrapText(content, contentWidth)

	// Add each line of wrapped text as a separate Label
	for i, line := range wrappedContent {
		// Position labels starting from top-left (0,0) relative to content area
		// Ensure we don't exceed the window's content height
		if i < height-2 { // Account for top/bottom borders
			label := NewLabel(line, 0, i, contentColor) // Use provided contentColor
			win.AddElement(label)
		} else {
			break // Stop adding lines if window height is exceeded
		}
	}

	// Render the window
	win.Render()

	// Note: The original PrintWindow printed a newline after the content banner.
	// The new Render method places the window absolutely, so a newline might not be needed
	// or desired depending on how it's used in the application flow.
	// If a newline is needed to move the cursor below the window, add:
	// fmt.Print(MoveCursorCmd(winY+height, 0)) // Move cursor below the window
}
