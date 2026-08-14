package views

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "github.com/cookiengineer/systemintegrity/structs"
import utils_fmt "github.com/cookiengineer/systemintegrity/utils/fmt"
import "fmt"
import "strings"

type Devices struct {
	*gtk.Box
	controller  *controllers.App
	report_view *gtk.TextView
}

func NewDevices(controller *controllers.App) *Devices {

	box := gtk.NewBox(gtk.OrientationVertical, 12)
	box.SetMarginStart(16)
	box.SetMarginEnd(16)
	box.SetMarginTop(16)
	box.SetMarginBottom(16)

	title := gtk.NewLabel("Hardware Integrity")
	title.SetHAlign(gtk.AlignStart)
	box.Append(title.AsPtr())

	report_wrapper := gtk.NewScrolledWindow()
	report_wrapper.SetPolicy(gtk.PolicyAutomatic, gtk.PolicyAutomatic)
	report_wrapper.SetHExpand(true)
	report_wrapper.SetVExpand(true)

	report_view := gtk.NewTextView()
	report_view.SetEditable(false)
	report_view.SetCursorVisible(false)
	report_view.SetMonospace(true)
	report_view.SetTerminalStyle()

	report_wrapper.SetChild(report_view.AsPtr())
	box.Append(report_wrapper.AsPtr())

	view := &Devices{
		Box:         box,
		controller:  controller,
		report_view: report_view,
	}

	return view

}

func (view *Devices) Refresh(system *structs.System) {

	var builder strings.Builder

	builder.WriteString("BIOS\n")
	builder.WriteString(view.render_device(system.BIOS, "  "))

	builder.WriteString("\nBoard\n")
	builder.WriteString(view.render_device(system.Board, "  "))

	builder.WriteString("\nDevices (" + fmt.Sprintf("%d", len(system.Devices)) + ")\n")

	for d := 0; d < len(system.Devices); d++ {
		builder.WriteString(view.render_device(system.Devices[d], "  "))
	}

	builder.WriteString("\nDrives (" + fmt.Sprintf("%d", len(system.Drives)) + ")\n")

	for d := 0; d < len(system.Drives); d++ {
		builder.WriteString(view.render_drive(system.Drives[d], "  "))
	}

	view.report_view.Clear()
	view.report_view.Append(builder.String())

}

func (view *Devices) render_device(device structs.Device, indent string) string {

	var builder strings.Builder

	if device.Name == "" && device.Bus == "" {
		builder.WriteString(indent + "Unknown\n")
		return builder.String()
	}

	builder.WriteString(indent + "Name: " + device.Name + "\n")

	if device.Bus != "" {
		builder.WriteString(indent + "Bus: " + device.Bus + "\n")
	}

	if device.System != nil {

		if device.System.Vendor != "" {
			builder.WriteString(indent + "Vendor: " + device.System.Vendor + "\n")
		}

		if device.System.Device != "" {
			builder.WriteString(indent + "Device: " + device.System.Device + "\n")
		}

	}

	if device.Subsystem != nil && device.Subsystem.Vendor != "" && device.Subsystem.Device != "" {
		builder.WriteString(indent + "Subsystem: " + device.Subsystem.Vendor + ":" + device.Subsystem.Device + "\n")
	}

	return builder.String()

}

func (view *Devices) render_drive(drive structs.Drive, indent string) string {

	var builder strings.Builder

	builder.WriteString(indent + "Name: " + drive.Name + "\n")

	if drive.Mountpoint != "" {
		builder.WriteString(indent + "Mountpoint: " + drive.Mountpoint + "\n")
	}

	if drive.Type != "" {
		builder.WriteString(indent + "Type: " + drive.Type + "\n")
	}

	if drive.Size > 0 {
		builder.WriteString(indent + "Size: " + utils_fmt.FormatBytes(drive.Size) + "\n")
	}

	if drive.Used > 0 {
		builder.WriteString(indent + "Used: " + utils_fmt.FormatBytes(drive.Used) + "\n")
	}

	if drive.Free > 0 {
		builder.WriteString(indent + "Free: " + utils_fmt.FormatBytes(drive.Free) + "\n")
	}

	builder.WriteString("\n")

	return builder.String()

}

