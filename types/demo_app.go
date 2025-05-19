package types

// DemoApp represents a demo application in the Window-Go engine
type DemoApp struct {
	ID          int
	Name        string
	Description string
	RunApp      func()
}
