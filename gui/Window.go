package gui

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "github.com/cookiengineer/systemintegrity/gui/views"

type Window struct {
	*gtk.Window
	controller *controllers.App
	sidebar    *gtk.StackSidebar
	Packages   *views.Packages
	Boot       *views.Boot
	Devices    *views.Devices
}

func NewWindow(app *gtk.Application, controller *controllers.App) *Window {

	window := gtk.NewWindow(app)
	window.SetTitle("System Integrity Check")
	window.SetDefaultSize(960, 680)

	stack := gtk.NewStack()
	stack.SetHExpand(true)
	stack.SetVExpand(true)

	sidebar := gtk.NewStackSidebar()
	sidebar.SetStack(stack)
	sidebar.SetSizeRequest(220, -1)
	sidebar.SetSensitive(false)

	welcome  := views.NewWelcome(controller)
	packages := views.NewPackages(controller)
	boot     := views.NewBoot(controller)
	devices  := views.NewDevices(controller)

	stack.AddTitled(welcome.AsPtr(),  "welcome",  "Welcome")
	stack.AddTitled(packages.AsPtr(), "packages", "Packages")
	stack.AddTitled(boot.AsPtr(),     "boot",     "Boot")
	stack.AddTitled(devices.AsPtr(),  "devices",  "Devices")

	stack.SetVisibleChildName("welcome")

	layout := gtk.NewBox(gtk.OrientationHorizontal, 0)
	layout.Append(sidebar.AsPtr())
	layout.Append(stack.AsPtr())

	window.SetChild(layout.AsPtr())

	self := &Window{
		Window:     window,
		controller: controller,
		sidebar:    sidebar,
		Packages:   packages,
		Boot:       boot,
		Devices:    devices,
	}

	welcome.SetOnCollected(func() {
		self.Refresh()
		self.sidebar.SetSensitive(true)
	})

	return self

}

func (window *Window) Refresh() {

	system := window.controller.System()

	window.Packages.Refresh(system)
	window.Boot.Refresh(system)
	window.Devices.Refresh(system)

}
