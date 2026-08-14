package views

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "github.com/cookiengineer/systemintegrity/structs"
import "strings"

type Boot struct {
	*gtk.Box
	controller  *controllers.App
	description *gtk.Label
}

func NewBoot(controller *controllers.App) *Boot {

	box := gtk.NewBox(gtk.OrientationVertical, 12)
	box.SetMarginStart(16)
	box.SetMarginEnd(16)
	box.SetMarginTop(16)
	box.SetMarginBottom(16)

	title := gtk.NewLabel("Boot Integrity")
	title.SetHAlign(gtk.AlignStart)
	box.Append(title.AsPtr())

	description := gtk.NewLabel("No data collected yet.")
	description.SetWrap(true)
	description.SetXAlign(0.0)
	description.SetVAlign(gtk.AlignStart)
	box.Append(description.AsPtr())

	view := &Boot{
		Box:         box,
		controller:  controller,
		description: description,
	}

	return view

}

func (view *Boot) Refresh(system *structs.System) {

	boot := system.Boot

	if boot.IsValid() == false {
		view.description.SetText("No boot configuration collected.")
		return
	}

	var builder strings.Builder

	if boot.Kernel != "" {
		builder.WriteString("Kernel: " + boot.Kernel + "\n")
	}

	if boot.KernelVersion != "" {
		builder.WriteString("Kernel Version: " + boot.KernelVersion + "\n")
	}

	if boot.KernelArchitecture != "" {
		builder.WriteString("Architecture: " + boot.KernelArchitecture + "\n")
	}

	if boot.Bootloader != "" {
		builder.WriteString("Bootloader: " + boot.Bootloader + "\n")
	}

	if boot.Mode != "" {
		builder.WriteString("Mode: " + boot.Mode + "\n")
	}

	if boot.SecureBoot != "" {
		builder.WriteString("Secure Boot: " + boot.SecureBoot + "\n")
	}

	if boot.Initramfs != "" {
		builder.WriteString("Initramfs: " + boot.Initramfs + "\n")
	}

	if boot.ESP != "" {
		builder.WriteString("EFI System Partition: " + boot.ESP + "\n")
	}

	view.description.SetText(builder.String())

}
