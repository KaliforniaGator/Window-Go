# Window-Go

## Features

### Window Core Features

*   **Customization:**
    *   Set window title and icon.
    *   Define position (X, Y) and dimensions (Width, Height).
    *   Choose from various box drawing styles (e.g., "single", "double").
    *   Customize colors for title, border, background, and default content text.
*   **Element Management:**
    *   Add and remove UI elements dynamically.
    *   Automatic focus management for interactive elements (Buttons, TextBoxes, CheckBoxes, RadioButtons, ScrollBars, Containers, TextAreas, MenuBars, Prompts).
    *   Elements are rendered based on a Z-index, allowing for overlapping (e.g., submenus, prompts).
*   **Interaction:**
    *   Built-in raw terminal input handling for keyboard events (arrow keys, Enter, Tab, Shift+Tab, Backspace, Delete, printable characters, Ctrl+C).
    *   Extensible key stroke handling via `KeyStrokeHandler` interface for custom window-level input processing.
*   **Rendering:**
    *   Efficient rendering using an internal buffer.
    *   Automatic cursor management: shows cursor for active input elements (TextBox, TextArea), hides otherwise.
    *   Smart title truncation with "..." for titles longer than available width.
    *   Basic support for wide character and emoji display width in titles.

### UI Elements

*   **Label:**
    *   Displays simple text.
    *   Customizable text color.
    *   Positionable (X, Y relative to window content area).
    *   Automatic word wrapping within the label's implicitly defined width (based on window content width and label's X position).
*   **Button:**
    *   Clickable button with customizable text.
    *   Define normal and active (focused) colors.
    *   Assign an action (callback function) to be executed on activation (Enter key).
    *   Fixed width.
*   **TextBox:**
    *   Single-line editable text input field.
    *   Normal and active (focused) color customization.
    *   Cursor management (visible when active, moves with input).
    *   Horizontal text scrolling if text exceeds width.
    *   Pristine state: default text can be cleared on first input.
*   **CheckBox:**
    *   Toggleable checkbox with a label.
    *   Normal and active (focused) color customization.
    *   Can be checked or unchecked.
*   **Spacer:**
    *   Provides vertical empty space for layout purposes.
*   **RadioButton & RadioGroup:**
    *   Allows selection of one option from a group.
    *   Each `RadioButton` has a label and an associated value.
    *   Normal and active (focused) color customization.
    *   `RadioGroup` manages the selection state of its buttons.
*   **ProgressBar:**
    *   Visual indicator of progress.
    *   Set current value and maximum value.
    *   Customizable colors for filled and unfilled portions.
    *   Optionally displays percentage text.
*   **GradientProgressBar:**
    *   Progress bar with a two-color gradient fill.
    *   Customizable start and end hex colors for the gradient.
    *   Customizable color for the unfilled portion.
    *   Optionally displays percentage text.
*   **ScrollBar:**
    *   Vertical scrollbar for indicating position within scrollable content.
    *   Customizable height, current value, and maximum value.
    *   Normal and active (focused) color customization.
    *   Can be set to visible or hidden.
    *   `OnScroll` callback triggered when value changes.
*   **Container:**
    *   Scrollable area for displaying a list of string content.
    *   Manages an internal `ScrollBar` which becomes visible if content exceeds container height.
    *   Supports item highlighting (via arrow keys when focused) and selection (via Enter key).
    *   Customizable colors for default content and selected/highlighted item.
    *   `OnItemSelected` callback triggered when an item is selected.
*   **TextArea:**
    *   Multi-line editable text input area.
    *   Supports vertical scrolling with an internal `ScrollBar`.
    *   Cursor management (visible when active, moves with input across lines and columns).
    *   Text manipulation: insert characters (including newlines), delete (Backspace), delete forward (Delete).
    *   Optional maximum character limit.
    *   Displays word count and character count (optionally).
    *   Normal and active (focused) color customization.
*   **MenuBar, Menu, MenuItem:**
    *   Hierarchical menu system.
    *   `MenuBar` is the top-level container.
    *   `Menu` can contain `MenuItem`s or other `Menu`s (submenus).
    *   `MenuItem`s have text, can trigger an action, or open a submenu.
    *   Customizable colors for items (normal/active) and menu background/borders.
    *   Keyboard navigation (arrows, Enter, Escape).
    *   Submenus appear with a Z-index above other elements.
