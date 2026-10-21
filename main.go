package main

import (
	"embed"
	"flag"
	"log"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// runs the application and logs any error that might occur.
//
// Startup configuration is supplied by App Builder through command-line flags.
func main() {
	settings := Settings{}
	flag.StringVar(&settings.Project, "project", "", "Path to the project .appDef file")
	flag.StringVar(&settings.Theme, "theme", "light", "Dialog theme: light or dark")
	flag.StringVar(&settings.Language, "lang", "en", "Dialog language ID")
	flag.Parse()

	projectService := &ProjectService{}
	if settings.Project != "" {
		projectService.folder = filepath.Dir(settings.Project)
	}

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "sdc",
		Description: "Upload App Builder project data to the support system",
		Services: []application.Service{
			application.NewService(projectService),
			application.NewService(&SettingsService{settings: settings}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:     "Upload Support Data",
		Width:     800,
		Height:    640,
		MinWidth:  640,
		MinHeight: 480,
		Mac: application.MacWindow{
			TitleBar: application.MacTitleBarDefault,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
