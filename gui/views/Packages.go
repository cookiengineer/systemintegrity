package views

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "github.com/cookiengineer/systemintegrity/structs"
import "fmt"
import "strings"

type Packages struct {
	*gtk.Box
	controller     *controllers.App
	description    *gtk.Label
	progress_label *gtk.Label
	report_content *gtk.Box
}

func NewPackages(controller *controllers.App) *Packages {

	box := gtk.NewBox(gtk.OrientationVertical, 12)
	box.SetMarginStart(16)
	box.SetMarginEnd(16)
	box.SetMarginTop(16)
	box.SetMarginBottom(16)

	title := gtk.NewLabel("")
	title.SetMarkup("<b>Packages</b>")
	title.SetHAlign(gtk.AlignStart)
	box.Append(title.AsPtr())

	description := gtk.NewLabel("No System Report collected yet.")
	description.SetXAlign(0.0)
	box.Append(description.AsPtr())

	refresh_button := gtk.NewButton("Re-verify Packages")
	refresh_button.SetHAlign(gtk.AlignStart)
	box.Append(refresh_button.AsPtr())

	progress_label := gtk.NewLabel("")
	progress_label.SetXAlign(0.0)
	box.Append(progress_label.AsPtr())

	report_wrapper := gtk.NewScrolledWindow()
	report_wrapper.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	report_wrapper.SetHExpand(true)
	report_wrapper.SetVExpand(true)

	report_content := gtk.NewBox(gtk.OrientationVertical, 4)
	report_content.SetMarginStart(8)
	report_content.SetMarginEnd(8)

	report_wrapper.SetChild(report_content.AsPtr())
	box.Append(report_wrapper.AsPtr())

	view := &Packages{
		Box:            box,
		controller:     controller,
		description:    description,
		progress_label: progress_label,
		report_content: report_content,
	}

	refresh_button.OnClick(func() {

		refresh_button.SetSensitive(false)
		progress_label.SetText("Verifying ...")

		controller.StartVerification(
			func(step string) {

				if controller.IsVerifying() == false {
					return
				}

				progress_label.SetText("Verifying: " + step)

			},
			func() {

				progress_label.SetText("Verification finished.")
				refresh_button.SetSensitive(true)

				view.Refresh(controller.System())

			},
		)

	})

	return view

}

func (view *Packages) Refresh(system *structs.System) {

	packages := len(system.Packages)
	affected := len(system.Verifications)

	files := 0

	for v := 0; v < len(system.Verifications); v++ {
		files += len(system.Verifications[v].Files)
	}

	view.description.SetText(fmt.Sprintf("%d of %d packages affected with %d changed files", affected, packages, files))
	view.report_content.Clear()

	if affected > 0 {

		for v := 0; v < len(system.Verifications); v++ {
			view.report_content.Append(view.render_package_verification(system.Verifications[v]).AsPtr())
		}

	} else {

		empty_label := gtk.NewLabel("No modified packages detected.")
		empty_label.SetXAlign(0.0)
		view.report_content.Append(empty_label.AsPtr())

	}

}

func (view *Packages) render_package_verification(verification structs.PackageVerification) *gtk.Expander {

	version := ""

	if verification.Version.IsValid() {
		version = " " + strings.TrimPrefix(verification.Version.String(), "0:")
	}

	label := fmt.Sprintf("%s%s (%d files)", verification.Name, version, len(verification.Files))

	wrapper := gtk.NewExpander(label)

	box := gtk.NewBox(gtk.OrientationVertical, 2)
	box.SetMarginStart(16)
	box.SetMarginBottom(4)

	for f := 0; f < len(verification.Files); f++ {

		file := verification.Files[f]
		line := gtk.NewLabel(file.Path + " - " + file.Reason)
		line.SetXAlign(0.0)
		line.SetWrap(true)

		box.Append(line.AsPtr())

	}

	wrapper.SetChild(box.AsPtr())

	return wrapper

}