*   **Prompt:**
    *   Displays messages to the user with interactive buttons.
    *   Two styles: `SingleLinePrompt` and `DialogBoxPrompt`.
    *   Customizable title, message, and button text/actions.
    *   Customizable colors for background, border, title, and message.
    *   `DialogBoxPrompt` can be modal, blocking interaction with elements behind it.
    *   Keyboard navigation between buttons (Left/Right arrows or Tab for non-modal).
    *   Renders with a high Z-index to appear above other content.

### Box Drawing Styles

The following box drawing styles are available:
* `single` - Single line borders (┌─┐│└┘)
* `double` - Double line borders (╔═╗║╚╝)
* `round` - Rounded corners (╭─╮│╰╯)
* `bold` - Bold lines (┏━┓┃┗┛)

### Color Support

Window-Go provides extensive color support through the `colors` package:

* Regular Colors: red, green, yellow, blue, purple, cyan, gray, white, black
* Bold Colors: bold_red, bold_green, bold_yellow, bold_blue, bold_purple, bold_cyan, bold_gray, bold_white, bold_black
* Background Colors: bg_red, bg_green, bg_yellow, bg_blue, bg_purple, bg_cyan, bg_gray, bg_white, bg_black
* Gray Shades: gray1, gray2, gray3, gray4, gray5
* Background Grays: bg_gray1, bg_gray2, bg_gray3, bg_gray4, bg_gray5

Colors can be combined using the + operator: `colors.BgBlue + colors.BoldWhite`

### Helper Functions

* `ClearScreen()` - Clears the terminal screen
* `ClearScreenAndBuffer()` - Clears both screen and scrollback buffer
* `MoveCursor(row, col)` - Positions cursor at specific coordinates
* `HideCursor()` / `ShowCursor()` - Controls cursor visibility
* `PrintColoredText()` - Print text with specified color
* `PrintError()` / `PrintSuccess()` / `PrintWarning()` / `PrintInfo()` / `PrintDebug()` / `PrintAlert()` - Print formatted messages
* `GetTerminalWidth()` / `GetTerminalHeight()` - Get terminal dimensions

### Demo Applications

The project includes several demo applications showcasing different features:

1. **Freedom Task** - Task management demo
   * CRUD operations for tasks
   * Priority levels with color coding
   * Progress tracking
   * Scrollable task list with selection

2. **Segmented Notes** - Note-taking demo
   * Split-pane interface
   * Note list with selection
   * Multi-line text editing
   * Word and character counting

3. **Menu Demo** - Menu system showcase
   * Hierarchical menu structure
   * Nested submenus
   * Keyboard navigation
   * Modal and non-modal interactions

4. **Dialog Demo** - Dialog system showcase
   * Single line prompts
   * Modal dialog boxes
   * Various dialog types (info, warning, error)
   * Custom dialog configurations

### Running the Demos

```bash
# Build and run a specific demo
window-go -app <number>

# Available demos:
window-go -app 1    # Freedom Task
window-go -app 2    # Segmented Notes
window-go -app 3    # Menu Demo
window-go -app 4    # Dialog Demo
```

### Common Key Bindings

* `Tab` / `Shift+Tab` - Navigate between interactive elements
* `Arrow Keys` - Navigate within elements (lists, menus)
* `Enter` - Activate buttons, select items
* `Escape` - Close menus, non-modal dialogs
* `Backspace` / `Delete` - Text editing
* `q` or `Ctrl+C` - Quit application

### Element Hierarchy and Z-Index

Elements are rendered in layers based on their Z-index:
1. Regular UI elements (default: 0)
2. MenuBar (100)
3. Menus and Submenus (150)
4. Prompts and Dialogs (1000)

### Custom Key Handling

Implement the `KeyStrokeHandler` interface to add custom keyboard handling:
```go
type KeyStrokeHandler interface {
    HandleKeyStroke(key []byte, w *Window) (handled bool, needsRender bool, shouldQuit bool)
}
```