package views

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/gui/controllers"
import "fmt"

type Welcome struct {
	*gtk.Box
	controller     *controllers.App
	progress_label *gtk.Label
	progress_bar   *gtk.ProgressBar
	on_collected   func()
}

func NewWelcome(controller *controllers.App) *Welcome {

	box := gtk.NewBox(gtk.OrientationVertical, 12)
	box.SetMarginStart(16)
	box.SetMarginEnd(16)
	box.SetMarginTop(16)
	box.SetMarginBottom(16)

	title := gtk.NewLabel("")
	title.SetMarkup("<b>System Integrity Check</b>")
	title.SetHAlign(gtk.AlignStart)
	box.Append(title.AsPtr())

	description := gtk.NewLabel("Collect packages, boot configuration and device details.")
	description.SetWrap(true)
	description.SetXAlign(0.0)
	box.Append(description.AsPtr())

	progress_bar := gtk.NewProgressBar()
	progress_bar.SetHExpand(true)
	progress_bar.SetShowText(true)
	box.Append(progress_bar.AsPtr())

	progress_label := gtk.NewLabel("")
	progress_label.SetXAlign(0.0)
	box.Append(progress_label.AsPtr())

	button := gtk.NewButton("Collect System Report")
	button.SetHAlign(gtk.AlignStart)
	box.Append(button.AsPtr())

	view := &Welcome{
		Box:            box,
		controller:     controller,
		progress_label: progress_label,
		progress_bar:   progress_bar,
		on_collected:   nil,
	}

	button.OnClick(func() {

		button.SetSensitive(false)

		progress_bar.SetFraction(0.0)
		progress_label.SetText("Collecting ...")

		controller.StartCollect(
			func(progress controllers.Progress) {

				if controller.IsCollecting() == false {
					return
				}

				progress_bar.SetFraction(progress.Fraction)
				progress_label.SetText(fmt.Sprintf("Step %d/%d: %s", progress.Completed, progress.Total, progress.Step))

			},
			func() {

				progress_label.SetText("Finished")
				progress_bar.SetFraction(1.0)

				button.SetSensitive(true)

				if view.on_collected != nil {
					view.on_collected()
				}

			},
		)

	})

	return view

}

func (view *Welcome) SetOnCollected(callback func()) {
	view.on_collected = callback
}
