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

### Interfaces for Extensibility

*   **`UIElement`:** Common interface for all renderable elements.
*   **`KeyStrokeHandler`:** Allows custom window-level keyboard input processing.
*   **`CursorManager`:** For elements that need to control cursor visibility and position (e.g., `TextBox`, `TextArea`).
*   **`ZIndexer`:** For elements that need to be rendered in a specific stacking order (e.g., `Menu`, `Prompt`).