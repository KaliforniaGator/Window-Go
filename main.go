package main

import (
	"flag"
	"fmt"
	"os"
	"window-go/tests"
	"window-go/types"
)

// demoApps stores all registered demo applications
var demoApps = make(map[int]types.DemoApp)

// registerDemoApp adds a demo app to the registry
func registerDemoApp(app types.DemoApp) {
	demoApps[app.ID] = app
}

// initializeDemoApps registers all available demo applications
func initializeDemoApps() {
	registerDemoApp(types.DemoApp{
		ID:          1,
		Name:        "Freedom Task",
		Description: "A task management demo application",
		RunApp:      tests.TestWindowApp,
	})

	registerDemoApp(types.DemoApp{
		ID:          2,
		Name:        "Segmented Notes",
		Description: "A note-taking demo application",
		RunApp:      tests.TestSegmentsApp,
	})
}

func printUsage() {
	fmt.Println("Window-Go Demo Apps")
	fmt.Println("\nUsage:")
	fmt.Printf("  window-go -app <number>\n\n")
	fmt.Println("Available Apps:")

	for id, app := range demoApps {
		fmt.Printf("  %d    %s - %s\n", id, app.Name, app.Description)
	}

	fmt.Println("\nExample:")
	fmt.Println("  window-go -app 1    # Run the Freedom Task demo")
}

func main() {
	// Initialize demo apps registry
	initializeDemoApps()

	// Define flags
	appID := flag.Int("app", 0, "ID of the demo app to run")

	// Parse command line arguments
	flag.Parse()

	// Check if valid app ID was provided
	if *appID == 0 {
		printUsage()
		os.Exit(1)
	}

	// Find and run the selected demo
	if app, exists := demoApps[*appID]; exists {
		fmt.Printf("Running %s demo...\n", app.Name)
		app.RunApp()
	} else {
		fmt.Printf("Error: No demo app found with ID %d\n", *appID)
		printUsage()
		os.Exit(1)
	}
}
