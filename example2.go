package main

import (
	"os"

	"github.com/brotholo/gotk3-layershell/layershell"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.simple", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	window := gtk.NewApplicationWindow(app)

	win := &window.Window
	layershell.InitForWindow(win)

	//  layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_LEFT, true)
	//  layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_TOP, true)
	//  layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_RIGHT, true)

	//  layershell.SetLayer(win, layershell.LAYER_SHELL_LAYER_BOTTOM)
	//  layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_TOP, 0)
	//  layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_LEFT, 0)
	//  layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_RIGHT, 0)

	//  layershell.SetExclusiveZone(win, 200)

	window.SetTitle("gotk4 Example")
	window.SetChild(gtk.NewLabel("Hello from Go!"))
	window.SetDefaultSize(400, 300)
	window.Show()

	//  winapp := gtk.NewApplicationWindow(app)
	//  fmt.Println(win.
	//  gtk.top
	//  if err != nil {
	//  log.Fatal("Unable to create window:", err)
	//  }
}
