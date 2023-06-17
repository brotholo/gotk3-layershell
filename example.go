package main

import (
	"os"

	"github.com/brotholo/gotk3-layershell/layershell"
	//  "github.com/gotk3/gotk3/gtk"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main2() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init()

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	//  win := gtk.NewWindow()
	app := gtk.NewApplication("stacip", gio.ApplicationFlagsNone)

	// Recursively show all widgets contained in this window.
	//  win.ShowAll()
	app.ConnectActivate(func() { activate2(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
	//  gtk.Window
	//  win.Sho
	//  win.
	//  fmt.Println("STOCAZZO")

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	//  gtk.Main()
}
func activate2(app *gtk.Application) {
	winapp := gtk.NewApplicationWindow(app)
	win := &winapp.Window
	//  fmt.Println(win.
	//  gtk.top
	//  if err != nil {
	//  log.Fatal("Unable to create window:", err)
	//  }
	layershell.InitForWindow(win)

	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_LEFT, true)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_TOP, true)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_RIGHT, true)

	layershell.SetLayer(win, layershell.LAYER_SHELL_LAYER_BOTTOM)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_TOP, 0)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_LEFT, 0)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_RIGHT, 0)

	layershell.SetExclusiveZone(win, 200)

	winapp.SetTitle("Simple Example")
	//  win.Connect("destroy", func() {
	//  //  gtk.MainQuit()
	//  //  gtk.Qui
	//  })

	// Create a new label widget to show in the window.
	l := gtk.NewLabel("Hello, gotk3!")
	//  l, err := gtk.NewLabel("Hello, gotk3!")
	//  if err != nil {
	//  log.Fatal("Unable to create label:", err)
	//  }

	// Add the label to the window.
	//  win.Add(l)
	winapp.SetChild(l)

	// Set the default window size.
	winapp.SetDefaultSize(800, 30)
	winapp.Show()
}
