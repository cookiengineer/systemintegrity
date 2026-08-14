package main

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "os"

func main() {

	app := gtk.NewApplication("engineer.cookie.systemintegrity")

	app.OnActivate(func() {

		controller := controllers.NewApp()
		window     := gui.NewWindow(app, controller)

		window.Present()

	})

	os.Exit(app.Run())

}
